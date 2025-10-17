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

// SubnetService wraps common Subnet contract calls.
type SubnetService struct {
	contract *subnetcontract.Subnet
	txMgr    *txmgr.Manager
}

// NewSubnetService constructs the service using a bound Subnet contract instance and TxManager.
func NewSubnetService(contract *subnetcontract.Subnet, txm *txmgr.Manager) *SubnetService {
	return &SubnetService{contract: contract, txMgr: txm}
}

// NewSubnetServiceByAddress binds a Subnet by contract address and returns a service instance.
func NewSubnetServiceByAddress(backend bind.ContractBackend, addr common.Address, txm *txmgr.Manager) (*SubnetService, error) {
	contract, err := subnetcontract.NewSubnet(addr, backend)
	if err != nil {
		return nil, err
	}
	return &SubnetService{contract: contract, txMgr: txm}, nil
}

// RegisterParticipantParams describes configurable options for RegisterParticipant calls.
type RegisterParticipantParams struct {
	Domain      string
	Endpoint    string
	MetadataURI string
	Value       *big.Int // Optional: native token amount required for registration (if any)
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

// RegisterValidator registers a validator using RegisterParticipant.
func (s *SubnetService) RegisterValidator(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantValidator, params)
}

// RegisterAgent registers an agent using RegisterParticipant.
func (s *SubnetService) RegisterAgent(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantAgent, params)
}

// RegisterMatcher registers a matcher using RegisterParticipant.
func (s *SubnetService) RegisterMatcher(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error) {
	return s.registerParticipant(ctx, ParticipantMatcher, params)
}

// RegisterParticipantERC20Params describes input for RegisterParticipantERC20.
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

// RegisterValidatorERC20 registers a validator via ERC20 staking.
func (s *SubnetService) RegisterValidatorERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantValidator, params)
}

// RegisterAgentERC20 registers an agent via ERC20 staking.
func (s *SubnetService) RegisterAgentERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantAgent, params)
}

// RegisterMatcherERC20 registers a matcher via ERC20 staking.
func (s *SubnetService) RegisterMatcherERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error) {
	return s.registerParticipantERC20(ctx, ParticipantMatcher, params)
}

// ListActiveParticipants returns a list of active participants of the specified type.
func (s *SubnetService) ListActiveParticipants(ctx context.Context, participantType ParticipantType) ([]subnetcontract.DataStructuresParticipantInfo, error) {
	return s.contract.ListActiveParticipants(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// IsActiveParticipant queries whether the specified address is active for a participant type.
func (s *SubnetService) IsActiveParticipant(ctx context.Context, addr common.Address, participantType ParticipantType) (bool, error) {
	return s.contract.IsActiveParticipant(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantInfo returns participant details.
func (s *SubnetService) GetParticipantInfo(ctx context.Context, addr common.Address, participantType ParticipantType) (subnetcontract.DataStructuresParticipantInfo, error) {
	return s.contract.GetParticipantInfo(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantStakeInfo returns the participant's staking details in the subnet.
func (s *SubnetService) GetParticipantStakeInfo(ctx context.Context, addr common.Address, participantType ParticipantType) (subnetcontract.DataStructuresStakeInfo, error) {
	return s.contract.GetParticipantStakeInfo(&bind.CallOpts{Context: ctx}, addr, uint8(participantType))
}

// GetParticipantCount returns the number of participants of the specified type.
func (s *SubnetService) GetParticipantCount(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.GetParticipantCount(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// ======================== Read-only methods: Roles & Permissions ========================

// AdminRole returns the admin role hash.
func (s *SubnetService) AdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.ADMINROLE(&bind.CallOpts{Context: ctx})
}

// DefaultAdminRole returns the default admin role hash.
func (s *SubnetService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin returns the admin role for the specified role.
func (s *SubnetService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole checks if an account has the specified role.
func (s *SubnetService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== Read-only methods: Subnet configuration queries ========================

// GetSubnetInfo returns the subnet's complete information.
func (s *SubnetService) GetSubnetInfo(ctx context.Context) (subnetcontract.DataStructuresSubnetInfo, error) {
	return s.contract.GetSubnetInfo(&bind.CallOpts{Context: ctx})
}

// GetBidFrequencyLimit returns the subnet's bid frequency limit.
func (s *SubnetService) GetBidFrequencyLimit(ctx context.Context) (*big.Int, error) {
	return s.contract.GetBidFrequencyLimit(&bind.CallOpts{Context: ctx})
}

// GetCheckpointPolicy returns the subnet's checkpoint policy configuration.
func (s *SubnetService) GetCheckpointPolicy(ctx context.Context) (subnetcontract.DataStructuresCheckpointPolicy, error) {
	return s.contract.GetCheckpointPolicy(&bind.CallOpts{Context: ctx})
}

// GetSignatureThreshold returns the subnet's signature threshold configuration.
func (s *SubnetService) GetSignatureThreshold(ctx context.Context) (subnetcontract.DataStructuresSignatureThreshold, error) {
	return s.contract.GetSignatureThreshold(&bind.CallOpts{Context: ctx})
}

// GetStakeGovernanceConfig returns the subnet's staking governance configuration.
func (s *SubnetService) GetStakeGovernanceConfig(ctx context.Context) (subnetcontract.DataStructuresStakeGovernanceConfig, error) {
	return s.contract.GetStakeGovernanceConfig(&bind.CallOpts{Context: ctx})
}

// GetSlashingRates returns the slashing rate list.
func (s *SubnetService) GetSlashingRates(ctx context.Context) ([]*big.Int, error) {
	return s.contract.GetSlashingRates(&bind.CallOpts{Context: ctx})
}

// GetUnstakeLockPeriod returns the unlock period.
func (s *SubnetService) GetUnstakeLockPeriod(ctx context.Context) (*big.Int, error) {
	return s.contract.GetUnstakeLockPeriod(&bind.CallOpts{Context: ctx})
}

// GetStakingToken returns the staking token address.
func (s *SubnetService) GetStakingToken(ctx context.Context) (common.Address, error) {
	return s.contract.GetStakingToken(&bind.CallOpts{Context: ctx})
}

// GetRequireKYC returns whether the subnet requires KYC.
func (s *SubnetService) GetRequireKYC(ctx context.Context) (bool, error) {
	return s.contract.GetRequireKYC(&bind.CallOpts{Context: ctx})
}

// GetMinimumStakeByType returns the minimum stake requirement for the specified participant type.
func (s *SubnetService) GetMinimumStakeByType(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.GetMinimumStakeByType(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// ======================== Read-only methods: Participant statistics queries ========================

// GetTotalParticipantCount returns the total count of participants of all types.
func (s *SubnetService) GetTotalParticipantCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetTotalParticipantCount(&bind.CallOpts{Context: ctx})
}

// GetActiveParticipantCount returns the number of active participants of the specified type.
func (s *SubnetService) GetActiveParticipantCount(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.GetActiveParticipantCount(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// ActiveParticipantCounts returns the active participant count for the specified type (mapping accessor).
func (s *SubnetService) ActiveParticipantCounts(ctx context.Context, participantType ParticipantType) (*big.Int, error) {
	return s.contract.ActiveParticipantCounts(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// HasAvailableSlots checks if the specified participant type has available slots.
func (s *SubnetService) HasAvailableSlots(ctx context.Context, participantType ParticipantType) (bool, error) {
	return s.contract.HasAvailableSlots(&bind.CallOpts{Context: ctx}, uint8(participantType))
}

// Paused returns whether the subnet is in a paused state.
func (s *SubnetService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface checks if the contract supports the specified interface (ERC165).
func (s *SubnetService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ParticipantListByType returns the address at the specified index in the participant list of the specified type (mapping accessor).
func (s *SubnetService) ParticipantListByType(ctx context.Context, participantType ParticipantType, index *big.Int) (common.Address, error) {
	return s.contract.ParticipantListByType(&bind.CallOpts{Context: ctx}, uint8(participantType), index)
}

// ======================== Write methods: Participant management ========================

// ApproveParticipant approves participant registration (requires ADMIN_ROLE).
func (s *SubnetService) ApproveParticipant(ctx context.Context, participantAddr common.Address, participantType ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ApproveParticipant(opts, participantAddr, uint8(participantType))
	})
}

// RejectParticipant rejects participant registration (requires ADMIN_ROLE).
func (s *SubnetService) RejectParticipant(ctx context.Context, participantAddr common.Address, participantType ParticipantType, reason string) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RejectParticipant(opts, participantAddr, uint8(participantType), reason)
	})
}

// SuspendParticipant suspends a participant (requires ADMIN_ROLE).
func (s *SubnetService) SuspendParticipant(ctx context.Context, participantType ParticipantType, reason string) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SuspendParticipant(opts, uint8(participantType), reason)
	})
}

// ResumeParticipant resumes a suspended participant (requires ADMIN_ROLE).
func (s *SubnetService) ResumeParticipant(ctx context.Context, participantAddr common.Address, participantType ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ResumeParticipant(opts, participantAddr, uint8(participantType))
	})
}

// UpdateParticipantActivity updates the participant activity timestamp (requires ADMIN_ROLE).
func (s *SubnetService) UpdateParticipantActivity(ctx context.Context, participantType ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.UpdateParticipantActivity(opts, uint8(participantType))
	})
}

// ======================== Write methods: Role management ========================

// GrantRole grants the specified role to an account (requires role admin permission).
func (s *SubnetService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole revokes the specified role from an account (requires role admin permission).
func (s *SubnetService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole renounces the caller's specified role.
func (s *SubnetService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// ======================== Write methods: Subnet management ========================

// Pause pauses the subnet (requires ADMIN_ROLE).
func (s *SubnetService) Pause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Pause(opts)
	})
}

// Unpause unpauses the subnet (requires ADMIN_ROLE).
func (s *SubnetService) Unpause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Unpause(opts)
	})
}

// Deprecate marks the subnet as deprecated (requires ADMIN_ROLE).
func (s *SubnetService) Deprecate(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Deprecate(opts)
	})
}

// SetStakeConfig sets the subnet's staking configuration (requires ADMIN_ROLE).
func (s *SubnetService) SetStakeConfig(ctx context.Context, stakeConfig subnetcontract.DataStructuresStakeGovernanceConfig) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetStakeConfig(opts, stakeConfig)
	})
}

// Initialize initializes the subnet contract (call only once after deployment).
func (s *SubnetService) Initialize(ctx context.Context, subnetInfo subnetcontract.DataStructuresSubnetInfo, factory common.Address, stakingManager common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("subnet: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, subnetInfo, factory, stakingManager)
	})
}
