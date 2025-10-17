# Configuration & Environment

**[中文文档](config-zh.md)** | English

This document explains SDK configuration methods, environment variable conventions, and local development setup.

## Basic Configuration

### Required Configuration

- `PIN_RPC_URL` (**required**): Ethereum node RPC endpoint

### Optional Configuration

- `PIN_PRIVATE_KEY` (optional): Private key with `0x` prefix (skip if using remote signer)
- `PIN_NETWORK` (optional): `base` | `base_sepolia` | `local` (auto-detected from RPC chainId if not set)

Contract addresses can be provided via environment variables or code-level configuration. See [Contract Addresses & Networks](addresses.md).

## Configuration Priority

The SDK resolves configuration in the following order (highest to lowest priority):

1. **Code-level configuration**: Explicit parameters passed to `sdk.Config`
2. **Environment variables**: Network-specific environment variables
3. **SDK defaults**: Built-in mock/local addresses (fallback)

## Environment Variables

### .env File Template

See `.env.example` in the repository root. Copy it to `.env` and fill in your values:

```ini
# ============================================
# Network Configuration
# ============================================
PIN_NETWORK=base_sepolia
PIN_RPC_URL=https://sepolia.base.org
PIN_PRIVATE_KEY=0xYourPrivateKeyHere

# ============================================
# Base Sepolia Contract Addresses
# (Uses SDK default mock addresses if not provided)
# ============================================
PIN_BASE_SEPOLIA_INTENT_MANAGER=0x5555555555555555555555555555555555555555
PIN_BASE_SEPOLIA_SUBNET_FACTORY=0x6666666666666666666666666666666666666666
PIN_BASE_SEPOLIA_STAKING_MANAGER=0x7777777777777777777777777777777777777777
PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER=0x8888888888888888888888888888888888888888

# ============================================
# Optional: Local Development Configuration
# (Takes effect when PIN_NETWORK=local or RPC chainId=31337)
# ============================================
# PIN_LOCAL_INTENT_MANAGER=0x9999999999999999999999999999999999999999
# PIN_LOCAL_SUBNET_FACTORY=0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
# PIN_LOCAL_STAKING_MANAGER=0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
# PIN_LOCAL_CHECKPOINT_MANAGER=0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC
```

### Loading Environment Variables

```go
import (
    "os"
    "github.com/joho/godotenv"
    sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
)

func main() {
    // Load .env file (optional, can use system environment variables directly)
    _ = godotenv.Load()

    client, err := sdk.NewClient(context.Background(), sdk.Config{
        RPCURL:        os.Getenv("PIN_RPC_URL"),
        PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
        Network:       os.Getenv("PIN_NETWORK"), // Optional
    })
    // ...
}
```

## Local Development Setup

### Using Hardhat/Anvil

Recommended setup for local development:

```bash
# Start Anvil (Foundry)
anvil --chain-id 31337

# Or Hardhat
npx hardhat node
```

### Local Configuration

```ini
PIN_NETWORK=local
PIN_RPC_URL=http://127.0.0.1:8545
PIN_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Optional: Override with your deployed contract addresses
PIN_LOCAL_INTENT_MANAGER=0xYourLocalIntentManager
PIN_LOCAL_SUBNET_FACTORY=0xYourLocalSubnetFactory
PIN_LOCAL_STAKING_MANAGER=0xYourLocalStakingManager
PIN_LOCAL_CHECKPOINT_MANAGER=0xYourLocalCheckpointManager
```

> **Tip**: When using local test networks, the SDK will use built-in default addresses if no overrides are provided. These are suitable for development but should be replaced with actual deployment addresses for testing contract integration.

## Advanced Configuration

### Code-Level Configuration

For maximum flexibility, pass configuration directly to the SDK:

```go
import (
    "math/big"
    "time"

    sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
    "github.com/ethereum/go-ethereum/common"
)

cfg := sdk.Config{
    // Basic connectivity
    RPCURL:        "https://sepolia.base.org",
    PrivateKeyHex: "0x...",
    Network:       "base_sepolia",

    // Custom contract addresses (highest priority)
    Addresses: &sdk.AddressBook{
        Network:           "base_sepolia",
        IntentManager:     common.HexToAddress("0x..."),
        SubnetFactory:     common.HexToAddress("0x..."),
        StakingManager:    common.HexToAddress("0x..."),
        CheckpointManager: common.HexToAddress("0x..."),
    },

    // Transaction management options
    Tx: &sdk.TxOptions{
        Use1559:            ptr(true),
        GasLimitMultiplier: ptr(1.2),               // 20% safety margin
        GasCeil:            ptr(uint64(2_000_000)), // Max 2M gas
        ReplaceStuck:       ptr(true),
        ReplaceAfter:       ptr(30 * time.Second),
        BumpPercent:        ptr(12.5),
        NoSend:             ptr(false),
    },
}

client, err := sdk.NewClient(context.Background(), cfg)

// Helper function
func ptr[T any](v T) *T { return &v }
```

### Transaction Options

Complete configuration options for `TxOptions`:

```go
type TxOptions struct {
    // EIP-1559 Configuration
    Use1559              *bool     // Enable EIP-1559 (default: true)
    MaxFeePerGas         *big.Int  // Max fee per gas (auto-calculated if nil)
    MaxPriorityFeePerGas *big.Int  // Max priority fee (auto-calculated if nil)

    // Gas Estimation
    GasLimit             *uint64   // Fixed gas limit (0 = auto-estimate)
    GasLimitMultiplier   *float64  // Safety multiplier for estimates (default: 1.5)
    GasCeil              *uint64   // Maximum allowed gas (0 = no limit)

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

See [Transaction Management](txmanager.md) for detailed TxManager documentation.

### Custom Signer

For advanced use cases (hardware wallets, remote signing, etc.):

```go
import "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/signer"

// Implement the Signer interface
type CustomSigner struct {
    addr common.Address
    // ... your signing logic
}

func (s *CustomSigner) Address() common.Address {
    return s.addr
}

func (s *CustomSigner) SignDigest(digest [32]byte) ([]byte, error) {
    // Your custom signing implementation
    // Must return 65-byte signature compatible with ECDSA recovery
    return signature, nil
}

// Use custom signer with SDK
cfg := sdk.Config{
    RPCURL: "https://sepolia.base.org",
    Signer: &CustomSigner{addr: common.HexToAddress("0x...")},
}
client, err := sdk.NewClient(context.Background(), cfg)
```

## Network-Specific Configurations

### Base Mainnet Production

```ini
PIN_NETWORK=base
PIN_RPC_URL=https://mainnet.base.org
PIN_PRIVATE_KEY=0x...

PIN_BASE_INTENT_MANAGER=0xProductionIntentManager
PIN_BASE_SUBNET_FACTORY=0xProductionSubnetFactory
PIN_BASE_STAKING_MANAGER=0xProductionStakingManager
PIN_BASE_CHECKPOINT_MANAGER=0xProductionCheckpointManager
```

### Base Sepolia Testing

```ini
PIN_NETWORK=base_sepolia
PIN_RPC_URL=https://sepolia.base.org
PIN_PRIVATE_KEY=0x...

PIN_BASE_SEPOLIA_INTENT_MANAGER=0xTestIntentManager
PIN_BASE_SEPOLIA_SUBNET_FACTORY=0xTestSubnetFactory
PIN_BASE_SEPOLIA_STAKING_MANAGER=0xTestStakingManager
PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER=0xTestCheckpointManager
```

### Local Development

```ini
PIN_NETWORK=local
PIN_RPC_URL=http://127.0.0.1:8545
PIN_PRIVATE_KEY=0xLocalDevPrivateKey

# Optional: Use defaults or specify deployed addresses
# PIN_LOCAL_INTENT_MANAGER=0x...
```

## Configuration Validation

The SDK performs validation on initialization:

```go
client, err := sdk.NewClient(ctx, cfg)
if err != nil {
    // Common validation errors:
    // - Invalid RPC URL or connection failure
    // - Invalid private key format
    // - Network/chainId mismatch
    // - Missing required contract addresses
    log.Fatalf("SDK initialization failed: %v", err)
}

// Verify configuration
log.Printf("Connected to: %s (chainId: %s)", client.Network, client.ChainID)
log.Printf("Using address: %s", client.Signer.Address())
log.Printf("IntentManager: %s", client.Addresses.IntentManager)
```

## Best Practices

### Development
- Use `.env` file for local configuration
- Leverage default addresses for rapid prototyping
- Enable `NoSend` for transaction testing

### Staging/Testing
- Use Base Sepolia testnet
- Set explicit contract addresses via environment variables
- Enable stuck transaction replacement for reliability

### Production
- **Never commit** private keys or `.env` files to version control
- Use environment variables from secure secret management
- Set explicit addresses via code-level or environment configuration
- Configure appropriate gas limits and fee strategies
- Enable monitoring and logging

## Troubleshooting

### Issue: "Invalid private key"
- Ensure key starts with `0x`
- Verify key is 64 hex characters (after `0x`)
- Check for typos or extra whitespace

### Issue: "Connection refused" or "Context deadline exceeded"
- Verify RPC URL is correct and accessible
- Check network connectivity
- Try alternative RPC providers

### Issue: "Network mismatch"
- Ensure `PIN_NETWORK` matches the chainId of your RPC endpoint
- Or remove `PIN_NETWORK` to enable auto-detection

### Issue: "Contract call reverted"
- Verify contract addresses are correct for your network
- Ensure contracts are deployed at specified addresses
- Check if contracts are paused or access-controlled

## Related Documentation

- [Contract Addresses & Networks](addresses.md) - Address configuration details
- [Transaction Management](txmanager.md) - TxManager options and strategies
- [Quick Start](quickstart.md) - Getting started examples
- [Troubleshooting](troubleshooting.md) - Common issues and solutions
