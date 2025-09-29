package signer

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Signer 抽象用于交易签名与 EIP-191 摘要签名。
type Signer interface {
	Address() common.Address
	SignDigest(digest [32]byte) ([]byte, error)
	SignTransaction(ctx context.Context, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
}

// ErrInvalidSignerAddress 当 signer 与请求地址不匹配时返回。
var ErrInvalidSignerAddress = errors.New("signer: invalid address")
