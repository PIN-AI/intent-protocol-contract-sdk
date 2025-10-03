package crypto

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// 辅助函数：生成测试用的CommitmentRoots
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

// 辅助函数：生成测试用的DACommitment
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

	// 验证长度：8个根 * 32字节 = 256字节
	assert.Equal(t, 256, len(encoded), "编码后应为256字节")

	// 验证每个根的位置
	assert.Equal(t, roots.AgentRoot[:], encoded[0:32], "AgentRoot位置错误")
	assert.Equal(t, roots.AgentServiceRoot[:], encoded[32:64], "AgentServiceRoot位置错误")
	assert.Equal(t, roots.RankRoot[:], encoded[64:96], "RankRoot位置错误")
	assert.Equal(t, roots.MetricsRoot[:], encoded[96:128], "MetricsRoot位置错误")
	assert.Equal(t, roots.DataUsageRoot[:], encoded[128:160], "DataUsageRoot位置错误")
	assert.Equal(t, roots.StateRoot[:], encoded[160:192], "StateRoot位置错误")
	assert.Equal(t, roots.EventRoot[:], encoded[192:224], "EventRoot位置错误")
	assert.Equal(t, roots.CrossSubnetRoot[:], encoded[224:256], "CrossSubnetRoot位置错误")
}

func TestEncodeCommitmentRoots_EmptyRoots(t *testing.T) {
	roots := CommitmentRoots{} // 全部为零值
	encoded := EncodeCommitmentRoots(roots)

	assert.Equal(t, 256, len(encoded))
	// 所有字节应为0
	for i, b := range encoded {
		assert.Equal(t, byte(0), b, "索引%d应为0", i)
	}
}

func TestEncodeDACommitments_Empty(t *testing.T) {
	commitments := []DACommitment{}
	encoded, err := EncodeDACommitments(commitments)

	require.NoError(t, err)
	assert.NotEmpty(t, encoded, "空数组也应产生有效编码")
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
			SizeHint:      nil, // nil应被normalizeBigInt处理为0
			SegmentHashes: [][32]byte{{0x01}},
			Expiry:        nil, // nil应被normalizeBigInt处理为0
		},
	}

	encoded, err := EncodeDACommitments(commitments)

	require.NoError(t, err, "nil big.Int应被正确处理")
	assert.NotEmpty(t, encoded)
}

func TestEncodeEpochSlot(t *testing.T) {
	slot := EpochSlot{
		Epoch: 12345,
		Slot:  67890,
	}

	encoded := EncodeEpochSlot(slot)

	// 验证长度：2个uint64 * 32字节（左填充）= 64字节
	assert.Equal(t, 64, len(encoded), "编码后应为64字节")

	// 验证每个字段占32字节
	epochBytes := encoded[0:32]
	slotBytes := encoded[32:64]

	assert.NotEmpty(t, epochBytes)
	assert.NotEmpty(t, slotBytes)
}

func TestEncodeEpochSlot_Zero(t *testing.T) {
	slot := EpochSlot{Epoch: 0, Slot: 0}
	encoded := EncodeEpochSlot(slot)

	assert.Equal(t, 64, len(encoded))
	// 所有字节应为0
	for i, b := range encoded {
		assert.Equal(t, byte(0), b, "索引%d应为0", i)
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

	assert.Equal(t, digest1, digest2, "相同输入应产生相同的digest")
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
	roots2.AgentRoot = [32]byte{0xFF} // 修改一个根
	input2.Roots = roots2

	digest1, err1 := ComputeCheckpointDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeCheckpointDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "不同Roots应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同DACommitments应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同Epoch应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同Stats应产生不同的digest（通过statsHash）")
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

	assert.NotEqual(t, digest1, digest2, "不同chain_id应产生不同的digest（防重放）")
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

	assert.NotEqual(t, digest1, digest2, "不同contract应产生不同的digest（防跨合约重放）")
}

func TestNormalizeBigInt(t *testing.T) {
	// 测试非nil值
	value := big.NewInt(12345)
	result := normalizeBigInt(value)
	assert.Equal(t, value, result, "非nil值应原样返回")

	// 测试nil值
	nilResult := normalizeBigInt(nil)
	assert.NotNil(t, nilResult, "nil应被转换为非nil")
	assert.Equal(t, big.NewInt(0), nilResult, "nil应被转换为0")
}
