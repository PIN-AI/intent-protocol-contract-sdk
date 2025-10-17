# Troubleshooting

**[中文文档](troubleshooting-zh.md)** | English

Common issues and solutions when working with the PIN Intent Protocol SDK.

## Table of Contents

1. [Connection Issues](#connection-issues)
2. [Transaction Errors](#transaction-errors)
3. [Signature Issues](#signature-issues)
4. [Gas Problems](#gas-problems)
5. [Contract Interaction](#contract-interaction)
6. [Configuration Problems](#configuration-problems)

## Connection Issues

### Issue: "Connection refused" or "dial tcp: connection refused"

**Symptoms**:
```
Error: Post "http://127.0.0.1:8545": dial tcp 127.0.0.1:8545: connect: connection refused
```

**Causes**:
- RPC node is not running
- Wrong RPC URL
- Firewall blocking connection
- Network connectivity issues

**Solutions**:

1. **Verify RPC URL**:
```go
// Test connection
client, err := ethclient.Dial(os.Getenv("PIN_RPC_URL"))
if err != nil {
    log.Fatal("Cannot connect to RPC:", err)
}
defer client.Close()

// Check if node is responding
blockNumber, err := client.BlockNumber(context.Background())
if err != nil {
    log.Fatal("RPC not responding:", err)
}
log.Printf("Connected! Current block: %d", blockNumber)
```

2. **Check local node status**:
```bash
# If using Anvil
anvil --chain-id 31337

# If using Hardhat
npx hardhat node

# Test connectivity
curl -X POST -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1"}' \
  http://127.0.0.1:8545
```

3. **Try alternative RPC providers**:
```go
// Use public RPC endpoints
rpcURLs := []string{
    "https://sepolia.base.org",
    "https://base-sepolia.drpc.org",
    "https://base-sepolia.public.blastapi.io",
}

for _, url := range rpcURLs {
    client, err := sdk.NewClient(ctx, sdk.Config{RPCURL: url, ...})
    if err == nil {
        break // Success
    }
}
```

### Issue: "Context deadline exceeded"

**Symptoms**:
```
Error: context deadline exceeded
```

**Causes**:
- Network timeout
- RPC node overloaded
- Slow transaction confirmation

**Solutions**:

1. **Increase context timeout**:
```go
// Increase timeout for slow networks
ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
defer cancel()

tx, err := client.Intent.SubmitIntent(ctx, params)
```

2. **Use retry logic**:
```go
func submitWithRetry(params sdk.SubmitIntentParams) (*types.Transaction, error) {
    maxRetries := 3
    for i := 0; i < maxRetries; i++ {
        ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
        defer cancel()

        tx, err := client.Intent.SubmitIntent(ctx, params)
        if err == nil {
            return tx, nil
        }

        if !strings.Contains(err.Error(), "deadline") {
            return nil, err // Non-timeout error
        }

        time.Sleep(time.Duration(i+1) * 5 * time.Second)
    }
    return nil, fmt.Errorf("max retries exceeded")
}
```

## Transaction Errors

### Issue: "Nonce too low"

**Symptoms**:
```
Error: nonce too low: address 0x..., tx: 5, state: 6
```

**Causes**:
- Nonce already used by confirmed transaction
- Multiple clients using same address
- Nonce source mismatch

**Solutions**:

1. **Switch to latest nonce**:
```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        NonceSource: ptr("latest"), // More conservative
    },
}
```

2. **Check pending transactions**:
```bash
# Check account nonce
cast nonce $ADDRESS --rpc-url $RPC_URL

# Check pending transactions
cast pending --from $ADDRESS --rpc-url $RPC_URL
```

3. **Clear stuck transactions**:
```go
// Get pending nonce
pendingNonce, err := client.Backend.PendingNonceAt(ctx, client.Signer.Address())

// Get confirmed nonce
confirmedNonce, err := client.Backend.NonceAt(ctx, client.Signer.Address(), nil)

log.Printf("Pending: %d, Confirmed: %d", pendingNonce, confirmedNonce)

// If gap exists, transactions are stuck
// Enable replacement or wait for confirmation
```

### Issue: "Transaction underpriced"

**Symptoms**:
```
Error: replacement transaction underpriced
Error: transaction underpriced
```

**Causes**:
- Gas price too low
- Network congestion
- Insufficient gas bump for replacement

**Solutions**:

1. **Use automatic fee calculation**:
```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        Use1559: ptr(true), // Enable EIP-1559
        // Don't set MaxFeePerGas manually
    },
}
```

2. **Increase bump percentage**:
```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        ReplaceStuck: ptr(true),
        BumpPercent:  ptr(15.0), // Increase from 10% to 15%
    },
}
```

3. **Check current gas prices**:
```go
// Get suggested gas price
gasPrice, err := client.Backend.SuggestGasPrice(ctx)
log.Printf("Suggested gas price: %s gwei", new(big.Int).Div(gasPrice, big.NewInt(1e9)))

// Get suggested tip
gasTipCap, err := client.Backend.SuggestGasTipCap(ctx)
log.Printf("Suggested tip: %s gwei", new(big.Int).Div(gasTipCap, big.NewInt(1e9)))
```

### Issue: "Insufficient funds"

**Symptoms**:
```
Error: insufficient funds for gas * price + value
```

**Causes**:
- Low account balance
- Underestimated gas requirements
- Concurrent transactions depleting balance

**Solutions**:

1. **Check account balance**:
```go
balance, err := client.Backend.BalanceAt(ctx, client.Signer.Address(), nil)
if err != nil {
    return err
}

ethBalance := new(big.Float).Quo(
    new(big.Float).SetInt(balance),
    big.NewFloat(1e18),
)
log.Printf("Account balance: %s ETH", ethBalance.String())

// Calculate required funds
gasLimit := tx.Gas()
gasPrice := tx.GasPrice()
value := tx.Value()

required := new(big.Int).Mul(big.NewInt(int64(gasLimit)), gasPrice)
required.Add(required, value)

if balance.Cmp(required) < 0 {
    return fmt.Errorf("insufficient balance: have %s, need %s", balance, required)
}
```

2. **Request testnet funds**:
```bash
# Base Sepolia faucet
https://faucet.quicknode.com/base/sepolia

# Alternative faucets
https://www.alchemy.com/faucets/base-sepolia
```

## Signature Issues

### Issue: "Invalid signature"

**Symptoms**:
```
Error: invalid signature
Error: signature verification failed
```

**Causes**:
- Incorrect digest computation
- Wrong contract address in digest
- Mismatched chain ID
- V value not normalized

**Solutions**:

1. **Use SDK digest computation**:
```go
// ❌ Don't compute manually
digest := keccak256(...)

// ✅ Use SDK methods
digest, err := client.Intent.ComputeDigest(input)
```

2. **Verify digest components**:
```go
input := sdkcrypto.SignedIntentInput{
    IntentID:     intentID,
    SubnetID:     subnetID,
    Requester:    client.Signer.Address(), // Must match signer
    IntentType:   "book_flight",           // Must match exactly
    ParamsHash:   paramsHash,
    Deadline:     deadline,
    PaymentToken: token,
    Amount:       amount,
}

digest, err := client.Intent.ComputeDigest(input)

log.Printf("Digest: %x", digest)
log.Printf("Contract: %s", client.Addresses.IntentManager)
log.Printf("ChainID: %s", client.ChainID)
```

3. **Check signature format**:
```go
func verifySignatureFormat(sig []byte) error {
    if len(sig) != 65 {
        return fmt.Errorf("signature must be 65 bytes, got %d", len(sig))
    }

    v := sig[64]
    if v != 27 && v != 28 && v != 0 && v != 1 {
        return fmt.Errorf("invalid v value: %d", v)
    }

    return nil
}
```

### Issue: "Signer mismatch"

**Symptoms**:
```
Error: signer does not match requester
```

**Causes**:
- Wrong private key used
- Requester field doesn't match signer address

**Solutions**:

1. **Verify signer address**:
```go
signerAddr := client.Signer.Address()
log.Printf("Signer address: %s", signerAddr.Hex())

// Ensure input.Requester matches
if input.Requester != signerAddr {
    return fmt.Errorf("requester mismatch: expected %s, got %s",
        signerAddr.Hex(), input.Requester.Hex())
}
```

2. **Check private key**:
```bash
# Derive address from private key
cast wallet address --private-key $PRIVATE_KEY

# Should match expected address
```

## Gas Problems

### Issue: "Out of gas"

**Symptoms**:
```
Error: out of gas
Error: execution reverted: gas required exceeds allowance
```

**Causes**:
- Gas limit too low
- Complex contract operations
- Batch too large

**Solutions**:

1. **Increase gas multiplier**:
```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        GasLimitMultiplier: ptr(1.5), // Increase from 1.2 to 1.5
    },
}
```

2. **Test with dry-run**:
```go
// Estimate gas first
cfg.Tx.NoSend = ptr(true)
tx, err := client.Intent.SubmitIntent(ctx, params)
if err != nil {
    return err
}

log.Printf("Estimated gas: %d", tx.Gas())

// If too high, split batch or adjust params
```

3. **Check gas ceiling**:
```go
// Temporarily remove ceiling for testing
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        GasCeil: ptr(uint64(0)), // No limit
    },
}

tx, err := client.Intent.SubmitIntent(ctx, params)
log.Printf("Actual gas used: %d", tx.Gas())
```

### Issue: "Gas ceiling exceeded"

**Symptoms**:
```
Error: estimated gas (2500000) exceeds ceiling (2000000)
```

**Causes**:
- Batch too large
- Complex operation
- Gas ceiling too restrictive

**Solutions**:

1. **Split batch**:
```go
func submitInBatches(items []sdk.SignedIntent, batchSize int) error {
    for i := 0; i < len(items); i += batchSize {
        end := i + batchSize
        if end > len(items) {
            end = len(items)
        }

        batch := items[i:end]
        tx, err := client.Intent.SubmitIntentsBySignatures(ctx, sdk.SubmitIntentBatchParams{
            Items: batch,
        })
        if err != nil {
            return err
        }

        log.Printf("Batch %d submitted: %s", i/batchSize+1, tx.Hash().Hex())
    }
    return nil
}
```

2. **Increase ceiling**:
```go
cfg := sdk.Config{
    Tx: &sdk.TxOptions{
        GasCeil: ptr(uint64(5_000_000)), // Increase to 5M
    },
}
```

## Contract Interaction

### Issue: "Contract call reverted"

**Symptoms**:
```
Error: execution reverted
Error: execution reverted: IntentManager: intent expired
```

**Causes**:
- Contract validation failed
- Insufficient approvals
- Contract paused
- Invalid parameters

**Solutions**:

1. **Check contract state**:
```go
// Check if contract is paused
paused, err := client.Intent.Paused(ctx)
if paused {
    return fmt.Errorf("contract is paused")
}

// Check if intent exists
exists, err := client.Intent.IntentExists(ctx, intentID)
if !exists {
    return fmt.Errorf("intent does not exist")
}

// Check if intent expired
expired, err := client.Intent.IsIntentExpired(ctx, intentID)
if expired {
    return fmt.Errorf("intent has expired")
}
```

2. **Check ERC20 approvals**:
```go
// For ERC20 payments, check allowance
import "github.com/ethereum/go-ethereum/accounts/abi/bind"

erc20, err := NewERC20(tokenAddress, client.Backend)
allowance, err := erc20.Allowance(&bind.CallOpts{},
    client.Signer.Address(),
    client.Addresses.IntentManager,
)

if allowance.Cmp(amount) < 0 {
    // Need to approve first
    approveTx, err := erc20.Approve(
        client.TxManager.NewTransactor(),
        client.Addresses.IntentManager,
        amount,
    )
}
```

3. **Decode revert reason**:
```go
// If transaction reverted, get reason
receipt, _ := client.Backend.TransactionReceipt(ctx, tx.Hash())
if receipt.Status == 0 {
    // Call with same params to get revert reason
    _, err := client.Backend.CallContract(ctx, ethereum.CallMsg{
        From: client.Signer.Address(),
        To:   tx.To(),
        Data: tx.Data(),
    }, receipt.BlockNumber)

    log.Printf("Revert reason: %v", err)
}
```

### Issue: "Method not found"

**Symptoms**:
```
Error: method not found
```

**Causes**:
- Wrong contract address
- Contract not deployed
- ABI mismatch

**Solutions**:

1. **Verify contract addresses**:
```go
log.Printf("IntentManager: %s", client.Addresses.IntentManager)
log.Printf("SubnetFactory: %s", client.Addresses.SubnetFactory)

// Check if code exists at address
code, err := client.Backend.CodeAt(ctx, client.Addresses.IntentManager, nil)
if len(code) == 0 {
    return fmt.Errorf("no contract at address %s", client.Addresses.IntentManager)
}
```

2. **Check network**:
```go
chainID, err := client.Backend.ChainID(ctx)
log.Printf("Connected to chain ID: %s", chainID)

// Verify matches expected network
if chainID.Cmp(big.NewInt(84532)) != 0 {
    return fmt.Errorf("wrong network: expected Base Sepolia (84532), got %s", chainID)
}
```

## Configuration Problems

### Issue: "Invalid private key"

**Symptoms**:
```
Error: invalid private key
Error: invalid hex string
```

**Causes**:
- Missing 0x prefix
- Wrong length
- Invalid hex characters

**Solutions**:

1. **Validate format**:
```go
func validatePrivateKey(key string) error {
    // Must start with 0x
    if !strings.HasPrefix(key, "0x") {
        return fmt.Errorf("private key must start with 0x")
    }

    // Remove 0x prefix
    key = key[2:]

    // Must be 64 hex characters
    if len(key) != 64 {
        return fmt.Errorf("private key must be 64 hex characters, got %d", len(key))
    }

    // Must be valid hex
    _, err := hex.DecodeString(key)
    if err != nil {
        return fmt.Errorf("invalid hex characters: %w", err)
    }

    return nil
}
```

2. **Test key**:
```bash
# Derive address from key
cast wallet address --private-key $PRIVATE_KEY

# Should output valid Ethereum address
```

### Issue: "Network mismatch"

**Symptoms**:
```
Error: network mismatch
Error: wrong chain ID
```

**Causes**:
- RPC URL doesn't match PIN_NETWORK
- Wrong contract addresses for network

**Solutions**:

1. **Auto-detect network**:
```go
// Don't specify Network, let SDK auto-detect
cfg := sdk.Config{
    RPCURL:        os.Getenv("PIN_RPC_URL"),
    PrivateKeyHex: os.Getenv("PIN_PRIVATE_KEY"),
    // Network: "",  // Omit to auto-detect
}
```

2. **Verify configuration**:
```go
client, err := sdk.NewClient(ctx, cfg)
if err != nil {
    return err
}

log.Printf("Network: %s", client.Network)
log.Printf("Chain ID: %s", client.ChainID)
log.Printf("Addresses:")
log.Printf("  IntentManager: %s", client.Addresses.IntentManager)
log.Printf("  SubnetFactory: %s", client.Addresses.SubnetFactory)
```

## Debugging Tips

### Enable Verbose Logging

```go
import "log"

// Set verbose logging
log.SetFlags(log.LstdFlags | log.Lshortfile)

// Log all operations
func submitWithLogging(params sdk.SubmitIntentParams) error {
    log.Printf("Submitting intent: %+v", params)

    tx, err := client.Intent.SubmitIntent(ctx, params)
    if err != nil {
        log.Printf("Error: %v", err)
        return err
    }

    log.Printf("Transaction sent: %s", tx.Hash().Hex())
    log.Printf("  Nonce: %d", tx.Nonce())
    log.Printf("  Gas: %d", tx.Gas())
    log.Printf("  Gas Price: %s", tx.GasPrice())

    return nil
}
```

### Inspect Transactions

```bash
# View transaction details
cast tx $TX_HASH --rpc-url $RPC_URL

# View transaction receipt
cast receipt $TX_HASH --rpc-url $RPC_URL

# Decode transaction data
cast 4byte-decode $TX_DATA
```

### Test in Isolation

```go
// Test each component separately
func TestComponents(t *testing.T) {
    // 1. Test RPC connection
    _, err := ethclient.Dial(rpcURL)
    require.NoError(t, err)

    // 2. Test signer
    signer, err := sdk.NewSigner(privateKey)
    require.NoError(t, err)

    // 3. Test contract binding
    intent, err := intentmanager.NewIntentmanager(address, client)
    require.NoError(t, err)

    // 4. Test read operation
    _, err = intent.IntentExists(nil, intentID)
    require.NoError(t, err)
}
```

## Getting Help

If you're still experiencing issues:

1. **Check logs**: Review error messages carefully
2. **Search issues**: Check [GitHub Issues](https://github.com/PIN-AI/intent-protocol-contract-sdk/issues)
3. **Minimal reproduction**: Create minimal code that reproduces the issue
4. **Provide context**: Include SDK version, Go version, network, error messages
5. **Ask for help**: Open a GitHub issue with all relevant details

## Related Documentation

- [Configuration & Environment](config.md) - Configuration guide
- [Transaction Management](txmanager.md) - Transaction handling
- [Signing Specification](signing.md) - Signature details
- [Best Practices](best-practices.md) - Production recommendations
