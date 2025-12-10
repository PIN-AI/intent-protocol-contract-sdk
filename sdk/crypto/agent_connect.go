package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	agentConnectTypeHash = crypto.Keccak256Hash([]byte("PIN_AGENT_CONNECT_V1(address,uint256,bytes32)"))
)

// AgentConnectInput represents the data required to construct an AgentConnect digest.
type AgentConnectInput struct {
	AgentAddress common.Address
	Timestamp    *big.Int
	RandomNonce  [32]byte
}

// ComputeAgentConnectDigest computes the digest required for Agent connection signature.
// This digest is used when an Agent connects to PIN AI network for authentication.
//
// The digest is computed using direct byte concatenation (not ABI encoding) to match
// the server-side implementation:
//   digest = keccak256(
//       leftPad(typeHash, 32) ||
//       leftPad(agentAddress, 32) ||
//       leftPad(timestamp, 32) ||
//       randomNonce
//   )
func ComputeAgentConnectDigest(input AgentConnectInput) ([32]byte, error) {
	var zero [32]byte
	if input.Timestamp == nil {
		return zero, errors.New("crypto: nil timestamp")
	}

	// Direct byte concatenation matching server implementation
	digest := crypto.Keccak256Hash(
		common.LeftPadBytes(agentConnectTypeHash[:], 32),
		common.LeftPadBytes(input.AgentAddress.Bytes(), 32),
		common.LeftPadBytes(input.Timestamp.Bytes(), 32),
		input.RandomNonce[:],
	)

	return digest, nil
}
