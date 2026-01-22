package sdk

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	sdkcrypto "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
	pb "github.com/PIN-AI/pin-protocol-proto/rootlayer/proto"
	"github.com/ethereum/go-ethereum/common"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
)

func hashRootLayerParams(params *pb.IntentParams) ([32]byte, error) {
	var out [32]byte
	if params == nil {
		return out, errors.New("sdk: params is required")
	}
	if len(params.IntentRaw) == 0 {
		return out, errors.New("sdk: params.intent_raw cannot be empty")
	}

	h := gethcrypto.Keccak256Hash(params.IntentRaw, params.Metadata)
	copy(out[:], h[:])
	return out, nil
}

func parseHexAddress(value string) (common.Address, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return common.Address{}, errors.New("empty address")
	}
	if !common.IsHexAddress(trimmed) {
		return common.Address{}, fmt.Errorf("invalid address: %s", value)
	}
	return common.HexToAddress(trimmed), nil
}

func parseOptionalHexAddress(value string) (common.Address, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return common.Address{}, nil
	}
	return parseHexAddress(trimmed)
}

func parseUint256(value string) (*big.Int, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return nil, errors.New("uint256 value is required")
	}
	base := 10
	if strings.HasPrefix(trimmed, "0x") || strings.HasPrefix(trimmed, "0X") {
		base = 16
		trimmed = trimmed[2:]
		if trimmed == "" {
			return nil, errors.New("uint256 value is required")
		}
	}
	v, ok := new(big.Int).SetString(trimmed, base)
	if !ok {
		return nil, fmt.Errorf("invalid uint256: %s", value)
	}
	if v.Sign() < 0 {
		return nil, errors.New("uint256 cannot be negative")
	}
	return v, nil
}

// SubmitIntentToRootLayer submits a standard intent to RootLayer. If Signature is empty, it will be auto-signed.
func (c *Client) SubmitIntentToRootLayer(ctx context.Context, req *pb.SubmitIntentRequest) (*pb.SubmitIntentResponse, error) {
	if c == nil {
		return nil, errors.New("sdk: nil client")
	}
	if c.RootLayer == nil {
		return nil, errors.New("sdk: RootLayerURL not configured")
	}
	if req == nil {
		return nil, errors.New("sdk: nil request")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	if strings.TrimSpace(req.Requester) == "" {
		req.Requester = c.Signer.Address().Hex()
	}
	if strings.TrimSpace(req.SettleChain) == "" {
		req.SettleChain = string(c.Network)
	}
	if strings.TrimSpace(req.TipsToken) == "" {
		req.TipsToken = ZeroAddress.Hex()
	}
	if strings.TrimSpace(req.Tips) == "" {
		req.Tips = "0"
	}
	if strings.TrimSpace(req.BudgetToken) == "" {
		req.BudgetToken = ZeroAddress.Hex()
	}
	if strings.TrimSpace(req.Budget) == "" {
		req.Budget = "0"
	}

	if len(req.Signature) == 0 {
		intentID, err := Bytes32FromHex(req.IntentId)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse intent_id: %w", err)
		}
		subnetID, err := Bytes32FromHex(req.SubnetId)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse subnet_id: %w", err)
		}
		requester, err := parseHexAddress(req.Requester)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse requester: %w", err)
		}
		budgetToken, err := parseOptionalHexAddress(req.BudgetToken)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse budget_token: %w", err)
		}
		budget, err := parseUint256(req.Budget)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse budget: %w", err)
		}
		paramsHash, err := hashRootLayerParams(req.Params)
		if err != nil {
			return nil, err
		}

		input := sdkcrypto.SignedIntentInput{
			IntentID:     intentID,
			SubnetID:     subnetID,
			Requester:    requester,
			IntentType:   req.IntentType,
			ParamsHash:   paramsHash,
			Deadline:     big.NewInt(req.Deadline),
			PaymentToken: budgetToken,
			Amount:       budget,
		}
		sig, err := c.Intent.SignIntent(input)
		if err != nil {
			return nil, err
		}
		req.Signature = sig
	}

	return c.RootLayer.IntentPool.SubmitIntent(ctx, req)
}

// SubmitDirectIntentToRootLayer submits a Direct Mode intent to RootLayer. If Signature is empty, it will be auto-signed.
func (c *Client) SubmitDirectIntentToRootLayer(ctx context.Context, req *pb.SubmitDirectIntentRequest) (*pb.SubmitDirectIntentResponse, error) {
	if c == nil {
		return nil, errors.New("sdk: nil client")
	}
	if c.RootLayer == nil {
		return nil, errors.New("sdk: RootLayerURL not configured")
	}
	if req == nil {
		return nil, errors.New("sdk: nil request")
	}
	if ctx == nil {
		ctx = context.Background()
	}

	if strings.TrimSpace(req.Requester) == "" {
		req.Requester = c.Signer.Address().Hex()
	}
	if strings.TrimSpace(req.SettleChain) == "" {
		req.SettleChain = string(c.Network)
	}
	if strings.TrimSpace(req.PaymentToken) == "" {
		req.PaymentToken = ZeroAddress.Hex()
	}
	if strings.TrimSpace(req.Amount) == "" {
		req.Amount = "0"
	}
	if strings.TrimSpace(req.TargetAgentId) == "" {
		return nil, errors.New("sdk: target_agent_id is required")
	}

	if len(req.Signature) == 0 {
		intentID, err := Bytes32FromHex(req.IntentId)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse intent_id: %w", err)
		}
		subnetID, err := Bytes32FromHex(req.SubnetId)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse subnet_id: %w", err)
		}
		requester, err := parseHexAddress(req.Requester)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse requester: %w", err)
		}
		targetAgent, err := parseHexAddress(req.TargetAgent)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse target_agent: %w", err)
		}
		paymentToken, err := parseOptionalHexAddress(req.PaymentToken)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse payment_token: %w", err)
		}
		amount, err := parseUint256(req.Amount)
		if err != nil {
			return nil, fmt.Errorf("sdk: parse amount: %w", err)
		}
		paramsHash, err := hashRootLayerParams(req.Params)
		if err != nil {
			return nil, err
		}

		input := sdkcrypto.DirectIntentInput{
			IntentID:     intentID,
			SubnetID:     subnetID,
			Requester:    requester,
			IntentType:   req.IntentType,
			ParamsHash:   paramsHash,
			Deadline:     big.NewInt(req.Deadline),
			PaymentToken: paymentToken,
			Amount:       amount,
			TargetAgent:  targetAgent,
		}
		sig, err := c.Intent.SignDirectIntent(input)
		if err != nil {
			return nil, err
		}
		req.Signature = sig
	}

	return c.RootLayer.DirectIntent.SubmitDirectIntent(ctx, req)
}
