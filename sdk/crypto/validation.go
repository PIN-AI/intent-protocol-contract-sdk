package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	// ValidationTypeHash for single validation (v2.2 and earlier)
	validationTypeHash = crypto.Keccak256Hash([]byte("PIN_VALIDATION_V1(bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address,uint256)"))

	// ValidationBatchTypeHash for batch validation (v2.3+)
	validationBatchTypeHash = crypto.Keccak256Hash([]byte("PIN_VALIDATION_BATCH_V1(bytes32,bytes32,uint64,bytes32,address,uint256)"))

	validationDigestArgs = abi.Arguments{
		{Type: typeBytes32}, // typeHash
		{Type: typeBytes32}, // IntentID
		{Type: typeBytes32}, // AssignmentID
		{Type: typeBytes32}, // SubnetID
		{Type: typeAddress}, // Agent
		{Type: typeBytes32}, // ResultHash
		{Type: typeBytes32}, // ProofHash
		{Type: typeUint64},  // RootHeight
		{Type: typeBytes32}, // RootHash
		{Type: typeAddress}, // contract
		{Type: typeUint256}, // chainID
	}

	validationBatchDigestArgs = abi.Arguments{
		{Type: typeBytes32}, // typeHash
		{Type: typeBytes32}, // SubnetID
		{Type: typeBytes32}, // ItemsHash
		{Type: typeUint64},  // RootHeight
		{Type: typeBytes32}, // RootHash
		{Type: typeAddress}, // contract
		{Type: typeUint256}, // chainID
	}

	validationItemArgs = abi.Arguments{
		{Type: typeBytes32}, // IntentID
		{Type: typeBytes32}, // AssignmentID
		{Type: typeAddress}, // Agent
		{Type: typeBytes32}, // ResultHash
		{Type: typeBytes32}, // ProofHash
	}
)

// ValidationInput models the payload a validator signs for a validation bundle.
type ValidationInput struct {
	IntentID     [32]byte
	AssignmentID [32]byte
	SubnetID     [32]byte
	Agent        common.Address
	ResultHash   [32]byte
	ProofHash    [32]byte
	RootHeight   uint64
	RootHash     [32]byte
}

// ValidationBatchInput models the payload for batch validation (v2.3+).
type ValidationBatchInput struct {
	SubnetID   [32]byte
	ItemsHash  [32]byte
	RootHeight uint64
	RootHash   [32]byte
}

// ValidationItem represents a single validation item within a batch.
type ValidationItem struct {
	IntentID     [32]byte
	AssignmentID [32]byte
	Agent        common.Address
	ResultHash   [32]byte
	ProofHash    [32]byte
}

// ComputeValidationDigest returns the keccak256 hash for a given validation input (single validation, v2.2 and earlier).
func ComputeValidationDigest(input ValidationInput, contract common.Address, chainID *big.Int) ([32]byte, error) {
	var zero [32]byte
	if chainID == nil {
		return zero, errors.New("crypto: nil chain id")
	}

	encoded, err := validationDigestArgs.Pack(
		validationTypeHash,
		input.IntentID,
		input.AssignmentID,
		input.SubnetID,
		input.Agent,
		input.ResultHash,
		input.ProofHash,
		input.RootHeight,
		input.RootHash,
		contract,
		chainID,
	)
	if err != nil {
		return zero, err
	}

	return bytesToBytes32(crypto.Keccak256Hash(encoded).Bytes()), nil
}

// ComputeItemsHash calculates keccak256(abi.encode(items)) for batch validation.
// This must match Solidity's abi.encode behavior for the items array.
func ComputeItemsHash(items []ValidationItem) ([32]byte, error) {
	var zero [32]byte
	if len(items) == 0 {
		return zero, errors.New("crypto: empty items array")
	}

	// Solidity's abi.encode for struct array manually constructed:
	// - 32 bytes: offset to array data (0x20)
	// - 32 bytes: array length
	// - N * 160 bytes: each item data (5 fields × 32 bytes each)
	var encoded []byte

	// 1. Offset to array data (always 0x20 for single parameter)
	offset := make([]byte, 32)
	offset[31] = 0x20
	encoded = append(encoded, offset...)

	// 2. Array length
	lengthBytes := make([]byte, 32)
	lengthBig := big.NewInt(int64(len(items)))
	lengthBig.FillBytes(lengthBytes)
	encoded = append(encoded, lengthBytes...)

	// 3. Each item's fields (all 32-byte aligned)
	for _, item := range items {
		encoded = append(encoded, item.IntentID[:]...)                         // bytes32
		encoded = append(encoded, item.AssignmentID[:]...)                     // bytes32
		encoded = append(encoded, common.LeftPadBytes(item.Agent.Bytes(), 32)...) // address (left-padded)
		encoded = append(encoded, item.ResultHash[:]...)                       // bytes32
		encoded = append(encoded, item.ProofHash[:]...)                        // bytes32
	}

	return bytesToBytes32(crypto.Keccak256Hash(encoded).Bytes()), nil
}

// ComputeValidationBatchDigest returns the keccak256 hash for batch validation (v2.3+).
func ComputeValidationBatchDigest(input ValidationBatchInput, contract common.Address, chainID *big.Int) ([32]byte, error) {
	var zero [32]byte
	if chainID == nil {
		return zero, errors.New("crypto: nil chain id")
	}

	encoded, err := validationBatchDigestArgs.Pack(
		validationBatchTypeHash,
		input.SubnetID,
		input.ItemsHash,
		input.RootHeight,
		input.RootHash,
		contract,
		chainID,
	)
	if err != nil {
		return zero, err
	}

	return bytesToBytes32(crypto.Keccak256Hash(encoded).Bytes()), nil
}
