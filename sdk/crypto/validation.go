package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	validationTypeHash = crypto.Keccak256Hash([]byte("PIN_VALIDATION_V1(bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address,uint256)"))

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

// ComputeValidationDigest returns the keccak256 hash for a given validation input.
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
