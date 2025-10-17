# Contract Addresses & Networks

**[中文文档](addresses-zh.md)** | English

The SDK selects contract addresses through a three-tier system: "default addresses + environment variable overrides + code-level overrides".

## Supported Networks

- **base** (mainnet, chainId=8453)
- **base_sepolia** (testnet, chainId=84532)
- **local** (local development, recommended for Hardhat/Anvil, chainId=31337)

## Override Priority

The SDK resolves contract addresses in the following order (highest to lowest priority):

1. **Code-level override**: Pass custom `AddressBook` when initializing SDK (highest priority)
2. **Environment variables**: Inject addresses via network-specific environment variables
3. **SDK default addresses**: Built-in mock/local addresses (fallback)

> **Important**: For production environments, always set real deployed addresses via environment variables or code-level overrides.

## Environment Variable Naming

### General Configuration

- `PIN_NETWORK`: Network identifier - `base` | `base_sepolia` | `local` (optional, auto-detected from RPC chainId if not set)
- `PIN_RPC_URL`: Ethereum RPC endpoint URL
- `PIN_PRIVATE_KEY`: Hex-encoded private key with `0x` prefix (or use custom Signer)

### Contract Addresses (by Network)

#### Base Mainnet
- `PIN_BASE_INTENT_MANAGER`
- `PIN_BASE_SUBNET_FACTORY`
- `PIN_BASE_STAKING_MANAGER`
- `PIN_BASE_CHECKPOINT_MANAGER`

#### Base Sepolia Testnet
- `PIN_BASE_SEPOLIA_INTENT_MANAGER`
- `PIN_BASE_SEPOLIA_SUBNET_FACTORY`
- `PIN_BASE_SEPOLIA_STAKING_MANAGER`
- `PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER`

#### Local Development
- `PIN_LOCAL_INTENT_MANAGER`
- `PIN_LOCAL_SUBNET_FACTORY`
- `PIN_LOCAL_STAKING_MANAGER`
- `PIN_LOCAL_CHECKPOINT_MANAGER`

## Default Addresses (Mock & Local)

When addresses are not provided via code or environment variables, the SDK falls back to these defaults:

### Base Mainnet (8453)
- IntentManager: `0x1111111111111111111111111111111111111111`
- SubnetFactory: `0x2222222222222222222222222222222222222222`
- StakingManager: `0x3333333333333333333333333333333333333333`
- CheckpointManager: `0x4444444444444444444444444444444444444444`

### Base Sepolia Testnet (84532)
- IntentManager: `0x5555555555555555555555555555555555555555`
- SubnetFactory: `0x6666666666666666666666666666666666666666`
- StakingManager: `0x7777777777777777777777777777777777777777`
- CheckpointManager: `0x8888888888888888888888888888888888888888`

### Local Development (31337)
- IntentManager: `0x9999999999999999999999999999999999999999`
- SubnetFactory: `0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
- StakingManager: `0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB`
- CheckpointManager: `0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC`

> **Tip**: The `local` network is intended for local development and integration testing. If you have actual deployed addresses locally, override the defaults using environment variables.

## Configuration Examples

### Environment Variables (`.env` file)

```ini
# Network configuration
PIN_NETWORK=base_sepolia
PIN_RPC_URL=https://sepolia.base.org
PIN_PRIVATE_KEY=0xYourPrivateKeyHere

# Base Sepolia contract addresses
PIN_BASE_SEPOLIA_INTENT_MANAGER=0x5555555555555555555555555555555555555555
PIN_BASE_SEPOLIA_SUBNET_FACTORY=0x6666666666666666666666666666666666666666
PIN_BASE_SEPOLIA_STAKING_MANAGER=0x7777777777777777777777777777777777777777
PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER=0x8888888888888888888888888888888888888888
```

### Code-Level Override

```go
import (
    sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
    "github.com/ethereum/go-ethereum/common"
)

cfg := sdk.Config{
    RPCURL:        "https://sepolia.base.org",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Network:       "base_sepolia",

    // Code-level override (highest priority)
    Addresses: &sdk.AddressBook{
        Network:           "base_sepolia",
        IntentManager:     common.HexToAddress("0xYourIntentManager"),
        SubnetFactory:     common.HexToAddress("0xYourSubnetFactory"),
        StakingManager:    common.HexToAddress("0xYourStakingManager"),
        CheckpointManager: common.HexToAddress("0xYourCheckpointManager"),
    },
}

client, err := sdk.NewClient(context.Background(), cfg)
```

## Network Auto-Detection

If `PIN_NETWORK` is not specified, the SDK will:

1. Connect to the RPC endpoint specified in `PIN_RPC_URL`
2. Query the `chainId` from the blockchain
3. Automatically map to the corresponding network:
   - `8453` → `base`
   - `84532` → `base_sepolia`
   - `31337` → `local`

This allows for flexible configuration where only the RPC URL needs to be changed to switch networks.

## FAQ

### Q: I configured `PIN_NETWORK` but addresses aren't taking effect?

**A**: Check the following:
1. Verify you've set the network-specific environment variables (e.g., `PIN_BASE_SEPOLIA_INTENT_MANAGER` for Base Sepolia)
2. Ensure the `PIN_NETWORK` value matches the actual chainId of your RPC endpoint
3. The SDK will also auto-detect network from RPC chainId, which may override `PIN_NETWORK` if they don't match

### Q: How does the SDK choose addresses when multiple networks are configured?

**A**: The SDK will:
1. First try to use code-level overrides if provided in `Config.Addresses`
2. Then read environment variables matching the current network (from `PIN_NETWORK` or RPC chainId)
3. Finally fall back to built-in defaults if no overrides are found

### Q: Can I use different contract addresses for the same network?

**A**: Yes, for testing or multi-deployment scenarios:
- Use code-level overrides with `Config.Addresses` for maximum flexibility
- Or use different environment variable sets and swap `.env` files

### Q: What happens if I don't set any addresses?

**A**: The SDK will use built-in default addresses. These are mock addresses suitable for:
- Local development testing
- Integration testing
- SDK API exploration

**Do not use default addresses in production.**

## Related Documentation

- [Configuration & Environment](config.md) - Complete configuration guide
- [Quick Start](quickstart.md) - Setup examples
- `.env.example` - Environment variable template
