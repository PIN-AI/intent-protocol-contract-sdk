# 配置与环境变量

本节介绍 SDK 的配置方式、环境变量约定与本地开发配置（local）。

## 基本配置项

- `PIN_RPC_URL`（必需）：以太坊节点 RPC 地址
- `PIN_PRIVATE_KEY`（可选）：0x 开头的私钥（如使用远程签名器可不提供）
- `PIN_NETWORK`（可选）：`base` | `base_sepolia` | `local`，不传则 SDK 会通过 RPC 自动探测 `chainId`

合约地址通过环境变量或代码传入，见 docs/addresses.md。

## 环境变量优先级

1. 代码级覆盖（在初始化 SDK 时显式传入 AddressBook）
2. 环境变量（按当前网络选择相应前缀）
3. SDK 内置默认地址（Mock 或 local 默认）

## .env 示例

见仓库根目录 `.env.example`。你可以复制为 `.env` 并填入真实值。

```ini
# 基础网络配置
PIN_NETWORK=base_sepolia
PIN_RPC_URL=https://sepolia.base.org
PIN_PRIVATE_KEY=0x你的私钥

# base_sepolia 合约地址（如未提供，将使用SDK默认mock地址）
PIN_BASE_SEPOLIA_INTENT_MANAGER=0x5555555555555555555555555555555555555555
PIN_BASE_SEPOLIA_SUBNET_FACTORY=0x6666666666666666666666666666666666666666
PIN_BASE_SEPOLIA_STAKING_MANAGER=0x7777777777777777777777777777777777777777
PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER=0x8888888888888888888888888888888888888888

# 可选：本地开发（local）配置
# 当 PIN_NETWORK=local 或 RPC chainId=31337 时生效
PIN_LOCAL_INTENT_MANAGER=0x9999999999999999999999999999999999999999
PIN_LOCAL_SUBNET_FACTORY=0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
PIN_LOCAL_STAKING_MANAGER=0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
PIN_LOCAL_CHECKPOINT_MANAGER=0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC
```

## 本地开发（local）

- 推荐使用 Hardhat/Anvil，chainId=31337
- 若你在本地重新部署了合约，请更新对应的 `PIN_LOCAL_*` 环境变量
- 未配置时，SDK 将使用内置的 local 默认地址（仅用于开发）

## 代码级覆盖（示意）

后续 SDK 初始化会支持通过 `AddressBook` 传入覆盖（示意 API）：

```go
cfg := sdk.Config{
    RPCURL: "http://127.0.0.1:8545",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Network: "local",
    Addresses: sdk.AddressBook{
        Network: "local",
        IntentManager:      common.HexToAddress("0x..."),
        SubnetFactory:      common.HexToAddress("0x..."),
        StakingManager:     common.HexToAddress("0x..."),
        CheckpointManager:  common.HexToAddress("0x..."),
    },
}
client, err := sdk.NewClient(cfg)
```

> 上述 API 将在实现阶段提供；当前文档用于先行约定配置方式。
