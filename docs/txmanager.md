# 交易管理（TxManager）

TxManager 负责封装交易的 nonce、gas 估算、EIP-1559 费用、超时重试与替换策略。

## 目标

- 默认安全的 EIP-1559 参数选择
- 并发安全的 nonce 分配（基于 pending 或 latest）
- 可配置的替换策略（交易卡住时提升 gas 再发）
- 支持 Dry-Run（NoSend）模式用于离线构造与审计

## 配置项（计划）

```go
type TxConfig struct {
    Use1559           bool          // 默认 true
    GasLimit          uint64        // 0=自动估算
    GasLimitMultiplier float64      // 估算乘数，默认 1.1

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

## 用法示例（计划 API）

```go
cfg := sdk.Config{
    RPCURL: "...",
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    Tx: sdk.TxConfig{
        Use1559: true,
        GasLimitMultiplier: 1.1,
        ReplaceStuck: true,
        ReplaceAfter: 30 * time.Second,
        BumpPercent: 12.5,
    },
}
client, err := sdk.NewClient(cfg)
```

> 具体实现会结合 go-ethereum 的 `SuggestGasTipCap`、`SuggestGasPrice`、`EstimateGas` 等，确保与 Base L2 的计费模型兼容。
