# Transaction Management (TxManager)

**[中文文档](txmanager-zh.md)** | English

TxManager handles transaction nonce management, gas estimation, EIP-1559 fee configuration, timeout retry, and replacement strategies.

## Overview

The TxManager provides:

- **Safe EIP-1559 defaults** - Automatic fee calculation based on network conditions
- **Concurrent-safe nonce allocation** - Based on `pending` or `latest` nonce source
- **Automatic gas estimation** - With configurable safety margin (`GasLimitMultiplier`)
- **Configurable replacement strategy** - Bump gas and resend stuck transactions
- **Dry-run mode** - Construct transactions without broadcasting (`NoSend`)

## Gas Estimation & Safety Margin

TxManager provides automatic gas estimation with safety margins to avoid out-of-gas errors.

### How It Works

1. **When `GasLimit==0` and `GasLimitMultiplier > 1.0`**:
   - TxManager performs a trial run with `NoSend=true` to get estimated gas
   - Applies `GasLimitMultiplier` to the estimate (default 1.5 = 50% margin)
   - Checks if result exceeds `GasCeil` limit (if configured)
   - Sets the adjusted value as actual `GasLimit`
   - Broadcasts the transaction

2. **When `GasLimit > 0`**:
   - Uses the specified fixed value directly, no estimation

3. **Recommended Configuration**:
   - Production: `GasLimit=0`, `GasLimitMultiplier=1.2` (20% margin)
   - Batch operations: `GasLimitMultiplier=1.3` (30% margin)

### Configuration Options

```go
type TxOptions struct {
    // EIP-1559 Configuration
    Use1559              *bool     // Enable EIP-1559 (default: true)
    MaxFeePerGas         *big.Int  // Max fee per gas (auto-calculated if nil)
    MaxPriorityFeePerGas *big.Int  // Max priority fee (auto-calculated if nil)

    // Gas Estimation
    GasLimit             *uint64   // Fixed gas limit (0 = auto-estimate, default: 0)
    GasLimitMultiplier   *float64  // Safety multiplier for estimates (default: 1.5)
    GasCeil              *uint64   // Maximum allowed gas (0 = no limit, default: 0)

    // Nonce Management
    NonceSource          *string   // "pending" or "latest" (default: "pending")

    // Stuck Transaction Replacement
    ReplaceStuck         *bool          // Enable automatic replacement (default: false)
    ReplaceAfter         *time.Duration // Wait time before replacement (e.g., 30s)
    BumpPercent          *float64       // Gas price increase percentage (e.g., 12.5)

    // Development Options
    NoSend               *bool     // Construct but don't broadcast tx (default: false)
}
```

### Gas Estimation Example

```go
import "time"

cfg := sdk.Config{
    RPCURL:        "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        GasLimitMultiplier: ptr(1.2), // 20% safety margin
    },
}
client, _ := sdk.NewClient(ctx, cfg)

// ValidateIntentsBySignatures will automatically estimate and apply 1.2x margin
tx, err := client.Validation.ValidateIntentsBySignatures(ctx, bundles)
```

## GasCeil Protection

`GasCeil` prevents transactions with abnormally high gas estimates from being sent.

### Behavior

- **Trigger condition**: Estimated gas (after multiplier) exceeds `GasCeil`
- **Action**: Refuses to send transaction, returns detailed error with:
  - Raw estimate
  - Adjusted estimate (with multiplier)
  - Multiplier value
  - Ceiling limit
- **Use cases**:
  - Batch operations (prevent gas explosion from too many items)
  - Production environments (control per-transaction cost ceiling)
  - Avoid approaching block gas limit (Base L2 = 30M)

### Recommended GasCeil Values

- Single operations: `1,000,000` (1M)
- Batch operations (<10 items): `2,000,000` (2M)
- Batch operations (>10 items): `5,000,000` (5M)

### GasCeil Example

```go
cfg := sdk.Config{
    RPCURL:        "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        GasLimitMultiplier: ptr(1.2),              // 20% margin
        GasCeil:            ptr(uint64(2000000)),  // 2M ceiling
    },
}
client, _ := sdk.NewClient(ctx, cfg)

// Automatically estimates and applies 1.2x margin, rejects if > 2M
tx, err := client.Validation.ValidateIntentsBySignatures(ctx, bundles)
if err != nil {
    // Check for gas ceiling error
    var gasCeilErr *txmgr.ErrGasCeilExceeded
    if errors.As(err, &gasCeilErr) {
        log.Printf("Gas too high: raw=%d, adjusted=%d (%.2fx), ceil=%d",
            gasCeilErr.Raw, gasCeilErr.Adjusted, gasCeilErr.Multiplier, gasCeilErr.Ceil)
        // Consider splitting batch or adjusting parameters
    }
    return err
}
```

## EIP-1559 Fee Management

### Automatic Fee Calculation

When `MaxFeePerGas` and `MaxPriorityFeePerGas` are not specified, TxManager automatically calculates them based on network conditions:

```go
// Automatic fee calculation (default)
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        Use1559: ptr(true), // EIP-1559 enabled by default
    },
}
```

The TxManager:
1. Queries `eth_maxPriorityFeePerGas` for priority fee
2. Queries `eth_gasPrice` for base fee reference
3. Calculates `MaxFeePerGas = baseFee * 2 + priorityFee`

### Manual Fee Configuration

For precise control:

```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        Use1559:              ptr(true),
        MaxFeePerGas:         big.NewInt(100_000_000_000), // 100 gwei
        MaxPriorityFeePerGas: big.NewInt(2_000_000_000),   // 2 gwei
    },
}
```

### Legacy Transaction Mode

To use legacy (non-EIP-1559) transactions:

```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        Use1559: ptr(false), // Disable EIP-1559
    },
}
```

> **Note**: Base L2 supports EIP-1559. Using legacy mode is not recommended unless required for specific use cases.

## Nonce Management

### Nonce Source Options

- **`pending` (default)**: Includes pending transactions, suitable for rapid sequential submissions
- **`latest`**: Only confirmed transactions, more conservative

```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        NonceSource: ptr("pending"), // or "latest"
    },
}
```

### Concurrent Safety

TxManager uses internal locking to ensure nonce uniqueness across concurrent transactions:

```go
// Safe: Multiple goroutines can submit concurrently
go client.Intent.SubmitIntent(ctx, params1)
go client.Intent.SubmitIntent(ctx, params2)
go client.Intent.SubmitIntent(ctx, params3)
```

### Nonce Behavior

| Nonce Source | Behavior | Use Case |
|--------------|----------|----------|
| `pending` | Includes unconfirmed txs | High-throughput submission |
| `latest` | Only confirmed txs | Conservative, lower risk of nonce gaps |

## Stuck Transaction Replacement

When transactions remain unconfirmed for extended periods, TxManager can automatically replace them with higher gas prices.

### Configuration

```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        ReplaceStuck: ptr(true),                     // Enable replacement
        ReplaceAfter: ptr(45 * time.Second),         // Wait 45s before replacing
        BumpPercent:  ptr(12.5),                     // Increase gas by 12.5%
    },
}
```

### How It Works

1. Submit transaction with initial gas price
2. Wait for `ReplaceAfter` duration
3. If still pending:
   - Calculate new gas price: `newPrice = oldPrice * (1 + BumpPercent/100)`
   - Submit replacement transaction with same nonce
4. Repeat until confirmed or max retries reached

### Replacement Example

```go
cfg := sdk.Config{
    RPCURL:        "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        ReplaceStuck: ptr(true),
        ReplaceAfter: ptr(30 * time.Second),
        BumpPercent:  ptr(10.0), // 10% gas increase per replacement
    },
}

client, _ := sdk.NewClient(ctx, cfg)

// If transaction is stuck, it will be automatically replaced every 30s
// with 10% higher gas until confirmed
tx, err := client.Intent.SubmitIntent(ctx, params)
```

### Best Practices for Replacement

- **ReplaceAfter**: Set based on network congestion (15-60s typical)
- **BumpPercent**: 10-20% recommended (too high = expensive, too low = still stuck)
- **Monitor costs**: Replacements increase transaction costs
- **Production**: Enable for critical transactions, consider disabling for batch operations

## Dry-Run Mode (NoSend)

Construct and validate transactions without broadcasting to the network.

### Use Cases

- **Cost estimation**: Calculate gas without spending funds
- **Transaction inspection**: Review transaction data before sending
- **Offline signing**: Prepare transactions for later broadcast
- **Testing**: Validate transaction construction in CI/CD

### Configuration

```go
cfg := sdk.Config{
    RPCURL:        "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        NoSend: ptr(true), // Enable dry-run mode
    },
}

client, _ := sdk.NewClient(ctx, cfg)

// Transaction is constructed but not sent
tx, err := client.Intent.SubmitIntent(ctx, params)
if err != nil {
    log.Fatalf("Transaction construction failed: %v", err)
}

// Inspect transaction
log.Printf("Transaction data: %x", tx.Data())
log.Printf("Gas limit: %d", tx.Gas())
log.Printf("Gas price: %s", tx.GasPrice())
log.Printf("Value: %s", tx.Value())
```

### Transaction Object

The returned transaction object contains all fields populated:

```go
tx, err := client.Intent.SubmitIntent(ctx, params)

// Access transaction fields
nonce := tx.Nonce()
to := tx.To()
value := tx.Value()
gasLimit := tx.Gas()
gasPrice := tx.GasPrice() // Legacy
gasTipCap := tx.GasTipCap() // EIP-1559
gasFeeCap := tx.GasFeeCap() // EIP-1559
data := tx.Data()

// Sign offline (if needed)
signedTx, err := types.SignTx(tx, signer, privateKey)

// Broadcast later
err = client.Backend.SendTransaction(ctx, signedTx)
```

## Usage Examples

### Basic Configuration (Recommended)

```go
cfg := sdk.Config{
    RPCURL:        "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        Use1559:            ptr(true),
        GasLimitMultiplier: ptr(1.2),              // 20% safety margin
        GasCeil:            ptr(uint64(2000000)),  // 2M ceiling
    },
}
client, err := sdk.NewClient(ctx, cfg)

// Transactions will use automatic EIP-1559 fees with 20% gas margin
tx, err := client.Intent.SubmitIntent(ctx, params)
```

### Production Configuration (Full)

```go
cfg := sdk.Config{
    RPCURL:        "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        // EIP-1559
        Use1559: ptr(true),
        // Auto-calculated if nil

        // Gas estimation
        GasLimitMultiplier: ptr(1.2),              // 20% safety margin
        GasCeil:            ptr(uint64(2000000)),  // 2M ceiling for batch ops

        // Nonce
        NonceSource: ptr("pending"), // High throughput

        // Replacement
        ReplaceStuck: ptr(true),
        ReplaceAfter: ptr(30 * time.Second),
        BumpPercent:  ptr(12.5),

        // Not in dry-run
        NoSend: ptr(false),
    },
}

client, err := sdk.NewClient(ctx, cfg)

// All transactions inherit these settings
tx, err := client.Intent.SubmitIntent(ctx, params)
```

### Development/Testing Configuration

```go
cfg := sdk.Config{
    RPCURL:        "http://127.0.0.1:8545", // Local node
    PrivateKeyHex: "0x...",
    Tx: &sdk.TxOptions{
        Use1559:            ptr(true),
        GasLimitMultiplier: ptr(1.5),   // Higher margin for testing
        NoSend:             ptr(true),  // Dry-run mode
    },
}

client, err := sdk.NewClient(ctx, cfg)

// Test transaction construction without sending
tx, err := client.Intent.SubmitIntent(ctx, params)
log.Printf("Estimated gas: %d", tx.Gas())
```

### Batch Operations Configuration

```go
cfg := sdk.Config{
    RPCURL:        "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        GasLimitMultiplier: ptr(1.3),              // 30% margin for batch
        GasCeil:            ptr(uint64(5000000)),  // 5M ceiling (larger batch)
        ReplaceStuck:       ptr(false),            // Disable replacement for batch
    },
}

client, err := sdk.NewClient(ctx, cfg)

// Batch submission with higher gas margin
tx, err := client.Intent.SubmitIntentsBySignatures(ctx, batchParams)
```

## Transaction Lifecycle

```
┌─────────────────────────────────────────────────────────────┐
│ 1. Construct Transaction                                     │
│    • Encode contract call data                               │
│    • Get nonce (pending/latest)                              │
│    • Estimate gas (if GasLimit==0)                           │
│    • Apply GasLimitMultiplier                                │
│    • Check GasCeil                                           │
│    • Calculate fees (if not specified)                       │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│ 2. Sign Transaction                                          │
│    • Create transaction object                               │
│    • Sign with private key                                   │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│ 3. Broadcast (if NoSend==false)                              │
│    • Send to network via RPC                                 │
│    • Return transaction hash                                 │
└───────────────────────┬─────────────────────────────────────┘
                        │
                        ▼
┌─────────────────────────────────────────────────────────────┐
│ 4. Monitor & Replace (if ReplaceStuck==true)                 │
│    • Wait ReplaceAfter duration                              │
│    • Check if still pending                                  │
│    • If stuck: Bump gas by BumpPercent, resend               │
│    • Repeat until confirmed                                  │
└─────────────────────────────────────────────────────────────┘
```

## Performance Considerations

### Gas Estimation Overhead

- **Auto-estimation**: Adds one extra RPC call (`eth_estimateGas`)
- **Impact**: ~50-100ms per transaction
- **Optimization**: Use fixed `GasLimit` for repeated similar operations

### Nonce Source Choice

| Source | RPC Calls | Latency | Concurrency Safety |
|--------|-----------|---------|-------------------|
| `pending` | 1 | Low | High (TxManager locks) |
| `latest` | 1 | Low | Medium (may have gaps) |

### Replacement Overhead

- Each replacement: Full transaction submission + gas bump
- Monitor replacement frequency to avoid excessive costs
- Consider disabling for non-critical batch operations

## Troubleshooting

### Issue: "Nonce too low"

**Cause**: Nonce already used by a confirmed transaction.

**Solutions**:
- Check for duplicate transaction submissions
- Verify `NonceSource` setting
- Use `latest` nonce source if `pending` causes issues
- Clear pending transactions before retrying

### Issue: "Transaction underpriced"

**Cause**: Gas price too low for current network conditions.

**Solutions**:
- Enable automatic fee calculation (don't set `MaxFeePerGas` manually)
- Increase `BumpPercent` if using replacement
- Check network congestion and adjust fees

### Issue: "Out of gas"

**Cause**: Gas limit too low to execute transaction.

**Solutions**:
- Increase `GasLimitMultiplier` (try 1.5 or 2.0)
- Check if `GasCeil` is too restrictive
- Verify contract call parameters are correct
- Test with `NoSend=true` to see estimated gas

### Issue: "Replacement transaction underpriced"

**Cause**: `BumpPercent` too low (must be > 10% on most networks).

**Solutions**:
- Increase `BumpPercent` to at least 12.5%
- Some networks require 10-20% minimum increase
- Check network-specific replacement rules

### Issue: "Transaction stuck forever"

**Cause**: Replacement not enabled or not working.

**Solutions**:
- Enable replacement: `ReplaceStuck=true`
- Adjust `ReplaceAfter` and `BumpPercent`
- Monitor gas prices on network
- Consider manual transaction replacement

## Helper Function

```go
// Pointer helper for configuration
func ptr[T any](v T) *T {
    return &v
}

// Usage
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        GasLimitMultiplier: ptr(1.2),
        GasCeil:            ptr(uint64(2000000)),
        ReplaceStuck:       ptr(true),
    },
}
```

## Implementation Details

The TxManager integrates with go-ethereum's:
- `SuggestGasTipCap` - Priority fee estimation
- `SuggestGasPrice` - Base fee estimation
- Built-in gas estimation - via contract call simulation
- EIP-1559 fee market - Compatible with Base L2 fee model

## Related Documentation

- [Configuration & Environment](config.md) - TxOptions configuration
- [Quick Start](quickstart.md) - Transaction examples
- [Best Practices](best-practices.md) - Production recommendations
- [Troubleshooting](troubleshooting.md) - Transaction issue debugging
