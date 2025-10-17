package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
)

// ErrGasCeilExceeded is returned when estimated gas exceeds the configured ceiling.
type ErrGasCeilExceeded struct {
	Raw        uint64  // Raw estimated value
	Adjusted   uint64  // Value after applying multiplier
	Multiplier float64 // Multiplier used
	Ceil       uint64  // Configured ceiling
}

func (e *ErrGasCeilExceeded) Error() string {
	return fmt.Sprintf("txmgr: estimated gas %d (raw: %d, multiplier: %.2fx) exceeds ceiling %d",
		e.Adjusted, e.Raw, e.Multiplier, e.Ceil)
}

// Config controls the behavior of TxManager.
type Config struct {
	Use1559            bool
	GasLimit           uint64
	GasLimitMultiplier float64
	GasCeil            uint64 // Gas estimation ceiling; rejects if exceeded (0 = no limit)
	MaxFeePerGas       *big.Int
	MaxPriorityFeeCap  *big.Int
	NonceSource        string // "pending" (default) or "latest"
	ReplaceStuck       bool
	ReplaceAfter       time.Duration
	BumpPercent        float64
	NoSend             bool
}

// DefaultConfig returns the default configuration.
func DefaultConfig() Config {
	return Config{
		Use1559:            true,
		GasLimitMultiplier: 1.5,
		NonceSource:        "pending",
		ReplaceStuck:       false,
		ReplaceAfter:       45 * time.Second,
		BumpPercent:        12.5,
		NoSend:             false,
	}
}

// Manager handles transaction signing, gas parameter configuration, and optional replacement logic.
type Manager struct {
	client  *ethclient.Client
	signer  signer.Signer
	chainID *big.Int
	cfg     Config
	stopCh  chan struct{}
	wg      sync.WaitGroup
}

// New creates a Manager. The caller must call Close when done.
func New(client *ethclient.Client, chainID *big.Int, signer signer.Signer, cfg Config) *Manager {
	m := &Manager{
		client:  client,
		signer:  signer,
		chainID: new(big.Int).Set(chainID),
		cfg:     cfg,
		stopCh:  make(chan struct{}),
	}
	return m
}

// Close releases internal goroutines.
func (m *Manager) Close() {
	select {
	case <-m.stopCh:
		// already closed
	default:
		close(m.stopCh)
	}
	m.wg.Wait()
}

// TxExecutor is the execution function passed by high-level callers.
type TxExecutor func(opts *bind.TransactOpts) (*types.Transaction, error)

// Send constructs TransactOpts, executes, and optionally triggers transaction replacement.
func (m *Manager) Send(ctx context.Context, exec TxExecutor) (*types.Transaction, error) {
	if exec == nil {
		return nil, errors.New("txmgr: nil executor")
	}
	opts, err := m.prepareOpts(ctx)
	if err != nil {
		return nil, err
	}

	// If no fixed GasLimit is set and a multiplier is configured, estimate first then apply multiplier
	if m.cfg.GasLimit == 0 && m.cfg.GasLimitMultiplier > 1.0 {
		estimateOpts := *opts
		estimateOpts.NoSend = true
		estimateTx, estimateErr := exec(&estimateOpts)
		if estimateErr == nil && estimateTx != nil && estimateTx.Gas() > 0 {
			// Apply GasLimitMultiplier safety margin
			estimatedGas := estimateTx.Gas()
			gasWithBuffer := uint64(float64(estimatedGas) * m.cfg.GasLimitMultiplier)

			// Check if GasCeil ceiling is exceeded
			if m.cfg.GasCeil > 0 && gasWithBuffer > m.cfg.GasCeil {
				return nil, &ErrGasCeilExceeded{
					Raw:        estimatedGas,
					Adjusted:   gasWithBuffer,
					Multiplier: m.cfg.GasLimitMultiplier,
					Ceil:       m.cfg.GasCeil,
				}
			}

			opts.GasLimit = gasWithBuffer
		}
		// If estimation fails, continue with original opts (let abigen estimate)
	}

	if m.cfg.NoSend {
		opts.NoSend = true
	}
	tx, err := exec(opts)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, errors.New("txmgr: executor returned nil transaction")
	}
	if m.cfg.NoSend || !m.cfg.ReplaceStuck || m.cfg.ReplaceAfter <= 0 || m.cfg.BumpPercent <= 0 {
		return tx, nil
	}
	m.scheduleReplacement(ctx, tx)
	return tx, nil
}

func (m *Manager) prepareOpts(ctx context.Context) (*bind.TransactOpts, error) {
	from := m.signer.Address()
	opts := &bind.TransactOpts{
		From:    from,
		Context: ctx,
		Signer: func(addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if addr != from {
				return nil, signer.ErrInvalidSignerAddress
			}
			return m.signer.SignTransaction(ctx, tx, m.chainID)
		},
	}

	var nonce uint64
	var err error
	switch strings.ToLower(m.cfg.NonceSource) {
	case "", "pending":
		nonce, err = m.client.PendingNonceAt(ctx, from)
	case "latest":
		nonce, err = m.client.NonceAt(ctx, from, nil)
	default:
		err = errors.New("txmgr: invalid nonce source")
	}
	if err != nil {
		return nil, err
	}
	opts.Nonce = new(big.Int).SetUint64(nonce)

	if m.cfg.GasLimit > 0 {
		opts.GasLimit = m.cfg.GasLimit
	}

	if m.cfg.Use1559 {
		if m.cfg.MaxFeePerGas != nil {
			opts.GasFeeCap = new(big.Int).Set(m.cfg.MaxFeePerGas)
		} else {
			feeCap, err := m.suggestFeeCap(ctx)
			if err != nil {
				return nil, err
			}
			opts.GasFeeCap = feeCap
		}
		if m.cfg.MaxPriorityFeeCap != nil {
			opts.GasTipCap = new(big.Int).Set(m.cfg.MaxPriorityFeeCap)
		} else {
			tip, err := m.client.SuggestGasTipCap(ctx)
			if err != nil {
				return nil, err
			}
			opts.GasTipCap = tip
		}
	} else {
		if m.cfg.MaxFeePerGas != nil {
			opts.GasPrice = new(big.Int).Set(m.cfg.MaxFeePerGas)
		} else {
			price, err := m.client.SuggestGasPrice(ctx)
			if err != nil {
				return nil, err
			}
			opts.GasPrice = price
		}
	}
	return opts, nil
}

func (m *Manager) suggestFeeCap(ctx context.Context) (*big.Int, error) {
	tip, err := m.client.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, err
	}
	base, err := m.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	feeCap := new(big.Int).Add(base, new(big.Int).Mul(tip, big.NewInt(2)))
	return feeCap, nil
}

func (m *Manager) scheduleReplacement(parentCtx context.Context, original *types.Transaction) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		timer := time.NewTimer(m.cfg.ReplaceAfter)
		defer timer.Stop()
		select {
		case <-timer.C:
		case <-parentCtx.Done():
			return
		case <-m.stopCh:
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := m.client.TransactionReceipt(ctx, original.Hash())
		if err == nil {
			return
		}
		if !errors.Is(err, ethereum.NotFound) {
			return
		}
		if err := m.bumpAndSend(original); err != nil {
			return
		}
	}()
}

func (m *Manager) bumpAndSend(original *types.Transaction) error {
	newTx, err := m.bumpTransaction(original)
	if err != nil {
		return err
	}
	signed, err := m.signer.SignTransaction(context.Background(), newTx, m.chainID)
	if err != nil {
		return err
	}
	return m.client.SendTransaction(context.Background(), signed)
}

func (m *Manager) bumpTransaction(tx *types.Transaction) (*types.Transaction, error) {
	bump := m.cfg.BumpPercent
	if bump <= 0 {
		return nil, errors.New("txmgr: invalid bump percent")
	}
	switch tx.Type() {
	case types.DynamicFeeTxType:
		feeCap := bumpBig(tx.GasFeeCap(), bump)
		tipCap := bumpBig(tx.GasTipCap(), bump)
		return types.NewTx(&types.DynamicFeeTx{
			ChainID:    tx.ChainId(),
			Nonce:      tx.Nonce(),
			GasTipCap:  tipCap,
			GasFeeCap:  feeCap,
			Gas:        tx.Gas(),
			To:         tx.To(),
			Value:      tx.Value(),
			Data:       tx.Data(),
			AccessList: tx.AccessList(),
		}), nil
	case types.AccessListTxType:
		price := bumpBig(tx.GasPrice(), bump)
		return types.NewTx(&types.AccessListTx{
			ChainID:    tx.ChainId(),
			Nonce:      tx.Nonce(),
			GasPrice:   price,
			Gas:        tx.Gas(),
			To:         tx.To(),
			Value:      tx.Value(),
			Data:       tx.Data(),
			AccessList: tx.AccessList(),
		}), nil
	default:
		price := bumpBig(tx.GasPrice(), bump)
		return types.NewTx(&types.LegacyTx{
			Nonce:    tx.Nonce(),
			GasPrice: price,
			Gas:      tx.Gas(),
			To:       tx.To(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		}), nil
	}
}

func bumpBig(value *big.Int, percent float64) *big.Int {
	if value == nil {
		return nil
	}
	if percent <= 0 {
		return new(big.Int).Set(value)
	}
	base := new(big.Float).SetPrec(256).SetInt(value)
	multiplier := big.NewFloat(1 + percent/100)
	multiplier.SetPrec(256)
	base.Mul(base, multiplier)
	base.Add(base, big.NewFloat(0.5))
	result := new(big.Int)
	base.Int(result)
	if result.Cmp(value) <= 0 {
		return new(big.Int).Add(value, big.NewInt(1))
	}
	return result
}
