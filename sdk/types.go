package sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// IntentStatus mirrors DataStructures.IntentStatus in the contracts.
type IntentStatus uint8

const (
	IntentStatusPending    IntentStatus = 0
	IntentStatusProcessing IntentStatus = 1
	IntentStatusCompleted  IntentStatus = 2
	IntentStatusExpired    IntentStatus = 3
	IntentStatusFail       IntentStatus = 4
)

// AssignmentStatus mirrors DataStructures.AssignmentStatus in the contracts.
type AssignmentStatus uint8

const (
	AssignmentStatusUnspecified AssignmentStatus = 0
	AssignmentStatusActive      AssignmentStatus = 1
	AssignmentStatusFailed      AssignmentStatus = 2
)

// AssignmentInfo is a high-level representation of assignment metadata returned by the contract.
type AssignmentInfo struct {
	AssignmentID [32]byte
	IntentID     [32]byte
	BidID        [32]byte
	Agent        common.Address
	Status       AssignmentStatus
	Matcher      common.Address
	AssignedAt   *big.Int
}

// ValidationItem represents a single validation item (without signatures) within a batch.
// Corresponds to IIntentManager.ValidationItemData in contracts.
type ValidationItem struct {
	IntentID     [32]byte
	AssignmentID [32]byte
	Agent        common.Address
	ResultHash   [32]byte
	ProofHash    [32]byte
}

// ValidationBatch represents batch validation data with shared signatures.
// Corresponds to IIntentManager.ValidationBatchData in contracts.
type ValidationBatch struct {
	SubnetID   [32]byte
	ItemsHash  [32]byte          // keccak256(abi.encode(items)), computed by SDK
	RootHeight uint64
	RootHash   [32]byte
	Items      []ValidationItem  // Multiple validation items
	Validators []common.Address  // Shared validator addresses
	Signatures [][]byte          // Shared signatures
}

// ValidationBundle captures the data required to submit a single validation on-chain (backward compatible).
// Corresponds to IIntentManager.ValidationBundleData in contracts.
type ValidationBundle struct {
	IntentID     [32]byte
	AssignmentID [32]byte
	SubnetID     [32]byte
	Agent        common.Address
	ResultHash   [32]byte
	ProofHash    [32]byte
	RootHeight   uint64
	RootHash     [32]byte
	Validators   []common.Address
	Signatures   [][]byte
}
