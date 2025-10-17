package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	intentmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/intentmanager"
	cryptoHelpers "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// SubmitIntentParams describes the parameters required for submitIntent.
type SubmitIntentParams struct {
	IntentID     [32]byte
	SubnetID     [32]byte
	IntentType   string
	ParamsHash   [32]byte
	Deadline     *big.Int
	PaymentToken common.Address
	Amount       *big.Int
	Value        *big.Int // Optional, defaults to Amount for ETH payments
}

// SignedIntent describes a single intent's information and signature for batch submission.
type SignedIntent struct {
	Data      cryptoHelpers.SignedIntentInput
	Signature []byte
}

// SubmitIntentBatchParams describes batch submission parameters.
type SubmitIntentBatchParams struct {
	Items         []SignedIntent
	TotalEthValue *big.Int // Optional, auto-calculates sum of PaymentToken==0 amounts if not provided
}

// IntentInfo is consistent with the contract return structure.
type IntentInfo = intentmanager.DataStructuresIntentInfo

// IntentService provides high-level IntentManager operations.
type IntentService struct {
	backend      bind.ContractBackend
	txManager    *txmgr.Manager
	contract     *intentmanager.IntentManager
	signer       signer.Signer
	chainID      *big.Int
	contractAddr common.Address
}

// NewIntentService constructs an IntentService instance.
func NewIntentService(backend bind.ContractBackend, txm *txmgr.Manager, contract *intentmanager.IntentManager, sig signer.Signer, chainID *big.Int, contractAddr common.Address) *IntentService {
	return &IntentService{
		backend:      backend,
		txManager:    txm,
		contract:     contract,
		signer:       sig,
		chainID:      new(big.Int).Set(chainID),
		contractAddr: contractAddr,
	}
}

// SubmitIntent calls the contract's submitIntent method.
func (s *IntentService) SubmitIntent(ctx context.Context, params SubmitIntentParams) (*types.Transaction, error) {
	if params.Deadline == nil {
		return nil, errors.New("intent: deadline is required")
	}
	if params.Amount == nil {
		return nil, errors.New("intent: amount is required")
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		if params.Value != nil {
			opts.Value = params.Value
		} else if params.PaymentToken == ZeroAddress {
			opts.Value = params.Amount
		}
		return s.contract.SubmitIntent(opts, params.IntentID, params.SubnetID, params.IntentType, params.ParamsHash, params.Deadline, params.PaymentToken, params.Amount)
	})
}

// SubmitIntentsBySignatures submits intents in batch.
func (s *IntentService) SubmitIntentsBySignatures(ctx context.Context, params SubmitIntentBatchParams) (*types.Transaction, error) {
	if len(params.Items) == 0 {
		return nil, errors.New("intent: empty batch")
	}
	intents := make([]intentmanager.IIntentManagerIntentData, len(params.Items))
	signatures := make([][]byte, len(params.Items))
	totalEth := new(big.Int)
	for i, item := range params.Items {
		if item.Data.Deadline == nil {
			return nil, errors.New("intent: nil deadline in batch item")
		}
		if item.Data.Amount == nil {
			return nil, errors.New("intent: nil amount in batch item")
		}
		intents[i] = intentmanager.IIntentManagerIntentData{
			IntentId:     item.Data.IntentID,
			SubnetId:     item.Data.SubnetID,
			Requester:    item.Data.Requester,
			IntentType:   item.Data.IntentType,
			ParamsHash:   item.Data.ParamsHash,
			Deadline:     item.Data.Deadline,
			PaymentToken: item.Data.PaymentToken,
			Amount:       item.Data.Amount,
		}
		signatures[i] = item.Signature
		if item.Data.PaymentToken == ZeroAddress {
			totalEth.Add(totalEth, item.Data.Amount)
		}
	}
	if params.TotalEthValue != nil {
		totalEth = new(big.Int).Set(params.TotalEthValue)
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		opts.Value = totalEth
		return s.contract.SubmitIntentsBySignatures(opts, intents, signatures)
	})
}

// FailIntent calls the failIntent method.
func (s *IntentService) FailIntent(ctx context.Context, intentID [32]byte, reason string, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.FailIntent(opts, intentID, reason, validators, signatures)
	})
}

// ProcessExpiredIntent calls the processExpiredIntent method.
func (s *IntentService) ProcessExpiredIntent(ctx context.Context, intentID [32]byte) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ProcessExpiredIntent(opts, intentID)
	})
}

// BatchProcessExpiredIntents calls the batchProcessExpiredIntents method.
func (s *IntentService) BatchProcessExpiredIntents(ctx context.Context, ids [][32]byte) (*types.Transaction, error) {
	if len(ids) == 0 {
		return nil, errors.New("intent: empty ids")
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.BatchProcessExpiredIntents(opts, ids)
	})
}

// GetIntentInfo reads intent information.
func (s *IntentService) GetIntentInfo(ctx context.Context, intentID [32]byte) (IntentInfo, error) {
	return s.contract.GetIntentInfo(&bind.CallOpts{Context: ctx}, intentID)
}

// BatchGetIntentInfo reads intent information in batch.
// Fetches multiple Intent details in one call, avoiding multiple RPC calls.
// Return array order matches input ids order; non-existent Intents return zero-value structs.
func (s *IntentService) BatchGetIntentInfo(ctx context.Context, ids [][32]byte) ([]IntentInfo, error) {
	if len(ids) == 0 {
		return []IntentInfo{}, nil
	}
	return s.contract.BatchGetIntentInfo(&bind.CallOpts{Context: ctx}, ids)
}

// GetIntentsByStatus reads the Intent ID list by status.
func (s *IntentService) GetIntentsByStatus(ctx context.Context, status IntentStatus) ([][32]byte, error) {
	return s.contract.GetIntentsByStatus(&bind.CallOpts{Context: ctx}, uint8(status))
}

// GetSubnetIntents queries all Intent IDs for a specific subnet.
func (s *IntentService) GetSubnetIntents(ctx context.Context, subnetID [32]byte) ([][32]byte, error) {
	return s.contract.GetSubnetIntents(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetPendingIntentCount queries the pending intent count. Queries all if subnetID is nil.
func (s *IntentService) GetPendingIntentCount(ctx context.Context, subnetID *[32]byte) (*big.Int, error) {
	var id [32]byte
	if subnetID != nil {
		id = *subnetID
	}
	return s.contract.GetPendingIntentCount(&bind.CallOpts{Context: ctx}, id)
}

// GetTotalIntentCount queries the total intent count.
func (s *IntentService) GetTotalIntentCount(ctx context.Context) (*big.Int, error) {
	return s.contract.GetTotalIntentCount(&bind.CallOpts{Context: ctx})
}

// GetRequiredValidatorCount queries the subnet validator threshold.
func (s *IntentService) GetRequiredValidatorCount(ctx context.Context, subnetID [32]byte) (*big.Int, error) {
	return s.contract.GetRequiredValidatorCount(&bind.CallOpts{Context: ctx}, subnetID)
}

// ComputeDigest computes the digest for a single intent, facilitating batch signing.
func (s *IntentService) ComputeDigest(input cryptoHelpers.SignedIntentInput) ([32]byte, error) {
	return cryptoHelpers.ComputeIntentDigest(input, s.contractAddr, s.chainID)
}

// SignDigest performs EIP-191 signing on the digest using the SDK signer.
func (s *IntentService) SignDigest(digest [32]byte) ([]byte, error) {
	return s.signer.SignDigest(digest)
}

// SignIntent wraps ComputeDigest + SignDigest to complete Intent signing in one step.
// Simplifies batch signing workflow, suitable for SubmitIntentsBySignatures scenarios.
func (s *IntentService) SignIntent(input cryptoHelpers.SignedIntentInput) ([]byte, error) {
	digest, err := s.ComputeDigest(input)
	if err != nil {
		return nil, err
	}
	return s.SignDigest(digest)
}

// ======================== Read-only methods: Roles & Permissions ========================

// DefaultAdminRole returns the default admin role hash.
func (s *IntentService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GovernanceRole returns the governance role hash.
func (s *IntentService) GovernanceRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GOVERNANCEROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin returns the admin role for the specified role.
func (s *IntentService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole checks if an account has the specified role.
func (s *IntentService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== Read-only methods: Intent state queries ========================

// IntentExists checks if an Intent exists.
func (s *IntentService) IntentExists(ctx context.Context, intentID [32]byte) (bool, error) {
	return s.contract.IntentExists(&bind.CallOpts{Context: ctx}, intentID)
}

// IsIntentExpired checks if an Intent has expired.
func (s *IntentService) IsIntentExpired(ctx context.Context, intentID [32]byte) (bool, error) {
	return s.contract.IsIntentExpired(&bind.CallOpts{Context: ctx}, intentID)
}

// ======================== Read-only methods: Global configuration queries ========================

// DefaultMaxDuration returns the default maximum Intent duration.
func (s *IntentService) DefaultMaxDuration(ctx context.Context) (*big.Int, error) {
	return s.contract.DEFAULTMAXDURATION(&bind.CallOpts{Context: ctx})
}

// DefaultMinBudget returns the default minimum Intent budget.
func (s *IntentService) DefaultMinBudget(ctx context.Context) (*big.Int, error) {
	return s.contract.DEFAULTMINBUDGET(&bind.CallOpts{Context: ctx})
}

// GetMaxIntentDuration returns the current maximum Intent duration configuration.
func (s *IntentService) GetMaxIntentDuration(ctx context.Context) (*big.Int, error) {
	return s.contract.GetMaxIntentDuration(&bind.CallOpts{Context: ctx})
}

// GetMinIntentBudget returns the current minimum Intent budget configuration.
func (s *IntentService) GetMinIntentBudget(ctx context.Context) (*big.Int, error) {
	return s.contract.GetMinIntentBudget(&bind.CallOpts{Context: ctx})
}

// GetTotalEscrowBalance returns the total escrow balance for the specified token.
func (s *IntentService) GetTotalEscrowBalance(ctx context.Context, token common.Address) (*big.Int, error) {
	return s.contract.GetTotalEscrowBalance(&bind.CallOpts{Context: ctx}, token)
}

// Paused returns whether the IntentManager is in a paused state.
func (s *IntentService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface checks if the contract supports the specified interface (ERC165).
func (s *IntentService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ======================== Write methods: Emergency control ========================

// EmergencyPause emergency pauses the IntentManager (requires GOVERNANCE_ROLE).
func (s *IntentService) EmergencyPause(ctx context.Context) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyPause(opts)
	})
}

// EmergencyUnpause lifts the IntentManager emergency pause (requires GOVERNANCE_ROLE).
func (s *IntentService) EmergencyUnpause(ctx context.Context) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyUnpause(opts)
	})
}

// EmergencyRefundBatch performs emergency batch refund (requires GOVERNANCE_ROLE).
func (s *IntentService) EmergencyRefundBatch(ctx context.Context, intentIDs [][32]byte, reason string) (*types.Transaction, error) {
	if len(intentIDs) == 0 {
		return nil, errors.New("intent: empty intent IDs")
	}
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.EmergencyRefundBatch(opts, intentIDs, reason)
	})
}

// ======================== Write methods: Role management ========================

// GrantRole grants the specified role to an account (requires role admin permission).
func (s *IntentService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole revokes the specified role from an account (requires role admin permission).
func (s *IntentService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole renounces the caller's specified role.
func (s *IntentService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// ======================== Write methods: Global configuration management ========================

// SetMaxIntentDuration sets the maximum Intent duration (requires GOVERNANCE_ROLE).
func (s *IntentService) SetMaxIntentDuration(ctx context.Context, maxDuration *big.Int) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetMaxIntentDuration(opts, maxDuration)
	})
}

// SetMinIntentBudget sets the minimum Intent budget (requires GOVERNANCE_ROLE).
func (s *IntentService) SetMinIntentBudget(ctx context.Context, minBudget *big.Int) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetMinIntentBudget(opts, minBudget)
	})
}

// Initialize initializes the IntentManager contract (call only once after deployment).
func (s *IntentService) Initialize(ctx context.Context, admin common.Address, subnetFactory common.Address) (*types.Transaction, error) {
	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, admin, subnetFactory)
	})
}
