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

// ErrGasCeilExceeded 在估算 gas 超过配置上限时返回。
type ErrGasCeilExceeded struct {
	Raw        uint64  // 原始估算值
	Adjusted   uint64  // 应用乘数后的值
	Multiplier float64 // 使用的乘数
	Ceil       uint64  // 配置的上限
}

func (e *ErrGasCeilExceeded) Error() string {
	return fmt.Sprintf("txmgr: estimated gas %d (raw: %d, multiplier: %.2fx) exceeds ceiling %d",
		e.Adjusted, e.Raw, e.Multiplier, e.Ceil)
}

// Config 控制 TxManager 的行为。
type Config struct {
	Use1559            bool
	GasLimit           uint64
	GasLimitMultiplier float64
	GasCeil            uint64 // gas 估算上限，超过则拒绝发送（0=不限制）
	MaxFeePerGas       *big.Int
	MaxPriorityFeeCap  *big.Int
	NonceSource        string // "pending" (default) or "latest"
	ReplaceStuck       bool
	ReplaceAfter       time.Duration
	BumpPercent        float64
	NoSend             bool
}

// DefaultConfig 返回默认配置。
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

// Manager 负责交易签名、gas 参数设置与可选的替换逻辑。
type Manager struct {
	client  *ethclient.Client
	signer  signer.Signer
	chainID *big.Int
	cfg     Config
	stopCh  chan struct{}
	wg      sync.WaitGroup
}

// New 创建 Manager。调用方需要在不使用时调用 Close。
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

// Close 释放内部 goroutine。
func (m *Manager) Close() {
	select {
	case <-m.stopCh:
		// already closed
	default:
		close(m.stopCh)
	}
	m.wg.Wait()
}

// TxExecutor 是高层调用传入的执行函数。
type TxExecutor func(opts *bind.TransactOpts) (*types.Transaction, error)

// Send 构造 TransactOpts，执行并可选地触发交易替换。
func (m *Manager) Send(ctx context.Context, exec TxExecutor) (*types.Transaction, error) {
	if exec == nil {
		return nil, errors.New("txmgr: nil executor")
	}
	opts, err := m.prepareOpts(ctx)
	if err != nil {
		return nil, err
	}

	// 如果未设置固定 GasLimit 且配置了乘数，先估算再应用乘数
	if m.cfg.GasLimit == 0 && m.cfg.GasLimitMultiplier > 1.0 {
		estimateOpts := *opts
		estimateOpts.NoSend = true
		estimateTx, estimateErr := exec(&estimateOpts)
		if estimateErr == nil && estimateTx != nil && estimateTx.Gas() > 0 {
			// 应用 GasLimitMultiplier 安全余量
			estimatedGas := estimateTx.Gas()
			gasWithBuffer := uint64(float64(estimatedGas) * m.cfg.GasLimitMultiplier)

			// 检查是否超过 GasCeil 上限
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
		// 如果估算失败，继续使用原 opts（让 abigen 自己估算）
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
