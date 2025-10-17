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

// ParticipantType aligns with DataStructures.ParticipantType.
type ParticipantType uint8

const (
	ParticipantValidator   ParticipantType = 0
	ParticipantAgent       ParticipantType = 1
	ParticipantMatcher     ParticipantType = 2
	ParticipantSubnetOwner ParticipantType = 3
)

// StakingService provides common StakingManager operations.
type StakingService struct {
	contract *stakingmanager.StakingManager
	txMgr    *txmgr.Manager
}

// NewStakingService creates a new service instance.
func NewStakingService(contract *stakingmanager.StakingManager) *StakingService {
	return &StakingService{contract: contract}
}

// AttachTxManager attaches a TxManager for write operations.
func (s *StakingService) AttachTxManager(txm *txmgr.Manager) {
	s.txMgr = txm
}

// GetStakeInfo queries stake information.
func (s *StakingService) GetStakeInfo(ctx context.Context, user, token common.Address, role ParticipantType) (stakingmanager.DataStructuresStakeInfo, error) {
	return s.contract.GetStakeInfo(&bind.CallOpts{Context: ctx}, user, token, uint8(role))
}

// GetUnlockPeriod queries the unlock period.
func (s *StakingService) GetUnlockPeriod(ctx context.Context) (*big.Int, error) {
	return s.contract.GetUnlockPeriod(&bind.CallOpts{Context: ctx})
}

// DepositStakeFor calls depositStakeFor. Caller must have allowance or send ETH.
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

// RequestUnstake initiates an unstake request.
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

// Withdraw withdraws completed unstake amounts.
func (s *StakingService) Withdraw(ctx context.Context, token common.Address, role ParticipantType) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Withdraw(opts, token, uint8(role))
	})
}

// ======================== Read-only methods: Roles & Permissions ========================

// DefaultAdminRole returns the default admin role hash.
func (s *StakingService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GovernanceRole returns the governance role hash.
func (s *StakingService) GovernanceRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GOVERNANCEROLE(&bind.CallOpts{Context: ctx})
}

// SlasherRole returns the slasher role hash.
func (s *StakingService) SlasherRole(ctx context.Context) ([32]byte, error) {
	return s.contract.SLASHERROLE(&bind.CallOpts{Context: ctx})
}

// StakingAdminRole returns the staking admin role hash.
func (s *StakingService) StakingAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.STAKINGADMINROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin returns the admin role for the specified role.
func (s *StakingService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole checks if an account has the specified role.
func (s *StakingService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== Read-only methods: Configuration queries ========================

// GetStakingToken returns the staking token address.
func (s *StakingService) GetStakingToken(ctx context.Context) (common.Address, error) {
	return s.contract.GetStakingToken(&bind.CallOpts{Context: ctx})
}

// Paused returns whether the StakingManager is in a paused state.
func (s *StakingService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface checks if the contract supports the specified interface (ERC165).
func (s *StakingService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ======================== Write methods: Role management ========================

// GrantRole grants the specified role to an account (requires role admin permission).
func (s *StakingService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole revokes the specified role from an account (requires role admin permission).
func (s *StakingService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole renounces the caller's specified role.
func (s *StakingService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// SetGovernanceRole sets the governance role account (requires DEFAULT_ADMIN_ROLE).
func (s *StakingService) SetGovernanceRole(ctx context.Context, governance common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetGovernanceRole(opts, governance)
	})
}

// SetSlasherRole sets the slasher role account (requires DEFAULT_ADMIN_ROLE).
func (s *StakingService) SetSlasherRole(ctx context.Context, slasher common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetSlasherRole(opts, slasher)
	})
}

// ======================== Write methods: Configuration management ========================

// SetStakingToken sets the staking token address (requires GOVERNANCE_ROLE).
func (s *StakingService) SetStakingToken(ctx context.Context, token common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetStakingToken(opts, token)
	})
}

// SetUnlockPeriod sets the unlock period (requires GOVERNANCE_ROLE).
func (s *StakingService) SetUnlockPeriod(ctx context.Context, period *big.Int) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetUnlockPeriod(opts, period)
	})
}

// ======================== Write methods: Emergency control & Slashing ========================

// Pause pauses the StakingManager (requires GOVERNANCE_ROLE).
func (s *StakingService) Pause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Pause(opts)
	})
}

// Unpause unpauses the StakingManager (requires GOVERNANCE_ROLE).
func (s *StakingService) Unpause(ctx context.Context) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Unpause(opts)
	})
}

// Slash executes staking penalties (requires SLASHER_ROLE).
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

// EmergencyWithdraw performs emergency fund withdrawal (requires GOVERNANCE_ROLE).
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

// Initialize initializes the StakingManager contract (call only once after deployment).
func (s *StakingService) Initialize(ctx context.Context, admin common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("staking: tx manager not attached")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, admin)
	})
}
