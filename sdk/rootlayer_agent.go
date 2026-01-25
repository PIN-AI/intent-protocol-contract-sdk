package sdk

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	pb "github.com/PIN-AI/pin-protocol-proto/rootlayer/proto"
	"google.golang.org/grpc"
)

// AgentSession represents an authenticated Agent connection to RootLayer Direct Mode.
// It wraps the long-lived AgentConnect stream and provides helpers for Heartbeat and result submission.
type AgentSession struct {
	AgentAddress string
	AgentID      string

	Stream grpc.ServerStreamingClient[pb.DirectIntentPush]

	relayer pb.RelayerServiceClient
}

// AgentConnectToRootLayer establishes the AgentConnect stream to RootLayer.
// The signature is generated using PIN_AGENT_CONNECT_V2 and includes agent_id.
func (c *Client) AgentConnectToRootLayer(ctx context.Context, agentID string, clientVersion string) (*AgentSession, error) {
	if c == nil {
		return nil, errors.New("sdk: nil client")
	}
	if c.RootLayer == nil || c.RootLayer.Relayer == nil {
		return nil, errors.New("sdk: RootLayerURL not configured")
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
	sig, ts, nonce, err := c.SignAgentConnectNow(agentAddr, agentIDNorm)
	if err != nil {
		return nil, err
	}

	stream, err := c.RootLayer.Relayer.AgentConnect(ctx, &pb.AgentConnectRequest{
		AgentAddress:  agentAddr.Hex(),
		Signature:     sig,
		ClientVersion: clientVersion,
		Timestamp:     ts,
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

// Recv blocks until the next DirectIntentPush arrives on the AgentConnect stream.
func (s *AgentSession) Recv() (*pb.DirectIntentPush, error) {
	if s == nil || s.Stream == nil {
		return nil, errors.New("sdk: nil agent session stream")
	}
	return s.Stream.Recv()
}

// Heartbeat sends a HeartbeatRequest for this session.
func (s *AgentSession) Heartbeat(ctx context.Context) error {
	if s == nil || s.relayer == nil {
		return errors.New("sdk: nil agent session")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	_, err := s.relayer.Heartbeat(ctx, &pb.HeartbeatRequest{
		AgentAddress: s.AgentAddress,
		Timestamp:    time.Now().Unix(),
		AgentId:      s.AgentID,
	})
	return err
}

// SubmitDirectResult submits an execution result to RootLayer.
func (s *AgentSession) SubmitDirectResult(ctx context.Context, req *pb.DirectResultRequest) (*pb.DirectResultResponse, error) {
	if s == nil || s.relayer == nil {
		return nil, errors.New("sdk: nil agent session")
	}
	if ctx == nil {
		ctx = context.Background()
	}
	if req == nil {
		return nil, errors.New("sdk: nil direct result request")
	}
	return s.relayer.SubmitDirectResult(ctx, req)
}

// SubmitDirectResultFromPush is a convenience helper that submits a result using fields from the push.
func (s *AgentSession) SubmitDirectResultFromPush(ctx context.Context, push *pb.DirectIntentPush, resultData []byte, success bool, errorMessage string) (*pb.DirectResultResponse, error) {
	if push == nil {
		return nil, errors.New("sdk: nil direct intent push")
	}
	return s.SubmitDirectResult(ctx, &pb.DirectResultRequest{
		IntentId:      push.IntentId,
		AgentAddress:  s.AgentAddress,
		Success:       success,
		ResultData:    resultData,
		ErrorMessage:  errorMessage,
		Timestamp:     time.Now().Unix(),
		TargetAgentId: push.TargetAgentId,
		SubnetId:      push.SubnetId,
	})
}
