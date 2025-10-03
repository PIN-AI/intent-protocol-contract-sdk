package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	assignmentTypeHash = crypto.Keccak256Hash([]byte("PIN_ASSIGNMENT_V1(bytes32,bytes32,bytes32,address,uint8,address,address,uint256)"))

	assignmentDigestArgs = abi.Arguments{
		{Type: typeBytes32}, // typeHash
		{Type: typeBytes32}, // AssignmentID
		{Type: typeBytes32}, // IntentID
		{Type: typeBytes32}, // BidID
		{Type: typeAddress}, // Agent
		{Type: typeUint8},   // Status
		{Type: typeAddress}, // Matcher
		{Type: typeAddress}, // contract
		{Type: typeUint256}, // chainID
	}
)

// AssignmentInput captures the data required to compute the assignment digest.
type AssignmentInput struct {
	AssignmentID [32]byte
	IntentID     [32]byte
	BidID        [32]byte
	Agent        common.Address
	Status       uint8
	Matcher      common.Address
}

// ComputeAssignmentDigest returns the keccak256 hash for a given assignment input.
func ComputeAssignmentDigest(input AssignmentInput, contract common.Address, chainID *big.Int) ([32]byte, error) {
	var zero [32]byte
	if chainID == nil {
		return zero, errors.New("crypto: nil chain id")
	}

	encoded, err := assignmentDigestArgs.Pack(
		assignmentTypeHash,
		input.AssignmentID,
		input.IntentID,
		input.BidID,
		input.Agent,
		input.Status,
		input.Matcher,
		contract,
		chainID,
	)
	if err != nil {
		return zero, err
	}

	return bytesToBytes32(crypto.Keccak256Hash(encoded).Bytes()), nil
}
