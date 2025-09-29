package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
	sdkcrypto "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// Auto-load .env from CWD if present (simple parser, no external deps)
	_ = loadDotEnv(".env")

	cfg := runConfig{
		rpcURL:       strings.TrimSpace(envOr("PIN_RPC_URL", envOr("RPC_URL", ""))),
		privateKey:   strings.TrimSpace(envOr("PIN_PRIVATE_KEY", envOr("PRIVATE_KEY", ""))),
		network:      strings.TrimSpace(envOr("PIN_NETWORK", "")),
		intentIDHex:  strings.TrimSpace(os.Getenv("INTENT_ID")),
		subnetIDHex:  strings.TrimSpace(os.Getenv("SUBNET_ID")),
		intentType:   strings.TrimSpace(envOr("INTENT_TYPE", "test_intent")),
		paramsHash:   strings.TrimSpace(os.Getenv("PARAMS_HASH")),
		paramsJSON:   envOr("PARAMS_JSON", "{}"),
		deadline:     envDuration("DEADLINE_OFFSET", 24*time.Hour),
		amountWei:    strings.TrimSpace(envOr("AMOUNT_WEI", "1000000000000000")), // default 0.001 ETH
		valueWei:     strings.TrimSpace(os.Getenv("VALUE_WEI")),
		paymentToken: strings.TrimSpace(os.Getenv("PAYMENT_TOKEN")),
		dryRun:       envBool("DRY_RUN", true),
		timeout:      envDuration("TIMEOUT", 45*time.Second),
	}

	if err := runSendIntent(cfg); err != nil {
		log.Fatalf("send intent: %v", err)
	}
}

type runConfig struct {
	rpcURL       string
	privateKey   string
	network      string
	intentIDHex  string
	subnetIDHex  string
	intentType   string
	paramsHash   string
	paramsJSON   string
	deadline     time.Duration
	amountWei    string
	valueWei     string
	paymentToken string
	dryRun       bool
	timeout      time.Duration
}

func runSendIntent(cfg runConfig) error {
	if cfg.rpcURL == "" {
		return errors.New("missing RPC endpoint: set PIN_RPC_URL or RPC_URL in env/.env")
	}
	if cfg.privateKey == "" {
		return errors.New("missing private key: set PIN_PRIVATE_KEY or PRIVATE_KEY in env/.env")
	}

	// intent ID: use provided or generate random 32 bytes
	var intentID [32]byte
	var err error
	if cfg.intentIDHex != "" {
		intentID, err = sdk.Bytes32FromHex(cfg.intentIDHex)
		if err != nil {
			return fmt.Errorf("parse intent-id: %w", err)
		}
	} else {
		rb := make([]byte, 32)
		if _, err := rand.Read(rb); err != nil {
			return fmt.Errorf("rand intent-id: %w", err)
		}
		intentID, _ = sdk.Bytes32FromHex("0x" + hex.EncodeToString(rb))
		log.Printf("INTENT_ID not set; generated %s", "0x"+hex.EncodeToString(rb))
	}

	// subnet ID: use provided or pick first active from factory
	var subnetID [32]byte
	haveSubnet := false
	if cfg.subnetIDHex != "" {
		subnetID, err = sdk.Bytes32FromHex(cfg.subnetIDHex)
		if err != nil {
			return fmt.Errorf("parse subnet-id: %w", err)
		}
		haveSubnet = true
	}

	paramsHash, err := deriveParamsHash(cfg.paramsHash, cfg.paramsJSON)
	if err != nil {
		return err
	}

	amount, err := parseBigInt(cfg.amountWei)
	if err != nil {
		return fmt.Errorf("parse amount-wei: %w", err)
	}
	value, err := parseOptionalBigInt(cfg.valueWei)
	if err != nil {
		return fmt.Errorf("parse value-wei: %w", err)
	}

	paymentToken, err := parseAddress(cfg.paymentToken)
	if err != nil {
		return err
	}

	deadline := big.NewInt(time.Now().Add(cfg.deadline).Unix())

	ctx, cancel := context.WithTimeout(context.Background(), cfg.timeout)
	defer cancel()

	sdkCfg := sdk.Config{
		RPCURL:        cfg.rpcURL,
		PrivateKeyHex: cfg.privateKey,
		Network:       cfg.network,
	}
	if cfg.dryRun {
		sdkCfg.Tx = &sdk.TxOptions{NoSend: boolPtr(true)}
	}

	client, err := sdk.NewClient(ctx, sdkCfg)
	if err != nil {
		return fmt.Errorf("init sdk client: %w", err)
	}
	defer client.Close()

	// Pick a subnet if none provided
	if !haveSubnet {
		ids, err := client.SubnetFactory.EnumerateSubnets(ctx)
		if err != nil {
			return fmt.Errorf("enumerate subnets: %w", err)
		}
		for _, id := range ids {
			active, _ := client.SubnetFactory.IsSubnetActive(ctx, id)
			if active {
				subnetID = id
				haveSubnet = true
				log.Printf("SUBNET_ID not set; using active subnet %#x", subnetID)
				break
			}
		}
		if !haveSubnet {
			return errors.New("no active subnet found; set SUBNET_ID")
		}
	}

	params := sdk.SubmitIntentParams{
		IntentID:     intentID,
		SubnetID:     subnetID,
		IntentType:   cfg.intentType,
		ParamsHash:   paramsHash,
		Deadline:     deadline,
		PaymentToken: paymentToken,
		Amount:       amount,
		Value:        value,
	}

	// fmt.Printf("submitting intent: %#v\n", params)

	tx, err := client.IntentManager.SubmitIntent(ctx, params)
	if err != nil {
		return fmt.Errorf("submit intent: %w", err)
	}

	if cfg.dryRun {
		log.Printf("dry-run enabled: transaction prepared with nonce %d", tx.Nonce())
		log.Printf("gas limit=%d maxFeePerGas=%s maxPriorityFeePerGas=%s", tx.Gas(), tx.GasFeeCap(), tx.GasTipCap())
	} else {
		log.Printf("broadcasted transaction: %s", tx.Hash())
	}

	// --- Submit by signature (1 item) ---
	// Use a different intent id to avoid potential duplication
	var intentID2 [32]byte
	if v := strings.TrimSpace(os.Getenv("SIGNED_INTENT_ID")); v != "" {
		if parsed, perr := sdk.Bytes32FromHex(v); perr == nil {
			intentID2 = parsed
		} else {
			return fmt.Errorf("parse SIGNED_INTENT_ID: %w", perr)
		}
	} else {
		rb := make([]byte, 32)
		if _, err := rand.Read(rb); err != nil {
			return fmt.Errorf("rand signed-intent-id: %w", err)
		}
		intentID2, _ = sdk.Bytes32FromHex("0x" + hex.EncodeToString(rb))
		log.Printf("SIGNED_INTENT_ID not set; generated %s", "0x"+hex.EncodeToString(rb))
	}

	input := sdkcrypto.SignedIntentInput{
		IntentID:     intentID2,
		SubnetID:     subnetID,
		Requester:    client.Signer.Address(),
		IntentType:   cfg.intentType,
		ParamsHash:   paramsHash,
		Deadline:     deadline,
		PaymentToken: paymentToken,
		Amount:       amount,
	}
	digest, err := client.IntentManager.ComputeDigest(input)
	if err != nil {
		return fmt.Errorf("compute digest: %w", err)
	}
	sig, err := client.IntentManager.SignDigest(digest)
	if err != nil {
		return fmt.Errorf("sign digest: %w", err)
	}

	btx, err := client.IntentManager.SubmitIntentsBySignatures(ctx, sdk.SubmitIntentBatchParams{
		Items: []sdk.SignedIntent{{Data: input, Signature: sig}},
		// TotalEthValue nil -> SDK auto sums ETH when PaymentToken==0
	})
	if err != nil {
		return fmt.Errorf("submit by signatures: %w", err)
	}

	if cfg.dryRun {
		log.Printf("(batch) dry-run: tx nonce %d, gas=%d", btx.Nonce(), btx.Gas())
	} else {
		log.Printf("(batch) broadcasted transaction: %s", btx.Hash())
	}
	return nil
}

func envOr(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}
func envBool(key string, fallback bool) bool {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	switch strings.ToLower(v) {
	case "1", "true", "yes", "y":
		return true
	case "0", "false", "no", "n":
		return false
	default:
		return fallback
	}
}
func envDuration(key string, fallback time.Duration) time.Duration {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return fallback
	}
	return d
}

func deriveParamsHash(hexValue, jsonPayload string) ([32]byte, error) {
	if strings.TrimSpace(hexValue) != "" {
		out, err := sdk.Bytes32FromHex(hexValue)
		if err != nil {
			return out, fmt.Errorf("parse params-hash: %w", err)
		}
		return out, nil
	}
	return sdk.HashBytes([]byte(jsonPayload)), nil
}

func parseBigInt(input string) (*big.Int, error) {
	val, err := parseOptionalBigInt(input)
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, errors.New("empty value")
	}
	return val, nil
}

func parseOptionalBigInt(input string) (*big.Int, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return nil, nil
	}
	if strings.HasPrefix(trimmed, "0x") || strings.HasPrefix(trimmed, "0X") {
		trimmed = trimmed[2:]
		if trimmed == "" {
			return nil, errors.New("invalid hex big integer")
		}
		val, ok := new(big.Int).SetString(trimmed, 16)
		if !ok {
			return nil, errors.New("invalid hex big integer")
		}
		return val, nil
	}
	val, ok := new(big.Int).SetString(trimmed, 10)
	if !ok {
		return nil, errors.New("invalid decimal big integer")
	}
	return val, nil
}

func parseAddress(hexAddr string) (common.Address, error) {
	trimmed := strings.TrimSpace(hexAddr)
	if trimmed == "" {
		return common.Address{}, nil
	}
	if !common.IsHexAddress(trimmed) {
		return common.Address{}, fmt.Errorf("invalid payment-token address: %s", hexAddr)
	}
	return common.HexToAddress(trimmed), nil
}

func boolPtr(v bool) *bool {
	return &v
}

// Minimal .env loader: supports KEY=VALUE, ignores comments and blanks.
func loadDotEnv(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// allow inline comments only if not within quotes is tricky; keep simple
		if i := strings.Index(line, "="); i >= 0 {
			k := strings.TrimSpace(line[:i])
			v := strings.TrimSpace(line[i+1:])
			v = strings.Trim(v, "\"'")
			os.Setenv(k, v)
		}
	}
	return s.Err()
}
