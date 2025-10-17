package sdk

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	checkpointmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/checkpointmanager"
	intentmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/intentmanager"
	stakingmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/stakingmanager"
	subnetcontract "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnet"
	subnetfactory "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnetfactory"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/addressbook"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// Config provides parameters required to initialize the Client.
type Config struct {
	RPCURL        string
	PrivateKeyHex string
	Network       string
	Addresses     *addressbook.Addresses
	Tx            *TxOptions
	Signer        signer.Signer
}

// TxOptions describes optional settings for TxManager.
type TxOptions struct {
	Use1559            *bool
	GasLimit           *uint64
	GasLimitMultiplier *float64
	GasCeil            *uint64 // Gas estimation ceiling; rejects if exceeded (0 = no limit)
	MaxFeePerGas       *big.Int
	MaxPriorityFeeCap  *big.Int
	NonceSource        *string
	ReplaceStuck       *bool
	ReplaceAfter       *time.Duration
	BumpPercent        *float64
	NoSend             *bool
}

// Client exposes high-level contract wrappers and low-level connections.
type Client struct {
	Backend   *ethclient.Client
	ChainID   *big.Int
	Network   addressbook.Network
	Addresses addressbook.Addresses
	Signer    signer.Signer
	TxManager *txmgr.Manager

	Intent            *IntentService
	Assignment        *AssignmentService
	Validation        *ValidationService
	SubnetFactory     *SubnetFactoryService
	StakingManager    *StakingService
	CheckpointManager *CheckpointService
}

// NewClient creates a complete SDK instance based on Config.
func NewClient(ctx context.Context, cfg Config) (*Client, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if strings.TrimSpace(cfg.RPCURL) == "" {
		return nil, errors.New("sdk: RPCURL is required")
	}
	backend, err := ethclient.DialContext(ctx, cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("sdk: dial rpc: %w", err)
	}
	chainID, err := backend.ChainID(ctx)
	if err != nil {
		return nil, fmt.Errorf("sdk: fetch chain id: %w", err)
	}

	var netName addressbook.Network
	if strings.TrimSpace(cfg.Network) != "" {
		netName, err = addressbook.NormalizeNetwork(cfg.Network)
		if err != nil {
			return nil, err
		}
	} else {
		netName, err = addressbook.FromChainID(chainID)
		if err != nil {
			return nil, err
		}
	}

	var signing signer.Signer
	if cfg.Signer != nil {
		signing = cfg.Signer
	} else if strings.TrimSpace(cfg.PrivateKeyHex) != "" {
		signing, err = signer.NewLocalSigner(cfg.PrivateKeyHex)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("sdk: signer or private key required")
	}

	addresses, err := addressbook.Resolve(netName, cfg.Addresses)
	if err != nil {
		return nil, err
	}

	txCfg := txmgr.DefaultConfig()
	applyTxOptions(&txCfg, cfg.Tx)
	txManager := txmgr.New(backend, chainID, signing, txCfg)

	intentContract, err := intentmanager.NewIntentManager(addresses.IntentManager, backend)
	if err != nil {
		txManager.Close()
		return nil, fmt.Errorf("sdk: bind intent manager: %w", err)
	}
	subnetFactoryContract, err := subnetfactory.NewSubnetFactory(addresses.SubnetFactory, backend)
	if err != nil {
		txManager.Close()
		return nil, fmt.Errorf("sdk: bind subnet factory: %w", err)
	}
	stakingContract, err := stakingmanager.NewStakingManager(addresses.StakingManager, backend)
	if err != nil {
		txManager.Close()
		return nil, fmt.Errorf("sdk: bind staking manager: %w", err)
	}
	checkpointContract, err := checkpointmanager.NewCheckpointManager(addresses.CheckpointManager, backend)
	if err != nil {
		txManager.Close()
		return nil, fmt.Errorf("sdk: bind checkpoint manager: %w", err)
	}

	stakingService := NewStakingService(stakingContract)
	stakingService.AttachTxManager(txManager)

	checkpointService := NewCheckpointService(checkpointContract, backend, txManager, signing, chainID, addresses.CheckpointManager)

	client := &Client{
		Backend:           backend,
		ChainID:           chainID,
		Network:           netName,
		Addresses:         addresses,
		Signer:            signing,
		TxManager:         txManager,
		Intent:            NewIntentService(backend, txManager, intentContract, signing, chainID, addresses.IntentManager),
		Assignment:        NewAssignmentService(backend, txManager, intentContract, signing, chainID, addresses.IntentManager),
		Validation:        NewValidationService(backend, txManager, intentContract, signing, chainID, addresses.IntentManager),
		SubnetFactory:     NewSubnetFactoryService(backend, txManager, subnetFactoryContract, signing, chainID),
		StakingManager:    stakingService,
		CheckpointManager: checkpointService,
	}
	return client, nil
}

// Close closes the TxManager (RPC connection is managed by the caller).
func (c *Client) Close() {
	if c == nil {
		return
	}
	if c.TxManager != nil {
		c.TxManager.Close()
	}
}

func applyTxOptions(base *txmgr.Config, opts *TxOptions) {
	if base == nil || opts == nil {
		return
	}
	if opts.Use1559 != nil {
		base.Use1559 = *opts.Use1559
	}
	if opts.GasLimit != nil {
		base.GasLimit = *opts.GasLimit
	}
	if opts.GasLimitMultiplier != nil && *opts.GasLimitMultiplier > 0 {
		base.GasLimitMultiplier = *opts.GasLimitMultiplier
	}
	if opts.GasCeil != nil {
		base.GasCeil = *opts.GasCeil
	}
	if opts.MaxFeePerGas != nil {
		base.MaxFeePerGas = new(big.Int).Set(opts.MaxFeePerGas)
	}
	if opts.MaxPriorityFeeCap != nil {
		base.MaxPriorityFeeCap = new(big.Int).Set(opts.MaxPriorityFeeCap)
	}
	if opts.NonceSource != nil && strings.TrimSpace(*opts.NonceSource) != "" {
		base.NonceSource = strings.ToLower(strings.TrimSpace(*opts.NonceSource))
	}
	if opts.ReplaceStuck != nil {
		base.ReplaceStuck = *opts.ReplaceStuck
	}
	if opts.ReplaceAfter != nil {
		base.ReplaceAfter = *opts.ReplaceAfter
	}
	if opts.BumpPercent != nil && *opts.BumpPercent > 0 {
		base.BumpPercent = *opts.BumpPercent
	}
	if opts.NoSend != nil {
		base.NoSend = *opts.NoSend
	}
}

// SubnetServiceByAddress creates a SubnetService directly by subnet contract address.
func (c *Client) SubnetServiceByAddress(addr common.Address) (*SubnetService, error) {
	if c == nil {
		return nil, errors.New("sdk: nil client")
	}
	contract, err := subnetcontract.NewSubnet(addr, c.Backend)
	if err != nil {
		return nil, err
	}
	return NewSubnetService(contract, c.TxManager), nil
}

// SubnetServiceByID queries the subnet address by ID and creates a SubnetService.
func (c *Client) SubnetServiceByID(ctx context.Context, subnetID [32]byte) (*SubnetService, error) {
	if c == nil {
		return nil, errors.New("sdk: nil client")
	}
	if c.SubnetFactory == nil {
		return nil, errors.New("sdk: subnet factory not initialised")
	}
	return c.SubnetFactory.NewSubnetServiceByID(ctx, subnetID)
}
