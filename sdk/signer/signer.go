package signer

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Signer abstracts transaction signing and EIP-191 digest signing.
type Signer interface {
	Address() common.Address
	SignDigest(digest [32]byte) ([]byte, error)
	SignTransaction(ctx context.Context, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error)
}

// ErrInvalidSignerAddress is returned when signer does not match the requested address.
var ErrInvalidSignerAddress = errors.New("signer: invalid address")
