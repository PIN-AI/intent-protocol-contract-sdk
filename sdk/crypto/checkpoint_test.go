package crypto

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function: generates test CommitmentRoots
func generateTestRoots() CommitmentRoots {
	return CommitmentRoots{
		AgentRoot:        [32]byte{0x01},
		AgentServiceRoot: [32]byte{0x02},
		RankRoot:         [32]byte{0x03},
		MetricsRoot:      [32]byte{0x04},
		DataUsageRoot:    [32]byte{0x05},
		StateRoot:        [32]byte{0x06},
		EventRoot:        [32]byte{0x07},
		CrossSubnetRoot:  [32]byte{0x08},
	}
}

// Helper function: generates test DACommitments
func generateTestDACommitments() []DACommitment {
	return []DACommitment{
		{
			Kind:          "ipfs",
			Pointer:       "QmTest123",
			SizeHint:      big.NewInt(1024),
			SegmentHashes: [][32]byte{{0x01}, {0x02}},
			Expiry:        big.NewInt(1234567890),
		},
		{
			Kind:          "arweave",
			Pointer:       "ArTest456",
			SizeHint:      big.NewInt(2048),
			SegmentHashes: [][32]byte{{0x03}},
			Expiry:        big.NewInt(9876543210),
		},
	}
}

func TestEncodeCommitmentRoots(t *testing.T) {
	roots := generateTestRoots()
	encoded := EncodeCommitmentRoots(roots)

	// Verify length: 8 roots * 32 bytes = 256 bytes
	assert.Equal(t, 256, len(encoded), "Should be 256 bytes after encoding")

	// Verify each root's position
	assert.Equal(t, roots.AgentRoot[:], encoded[0:32], "AgentRoot position error")
	assert.Equal(t, roots.AgentServiceRoot[:], encoded[32:64], "AgentServiceRoot position error")
	assert.Equal(t, roots.RankRoot[:], encoded[64:96], "RankRoot position error")
	assert.Equal(t, roots.MetricsRoot[:], encoded[96:128], "MetricsRoot position error")
	assert.Equal(t, roots.DataUsageRoot[:], encoded[128:160], "DataUsageRoot position error")
	assert.Equal(t, roots.StateRoot[:], encoded[160:192], "StateRoot position error")
	assert.Equal(t, roots.EventRoot[:], encoded[192:224], "EventRoot position error")
	assert.Equal(t, roots.CrossSubnetRoot[:], encoded[224:256], "CrossSubnetRoot position error")
}

func TestEncodeCommitmentRoots_EmptyRoots(t *testing.T) {
	roots := CommitmentRoots{} // all zero values
	encoded := EncodeCommitmentRoots(roots)

	assert.Equal(t, 256, len(encoded))
	// All bytes should be 0
	for i, b := range encoded {
		assert.Equal(t, byte(0), b, "index %d should be 0", i)
	}
}

func TestEncodeDACommitments_Empty(t *testing.T) {
	commitments := []DACommitment{}
	encoded, err := EncodeDACommitments(commitments)

	require.NoError(t, err)
	assert.NotEmpty(t, encoded, "empty array should also produce valid encoding")
}

func TestEncodeDACommitments_Single(t *testing.T) {
	commitments := []DACommitment{
		{
			Kind:          "ipfs",
			Pointer:       "QmTest123",
			SizeHint:      big.NewInt(1024),
			SegmentHashes: [][32]byte{{0x01}},
			Expiry:        big.NewInt(1234567890),
		},
	}

	encoded, err := EncodeDACommitments(commitments)

	require.NoError(t, err)
	assert.NotEmpty(t, encoded)
}

func TestEncodeDACommitments_Multiple(t *testing.T) {
	commitments := generateTestDACommitments()
	encoded, err := EncodeDACommitments(commitments)

	require.NoError(t, err)
	assert.NotEmpty(t, encoded)
}

func TestEncodeDACommitments_NilBigInt(t *testing.T) {
	commitments := []DACommitment{
		{
			Kind:          "ipfs",
			Pointer:       "QmTest",
			SizeHint:      nil, // nil should be handled by normalizeBigInt as 0
			SegmentHashes: [][32]byte{{0x01}},
			Expiry:        nil, // nil should be handled by normalizeBigInt as 0
		},
	}

	encoded, err := EncodeDACommitments(commitments)

	require.NoError(t, err, "nil big.Int should be handled correctly")
	assert.NotEmpty(t, encoded)
}

func TestEncodeEpochSlot(t *testing.T) {
	slot := EpochSlot{
		Epoch: 12345,
		Slot:  67890,
	}

	encoded := EncodeEpochSlot(slot)

	// Verify length: 2 uint64 * 32 bytes (left-padded) = 64 bytes
	assert.Equal(t, 64, len(encoded), "Should be 64 bytes after encoding")

	// Verify each field occupies 32 bytes
	epochBytes := encoded[0:32]
	slotBytes := encoded[32:64]

	assert.NotEmpty(t, epochBytes)
	assert.NotEmpty(t, slotBytes)
}

func TestEncodeEpochSlot_Zero(t *testing.T) {
	slot := EpochSlot{Epoch: 0, Slot: 0}
	encoded := EncodeEpochSlot(slot)

	assert.Equal(t, 64, len(encoded))
	// All bytes should be 0
	for i, b := range encoded {
		assert.Equal(t, byte(0), b, "Index %d should be 0", i)
	}
}

func TestComputeCheckpointDigest_Valid(t *testing.T) {
	input := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(time.Now().Unix()),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        generateTestDACommitments(),
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte(`{"key":"value"}`),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")
	chainID := big.NewInt(84532)

	digest, err := ComputeCheckpointDigest(input, contract, chainID)

	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, digest)
	assert.Len(t, digest, 32)
}

func TestComputeCheckpointDigest_NilChainID(t *testing.T) {
	input := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(time.Now().Unix()),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        []DACommitment{},
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte("{}"),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")

	digest, err := ComputeCheckpointDigest(input, contract, nil)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil chain id", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeCheckpointDigest_NilTimestamp(t *testing.T) {
	input := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            nil,
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        []DACommitment{},
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte("{}"),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")
	chainID := big.NewInt(84532)

	digest, err := ComputeCheckpointDigest(input, contract, chainID)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil timestamp", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeCheckpointDigest_Deterministic(t *testing.T) {
	input := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(1234567890),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        generateTestDACommitments(),
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte(`{"key":"value"}`),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")
	chainID := big.NewInt(84532)

	digest1, err1 := ComputeCheckpointDigest(input, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input, contract, chainID)
	require.NoError(t, err2)

	assert.Equal(t, digest1, digest2, "same input should produce the same digest")
}

func TestComputeCheckpointDigest_DifferentRoots(t *testing.T) {
	baseInput := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(1234567890),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		DACommitments:        []DACommitment{},
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte("{}"),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.Roots = generateTestRoots()

	input2 := baseInput
	roots2 := generateTestRoots()
	roots2.AgentRoot = [32]byte{0xFF} // modify one root
	input2.Roots = roots2

	digest1, err1 := ComputeCheckpointDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different Roots should produce different digest")
}

func TestComputeCheckpointDigest_DifferentDACommitments(t *testing.T) {
	baseInput := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(1234567890),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte("{}"),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.DACommitments = []DACommitment{}

	input2 := baseInput
	input2.DACommitments = generateTestDACommitments()

	digest1, err1 := ComputeCheckpointDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different DACommitments should produce different digest")
}

func TestComputeCheckpointDigest_DifferentEpoch(t *testing.T) {
	baseInput := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(1234567890),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        []DACommitment{},
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte("{}"),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.Epoch = 12345

	input2 := baseInput
	input2.Epoch = 99999

	digest1, err1 := ComputeCheckpointDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different Epoch should produce different digest")
}

func TestComputeCheckpointDigest_DifferentStats(t *testing.T) {
	baseInput := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(1234567890),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        []DACommitment{},
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.Stats = []byte(`{"key":"value1"}`)

	input2 := baseInput
	input2.Stats = []byte(`{"key":"value2"}`)

	digest1, err1 := ComputeCheckpointDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different Stats should produce different digest (via statsHash)")
}

func TestComputeCheckpointDigest_DifferentChainID(t *testing.T) {
	input := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(1234567890),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        []DACommitment{},
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte("{}"),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	contract := common.HexToAddress("0xd4444444444444444444444444444444444444")

	digest1, err1 := ComputeCheckpointDigest(input, contract, big.NewInt(84532))
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input, contract, big.NewInt(11155111))
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different chain_id should produce different digest (replay protection)")
}

func TestComputeCheckpointDigest_DifferentContract(t *testing.T) {
	input := CheckpointInput{
		SubnetID:             [32]byte{0x01},
		Epoch:                12345,
		ParentCheckpointHash: [32]byte{0x02},
		Timestamp:            big.NewInt(1234567890),
		Version:              1,
		ParamsHash:           [32]byte{0x03},
		Roots:                generateTestRoots(),
		DACommitments:        []DACommitment{},
		EpochSlot:            EpochSlot{Epoch: 12345, Slot: 0},
		Stats:                []byte("{}"),
		AuxHash:              [32]byte{0x09},
		AssignmentsRoot:      [32]byte{0x0A},
		ValidationCommitment: [32]byte{0x0B},
		PolicyRoot:           [32]byte{0x0C},
	}
	chainID := big.NewInt(84532)

	contract1 := common.HexToAddress("0xd1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1")
	contract2 := common.HexToAddress("0xd2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2d2")

	digest1, err1 := ComputeCheckpointDigest(input, contract1, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input, contract2, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different contract should produce different digest (cross-contract replay protection)")
}

func TestNormalizeBigInt(t *testing.T) {
	// Test non-nil value
	value := big.NewInt(12345)
	result := normalizeBigInt(value)
	assert.Equal(t, value, result, "non-nil value should be returned as is")

	// Test nil value
	nilResult := normalizeBigInt(nil)
	assert.NotNil(t, nilResult, "nil should be converted to non-nil")
	assert.Equal(t, big.NewInt(0), nilResult, "nil should be converted to 0")
}
