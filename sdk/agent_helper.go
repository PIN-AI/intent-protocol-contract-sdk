package sdk

import (
	"crypto/rand"
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
//   - timestamp: Unix timestamp for the connection (use time.Now().Unix() for current time)
//   - randomNonce: 32-byte random nonce for replay protection (if nil, auto-generates)
//
// Returns the signature bytes and any error encountered.
func (c *Client) SignAgentConnect(agentAddress common.Address, timestamp int64, randomNonce *[32]byte) ([]byte, error) {
	var nonce [32]byte
	if randomNonce != nil {
		nonce = *randomNonce
	} else {
		// Auto-generate random nonce
		if _, err := rand.Read(nonce[:]); err != nil {
			return nil, err
		}
	}

	input := cryptoHelpers.AgentConnectInput{
		AgentAddress: agentAddress,
		Timestamp:    big.NewInt(timestamp),
		RandomNonce:  nonce,
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
// Equivalent to: client.SignAgentConnect(agentAddress, time.Now().Unix(), nil)
func (c *Client) SignAgentConnectNow(agentAddress common.Address) (signature []byte, timestamp int64, nonce [32]byte, err error) {
	timestamp = time.Now().Unix()
	if _, err = rand.Read(nonce[:]); err != nil {
		return
	}

	input := cryptoHelpers.AgentConnectInput{
		AgentAddress: agentAddress,
		Timestamp:    big.NewInt(timestamp),
		RandomNonce:  nonce,
	}

	var digest [32]byte
	digest, err = cryptoHelpers.ComputeAgentConnectDigest(input)
	if err != nil {
		return
	}

	signature, err = c.Signer.SignDigest(digest)
	return
}
