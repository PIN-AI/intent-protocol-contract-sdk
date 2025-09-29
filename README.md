# intent-protocol-contract-sdk

PIN AI Intent Protocol Rootlayer 智能合约 Go SDK。

面向 RootLayer 与 Subnet 的链上交互封装，预置 Base 主网/测试网与本地开发网络的合约地址（可通过环境变量与代码覆盖），提供 EIP‑1559 交易管理与 EIP‑191 摘要签名工具，并暴露常用的高层 API（提交/批量提交 Intent、查询子网/验证器/检查点、质押操作等）。

• 快速导航：`docs/README.md`
• 环境变量示例：`.env.example`

---

## 功能特性

- 强类型合约绑定：`IntentManager`、`SubnetFactory`、`Subnet`、`StakingManager`、`CheckpointManager`
- 高层 API：
  - Intent：提交/批量签名提交、完成/失败标记、过期处理、状态查询
  - SubnetFactory：活跃状态/合约地址查询、列表、预测地址、暂停/恢复/废弃（需权限）
  - Staking：质押/解押/提现、质押信息查询（管理员函数不做“易用封装”，仅底层绑定）
  - Checkpoint：查询检查点与证明
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
  tx, err := client.IntentManager.SubmitIntent(ctx, params)
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

digest, _ := client.IntentManager.ComputeDigest(input)
sig, _ := client.IntentManager.SignDigest(digest)

tx, err := client.IntentManager.SubmitIntentsBySignatures(ctx, sdk.SubmitIntentBatchParams{
  Items: []sdk.SignedIntent{{Data: input, Signature: sig}},
  // ETH 总额可省略，SDK 会自动相加 PaymentToken==0 的金额
})
```

更多示例见 `docs/quickstart.md`。

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
- 工具函数：`sdkcrypto.ComputeIntentDigest()`、`client.IntentManager.SignDigest()`
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
