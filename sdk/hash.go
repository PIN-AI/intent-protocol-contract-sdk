package sdk

import sdkcrypto "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"

// HashBytes 暴露 keccak256 计算。
func HashBytes(data []byte) [32]byte {
	return sdkcrypto.HashBytes(data)
}
