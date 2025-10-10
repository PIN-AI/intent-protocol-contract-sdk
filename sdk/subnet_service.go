package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	subnetcontract "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnet"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// SubnetService 封装 Subnet 合约的常用调用。
type SubnetService struct {
	contract *subnetcontract.Subnet
	txMgr    *txmgr.Manager
}

// NewSubnetService 使用已绑定的 Subnet 合约实例与 TxManager 构造服务。
func NewSubnetService(contract *subnetcontract.Subnet, txm *txmgr.Manager) *SubnetService {
	return &SubnetService{contract: contract, txMgr: txm}
}

// NewSubnetServiceByAddress 通过合约地址绑定 Subnet 并返回服务实例。
func NewSubnetServiceByAddress(backend bind.ContractBackend, addr common.Address, txm *txmgr.Manager) (*SubnetService, error) {
	contract, err := subnetcontract.NewSubnet(addr, backend)
	if err != nil {
		return nil, err
	}
	return &SubnetService{contract: contract, txMgr: txm}, nil
}

// RegisterParticipantParams 描述 RegisterParticipant 调用的可配置项。
type RegisterParticipantParams struct {
	Domain      string
	Endpoint    string
	MetadataURI string
	Value       *big.Int // 可选：注册所需的原生代币数量（如有）
}

func (s *SubnetService) registerParticipant(ctx context.Context, participantType ParticipantType, params RegisterParticipantParams) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		if params.Value != nil {
			opts.Value = params.Value
		}
		return s.contract.RegisterParticipant(opts, uint8(participantType), params.Domain, params.Endpoint, params.MetadataURI)
	})
}

// RegisterValidator 使用 RegisterParticipant 注册验证者。
func (s *SubnetService) RegisterValidator(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantValidator, params)
}

// RegisterAgent 使用 RegisterParticipant 注册代理。
func (s *SubnetService) RegisterAgent(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantAgent, params)
}

// RegisterMatcher 使用 RegisterParticipant 注册匹配器。
func (s *SubnetService) RegisterMatcher(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantMatcher, params)
}

// RegisterParticipantERC20Params 描述 RegisterParticipantERC20 的输入。
type RegisterParticipantERC20Params struct {
	Domain      string
	Endpoint    string
	MetadataURI string
	Amount      *big.Int
}

func (s *SubnetService) registerParticipantERC20(ctx context.Context, participantType ParticipantType, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	if params.Amount == nil {
		return nil, errors.New("subnet: amount is required for ERC20 registration")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RegisterParticipantERC20(opts, uint8(participantType), params.Amount, params.Domain, params.Endpoint, params.MetadataURI)
	})
}

// RegisterValidatorERC20 通过 ERC20 质押注册验证者。
func (s *SubnetService) RegisterValidatorERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantValidator, params)
}

// RegisterAgentERC20 通过 ERC20 质押注册代理。
func (s *SubnetService) RegisterAgentERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantAgent, params)
}

// RegisterMatcherERC20 通过 ERC20 质押注册匹配器。
func (s *SubnetService) RegisterMatcherERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantMatcher, params)
}

// ListActiveParticipants 返回指定类型的活跃参与者列表。
func (s *SubnetService) ListActiveParticipants(ctx context.Context, participantType ParticipantType) ([]subnetcontract.DataStructuresParticipantInfo, error) {
	return s.contract.ListActiveParticipants(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// IsActiveParticipant 查询指定地址在某参与者类型下是否活跃。
func (s *SubnetService) IsActiveParticipant(ctx context.Context, addr common.Address, participantType ParticipantType) (bool, error) {
	return s.contract.IsActiveParticipant(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantInfo 返回参与者详情。
func (s *SubnetService) GetParticipantInfo(ctx context.Context, addr common.Address, participantType ParticipantType) (subnetcontract.DataStructuresParticipantInfo, error) {
	return s.contract.GetParticipantInfo(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantStakeInfo 返回参与者在子网中的质押详情。
func (s *SubnetService) GetParticipantStakeInfo(ctx context.Context, addr common.Address, participantType ParticipantType) (subnetcontract.DataStructuresStakeInfo, error) {
	return s.contract.GetParticipantStakeInfo(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantCount 返回指定类型的参与者数量。
func (s *SubnetService) GetParticipantCount(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.GetParticipantCount(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// ======================== 只读方法：角色与权限 ========================

// AdminRole 返回管理员角色哈希。
func (s *SubnetService) AdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.ADMINROLE(&bind.CallOpts{Context: ctx})
}

// DefaultAdminRole 返回默认管理员角色哈希。
func (s *SubnetService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin 返回指定角色的管理员角色。
func (s *SubnetService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole 检查账户是否拥有指定角色。
func (s *SubnetService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== 只读方法：子网配置查询 ========================

// GetSubnetInfo 返回子网的完整信息。
func (s *SubnetService) GetSubnetInfo(ctx context.Context) (subnetcontract.DataStructuresSubnetInfo, error) {
	return s.contract.GetSubnetInfo(&bind.CallOpts{Context: ctx})
}

// GetBidFrequencyLimit 返回子网的 bid 频率限制。
func (s *SubnetService) GetBidFrequencyLimit(ctx context.Context) (*big.Int, error) {
	return s.contract.GetBidFrequencyLimit(&bind.CallOpts{Context: ctx})
}

// GetCheckpointPolicy 返回子网的 checkpoint 策略配置。
func (s *SubnetService) GetCheckpointPolicy(ctx context.Context) (subnetcontract.DataStructuresCheckpointPolicy, error) {
	return s.contract.GetCheckpointPolicy(&bind.CallOpts{Context: ctx})
}

// GetSignatureThreshold 返回子网的签名阈值配置。
func (s *SubnetService) GetSignatureThreshold(ctx context.Context) (subnetcontract.DataStructuresSignatureThreshold, error) {
	return s.contract.GetSignatureThreshold(&bind.CallOpts{Context: ctx})
}

// GetStakeGovernanceConfig 返回子网的质押治理配置。
func (s *SubnetService) GetStakeGovernanceConfig(ctx context.Context) (subnetcontract.DataStructuresStakeGovernanceConfig, error) {
	return s.contract.GetStakeGovernanceConfig(&bind.CallOpts{Context: ctx})
}

// GetSlashingRates 返回惩罚比率列表。
func (s *SubnetService) GetSlashingRates(ctx context.Context) ([]*big.Int, error) {
	return s.contract.GetSlashingRates(&bind.CallOpts{Context: ctx})
}

// GetUnstakeLockPeriod 返回解锁期。
func (s *SubnetService) GetUnstakeLockPeriod(ctx context.Context) (*big.Int, error) {
	return s.contract.GetUnstakeLockPeriod(&bind.CallOpts{Context: ctx})
}

// GetStakingToken 返回质押代币地址。
func (s *SubnetService) GetStakingToken(ctx context.Context) (common.Address, error) {
	return s.contract.GetStakingToken(&bind.CallOpts{Context: ctx})
}

// GetRequireKYC 返回子网是否要求 KYC。
func (s *SubnetService) GetRequireKYC(ctx context.Context) (bool, error) {
	return s.contract.GetRequireKYC(&bind.CallOpts{Context: ctx})
}

// GetMinimumStakeByType 返回指定参与者类型的最小质押要求。
func (s *SubnetService) GetMinimumStakeByType(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.GetMinimumStakeByType(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// ======================== 只读方法：参与者统计查询 ========================

// GetTotalParticipantCount 返回所有类型参与者的总数。
func (s *SubnetService) GetTotalParticipantCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetTotalParticipantCount(&bind.CallOpts{Context: ctx})
}

// GetActiveParticipantCount 返回指定类型的活跃参与者数量。
func (s *SubnetService) GetActiveParticipantCount(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.GetActiveParticipantCount(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// ActiveParticipantCounts 返回指定类型的活跃参与者计数（mapping 访问器）。
func (s *SubnetService) ActiveParticipantCounts(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.ActiveParticipantCounts(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// HasAvailableSlots 检查指定参与者类型是否还有可用槽位。
func (s *SubnetService) HasAvailableSlots(ctx context.Context, participantType ParticipantType) (bool, error) {
	return s.contract.HasAvailableSlots(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// Paused 返回子网是否处于暂停状态。
func (s *SubnetService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface 检查合约是否支持指定接口（ERC165）。
func (s *SubnetService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ParticipantListByType 返回指定类型参与者列表中指定索引的地址（mapping 访问器）。
func (s *SubnetService) ParticipantListByType(ctx context.Context, participantType ParticipantType, index *big.Int) (common.Address, error) {
	return s.contract.ParticipantListByType(&bind.CallOpts{Context: ctx}, uint8(participantType), index)
}

// ======================== 写入方法：参与者管理 ========================

// ApproveParticipant 批准参与者注册（需要 ADMIN_ROLE）。
func (s *SubnetService) ApproveParticipant(ctx context.Context, participantAddr common.Address, participantType ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ApproveParticipant(opts, participantAddr, uint8(participantType))
	})
}

// RejectParticipant 拒绝参与者注册（需要 ADMIN_ROLE）。
func (s *SubnetService) RejectParticipant(ctx context.Context, participantAddr common.Address, participantType ParticipantType, reason string) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RejectParticipant(opts, participantAddr, uint8(participantType), reason)
	})
}

// SuspendParticipant 暂停参与者（需要 ADMIN_ROLE）。
func (s *SubnetService) SuspendParticipant(ctx context.Context, participantType ParticipantType, reason string) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SuspendParticipant(opts, uint8(participantType), reason)
	})
}

// ResumeParticipant 恢复被暂停的参与者（需要 ADMIN_ROLE）。
func (s *SubnetService) ResumeParticipant(ctx context.Context, participantAddr common.Address, participantType ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ResumeParticipant(opts, participantAddr, uint8(participantType))
	})
}

// UpdateParticipantActivity 更新参与者活动时间（需要 ADMIN_ROLE）。
func (s *SubnetService) UpdateParticipantActivity(ctx context.Context, participantType ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.UpdateParticipantActivity(opts, uint8(participantType))
	})
}

// ======================== 写入方法：角色管理 ========================

// GrantRole 授予账户指定角色（需要角色管理员权限）。
func (s *SubnetService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole 撤销账户的指定角色（需要角色管理员权限）。
func (s *SubnetService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole 放弃自己的指定角色。
func (s *SubnetService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// ======================== 写入方法：子网管理 ========================

// Pause 暂停子网（需要 ADMIN_ROLE）。
func (s *SubnetService) Pause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Pause(opts)
	})
}

// Unpause 解除子网暂停（需要 ADMIN_ROLE）。
func (s *SubnetService) Unpause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Unpause(opts)
	})
}

// Deprecate 标记子网为废弃状态（需要 ADMIN_ROLE）。
func (s *SubnetService) Deprecate(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Deprecate(opts)
	})
}

// SetStakeConfig 设置子网的质押配置（需要 ADMIN_ROLE）。
func (s *SubnetService) SetStakeConfig(ctx context.Context, stakeConfig subnetcontract.DataStructuresStakeGovernanceConfig) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetStakeConfig(opts, stakeConfig)
	})
}

// Initialize 初始化子网合约（仅在部署后调用一次）。
func (s *SubnetService) Initialize(ctx context.Context, subnetInfo subnetcontract.DataStructuresSubnetInfo, factory common.Address, stakingManager common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, subnetInfo, factory, stakingManager)
	})
}
