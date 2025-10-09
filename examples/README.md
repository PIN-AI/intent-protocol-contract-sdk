# Examples

These entry points demonstrate common RootLayer workflows using the Go SDK. All commands assume you run them from the repository root.

## Shared Setup

1. Copy `.env.example` to `.env` and fill in at least:
   - `PIN_RPC_URL`
   - `PIN_PRIVATE_KEY`
   - `PIN_NETWORK` (`base`, `base_sepolia`, or `local`).
2. Export any runtime overrides in your shell or edit the `.env` file.
3. Execute scripts with `go run examples/<name>/main.go`. When you are ready to ship binaries, prefer `go build -o <name> examples/<name>/`.

## Scripts

### `send_intent`

Submits an intent directly or via signer-provided signatures.

```bash
go run examples/send_intent/main.go
```

Required environment:
- `PIN_RPC_URL`, `PIN_PRIVATE_KEY`, `PIN_NETWORK`

Notable overrides:
- `INTENT_ID` (defaults to random)
- `SUBNET_ID` (auto-selects first active subnet if empty)
- `PARAMS_JSON` or `PARAMS_HASH`
- `AMOUNT_WEI`, `VALUE_WEI`, `PAYMENT_TOKEN`
- `DRY_RUN=true|false`

### `assign_intents`

Signs and submits a matcher assignment batch.

```bash
go run examples/assign_intents/main.go
```

Required environment:
- `INTENT_ID`
- `AGENT_ADDRESS`
- `PIN_RPC_URL`, `PIN_PRIVATE_KEY`, `PIN_NETWORK`

Optional overrides:
- `ASSIGNMENT_ID`, `BID_ID` (random if omitted)
- `MATCHER_ADDRESS` (defaults to local signer)
- `ASSIGNMENT_STATUS` (`active`, `failed`, `unspecified`)
- `DRY_RUN=true|false`

### `validate_intents`

Builds a validation digest and submits validator bundles.

```bash
go run examples/validate_intents/main.go
```

Required environment:
- `INTENT_ID`
- `ASSIGNMENT_ID`
- `SUBNET_ID`
- `AGENT_ADDRESS`
- `RESULT_HASH`
- `PROOF_HASH`
- `ROOT_HASH`
- `PIN_RPC_URL`, `PIN_PRIVATE_KEY`, `PIN_NETWORK`

Optional overrides:
- `ROOT_HEIGHT` (defaults to `0`)
- `VALIDATORS` comma-separated address list (defaults to signer)
- `VALIDATOR_SIGNATURES` comma-separated hex if using external validators
- `DRY_RUN=true|false`

### `register_agent`

Registers the signer as an agent on a subnet.

```bash
go run examples/register_agent/main.go
```

Required environment:
- `PIN_RPC_URL`, `PIN_PRIVATE_KEY`, `PIN_NETWORK`

Optional overrides:
- `SUBNET_ID` (auto-selects active subnet if empty)
- `AGENT_DOMAIN`, `AGENT_ENDPOINT`, `AGENT_METADATA_URI`
- `AGENT_VALUE_WEI` (stake amount, defaults to `0.1` ETH in wei)
- `DRY_RUN=true|false`

### `list_subnets`

Lists subnet IDs and status for the configured network.

```bash
go run examples/list_subnets/main.go
```

Requires only the core environment variables (`PIN_RPC_URL`, `PIN_NETWORK`).

## Tips

- Set `DRY_RUN=false` to broadcast transactions once you are satisfied with the configuration.
- The scripts reuse the SDK client's `.env` loader; `.env` values take precedence over shell defaults.
- Use `go test ./sdk/...` with `PIN_NETWORK=local` to validate SDK logic alongside these flows.
