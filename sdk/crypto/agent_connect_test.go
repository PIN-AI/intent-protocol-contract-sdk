package crypto

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func TestComputeAgentConnectDigest_Valid(t *testing.T) {
	input := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Timestamp:    big.NewInt(time.Now().Unix()),
		RandomNonce:  [32]byte{1, 2, 3, 4, 5},
	}
	digest, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("ComputeAgentConnectDigest failed: %v", err)
	}
	if digest == [32]byte{} {
		t.Fatal("digest is zero")
	}
	t.Logf("Digest: %x", digest)
}

func TestComputeAgentConnectDigest_NilTimestamp(t *testing.T) {
	input := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Timestamp:    nil,
		RandomNonce:  [32]byte{1, 2, 3, 4, 5},
	}
	_, err := ComputeAgentConnectDigest(input)
	if err == nil {
		t.Fatal("expected error for nil timestamp, got nil")
	}
	if err.Error() != "crypto: nil timestamp" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestComputeAgentConnectDigest_Deterministic(t *testing.T) {
	input := AgentConnectInput{
		AgentAddress: common.HexToAddress("0xABCDEF1234567890ABCDEF1234567890ABCDEF12"),
		Timestamp:    big.NewInt(1704067200), // 2024-01-01 00:00:00 UTC
		RandomNonce:  [32]byte{0xAA, 0xBB, 0xCC, 0xDD},
	}
	digest1, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("first call failed: %v", err)
	}
	digest2, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("second call failed: %v", err)
	}
	if digest1 != digest2 {
		t.Fatalf("digests differ: %x != %x", digest1, digest2)
	}
}

func TestComputeAgentConnectDigest_DifferentAgent(t *testing.T) {
	timestamp := big.NewInt(time.Now().Unix())
	nonce := [32]byte{1, 2, 3}
	input1 := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		Timestamp:    timestamp,
		RandomNonce:  nonce,
	}
	input2 := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Timestamp:    timestamp,
		RandomNonce:  nonce,
	}
	digest1, _ := ComputeAgentConnectDigest(input1)
	digest2, _ := ComputeAgentConnectDigest(input2)
	if digest1 == digest2 {
		t.Fatal("digests should differ for different agents")
	}
}

func TestComputeAgentConnectDigest_DifferentTimestamp(t *testing.T) {
	agent := common.HexToAddress("0x1234567890123456789012345678901234567890")
	nonce := [32]byte{1, 2, 3}
	input1 := AgentConnectInput{
		AgentAddress: agent,
		Timestamp:    big.NewInt(1000000),
		RandomNonce:  nonce,
	}
	input2 := AgentConnectInput{
		AgentAddress: agent,
		Timestamp:    big.NewInt(2000000),
		RandomNonce:  nonce,
	}
	digest1, _ := ComputeAgentConnectDigest(input1)
	digest2, _ := ComputeAgentConnectDigest(input2)
	if digest1 == digest2 {
		t.Fatal("digests should differ for different timestamps")
	}
}

func TestComputeAgentConnectDigest_DifferentNonce(t *testing.T) {
	agent := common.HexToAddress("0x1234567890123456789012345678901234567890")
	timestamp := big.NewInt(time.Now().Unix())
	input1 := AgentConnectInput{
		AgentAddress: agent,
		Timestamp:    timestamp,
		RandomNonce:  [32]byte{1, 2, 3, 4, 5},
	}
	input2 := AgentConnectInput{
		AgentAddress: agent,
		Timestamp:    timestamp,
		RandomNonce:  [32]byte{5, 4, 3, 2, 1},
	}
	digest1, _ := ComputeAgentConnectDigest(input1)
	digest2, _ := ComputeAgentConnectDigest(input2)
	if digest1 == digest2 {
		t.Fatal("digests should differ for different nonces")
	}
}

func TestComputeAgentConnectDigest_ZeroAddress(t *testing.T) {
	input := AgentConnectInput{
		AgentAddress: common.Address{},
		Timestamp:    big.NewInt(time.Now().Unix()),
		RandomNonce:  [32]byte{1, 2, 3},
	}
	digest, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("ComputeAgentConnectDigest failed: %v", err)
	}
	if digest == [32]byte{} {
		t.Fatal("digest is zero")
	}
}

func TestComputeAgentConnectDigest_EmptyNonce(t *testing.T) {
	input := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Timestamp:    big.NewInt(time.Now().Unix()),
		RandomNonce:  [32]byte{},
	}
	digest, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("ComputeAgentConnectDigest failed: %v", err)
	}
	if digest == [32]byte{} {
		t.Fatal("digest is zero")
	}
}
