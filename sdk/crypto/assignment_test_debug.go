package crypto

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestComputeAssignmentDigest_DifferentContract_Debug(t *testing.T) {
	input := AssignmentInput{
		AssignmentID: [32]byte{0x01},
		IntentID:     [32]byte{0x02},
		BidID:        [32]byte{0x03},
		Agent:        common.HexToAddress("0xa111111111111111111111111111111111111111"),
		Status:       1,
		Matcher:      common.HexToAddress("0xb222222222222222222222222222222222222222"),
	}
	chainID := big.NewInt(84532)

	contract1 := common.HexToAddress("0x1111000000000000000000000000000000000001")
	contract2 := common.HexToAddress("0x1111000000000000000000000000000000000002")

	digest1, err1 := ComputeAssignmentDigest(input, contract1, chainID)
	require.NoError(t, err1)

	digest2, err2 := ComputeAssignmentDigest(input, contract2, chainID)
	require.NoError(t, err2)

	fmt.Printf("\n=== DEBUG OUTPUT ===\n")
	fmt.Printf("Contract1: %s\n", contract1.Hex())
	fmt.Printf("Contract2: %s\n", contract2.Hex())
	fmt.Printf("Digest1: %x\n", digest1)
	fmt.Printf("Digest2: %x\n", digest2)
	fmt.Printf("Equal: %v\n", digest1 == digest2)
	fmt.Printf("Arguments length: %d\n", len(assignmentDigestArgs))
	fmt.Printf("===================\n\n")

	if digest1 == digest2 {
		t.Errorf("不同contract应产生不同的digest！但它们相等")
	}
}
