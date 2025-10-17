# API Reference

**[中文文档](api-reference-zh.md)** | English

Complete API reference for the PIN Intent Protocol SDK.

## Table of Contents

1. [Client](#client)
2. [IntentService](#intentservice)
3. [AssignmentService](#assignmentservice)
4. [ValidationService](#validationservice)
5. [SubnetFactoryService](#subnetfactoryservice)
6. [SubnetService](#subnetservice)
7. [StakingService](#stakingservice)
8. [CheckpointService](#checkpointservice)
9. [TxManager](#txmanager)
10. [Signer](#signer)
11. [Types & Constants](#types--constants)

## Client

The main entry point to the SDK.

### NewClient

```go
func NewClient(ctx context.Context, cfg Config) (*Client, error)
```

Creates a new SDK client instance.

**Parameters**:
- `ctx`: Context for initialization
- `cfg`: Configuration options

**Returns**:
- `*Client`: Initialized client
- `error`: Error if initialization fails

**Example**:
```go
client, err := sdk.NewClient(ctx, sdk.Config{
    RPCURL:        "https://sepolia.base.org",
    PrivateKeyHex: "0x...",
    Network:       "base_sepolia",
})
```

### Client.Close

```go
func (c *Client) Close()
```

Closes the client and releases resources.

### Client.SubnetServiceByID

```go
func (c *Client) SubnetServiceByID(ctx context.Context, subnetID [32]byte) (*SubnetService, error)
```

Gets a SubnetService by querying the subnet address from SubnetFactory.

**Parameters**:
- `ctx`: Context
- `subnetID`: Subnet identifier

**Returns**:
- `*SubnetService`: Subnet service instance
- `error`: Error if subnet not found

### Client.SubnetServiceByAddress

```go
func (c *Client) SubnetServiceByAddress(addr common.Address) (*SubnetService, error)
```

Gets a SubnetService directly from a known subnet contract address.

**Parameters**:
- `addr`: Subnet contract address

**Returns**:
- `*SubnetService`: Subnet service instance
- `error`: Error if binding fails

## IntentService

Service for intent lifecycle management (22 methods).

### Submit Operations

#### SubmitIntent

```go
func (s *IntentService) SubmitIntent(ctx context.Context, params SubmitIntentParams) (*types.Transaction, error)
```

Submits a single intent to the IntentManager contract.

**Parameters**:
- `ctx`: Context
- `params`: Intent submission parameters
  - `IntentID`: Unique intent identifier (bytes32)
  - `SubnetID`: Target subnet identifier (bytes32)
  - `IntentType`: Intent type string
  - `ParamsHash`: Hash of intent parameters (bytes32)
  - `Deadline`: Unix timestamp deadline (*big.Int)
  - `PaymentToken`: Token address (common.Address, 0x0 for ETH)
  - `Amount`: Payment amount (*big.Int)
  - `Value`: ETH value to send (*big.Int, required if PaymentToken==0x0)

**Returns**:
- `*types.Transaction`: Submitted transaction
- `error`: Error if submission fails

**Example**:
```go
tx, err := client.Intent.SubmitIntent(ctx, sdk.SubmitIntentParams{
    IntentID:     sdk.MustBytes32FromHex("0x..."),
    SubnetID:     sdk.MustBytes32FromHex("0x..."),
    IntentType:   "book_flight",
    ParamsHash:   sdk.HashBytes([]byte("{...}")),
    Deadline:     big.NewInt(time.Now().Add(24*time.Hour).Unix()),
    PaymentToken: common.Address{},
    Amount:       big.NewInt(1e16),
    Value:        big.NewInt(1e16),
})
```

#### SubmitIntentsBySignatures

```go
func (s *IntentService) SubmitIntentsBySignatures(ctx context.Context, params SubmitIntentBatchParams) (*types.Transaction, error)
```

Submits multiple intents with signatures (batch submission).

**Parameters**:
- `ctx`: Context
- `params`: Batch submission parameters
  - `Items`: Array of `SignedIntent` structs
  - `TotalEthValue`: Total ETH value (optional, auto-calculated if nil)

**Returns**:
- `*types.Transaction`: Submitted transaction
- `error`: Error if submission fails

### Query Operations

#### IntentExists

```go
func (s *IntentService) IntentExists(ctx context.Context, intentID [32]byte) (bool, error)
```

Checks if an intent exists.

#### GetIntent

```go
func (s *IntentService) GetIntent(ctx context.Context, intentID [32]byte) (IntentInfo, error)
```

Retrieves intent details.

**Returns**:
```go
type IntentInfo struct {
    Requester    common.Address
    SubnetID     [32]byte
    IntentType   string
    ParamsHash   [32]byte
    Deadline     *big.Int
    PaymentToken common.Address
    Amount       *big.Int
    Status       IntentStatus
    // ... additional fields
}
```

#### IsIntentExpired

```go
func (s *IntentService) IsIntentExpired(ctx context.Context, intentID [32]byte) (bool, error)
```

Checks if an intent has expired.

#### GetIntentStatus

```go
func (s *IntentService) GetIntentStatus(ctx context.Context, intentID [32]byte) (IntentStatus, error)
```

Gets the current status of an intent.

**Returns**:
```go
type IntentStatus uint8

const (
    IntentStatusPending IntentStatus = iota
    IntentStatusAssigned
    IntentStatusValidated
    IntentStatusCompleted
    IntentStatusExpired
    IntentStatusFailed
    IntentStatusRefunded
)
```

### Management Operations

#### MarkIntentExpired

```go
func (s *IntentService) MarkIntentExpired(ctx context.Context, intentID [32]byte) (*types.Transaction, error)
```

Marks an expired intent for refund processing.

#### MarkIntentFailed

```go
func (s *IntentService) MarkIntentFailed(ctx context.Context, intentID [32]byte) (*types.Transaction, error)
```

Marks an intent as failed.

#### EmergencyRefundBatch

```go
func (s *IntentService) EmergencyRefundBatch(ctx context.Context, intentIDs [][32]byte) (*types.Transaction, error)
```

Emergency refund for multiple intents (requires GUARDIAN_ROLE).

### Role Management

#### HasRole

```go
func (s *IntentService) HasRole(ctx context.Context, role [32]byte, account common.Address) (bool, error)
```

Checks if an account has a specific role.

#### GrantRole

```go
func (s *IntentService) GrantRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error)
```

Grants a role to an account (requires ADMIN_ROLE).

#### RevokeRole

```go
func (s *IntentService) RevokeRole(ctx context.Context, role [32]byte, account common.Address) (*types.Transaction, error)
```

Revokes a role from an account (requires ADMIN_ROLE).

### Signing Helpers

#### ComputeDigest

```go
func (s *IntentService) ComputeDigest(input sdkcrypto.SignedIntentInput) ([32]byte, error)
```

Computes the EIP-191 digest for an intent.

#### SignDigest

```go
func (s *IntentService) SignDigest(digest [32]byte) ([]byte, error)
```

Signs a digest using the client's signer.

## AssignmentService

Service for matcher assignment operations.

### AssignIntentsBySignatures

```go
func (s *AssignmentService) AssignIntentsBySignatures(ctx context.Context, assignments []SignedAssignment) (*types.Transaction, error)
```

Assigns intents to agents with matcher signatures.

**Parameters**:
- `ctx`: Context
- `assignments`: Array of signed assignments

**Example**:
```go
assignment := sdk.AssignmentData{
    AssignmentID: assignmentID,
    IntentID:     intentID,
    BidID:        bidID,
    Agent:        agentAddress,
    Status:       sdk.AssignmentStatusActive,
    Matcher:      matcherAddress,
}

digest, _ := client.Assignment.ComputeDigest(assignment)
sig, _ := client.Assignment.SignDigest(digest)

tx, err := client.Assignment.AssignIntentsBySignatures(ctx, []sdk.SignedAssignment{
    {Data: assignment, Signature: sig},
})
```

### ComputeDigest

```go
func (s *AssignmentService) ComputeDigest(assignment AssignmentData) ([32]byte, error)
```

Computes the digest for an assignment.

### SignDigest

```go
func (s *AssignmentService) SignDigest(digest [32]byte) ([]byte, error)
```

Signs an assignment digest.

## ValidationService

Service for validator validation operations.

### ValidateIntentsBySignatures

```go
func (s *ValidationService) ValidateIntentsBySignatures(ctx context.Context, bundles []ValidationBundle) (*types.Transaction, error)
```

Submits validation bundles with validator signatures.

**Parameters**:
- `ctx`: Context
- `bundles`: Array of validation bundles

**Example**:
```go
bundle := sdk.ValidationBundle{
    IntentID:     intentID,
    AssignmentID: assignmentID,
    SubnetID:     subnetID,
    Agent:        agentAddress,
    ResultHash:   resultHash,
    ProofHash:    proofHash,
    RootHeight:   big.NewInt(123),
    RootHash:     rootHash,
    Validators:   []common.Address{validatorAddr},
    Signatures:   [][]byte{signature},
}

tx, err := client.Validation.ValidateIntentsBySignatures(ctx, []sdk.ValidationBundle{bundle})
```

### ComputeDigest

```go
func (s *ValidationService) ComputeDigest(bundle ValidationBundle) ([32]byte, error)
```

Computes the digest for a validation bundle.

## SubnetFactoryService

Service for subnet factory operations (30 methods).

### Subnet Creation

#### CreateSubnet

```go
func (s *SubnetFactoryService) CreateSubnet(ctx context.Context, params CreateSubnetParams) (*types.Transaction, error)
```

Creates a new subnet.

**Parameters**:
```go
type CreateSubnetParams struct {
    Name              string
    Description       string
    MinStakeValidator *big.Int
    MinStakeAgent     *big.Int
    MinStakeMatcher   *big.Int
    AutoApprove       bool
    Value             *big.Int // Initial stake
}
```

### Query Operations

#### GetTotalSubnetCount

```go
func (s *SubnetFactoryService) GetTotalSubnetCount(ctx context.Context) (*big.Int, error)
```

Gets the total number of subnets.

#### GetActiveSubnetCount

```go
func (s *SubnetFactoryService) GetActiveSubnetCount(ctx context.Context) (*big.Int, error)
```

Gets the number of active subnets.

#### GetSubnet

```go
func (s *SubnetFactoryService) GetSubnet(ctx context.Context, subnetID [32]byte) (SubnetInfo, error)
```

Gets subnet information by ID.

#### GetSubnetsByStatus

```go
func (s *SubnetFactoryService) GetSubnetsByStatus(ctx context.Context, status SubnetStatus) ([]SubnetInfo, error)
```

Gets all subnets with a specific status (efficient batch query).

**Parameters**:
```go
type SubnetStatus uint8

const (
    SubnetStatusInactive SubnetStatus = iota
    SubnetStatusActive
    SubnetStatusPaused
    SubnetStatusDeprecated
)
```

#### GetSubnetsByOwner

```go
func (s *SubnetFactoryService) GetSubnetsByOwner(ctx context.Context, owner common.Address) ([]SubnetInfo, error)
```

Gets all subnets owned by an address.

#### IsSubnetActive

```go
func (s *SubnetFactoryService) IsSubnetActive(ctx context.Context, subnetID [32]byte) (bool, error)
```

Checks if a subnet is active.

### Management Operations

#### PauseSubnet

```go
func (s *SubnetFactoryService) PauseSubnet(ctx context.Context, subnetID [32]byte) (*types.Transaction, error)
```

Pauses a subnet (requires GOVERNANCE_ROLE).

#### ResumeSubnet

```go
func (s *SubnetFactoryService) ResumeSubnet(ctx context.Context, subnetID [32]byte) (*types.Transaction, error)
```

Resumes a paused subnet.

#### DeprecateSubnet

```go
func (s *SubnetFactoryService) DeprecateSubnet(ctx context.Context, subnetID [32]byte) (*types.Transaction, error)
```

Marks a subnet as deprecated.

## SubnetService

Service for individual subnet operations (27 methods).

### Participant Registration

#### RegisterValidator

```go
func (s *SubnetService) RegisterValidator(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error)
```

Registers as a validator with ETH stake.

**Parameters**:
```go
type RegisterParticipantParams struct {
    Domain      string
    Endpoint    string
    MetadataURI string
    Value       *big.Int // ETH stake amount
}
```

#### RegisterValidatorERC20

```go
func (s *SubnetService) RegisterValidatorERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error)
```

Registers as a validator with ERC20 token stake.

**Parameters**:
```go
type RegisterParticipantERC20Params struct {
    Amount      *big.Int // Token amount
    Domain      string
    Endpoint    string
    MetadataURI string
}
```

#### RegisterAgent

```go
func (s *SubnetService) RegisterAgent(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error)
```

Registers as an agent with ETH stake.

#### RegisterAgentERC20

```go
func (s *SubnetService) RegisterAgentERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error)
```

Registers as an agent with ERC20 token stake.

#### RegisterMatcher

```go
func (s *SubnetService) RegisterMatcher(ctx context.Context, params RegisterParticipantParams) (*types.Transaction, error)
```

Registers as a matcher with ETH stake.

#### RegisterMatcherERC20

```go
func (s *SubnetService) RegisterMatcherERC20(ctx context.Context, params RegisterParticipantERC20Params) (*types.Transaction, error)
```

Registers as a matcher with ERC20 token stake.

### Participant Queries

#### GetParticipantCount

```go
func (s *SubnetService) GetParticipantCount(ctx context.Context, participantType ParticipantType) (*big.Int, error)
```

Gets the count of participants by type.

**Parameters**:
```go
type ParticipantType uint8

const (
    ParticipantValidator ParticipantType = iota
    ParticipantAgent
    ParticipantMatcher
)
```

#### IsActiveParticipant

```go
func (s *SubnetService) IsActiveParticipant(ctx context.Context, addr common.Address, pType ParticipantType) (bool, error)
```

Checks if an address is an active participant of a specific type.

#### ListActiveParticipants

```go
func (s *SubnetService) ListActiveParticipants(ctx context.Context, pType ParticipantType) ([]common.Address, error)
```

Lists all active participants of a specific type.

#### GetParticipantInfo

```go
func (s *SubnetService) GetParticipantInfo(ctx context.Context, addr common.Address) (ParticipantInfo, error)
```

Gets detailed participant information.

### Management Operations

#### ApproveParticipant

```go
func (s *SubnetService) ApproveParticipant(ctx context.Context, participant common.Address, pType ParticipantType) (*types.Transaction, error)
```

Approves a pending participant (requires GOVERNANCE_ROLE).

#### RejectParticipant

```go
func (s *SubnetService) RejectParticipant(ctx context.Context, participant common.Address) (*types.Transaction, error)
```

Rejects a pending participant.

#### SuspendParticipant

```go
func (s *SubnetService) SuspendParticipant(ctx context.Context, participant common.Address) (*types.Transaction, error)
```

Suspends a participant.

#### ResumeParticipant

```go
func (s *SubnetService) ResumeParticipant(ctx context.Context, participant common.Address) (*types.Transaction, error)
```

Resumes a suspended participant.

## StakingService

Service for staking operations (21 methods).

### Staking Operations

#### DepositStake

```go
func (s *StakingService) DepositStake(ctx context.Context, amount *big.Int) (*types.Transaction, error)
```

Deposits additional stake.

#### DepositStakeFor

```go
func (s *StakingService) DepositStakeFor(ctx context.Context, account common.Address, amount *big.Int) (*types.Transaction, error)
```

Deposits stake on behalf of another account.

#### RequestUnstake

```go
func (s *StakingService) RequestUnstake(ctx context.Context, amount *big.Int) (*types.Transaction, error)
```

Requests to unstake an amount (subject to unlock period).

#### Withdraw

```go
func (s *StakingService) Withdraw(ctx context.Context) (*types.Transaction, error)
```

Withdraws unlocked stake.

### Query Operations

#### GetStake

```go
func (s *StakingService) GetStake(ctx context.Context, account common.Address) (*big.Int, error)
```

Gets the staked amount for an account.

#### GetUnstakeRequest

```go
func (s *StakingService) GetUnstakeRequest(ctx context.Context, account common.Address) (UnstakeRequest, error)
```

Gets pending unstake request details.

**Returns**:
```go
type UnstakeRequest struct {
    Amount        *big.Int
    RequestTime   *big.Int
    UnlockTime    *big.Int
}
```

#### GetTotalStaked

```go
func (s *StakingService) GetTotalStaked(ctx context.Context) (*big.Int, error)
```

Gets the total staked amount across all participants.

### Slashing Operations

#### Slash

```go
func (s *StakingService) Slash(ctx context.Context, account common.Address, amount *big.Int) (*types.Transaction, error)
```

Slashes stake from an account (requires SLASHER_ROLE).

## CheckpointService

Service for checkpoint operations (18 methods).

### Submission Operations

#### SubmitCheckpoint

```go
func (s *CheckpointService) SubmitCheckpoint(ctx context.Context, checkpoint CheckpointData) (*types.Transaction, error)
```

Submits a new checkpoint.

**Parameters**:
```go
type CheckpointData struct {
    SubnetID    [32]byte
    Height      *big.Int
    StateRoot   [32]byte
    IntentCount *big.Int
    Timestamp   *big.Int
}
```

### Query Operations

#### GetLatestCheckpoint

```go
func (s *CheckpointService) GetLatestCheckpoint(ctx context.Context, subnetID [32]byte) (CheckpointInfo, error)
```

Gets the latest checkpoint for a subnet.

#### GetCheckpoint

```go
func (s *CheckpointService) GetCheckpoint(ctx context.Context, checkpointID [32]byte) (CheckpointInfo, error)
```

Gets checkpoint details by ID.

#### CanFinalizeCheckpoint

```go
func (s *CheckpointService) CanFinalizeCheckpoint(ctx context.Context, checkpointID [32]byte) (bool, error)
```

Checks if a checkpoint can be finalized.

### Management Operations

#### FinalizeCheckpoint

```go
func (s *CheckpointService) FinalizeCheckpoint(ctx context.Context, checkpointID [32]byte) (*types.Transaction, error)
```

Finalizes a checkpoint (requires GOVERNANCE_ROLE).

#### RevertCheckpoint

```go
func (s *CheckpointService) RevertCheckpoint(ctx context.Context, checkpointID [32]byte) (*types.Transaction, error)
```

Reverts a checkpoint (emergency operation).

## TxManager

Transaction management utilities.

See [Transaction Management](txmanager.md) for detailed documentation.

## Signer

Signature generation interface.

### Interface

```go
type Signer interface {
    Address() common.Address
    SignDigest(digest [32]byte) ([]byte, error)
}
```

### NewSigner

```go
func NewSigner(privateKeyHex string) (Signer, error)
```

Creates a new local signer from a private key.

## Types & Constants

### Common Types

```go
// Intent statuses
type IntentStatus uint8
const (
    IntentStatusPending IntentStatus = iota
    IntentStatusAssigned
    IntentStatusValidated
    IntentStatusCompleted
    IntentStatusExpired
    IntentStatusFailed
    IntentStatusRefunded
)

// Assignment statuses
type AssignmentStatus uint8
const (
    AssignmentStatusActive AssignmentStatus = iota
    AssignmentStatusCompleted
    AssignmentStatusCancelled
)

// Subnet statuses
type SubnetStatus uint8
const (
    SubnetStatusInactive SubnetStatus = iota
    SubnetStatusActive
    SubnetStatusPaused
    SubnetStatusDeprecated
)

// Participant types
type ParticipantType uint8
const (
    ParticipantValidator ParticipantType = iota
    ParticipantAgent
    ParticipantMatcher
)
```

### Utility Functions

#### HashBytes

```go
func HashBytes(data []byte) [32]byte
```

Computes keccak256 hash of data.

#### MustBytes32FromHex

```go
func MustBytes32FromHex(s string) [32]byte
```

Converts hex string to bytes32 (panics on error).

#### GenerateIntentID

```go
func GenerateIntentID() [32]byte
```

Generates a random intent ID.

#### GenerateAssignmentID

```go
func GenerateAssignmentID() [32]byte
```

Generates a random assignment ID.

## Error Handling

All methods return errors following Go conventions. Common error patterns:

```go
// Network errors
if strings.Contains(err.Error(), "connection") {
    // Handle network error
}

// Transaction errors
if strings.Contains(err.Error(), "reverted") {
    // Handle revert
}

// Validation errors
if strings.Contains(err.Error(), "invalid") {
    // Handle validation error
}
```

## Related Documentation

- [Quick Start](quickstart.md) - Usage examples
- [Roles and Workflows](roles-and-workflows.md) - Role-specific examples
- [Configuration](config.md) - Configuration options
- [Best Practices](best-practices.md) - Production recommendations
