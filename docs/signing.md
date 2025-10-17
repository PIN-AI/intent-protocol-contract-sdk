# Signing & Digest Specification

**[中文文档](signing-zh.md)** | English

To align with contract methods `IntentManager.submitIntentsBySignatures()` and `SignatureLib.verifySingleSignature()`, the SDK provides EIP-191 (Ethereum message prefix) signing utilities and digest computation methods.

## Digest Construction (Batch Signed Submission)

The contract calculates the digest for each Intent as follows:

```
typeHash = "PIN_INTENT_V1(bytes32,bytes32,address,bytes32,bytes32,uint256,address,uint256,address,uint256)"

digest = keccak256(abi.encode(
    keccak256(typeHash),
    intent_id,                     // bytes32
    subnet_id,                     // bytes32
    requester,                     // address
    keccak256(bytes(intent_type)), // bytes32
    params_hash,                   // bytes32
    deadline,                      // uint256
    payment_token,                 // address
    amount,                        // uint256
    address(this),                 // IntentManager contract address
    block.chainid                  // Current chain ID
))
```

The SDK automatically injects the `IntentManager` address and `chainId`. Callers only need to provide the other fields.

## Signing Method (EIP-191)

- Use Ethereum message prefix (EIP-191 / eth_sign) to sign the above `digest`:
  - `ethSignedMessage = keccak256("\x19Ethereum Signed Message:\n32" || digest)`
  - `signature = Sign(ethSignedMessage, privateKey)`
- The contract verifies using `SignatureLib.verifySingleSignature()`.

## Intent Digest Details

### TypeHash

The typeHash defines the structure of the data being signed:

```solidity
bytes32 constant TYPE_HASH = keccak256(
    "PIN_INTENT_V1(bytes32,bytes32,address,bytes32,bytes32,uint256,address,uint256,address,uint256)"
);
```

### Digest Fields

| Field | Type | Description |
|-------|------|-------------|
| `intent_id` | bytes32 | Unique identifier for the intent |
| `subnet_id` | bytes32 | Target subnet identifier |
| `requester` | address | Address submitting the intent |
| `keccak256(bytes(intent_type))` | bytes32 | Hash of intent type string |
| `params_hash` | bytes32 | Hash of intent parameters |
| `deadline` | uint256 | Unix timestamp deadline |
| `payment_token` | address | Token address (0x0 for ETH) |
| `amount` | uint256 | Payment amount |
| `address(this)` | address | IntentManager contract address |
| `block.chainid` | uint256 | Chain ID for replay protection |

### Example Computation

```go
import (
    sdkcrypto "github.com/PIN-AI/intent-protocol-contract-sdk/sdk/crypto"
)

input := sdkcrypto.SignedIntentInput{
    IntentID:     [32]byte{0x00, 0x01, ...},
    SubnetID:     [32]byte{0x01, 0x00, ...},
    Requester:    common.HexToAddress("0xRequester"),
    IntentType:   "book_flight",
    ParamsHash:   [32]byte{0x12, 0x34, ...},
    Deadline:     big.NewInt(1735689600),
    PaymentToken: common.Address{}, // 0x0 for ETH
    Amount:       big.NewInt(1e16),
}

// Compute digest (auto-injects contract address and chainId)
digest, err := client.Intent.ComputeDigest(input)
if err != nil {
    log.Fatal(err)
}

// Sign digest with EIP-191
signature, err := client.Intent.SignDigest(digest)
if err != nil {
    log.Fatal(err)
}

// signature is ready for submitIntentsBySignatures()
```

## Assignment Digest

For Matcher assignment operations:

```go
assignment := sdk.AssignmentData{
    AssignmentID: [32]byte{...},
    IntentID:     [32]byte{...},
    BidID:        [32]byte{...},
    Agent:        common.HexToAddress("0xAgent"),
    Status:       sdk.AssignmentStatusActive,
    Matcher:      common.HexToAddress("0xMatcher"),
}

digest, err := client.Assignment.ComputeDigest(assignment)
signature, err := client.Assignment.SignDigest(digest)

// Use in AssignIntentsBySignatures()
```

## Validation Digest

For Validator validation operations:

```go
bundle := sdk.ValidationBundle{
    IntentID:     [32]byte{...},
    AssignmentID: [32]byte{...},
    SubnetID:     [32]byte{...},
    Agent:        common.HexToAddress("0xAgent"),
    ResultHash:   [32]byte{...},
    ProofHash:    [32]byte{...},
    RootHeight:   big.NewInt(12345),
    RootHash:     [32]byte{...},
    Validators:   []common.Address{...},
}

digest, err := client.Validation.ComputeDigest(bundle)
signature, err := client.Validation.SignDigest(digest)

bundle.Signatures = [][]byte{signature}
// Use in ValidateIntentsBySignatures()
```

## Checkpoint Digest

For Checkpoint submission:

```go
checkpointData := sdk.CheckpointData{
    SubnetID:    [32]byte{...},
    Height:      big.NewInt(1000),
    StateRoot:   [32]byte{...},
    IntentCount: big.NewInt(42),
    Timestamp:   big.NewInt(time.Now().Unix()),
}

digest, err := client.CheckpointManager.ComputeDigest(checkpointData)
signature, err := client.CheckpointManager.SignDigest(digest)

// Use in SubmitCheckpoint()
```

## Important Notes

### V Value Normalization

- Some signature implementations return `v ∈ {0,1}`, which must be converted to `27/28`
- go-ethereum's `crypto.Sign` returns a 65-byte signature; the last byte may need adjustment depending on implementation
- **The SDK automatically handles v value normalization** when using `SignDigest()`

### Intent Type Hashing

- `intent_type` must be hashed first: `keccak256(bytes(intent_type))`
- Example: `book_flight` → `keccak256("book_flight")`
- The SDK handles this automatically in `ComputeDigest()`

### ETH Payment

- `payment_token == address(0)` indicates ETH payment
- For batch submission, `totalEthRequired` must equal the sum of all ETH intent amounts
- **The SDK auto-calculates this** unless manually specified

### Replay Protection

The digest binds to:
- Contract address (`address(this)`) - prevents cross-contract replay
- Chain ID (`block.chainid`) - prevents cross-chain replay

This ensures signatures are only valid for:
- The specific contract instance
- The specific blockchain network

## SDK Utility Functions

### Planned API (Example)

```go
// Compute digest (auto-injects contract address and chainId)
func (c *Client) ComputeIntentDigest(
    input SignedIntentInput,
) ([32]byte, error)

// Sign digest with EIP-191
func (c *Client) SignIntentDigest(
    digest [32]byte,
) ([]byte, error)

// Signer interface: supports local private key signers and pluggable remote signers
type Signer interface {
    Address() common.Address
    SignDigest([32]byte) ([]byte, error)
}
```

### Current Implementation

The SDK currently provides:

```go
// IntentService methods
digest, err := client.Intent.ComputeDigest(input)
signature, err := client.Intent.SignDigest(digest)

// AssignmentService methods
digest, err := client.Assignment.ComputeDigest(assignment)
signature, err := client.Assignment.SignDigest(digest)

// ValidationService methods
digest, err := client.Validation.ComputeDigest(bundle)
signature, err := client.Validation.SignDigest(digest)

// CheckpointService methods
digest, err := client.CheckpointManager.ComputeDigest(checkpointData)
signature, err := client.CheckpointManager.SignDigest(digest)
```

## Signature Format

### Standard ECDSA Signature

The signature is a 65-byte array:
- `r` (32 bytes): First component of ECDSA signature
- `s` (32 bytes): Second component of ECDSA signature
- `v` (1 byte): Recovery ID (27 or 28)

### Signature Verification

The contract uses `SignatureLib.verifySingleSignature()` which:
1. Extracts `r`, `s`, `v` from the signature
2. Recovers the signer address using `ecrecover(digest, v, r, s)`
3. Compares recovered address with expected signer

### Example Verification (Conceptual)

```go
// On-chain verification (conceptual Solidity)
bytes32 digest = computeDigest(...);
address signer = ecrecover(digest, v, r, s);
require(signer == expectedSigner, "Invalid signature");
```

## Testing & Validation

The SDK includes comprehensive tests for digest computation:

```bash
# Run signing tests
go test -v ./sdk/crypto/

# All 52 digest calculation tests pass
```

### Test Coverage

- Intent digest computation
- Assignment digest computation
- Validation digest computation
- Checkpoint digest computation
- EIP-191 signature generation
- Signature verification alignment with contracts

## EIP-712 Support (Future)

> If the contract adds EIP-712 support in the future, the SDK will provide `ComputeEIP712Digest` methods for structured data signing.

EIP-712 would provide:
- Better wallet UI for signature prompts
- Structured data display in MetaMask/hardware wallets
- Domain separation with chain and contract info

## Common Issues & Solutions

### Issue: "Invalid signature"

**Possible causes:**
1. Incorrect digest computation (field order, missing hash, wrong contract address)
2. V value not normalized (must be 27 or 28)
3. Wrong chain ID used
4. Signature from different private key

**Solution:**
- Use SDK's built-in `ComputeDigest()` and `SignDigest()` methods
- Verify contract address matches `client.Addresses.IntentManager`
- Ensure chainId matches network: `client.ChainID`

### Issue: "Signature verification failed"

**Possible causes:**
1. Digest mismatch between client and contract
2. IntentType not properly hashed
3. Using wrong contract address in digest

**Solution:**
- Double-check all input fields match exactly
- Verify `intent_type` is hashed with `keccak256(bytes(intent_type))`
- Ensure using correct `IntentManager` address from `client.Addresses`

### Issue: "Signer mismatch"

**Possible causes:**
1. Signed with wrong private key
2. Requester address doesn't match signer

**Solution:**
- Verify `input.Requester == client.Signer.Address()`
- Check private key corresponds to expected address

## Best Practices

1. **Always use SDK digest computation** - Ensures correct field ordering and hashing
2. **Verify before batch submission** - Test single signature before large batches
3. **Log digest and signature** (in development) - Helps debug verification issues
4. **Use correct contract address** - Different networks have different addresses
5. **Match chain ID** - Ensure client connected to correct network

## Related Documentation

- [Quick Start](quickstart.md) - Signing examples
- [API Reference](api-reference.md) - Complete signing method reference
- [Troubleshooting](troubleshooting.md) - Signature issue debugging
