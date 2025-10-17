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

// SubnetInfoSummary summarizes return values from SubnetFactory.getSubnetInfo.
type SubnetInfoSummary struct {
	Info           subnetfactory.DataStructuresSubnetInfo
	ValidatorCount *big.Int
	AgentCount     *big.Int
	Status         uint8
}

// SubnetFactoryService wraps SubnetFactory read and write operations.
type SubnetFactoryService struct {
	backend  bind.ContractBackend
	txMgr    *txmgr.Manager
	contract *subnetfactory.SubnetFactory
	signer   signer.Signer
	chainID  *big.Int
}

// NewSubnetFactoryService constructs the service.
func NewSubnetFactoryService(backend bind.ContractBackend, txm *txmgr.Manager, contract *subnetfactory.SubnetFactory, sig signer.Signer, chainID *big.Int) *SubnetFactoryService {
	return &SubnetFactoryService{
		backend:  backend,
		txMgr:    txm,
		contract: contract,
		signer:   sig,
		chainID:  new(big.Int).Set(chainID),
	}
}

// GetSubnetContract returns the subnet contract address.
func (s *SubnetFactoryService) GetSubnetContract(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.SubnetContracts(&bind.CallOpts{Context: ctx}, subnetID)
}

// IsSubnetActive queries whether the subnet is in ACTIVE status.
func (s *SubnetFactoryService) IsSubnetActive(ctx context.Context, subnetID [32]byte) (bool, error) {
	return s.contract.IsSubnetActive(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetSubnetInfo returns detailed subnet information and statistics.
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

// TotalCreated returns the total number of subnets created.
func (s *SubnetFactoryService) TotalCreated(ctx context.Context) (*big.Int, error) {
	return s.contract.TotalCreated(&bind.CallOpts{Context: ctx})
}

// EnumerateSubnets returns all current subnet IDs.
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

// GetActiveSubnetCount returns the number of active subnets.
func (s *SubnetFactoryService) GetActiveSubnetCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetActiveSubnetCount(&bind.CallOpts{Context: ctx})
}

// GetCreationStats returns totalCreated (legacy interface compatibility).
func (s *SubnetFactoryService) GetCreationStats(ctx context.Context) (*big.Int, error) {
	return s.contract.GetCreationStats(&bind.CallOpts{Context: ctx})
}

// PredictSubnetAddress predicts the address using clone determinism.
func (s *SubnetFactoryService) PredictSubnetAddress(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.PredictSubnetAddress(&bind.CallOpts{Context: ctx}, subnetID)
}

// BindSubnet returns a Subnet contract instance for querying participants and other information.
func (s *SubnetFactoryService) BindSubnet(subnetAddr common.Address) (*subnetcontract.Subnet, error) {
	return subnetcontract.NewSubnet(subnetAddr, s.backend)
}

// NewSubnetServiceByAddress creates a SubnetService by subnet contract address.
func (s *SubnetFactoryService) NewSubnetServiceByAddress(subnetAddr common.Address) (*SubnetService, error) {
	contract, err := s.BindSubnet(subnetAddr)
	if err != nil {
		return nil, err
	}
	return NewSubnetService(contract, s.txMgr), nil
}

// NewSubnetServiceByID queries the subnet address by ID and creates a SubnetService.
func (s *SubnetFactoryService) NewSubnetServiceByID(ctx context.Context, subnetID [32]byte) (*SubnetService, error) {
	addr, err := s.GetSubnetContract(ctx, subnetID)
	if err != nil {
		return nil, err
	}
	return s.NewSubnetServiceByAddress(addr)
}

// PauseSubnet calls pauseSubnet (requires GUARDIAN_ROLE).
func (s *SubnetFactoryService) PauseSubnet(ctx context.Context, subnetID [32]byte, reason string) (*types.Transaction, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, errors.New("subnet: pause reason required")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.PauseSubnet(opts, subnetID, reason)
	})
}

// ResumeSubnet calls resumeSubnet.
func (s *SubnetFactoryService) ResumeSubnet(ctx context.Context, subnetID [32]byte) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ResumeSubnet(opts, subnetID)
	})
}

// DeprecateSubnet calls deprecateSubnet.
func (s *SubnetFactoryService) DeprecateSubnet(ctx context.Context, subnetID [32]byte, reason string) (*types.Transaction, error) {
	if strings.TrimSpace(reason) == "" {
		return nil, errors.New("subnet: deprecate reason required")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.DeprecateSubnet(opts, subnetID, reason)
	})
}

// GetActiveParticipants batch retrieves active participant info for specified subnets.
// Returns three types of active participants for each subnet: VALIDATOR, AGENT, MATCHER.
// The return array length always equals the input length; empty arrays are returned for non-existent subnets.
func (s *SubnetFactoryService) GetActiveParticipants(ctx context.Context, subnetIDs [][32]byte) ([]subnetfactory.DataStructuresSubnetParticipants, error) {
	return s.contract.GetActiveParticipants(&bind.CallOpts{Context: ctx}, subnetIDs)
}

// ======================== Read-only methods: Roles & Permissions ========================

// DefaultAdminRole returns the default admin role hash.
func (s *SubnetFactoryService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GuardianRole returns the guardian role hash.
func (s *SubnetFactoryService) GuardianRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GUARDIANROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin returns the admin role for the specified role.
func (s *SubnetFactoryService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole checks if an account has the specified role.
func (s *SubnetFactoryService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== Read-only methods: Subnet queries & filtering ========================

// FindSubnetByName finds a subnet ID by canonical name.
func (s *SubnetFactoryService) FindSubnetByName(ctx context.Context, canonicalName string) ([32]byte, error) {
	return s.contract.FindSubnetByName(&bind.CallOpts{Context: ctx}, canonicalName)
}

// GetAllSubnetInfo returns complete information for all subnets.
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

// GetSubnetsByOwner returns all subnet IDs owned by the specified owner.
func (s *SubnetFactoryService) GetSubnetsByOwner(ctx context.Context, owner common.Address) ([][32]byte, error) {
	return s.contract.GetSubnetsByOwner(&bind.CallOpts{Context: ctx}, owner)
}

// GetSubnetsByStatus returns all subnet IDs with the specified status.
func (s *SubnetFactoryService) GetSubnetsByStatus(ctx context.Context, status uint8) ([][32]byte, error) {
	return s.contract.GetSubnetsByStatus(&bind.CallOpts{Context: ctx}, status)
}

// GetTotalSubnetCount returns the total subnet count (equivalent to TotalCreated).
func (s *SubnetFactoryService) GetTotalSubnetCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetTotalSubnetCount(&bind.CallOpts{Context: ctx})
}

// ListAllSubnets returns all subnet IDs (equivalent to EnumerateSubnets).
func (s *SubnetFactoryService) ListAllSubnets(ctx context.Context) ([][32]byte, error) {
	return s.contract.ListAllSubnets(&bind.CallOpts{Context: ctx})
}

// SubnetExists checks if a subnet exists.
func (s *SubnetFactoryService) SubnetExists(ctx context.Context, subnetID [32]byte) (bool, error) {
	return s.contract.SubnetExists(&bind.CallOpts{Context: ctx}, subnetID)
}

// ======================== Read-only methods: Subnet configuration queries ========================

// GetSubnetStakeConfig returns the subnet's staking governance configuration.
func (s *SubnetFactoryService) GetSubnetStakeConfig(ctx context.Context, subnetID [32]byte) (subnetfactory.DataStructuresStakeGovernanceConfig, error) {
	return s.contract.GetSubnetStakeConfig(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetSubnetThreshold returns the subnet's signature threshold configuration.
func (s *SubnetFactoryService) GetSubnetThreshold(ctx context.Context, subnetID [32]byte) (subnetfactory.DataStructuresSignatureThreshold, error) {
	return s.contract.GetSubnetThreshold(&bind.CallOpts{Context: ctx}, subnetID)
}

// ======================== Read-only methods: Global configuration queries ========================

// GetMinStakeCreateSubnet returns the minimum stake requirement for creating a subnet.
func (s *SubnetFactoryService) GetMinStakeCreateSubnet(ctx context.Context) (*big.Int, error) {
	return s.contract.GetMinStakeCreateSubnet(&bind.CallOpts{Context: ctx})
}

// GetStakingManager returns the StakingManager contract address.
func (s *SubnetFactoryService) GetStakingManager(ctx context.Context) (common.Address, error) {
	return s.contract.GetStakingManager(&bind.CallOpts{Context: ctx})
}

// GetSubnetImplementation returns the subnet implementation contract address.
func (s *SubnetFactoryService) GetSubnetImplementation(ctx context.Context) (common.Address, error) {
	return s.contract.GetSubnetImplementation(&bind.CallOpts{Context: ctx})
}

// StakingManager returns the StakingManager contract address (field accessor).
func (s *SubnetFactoryService) StakingManager(ctx context.Context) (common.Address, error) {
	return s.contract.StakingManager(&bind.CallOpts{Context: ctx})
}

// SubnetImplementation returns the subnet implementation contract address (field accessor).
func (s *SubnetFactoryService) SubnetImplementation(ctx context.Context) (common.Address, error) {
	return s.contract.SubnetImplementation(&bind.CallOpts{Context: ctx})
}

// Paused returns whether SubnetFactory is in a paused state.
func (s *SubnetFactoryService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface checks if the contract supports the specified interface (ERC165).
func (s *SubnetFactoryService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// AllSubnets returns the subnet ID at the specified index position (for iteration).
func (s *SubnetFactoryService) AllSubnets(ctx context.Context, index *big.Int) ([32]byte, error) {
	return s.contract.AllSubnets(&bind.CallOpts{Context: ctx}, index)
}

// SubnetContracts returns the contract address for a subnet ID (mapping accessor).
func (s *SubnetFactoryService) SubnetContracts(ctx context.Context, subnetID [32]byte) (common.Address, error) {
	return s.contract.SubnetContracts(&bind.CallOpts{Context: ctx}, subnetID)
}

// ======================== Write methods: Subnet creation ========================

// CreateSubnet creates a new subnet (must satisfy minimum stake requirements).
func (s *SubnetFactoryService) CreateSubnet(ctx context.Context, createInfo subnetfactory.DataStructuresCreateSubnetInfo) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.CreateSubnet(opts, createInfo)
	})
}

// ======================== Write methods: Emergency control ========================

// EmergencyPause emergency pauses SubnetFactory (requires GUARDIAN_ROLE).
func (s *SubnetFactoryService) EmergencyPause(ctx context.Context) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyPause(opts)
	})
}

// EmergencyUnpause lifts the SubnetFactory emergency pause (requires GUARDIAN_ROLE).
func (s *SubnetFactoryService) EmergencyUnpause(ctx context.Context) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyUnpause(opts)
	})
}

// ======================== Write methods: Role management ========================

// GrantRole grants the specified role to an account (requires role admin permission).
func (s *SubnetFactoryService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole revokes the specified role from an account (requires role admin permission).
func (s *SubnetFactoryService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole renounces the caller's specified role.
func (s *SubnetFactoryService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// ======================== Write methods: Global configuration management ========================

// SetMinStakeCreateSubnet sets the minimum stake requirement for creating a subnet (requires DEFAULT_ADMIN_ROLE).
func (s *SubnetFactoryService) SetMinStakeCreateSubnet(ctx context.Context, amount *big.Int) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetMinStakeCreateSubnet(opts, amount)
	})
}

// SetStakingManager sets the StakingManager contract address (requires DEFAULT_ADMIN_ROLE).
func (s *SubnetFactoryService) SetStakingManager(ctx context.Context, stakingManager common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetStakingManager(opts, stakingManager)
	})
}

// SetSubnetImplementation sets the subnet implementation contract address (requires DEFAULT_ADMIN_ROLE).
func (s *SubnetFactoryService) SetSubnetImplementation(ctx context.Context, newImplementation common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetSubnetImplementation(opts, newImplementation)
	})
}

// Initialize initializes the SubnetFactory contract (call only once after deployment).
func (s *SubnetFactoryService) Initialize(ctx context.Context, guardian common.Address, subnetImplementation common.Address, stakingManager common.Address) (*types.Transaction, error) {
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, guardian, subnetImplementation, stakingManager)
	})
}
