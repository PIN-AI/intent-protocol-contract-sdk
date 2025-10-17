package crypto

import "github.com/ethereum/go-ethereum/crypto"

// HashBytes computes keccak256 for arbitrary data.
func HashBytes(data []byte) [32]byte {
	return bytesToBytes32(crypto.Keccak256Hash(data).Bytes())
}

// bytesToBytes32 converts []byte to [32]byte, an internal helper function.
func bytesToBytes32(b []byte) [32]byte {
	var out [32]byte
	copy(out[:], b)
	return out
}
