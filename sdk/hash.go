package sdk

import sdkcrypto "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"

// HashBytes exposes keccak256 computation.
func HashBytes(data []byte) [32]byte {
	return sdkcrypto.HashBytes(data)
}
