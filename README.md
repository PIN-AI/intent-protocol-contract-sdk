# intent-protocol-contract-sdk

Go SDK for PIN AI Intent Protocol RootLayer smart contracts.

Provides comprehensive on-chain interaction wrappers for RootLayer and Subnet, with pre-configured contract addresses for Base mainnet/testnet and local development networks (configurable via environment variables or code). Features EIP-1559 transaction management, EIP-191 digest signing utilities, and exposes high-level APIs for common operations (submit/batch submit Intents, query subnets/validators/checkpoints, staking operations, etc.).

• Quick Navigation: `docs/README.md`
• Examples: `examples/`
• Environment Variables: `.env.example`

---

## Features

- **Strongly-Typed Contract Bindings**: `IntentManager`, `SubnetFactory`, `Subnet`, `StakingManager`, `CheckpointManager`
- **Complete Service Layer** (2025-01 refactor with 138 new methods):
  - **IntentService** (22 methods): Submit/batch signed submit, expiry handling, status queries, role management, emergency controls
  - **AssignmentService**: Batch signed assignment, digest construction, matcher signature helpers
  - **ValidationService**: Batch validation submission, validator signature digest
  - **SubnetFactoryService** (30 methods): Subnet creation, query by status/owner, batch participant retrieval, pause/resume/deprecate
  - **SubnetService** (27 methods): Register validator/agent/matcher (ETH or ERC20 staking), participant management, config queries
  - **StakingService** (21 methods): Stake/unstake/withdraw, staking info queries, slashing, role and config management
  - **CheckpointService** (18 methods): Query, verify, submit, finalize, revert checkpoints
- **Configurable TxManager**: EIP-1559 fees, nonce source, stuck transaction replacement (gas bump), dry-run
- **Signing & Hashing**: EIP-191 (eth_sign) digest signing, batch submission digest construction, EIP-712 reserved
- **Networks & Addresses**: Pre-configured for `base`/`base_sepolia`/`local`, supports environment variables and code-level overrides

## Installation & Requirements

- **Go Version**: `go 1.24+` (see `go.mod`)
- **Get SDK**:

```bash
go get github.com/PIN-AI/intent-protocol-contract-sdk@latest
```

- **Build & Test**:

```bash
# Build all packages
go build ./...

# Run all tests
go test ./...

# Run specific package tests
go test -v ./sdk/crypto/
```

## Quick Start

1) Prepare environment variables (refer to `.env.example`)

```ini
PIN_NETWORK=base_sepolia
PIN_RPC_URL=https://sepolia.base.org
PIN_PRIVATE_KEY=0xYourPrivateKey

# Optional: Override contract addresses per network (uses mock addresses if not set)
PIN_BASE_SEPOLIA_INTENT_MANAGER=0x5555...
PIN_BASE_SEPOLIA_SUBNET_FACTORY=0x6666...
PIN_BASE_SEPOLIA_STAKING_MANAGER=0x7777...
PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER=0x8888...
```

2) Initialize Client and submit an Intent (ETH payment example)

```go
package main

import (
  "context"
  "log"
  "math/big"
  "os"
  "time"

  sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
  "github.com/ethereum/go-ethereum/common"
)

func main() {
  ctx := context.Background()
  client, err := sdk.NewClient(ctx, sdk.Config{
    RPCURL:        os.Getenv("PIN_RPC_URL"),
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Network:       os.Getenv("PIN_NETWORK"), // Optional: auto-detect from chainId
  })
  if err != nil { log.Fatal(err) }
  defer client.Close()

  params := sdk.SubmitIntentParams{
    IntentID:   sdk.MustBytes32FromHex("0x..."),
    SubnetID:   sdk.MustBytes32FromHex("0x..."),
    IntentType: "book_flight",
    ParamsHash: sdk.HashBytes([]byte("{...}")),
    Deadline:   big.NewInt(time.Now().Add(24*time.Hour).Unix()),
    // 0 address = ETH; for ERC20, set token address and leave Value empty
    PaymentToken: common.Address{},
    Amount:       big.NewInt(1e16), // 0.01 ETH
    Value:        big.NewInt(1e16), // Must match Amount for ETH
  }
  tx, err := client.Intent.SubmitIntent(ctx, params)
  if err != nil { log.Fatal(err) }
  log.Printf("submitted: %s", tx.Hash())
}
```

3) Batch submission with signatures (digest construction + EIP-191 signing)

```go
input := sdkcrypto.SignedIntentInput{
  IntentID:     sdk.MustBytes32FromHex("0x..."),
  SubnetID:     sdk.MustBytes32FromHex("0x..."),
  Requester:    client.Signer.Address(),
  IntentType:   "book_flight",
  ParamsHash:   sdk.MustBytes32FromHex("0x..."),
  Deadline:     big.NewInt(....),
  PaymentToken: common.Address{},
  Amount:       big.NewInt(1e16),
}

digest, _ := client.Intent.ComputeDigest(input)
sig, _ := client.Intent.SignDigest(digest)

tx, err := client.Intent.SubmitIntentsBySignatures(ctx, sdk.SubmitIntentBatchParams{
  Items: []sdk.SignedIntent{{Data: input, Signature: sig}},
  // TotalEthValue optional: SDK auto-sums PaymentToken==0 amounts
})
```

4) Assignment & Validation (Matcher & Validator)

```go
assignment := sdk.AssignmentData{
  AssignmentID: sdk.MustBytes32FromHex("0x..."),
  IntentID:     sdk.MustBytes32FromHex("0x..."),
  BidID:        sdk.MustBytes32FromHex("0x..."),
  Agent:        common.HexToAddress("0xAgent"),
  Status:       sdk.AssignmentStatusActive,
  Matcher:      client.Signer.Address(),
}

digest, _ = client.Assignment.ComputeDigest(assignment)
sig, _ = client.Assignment.SignDigest(digest)

_, _ = client.Assignment.AssignIntentsBySignatures(ctx, []sdk.SignedAssignment{{Data: assignment, Signature: sig}})

bundle := sdk.ValidationBundle{
  IntentID:     assignment.IntentID,
  AssignmentID: assignment.AssignmentID,
  SubnetID:     sdk.MustBytes32FromHex("0x..."),
  Agent:        assignment.Agent,
  ResultHash:   sdk.HashBytes([]byte("result")),
  ProofHash:    sdk.HashBytes([]byte("proof")),
  RootHeight:   123,
  RootHash:     sdk.HashBytes([]byte("root")),
  Validators:   []common.Address{common.HexToAddress("0xValidator1")},
  Signatures:   [][]byte{sig}, // Real validator signatures required
}

_, _ = client.Validation.ValidateIntentsBySignatures(ctx, []sdk.ValidationBundle{bundle})
```

More examples in `docs/quickstart.md`.

### Example Scripts

- `examples/list_subnets`: **Optimized** listing of active subnets, uses `GetSubnetsByStatus` for efficient queries, displays detailed participant info (address, reputation, endpoint, etc.), avoids redundant RPC calls.
- `examples/send_intent`: Environment-driven CLI, default dry-run, demonstrates single submission and batch submission with signatures (requires `PIN_RPC_URL`/`PIN_PRIVATE_KEY`, optional `SUBNET_ID`, `SIGNED_INTENT_ID`, etc.).
- `examples/assign_intents`: Matcher-side batch assignment script, supports auto-signing single assignment or using external signatures (requires `INTENT_ID`/`AGENT_ADDRESS`, etc.).
- `examples/validate_intents`: Validator-side validation script, can compute digest and sign for single validator or load pre-signed multi-validator bundles.
- `examples/register_agent`: Agent registration script based on Subnet contract, configurable `AGENT_DOMAIN`/`AGENT_ENDPOINT`/`AGENT_METADATA_URI` and `AGENT_VALUE_WEI` (default dry-run, can override 0 stake requirement).
- `examples/complete_workflow` (planned): Demonstrates end-to-end flow from Intent → Assignment → Validation → Checkpoint.

## Connecting to Subnet & Role Registration

### Connecting to Subnet Contract (two methods)

1) Query address by subnet ID and get service

```go
subnetID := sdk.MustBytes32FromHex("0x...")
subnetSvc, err := client.SubnetServiceByID(ctx, subnetID)
if err != nil { log.Fatal(err) }
```

2) Bind directly if subnet contract address is known

```go
subnetAddr := common.HexToAddress("0xSubnET...")
subnetSvc, err := client.SubnetServiceByAddress(subnetAddr)
if err != nil { log.Fatal(err) }
```

Both methods internally share the `Client`'s `TxManager` and `Signer`, maintaining consistent nonce and 1559 fee strategy.

### Registering 3 Participant Roles

The Subnet contract's `RegisterParticipant`/`RegisterParticipantERC20` are wrapped as 6 convenience methods:

- `RegisterValidator` / `RegisterValidatorERC20`
- `RegisterAgent` / `RegisterAgentERC20`
- `RegisterMatcher` / `RegisterMatcherERC20`

ETH staking registration (sends native token to contract, can satisfy minimum stake):

```go
tx, err := subnetSvc.RegisterValidator(ctx, sdk.RegisterParticipantParams{
  Domain:      "example.org",
  Endpoint:    "https://validator.example.org",
  MetadataURI: "ipfs://...",
  Value:       big.NewInt(1e18), // e.g. 1 ETH, check registration fee
})
if err != nil { log.Fatal(err) }
log.Printf("validator registered, tx=%s", tx.Hash())

// Same for Agent and Matcher:
_, _ = subnetSvc.RegisterAgent(ctx, sdk.RegisterParticipantParams{Domain: "...", Endpoint: "...", MetadataURI: "...", Value: big.NewInt(0)})
_, _ = subnetSvc.RegisterMatcher(ctx, sdk.RegisterParticipantParams{Domain: "...", Endpoint: "...", MetadataURI: "...", Value: big.NewInt(0)})
```

ERC20 staking registration (uses governance-configured staking token; ensure allowance is set before calling):

```go
tx, err := subnetSvc.RegisterValidatorERC20(ctx, sdk.RegisterParticipantERC20Params{
  Amount:      big.NewInt(1_000_000_000_000_000_000), // 1 token (adjust for decimals)
  Domain:      "example.org",
  Endpoint:    "https://validator.example.org",
  MetadataURI: "ipfs://...",
})
if err != nil { log.Fatal(err) }
log.Printf("validator (ERC20) registered, tx=%s", tx.Hash())
```

Tips:
- Specific minimum stake and approval requirements are determined by subnet's `StakeGovernanceConfig` and `autoApprove` settings; manual approval by subnet governance may be required if not auto-approved.
- ERC20 path typically requires `approve` authorization to the staking receiver (e.g., `StakingManager` or `Subnet`, depending on contract implementation).
- Query methods available:
  - `subnetSvc.IsActiveParticipant(ctx, addr, sdk.ParticipantValidator)`
  - `subnetSvc.ListActiveParticipants(ctx, sdk.ParticipantValidator)`
  - `subnetSvc.GetParticipantInfo/GetParticipantStakeInfo/GetParticipantCount`

## Networks & Addresses

- Pre-configured networks: `base`(8453), `base_sepolia`(84532), `local`(31337)
- Override priority: Code-level override > Environment variables > SDK defaults (mock/local)
- Environment variable naming (examples):
  - `PIN_BASE_INTENT_MANAGER` / `PIN_BASE_SEPOLIA_SUBNET_FACTORY` / `PIN_LOCAL_CHECKPOINT_MANAGER`
- See: `docs/addresses.md` and `.env.example`

## TxManager (Transaction Management)

Configurable EIP-1559 transaction manager supporting:

- Nonce source: `pending`/`latest`
- Gas estimation with multiplier, suggested fees, `MaxFeePerGas`/`MaxPriorityFee`
- Replacement strategy: After `ReplaceAfter` timeout, bump by `BumpPercent` and resend
- Dry-run: `NoSend` constructs transaction without sending

Configure in `sdk.Config.Tx` or use defaults. See `docs/txmanager.md` for details.

## Signing & Digest

- Batch signing digest:
  - typeHash: `PIN_INTENT_V1(bytes32,bytes32,address,bytes32,bytes32,uint256,address,uint256,address,uint256)`
  - digest: `keccak256(abi.encode(typeHash, intent_id, subnet_id, requester, keccak256(bytes(intent_type)), params_hash, deadline, payment_token, amount, address(this), chainid))`
- Off-chain signing: EIP-191 (eth_sign prefix), aligned with contract `SignatureLib.verifySingleSignature()`
- Utility functions: `sdkcrypto.ComputeIntentDigest()`, `client.Intent.SignDigest()`
- Other digests: `client.Assignment.ComputeDigest()`, `client.Validation.ComputeDigest()`, `client.CheckpointManager.ComputeDigest()`
- See: `docs/signing.md`

## Directory Structure

- `sdk/`: Public API (Client, TxManager, Signer, AddressBook, high-level Services)
  - 7 complete Services: Intent, Assignment, Validation, SubnetFactory, Subnet, Staking, Checkpoint
  - All Services implement full contract ABI method coverage (read + write)
- `contracts/`: abigen bindings organized by contract module (pre-generated)
- `examples/`: Runnable example scripts
- `docs/`: Documentation and specifications
- `CLAUDE.md`: Claude Code project guide (architecture, commands, conventions)

## Generating Contract Bindings (Developers)

Bindings are provided with the repository. To regenerate from `RootLayer/artifacts`, refer to:

```bash
# Requires abigen and jq
for name in IntentManager SubnetFactory Subnet StakingManager CheckpointManager; do
  jq -r '.abi' /path/to/RootLayer/artifacts/contracts/${name}.sol/${name}.json > /tmp/${name}.abi
  pkg=$(echo "$name" | tr '[:upper:]' '[:lower:]')
  abigen --abi /tmp/${name}.abi --pkg $pkg --type ${name} --out contracts/$pkg/${pkg}.go
done
```

## Service Layer Method Categories

All Services provide complete method coverage, categorized as:

### Read-Only Methods
- **Roles & Permissions**: `DefaultAdminRole()`, `GovernanceRole()`, `HasRole()`, `GetRoleAdmin()`
- **Config Queries**: `GetMinStakeCreateSubnet()`, `GetStakingManager()`, `GetMaxIntentDuration()`
- **State Queries**: `IsSubnetActive()`, `IntentExists()`, `Paused()`, `CanFinalizeCheckpoint()`
- **Stats Queries**: `GetTotalSubnetCount()`, `GetActiveSubnetCount()`, `GetParticipantCount()`
- **Batch Queries**: `GetSubnetsByStatus()`, `GetSubnetsByOwner()`, `GetAllSubnetInfo()`

### Write Methods
- **Role Management**: `GrantRole()`, `RevokeRole()`, `RenounceRole()`
- **Emergency Controls**: `EmergencyPause()`, `EmergencyUnpause()`, `EmergencyRefundBatch()`
- **Config Management**: `SetMinStakeCreateSubnet()`, `SetMaxIntentDuration()`, `SetStakingToken()`
- **Subnet Management**: `CreateSubnet()`, `PauseSubnet()`, `ResumeSubnet()`, `DeprecateSubnet()`
- **Participant Management**: `ApproveParticipant()`, `RejectParticipant()`, `SuspendParticipant()`
- **Staking & Slashing**: `Slash()`, `DepositStakeFor()`, `RequestUnstake()`, `Withdraw()`

See `CLAUDE.md` for complete descriptions.

## FAQ

- **Addresses not taking effect?**
  - Check if network-matching environment variables are set; or use `sdk.Config.Addresses` for code-level override
- **Transaction stuck?**
  - Enable replacement strategy (`ReplaceStuck=true`, set `ReplaceAfter` and `BumpPercent`)
- **EIP-191 v value?**
  - SDK local signer normalizes `v∈{0,1}` to `27/28`
- **Need more methods?**
  - All contract ABI methods are fully implemented, refer to `CLAUDE.md` method categories

## Compatibility & Security

- **Complete Method Coverage**: All admin/governance functions implemented, including role management, emergency controls, config management
- **Access Control**: Strictly follow `GUARDIAN_ROLE`/`GOVERNANCE_ROLE`/`ADMIN_ROLE` permission requirements
- **Signature Security**: Strictly aligned with contract signature/hash, digest binds `address(this)` and `chainId`, providing cross-chain/cross-contract replay protection
- **Test Coverage**: All digest calculations verified by 52 unit tests

## Reference Documentation

- Documentation Index: `docs/README.md`
- Addresses & Networks: `docs/addresses.md`
- Configuration & Environment: `docs/config.md`
- Signing Specification: `docs/signing.md`
- Transaction Management: `docs/txmanager.md`
- Quick Start: `docs/quickstart.md`
- API Reference: `docs/api-reference.md`
- Best Practices: `docs/best-practices.md`
- Troubleshooting: `docs/troubleshooting.md`
