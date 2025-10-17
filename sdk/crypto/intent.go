package crypto

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	intentTypeHash = crypto.Keccak256Hash([]byte("PIN_INTENT_V1(bytes32,bytes32,address,bytes32,bytes32,uint256,address,uint256,address,uint256)"))

	// arguments: typehash, intent_id, subnet_id, requester, intent_type_hash, params_hash, deadline, payment_token, amount, contract, chainId
	intentDigestArgs = abi.Arguments{
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeAddress},
		{Type: typeBytes32},
		{Type: typeBytes32},
		{Type: typeUint256},
		{Type: typeAddress},
		{Type: typeUint256},
		{Type: typeAddress},
		{Type: typeUint256},
	}
)

// SignedIntentInput represents the data required to construct an Intent digest.
type SignedIntentInput struct {
	IntentID     [32]byte
	SubnetID     [32]byte
	Requester    common.Address
	IntentType   string
	ParamsHash   [32]byte
	Deadline     *big.Int
	PaymentToken common.Address
	Amount       *big.Int
}

// ComputeIntentDigest computes the digest required for submitIntentsBySignatures.
func ComputeIntentDigest(input SignedIntentInput, contract common.Address, chainID *big.Int) ([32]byte, error) {
	var zero [32]byte
	if chainID == nil {
		return zero, errors.New("crypto: nil chain id")
	}
	if input.Deadline == nil {
		return zero, errors.New("crypto: nil deadline")
	}
	if input.Amount == nil {
		return zero, errors.New("crypto: nil amount")
	}
	intentTypeDigest := crypto.Keccak256Hash([]byte(input.IntentType))
	encoded, err := intentDigestArgs.Pack(
		intentTypeHash,
		input.IntentID,
		input.SubnetID,
		input.Requester,
		intentTypeDigest,
		input.ParamsHash,
		input.Deadline,
		input.PaymentToken,
		input.Amount,
		contract,
		chainID,
	)
	if err != nil {
		return zero, err
	}
	return bytesToBytes32(crypto.Keccak256Hash(encoded).Bytes()), nil
}
