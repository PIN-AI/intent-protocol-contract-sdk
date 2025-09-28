# 快速开始（规划中的 API）

本节给出 SDK 的典型使用姿势。实际代码将在实现阶段提供，接口会尽量保持一致。

## 1. 安装（预期）

```bash
go get github.com/PIN-AI/intent-protocol-contract-sdk@latest
```

## 2. 配置环境变量

参见 `.env.example` 与 docs/config.md。最少需要：

- `PIN_RPC_URL`
- `PIN_PRIVATE_KEY`（或使用自定义 Signer）
- （可选）`PIN_NETWORK` 与合约地址变量

## 3. 初始化 Client（示意）

```go
package main

import (
    "context"
    "log"
    "math/big"
    "os"

    sdk "github.com/PIN-AI/intent-protocol-contract-sdk/sdk"
)

func main() {
    cfg := sdk.Config{
        RPCURL:        os.Getenv("PIN_RPC_URL"),
        PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
        Network:       os.Getenv("PIN_NETWORK"), // 可选：base | base_sepolia | local
        // Addresses: 可选：代码级覆盖合约地址
        Tx: sdk.TxConfig{ Use1559: true },
    }

    client, err := sdk.NewClient(cfg)
    if err != nil {
        log.Fatalf("init sdk: %v", err)
    }

    _ = client // 后续使用 client.Intent 等高层接口
}
```

## 4. 提交 Intent（ETH 预算，示意）

```go
params := sdk.SubmitIntentParams{
    IntentID:     sdk.MustBytes32FromHex("0x..."),
    SubnetID:     sdk.MustBytes32FromHex("0x..."),
    IntentType:   "book_flight",
    ParamsHash:   sdk.HashParams([]byte("{...}")),
    DeadlineUnix: time.Now().Add(24*time.Hour).Unix(),
    PaymentToken: common.Address{}, // 0 地址表示 ETH
    AmountWei:    big.NewInt(1e16),
    Value:        big.NewInt(1e16), // 与 AmountWei 一致
}
tx, err := client.Intent.SubmitIntent(ctx, params)
```

## 5. 批量签名并提交（示意）

```go
input := sdk.SignedIntentInput{
    IntentID:   sdk.MustBytes32FromHex("0x..."),
    SubnetID:   sdk.MustBytes32FromHex("0x..."),
    Requester:  signer.Address(),
    IntentType: "book_flight",
    ParamsHash: sdk.MustBytes32FromHex("0x..."),
    Deadline:   big.NewInt(deadlineUnix),
    PaymentToken: common.Address{},
    Amount:     big.NewInt(1e16),
}

digest, _ := sdkcrypto.ComputeIntentDigest(input, client.Addresses.IntentManager, client.ChainID)
sig, _ := client.Signer.SignDigest(digest)

tx, err := client.Intent.SubmitIntentsBySignatures(ctx, []sdk.SignedIntent{{Data: input, Signature: sig}}, big.NewInt(1e16))
```

> 以上为示意 API，落地实现时我们会保证与合约签名校验逻辑完全一致。
