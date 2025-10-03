package crypto

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComputeAssignmentDigest_Valid(t *testing.T) {
	input := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1, // ACTIVE
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest, err := ComputeAssignmentDigest(input, contract, chainID)

	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, digest)
	assert.Len(t, digest, 32)
}

func TestComputeAssignmentDigest_NilChainID(t *testing.T) {
	input := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digest, err := ComputeAssignmentDigest(input, contract, nil)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil chain id", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeAssignmentDigest_Deterministic(t *testing.T) {
	input := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest1, err1 := ComputeAssignmentDigest(input, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input, contract, chainID)
	require.NoError(t, err2)

	assert.Equal(t, digest1, digest2, "相同输入应产生相同的digest")
}

func TestComputeAssignmentDigest_DifferentStatus(t *testing.T) {
	baseInput := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.Status = 1 // ACTIVE

	input2 := baseInput
	input2.Status = 2 // FAILED

	digest1, err1 := ComputeAssignmentDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "不同Status应产生不同的digest")
}

func TestComputeAssignmentDigest_DifferentAgent(t *testing.T) {
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}

	input2 := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2a2"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}

	digest1, err1 := ComputeAssignmentDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "不同Agent应产生不同的digest")
}

func TestComputeAssignmentDigest_DifferentMatcher(t *testing.T) {
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}

	input2 := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2b2"),
	}

	digest1, err1 := ComputeAssignmentDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "不同Matcher应产生不同的digest")
}

func TestComputeAssignmentDigest_DifferentChainID(t *testing.T) {
	input := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digest1, err1 := ComputeAssignmentDigest(input, contract, big.NewInt(84532))
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input, contract, big.NewInt(11155111))
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "不同chain_id应产生不同的digest（防重放）")
}

func TestComputeAssignmentDigest_DifferentContract(t *testing.T) {
	input := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}
	chainID := big.NewInt(84532)

	contract1 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	contract2 := common.HexToAddress("0x2222222222222222222222222222222222222222")

	digest1, err1 := ComputeAssignmentDigest(input, contract1, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input, contract2, chainID)
	require.NoError(t, err2)

	t.Logf("Arguments count: %d", len(assignmentDigestArgs))
	t.Logf("Contract1: %s", contract1.Hex())
	t.Logf("Contract2: %s", contract2.Hex())
	t.Logf("Digest1: %x", digest1)
	t.Logf("Digest2: %x", digest2)

	assert.NotEqual(t, digest1, digest2, "不同contract应产生不同的digest（防跨合约重放）")
}

func TestComputeAssignmentDigest_AllIDs(t *testing.T) {
	baseInput := AssignmentInput{
		Agent:   common.HexToAddress("0xa1111111111111111111111111111111111111"),
		Status:  1,
		Matcher: common.HexToAddress("0xb2222222222222222222222222222222222222"),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	// 测试不同的AssignmentID
	input1 := baseInput
	input1.AssignmentID = [32]byte{0x01}
	input1.IntentID = [32]byte{0x02}
	input1.BidID = [32]byte{0x03}

	input2 := baseInput
	input2.AssignmentID = [32]byte{0xFF}
	input2.IntentID = [32]byte{0x02}
	input2.BidID = [32]byte{0x03}

	digest1, err1 := ComputeAssignmentDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "不同AssignmentID应产生不同的digest")

	// 测试不同的IntentID
	input3 := baseInput
	input3.AssignmentID = [32]byte{0x01}
	input3.IntentID = [32]byte{0xFF}
	input3.BidID = [32]byte{0x03}

	digest3, err3 := ComputeAssignmentDigest(input3, contract, chainID)
	require.NoError(t, err3)

	assert.NotEqual(t, digest1, digest3, "不同IntentID应产生不同的digest")

	// 测试不同的BidID
	input4 := baseInput
	input4.AssignmentID = [32]byte{0x01}
	input4.IntentID = [32]byte{0x02}
	input4.BidID = [32]byte{0xFF}

	digest4, err4 := ComputeAssignmentDigest(input4, contract, chainID)
	require.NoError(t, err4)

	assert.NotEqual(t, digest1, digest4, "不同BidID应产生不同的digest")
}
