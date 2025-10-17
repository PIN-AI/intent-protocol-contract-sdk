package sdk

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Bytes32FromHex converts a 0x-prefixed hex string to [32]byte.
func Bytes32FromHex(input string) ([32]byte, error) {
	var out [32]byte
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return out, fmt.Errorf("sdk: empty hex string")
	}
	if strings.HasPrefix(trimmed, "0x") || strings.HasPrefix(trimmed, "0X") {
		trimmed = trimmed[2:]
	}
	b, err := hex.DecodeString(trimmed)
	if err != nil {
		return out, fmt.Errorf("sdk: decode hex: %w", err)
	}
	if len(b) != 32 {
		return out, fmt.Errorf("sdk: expect 32 bytes got %d", len(b))
	}
	copy(out[:], b)
	return out, nil
}

// MustBytes32FromHex is the panic version of Bytes32FromHex.
func MustBytes32FromHex(input string) [32]byte {
	res, err := Bytes32FromHex(input)
	if err != nil {
		panic(err)
	}
	return res
}

// ZeroAddress is the zero address constant.
var ZeroAddress = common.Address{}
