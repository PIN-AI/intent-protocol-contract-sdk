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
	fmt.Printf("total created: %s\n", total.String())

	ids, err := client.SubnetFactory.EnumerateSubnets(ctx)
	if err != nil {
		log.Fatalf("enumerate: %v", err)
	}
	for i, id := range ids {
		active, _ := client.SubnetFactory.IsSubnetActive(ctx, id)
		info, _ := client.SubnetFactory.GetSubnetInfo(ctx, id)
		fmt.Printf("%d) id=%#x active=%v validators=%s agents=%s status=%d\n", i, id, active, bi(info.ValidatorCount), bi(info.AgentCount), info.Status)
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
