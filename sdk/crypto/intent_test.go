package crypto

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComputeIntentDigest_Valid(t *testing.T) {
	input := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		IntentType:   "test-intent",
		ParamsHash:   [32]byte{0x03},
		Deadline:     big.NewInt(1234567890),
		PaymentToken: common.Address{},                // Zero address (ETH)
		Amount:       big.NewInt(1000000000000000000), // 1 ETH
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532) // Base Sepolia

	digest, err := ComputeIntentDigest(input, contract, chainID)

	require.NoError(t, err)
	assert.NotEqual(t, [32]byte{}, digest, "digest should not be zero")
	assert.Len(t, digest, 32, "digest should be 32 bytes")
}

func TestComputeIntentDigest_NilChainID(t *testing.T) {
	input := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		IntentType:   "test-intent",
		ParamsHash:   [32]byte{0x03},
		Deadline:     big.NewInt(1234567890),
		PaymentToken: common.Address{},
		Amount:       big.NewInt(1000000000000000000),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digest, err := ComputeIntentDigest(input, contract, nil)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil chain id", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeIntentDigest_NilDeadline(t *testing.T) {
	input := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		IntentType:   "test-intent",
		ParamsHash:   [32]byte{0x03},
		Deadline:     nil,
		PaymentToken: common.Address{},
		Amount:       big.NewInt(1000000000000000000),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest, err := ComputeIntentDigest(input, contract, chainID)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil deadline", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeIntentDigest_NilAmount(t *testing.T) {
	input := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		IntentType:   "test-intent",
		ParamsHash:   [32]byte{0x03},
		Deadline:     big.NewInt(1234567890),
		PaymentToken: common.Address{},
		Amount:       nil,
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	digest, err := ComputeIntentDigest(input, contract, chainID)

	require.Error(t, err)
	assert.Equal(t, "crypto: nil amount", err.Error())
	assert.Equal(t, [32]byte{}, digest)
}

func TestComputeIntentDigest_Deterministic(t *testing.T) {
	input := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		IntentType:   "test-intent",
		ParamsHash:   [32]byte{0x03},
		Deadline:     big.NewInt(1234567890),
		PaymentToken: common.Address{},
		Amount:       big.NewInt(1000000000000000000),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	// compute twice
	digest1, err1 := ComputeIntentDigest(input, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeIntentDigest(input, contract, chainID)
	require.NoError(t, err2)

	assert.Equal(t, digest1, digest2, "same input should produce the same digest (deterministic)")
}

func TestComputeIntentDigest_DifferentChainID(t *testing.T) {
	input := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		IntentType:   "test-intent",
		ParamsHash:   [32]byte{0x03},
		Deadline:     big.NewInt(1234567890),
		PaymentToken: common.Address{},
		Amount:       big.NewInt(1000000000000000000),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")

	digestBaseSepolia, err1 := ComputeIntentDigest(input, contract, big.NewInt(84532)) // Base Sepolia
	require.NoError(t, err1)

	digestSepolia, err2 := ComputeIntentDigest(input, contract, big.NewInt(11155111)) // Sepolia
	require.NoError(t, err2)

	assert.NotEqual(t, digestBaseSepolia, digestSepolia, "different chain_id should produce different digest (replay protection)")
}

func TestComputeIntentDigest_DifferentContract(t *testing.T) {
	input := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		IntentType:   "test-intent",
		ParamsHash:   [32]byte{0x03},
		Deadline:     big.NewInt(1234567890),
		PaymentToken: common.Address{},
		Amount:       big.NewInt(1000000000000000000),
	}
	chainID := big.NewInt(84532)

	contract1 := common.HexToAddress("0xc1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1")
	contract2 := common.HexToAddress("0xc2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2c2")

	digest1, err1 := ComputeIntentDigest(input, contract1, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeIntentDigest(input, contract2, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different contract_address should produce different digest (cross-contract replay protection)")
}

func TestComputeIntentDigest_DifferentIntentType(t *testing.T) {
	baseInput := SignedIntentInput{
		IntentID:     [32]byte{0x01},
		SubnetID:     [32]byte{0x02},
		Requester:    common.HexToAddress("0x1234567890123456789012345678901234567890"),
		ParamsHash:   [32]byte{0x03},
		Deadline:     big.NewInt(1234567890),
		PaymentToken: common.Address{},
		Amount:       big.NewInt(1000000000000000000),
	}
	contract := common.HexToAddress("0xc3333333333333333333333333333333333333")
	chainID := big.NewInt(84532)

	input1 := baseInput
	input1.IntentType = "intent-type-1"

	input2 := baseInput
	input2.IntentType = "intent-type-2"

	digest1, err1 := ComputeIntentDigest(input1, contract, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeIntentDigest(input2, contract, chainID)
	require.NoError(t, err2)

	assert.NotEqual(t, digest1, digest2, "different IntentType should produce different digest")
}

func TestHashBytes(t *testing.T) {
	data := []byte("test data")
	hash := HashBytes(data)

	assert.NotEqual(t, [32]byte{}, hash)
	assert.Len(t, hash, 32)

	// deterministic test
	hash2 := HashBytes(data)
	assert.Equal(t, hash, hash2)

	// different data produces different hash
	differentData := []byte("different data")
	differentHash := HashBytes(differentData)
	assert.NotEqual(t, hash, differentHash)
}

func TestHashBytes_Empty(t *testing.T) {
	emptyHash := HashBytes([]byte{})
	assert.NotEqual(t, [32]byte{}, emptyHash, "empty data should also produce valid hash")

	nilHash := HashBytes(nil)
	assert.NotEqual(t, [32]byte{}, nilHash, "nil data should also produce valid hash")

	// empty array and nil should produce same hash
	assert.Equal(t, emptyHash, nilHash)
}
