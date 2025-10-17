package signer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// LocalSigner implements Signer using a local private key.
type LocalSigner struct {
	key  *ecdsa.PrivateKey
	addr common.Address
}

// NewLocalSigner parses a 0x-prefixed private key string.
func NewLocalSigner(hexKey string) (*LocalSigner, error) {
	trimmed := strings.TrimSpace(hexKey)
	if trimmed == "" {
		return nil, fmt.Errorf("signer: empty private key")
	}
	if strings.HasPrefix(trimmed, "0x") || strings.HasPrefix(trimmed, "0X") {
		trimmed = trimmed[2:]
	}
	key, err := crypto.HexToECDSA(trimmed)
	if err != nil {
		return nil, fmt.Errorf("signer: parse private key: %w", err)
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return &LocalSigner{key: key, addr: addr}, nil
}

// MustNewLocalSigner is the panic wrapper for NewLocalSigner.
func MustNewLocalSigner(hexKey string) *LocalSigner {
	s, err := NewLocalSigner(hexKey)
	if err != nil {
		panic(err)
	}
	return s
}

func (s *LocalSigner) Address() common.Address {
	return s.addr
}

// SignDigest performs signing according to EIP-191 (eth_sign).
func (s *LocalSigner) SignDigest(digest [32]byte) ([]byte, error) {
	msg := accounts.TextHash(digest[:])
	sig, err := crypto.Sign(msg, s.key)
	if err != nil {
		return nil, err
	}
	if sig[64] < 27 {
		sig[64] += 27
	}
	return sig, nil
}

// SignTransaction signs a transaction.
func (s *LocalSigner) SignTransaction(_ context.Context, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	if tx == nil {
		return nil, fmt.Errorf("signer: nil transaction")
	}
	if chainID == nil {
		return nil, fmt.Errorf("signer: nil chain id")
	}
	return types.SignTx(tx, types.LatestSignerForChainID(chainID), s.key)
}
