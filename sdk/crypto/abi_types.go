package crypto

import "github.com/ethereum/go-ethereum/accounts/abi"

// Shared ABI type definitions, used for intent/assignment/validation/checkpoint digest construction
var (
	typeBytes32, _ = abi.NewType("bytes32", "", nil)
	typeAddress, _ = abi.NewType("address", "", nil)
	typeUint256, _ = abi.NewType("uint256", "", nil)
	typeUint8, _   = abi.NewType("uint8", "", nil)
	typeUint32, _  = abi.NewType("uint32", "", nil)
	typeUint64, _  = abi.NewType("uint64", "", nil)
)
