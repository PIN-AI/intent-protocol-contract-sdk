# intent-protocol-contract-sdk

PIN AI Intent Protocol Rootlayer 智能合约 Go SDK。

面向 RootLayer 与 Subnet 的链上交互封装，预置 Base 主网/测试网与本地开发网络的合约地址（可通过环境变量与代码覆盖），提供 EIP‑1559 交易管理与 EIP‑191 摘要签名工具，并暴露常用的高层 API（提交/批量提交 Intent、查询子网/验证器/检查点、质押操作等）。

• 快速导航：`docs/README.md`
• 示例代码: `example/`
• 环境变量示例：`.env.example`

---

## 功能特性

- 强类型合约绑定：`IntentManager`、`SubnetFactory`、`Subnet`、`StakingManager`、`CheckpointManager`
- 高层 API：
  - Intent：提交/批量签名提交、过期处理、状态查询
  - Assignment：批量签名分配、digest 构造、matcher 签名辅助
  - Validation：批量验证提交、validator 签名 digest
  - SubnetFactory：活跃状态/合约地址查询、列表、预测地址、暂停/恢复/废弃（需权限）
  - Subnet：注册验证者/代理/匹配器（ETH 或 ERC20 质押）、查询参与者/质押信息、活跃状态
  - Staking：质押/解押/提现、质押信息查询（管理员函数不做“易用封装”，仅底层绑定）
  - Checkpoint：查询、digest 构造、checkpoint 提交
- 可配置 TxManager：EIP‑1559 费用、nonce 源、卡顿替换（gas bump）、dry‑run
- 签名与哈希：EIP‑191（eth_sign）摘要签名、批量提交 digest 构造，预留 EIP‑712
- 网络与地址：`base`/`base_sepolia`/`local` 预置，支持环境变量和代码级覆盖

## 安装与环境

- 要求 Go：`go 1.24+`（见 `go.mod`）
- 获取 SDK：

```bash
go get github.com/PIN-AI/intent-protocol-contract-sdk@latest
```

## 快速开始

1) 准备环境变量（可参考 `.env.example`）

```ini
PIN_NETWORK=base_sepolia
PIN_RPC_URL=https://sepolia.base.org
PIN_PRIVATE_KEY=0x你的私钥

# 可选：按网络覆盖合约地址（未设置时用默认 mock 地址）
PIN_BASE_SEPOLIA_INTENT_MANAGER=0x5555...
PIN_BASE_SEPOLIA_SUBNET_FACTORY=0x6666...
PIN_BASE_SEPOLIA_STAKING_MANAGER=0x7777...
PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER=0x8888...
```

2) 初始化 Client 并提交 Intent（ETH 预算示例）

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
    Network:       os.Getenv("PIN_NETWORK"), // 可空：自动根据 chainId 识别
  })
  if err != nil { log.Fatal(err) }
  defer client.Close()

  params := sdk.SubmitIntentParams{
    IntentID:   sdk.MustBytes32FromHex("0x..."),
    SubnetID:   sdk.MustBytes32FromHex("0x..."),
    IntentType: "book_flight",
    ParamsHash: sdk.HashBytes([]byte("{...}")),
    Deadline:   big.NewInt(time.Now().Add(24*time.Hour).Unix()),
    // 0 地址=ETH；ERC20 请填代币地址并将 Value 置空
    PaymentToken: common.Address{},
    Amount:       big.NewInt(1e16), // 0.01 ETH
    Value:        big.NewInt(1e16), // ETH 需与 Amount 一致
  }
  tx, err := client.Intent.SubmitIntent(ctx, params)
  if err != nil { log.Fatal(err) }
  log.Printf("submitted: %s", tx.Hash())
}
```

3) 批量签名提交（构造 digest + EIP‑191 签名）

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
  // ETH 总额可省略，SDK 会自动相加 PaymentToken==0 的金额
})
```

4) 分配与验证（Matcher 与 Validator）

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
  Signatures:   [][]byte{sig}, // 需真实 validator 签名
}

_, _ = client.Validation.ValidateIntentsBySignatures(ctx, []sdk.ValidationBundle{bundle})
```

更多示例见 `docs/quickstart.md`。

### 示例脚本

- `examples/send_intent`: 纯环境变量驱动的 CLI，默认 dry-run，演示单条提交与带签名的批量提交（需 `PIN_RPC_URL`/`PIN_PRIVATE_KEY`，可选 `SUBNET_ID`、`SIGNED_INTENT_ID` 等）。
- `examples/assign_intents`: Matcher 端批量分配脚本，支持自动签名单条分配或使用外部签名（需 `INTENT_ID`/`AGENT_ADDRESS` 等）。
- `examples/validate_intents`: Validator 端验证脚本，可计算 digest 并签名单验证者或加载预签名的多验证者 bundle。
- `examples/register_agent`: 基于 Subnet 合约的代理注册脚本，可按需设置 `AGENT_DOMAIN`/`AGENT_ENDPOINT`/`AGENT_METADATA_URI` 以及 `AGENT_VALUE_WEI`（默认 dry-run，可覆盖 0 值质押需求）。
- `examples/list_subnets`: 列出当前网络下的子网 ID 与活跃状态，便于在其他脚本中复用。
- `examples/complete_workflow`（规划中）：演示 Intent → Assignment → Validation → Checkpoint 的端到端流程。

## 连接 Subnet 与角色注册

### 连接子网合约（两种方式）

1) 通过子网 ID 查询地址并获取服务

```go
subnetID := sdk.MustBytes32FromHex("0x...")
subnetSvc, err := client.SubnetServiceByID(ctx, subnetID)
if err != nil { log.Fatal(err) }
```

2) 已知子网合约地址直接绑定

```go
subnetAddr := common.HexToAddress("0xSubnET...")
subnetSvc, err := client.SubnetServiceByAddress(subnetAddr)
if err != nil { log.Fatal(err) }
```

两种方式内部都会共享 `Client` 的 `TxManager` 与 `Signer`，保持一致的 nonce 与 1559 费用策略。

### 注册 3 种参与者角色

Subnet 合约的 `RegisterParticipant`/`RegisterParticipantERC20` 已封装为 6 个便捷方法：

- `RegisterValidator` / `RegisterValidatorERC20`
- `RegisterAgent` / `RegisterAgentERC20`
- `RegisterMatcher` / `RegisterMatcherERC20`

ETH 质押注册（向合约发送原生代币，可用于满足最小质押）：

```go
tx, err := subnetSvc.RegisterValidator(ctx, sdk.RegisterParticipantParams{
  Domain:      "example.org",
  Endpoint:    "https://validator.example.org",
  MetadataURI: "ipfs://...",
  Value:       big.NewInt(1e18), // 例如 1 ETH，按照注册费用查询
})
if err != nil { log.Fatal(err) }
log.Printf("validator registered, tx=%s", tx.Hash())

// 代理（Agent）与匹配器（Matcher）同理：
_, _ = subnetSvc.RegisterAgent(ctx, sdk.RegisterParticipantParams{Domain: "...", Endpoint: "...", MetadataURI: "...", Value: big.NewInt(0)})
_, _ = subnetSvc.RegisterMatcher(ctx, sdk.RegisterParticipantParams{Domain: "...", Endpoint: "...", MetadataURI: "...", Value: big.NewInt(0)})
```

ERC20 质押注册（使用治理配置的质押代币；调用前需确保已对需要拉取资金的合约设置 allowance）：

```go
tx, err := subnetSvc.RegisterValidatorERC20(ctx, sdk.RegisterParticipantERC20Params{
  Amount:      big.NewInt(1_000_000_000_000_000_000), // 1 token（按代币精度）
  Domain:      "example.org",
  Endpoint:    "https://validator.example.org",
  MetadataURI: "ipfs://...",
})
if err != nil { log.Fatal(err) }
log.Printf("validator (ERC20) registered, tx=%s", tx.Hash())
```

提示：
- 具体最小质押与注册审批由子网的 `StakeGovernanceConfig` 与 `autoApprove` 等参数决定；如为手动审批，注册后需要子网治理批准。
- ERC20 路径通常需要对质押接收方（例如 `StakingManager` 或 `Subnet`，以合约实现为准）进行 `approve` 授权。
- 可用查询：
  - `subnetSvc.IsActiveParticipant(ctx, addr, sdk.ParticipantValidator)`
  - `subnetSvc.ListActiveParticipants(ctx, sdk.ParticipantValidator)`
  - `subnetSvc.GetParticipantInfo/GetParticipantStakeInfo/GetParticipantCount`

## 网络与地址

- 预置网络：`base`(8453)、`base_sepolia`(84532)、`local`(31337)
- 覆盖优先级：代码级覆盖 > 环境变量 > SDK 默认（mock/local）
- 环境变量命名（示例）：
  - `PIN_BASE_INTENT_MANAGER` / `PIN_BASE_SEPOLIA_SUBNET_FACTORY` / `PIN_LOCAL_CHECKPOINT_MANAGER`
- 详见：`docs/addresses.md` 与 `.env.example`

## TxManager（交易管理）

可配置的 EIP‑1559 交易管理器，支持：

- nonce 源：`pending`/`latest`
- Gas 估算与乘数、建议费率、`MaxFeePerGas`/`MaxPriorityFee`
- 替换策略：超时 `ReplaceAfter` 后提升 `BumpPercent` 并重发
- Dry‑run：`NoSend` 仅构造、不发送

在 `sdk.Config.Tx` 中设置，或使用默认策略。详情见 `docs/txmanager.md`。

## 签名与 Digest

- 批量签名摘要：
  - typeHash: `PIN_INTENT_V1(bytes32,bytes32,address,bytes32,bytes32,uint256,address,uint256,address,uint256)`
  - digest: `keccak256(abi.encode(typeHash, intent_id, subnet_id, requester, keccak256(bytes(intent_type)), params_hash, deadline, payment_token, amount, address(this), chainid))`
- 链下签名：EIP‑191（eth_sign 前缀），与合约 `SignatureLib.verifySingleSignature()` 对齐
- 工具函数：`sdkcrypto.ComputeIntentDigest()`、`client.Intent.SignDigest()`
- 其他摘要：`client.Assignment.ComputeDigest()`、`client.Validation.ComputeDigest()`、`client.CheckpointManager.ComputeDigest()`
- 详见：`docs/signing.md`

## 目录结构

- `sdk/`：对外 API（Client、TxManager、Signer、AddressBook、高层 Service）
- `contracts/`：按合约模块分目录的 abigen 绑定（已内置）
- `docs/`：中文文档与规范

## 生成合约绑定（开发者）

绑定已随仓库提供。如需从 `RootLayer/artifacts` 再生，可参考：

```bash
# 需要 abigen 与 jq
for name in IntentManager SubnetFactory Subnet StakingManager CheckpointManager; do
  jq -r '.abi' /path/to/RootLayer/artifacts/contracts/${name}.sol/${name}.json > /tmp/${name}.abi
  pkg=$(echo "$name" | tr '[:upper:]' '[:lower:]')
  abigen --abi /tmp/${name}.abi --pkg $pkg --type ${name} --out contracts/$pkg/${pkg}.go
done
```

## 常见问题

- 地址未生效？
  - 检查是否设置了匹配网络的环境变量；或使用 `sdk.Config.Addresses` 进行代码级覆盖
- 交易卡住？
  - 开启替换策略（`ReplaceStuck=true`，设置 `ReplaceAfter` 与 `BumpPercent`）
- EIP‑191 v 值？
  - SDK 本地签名器会将 `v∈{0,1}` 归一化为 `27/28`

## 兼容性与安全

- 管理员/治理函数不做“易用封装”，仅暴露底层绑定接口；请严格遵守 `GUARDIAN_ROLE`/`GOVERNANCE_ROLE` 的访问控制
- 与合约的签名/哈希严格对齐，digest 绑定 `address(this)` 与 `chainId`，具备跨链/跨合约防重放能力

## 参考文档

- 文档导航：`docs/README.md`
- 地址与网络：`docs/addresses.md`
- 配置与环境：`docs/config.md`
- 签名规范：`docs/signing.md`
- 交易管理：`docs/txmanager.md`
- 快速开始：`docs/quickstart.md`
