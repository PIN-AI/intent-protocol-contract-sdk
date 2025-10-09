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

// ValidateIntentsBySignatures submits validator bundles to finalize intents.
func (s *ValidationService) ValidateIntentsBySignatures(ctx context.Context, bundles []ValidationBundle) (*types.Transaction, error) {
	if len(bundles) == 0 {
		return nil, errors.New("validation: empty bundle batch")
	}

	contractBundles := make([]intentmanager.IIntentManagerValidationBundleData, len(bundles))

	for i, b := range bundles {
		if len(b.Validators) != len(b.Signatures) {
			return nil, errors.New("validation: validator/signature length mismatch")
		}
		contractBundles[i] = intentmanager.IIntentManagerValidationBundleData{
			IntentId:     b.IntentID,
			AssignmentId: b.AssignmentID,
			SubnetId:     b.SubnetID,
			Agent:        b.Agent,
			ResultHash:   b.ResultHash,
			ProofHash:    b.ProofHash,
			RootHeight:   b.RootHeight,
			RootHash:     b.RootHash,
			Validators:   b.Validators,
			Signatures:   b.Signatures,
		}
	}

	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.ValidateIntentsBySignatures(opts, contractBundles)
	})
}

// ComputeDigest builds the validation digest for a bundle.
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

// SignDigest signs the validation digest using the configured SDK signer.
// Useful for single validator testing or when aggregating signatures client-side.
func (s *ValidationService) SignDigest(digest [32]byte) ([]byte, error) {
	if s.signer == nil {
		return nil, errors.New("validation: signer not configured")
	}
	return s.signer.SignDigest(digest)
}

// SignValidation 封装 ComputeDigest + SignDigest，一步完成 Validation 签名。
// 简化 Validator 签名流程，适用于 ValidateIntentsBySignatures 场景。
func (s *ValidationService) SignValidation(bundle ValidationBundle) ([]byte, error) {
	digest, err := s.ComputeDigest(bundle)
	if err != nil {
		return nil, err
	}
	return s.SignDigest(digest)
}
