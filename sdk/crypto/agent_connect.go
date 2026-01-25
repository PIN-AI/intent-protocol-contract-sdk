package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	agentConnectTypeHash = crypto.Keccak256Hash([]byte("PIN_AGENT_CONNECT_V2(address,uint256,bytes32,uint256)"))
)

// AgentConnectInput represents the data required to construct an AgentConnect digest.
type AgentConnectInput struct {
	AgentAddress common.Address
	Timestamp    *big.Int
	RandomNonce  [32]byte
	AgentID      *big.Int
}

// ComputeAgentConnectDigest computes the digest required for Agent connection signature.
// This digest is used when an Agent connects to PIN AI network for authentication.
//
// The digest is computed using direct byte concatenation (not ABI encoding) to match
// the server-side implementation:
//
//	digest = keccak256(
//	    leftPad(typeHash, 32) ||
//	    leftPad(agentAddress, 32) ||
//	    leftPad(timestamp, 32) ||
//	    randomNonce ||
//	    leftPad(agent_id, 32)
//	)
func ComputeAgentConnectDigest(input AgentConnectInput) ([32]byte, error) {
	var zero [32]byte
	if input.Timestamp == nil {
		return zero, errors.New("crypto: nil timestamp")
	}
	if input.AgentID == nil {
		return zero, errors.New("crypto: nil agent_id")
	}
	if input.AgentID.Sign() < 0 {
		return zero, errors.New("crypto: agent_id must be non-negative")
	}
	if input.AgentID.BitLen() > 256 {
		return zero, errors.New("crypto: agent_id exceeds uint256 range")
	}

	// Direct byte concatenation matching server implementation
	digest := crypto.Keccak256Hash(
		common.LeftPadBytes(agentConnectTypeHash[:], 32),
		common.LeftPadBytes(input.AgentAddress.Bytes(), 32),
		common.LeftPadBytes(input.Timestamp.Bytes(), 32),
		input.RandomNonce[:],
		common.LeftPadBytes(input.AgentID.Bytes(), 32),
	)

	return digest, nil
}
