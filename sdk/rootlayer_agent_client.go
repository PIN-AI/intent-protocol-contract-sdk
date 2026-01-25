package sdk

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	sdkcrypto "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
	"github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"
	pb "github.com/PIN-AI/pin-protocol-proto/rootlayer/proto"
)

// RootLayerAgentClient is a lightweight client for Agent-side Direct Mode interactions.
// It does NOT require any on-chain RPC configuration.
type RootLayerAgentClient struct {
	RootLayer *RootLayerClient
	Signer    signer.Signer
}

func NewRootLayerAgentClient(ctx context.Context, rootLayerURL string, signing signer.Signer) (*RootLayerAgentClient, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if strings.TrimSpace(rootLayerURL) == "" {
		return nil, errors.New("sdk: RootLayerURL is required")
	}
	if signing == nil {
		return nil, errors.New("sdk: signer is required")
	}

	rl, err := newRootLayerClient(ctx, rootLayerURL)
	if err != nil {
		return nil, err
	}
	return &RootLayerAgentClient{RootLayer: rl, Signer: signing}, nil
}

func (c *RootLayerAgentClient) Close() error {
	if c == nil || c.RootLayer == nil {
		return nil
	}
	return c.RootLayer.Close()
}

// AgentConnect establishes the AgentConnect stream to RootLayer.
func (c *RootLayerAgentClient) AgentConnect(ctx context.Context, agentID string, clientVersion string) (*AgentSession, error) {
	if c == nil {
		return nil, errors.New("sdk: nil rootlayer agent client")
	}
	if c.RootLayer == nil || c.RootLayer.Relayer == nil {
		return nil, errors.New("sdk: RootLayerURL not configured")
	}
	if c.Signer == nil {
		return nil, errors.New("sdk: signer is required")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if strings.TrimSpace(clientVersion) == "" {
		clientVersion = "intent-protocol-contract-sdk"
	}

	agentIDInt, err := parseUint256(agentID)
	if err != nil {
		return nil, fmt.Errorf("sdk: parse agent_id: %w", err)
	}
	if agentIDInt.BitLen() > 256 {
		return nil, errors.New("sdk: agent_id exceeds uint256 range")
	}
	agentIDNorm := agentIDInt.String()

	agentAddr := c.Signer.Address()
	timestamp := time.Now().Unix()
	var nonce [32]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return nil, fmt.Errorf("sdk: generate nonce: %w", err)
	}

	digest, err := sdkcrypto.ComputeAgentConnectDigest(sdkcrypto.AgentConnectInput{
		AgentAddress: agentAddr,
		Timestamp:    big.NewInt(timestamp),
		RandomNonce:  nonce,
		AgentID:      agentIDInt,
	})
	if err != nil {
		return nil, err
	}

	sig, err := c.Signer.SignDigest(digest)
	if err != nil {
		return nil, err
	}

	stream, err := c.RootLayer.Relayer.AgentConnect(ctx, &pb.AgentConnectRequest{
		AgentAddress:  agentAddr.Hex(),
		Signature:     sig,
		ClientVersion: clientVersion,
		Timestamp:     timestamp,
		RandomNonce:   nonce[:],
		AgentId:       agentIDNorm,
	})
	if err != nil {
		return nil, err
	}

	return &AgentSession{
		AgentAddress: agentAddr.Hex(),
		AgentID:      agentIDNorm,
		Stream:       stream,
		relayer:      c.RootLayer.Relayer,
	}, nil
}
