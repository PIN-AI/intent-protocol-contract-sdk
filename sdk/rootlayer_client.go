package sdk

import (
	"context"
	"errors"
	"strings"

	pb "github.com/PIN-AI/pin-protocol-proto/rootlayer/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RootLayerClient wraps gRPC stubs for RootLayer APIs.
type RootLayerClient struct {
	conn *grpc.ClientConn

	IntentPool   pb.IntentPoolServiceClient
	DirectIntent pb.DirectIntentServiceClient
	Relayer      pb.RelayerServiceClient
}

func newRootLayerClient(ctx context.Context, rootLayerURL string) (*RootLayerClient, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	endpoint := strings.TrimSpace(rootLayerURL)
	if endpoint == "" {
		return nil, errors.New("sdk: RootLayerURL is required")
	}

	conn, err := grpc.DialContext(ctx, endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &RootLayerClient{
		conn:         conn,
		IntentPool:   pb.NewIntentPoolServiceClient(conn),
		DirectIntent: pb.NewDirectIntentServiceClient(conn),
		Relayer:      pb.NewRelayerServiceClient(conn),
	}, nil
}

func (c *RootLayerClient) Close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}
