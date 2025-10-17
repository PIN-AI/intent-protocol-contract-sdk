# 交易管理（TxManager）

TxManager 负责封装交易的 nonce、gas 估算、EIP-1559 费用、超时重试与替换策略。

## 目标

- 默认安全的 EIP-1559 参数选择
- 并发安全的 nonce 分配（基于 pending 或 latest）
- 自动 gas 估算与安全余量（`GasLimitMultiplier`）
- 可配置的替换策略（交易卡住时提升 gas 再发）
- 支持 Dry-Run（NoSend）模式用于离线构造与审计

## Gas 估算与安全余量

TxManager 提供自动 gas 估算并应用安全余量，避免 out-of-gas 错误：

### 工作机制

1. **当 `GasLimit==0` 且 `GasLimitMultiplier > 1.0` 时**：
   - TxManager 先用 `NoSend=true` 执行一次交易获取估算值
   - 对估算值应用 `GasLimitMultiplier`（默认 1.5，即 50% 余量）
   - 检查是否超过 `GasCeil` 上限（若配置）
   - 将结果设置为实际 `GasLimit`
   - 再真正发送交易

2. **当 `GasLimit > 0` 时**：
   - 直接使用指定的固定值，不估算

3. **推荐配置**：
   - 生产环境：`GasLimit=0`，`GasLimitMultiplier=1.2`（20% 余量）
   - 批量操作：`GasLimitMultiplier=1.3`（30% 余量）

### 示例

```go
cfg := sdk.Config{
    RPCURL: "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        GasLimitMultiplier: ptr(1.2), // 20% 安全余量
    },
}
client, _ := sdk.NewClient(ctx, cfg)

// ValidateIntentsBySignatures 会自动估算并应用 1.2x 余量
tx, err := client.Validation.ValidateIntentsBySignatures(ctx, bundles)
```

## 配置项

```go
type TxConfig struct {
    Use1559           bool          // 默认 true
    GasLimit          uint64        // 0=自动估算（推荐），>0=固定值
    GasLimitMultiplier float64      // 估算安全系数，默认 1.5
    GasCeil           uint64        // gas 估算上限，超过则拒绝发送（0=不限制）

    MaxFeePerGas         *big.Int   // 不填则根据建议值推导
    MaxPriorityFeePerGas *big.Int   // 不填则使用 SuggestGasTipCap

    NonceSource string   // "pending" | "latest"，默认 "pending"

    // 替换/重试
    ReplaceStuck  bool
    ReplaceAfter  time.Duration // 例如 30s
    BumpPercent   float64       // 例如 12.5（提升 12.5%）

    // 调试
    NoSend bool // 仅构造交易，不发送
}
```

### GasCeil 保护

`GasCeil` 用于防止 gas 估算异常高的情况：

- **触发条件**：估算值（应用乘数后）超过 `GasCeil`
- **行为**：拒绝发送交易，返回详细错误信息（包含原始估算值、乘数、上限）
- **适用场景**：
  - 批量操作（防止 items 过多导致 gas 爆炸）
  - 生产环境（控制单次交易成本上限）
  - 防止接近区块 gas limit（Base L2 为 30M）
- **推荐配置**：
  - 单条操作：`1000000`（100万）
  - 批量操作（<10 items）：`2000000`（200万）
  - 批量操作（>10 items）：`5000000`（500万）

## 用法示例

### 基础配置（推荐）

```go
cfg := sdk.Config{
    RPCURL: "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        Use1559: ptr(true),
        GasLimitMultiplier: ptr(1.2),    // 20% 安全余量
        GasCeil: ptr(uint64(2000000)),   // 200万上限（批量操作）
    },
}
client, err := sdk.NewClient(ctx, cfg)

// 自动估算并应用 1.2x 余量，超过 200万 则拒绝
tx, err := client.Validation.ValidateIntentsBySignatures(ctx, bundles)
if err != nil {
    // 如果是 gas 超限错误，可以检查详细信息
    var gasCeilErr *txmgr.ErrGasCeilExceeded
    if errors.As(err, &gasCeilErr) {
        log.Printf("Gas too high: raw=%d, adjusted=%d (%.2fx), ceil=%d",
            gasCeilErr.Raw, gasCeilErr.Adjusted, gasCeilErr.Multiplier, gasCeilErr.Ceil)
        // 考虑分批或调整参数
    }
    return err
}
```

### 生产环境完整配置

```go
cfg := sdk.Config{
    RPCURL: "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: &sdk.TxOptions{
        Use1559: ptr(true),
        GasLimitMultiplier: ptr(1.2),    // 20% 安全余量
        GasCeil: ptr(uint64(2000000)),   // 200万上限
        ReplaceStuck: ptr(true),         // 开启卡顿替换
        ReplaceAfter: ptr(30 * time.Second),
        BumpPercent: ptr(12.5),
    },
}
client, err := sdk.NewClient(ctx, cfg)

// 使用示例
tx, err := client.Validation.ValidateIntentsBySignatures(ctx, bundles)
```

> 实现结合了 go-ethereum 的 `SuggestGasTipCap`、`SuggestGasPrice` 与自动 gas 估算，确保与 Base L2 的计费模型兼容。
