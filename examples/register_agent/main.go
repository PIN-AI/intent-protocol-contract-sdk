package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/PIN-AI/intent-protocol-contract-sdk/contracts/subnetfactory"
	sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
)

func main() {
	_ = loadDotEnv(".env")

	cfg := runConfig{
		rpcURL:      strings.TrimSpace(envOr("PIN_RPC_URL", envOr("RPC_URL", ""))),
		privateKey:  strings.TrimSpace(envOr("PIN_PRIVATE_KEY", envOr("PRIVATE_KEY", ""))),
		network:     strings.TrimSpace(envOr("PIN_NETWORK", "")),
		subnetIDHex: strings.TrimSpace(os.Getenv("SUBNET_ID")),
		domain:      strings.TrimSpace(envOr("AGENT_DOMAIN", "agent.example")),
		endpoint:    strings.TrimSpace(envOr("AGENT_ENDPOINT", "https://agent.example/api")),
		metadataURI: strings.TrimSpace(envOr("AGENT_METADATA_URI", "ipfs://agent-metadata")),
		// 0.1 eth in wei
		valueWei: strings.TrimSpace(envOr("AGENT_VALUE_WEI", "100000000000000000")),
		dryRun:   envBool("DRY_RUN", false),
		timeout:  envDuration("TIMEOUT", 45*time.Second),
	}

	if err := runRegisterAgent(cfg); err != nil {
		log.Fatalf("register agent: %v", err)
	}
}

type runConfig struct {
	rpcURL      string
	privateKey  string
	network     string
	subnetIDHex string
	domain      string
	endpoint    string
	metadataURI string
	valueWei    string
	dryRun      bool
	timeout     time.Duration
}

func runRegisterAgent(cfg runConfig) error {
	if cfg.rpcURL == "" {
		return errors.New("missing RPC endpoint: set PIN_RPC_URL or RPC_URL")
	}
	if cfg.privateKey == "" {
		return errors.New("missing private key: set PIN_PRIVATE_KEY or PRIVATE_KEY")
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

	// locate subnet id
	var subnetID [32]byte
	haveSubnet := false
	if cfg.subnetIDHex != "" {
		subnetID, err = sdk.Bytes32FromHex(cfg.subnetIDHex)
		if err != nil {
			return fmt.Errorf("parse SUBNET_ID: %w", err)
		}
		haveSubnet = true
	}

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

	subnetAddr, err := client.SubnetFactory.GetSubnetContract(ctx, subnetID)
	if err != nil {
		return fmt.Errorf("fetch subnet contract: %w", err)
	}

	subnetSvc, err := sdk.NewSubnetServiceByAddress(client.Backend, subnetAddr, client.TxManager)
	if err != nil {
		return fmt.Errorf("bind subnet service: %w", err)
	}

	participants, err := client.SubnetFactory.GetActiveParticipants(ctx, [][32]byte{subnetID})
	if err != nil {
		return fmt.Errorf("fetch active participants: %w", err)
	}
	if len(participants) > 0 {
		printParticipants(participants[0])
	}

	value, err := parseOptionalBigInt(cfg.valueWei)
	if err != nil {
		return fmt.Errorf("parse AGENT_VALUE_WEI: %w", err)
	}

	params := sdk.RegisterParticipantParams{
		Domain:      cfg.domain,
		Endpoint:    cfg.endpoint,
		MetadataURI: cfg.metadataURI,
		Value:       value,
	}

	tx, err := subnetSvc.RegisterAgent(ctx, params)
	if err != nil {
		return fmt.Errorf("register agent: %w", err)
	}

	log.Printf("subnet contract: %s", subnetAddr)
	log.Printf("agent domain=%q endpoint=%q metadata=%q", params.Domain, params.Endpoint, params.MetadataURI)

	if cfg.dryRun {
		log.Printf("dry-run enabled: tx nonce %d gas=%d", tx.Nonce(), tx.Gas())
	} else {
		log.Printf("broadcasted transaction: %s", tx.Hash())
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

func boolPtr(v bool) *bool { return &v }

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

func printParticipants(p subnetfactory.DataStructuresSubnetParticipants) {
	log.Printf("========== Subnet Participants ==========")
	log.Printf("Subnet ID: %#x", p.SubnetId)
	log.Printf("")

	log.Printf("--- Validators (%d) ---", len(p.Validators))
	for i, v := range p.Validators {
		printParticipantInfo(i+1, v)
	}

	log.Printf("")
	log.Printf("--- Agents (%d) ---", len(p.Agents))
	for i, a := range p.Agents {
		printParticipantInfo(i+1, a)
	}

	log.Printf("")
	log.Printf("--- Matchers (%d) ---", len(p.Matchers))
	for i, m := range p.Matchers {
		printParticipantInfo(i+1, m)
	}
	log.Printf("====================================")
}

func printParticipantInfo(idx int, info subnetfactory.DataStructuresParticipantInfo) {
	log.Printf("  [%d] Owner: %s", idx, info.Owner.Hex())
	log.Printf("      Status: %d | Type: %d | Reputation: %s", info.Status, info.ParticipantType, info.ReputationScore.String())
	if info.Endpoint != "" {
		log.Printf("      Endpoint: %s", info.Endpoint)
	}
	if info.Domain != "" {
		log.Printf("      Domain: %s", info.Domain)
	}
	if info.MetadataUri != "" {
		log.Printf("      Metadata: %s", info.MetadataUri)
	}
}
