package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	agentConnectTypeHash = crypto.Keccak256Hash([]byte("PIN_AGENT_CONNECT_V1(address,uint256,bytes32)"))

	// arguments: typehash, agent_address, timestamp, random_nonce
	agentConnectDigestArgs = abi.Arguments{
		{Type: typeBytes32},
		{Type: typeAddress},
		{Type: typeUint256},
		{Type: typeBytes32},
	}
)

// AgentConnectInput represents the data required to construct an AgentConnect digest.
type AgentConnectInput struct {
	AgentAddress common.Address
	Timestamp    *big.Int
	RandomNonce  [32]byte
}

// ComputeAgentConnectDigest computes the digest required for Agent connection signature.
// This digest is used when an Agent connects to PIN AI network for authentication.
func ComputeAgentConnectDigest(input AgentConnectInput) ([32]byte, error) {
	var zero [32]byte
	if input.Timestamp == nil {
		return zero, errors.New("crypto: nil timestamp")
	}
	encoded, err := agentConnectDigestArgs.Pack(
		agentConnectTypeHash,
		input.AgentAddress,
		input.Timestamp,
		input.RandomNonce,
	)
	if err != nil {
		return zero, err
	}
	return bytesToBytes32(crypto.Keccak256Hash(encoded).Bytes()), nil
}
