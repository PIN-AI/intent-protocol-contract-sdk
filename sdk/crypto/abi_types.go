package crypto

import "github.com/ethereum/go-ethereum/accounts/abi"

// 共享的 ABI type 定义，供 intent/assignment/validation/checkpoint digest 构造使用
var (
	typeBytes32, _ = abi.NewType("bytes32", "", nil)
	typeAddress, _ = abi.NewType("address", "", nil)
	typeUint256, _ = abi.NewType("uint256", "", nil)
	typeUint8, _   = abi.NewType("uint8", "", nil)
	typeUint32, _  = abi.NewType("uint32", "", nil)
	typeUint64, _  = abi.NewType("uint64", "", nil)
)
