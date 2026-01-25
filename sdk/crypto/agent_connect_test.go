package crypto

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestComputeAgentConnectDigest_Valid(t *testing.T) {
	input := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Timestamp:    big.NewInt(time.Now().Unix()),
		RandomNonce:  [32]byte{1, 2, 3, 4, 5},
		AgentID:      big.NewInt(1),
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

// TestComputeAgentConnectDigest_ServerCompatibility verifies our implementation
// matches the server-side digest calculation exactly
func TestComputeAgentConnectDigest_ServerCompatibility(t *testing.T) {
	// Use the same test data
	agentAddr := common.HexToAddress("0xABCDEF1234567890ABCDEF1234567890ABCDEF12")
	timestamp := int64(1704067200) // 2024-01-01 00:00:00 UTC
	randomNonce := [32]byte{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}
	agentID := big.NewInt(10)

	// Calculate digest using SDK
	input := AgentConnectInput{
		AgentAddress: agentAddr,
		Timestamp:    big.NewInt(timestamp),
		RandomNonce:  randomNonce,
		AgentID:      agentID,
	}
	sdkDigest, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("SDK digest failed: %v", err)
	}

	// Calculate digest using server method (direct byte concatenation)
	typeHashDef := "PIN_AGENT_CONNECT_V2(address,uint256,bytes32,uint256)"
	typeHash := crypto.Keccak256Hash([]byte(typeHashDef))
	serverDigest := crypto.Keccak256Hash(
		common.LeftPadBytes(typeHash[:], 32),
		common.LeftPadBytes(agentAddr.Bytes(), 32),
		common.LeftPadBytes(big.NewInt(timestamp).Bytes(), 32),
		randomNonce[:],
		common.LeftPadBytes(agentID.Bytes(), 32),
	)

	// Verify they match
	if sdkDigest != serverDigest {
		t.Errorf("Digest mismatch!\nSDK:    %x\nServer: %x", sdkDigest, serverDigest)
	}

	t.Logf("✓ SDK digest matches server digest: %x", sdkDigest)
}

// TestAgentConnectSignature_FullFlow tests the complete signature flow
// matching the server-side implementation
func TestAgentConnectSignature_FullFlow(t *testing.T) {
	// Generate key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatal("Failed to cast public key")
	}

	agentAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	timestamp := time.Now().Unix()
	randomNonce := [32]byte{0x01, 0x02, 0x03, 0x04, 0x05}

	// Compute digest using SDK
	input := AgentConnectInput{
		AgentAddress: agentAddr,
		Timestamp:    big.NewInt(timestamp),
		RandomNonce:  randomNonce,
		AgentID:      big.NewInt(1),
	}
	digest, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("ComputeAgentConnectDigest failed: %v", err)
	}

	// Apply EIP-191 prefix and sign
	ethSignedHash := crypto.Keccak256Hash(
		[]byte("\x19Ethereum Signed Message:\n32"),
		digest[:],
	)
	signature, err := crypto.Sign(ethSignedHash.Bytes(), privateKey)
	if err != nil {
		t.Fatalf("Signing failed: %v", err)
	}

	// Verify signature
	pubKey, err := crypto.SigToPub(ethSignedHash.Bytes(), signature)
	if err != nil {
		t.Fatalf("Failed to recover public key: %v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if recoveredAddr != agentAddr {
		t.Errorf("Address mismatch!\nExpected: %s\nRecovered: %s", agentAddr.Hex(), recoveredAddr.Hex())
	}

	t.Logf("✓ Signature verified successfully")
	t.Logf("  Agent Address: %s", agentAddr.Hex())
	t.Logf("  Digest: %x", digest)
	t.Logf("  Signature: %x", signature)
}

func TestComputeAgentConnectDigest_NilTimestamp(t *testing.T) {
	input := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x1234567890123456789012345678901234567890"),
		Timestamp:    nil,
		RandomNonce:  [32]byte{1, 2, 3, 4, 5},
		AgentID:      big.NewInt(1),
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
		AgentID:      big.NewInt(1),
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
		AgentID:      big.NewInt(1),
	}
	input2 := AgentConnectInput{
		AgentAddress: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		Timestamp:    timestamp,
		RandomNonce:  nonce,
		AgentID:      big.NewInt(1),
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
		AgentID:      big.NewInt(1),
	}
	input2 := AgentConnectInput{
		AgentAddress: agent,
		Timestamp:    big.NewInt(2000000),
		RandomNonce:  nonce,
		AgentID:      big.NewInt(1),
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
		AgentID:      big.NewInt(1),
	}
	input2 := AgentConnectInput{
		AgentAddress: agent,
		Timestamp:    timestamp,
		RandomNonce:  [32]byte{5, 4, 3, 2, 1},
		AgentID:      big.NewInt(1),
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
		AgentID:      big.NewInt(1),
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
		AgentID:      big.NewInt(1),
	}
	digest, err := ComputeAgentConnectDigest(input)
	if err != nil {
		t.Fatalf("ComputeAgentConnectDigest failed: %v", err)
	}
	if digest == [32]byte{} {
		t.Fatal("digest is zero")
	}
}
