package main

import (
	"bufio"
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
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
		bidIDHex:        strings.TrimSpace(os.Getenv("BID_ID")),
		agentAddr:       strings.TrimSpace(os.Getenv("AGENT_ADDRESS")),
		matcherAddr:     strings.TrimSpace(os.Getenv("MATCHER_ADDRESS")),
		statusInput:     strings.TrimSpace(envOr("ASSIGNMENT_STATUS", "active")),
		dryRun:          envBool("DRY_RUN", true),
		timeout:         envDuration("TIMEOUT", 45*time.Second),
	}

	if err := runAssignIntents(cfg); err != nil {
		log.Fatalf("assign intents: %v", err)
	}
}

type runConfig struct {
	rpcURL          string
	privateKey      string
	network         string
	intentIDHex     string
	assignmentIDHex string
	bidIDHex        string
	agentAddr       string
	matcherAddr     string
	statusInput     string
	dryRun          bool
	timeout         time.Duration
}

func runAssignIntents(cfg runConfig) error {
	if cfg.rpcURL == "" {
		return errors.New("missing RPC endpoint: set PIN_RPC_URL or RPC_URL in env/.env")
	}
	if cfg.privateKey == "" {
		return errors.New("missing private key: set PIN_PRIVATE_KEY or PRIVATE_KEY in env/.env")
	}
	if cfg.intentIDHex == "" {
		return errors.New("missing INTENT_ID: provide the intent to assign")
	}
	if cfg.agentAddr == "" {
		return errors.New("missing AGENT_ADDRESS: provide the assigned agent address")
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

	assignmentID, err := bytes32OrRandom(cfg.assignmentIDHex, "ASSIGNMENT_ID")
	if err != nil {
		return err
	}

	bidID, err := bytes32OrRandom(cfg.bidIDHex, "BID_ID")
	if err != nil {
		return err
	}

	agent, err := parseAddressRequired(cfg.agentAddr, "AGENT_ADDRESS")
	if err != nil {
		return err
	}

	matcher := client.Signer.Address()
	if cfg.matcherAddr != "" {
		matcher, err = parseAddressRequired(cfg.matcherAddr, "MATCHER_ADDRESS")
		if err != nil {
			return err
		}
	}

	status, err := parseAssignmentStatus(cfg.statusInput)
	if err != nil {
		return fmt.Errorf("parse ASSIGNMENT_STATUS: %w", err)
	}

	data := sdk.AssignmentData{
		AssignmentID: assignmentID,
		IntentID:     intentID,
		BidID:        bidID,
		Agent:        agent,
		Status:       status,
		Matcher:      matcher,
	}

	digest, err := client.Assignment.ComputeDigest(data)
	if err != nil {
		return fmt.Errorf("compute digest: %w", err)
	}

	sig, err := client.Assignment.SignDigest(digest)
	if err != nil {
		return fmt.Errorf("sign digest: %w", err)
	}

	tx, err := client.Assignment.AssignIntentsBySignatures(ctx, []sdk.SignedAssignment{{Data: data, Signature: sig}})
	if err != nil {
		return fmt.Errorf("submit assignments: %w", err)
	}

	log.Printf("assignment id=%#x intent id=%#x bid id=%#x", data.AssignmentID, data.IntentID, data.BidID)
	log.Printf("agent=%s matcher=%s status=%d", data.Agent.Hex(), data.Matcher.Hex(), data.Status)

	if cfg.dryRun {
		log.Printf("dry-run enabled: tx nonce %d gas=%d", tx.Nonce(), tx.Gas())
	} else {
		log.Printf("broadcasted transaction: %s", tx.Hash())
	}

	return nil
}

func bytes32OrRandom(hexValue, field string) ([32]byte, error) {
	if strings.TrimSpace(hexValue) == "" {
		randVal, err := randomBytes32()
		if err != nil {
			return randVal, fmt.Errorf("generate %s: %w", field, err)
		}
		log.Printf("%s not set; generated %s", field, "0x"+hex.EncodeToString(randVal[:]))
		return randVal, nil
	}
	return parseBytes32(hexValue, field)
}

func parseBytes32(hexValue, field string) ([32]byte, error) {
	val, err := sdk.Bytes32FromHex(hexValue)
	if err != nil {
		return val, fmt.Errorf("parse %s: %w", field, err)
	}
	return val, nil
}

func randomBytes32() ([32]byte, error) {
	var out [32]byte
	if _, err := rand.Read(out[:]); err != nil {
		return out, err
	}
	return out, nil
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

func parseAssignmentStatus(input string) (sdk.AssignmentStatus, error) {
	trimmed := strings.TrimSpace(strings.ToLower(input))
	switch trimmed {
	case "", "1", "active":
		return sdk.AssignmentStatusActive, nil
	case "0", "unspecified":
		return sdk.AssignmentStatusUnspecified, nil
	case "2", "failed":
		return sdk.AssignmentStatusFailed, nil
	default:
		return 0, fmt.Errorf("unknown status %q", input)
	}
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
