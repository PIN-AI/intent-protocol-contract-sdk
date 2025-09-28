# 合约地址与网络配置

SDK 通过“预置默认地址 + 环境变量覆盖 + 代码级覆盖”的方式选择合约地址。

- 预置网络：
  - base（主网，chainId=8453）
  - base_sepolia（测试网，chainId=84532）
  - local（本地开发，推荐 Hardhat/Anvil，chainId=31337）

## 覆盖优先级

1. 代码级覆盖：在初始化 SDK 时传入自定义 AddressBook（最高优先级）
2. 环境变量：通过固定命名的环境变量注入地址
3. SDK 内置默认地址（Mock 或 local 默认地址）

> 注意：生产环境请务必通过环境变量或代码覆盖来设置真实部署地址。

## 环境变量命名

- 通用：
  - `PIN_NETWORK`：`base` | `base_sepolia` | `local`（可选，不传则由 RPC 自动探测 chainId 匹配）
  - `PIN_RPC_URL`：以太坊 RPC 端点
  - `PIN_PRIVATE_KEY`：十六进制私钥（0x 前缀），或使用自定义 Signer

- 合约地址（按网络区分）：
  - base：
    - `PIN_BASE_INTENT_MANAGER`
    - `PIN_BASE_SUBNET_FACTORY`
    - `PIN_BASE_STAKING_MANAGER`
    - `PIN_BASE_CHECKPOINT_MANAGER`
  - base_sepolia：
    - `PIN_BASE_SEPOLIA_INTENT_MANAGER`
    - `PIN_BASE_SEPOLIA_SUBNET_FACTORY`
    - `PIN_BASE_SEPOLIA_STAKING_MANAGER`
    - `PIN_BASE_SEPOLIA_CHECKPOINT_MANAGER`
  - local：
    - `PIN_LOCAL_INTENT_MANAGER`
    - `PIN_LOCAL_SUBNET_FACTORY`
    - `PIN_LOCAL_STAKING_MANAGER`
    - `PIN_LOCAL_CHECKPOINT_MANAGER`

## 默认地址（Mock 与 Local）

当未通过代码或环境变量提供地址时，SDK 将回落到以下默认地址：

- base（8453）：
  - IntentManager：`0x1111111111111111111111111111111111111111`
  - SubnetFactory：`0x2222222222222222222222222222222222222222`
  - StakingManager：`0x3333333333333333333333333333333333333333`
  - CheckpointManager：`0x4444444444444444444444444444444444444444`

- base_sepolia（84532）：
  - IntentManager：`0x5555555555555555555555555555555555555555`
  - SubnetFactory：`0x6666666666666666666666666666666666666666`
  - StakingManager：`0x7777777777777777777777777777777777777777`
  - CheckpointManager：`0x8888888888888888888888888888888888888888`

- local（31337）：
  - IntentManager：`0x9999999999999999999999999999999999999999`
  - SubnetFactory：`0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA`
  - StakingManager：`0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB`
  - CheckpointManager：`0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC`

> 提示：`local` 用于本地开发与集成测试。若你本地有实际部署的地址，请用环境变量覆盖默认值。

## 常见问题

- Q: 配置了 `PIN_NETWORK` 但地址没生效？
  - A: 检查是否设置了对应网络前缀的环境变量名；此外，SDK 也会通过 RPC 读取 `chainId` 做自动匹配。
- Q: 多网络地址同时配置时如何选择？
  - A: SDK 会优先读取与当前网络匹配（`PIN_NETWORK` 或链上 `chainId`）的环境变量组。
