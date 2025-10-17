package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	checkpointmanager "github.com/PIN-AI/intent-protocol-contract-sdk/contracts/checkpointmanager"
	cryptoHelpers "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/txmgr"
)

// CheckpointState is an alias for the contract return structure.
type CheckpointState = checkpointmanager.DataStructuresCheckpointState

// CheckpointService provides common CheckpointManager operations (read and write).
type CheckpointService struct {
	contract     *checkpointmanager.CheckpointManager
	backend      bind.ContractBackend
	txMgr        *txmgr.Manager
	signer       signer.Signer
	chainID      *big.Int
	contractAddr common.Address
}

// NewCheckpointService creates a new service instance.
func NewCheckpointService(
	contract *checkpointmanager.CheckpointManager,
	backend bind.ContractBackend,
	txm *txmgr.Manager,
	sig signer.Signer,
	chainID *big.Int,
	contractAddr common.Address,
) *CheckpointService {
	return &CheckpointService{
		contract:     contract,
		backend:      backend,
		txMgr:        txm,
		signer:       sig,
		chainID:      new(big.Int).Set(chainID),
		contractAddr: contractAddr,
	}
}

// AttachTxManager allows updating the TxManager at runtime.
func (s *CheckpointService) AttachTxManager(txm *txmgr.Manager) {
	s.txMgr = txm
}

// GetCheckpoint queries the checkpoint state for a specified subnet/epoch.
func (s *CheckpointService) GetCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (CheckpointState, error) {
	return s.contract.GetCheckpoint(&bind.CallOpts{Context: ctx}, subnetID, epoch)
}

// GetCheckpointProof returns the proof corresponding to the checkpoint.
func (s *CheckpointService) GetCheckpointProof(ctx context.Context, subnetID [32]byte, epoch uint64) ([]byte, error) {
	return s.contract.GetCheckpointProof(&bind.CallOpts{Context: ctx}, subnetID, epoch)
}

// CheckpointHeader represents the header data required for checkpoint submission.
type CheckpointHeader struct {
	SubnetID             [32]byte
	Epoch                uint64
	ParentCheckpointHash [32]byte
	Timestamp            *big.Int
	Version              uint32
	ParamsHash           [32]byte
	Roots                cryptoHelpers.CommitmentRoots
	DACommitments        []cryptoHelpers.DACommitment
	EpochSlot            cryptoHelpers.EpochSlot
	Stats                []byte
	AuxHash              [32]byte
	AssignmentsRoot      [32]byte
	ValidationCommitment [32]byte
	PolicyRoot           [32]byte
}

// SubmitCheckpoint calls the contract's submitCheckpoint method.
func (s *CheckpointService) SubmitCheckpoint(
	ctx context.Context,
	header CheckpointHeader,
	validators []common.Address,
	signatures [][]byte,
) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	if len(validators) != len(signatures) {
		return nil, errors.New("checkpoint: validators/signatures length mismatch")
	}
	contractHeader := convertCheckpointHeader(header)
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SubmitCheckpoint(opts, contractHeader, validators, signatures)
	})
}

// ComputeCheckpointHash computes the checkpoint hash via contract.
func (s *CheckpointService) ComputeCheckpointHash(ctx context.Context, header CheckpointHeader) ([32]byte, error) {
	return s.contract.ComputeCheckpointHash(&bind.CallOpts{Context: ctx}, convertCheckpointHeader(header))
}

// ComputeDigest computes the signature digest required for submission.
func (s *CheckpointService) ComputeDigest(header CheckpointHeader) ([32]byte, error) {
	return cryptoHelpers.ComputeCheckpointDigest(cryptoHelpers.CheckpointInput{
		SubnetID:             header.SubnetID,
		Epoch:                header.Epoch,
		ParentCheckpointHash: header.ParentCheckpointHash,
		Timestamp:            header.Timestamp,
		Version:              header.Version,
		ParamsHash:           header.ParamsHash,
		Roots:                header.Roots,
		DACommitments:        header.DACommitments,
		EpochSlot:            header.EpochSlot,
		Stats:                header.Stats,
		AuxHash:              header.AuxHash,
		AssignmentsRoot:      header.AssignmentsRoot,
		ValidationCommitment: header.ValidationCommitment,
		PolicyRoot:           header.PolicyRoot,
	}, s.contractAddr, s.chainID)
}

// SignDigest signs the digest using the bound signer.
func (s *CheckpointService) SignDigest(digest [32]byte) ([]byte, error) {
	if s.signer == nil {
		return nil, errors.New("checkpoint: signer not configured")
	}
	return s.signer.SignDigest(digest)
}

func convertCheckpointHeader(header CheckpointHeader) checkpointmanager.DataStructuresCheckpointHeader {
	return checkpointmanager.DataStructuresCheckpointHeader{
		SubnetId:             header.SubnetID,
		Epoch:                header.Epoch,
		ParentCpHash:         header.ParentCheckpointHash,
		Timestamp:            header.Timestamp,
		Version:              header.Version,
		ParamsHash:           header.ParamsHash,
		Roots:                convertCommitmentRoots(header.Roots),
		DaCommitments:        convertDACommitments(header.DACommitments),
		EpochSlot:            convertEpochSlot(header.EpochSlot),
		Stats:                header.Stats,
		AuxHash:              header.AuxHash,
		AssignmentsRoot:      header.AssignmentsRoot,
		ValidationCommitment: header.ValidationCommitment,
		PolicyRoot:           header.PolicyRoot,
	}
}

func convertCommitmentRoots(roots cryptoHelpers.CommitmentRoots) checkpointmanager.DataStructuresCommitmentRoots {
	return checkpointmanager.DataStructuresCommitmentRoots{
		AgentRoot:        roots.AgentRoot,
		AgentServiceRoot: roots.AgentServiceRoot,
		RankRoot:         roots.RankRoot,
		MetricsRoot:      roots.MetricsRoot,
		DataUsageRoot:    roots.DataUsageRoot,
		StateRoot:        roots.StateRoot,
		EventRoot:        roots.EventRoot,
		CrossSubnetRoot:  roots.CrossSubnetRoot,
	}
}

func convertDACommitments(items []cryptoHelpers.DACommitment) []checkpointmanager.DataStructuresDACommitment {
	if len(items) == 0 {
		return nil
	}
	out := make([]checkpointmanager.DataStructuresDACommitment, len(items))
	for i, item := range items {
		out[i] = checkpointmanager.DataStructuresDACommitment{
			Kind:          item.Kind,
			Pointer:       item.Pointer,
			SizeHint:      normalizeBigInt(item.SizeHint),
			SegmentHashes: item.SegmentHashes,
			Expiry:        normalizeBigInt(item.Expiry),
		}
	}
	return out
}

func convertEpochSlot(slot cryptoHelpers.EpochSlot) checkpointmanager.DataStructuresEpochSlot {
	return checkpointmanager.DataStructuresEpochSlot{Epoch: slot.Epoch, Slot: slot.Slot}
}

func normalizeBigInt(v *big.Int) *big.Int {
	if v != nil {
		return v
	}
	return big.NewInt(0)
}

// ======================== Read-only methods: Roles & Permissions ========================

// DefaultAdminRole returns the default admin role hash.
func (s *CheckpointService) DefaultAdminRole(ctx context.Context) ([32]byte, error) {
	return s.contract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
}

// GovernanceRole returns the governance role hash.
func (s *CheckpointService) GovernanceRole(ctx context.Context) ([32]byte, error) {
	return s.contract.GOVERNANCEROLE(&bind.CallOpts{Context: ctx})
}

// GetRoleAdmin returns the admin role for the specified role.
func (s *CheckpointService) GetRoleAdmin(ctx context.Context, role [32]byte) ([32]byte, error) {
	return s.contract.GetRoleAdmin(&bind.CallOpts{Context: ctx}, role)
}

// HasRole checks if an account has the specified role.
func (s *CheckpointService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error) {
	return s.contract.HasRole(&bind.CallOpts{Context: ctx}, role, account)
}

// ======================== Read-only methods: Configuration & State queries ========================

// DefaultChallengeWindow returns the default challenge window.
func (s *CheckpointService) DefaultChallengeWindow(ctx context.Context) (*big.Int, error) {
	return s.contract.DEFAULTCHALLENGEWINDOW(&bind.CallOpts{Context: ctx})
}

// GetLatestCheckpoint returns the subnet's latest checkpoint.
func (s *CheckpointService) GetLatestCheckpoint(ctx context.Context, subnetID [32]byte) (CheckpointState, error) {
	return s.contract.GetLatestCheckpoint(&bind.CallOpts{Context: ctx}, subnetID)
}

// GetLatestFinalized returns the subnet's latest finalized checkpoint.
func (s *CheckpointService) GetLatestFinalized(ctx context.Context, subnetID [32]byte) (CheckpointState, error) {
	return s.contract.GetLatestFinalized(&bind.CallOpts{Context: ctx}, subnetID)
}

// CanFinalizeCheckpoint checks if the specified checkpoint can be finalized.
func (s *CheckpointService) CanFinalizeCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (bool, error) {
	return s.contract.CanFinalizeCheckpoint(&bind.CallOpts{Context: ctx}, subnetID, epoch)
}

// VerifyCheckpoint verifies the integrity and signatures of the specified checkpoint.
func (s *CheckpointService) VerifyCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (bool, string, error) {
	result, err := s.contract.VerifyCheckpoint(&bind.CallOpts{Context: ctx}, subnetID, epoch)
	if err != nil {
		return false, "", err
	}
	return result.Valid, result.Reason, nil
}

// VerifyCheckpointContinuity verifies if the checkpoint is continuous with the previous checkpoint.
func (s *CheckpointService) VerifyCheckpointContinuity(ctx context.Context, subnetID [32]byte, header CheckpointHeader) (bool, error) {
	return s.contract.VerifyCheckpointContinuity(&bind.CallOpts{Context: ctx}, subnetID, convertCheckpointHeader(header))
}

// Paused returns whether the CheckpointManager is in a paused state.
func (s *CheckpointService) Paused(ctx context.Context) (bool, error) {
	return s.contract.Paused(&bind.CallOpts{Context: ctx})
}

// SupportsInterface checks if the contract supports the specified interface (ERC165).
func (s *CheckpointService) SupportsInterface(ctx context.Context, interfaceID [4]byte) (bool, error) {
	return s.contract.SupportsInterface(&bind.CallOpts{Context: ctx}, interfaceID)
}

// ======================== Write methods: Checkpoint management ========================

// FinalizeCheckpoint finalizes the specified checkpoint (requires GOVERNANCE_ROLE).
func (s *CheckpointService) FinalizeCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.FinalizeCheckpoint(opts, subnetID, epoch)
	})
}

// RevertCheckpoint reverts the specified checkpoint (requires GOVERNANCE_ROLE).
func (s *CheckpointService) RevertCheckpoint(ctx context.Context, subnetID [32]byte, epoch uint64, reason string) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevertCheckpoint(opts, subnetID, epoch, reason)
	})
}

// ======================== Write methods: Role management ========================

// GrantRole grants the specified role to an account (requires role admin permission).
func (s *CheckpointService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.GrantRole(opts, role, account)
	})
}

// RevokeRole revokes the specified role from an account (requires role admin permission).
func (s *CheckpointService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RevokeRole(opts, role, account)
	})
}

// RenounceRole renounces the caller's specified role.
func (s *CheckpointService) RenounceRole(ctx context.Context, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.RenounceRole(opts, role, callerConfirmation)
	})
}

// SetGovernanceRole sets the governance role account (requires DEFAULT_ADMIN_ROLE).
func (s *CheckpointService) SetGovernanceRole(ctx context.Context, governance common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.SetGovernanceRole(opts, governance)
	})
}

// Initialize initializes the CheckpointManager contract (call only once after deployment).
func (s *CheckpointService) Initialize(ctx context.Context, admin common.Address, registry common.Address) (*types.Transaction, error) {
	if s.txMgr == nil {
		return nil, errors.New("checkpoint: tx manager not configured")
	}
	return s.txMgr.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.Initialize(opts, admin, registry)
	})
}
