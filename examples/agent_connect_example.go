package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
	cryptoHelpers "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
)

func main() {
	fmt.Println("=== PIN AI Agent Connect Signature Example ===")
	fmt.Println()

	// Initialize SDK client
	client, err := sdk.NewClient(context.Background(), sdk.Config{
		RPCURL:        os.Getenv("PIN_RPC_URL"),
		PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	agentAddr := client.Signer.Address()
	fmt.Printf("Agent Address: %s\n\n", agentAddr.Hex())

	// Method 1: Quick Start (Recommended)
	fmt.Println("Method 1: One-step signature generation")
	fmt.Println("---------------------------------------")
	signature1, timestamp1, nonce1, err := client.SignAgentConnectNow(agentAddr)
	if err != nil {
		log.Fatalf("SignAgentConnectNow failed: %v", err)
	}
	fmt.Printf("Signature:  %x\n", signature1)
	fmt.Printf("Timestamp:  %d\n", timestamp1)
	fmt.Printf("Nonce:      %x\n\n", nonce1)

	// Method 2: Manual Control
	fmt.Println("Method 2: Manual timestamp and nonce")
	fmt.Println("-------------------------------------")
	timestamp2 := time.Now().Unix()
	var nonce2 [32]byte
	if _, err := rand.Read(nonce2[:]); err != nil {
		log.Fatalf("Failed to generate nonce: %v", err)
	}

	signature2, err := client.SignAgentConnect(agentAddr, timestamp2, &nonce2)
	if err != nil {
		log.Fatalf("SignAgentConnect failed: %v", err)
	}
	fmt.Printf("Signature:  %x\n", signature2)
	fmt.Printf("Timestamp:  %d\n", timestamp2)
	fmt.Printf("Nonce:      %x\n\n", nonce2)

	// Method 3: Low-Level API (for advanced users)
	fmt.Println("Method 3: Low-level digest computation")
	fmt.Println("---------------------------------------")
	timestamp3 := time.Now().Unix()
	var nonce3 [32]byte
	if _, err := rand.Read(nonce3[:]); err != nil {
		log.Fatalf("Failed to generate nonce: %v", err)
	}

	input := cryptoHelpers.AgentConnectInput{
		AgentAddress: agentAddr,
		Timestamp:    big.NewInt(timestamp3),
		RandomNonce:  nonce3,
	}

	digest, err := cryptoHelpers.ComputeAgentConnectDigest(input)
	if err != nil {
		log.Fatalf("ComputeAgentConnectDigest failed: %v", err)
	}

	signature3, err := client.Signer.SignDigest(digest)
	if err != nil {
		log.Fatalf("SignDigest failed: %v", err)
	}

	fmt.Printf("Digest:     %x\n", digest)
	fmt.Printf("Signature:  %x\n", signature3)
	fmt.Printf("Timestamp:  %d\n", timestamp3)
	fmt.Printf("Nonce:      %x\n\n", nonce3)

	// Example: Prepare request for server
	fmt.Println("Example: Server Request Format")
	fmt.Println("-------------------------------")
	fmt.Printf("{\n")
	fmt.Printf("  \"agent_address\": \"%s\",\n", agentAddr.Hex())
	fmt.Printf("  \"signature\": \"%x\",\n", signature1)
	fmt.Printf("  \"timestamp\": %d,\n", timestamp1)
	fmt.Printf("  \"random_nonce\": \"%x\",\n", nonce1)
	fmt.Printf("  \"client_version\": \"v1.0.0\"\n")
	fmt.Printf("}\n")
}
