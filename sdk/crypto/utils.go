package crypto

import "github.com/ethereum/go-ethereum/crypto"

// HashBytes 计算任意数据的 keccak256。
func HashBytes(data []byte) [32]byte {
	return bytesToBytes32(crypto.Keccak256Hash(data).Bytes())
}

// bytesToBytes32 将 []byte 转换为 [32]byte，内部辅助函数。
func bytesToBytes32(b []byte) [32]byte {
	var out [32]byte
	copy(out[:], b)
	return out
}
