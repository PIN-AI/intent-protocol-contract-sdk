package sdk

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	subnetcontract "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnet"
	subnetfactory "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnetfactory"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// SubnetInfoSummary 汇总 SubnetFactory.getSubnetInfo 的返回值。
type SubnetInfoSummary struct {
	Info           subnetfactory.DataStructuresSubnetInfo
	ValidatorCount *big.Int
	AgentCount     *big.Int
	Status         uint8
}

// SubnetFactoryService 封装 SubnetFactory 读写操作。
type SubnetFactoryService struct {
	backend  bind.ContractBackend
	txMgr    *txmgr.Manager
	contract *subnetfactory.SubnetFactory
	signer   signer.Signer
	chainID  *big.Int
}

// NewSubnetFactoryService 构造服务。
func NewSubnetFactoryService(backend bind.ContractBackend, txm *txmgr.Manager, contract *subnetfactory.SubnetFactory, sig signer.Signer, chainID *big.Int) *SubnetFactoryService {
	return &SubnetFactoryService{
		backend:  backend,
		txMgr:    txm,
		contract: contract,
		signer:   sig,
		chainID:  new(big.Int).Set(chainID),
	}
}

// GetSubnetContract 返回子网合约地址。
func (s *SubnetFactoryService) GetSubnetContract(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.SubnetContracts(&bind.CallOpts{Context: ctx}, subnetID)
}

// IsSubnetActive 查询子网是否处于 ACTIVE 状态。
func (s *SubnetFactoryService) IsSubnetActive(ctx context.Context, subnetID [32]byte) (bool, error) {
	return s.contract.IsSubnetActive(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetSubnetInfo 返回子网详细信息与统计。
func (s *SubnetFactoryService) GetSubnetInfo(ctx context.Context, subnetID [32]byte) (SubnetInfoSummary, error) {
	resp, err := s.contract.GetSubnetInfo(&bind.CallOpts{Context: ctx}, subnetID)
	if err != nil {
		return SubnetInfoSummary{}, err
	}
	return SubnetInfoSummary{
		Info:           resp.Info,
		ValidatorCount: resp.ValidatorCount,
		AgentCount:     resp.AgentCount,
		Status:         resp.SubnetStatus,
	}, nil
}

// TotalCreated 返回累计创建的子网数量。
func (s *SubnetFactoryService) TotalCreated(ctx context.Context) (*big.Int, error) {
	return s.contract.TotalCreated(&bind.CallOpts{Context: ctx})
}

// EnumerateSubnets 返回当前所有子网 ID。
func (s *SubnetFactoryService) EnumerateSubnets(ctx context.Context) ([][32]byte, error) {
	total, err := s.contract.TotalCreated(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	count := total.Uint64()
	ids := make([][32]byte, 0, count)
	for i := uint64(0); i < count; i++ {
		id, err := s.contract.AllSubnets(&bind.CallOpts{Context: ctx}, new(big.Int).SetUint64(i))
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

// GetActiveSubnetCount 返回活跃子网数量。
func (s *SubnetFactoryService) GetActiveSubnetCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetActiveSubnetCount(&bind.CallOpts{Context: ctx})
}

// GetCreationStats 返回 totalCreated（兼容旧接口）。
func (s *SubnetFactoryService) GetCreationStats(ctx context.Context) (*big.Int, error) {
	return s.contract.GetCreationStats(&bind.CallOpts{Context: ctx})
}

// PredictSubnetAddress 使用 clone determinism 预测地址。
func (s *SubnetFactoryService) PredictSubnetAddress(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.PredictSubnetAddress(&bind.CallOpts{Context: ctx}, subnetID)
}

// BindSubnet 返回 Subnet 合约实例，便于查询参与者等信息。
func (s *SubnetFactoryService) BindSubnet(subnetAddr common.Address) (*subnetcontract.Subnet, error) {
	return subnetcontract.NewSubnet(subnetAddr, s.backend)
}

// NewSubnetServiceByAddress 根据子网合约地址创建 SubnetService。
func (s *SubnetFactoryService) NewSubnetServiceByAddress(subnetAddr common.Address) (*SubnetService, error) {
	contract, err := s.BindSubnet(subnetAddr)
	if err != nil {
		return nil, err
	}
	return NewSubnetService(contract, s.txMgr), nil
}

// NewSubnetServiceByID 通过子网 ID 查询地址并创建 SubnetService。
func (s *SubnetFactoryService) NewSubnetServiceByID(ctx context.Context, subnetID [32]byte) (*SubnetService, error) {
	addr, err := s.GetSubnetContract(ctx, subnetID)
	if err != nil {
		return nil, err
	}
	return s.NewSubnetServiceByAddress(addr)
}

// PauseSubnet 调用 pauseSubnet（需要 GUARDIAN_ROLE）。
func (s *SubnetFactoryService) PauseSubnet(ctx context.Context, subnetID [32]byte, reason string) (*types.Transaction, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, errors.New("subnet: pause reason required")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.PauseSubnet(opts, subnetID, reason)
	})
}

// ResumeSubnet 调用 resumeSubnet。
func (s *SubnetFactoryService) ResumeSubnet(ctx context.Context, subnetID [32]byte) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ResumeSubnet(opts, subnetID)
	})
}

// DeprecateSubnet 调用 deprecateSubnet。
func (s *SubnetFactoryService) DeprecateSubnet(ctx context.Context, subnetID [32]byte, reason string) (*types.Transaction, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, errors.New("subnet: deprecate reason required")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.DeprecateSubnet(opts, subnetID, reason)
	})
}

// GetActiveParticipants 批量获取指定子网的活跃参与者信息。
// 返回每个子网的三种活跃参与者：VALIDATOR、AGENT、MATCHER。
// 返回数组长度始终等于输入长度，即使部分子网不存在也会返回空数组。
func (s *SubnetFactoryService) GetActiveParticipants(ctx context.Context, subnetIDs [][32]byte) ([]subnetfactory.DataStructuresSubnetParticipants, error) {
	return s.contract.GetActiveParticipants(&bind.CallOpts{Context: ctx}, subnetIDs)
}

// ======================== 只读方法：角色与权限 ========================

// DefaultAdminRole 返回默认管理员角色哈希。
func (s *SubnetFactoryService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GuardianRole 返回监护人角色哈希。
func (s *SubnetFactoryService) GuardianRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GUARDIANROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin 返回指定角色的管理员角色。
func (s *SubnetFactoryService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole 检查账户是否拥有指定角色。
func (s *SubnetFactoryService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== 只读方法：子网查询与过滤 ========================

// FindSubnetByName 通过规范名称查找子网 ID。
func (s *SubnetFactoryService) FindSubnetByName(ctx context.Context, canonicalName string) ([32]byte, error) {
	return s.contract.FindSubnetByName(&bind.CallOpts{Context: ctx}, canonicalName)
}

// GetAllSubnetInfo 返回所有子网的完整信息。
func (s *SubnetFactoryService) GetAllSubnetInfo(ctx context.Context) (
	infos []subnetfactory.DataStructuresSubnetInfo,
	validatorCounts []*big.Int,
	agentCounts []*big.Int,
	statuses []uint8,
	err error,
) {
	result, err := s.contract.GetAllSubnetInfo(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return result.Infos, result.ValidatorCounts, result.AgentCounts, result.Statuses, nil
}

// GetSubnetsByOwner 返回指定所有者拥有的所有子网 ID。
func (s *SubnetFactoryService) GetSubnetsByOwner(ctx context.Context, owner common.Address) ([][32]byte, error) {
	return s.contract.GetSubnetsByOwner(&bind.CallOpts{Context: ctx}, owner)
}

// GetSubnetsByStatus 返回指定状态的所有子网 ID。
func (s *SubnetFactoryService) GetSubnetsByStatus(ctx context.Context, status uint8) ([][32]byte, error) {
	return s.contract.GetSubnetsByStatus(&bind.CallOpts{Context: ctx}, status)
}

// GetTotalSubnetCount 返回子网总数（等同于 TotalCreated）。
func (s *SubnetFactoryService) GetTotalSubnetCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetTotalSubnetCount(&bind.CallOpts{Context: ctx})
}

// ListAllSubnets 返回所有子网 ID（等同于 EnumerateSubnets）。
func (s *SubnetFactoryService) ListAllSubnets(ctx context.Context) ([][32]byte, error) {
	return s.contract.ListAllSubnets(&bind.CallOpts{Context: ctx})
}

// SubnetExists 检查子网是否存在。
func (s *SubnetFactoryService) SubnetExists(ctx context.Context, subnetID [32]byte) (bool, error) {
	return s.contract.SubnetExists(&bind.CallOpts{Context: ctx}, subnetID)
}

// ======================== 只读方法：子网配置查询 ========================

// GetSubnetStakeConfig 返回子网的质押治理配置。
func (s *SubnetFactoryService) GetSubnetStakeConfig(ctx context.Context, subnetID [32]byte) (subnetfactory.DataStructuresStakeGovernanceConfig, error) {
	return s.contract.GetSubnetStakeConfig(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetSubnetThreshold 返回子网的签名阈值配置。
func (s *SubnetFactoryService) GetSubnetThreshold(ctx context.Context, subnetID [32]byte) (subnetfactory.DataStructuresSignatureThreshold, error) {
	return s.contract.GetSubnetThreshold(&bind.CallOpts{Context: ctx}, subnetID)
}

// ======================== 只读方法：全局配置查询 ========================

// GetMinStakeCreateSubnet 返回创建子网的最小质押要求。
func (s *SubnetFactoryService) GetMinStakeCreateSubnet(ctx context.Context) (*big.Int, error) {
	return s.contract.GetMinStakeCreateSubnet(&bind.CallOpts{Context: ctx})
}

// GetStakingManager 返回 StakingManager 合约地址。
func (s *SubnetFactoryService) GetStakingManager(ctx context.Context) (common.Address, error) {
	return s.contract.GetStakingManager(&bind.CallOpts{Context: ctx})
}

// GetSubnetImplementation 返回子网实现合约地址。
func (s *SubnetFactoryService) GetSubnetImplementation(ctx context.Context) (common.Address, error) {
	return s.contract.GetSubnetImplementation(&bind.CallOpts{Context: ctx})
}

// StakingManager 返回 StakingManager 合约地址（字段访问器）。
func (s *SubnetFactoryService) StakingManager(ctx context.Context) (common.Address, error) {
	return s.contract.StakingManager(&bind.CallOpts{Context: ctx})
}

// SubnetImplementation 返回子网实现合约地址（字段访问器）。
func (s *SubnetFactoryService) SubnetImplementation(ctx context.Context) (common.Address, error) {
	return s.contract.SubnetImplementation(&bind.CallOpts{Context: ctx})
}

// Paused 返回 SubnetFactory 是否处于暂停状态。
func (s *SubnetFactoryService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface 检查合约是否支持指定接口（ERC165）。
func (s *SubnetFactoryService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// AllSubnets 返回指定索引位置的子网 ID（用于遍历）。
func (s *SubnetFactoryService) AllSubnets(ctx context.Context, index *big.Int) ([32]byte, error) {
	return s.contract.AllSubnets(&bind.CallOpts{Context: ctx}, index)
}

// SubnetContracts 返回子网 ID 对应的合约地址（mapping 访问器）。
func (s *SubnetFactoryService) SubnetContracts(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.SubnetContracts(&bind.CallOpts{Context: ctx}, subnetID)
}

// ======================== 写入方法：子网创建 ========================

// CreateSubnet 创建新子网（需要满足最小质押要求）。
func (s *SubnetFactoryService) CreateSubnet(ctx context.Context, createInfo subnetfactory.DataStructuresCreateSubnetInfo) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.CreateSubnet(opts, createInfo)
	})
}

// ======================== 写入方法：紧急控制 ========================

// EmergencyPause 紧急暂停 SubnetFactory（需要 GUARDIAN_ROLE）。
func (s *SubnetFactoryService) EmergencyPause(ctx context.Context) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyPause(opts)
	})
}

// EmergencyUnpause 解除 SubnetFactory 紧急暂停（需要 GUARDIAN_ROLE）。
func (s *SubnetFactoryService) EmergencyUnpause(ctx context.Context) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyUnpause(opts)
	})
}

// ======================== 写入方法：角色管理 ========================

// GrantRole 授予账户指定角色（需要角色管理员权限）。
func (s *SubnetFactoryService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole 撤销账户的指定角色（需要角色管理员权限）。
func (s *SubnetFactoryService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole 放弃自己的指定角色。
func (s *SubnetFactoryService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// ======================== 写入方法：全局配置管理 ========================

// SetMinStakeCreateSubnet 设置创建子网的最小质押要求（需要 DEFAULT_ADMIN_ROLE）。
func (s *SubnetFactoryService) SetMinStakeCreateSubnet(ctx context.Context, amount *big.Int) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetMinStakeCreateSubnet(opts, amount)
	})
}

// SetStakingManager 设置 StakingManager 合约地址（需要 DEFAULT_ADMIN_ROLE）。
func (s *SubnetFactoryService) SetStakingManager(ctx context.Context, stakingManager common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetStakingManager(opts, stakingManager)
	})
}

// SetSubnetImplementation 设置子网实现合约地址（需要 DEFAULT_ADMIN_ROLE）。
func (s *SubnetFactoryService) SetSubnetImplementation(ctx context.Context, newImplementation common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetSubnetImplementation(opts, newImplementation)
	})
}

// Initialize 初始化 SubnetFactory 合约（仅在部署后调用一次）。
func (s *SubnetFactoryService) Initialize(ctx context.Context, guardian common.Address, subnetImplementation common.Address, stakingManager common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, guardian, subnetImplementation, stakingManager)
	})
}
