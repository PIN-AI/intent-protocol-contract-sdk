# 签名与 Digest 规范

为对齐合约 `IntentManager.submitIntentsBySignatures()` 与 `SignatureLib.verifySingleSignature()`，SDK 将提供 EIP-191（以太坊消息前缀）签名工具与 digest 计算方法。

## Digest 构造（批量签名提交）

合约对每条 Intent 的摘要（digest）计算如下：

```
typeHash = "PIN_INTENT_V1(bytes32,bytes32,address,bytes32,bytes32,uint256,address,uint256,address,uint256)"

digest = keccak256(abi.encode(
    keccak256(typeHash),
    intent_id,                 // bytes32
    subnet_id,                 // bytes32
    requester,                 // address
    keccak256(bytes(intent_type)), // bytes32
    params_hash,               // bytes32
    deadline,                  // uint256
    payment_token,             // address
    amount,                    // uint256
    address(this),             // IntentManager 合约地址
    block.chainid              // 当前链ID
))
```

SDK 会自动注入 `IntentManager` 地址与 `chainId`，调用方只需提供其他字段。

## 签名方式（EIP-191）

- 使用以太坊消息前缀（EIP-191 / eth_sign）对上面的 `digest` 进行签名：
  - `ethSignedMessage = keccak256("\x19Ethereum Signed Message:\n32" || digest)`
  - `signature = Sign(ethSignedMessage, privateKey)`
- 合约端通过 `SignatureLib.verifySingleSignature()` 进行恢复与校验。

## 注意事项

- v 值标准化：部分签名实现返回 v∈{0,1}，应转换为 27/28；go-ethereum 的 `crypto.Sign` 返回 65 字节签名，末位需要按实现确认是否已为 27/28
- `intent_type` 需要先 `keccak256(bytes(intent_type))`
- `payment_token == address(0)` 表示使用 ETH（批量提交时 `totalEthRequired` 需与所有 ETH 意图金额之和一致）

## SDK 计划提供的工具（示意 API）

```go
// 计算 digest（自动注入 contract 与 chainId）
func ComputeIntentDigest(input SignedIntentInput, intentManager common.Address, chainID *big.Int) ([32]byte, error)

// 对 digest 进行 EIP-191 签名
func SignIntentDigest(digest [32]byte, s Signer) ([]byte, error)

// Signer 接口：支持本地私钥签名器与可插拔远程签名器
type Signer interface {
    Address() common.Address
    SignDigest([32]byte) ([]byte, error)
}
```

> 后续若合约支持 EIP-712，将在 SDK 中提供 `ComputeEIP712Digest` 等方法以便切换。
