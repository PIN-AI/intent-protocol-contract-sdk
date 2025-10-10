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

// ======================== 只读方法：角色与权限 ========================

// DefaultAdminRole 返回默认管理员角色哈希。
func (s *StakingService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GovernanceRole 返回治理角色哈希。
func (s *StakingService) GovernanceRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GOVERNANCEROLE(&bind.CallOpts{Context: ctx})
}

// SlasherRole 返回惩罚者角色哈希。
func (s *StakingService) SlasherRole(ctx context.Context) ([32]byte, error) {
	return s.contract.SLASHERROLE(&bind.CallOpts{Context: ctx})
}

// StakingAdminRole 返回质押管理员角色哈希。
func (s *StakingService) StakingAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.STAKINGADMINROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin 返回指定角色的管理员角色。
func (s *StakingService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole 检查账户是否拥有指定角色。
func (s *StakingService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== 只读方法：配置查询 ========================

// GetStakingToken 返回质押代币地址。
func (s *StakingService) GetStakingToken(ctx context.Context) (common.Address, error) {
	return s.contract.GetStakingToken(&bind.CallOpts{Context: ctx})
}

// Paused 返回 StakingManager 是否处于暂停状态。
func (s *StakingService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface 检查合约是否支持指定接口（ERC165）。
func (s *StakingService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ======================== 写入方法：角色管理 ========================

// GrantRole 授予账户指定角色（需要角色管理员权限）。
func (s *StakingService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole 撤销账户的指定角色（需要角色管理员权限）。
func (s *StakingService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole 放弃自己的指定角色。
func (s *StakingService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// SetGovernanceRole 设置治理角色账户（需要 DEFAULT_ADMIN_ROLE）。
func (s *StakingService) SetGovernanceRole(ctx context.Context, governance common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetGovernanceRole(opts, governance)
	})
}

// SetSlasherRole 设置惩罚者角色账户（需要 DEFAULT_ADMIN_ROLE）。
func (s *StakingService) SetSlasherRole(ctx context.Context, slasher common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetSlasherRole(opts, slasher)
	})
}

// ======================== 写入方法：配置管理 ========================

// SetStakingToken 设置质押代币地址（需要 GOVERNANCE_ROLE）。
func (s *StakingService) SetStakingToken(ctx context.Context, token common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetStakingToken(opts, token)
	})
}

// SetUnlockPeriod 设置解锁期（需要 GOVERNANCE_ROLE）。
func (s *StakingService) SetUnlockPeriod(ctx context.Context, period *big.Int) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetUnlockPeriod(opts, period)
	})
}

// ======================== 写入方法：紧急控制与惩罚 ========================

// Pause 暂停 StakingManager（需要 GOVERNANCE_ROLE）。
func (s *StakingService) Pause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Pause(opts)
	})
}

// Unpause 解除 StakingManager 暂停（需要 GOVERNANCE_ROLE）。
func (s *StakingService) Unpause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Unpause(opts)
	})
}

// Slash 执行质押惩罚（需要 SLASHER_ROLE）。
func (s *StakingService) Slash(ctx context.Context, user common.Address, token common.Address, role ParticipantType, amount *big.Int, reason string) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	if amount == nil {
		return nil, errors.New("staking: amount is nil")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Slash(opts, user, token, uint8(role), amount, reason)
	})
}

// EmergencyWithdraw 紧急提取资金（需要 GOVERNANCE_ROLE）。
func (s *StakingService) EmergencyWithdraw(ctx context.Context, recipient common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	if amount == nil {
		return nil, errors.New("staking: amount is nil")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyWithdraw(opts, recipient, token, amount)
	})
}

// Initialize 初始化 StakingManager 合约（仅在部署后调用一次）。
func (s *StakingService) Initialize(ctx context.Context, admin common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, admin)
	})
}
