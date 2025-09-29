package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	stakingmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/stakingmanager"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// ParticipantType 与 DataStructures.ParticipantType 对齐。
type ParticipantType uint8

const (
	ParticipantValidator   ParticipantType = 0
	ParticipantAgent       ParticipantType = 1
	ParticipantMatcher     ParticipantType = 2
	ParticipantSubnetOwner ParticipantType = 3
)

// StakingService 提供 StakingManager 的常用封装。
type StakingService struct {
	contract *stakingmanager.StakingManager
	txMgr    *txmgr.Manager
}

// NewStakingService 创建服务。
func NewStakingService(contract *stakingmanager.StakingManager) *StakingService {
	return &StakingService{contract: contract}
}

// AttachTxManager 为需要写操作的场景附加 TxManager。
func (s *StakingService) AttachTxManager(txm *txmgr.Manager) {
	s.txMgr = txm
}

// GetStakeInfo 查询 stake 信息。
func (s *StakingService) GetStakeInfo(ctx context.Context, user, token common.Address, role ParticipantType) (stakingmanager.DataStructuresStakeInfo, error) {
	return s.contract.GetStakeInfo(&bind.CallOpts{Context: ctx}, user, token, uint8(role))
}

// GetUnlockPeriod 查询解锁期。
func (s *StakingService) GetUnlockPeriod(ctx context.Context) (*big.Int, error) {
	return s.contract.GetUnlockPeriod(&bind.CallOpts{Context: ctx})
}

// DepositStakeFor 调用 depositStakeFor。调用者需具备额度或发送 ETH。
func (s *StakingService) DepositStakeFor(ctx context.Context, user, token common.Address, role ParticipantType, amount *big.Int) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	if amount == nil {
		return nil, errors.New("staking: amount is nil")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		if token == ZeroAddress {
			opts.Value = amount
		}
		return s.contract.DepositStakeFor(opts, user, token, uint8(role), amount)
	})
}

// RequestUnstake 发起解押申请。
func (s *StakingService) RequestUnstake(ctx context.Context, token common.Address, role ParticipantType, amount *big.Int) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	if amount == nil {
		return nil, errors.New("staking: amount is nil")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RequestUnstake(opts, token, uint8(role), amount)
	})
}

// Withdraw 提取完成的解押金额。
func (s *StakingService) Withdraw(ctx context.Context, token common.Address, role ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Withdraw(opts, token, uint8(role))
	})
}
