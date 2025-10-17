# Quick Start Guide

**[中文文档](quickstart-zh.md)** | English

This guide demonstrates typical SDK usage patterns with complete, working examples.

## Prerequisites

- Go 1.24+ installed
- Access to an Ethereum RPC endpoint (Base Sepolia recommended for testing)
- A funded wallet with private key (for transactions)

## Installation

```bash
go get github.com/PIN-AI/intent-protocol-contract-sdk@latest
```

## Configuration

### 1. Environment Setup

Create a `.env` file (see `.env.example`):

```ini
# Network Configuration
PIN_NETWORK=base_sepolia
PIN_RPC_URL=https://sepolia.base.org
PIN_PRIVATE_KEY=0xYourPrivateKey

# Optional: Override contract addresses (uses defaults if not set)
PIN_BASE_SEPOLIA_INTENT_MANAGER=0x5555555555555555555555555555555555555555
PIN_BASE_SEPOLIA_SUBNET_FACTORY=0x6666666666666666666666666666666666666666
PIN_BASE_SEPOLIA_STAKING_MANAGER=0x7777777777777777777777777777777777777777
PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER=0x8888888888888888888888888888888888888888
```

### 2. Initialize Client

```go
package main

import (
    "context"
    "log"
    "os"

    sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
    "github.com/joho/godotenv"
)

func main() {
    // Load environment variables
    _ = godotenv.Load()

    ctx := context.Background()
    client, err := sdk.NewClient(ctx, sdk.Config{
        RPCURL:        os.Getenv("PIN_RPC_URL"),
        PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
        Network:       os.Getenv("PIN_NETWORK"), // Optional: auto-detect from chainId
    })
    if err != nil {
        log.Fatalf("Failed to initialize SDK client: %v", err)
    }
    defer client.Close()

    log.Printf("Connected to network: %s (chainId: %s)",
        client.Network, client.ChainID.String())
    log.Printf("Using address: %s", client.Signer.Address().Hex())
}
```

## Common Use Cases

### 1. Submit a Single Intent (ETH Payment)

```go
import (
    "math/big"
    "time"

    sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
    "github.com/ethereum/go-ethereum/common"
)

func submitIntent(client *sdk.Client) error {
    ctx := context.Background()

    // Prepare intent parameters
    params := sdk.SubmitIntentParams{
        IntentID:   sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000001"),
        SubnetID:   sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000100"),
        IntentType: "book_flight",
        ParamsHash: sdk.HashBytes([]byte(`{"from":"NYC","to":"LAX","date":"2025-01-15"}`)),
        Deadline:   big.NewInt(time.Now().Add(24 * time.Hour).Unix()),

        // ETH payment: use zero address and set Value
        PaymentToken: common.Address{}, // 0x0 = ETH
        Amount:       big.NewInt(1e16),  // 0.01 ETH
        Value:        big.NewInt(1e16),  // Must match Amount for ETH
    }

    // Submit transaction
    tx, err := client.Intent.SubmitIntent(ctx, params)
    if err != nil {
        return fmt.Errorf("failed to submit intent: %w", err)
    }

    log.Printf("Intent submitted successfully!")
    log.Printf("Transaction hash: %s", tx.Hash().Hex())

    // Wait for confirmation (optional)
    receipt, err := bind.WaitMined(ctx, client.Backend, tx)
    if err != nil {
        return fmt.Errorf("failed to wait for transaction: %w", err)
    }

    if receipt.Status == 1 {
        log.Printf("Transaction confirmed in block: %d", receipt.BlockNumber.Uint64())
    } else {
        return fmt.Errorf("transaction failed")
    }

    return nil
}
```

### 2. Submit Intent with ERC20 Payment

```go
func submitIntentERC20(client *sdk.Client) error {
    ctx := context.Background()

    tokenAddress := common.HexToAddress("0xYourERC20TokenAddress")
    amount := big.NewInt(1e18) // 1 token (adjust for decimals)

    // IMPORTANT: Approve IntentManager to spend tokens first
    // This example assumes approval is already done

    params := sdk.SubmitIntentParams{
        IntentID:     sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000002"),
        SubnetID:     sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000100"),
        IntentType:   "book_hotel",
        ParamsHash:   sdk.HashBytes([]byte(`{"city":"Tokyo","checkin":"2025-01-20"}`)),
        Deadline:     big.NewInt(time.Now().Add(48 * time.Hour).Unix()),

        // ERC20 payment: set token address, leave Value empty
        PaymentToken: tokenAddress,
        Amount:       amount,
        Value:        nil, // No ETH sent
    }

    tx, err := client.Intent.SubmitIntent(ctx, params)
    if err != nil {
        return fmt.Errorf("failed to submit intent: %w", err)
    }

    log.Printf("ERC20 intent submitted: %s", tx.Hash().Hex())
    return nil
}
```

### 3. Batch Submit Intents with Signatures

Batch submission is more gas-efficient for multiple intents:

```go
import (
    sdkcrypto "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
)

func batchSubmitIntents(client *sdk.Client) error {
    ctx := context.Background()

    // Prepare multiple intents
    intents := []sdkcrypto.SignedIntentInput{
        {
            IntentID:     sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000003"),
            SubnetID:     sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000100"),
            Requester:    client.Signer.Address(),
            IntentType:   "book_flight",
            ParamsHash:   sdk.MustBytes32FromHex("0x1234567890123456789012345678901234567890123456789012345678901234"),
            Deadline:     big.NewInt(time.Now().Add(24 * time.Hour).Unix()),
            PaymentToken: common.Address{},
            Amount:       big.NewInt(1e16), // 0.01 ETH
        },
        {
            IntentID:     sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000004"),
            SubnetID:     sdk.MustBytes32FromHex("0x0000000000000000000000000000000000000000000000000000000000000100"),
            Requester:    client.Signer.Address(),
            IntentType:   "book_hotel",
            ParamsHash:   sdk.MustBytes32FromHex("0x4567890123456789012345678901234567890123456789012345678901234567"),
            Deadline:     big.NewInt(time.Now().Add(48 * time.Hour).Unix()),
            PaymentToken: common.Address{},
            Amount:       big.NewInt(2e16), // 0.02 ETH
        },
    }

    // Sign each intent
    signedIntents := make([]sdk.SignedIntent, len(intents))
    for i, input := range intents {
        // Compute digest
        digest, err := client.Intent.ComputeDigest(input)
        if err != nil {
            return fmt.Errorf("failed to compute digest for intent %d: %w", i, err)
        }

        // Sign digest
        sig, err := client.Intent.SignDigest(digest)
        if err != nil {
            return fmt.Errorf("failed to sign intent %d: %w", i, err)
        }

        signedIntents[i] = sdk.SignedIntent{
            Data:      input,
            Signature: sig,
        }
    }

    // Submit batch (SDK auto-calculates total ETH value)
    tx, err := client.Intent.SubmitIntentsBySignatures(ctx, sdk.SubmitIntentBatchParams{
        Items: signedIntents,
        // TotalEthValue optional: SDK sums all PaymentToken==0 amounts
    })
    if err != nil {
        return fmt.Errorf("failed to submit batch: %w", err)
    }

    log.Printf("Batch submitted successfully: %s", tx.Hash().Hex())
    return nil
}
```

### 4. Query Intents

```go
func queryIntent(client *sdk.Client, intentID [32]byte) error {
    ctx := context.Background()

    // Check if intent exists
    exists, err := client.Intent.IntentExists(ctx, intentID)
    if err != nil {
        return fmt.Errorf("failed to check intent existence: %w", err)
    }

    if !exists {
        log.Printf("Intent %x does not exist", intentID)
        return nil
    }

    // Get intent details
    intent, err := client.Intent.GetIntent(ctx, intentID)
    if err != nil {
        return fmt.Errorf("failed to get intent: %w", err)
    }

    log.Printf("Intent Details:")
    log.Printf("  Requester: %s", intent.Requester.Hex())
    log.Printf("  SubnetID: %x", intent.SubnetID)
    log.Printf("  IntentType: %s", intent.IntentType)
    log.Printf("  Status: %d", intent.Status)
    log.Printf("  Amount: %s", intent.Amount.String())
    log.Printf("  Deadline: %s", time.Unix(intent.Deadline.Int64(), 0))

    return nil
}
```

### 5. Register as a Validator

```go
func registerValidator(client *sdk.Client, subnetID [32]byte) error {
    ctx := context.Background()

    // Get subnet service
    subnetSvc, err := client.SubnetServiceByID(ctx, subnetID)
    if err != nil {
        return fmt.Errorf("failed to get subnet service: %w", err)
    }

    // Register with ETH stake
    tx, err := subnetSvc.RegisterValidator(ctx, sdk.RegisterParticipantParams{
        Domain:      "validator.example.com",
        Endpoint:    "https://validator.example.com/api",
        MetadataURI: "ipfs://QmValidator123...",
        Value:       big.NewInt(1e18), // 1 ETH stake
    })
    if err != nil {
        return fmt.Errorf("failed to register validator: %w", err)
    }

    log.Printf("Validator registered: %s", tx.Hash().Hex())

    // Wait for confirmation
    receipt, err := bind.WaitMined(ctx, client.Backend, tx)
    if err != nil {
        return err
    }

    if receipt.Status == 1 {
        log.Printf("Registration confirmed! You are now a validator.")
    }

    return nil
}
```

### 6. List Active Subnets

```go
func listActiveSubnets(client *sdk.Client) error {
    ctx := context.Background()

    // Get all active subnets (efficient batch query)
    subnets, err := client.SubnetFactory.GetSubnetsByStatus(ctx, sdk.SubnetStatusActive)
    if err != nil {
        return fmt.Errorf("failed to get active subnets: %w", err)
    }

    log.Printf("Found %d active subnet(s):", len(subnets))

    for i, subnet := range subnets {
        log.Printf("\nSubnet %d:", i+1)
        log.Printf("  ID: %x", subnet.ID)
        log.Printf("  Address: %s", subnet.Addr.Hex())
        log.Printf("  Owner: %s", subnet.Owner.Hex())
        log.Printf("  Name: %s", subnet.Name)
        log.Printf("  Status: %d", subnet.Status)

        // Get participant counts
        subnetSvc, err := client.SubnetServiceByAddress(subnet.Addr)
        if err != nil {
            continue
        }

        validatorCount, _ := subnetSvc.GetParticipantCount(ctx, sdk.ParticipantValidator)
        agentCount, _ := subnetSvc.GetParticipantCount(ctx, sdk.ParticipantAgent)
        matcherCount, _ := subnetSvc.GetParticipantCount(ctx, sdk.ParticipantMatcher)

        log.Printf("  Participants:")
        log.Printf("    Validators: %s", validatorCount.String())
        log.Printf("    Agents: %s", agentCount.String())
        log.Printf("    Matchers: %s", matcherCount.String())
    }

    return nil
}
```

### 7. Assign Intent (Matcher)

```go
func assignIntent(client *sdk.Client, intentID, bidID [32]byte, agentAddr common.Address) error {
    ctx := context.Background()

    // Create assignment
    assignmentID := sdk.HashBytes(append(intentID[:], bidID[:]...)) // Example ID generation

    assignment := sdk.AssignmentData{
        AssignmentID: assignmentID,
        IntentID:     intentID,
        BidID:        bidID,
        Agent:        agentAddr,
        Status:       sdk.AssignmentStatusActive,
        Matcher:      client.Signer.Address(),
    }

    // Compute digest and sign
    digest, err := client.Assignment.ComputeDigest(assignment)
    if err != nil {
        return fmt.Errorf("failed to compute digest: %w", err)
    }

    sig, err := client.Assignment.SignDigest(digest)
    if err != nil {
        return fmt.Errorf("failed to sign: %w", err)
    }

    // Submit assignment
    tx, err := client.Assignment.AssignIntentsBySignatures(ctx, []sdk.SignedAssignment{
        {Data: assignment, Signature: sig},
    })
    if err != nil {
        return fmt.Errorf("failed to assign intent: %w", err)
    }

    log.Printf("Intent assigned: %s", tx.Hash().Hex())
    return nil
}
```

### 8. Validate Intent (Validator)

```go
func validateIntent(client *sdk.Client, intentID, assignmentID, subnetID [32]byte, agentAddr common.Address) error {
    ctx := context.Background()

    // Create validation bundle
    bundle := sdk.ValidationBundle{
        IntentID:     intentID,
        AssignmentID: assignmentID,
        SubnetID:     subnetID,
        Agent:        agentAddr,
        ResultHash:   sdk.HashBytes([]byte("execution_result_data")),
        ProofHash:    sdk.HashBytes([]byte("execution_proof")),
        RootHeight:   big.NewInt(12345),
        RootHash:     sdk.HashBytes([]byte("state_root")),
        Validators:   []common.Address{client.Signer.Address()},
    }

    // Compute validation digest and sign
    digest, err := client.Validation.ComputeDigest(bundle)
    if err != nil {
        return fmt.Errorf("failed to compute digest: %w", err)
    }

    sig, err := client.Validation.SignDigest(digest)
    if err != nil {
        return fmt.Errorf("failed to sign: %w", err)
    }

    bundle.Signatures = [][]byte{sig}

    // Submit validation
    tx, err := client.Validation.ValidateIntentsBySignatures(ctx, []sdk.ValidationBundle{bundle})
    if err != nil {
        return fmt.Errorf("failed to validate intent: %w", err)
    }

    log.Printf("Intent validated: %s", tx.Hash().Hex())
    return nil
}
```

## Advanced Configuration

### Custom Transaction Options

```go
import "time"

client, err := sdk.NewClient(ctx, sdk.Config{
    RPCURL:        os.Getenv("PIN_RPC_URL"),
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        Use1559:            ptr(true),
        GasLimitMultiplier: ptr(1.2),           // 20% safety margin
        GasCeil:            ptr(uint64(2000000)), // Max 2M gas
        ReplaceStuck:       ptr(true),
        ReplaceAfter:       ptr(30 * time.Second),
        BumpPercent:        ptr(12.5),
    },
})
```

### Dry-Run Mode (No Transactions)

```go
client, err := sdk.NewClient(ctx, sdk.Config{
    RPCURL:        os.Getenv("PIN_RPC_URL"),
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        NoSend: ptr(true), // Only construct transactions, don't send
    },
})

// Transaction will be constructed but not broadcasted
tx, err := client.Intent.SubmitIntent(ctx, params)
// tx contains all details for offline signing/auditing
```

## Helper Functions

```go
// Pointer helpers for config
func ptr[T any](v T) *T {
    return &v
}
```

## Next Steps

- **Role-Specific Guides**: See [Roles and Workflows](roles-and-workflows.md) for detailed integration patterns
- **Production Setup**: Review [Best Practices](best-practices.md) before deploying
- **API Details**: Complete method reference in [API Reference](api-reference.md)
- **Troubleshooting**: Common issues and solutions in [Troubleshooting](troubleshooting.md)

## Example Scripts

Check the `examples/` directory for complete, runnable scripts:

- `examples/send_intent`: Intent submission with environment variables
- `examples/list_subnets`: Query and display subnet information
- `examples/register_agent`: Register as an agent participant
- `examples/assign_intents`: Matcher assignment workflow
- `examples/validate_intents`: Validator validation workflow
