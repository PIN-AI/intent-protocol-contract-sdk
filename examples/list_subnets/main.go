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
		validatorCount := "0"
		agentCount := "0"
		matcherCount := "0"

		// 直接从 GetActiveParticipants 结果获取数量，避免额外 RPC 调用
		if i < len(participants) {
			validatorCount = fmt.Sprintf("%d", len(participants[i].Validators))
			agentCount = fmt.Sprintf("%d", len(participants[i].Agents))
			matcherCount = fmt.Sprintf("%d", len(participants[i].Matchers))
		}

		fmt.Printf("\n%d) Subnet ID: %#x\n", i, id)
		fmt.Printf("   Summary: validators=%s agents=%s matchers=%s\n",
			validatorCount, agentCount, matcherCount)

		// 展示详细参与者信息
		if i < len(participants) {
			printParticipantsDetail(participants[i])
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

func printParticipantsDetail(p subnetfactory.DataStructuresSubnetParticipants) {
	// Validators
	if len(p.Validators) > 0 {
		fmt.Printf("\n   Validators (%d active):\n", len(p.Validators))
		for i, v := range p.Validators {
			if i >= 5 {
				fmt.Printf("      ... and %d more\n", len(p.Validators)-5)
				break
			}
			fmt.Printf("      [%d] Address: %s\n", i+1, v.Owner.Hex())
			fmt.Printf("          Reputation: %s, Registered: %s, Status: %d\n",
				v.ReputationScore.String(), v.RegisteredAt.String(), v.Status)
			if v.Endpoint != "" {
				fmt.Printf("          Endpoint: %s\n", v.Endpoint)
			}
			if v.Domain != "" {
				fmt.Printf("          Domain: %s\n", v.Domain)
			}
		}
	}

	// Agents
	if len(p.Agents) > 0 {
		fmt.Printf("\n   Agents (%d active):\n", len(p.Agents))
		for i, a := range p.Agents {
			if i >= 5 {
				fmt.Printf("      ... and %d more\n", len(p.Agents)-5)
				break
			}
			fmt.Printf("      [%d] Address: %s\n", i+1, a.Owner.Hex())
			fmt.Printf("          Reputation: %s, Registered: %s, Status: %d\n",
				a.ReputationScore.String(), a.RegisteredAt.String(), a.Status)
			if a.Endpoint != "" {
				fmt.Printf("          Endpoint: %s\n", a.Endpoint)
			}
			if a.Domain != "" {
				fmt.Printf("          Domain: %s\n", a.Domain)
			}
		}
	}

	// Matchers
	if len(p.Matchers) > 0 {
		fmt.Printf("\n   Matchers (%d active):\n", len(p.Matchers))
		for i, m := range p.Matchers {
			if i >= 5 {
				fmt.Printf("      ... and %d more\n", len(p.Matchers)-5)
				break
			}
			fmt.Printf("      [%d] Address: %s\n", i+1, m.Owner.Hex())
			fmt.Printf("          Reputation: %s, Registered: %s, Status: %d\n",
				m.ReputationScore.String(), m.RegisteredAt.String(), m.Status)
			if m.Endpoint != "" {
				fmt.Printf("          Endpoint: %s\n", m.Endpoint)
			}
			if m.Domain != "" {
				fmt.Printf("          Domain: %s\n", m.Domain)
			}
		}
	}
}
