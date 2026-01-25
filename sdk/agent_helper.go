package sdk

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	cryptoHelpers "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
)

// SignAgentConnect generates a signature for Agent connection to PIN AI network.
// This is used for Agent authentication when connecting to the network.
//
// Parameters:
//   - agentAddress: The Agent's wallet address (typically client.Signer.Address())
//   - agentID: ERC-8004 agent id (tokenId) as a uint256 string
//   - timestamp: Unix timestamp for the connection (use time.Now().Unix() for current time)
//   - randomNonce: 32-byte random nonce for replay protection (if nil, auto-generates)
//
// Returns the signature bytes and any error encountered.
func (c *Client) SignAgentConnect(agentAddress common.Address, agentID string, timestamp int64, randomNonce *[32]byte) ([]byte, error) {
	var nonce [32]byte
	if randomNonce != nil {
		nonce = *randomNonce
	} else {
		// Auto-generate random nonce
		if _, err := rand.Read(nonce[:]); err != nil {
			return nil, err
		}
	}

	agentIDInt, err := parseUint256(agentID)
	if err != nil {
		return nil, fmt.Errorf("sdk: parse agent_id: %w", err)
	}
	if agentIDInt.BitLen() > 256 {
		return nil, errors.New("sdk: agent_id exceeds uint256 range")
	}

	input := cryptoHelpers.AgentConnectInput{
		AgentAddress: agentAddress,
		Timestamp:    big.NewInt(timestamp),
		RandomNonce:  nonce,
		AgentID:      agentIDInt,
	}

	digest, err := cryptoHelpers.ComputeAgentConnectDigest(input)
	if err != nil {
		return nil, err
	}

	return c.Signer.SignDigest(digest)
}

// SignAgentConnectNow is a convenience wrapper that signs an Agent connection
// using the current timestamp and auto-generated random nonce.
//
// Equivalent to: client.SignAgentConnect(agentAddress, agentID, time.Now().Unix(), nil)
func (c *Client) SignAgentConnectNow(agentAddress common.Address, agentID string) (signature []byte, timestamp int64, nonce [32]byte, err error) {
	timestamp = time.Now().Unix()
	if _, err = rand.Read(nonce[:]); err != nil {
		return
	}

	agentIDInt, err := parseUint256(agentID)
	if err != nil {
		err = fmt.Errorf("sdk: parse agent_id: %w", err)
		return
	}
	if agentIDInt.BitLen() > 256 {
		err = errors.New("sdk: agent_id exceeds uint256 range")
		return
	}

	input := cryptoHelpers.AgentConnectInput{
		AgentAddress: agentAddress,
		Timestamp:    big.NewInt(timestamp),
		RandomNonce:  nonce,
		AgentID:      agentIDInt,
	}

	var digest [32]byte
	digest, err = cryptoHelpers.ComputeAgentConnectDigest(input)
	if err != nil {
		return
	}

	signature, err = c.Signer.SignDigest(digest)
	return
}
