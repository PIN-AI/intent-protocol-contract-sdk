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

	assert.Equal(t, digest1, digest2, "相同输入应产生相同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同ResultHash应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同ProofHash应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同RootHeight应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同RootHash应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同Agent应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同SubnetID应产生不同的digest")
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

	assert.NotEqual(t, digest1, digest2, "不同chain_id应产生不同的digest（防重放）")
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

	assert.NotEqual(t, digest1, digest2, "不同contract应产生不同的digest（防跨合约重放）")
}
