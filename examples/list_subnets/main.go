package main

import (
	"bufio"
	"context"
	"flag"
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

	rpc := envOr("RPC_URL", "")
	pk := envOr("PRIVATE_KEY", "")
	net := envOr("PIN_NETWORK", "")
	flag.StringVar(&rpc, "rpc", rpc, "RPC URL")
	flag.StringVar(&pk, "private-key", pk, "private key")
	flag.StringVar(&net, "network", net, "network alias")
	timeout := flag.Duration("timeout", 20*time.Second, "RPC timeout")
	flag.Parse()

	if rpc == "" || pk == "" {
		fmt.Printf("rpc=%s pk=%s net=%s\n", rpc, pk, net)
		log.Fatal("missing RPC_URL or PRIVATE_KEY")
	}
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	client, err := sdk.NewClient(ctx, sdk.Config{RPCURL: rpc, PrivateKeyHex: pk, Network: net})
	if err != nil {
		log.Fatalf("init sdk: %v", err)
	}
	defer client.Close()

	total, err := client.SubnetFactory.TotalCreated(ctx)
	if err != nil {
		log.Fatalf("total: %v", err)
	}
	activeCount, err := client.SubnetFactory.GetActiveSubnetCount(ctx)
	if err != nil {
		log.Fatalf("active count: %v", err)
	}
	fmt.Printf("total created: %s, active: %s\n", total.String(), activeCount.String())

	// 使用 GetSubnetsByStatus 获取 ACTIVE 状态的子网 (status=1)
	const SubnetStatusActive = uint8(1)
	ids, err := client.SubnetFactory.GetSubnetsByStatus(ctx, SubnetStatusActive)
	if err != nil {
		log.Fatalf("get active subnets: %v", err)
	}

	if len(ids) == 0 {
		fmt.Println("\nNo active subnets found.")
		return
	}

	var participants []subnetfactory.DataStructuresSubnetParticipants
	participants, err = client.SubnetFactory.GetActiveParticipants(ctx, ids)
	if err != nil {
		log.Printf("warning: failed to fetch participants: %v", err)
	}

	fmt.Printf("\nFound %d active subnet(s):\n", len(ids))

	for i, id := range ids {
		info, _ := client.SubnetFactory.GetSubnetInfo(ctx, id)

		matcherCount := "0"
		if i < len(participants) {
			matcherCount = fmt.Sprintf("%d", len(participants[i].Matchers))
		}

		fmt.Printf("\n%d) id=%#x validators=%s agents=%s matchers=%s status=%d\n",
			i, id, bi(info.ValidatorCount), bi(info.AgentCount), matcherCount, info.Status)

		// 展示详细参与者信息
		if i < len(participants) {
			printParticipantsCompact(participants[i])
		}
	}
}

func envOr(k, d string) string {
	v := strings.TrimSpace(os.Getenv(k))
	if v == "" {
		return d
	}
	return v
}
func bi(x *big.Int) string {
	if x == nil {
		return "0"
	}
	return x.String()
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

func printParticipantsCompact(p subnetfactory.DataStructuresSubnetParticipants) {
	fmt.Printf("   Validators: %d active\n", len(p.Validators))
	for i, v := range p.Validators {
		if i >= 3 {
			fmt.Printf("      ... and %d more\n", len(p.Validators)-3)
			break
		}
		fmt.Printf("      - %s (rep: %s)\n", v.Owner.Hex(), v.ReputationScore.String())
	}

	fmt.Printf("   Agents: %d active\n", len(p.Agents))
	for i, a := range p.Agents {
		if i >= 3 {
			fmt.Printf("      ... and %d more\n", len(p.Agents)-3)
			break
		}
		endpoint := a.Endpoint
		if endpoint == "" {
			endpoint = "(no endpoint)"
		}
		fmt.Printf("      - %s endpoint=%s\n", a.Owner.Hex(), endpoint)
	}

	fmt.Printf("   Matchers: %d active\n", len(p.Matchers))
	for i, m := range p.Matchers {
		if i >= 3 {
			fmt.Printf("      ... and %d more\n", len(p.Matchers)-3)
			break
		}
		fmt.Printf("      - %s\n", m.Owner.Hex())
	}
}
