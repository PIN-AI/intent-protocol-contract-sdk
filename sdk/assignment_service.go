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

// AssignmentData captures the matcher-signed assignment payload.
type AssignmentData struct {
	AssignmentID [32]byte
	IntentID     [32]byte
	BidID        [32]byte
	Agent        common.Address
	Status       AssignmentStatus
	Matcher      common.Address
}

// SignedAssignment wraps an assignment together with its signature.
type SignedAssignment struct {
	Data      AssignmentData
	Signature []byte
}

// AssignmentService provides high level accessors for assignment flows on IntentManager.
type AssignmentService struct {
	backend      bind.ContractBackend
	txManager    *txmgr.Manager
	contract     *intentmanager.IntentManager
	signer       signer.Signer
	chainID      *big.Int
	contractAddr common.Address
}

// NewAssignmentService constructs an AssignmentService instance.
func NewAssignmentService(
	backend bind.ContractBackend,
	txm *txmgr.Manager,
	contract *intentmanager.IntentManager,
	sig signer.Signer,
	chainID *big.Int,
	contractAddr common.Address,
) *AssignmentService {
	return &AssignmentService{
		backend:      backend,
		txManager:    txm,
		contract:     contract,
		signer:       sig,
		chainID:      new(big.Int).Set(chainID),
		contractAddr: contractAddr,
	}
}

// AssignIntentsBySignatures submits matcher-signed assignments in batch.
func (s *AssignmentService) AssignIntentsBySignatures(ctx context.Context, assignments []SignedAssignment) (*types.Transaction, error) {
	if len(assignments) == 0 {
		return nil, errors.New("assignment: empty batch")
	}

	contractAssignments := make([]intentmanager.IIntentManagerAssignmentData, len(assignments))
	signatures := make([][]byte, len(assignments))

	for i, item := range assignments {
		contractAssignments[i] = intentmanager.IIntentManagerAssignmentData{
			AssignmentId: item.Data.AssignmentID,
			IntentId:     item.Data.IntentID,
			BidId:        item.Data.BidID,
			Agent:        item.Data.Agent,
			Status:       uint8(item.Data.Status),
			Matcher:      item.Data.Matcher,
		}
		signatures[i] = item.Signature
	}

	return s.txManager.Send(ctx, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		opts.Context = ctx
		return s.contract.AssignIntentsBySignatures(opts, contractAssignments, signatures)
	})
}

// ComputeDigest calculates the assignment digest for matcher signatures.
func (s *AssignmentService) ComputeDigest(data AssignmentData) ([32]byte, error) {
	return cryptoHelpers.ComputeAssignmentDigest(cryptoHelpers.AssignmentInput{
		AssignmentID: data.AssignmentID,
		IntentID:     data.IntentID,
		BidID:        data.BidID,
		Agent:        data.Agent,
		Status:       uint8(data.Status),
		Matcher:      data.Matcher,
	}, s.contractAddr, s.chainID)
}

// SignDigest signs the provided digest using the configured SDK signer.
func (s *AssignmentService) SignDigest(digest [32]byte) ([]byte, error) {
	if s.signer == nil {
		return nil, errors.New("assignment: signer not configured")
	}
	return s.signer.SignDigest(digest)
}

// SignAssignment wraps ComputeDigest + SignDigest to complete Assignment signing in one step.
// Simplifies the Matcher batch assignment flow for AssignIntentsBySignatures scenarios.
func (s *AssignmentService) SignAssignment(data AssignmentData) ([]byte, error) {
	digest, err := s.ComputeDigest(data)
	if err != nil {
		return nil, err
	}
	return s.SignDigest(digest)
}
