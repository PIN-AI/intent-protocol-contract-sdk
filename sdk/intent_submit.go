package sdk

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	cryptoHelpers "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
)

// SubmitIntentParams is a convenience wrapper for submitting a single intent.
//
// Under the hood, this signs the intent digest and calls SubmitIntentsBySignatures
// with a batch size of 1.
type SubmitIntentParams struct {
	IntentID   [32]byte
	SubnetID   [32]byte
	IntentType string
	ParamsHash [32]byte
	Deadline   *big.Int

	PaymentToken common.Address
	Amount       *big.Int

	// Value is the ETH msg.value to send with the transaction.
	// For ETH payments (PaymentToken == ZeroAddress), Value should be equal to Amount.
	// For ERC20 payments, Value should be nil or 0.
	Value *big.Int

	// Optional override; defaults to the configured signer address.
	Requester common.Address
}

// SubmitIntent signs and submits a single intent.
func (s *IntentService) SubmitIntent(ctx context.Context, params SubmitIntentParams) (*types.Transaction, error) {
	if params.Deadline == nil {
		return nil, errors.New("intent: nil deadline")
	}
	if params.Amount == nil {
		return nil, errors.New("intent: nil amount")
	}

	requester := params.Requester
	if requester == (common.Address{}) {
		requester = s.signer.Address()
	}

	input := cryptoHelpers.SignedIntentInput{
		IntentID:     params.IntentID,
		SubnetID:     params.SubnetID,
		Requester:    requester,
		IntentType:   params.IntentType,
		ParamsHash:   params.ParamsHash,
		Deadline:     params.Deadline,
		PaymentToken: params.PaymentToken,
		Amount:       params.Amount,
	}

	sig, err := s.SignIntent(input)
	if err != nil {
		return nil, err
	}

	batch := SubmitIntentBatchParams{
		Items: []SignedIntent{{Data: input, Signature: sig}},
	}
	if params.PaymentToken == ZeroAddress {
		value := params.Value
		if value == nil {
			value = params.Amount
		}
		if value.Sign() != 0 && value.Cmp(params.Amount) != 0 {
			return nil, errors.New("intent: ETH value must equal amount")
		}
		batch.TotalEthValue = new(big.Int).Set(value)
	} else if params.Value != nil && params.Value.Sign() != 0 {
		return nil, errors.New("intent: ERC20 intent must not include ETH value")
	}

	return s.SubmitIntentsBySignatures(ctx, batch)
}

// SubmitDirectIntentParams is a convenience wrapper for submitting a single direct intent.
//
// Under the hood, this signs the direct intent digest and calls SubmitDirectIntentsBySignatures
// with a batch size of 1.
type SubmitDirectIntentParams struct {
	IntentID   [32]byte
	SubnetID   [32]byte
	IntentType string
	ParamsHash [32]byte
	Deadline   *big.Int

	PaymentToken common.Address
	Amount       *big.Int
	Value        *big.Int

	TargetAgent common.Address

	// Optional override; defaults to the configured signer address.
	Requester common.Address
}

// SubmitDirectIntent signs and submits a single direct intent.
func (s *IntentService) SubmitDirectIntent(ctx context.Context, params SubmitDirectIntentParams) (*types.Transaction, error) {
	if params.Deadline == nil {
		return nil, errors.New("intent: nil deadline")
	}
	if params.Amount == nil {
		return nil, errors.New("intent: nil amount")
	}
	if params.TargetAgent == (common.Address{}) {
		return nil, errors.New("intent: missing target agent")
	}

	requester := params.Requester
	if requester == (common.Address{}) {
		requester = s.signer.Address()
	}

	input := cryptoHelpers.DirectIntentInput{
		IntentID:     params.IntentID,
		SubnetID:     params.SubnetID,
		Requester:    requester,
		IntentType:   params.IntentType,
		ParamsHash:   params.ParamsHash,
		Deadline:     params.Deadline,
		PaymentToken: params.PaymentToken,
		Amount:       params.Amount,
		TargetAgent:  params.TargetAgent,
	}

	sig, err := s.SignDirectIntent(input)
	if err != nil {
		return nil, err
	}

	batch := SubmitDirectIntentBatchParams{
		Items: []SignedDirectIntent{{Data: input, Signature: sig}},
	}
	if params.PaymentToken == ZeroAddress {
		value := params.Value
		if value == nil {
			value = params.Amount
		}
		if value.Sign() != 0 && value.Cmp(params.Amount) != 0 {
			return nil, errors.New("intent: direct ETH value must equal amount")
		}
		batch.TotalEthValue = new(big.Int).Set(value)
	} else if params.Value != nil && params.Value.Sign() != 0 {
		return nil, errors.New("intent: direct ERC20 intent must not include ETH value")
	}

	return s.SubmitDirectIntentsBySignatures(ctx, batch)
}
