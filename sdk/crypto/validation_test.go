package crypto

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComputeValidationDigest_Valid(t *testing.T) {
	input := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest, err := ComputeValidationDigest(input, contract, chainID)

	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, digest)
	assert.Len(t, digest, 32)
}

func TestComputeValidationDigest_NilChainID(t *testing.T) {
	input := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digest, err := ComputeValidationDigest(input, contract, nil)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil chain id", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeValidationDigest_Deterministic(t *testing.T) {
	input := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest1, err1 := ComputeValidationDigest(input, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input, contract, chainID)
	require.NoError(t, err2)

	assert.Equal(t, digest1, digest2, "same input should produce the same digest")
}

func TestComputeValidationDigest_DifferentResultHash(t *testing.T) {
	baseInput := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.ResultHash = [32]byte{0x04}

	input2 := baseInput
	input2.ResultHash = [32]byte{0xFF}

	digest1, err1 := ComputeValidationDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different ResultHash should produce different digest")
}

func TestComputeValidationDigest_DifferentProofHash(t *testing.T) {
	baseInput := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.ProofHash = [32]byte{0x05}

	input2 := baseInput
	input2.ProofHash = [32]byte{0xFF}

	digest1, err1 := ComputeValidationDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different ProofHash should produce different digest")
}

func TestComputeValidationDigest_DifferentRootHeight(t *testing.T) {
	baseInput := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.RootHeight = 12345

	input2 := baseInput
	input2.RootHeight = 99999

	digest1, err1 := ComputeValidationDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different RootHeight should produce different digest")
}

func TestComputeValidationDigest_DifferentRootHash(t *testing.T) {
	baseInput := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.RootHash = [32]byte{0x06}

	input2 := baseInput
	input2.RootHash = [32]byte{0xFF}

	digest1, err1 := ComputeValidationDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different RootHash should produce different digest")
}

func TestComputeValidationDigest_DifferentAgent(t *testing.T) {
	baseInput := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.Agent = common.HexToAddress("0xa1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1")

	input2 := baseInput
	input2.Agent = common.HexToAddress("0xa2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2")

	digest1, err1 := ComputeValidationDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different Agent should produce different digest")
}

func TestComputeValidationDigest_DifferentSubnetID(t *testing.T) {
	baseInput := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.SubnetID = [32]byte{0x03}

	input2 := baseInput
	input2.SubnetID = [32]byte{0xFF}

	digest1, err1 := ComputeValidationDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different SubnetID should produce different digest")
}

func TestComputeValidationDigest_DifferentChainID(t *testing.T) {
	input := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digest1, err1 := ComputeValidationDigest(input, contract, big.NewInt(84532))
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input, contract, big.NewInt(11155111))
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different chain_id should produce different digest (replay protection)")
}

func TestComputeValidationDigest_DifferentContract(t *testing.T) {
	input := ValidationInput{
		IntentID:     [32]byte{0x01},
		AssignmentID: [32]byte{0x02},
		SubnetID:     [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		ResultHash:   [32]byte{0x04},
		ProofHash:    [32]byte{0x05},
		RootHeight:   12345,
		RootHash:     [32]byte{0x06},
	}
	chainID := big.NewInt(84532)

	contract1 := common.HexToAddress("0xc1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1")
	contract2 := common.HexToAddress("0xc2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2")

	digest1, err1 := ComputeValidationDigest(input, contract1, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationDigest(input, contract2, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different contract should produce different digest (cross-contract replay protection)")
}

// ========== Batch Validation Tests (v2.3+) ==========

func TestComputeItemsHash_Valid(t *testing.T) {
	items := []ValidationItem{
		{
			IntentID:     [32]byte{0x01},
			AssignmentID: [32]byte{0x02},
			Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
			ResultHash:   [32]byte{0x04},
			ProofHash:    [32]byte{0x05},
		},
		{
			IntentID:     [32]byte{0x11},
			AssignmentID: [32]byte{0x12},
			Agent:        common.HexToAddress("0xa2222222222222222222222222222222222222"),
			ResultHash:   [32]byte{0x14},
			ProofHash:    [32]byte{0x15},
		},
	}

	itemsHash, err := ComputeItemsHash(items)

	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, itemsHash)
	assert.Len(t, itemsHash, 32)
}

func TestComputeItemsHash_EmptyArray(t *testing.T) {
	items := []ValidationItem{}

	itemsHash, err := ComputeItemsHash(items)

	require.Error(t, err)
	assert.Equal(t, "crypto: empty items array", err.Error())
	assert.Equal(t, [32]byte{}, itemsHash)
}

func TestComputeItemsHash_SingleItem(t *testing.T) {
	items := []ValidationItem{
		{
			IntentID:     [32]byte{0x01},
			AssignmentID: [32]byte{0x02},
			Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
			ResultHash:   [32]byte{0x04},
			ProofHash:    [32]byte{0x05},
		},
	}

	itemsHash, err := ComputeItemsHash(items)

	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, itemsHash)
}

func TestComputeItemsHash_Deterministic(t *testing.T) {
	items := []ValidationItem{
		{
			IntentID:     [32]byte{0x01},
			AssignmentID: [32]byte{0x02},
			Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
			ResultHash:   [32]byte{0x04},
			ProofHash:    [32]byte{0x05},
		},
		{
			IntentID:     [32]byte{0x11},
			AssignmentID: [32]byte{0x12},
			Agent:        common.HexToAddress("0xa2222222222222222222222222222222222222"),
			ResultHash:   [32]byte{0x14},
			ProofHash:    [32]byte{0x15},
		},
	}

	hash1, err1 := ComputeItemsHash(items)
	require.NoError(t, err1)

	hash2, err2 := ComputeItemsHash(items)
	require.NoError(t, err2)

	assert.Equal(t, hash1, hash2, "same items should produce the same hash")
}

func TestComputeItemsHash_DifferentOrder(t *testing.T) {
	items1 := []ValidationItem{
		{
			IntentID:     [32]byte{0x01},
			AssignmentID: [32]byte{0x02},
			Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
			ResultHash:   [32]byte{0x04},
			ProofHash:    [32]byte{0x05},
		},
		{
			IntentID:     [32]byte{0x11},
			AssignmentID: [32]byte{0x12},
			Agent:        common.HexToAddress("0xa2222222222222222222222222222222222222"),
			ResultHash:   [32]byte{0x14},
			ProofHash:    [32]byte{0x15},
		},
	}

	items2 := []ValidationItem{
		{
			IntentID:     [32]byte{0x11},
			AssignmentID: [32]byte{0x12},
			Agent:        common.HexToAddress("0xa2222222222222222222222222222222222222"),
			ResultHash:   [32]byte{0x14},
			ProofHash:    [32]byte{0x15},
		},
		{
			IntentID:     [32]byte{0x01},
			AssignmentID: [32]byte{0x02},
			Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
			ResultHash:   [32]byte{0x04},
			ProofHash:    [32]byte{0x05},
		},
	}

	hash1, err1 := ComputeItemsHash(items1)
	require.NoError(t, err1)

	hash2, err2 := ComputeItemsHash(items2)
	require.NoError(t, err2)

	assert.NotEqual(t, hash1, hash2, "items in different order should produce different hash")
}

func TestComputeValidationBatchDigest_Valid(t *testing.T) {
	input := ValidationBatchInput{
		SubnetID:   [32]byte{0x03},
		ItemsHash:  [32]byte{0xAB, 0xCD},
		RootHeight: 12345,
		RootHash:   [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest, err := ComputeValidationBatchDigest(input, contract, chainID)

	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, digest)
	assert.Len(t, digest, 32)
}

func TestComputeValidationBatchDigest_NilChainID(t *testing.T) {
	input := ValidationBatchInput{
		SubnetID:   [32]byte{0x03},
		ItemsHash:  [32]byte{0xAB, 0xCD},
		RootHeight: 12345,
		RootHash:   [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digest, err := ComputeValidationBatchDigest(input, contract, nil)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil chain id", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeValidationBatchDigest_Deterministic(t *testing.T) {
	input := ValidationBatchInput{
		SubnetID:   [32]byte{0x03},
		ItemsHash:  [32]byte{0xAB, 0xCD},
		RootHeight: 12345,
		RootHash:   [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest1, err1 := ComputeValidationBatchDigest(input, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationBatchDigest(input, contract, chainID)
	require.NoError(t, err2)

	assert.Equal(t, digest1, digest2, "same input should produce the same digest")
}

func TestComputeValidationBatchDigest_DifferentItemsHash(t *testing.T) {
	baseInput := ValidationBatchInput{
		SubnetID:   [32]byte{0x03},
		RootHeight: 12345,
		RootHash:   [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.ItemsHash = [32]byte{0xAB, 0xCD}

	input2 := baseInput
	input2.ItemsHash = [32]byte{0xFF, 0xEE}

	digest1, err1 := ComputeValidationBatchDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationBatchDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different ItemsHash should produce different digest")
}

func TestComputeValidationBatchDigest_DifferentSubnetID(t *testing.T) {
	baseInput := ValidationBatchInput{
		ItemsHash:  [32]byte{0xAB, 0xCD},
		RootHeight: 12345,
		RootHash:   [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.SubnetID = [32]byte{0x03}

	input2 := baseInput
	input2.SubnetID = [32]byte{0xFF}

	digest1, err1 := ComputeValidationBatchDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationBatchDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different SubnetID should produce different digest")
}

func TestComputeValidationBatchDigest_DifferentChainID(t *testing.T) {
	input := ValidationBatchInput{
		SubnetID:   [32]byte{0x03},
		ItemsHash:  [32]byte{0xAB, 0xCD},
		RootHeight: 12345,
		RootHash:   [32]byte{0x06},
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digest1, err1 := ComputeValidationBatchDigest(input, contract, big.NewInt(84532))
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationBatchDigest(input, contract, big.NewInt(11155111))
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different chain_id should produce different digest (replay protection)")
}

func TestComputeValidationBatchDigest_DifferentContract(t *testing.T) {
	input := ValidationBatchInput{
		SubnetID:   [32]byte{0x03},
		ItemsHash:  [32]byte{0xAB, 0xCD},
		RootHeight: 12345,
		RootHash:   [32]byte{0x06},
	}
	chainID := big.NewInt(84532)

	contract1 := common.HexToAddress("0xc1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1")
	contract2 := common.HexToAddress("0xc2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2")

	digest1, err1 := ComputeValidationBatchDigest(input, contract1, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeValidationBatchDigest(input, contract2, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different contract should produce different digest (cross-contract replay protection)")
}
