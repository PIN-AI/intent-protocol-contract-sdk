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

// ValidationService orchestrates validation bundle submissions.
type ValidationService struct {
	backend      bind.ContractBackend
	txManager    *txmgr.Manager
	contract     *intentmanager.IntentManager
	signer       signer.Signer
	chainID      *big.Int
	contractAddr common.Address
}

// NewValidationService constructs a validation service.
func NewValidationService(
	backend bind.ContractBackend,
	txm *txmgr.Manager,
	contract *intentmanager.IntentManager,
	sig signer.Signer,
	chainID *big.Int,
	contractAddr common.Address,
) *ValidationService {
	return &ValidationService{
		backend:      backend,
		txManager:    txm,
		contract:     contract,
		signer:       sig,
		chainID:      new(big.Int).Set(chainID),
		contractAddr: contractAddr,
	}
}

// ValidateIntentsBySignatures submits batch validation bundles to finalize multiple intents (v2.3+).
// Each ValidationBatch contains multiple ValidationItems sharing the same validator signatures.
func (s *ValidationService) ValidateIntentsBySignatures(ctx context.Context, batches []ValidationBatch) (*types.Transaction, error) {
	if len(batches) == 0 {
		return nil, errors.New("validation: empty batch array")
	}

	contractBatches := make([]intentmanager.IIntentManagerValidationBatchData, len(batches))
	for i, b := range batches {
		if len(b.Validators) != len(b.Signatures) {
			return nil, errors.New("validation: validator/signature length mismatch")
		}
		if len(b.Items) == 0 {
			return nil, errors.New("validation: batch has no items")
		}

		// Convert ValidationItems to contract items
		contractItems := make([]intentmanager.IIntentManagerValidationItemData, len(b.Items))
		for j, item := range b.Items {
			contractItems[j] = intentmanager.IIntentManagerValidationItemData{
				IntentId:     item.IntentID,
				AssignmentId: item.AssignmentID,
				Agent:        item.Agent,
				ResultHash:   item.ResultHash,
				ProofHash:    item.ProofHash,
			}
		}

		contractBatches[i] = intentmanager.IIntentManagerValidationBatchData{
			SubnetId:   b.SubnetID,
			ItemsHash:  b.ItemsHash,
			RootHeight: b.RootHeight,
			RootHash:   b.RootHash,
			Items:      contractItems,
			Validators: b.Validators,
			Signatures: b.Signatures,
		}
	}

	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ValidateIntentsBySignatures(opts, contractBatches)
	})
}

// ValidateIntentBySignature submits a single validation bundle (backward compatible, v2.2 and earlier).
func (s *ValidationService) ValidateIntentBySignature(ctx context.Context, bundle ValidationBundle) (*types.Transaction, error) {
	if len(bundle.Validators) != len(bundle.Signatures) {
		return nil, errors.New("validation: validator/signature length mismatch")
	}

	contractBundle := intentmanager.IIntentManagerValidationBundleData{
		IntentId:     bundle.IntentID,
		AssignmentId: bundle.AssignmentID,
		SubnetId:     bundle.SubnetID,
		Agent:        bundle.Agent,
		ResultHash:   bundle.ResultHash,
		ProofHash:    bundle.ProofHash,
		RootHeight:   bundle.RootHeight,
		RootHash:     bundle.RootHash,
		Validators:   bundle.Validators,
		Signatures:   bundle.Signatures,
	}

	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ValidateIntentBySignature(opts, contractBundle)
	})
}

// ComputeDigest builds the validation digest for a single validation bundle (backward compatible).
func (s *ValidationService) ComputeDigest(bundle ValidationBundle) ([32]byte, error) {
	return cryptoHelpers.ComputeValidationDigest(cryptoHelpers.ValidationInput{
		IntentID:     bundle.IntentID,
		AssignmentID: bundle.AssignmentID,
		SubnetID:     bundle.SubnetID,
		Agent:        bundle.Agent,
		ResultHash:   bundle.ResultHash,
		ProofHash:    bundle.ProofHash,
		RootHeight:   bundle.RootHeight,
		RootHash:     bundle.RootHash,
	}, s.contractAddr, s.chainID)
}

// ComputeBatchDigest builds the validation digest for a batch validation (v2.3+).
func (s *ValidationService) ComputeBatchDigest(batch ValidationBatch) ([32]byte, error) {
	return cryptoHelpers.ComputeValidationBatchDigest(cryptoHelpers.ValidationBatchInput{
		SubnetID:   batch.SubnetID,
		ItemsHash:  batch.ItemsHash,
		RootHeight: batch.RootHeight,
		RootHash:   batch.RootHash,
	}, s.contractAddr, s.chainID)
}

// ComputeItemsHash computes keccak256(abi.encode(items)) for validation batch.
// This is required for constructing ValidationBatch.ItemsHash.
func (s *ValidationService) ComputeItemsHash(items []ValidationItem) ([32]byte, error) {
	cryptoItems := make([]cryptoHelpers.ValidationItem, len(items))
	for i, item := range items {
		cryptoItems[i] = cryptoHelpers.ValidationItem{
			IntentID:     item.IntentID,
			AssignmentID: item.AssignmentID,
			Agent:        item.Agent,
			ResultHash:   item.ResultHash,
			ProofHash:    item.ProofHash,
		}
	}
	return cryptoHelpers.ComputeItemsHash(cryptoItems)
}

// SignDigest signs the validation digest using the configured SDK signer.
// Useful for single validator testing or when aggregating signatures client-side.
func (s *ValidationService) SignDigest(digest [32]byte) ([]byte, error) {
	if s.signer == nil {
		return nil, errors.New("validation: signer not configured")
	}
	return s.signer.SignDigest(digest)
}

// SignValidation wraps ComputeDigest + SignDigest to complete Validation signing in one step.
// Simplifies the Validator signing flow for ValidateIntentsBySignatures scenarios.
func (s *ValidationService) SignValidation(bundle ValidationBundle) ([]byte, error) {
	digest, err := s.ComputeDigest(bundle)
	if err != nil {
		return nil, err
	}
	return s.SignDigest(digest)
}
