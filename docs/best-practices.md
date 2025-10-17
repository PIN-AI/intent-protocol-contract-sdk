# Best Practices

**[中文文档](best-practices-zh.md)** | English

Production-grade recommendations for deploying and operating with the PIN Intent Protocol SDK.

## Table of Contents

1. [Security](#security)
2. [Performance](#performance)
3. [Reliability](#reliability)
4. [Cost Optimization](#cost-optimization)
5. [Monitoring & Operations](#monitoring--operations)
6. [Development Workflow](#development-workflow)

## Security

### Private Key Management

**❌ Never**:
- Commit private keys to version control
- Store keys in plain text files
- Share keys across environments
- Log private keys

**✅ Do**:
```go
// Use environment variables from secure secret management
import "github.com/joho/godotenv"

// Load from secure vault in production
privateKey := os.Getenv("PIN_PRIVATE_KEY")

// Or use hardware wallet / remote signer
type HSMSigner struct {
    // Implement Signer interface with HSM
}
```

### Access Control

**Verify roles before operations**:
```go
// Check if address has required role before submission
hasRole, err := client.Intent.HasRole(ctx,
    client.Addresses.GuardianRole(),
    client.Signer.Address(),
)
if !hasRole {
    return fmt.Errorf("insufficient permissions")
}
```

### Input Validation

**Validate all user inputs**:
```go
func validateIntentParams(params sdk.SubmitIntentParams) error {
    // Check deadline is in future
    if params.Deadline.Cmp(big.NewInt(time.Now().Unix())) <= 0 {
        return fmt.Errorf("deadline must be in future")
    }

    // Check amount is positive
    if params.Amount.Sign() <= 0 {
        return fmt.Errorf("amount must be positive")
    }

    // Validate intent type format
    if !isValidIntentType(params.IntentType) {
        return fmt.Errorf("invalid intent type")
    }

    return nil
}
```

### Signature Verification

**Always verify signatures before accepting**:
```go
// When accepting signed intents from external sources
func verifySignature(input sdkcrypto.SignedIntentInput, sig []byte) error {
    digest, err := client.Intent.ComputeDigest(input)
    if err != nil {
        return err
    }

    // Verify signature matches expected signer
    recoveredAddr, err := crypto.RecoverAddress(digest[:], sig)
    if err != nil {
        return err
    }

    if recoveredAddr != input.Requester {
        return fmt.Errorf("signature mismatch")
    }

    return nil
}
```

## Performance

### Batch Operations

**Use batch methods for multiple items**:
```go
// ❌ Inefficient: Multiple individual transactions
for _, intent := range intents {
    tx, _ := client.Intent.SubmitIntent(ctx, intent)
}

// ✅ Efficient: Single batch transaction
signedIntents := make([]sdk.SignedIntent, len(intents))
for i, input := range intents {
    digest, _ := client.Intent.ComputeDigest(input)
    sig, _ := client.Intent.SignDigest(digest)
    signedIntents[i] = sdk.SignedIntent{Data: input, Signature: sig}
}

tx, err := client.Intent.SubmitIntentsBySignatures(ctx, sdk.SubmitIntentBatchParams{
    Items: signedIntents,
})
```

### Connection Reuse

**Reuse client connections**:
```go
// ❌ Don't create new client per request
func handleRequest() {
    client, _ := sdk.NewClient(ctx, cfg) // Expensive
    defer client.Close()
    // ... use client
}

// ✅ Reuse client across requests
var globalClient *sdk.Client

func init() {
    globalClient, _ = sdk.NewClient(context.Background(), cfg)
}

func handleRequest() {
    // Reuse globalClient
}
```

### Concurrent Operations

**Use goroutines for independent operations**:
```go
// Query multiple subnets concurrently
subnets := []common.Address{subnet1, subnet2, subnet3}
results := make(chan SubnetInfo, len(subnets))

for _, addr := range subnets {
    go func(addr common.Address) {
        svc, _ := client.SubnetServiceByAddress(addr)
        info, _ := svc.GetSubnetInfo(ctx)
        results <- info
    }(addr)
}

// Collect results
for i := 0; i < len(subnets); i++ {
    info := <-results
    // Process info
}
```

### Gas Optimization

**Optimize gas usage**:
```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        // Use optimal gas multiplier (not too high)
        GasLimitMultiplier: ptr(1.2), // 20% margin sufficient

        // Set ceiling to prevent waste
        GasCeil: ptr(uint64(2_000_000)),

        // Use pending nonce for high throughput
        NonceSource: ptr("pending"),
    },
}
```

## Reliability

### Error Handling

**Implement comprehensive error handling**:
```go
func submitWithRetry(ctx context.Context, params sdk.SubmitIntentParams) error {
    maxRetries := 3
    backoff := time.Second

    for attempt := 0; attempt < maxRetries; attempt++ {
        tx, err := client.Intent.SubmitIntent(ctx, params)

        if err == nil {
            // Success: wait for confirmation
            receipt, err := bind.WaitMined(ctx, client.Backend, tx)
            if err != nil {
                return fmt.Errorf("wait failed: %w", err)
            }

            if receipt.Status == 1 {
                return nil // Success
            }

            return fmt.Errorf("transaction reverted")
        }

        // Check if error is retryable
        if !isRetryable(err) {
            return fmt.Errorf("non-retryable error: %w", err)
        }

        // Exponential backoff
        time.Sleep(backoff)
        backoff *= 2
    }

    return fmt.Errorf("max retries exceeded")
}

func isRetryable(err error) bool {
    // Network errors are retryable
    if strings.Contains(err.Error(), "connection") {
        return true
    }

    // Nonce errors are retryable
    if strings.Contains(err.Error(), "nonce") {
        return true
    }

    return false
}
```

### Transaction Confirmation

**Always wait for confirmations**:
```go
import "github.com/ethereum/go-ethereum/accounts/abi/bind"

func submitAndConfirm(ctx context.Context, params sdk.SubmitIntentParams) error {
    // Submit transaction
    tx, err := client.Intent.SubmitIntent(ctx, params)
    if err != nil {
        return fmt.Errorf("submit failed: %w", err)
    }

    log.Printf("Transaction sent: %s", tx.Hash().Hex())

    // Wait for mining (1 confirmation)
    receipt, err := bind.WaitMined(ctx, client.Backend, tx)
    if err != nil {
        return fmt.Errorf("wait failed: %w", err)
    }

    // Check status
    if receipt.Status != 1 {
        return fmt.Errorf("transaction reverted")
    }

    log.Printf("Transaction confirmed in block: %d", receipt.BlockNumber.Uint64())

    // For critical operations: wait for more confirmations
    confirmations := 3
    for i := 1; i < confirmations; i++ {
        time.Sleep(12 * time.Second) // ~12s block time on Base

        currentBlock, _ := client.Backend.BlockNumber(ctx)
        if currentBlock >= receipt.BlockNumber.Uint64()+uint64(i) {
            log.Printf("Confirmation %d/%d", i+1, confirmations)
        }
    }

    return nil
}
```

### Stuck Transaction Handling

**Enable replacement for critical operations**:
```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        ReplaceStuck: ptr(true),
        ReplaceAfter: ptr(30 * time.Second),
        BumpPercent:  ptr(15.0), // 15% increase
    },
}

// Or manually replace if needed
func replaceTransaction(originalTx *types.Transaction) error {
    // Increase gas price
    newGasPrice := new(big.Int).Mul(
        originalTx.GasPrice(),
        big.NewInt(115), // 115% of original
    )
    newGasPrice.Div(newGasPrice, big.NewInt(100))

    // Create replacement with same nonce
    newTx := types.NewTransaction(
        originalTx.Nonce(),
        *originalTx.To(),
        originalTx.Value(),
        originalTx.Gas(),
        newGasPrice,
        originalTx.Data(),
    )

    // Sign and send
    // ...

    return nil
}
```

## Cost Optimization

### Minimize RPC Calls

**Cache frequently accessed data**:
```go
type CachedSubnetInfo struct {
    info      SubnetInfo
    timestamp time.Time
    ttl       time.Duration
}

var cache sync.Map // thread-safe map

func getSubnetInfoCached(ctx context.Context, addr common.Address) (SubnetInfo, error) {
    // Check cache
    if val, ok := cache.Load(addr); ok {
        cached := val.(CachedSubnetInfo)
        if time.Since(cached.timestamp) < cached.ttl {
            return cached.info, nil
        }
    }

    // Fetch from chain
    svc, _ := client.SubnetServiceByAddress(addr)
    info, err := svc.GetSubnetInfo(ctx)
    if err != nil {
        return SubnetInfo{}, err
    }

    // Update cache
    cache.Store(addr, CachedSubnetInfo{
        info:      info,
        timestamp: time.Now(),
        ttl:       5 * time.Minute,
    })

    return info, nil
}
```

### Use Read-Only Methods

**Prefer view functions over transactions**:
```go
// ❌ Expensive: Dry-run transaction to check
cfg.Tx.NoSend = ptr(true)
_, err := client.Intent.SubmitIntent(ctx, params)

// ✅ Cheap: Use read-only methods
exists, err := client.Intent.IntentExists(ctx, intentID)
isExpired, err := client.Intent.IsIntentExpired(ctx, intentID)
```

### Batch Queries

**Use batch query methods**:
```go
// ❌ Multiple individual queries
for _, id := range subnetIDs {
    subnet, _ := client.SubnetFactory.GetSubnet(ctx, id)
}

// ✅ Single batch query
subnets, err := client.SubnetFactory.GetSubnetsByStatus(ctx, sdk.SubnetStatusActive)
```

## Monitoring & Operations

### Logging

**Implement structured logging**:
```go
import "log/slog"

func submitIntent(params sdk.SubmitIntentParams) error {
    logger := slog.With(
        "intent_id", fmt.Sprintf("%x", params.IntentID),
        "subnet_id", fmt.Sprintf("%x", params.SubnetID),
    )

    logger.Info("submitting intent")

    tx, err := client.Intent.SubmitIntent(ctx, params)
    if err != nil {
        logger.Error("submission failed", "error", err)
        return err
    }

    logger.Info("intent submitted",
        "tx_hash", tx.Hash().Hex(),
        "gas_limit", tx.Gas(),
    )

    return nil
}
```

### Metrics

**Track key metrics**:
```go
import "github.com/prometheus/client_golang/prometheus"

var (
    intentsSubmitted = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "intents_submitted_total",
        Help: "Total number of intents submitted",
    })

    intentSubmissionDuration = prometheus.NewHistogram(prometheus.HistogramOpts{
        Name: "intent_submission_duration_seconds",
        Help: "Time taken to submit intent",
    })

    gasUsed = prometheus.NewHistogram(prometheus.HistogramOpts{
        Name: "transaction_gas_used",
        Help: "Gas used by transactions",
    })
)

func init() {
    prometheus.MustRegister(intentsSubmitted)
    prometheus.MustRegister(intentSubmissionDuration)
    prometheus.MustRegister(gasUsed)
}

func submitWithMetrics(params sdk.SubmitIntentParams) error {
    start := time.Now()
    defer func() {
        intentSubmissionDuration.Observe(time.Since(start).Seconds())
    }()

    tx, err := client.Intent.SubmitIntent(ctx, params)
    if err != nil {
        return err
    }

    intentsSubmitted.Inc()

    receipt, _ := bind.WaitMined(ctx, client.Backend, tx)
    gasUsed.Observe(float64(receipt.GasUsed))

    return nil
}
```

### Health Checks

**Implement health monitoring**:
```go
func healthCheck() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Check RPC connectivity
    _, err := client.Backend.BlockNumber(ctx)
    if err != nil {
        return fmt.Errorf("rpc unhealthy: %w", err)
    }

    // Check signer
    if client.Signer.Address() == (common.Address{}) {
        return fmt.Errorf("signer not initialized")
    }

    // Check contract accessibility
    _, err = client.Intent.Paused(ctx)
    if err != nil {
        return fmt.Errorf("contract inaccessible: %w", err)
    }

    return nil
}

// Run periodically
func startHealthMonitor() {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        if err := healthCheck(); err != nil {
            log.Printf("Health check failed: %v", err)
            // Alert or take corrective action
        }
    }
}
```

## Development Workflow

### Testing Strategy

**Use dry-run for tests**:
```go
func TestIntentSubmission(t *testing.T) {
    client, _ := sdk.NewClient(context.Background(), sdk.Config{
        RPCURL:        "http://127.0.0.1:8545",
        PrivateKeyHex: testPrivateKey,
        Tx: &sdk.TxOptions{
            NoSend: ptr(true), // Dry-run mode
        },
    })

    params := sdk.SubmitIntentParams{
        // ... test parameters
    }

    // Transaction constructed but not sent
    tx, err := client.Intent.SubmitIntent(context.Background(), params)
    require.NoError(t, err)
    require.NotNil(t, tx)

    // Verify transaction fields
    assert.Equal(t, uint64(0), tx.Nonce())
    assert.Greater(t, tx.Gas(), uint64(0))
}
```

### Environment Management

**Separate configurations per environment**:
```go
type Environment string

const (
    EnvDevelopment Environment = "development"
    EnvStaging     Environment = "staging"
    EnvProduction  Environment = "production"
)

func getConfig(env Environment) sdk.Config {
    switch env {
    case EnvDevelopment:
        return sdk.Config{
            RPCURL:        "http://127.0.0.1:8545",
            PrivateKeyHex: os.Getenv("DEV_PRIVATE_KEY"),
            Tx: &sdk.TxOptions{
                GasLimitMultiplier: ptr(2.0), // High margin
                NoSend:             ptr(false),
            },
        }

    case EnvStaging:
        return sdk.Config{
            RPCURL:        os.Getenv("STAGING_RPC_URL"),
            PrivateKeyHex: os.Getenv("STAGING_PRIVATE_KEY"),
            Tx: &sdk.TxOptions{
                GasLimitMultiplier: ptr(1.3),
                ReplaceStuck:       ptr(true),
                ReplaceAfter:       ptr(45 * time.Second),
            },
        }

    case EnvProduction:
        return sdk.Config{
            RPCURL:        os.Getenv("PROD_RPC_URL"),
            PrivateKeyHex: os.Getenv("PROD_PRIVATE_KEY"),
            Tx: &sdk.TxOptions{
                GasLimitMultiplier: ptr(1.2),
                GasCeil:            ptr(uint64(2_000_000)),
                ReplaceStuck:       ptr(true),
                ReplaceAfter:       ptr(30 * time.Second),
                BumpPercent:        ptr(12.5),
            },
        }

    default:
        panic("unknown environment")
    }
}
```

### CI/CD Integration

**Test contract integration in CI**:
```bash
#!/bin/bash
# ci-test.sh

set -e

# Start local node
anvil --chain-id 31337 --port 8545 &
ANVIL_PID=$!

# Wait for node to start
sleep 2

# Deploy contracts
forge script script/Deploy.s.sol --rpc-url http://127.0.0.1:8545 --broadcast

# Run tests
go test ./... -v

# Cleanup
kill $ANVIL_PID
```

## Production Checklist

Before deploying to production:

- [ ] **Security**
  - [ ] Private keys stored in secure vault
  - [ ] Access control roles verified
  - [ ] Input validation implemented
  - [ ] Signature verification in place

- [ ] **Configuration**
  - [ ] Correct network and RPC URL
  - [ ] Contract addresses verified
  - [ ] Gas parameters optimized
  - [ ] Transaction replacement enabled

- [ ] **Monitoring**
  - [ ] Structured logging implemented
  - [ ] Metrics collection configured
  - [ ] Health checks in place
  - [ ] Alerting rules defined

- [ ] **Testing**
  - [ ] Unit tests passing
  - [ ] Integration tests on testnet
  - [ ] Load testing completed
  - [ ] Failure scenarios tested

- [ ] **Operations**
  - [ ] Runbooks documented
  - [ ] Backup procedures defined
  - [ ] Rollback plan prepared
  - [ ] On-call rotation established

## Related Documentation

- [Configuration & Environment](config.md) - Configuration details
- [Transaction Management](txmanager.md) - TxManager best practices
- [Troubleshooting](troubleshooting.md) - Common issues and solutions
- [Quick Start](quickstart.md) - Getting started guide
