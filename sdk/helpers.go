package sdk

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// Bytes32FromHex 将 0x 前缀的字符串转换为 [32]byte。
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

// MustBytes32FromHex panic 版本。
func MustBytes32FromHex(input string) [32]byte {
	res, err := Bytes32FromHex(input)
	if err != nil {
		panic(err)
	}
	return res
}

// ZeroAddress 是 0 地址常量。
var ZeroAddress = common.Address{}
