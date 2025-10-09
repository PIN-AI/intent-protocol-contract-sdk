package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	_ = loadDotEnv(".env")

	cfg := runConfig{
		rpcURL:          strings.TrimSpace(envOr("PIN_RPC_URL", envOr("RPC_URL", ""))),
		privateKey:      strings.TrimSpace(envOr("PIN_PRIVATE_KEY", envOr("PRIVATE_KEY", ""))),
		network:         strings.TrimSpace(envOr("PIN_NETWORK", "")),
		intentIDHex:     strings.TrimSpace(os.Getenv("INTENT_ID")),
		assignmentIDHex: strings.TrimSpace(os.Getenv("ASSIGNMENT_ID")),
		subnetIDHex:     strings.TrimSpace(os.Getenv("SUBNET_ID")),
		agentAddr:       strings.TrimSpace(os.Getenv("AGENT_ADDRESS")),
		resultHashHex:   strings.TrimSpace(os.Getenv("RESULT_HASH")),
		proofHashHex:    strings.TrimSpace(os.Getenv("PROOF_HASH")),
		rootHeightStr:   strings.TrimSpace(envOr("ROOT_HEIGHT", "0")),
		rootHashHex:     strings.TrimSpace(os.Getenv("ROOT_HASH")),
		validatorList:   strings.TrimSpace(os.Getenv("VALIDATORS")),
		signatureList:   strings.TrimSpace(os.Getenv("VALIDATOR_SIGNATURES")),
		dryRun:          envBool("DRY_RUN", true),
		timeout:         envDuration("TIMEOUT", 45*time.Second),
	}

	if err := runValidateIntents(cfg); err != nil {
		log.Fatalf("validate intents: %v", err)
	}
}

type runConfig struct {
	rpcURL          string
	privateKey      string
	network         string
	intentIDHex     string
	assignmentIDHex string
	subnetIDHex     string
	agentAddr       string
	resultHashHex   string
	proofHashHex    string
	rootHeightStr   string
	rootHashHex     string
	validatorList   string
	signatureList   string
	dryRun          bool
	timeout         time.Duration
}

func runValidateIntents(cfg runConfig) error {
	if cfg.rpcURL == "" {
		return errors.New("missing RPC endpoint: set PIN_RPC_URL or RPC_URL in env/.env")
	}
	if cfg.privateKey == "" {
		return errors.New("missing private key: set PIN_PRIVATE_KEY or PRIVATE_KEY in env/.env")
	}
	if cfg.intentIDHex == "" {
		return errors.New("missing INTENT_ID")
	}
	if cfg.assignmentIDHex == "" {
		return errors.New("missing ASSIGNMENT_ID")
	}
	if cfg.subnetIDHex == "" {
		return errors.New("missing SUBNET_ID")
	}
	if cfg.agentAddr == "" {
		return errors.New("missing AGENT_ADDRESS")
	}
	if cfg.resultHashHex == "" {
		return errors.New("missing RESULT_HASH")
	}
	if cfg.proofHashHex == "" {
		return errors.New("missing PROOF_HASH")
	}
	if cfg.rootHashHex == "" {
		return errors.New("missing ROOT_HASH")
	}

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

	intentID, err := parseBytes32(cfg.intentIDHex, "INTENT_ID")
	if err != nil {
		return err
	}
	assignmentID, err := parseBytes32(cfg.assignmentIDHex, "ASSIGNMENT_ID")
	if err != nil {
		return err
	}
	subnetID, err := parseBytes32(cfg.subnetIDHex, "SUBNET_ID")
	if err != nil {
		return err
	}
	agent, err := parseAddressRequired(cfg.agentAddr, "AGENT_ADDRESS")
	if err != nil {
		return err
	}
	resultHash, err := parseBytes32(cfg.resultHashHex, "RESULT_HASH")
	if err != nil {
		return err
	}
	proofHash, err := parseBytes32(cfg.proofHashHex, "PROOF_HASH")
	if err != nil {
		return err
	}
	rootHash, err := parseBytes32(cfg.rootHashHex, "ROOT_HASH")
	if err != nil {
		return err
	}
	rootHeight, err := parseUint64(cfg.rootHeightStr)
	if err != nil {
		return fmt.Errorf("parse ROOT_HEIGHT: %w", err)
	}

	bundle := sdk.ValidationBundle{
		IntentID:     intentID,
		AssignmentID: assignmentID,
		SubnetID:     subnetID,
		Agent:        agent,
		ResultHash:   resultHash,
		ProofHash:    proofHash,
		RootHeight:   rootHeight,
		RootHash:     rootHash,
	}

	digest, err := client.Validation.ComputeDigest(bundle)
	if err != nil {
		return fmt.Errorf("compute digest: %w", err)
	}

	validators, signatures, err := prepareSignatures(cfg, client, digest)
	if err != nil {
		return err
	}
	bundle.Validators = validators
	bundle.Signatures = signatures

	tx, err := client.Validation.ValidateIntentsBySignatures(ctx, []sdk.ValidationBundle{bundle})
	if err != nil {
		return fmt.Errorf("submit validation bundle: %w", err)
	}

	log.Printf("validation digest=%s", "0x"+hex.EncodeToString(digest[:]))
	log.Printf("intent id=%#x assignment id=%#x subnet id=%#x", bundle.IntentID, bundle.AssignmentID, bundle.SubnetID)
	log.Printf("validators=%s", joinAddresses(bundle.Validators))

	if cfg.dryRun {
		log.Printf("dry-run enabled: tx nonce %d gas=%d", tx.Nonce(), tx.Gas())
	} else {
		log.Printf("broadcasted transaction: %s", tx.Hash())
	}
	return nil
}

func prepareSignatures(cfg runConfig, client *sdk.Client, digest [32]byte) ([]common.Address, [][]byte, error) {
	signerAddr := client.Signer.Address()
	addrInputs := splitCommaSeparated(cfg.validatorList)

	var validators []common.Address
	for _, addr := range addrInputs {
		if addr == "" {
			continue
		}
		parsed, err := parseAddressRequired(addr, "VALIDATORS")
		if err != nil {
			return nil, nil, err
		}
		validators = append(validators, parsed)
	}
	if len(validators) == 0 {
		validators = []common.Address{signerAddr}
	}

	sigInputs := splitCommaSeparated(cfg.signatureList)
	if len(sigInputs) > 0 {
		if len(sigInputs) != len(validators) {
			return nil, nil, fmt.Errorf("validator/signature count mismatch: %d vs %d", len(validators), len(sigInputs))
		}
		sigs := make([][]byte, len(sigInputs))
		for i, raw := range sigInputs {
			parsed, err := parseSignature(raw)
			if err != nil {
				return nil, nil, fmt.Errorf("parse signature %d: %w", i, err)
			}
			sigs[i] = parsed
		}
		return validators, sigs, nil
	}

	if len(validators) != 1 {
		return nil, nil, errors.New("provide VALIDATOR_SIGNATURES when using multiple validators")
	}
	if validators[0] != signerAddr {
		return nil, nil, fmt.Errorf("validator %s does not match local signer %s; provide VALIDATOR_SIGNATURES", validators[0].Hex(), signerAddr.Hex())
	}

	sig, err := client.Validation.SignDigest(digest)
	if err != nil {
		return nil, nil, fmt.Errorf("sign digest: %w", err)
	}
	return validators, [][]byte{sig}, nil
}

func parseSignature(input string) ([]byte, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return nil, errors.New("empty signature")
	}
	trimmed = strings.TrimPrefix(trimmed, "0x")
	if trimmed == "" {
		return nil, errors.New("invalid signature hex")
	}
	out, err := hex.DecodeString(trimmed)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func splitCommaSeparated(input string) []string {
	if strings.TrimSpace(input) == "" {
		return nil
	}
	parts := strings.Split(input, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		out = append(out, strings.TrimSpace(p))
	}
	return out
}

func joinAddresses(addrs []common.Address) string {
	out := make([]string, len(addrs))
	for i, addr := range addrs {
		out[i] = addr.Hex()
	}
	return strings.Join(out, ", ")
}

func parseUint64(input string) (uint64, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return 0, nil
	}
	if strings.HasPrefix(trimmed, "0x") || strings.HasPrefix(trimmed, "0X") {
		return strconv.ParseUint(trimmed[2:], 16, 64)
	}
	return strconv.ParseUint(trimmed, 10, 64)
}

func parseBytes32(hexValue, field string) ([32]byte, error) {
	val, err := sdk.Bytes32FromHex(hexValue)
	if err != nil {
		return val, fmt.Errorf("parse %s: %w", field, err)
	}
	return val, nil
}

func parseAddressRequired(hexAddr, field string) (common.Address, error) {
	trimmed := strings.TrimSpace(hexAddr)
	if trimmed == "" {
		return common.Address{}, fmt.Errorf("missing %s", field)
	}
	if !common.IsHexAddress(trimmed) {
		return common.Address{}, fmt.Errorf("invalid %s: %s", field, hexAddr)
	}
	return common.HexToAddress(trimmed), nil
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

func boolPtr(v bool) *bool {
	return &v
}

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
		if i := strings.Index(line, "="); i >= 0 {
			k := strings.TrimSpace(line[:i])
			v := strings.TrimSpace(line[i+1:])
			v = strings.Trim(v, "\"'")
			os.Setenv(k, v)
		}
	}
	return s.Err()
}
