// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package intentmanager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// DataStructuresIntentInfo is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresIntentInfo struct {
	IntentId              [32]byte
	SubnetId              [32]byte
	Requester             common.Address
	IntentType            string
	CreatedAt             *big.Int
	Deadline              *big.Int
	ParamsHash            [32]byte
	Budget                *big.Int
	PaymentToken          common.Address
	Status                uint8
	TargetAgent           common.Address
	ChallengeEndTimestamp *big.Int
	Disputed              bool
	DisputeDeposit        *big.Int
}

// IIntentManagerAssignmentData is an auto generated low-level Go binding around an user-defined struct.
type IIntentManagerAssignmentData struct {
	AssignmentId [32]byte
	IntentId     [32]byte
	BidId        [32]byte
	Agent        common.Address
	Status       uint8
	Matcher      common.Address
}

// IIntentManagerDirectIntentData is an auto generated low-level Go binding around an user-defined struct.
type IIntentManagerDirectIntentData struct {
	IntentId     [32]byte
	SubnetId     [32]byte
	Requester    common.Address
	IntentType   string
	ParamsHash   [32]byte
	Deadline     *big.Int
	PaymentToken common.Address
	Amount       *big.Int
	TargetAgent  common.Address
}

// IIntentManagerIntentData is an auto generated low-level Go binding around an user-defined struct.
type IIntentManagerIntentData struct {
	IntentId     [32]byte
	SubnetId     [32]byte
	Requester    common.Address
	IntentType   string
	ParamsHash   [32]byte
	Deadline     *big.Int
	PaymentToken common.Address
	Amount       *big.Int
}

// IIntentManagerValidationBatchData is an auto generated low-level Go binding around an user-defined struct.
type IIntentManagerValidationBatchData struct {
	SubnetId   [32]byte
	ItemsHash  [32]byte
	RootHeight uint64
	RootHash   [32]byte
	Items      []IIntentManagerValidationItemData
	Validators []common.Address
	Signatures [][]byte
}

// IIntentManagerValidationBundleData is an auto generated low-level Go binding around an user-defined struct.
type IIntentManagerValidationBundleData struct {
	IntentId     [32]byte
	AssignmentId [32]byte
	SubnetId     [32]byte
	Agent        common.Address
	ResultHash   [32]byte
	ProofHash    [32]byte
	RootHeight   uint64
	RootHash     [32]byte
	Validators   []common.Address
	Signatures   [][]byte
}

// IIntentManagerValidationItemData is an auto generated low-level Go binding around an user-defined struct.
type IIntentManagerValidationItemData struct {
	IntentId     [32]byte
	AssignmentId [32]byte
	Agent        common.Address
	ResultHash   [32]byte
	ProofHash    [32]byte
}

// IntentManagerMetaData contains all meta data concerning the IntentManager contract.
var IntentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssignmentAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssignmentIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengePeriodEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChallengePeriodNotEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateDigest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBudget\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientDeposit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentAlreadyDisputed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentDisputed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentNotPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IntentNotProcessing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDeadline\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIntent\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIntentStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidItemsHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParticipant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubnet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidValidatorCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoActiveValidators\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoClaimableBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatureVerificationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DirectIntentClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target_agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challenge_end_timestamp\",\"type\":\"uint256\"}],\"name\":\"DirectIntentSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit_amount\",\"type\":\"uint256\"}],\"name\":\"DisputeInitiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"agent_correct\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund_amount\",\"type\":\"uint256\"}],\"name\":\"DisputeResolved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"IntentAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"}],\"name\":\"IntentCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund_amount\",\"type\":\"uint256\"}],\"name\":\"IntentExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund_amount\",\"type\":\"uint256\"}],\"name\":\"IntentFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"old_status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"new_status\",\"type\":\"uint8\"}],\"name\":\"IntentStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"IntentSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max_duration\",\"type\":\"uint256\"}],\"name\":\"MaxIntentDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundClaimable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RefundClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ASSIGNMENT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_MAX_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INTENT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LEADER_VALIDATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATION_BATCH_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bid_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.AssignmentStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"}],\"internalType\":\"structIIntentManager.AssignmentData[]\",\"name\":\"assignments\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"assignIntentsBySignatures\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"assignment_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"name\":\"batchGetIntentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target_agent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challenge_end_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disputed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"dispute_deposit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.IntentInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"name\":\"batchProcessExpiredIntents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"claimDirectIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"claimRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"directIntentChallengePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"directIntentExecutionTimeout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeDepositRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"emergencyRefundBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"failIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getClaimableBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"getIntentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"target_agent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"challenge_end_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"disputed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"dispute_deposit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.IntentInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxIntentDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getRequiredValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"required_validators\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subnet_factory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"initiateDispute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"intentExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"isIntentExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"expired\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max_duration\",\"type\":\"uint256\"}],\"name\":\"setMaxIntentDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target_agent\",\"type\":\"address\"}],\"internalType\":\"structIIntentManager.DirectIntentData[]\",\"name\":\"intents\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"submitDirectIntentsBySignatures\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIIntentManager.IntentData[]\",\"name\":\"intents\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"submitIntentsBySignatures\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"result_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"proof_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"root_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structIIntentManager.ValidationBundleData\",\"name\":\"validation\",\"type\":\"tuple\"}],\"name\":\"validateIntentBySignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"items_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"root_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"result_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"proof_hash\",\"type\":\"bytes32\"}],\"internalType\":\"structIIntentManager.ValidationItemData[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structIIntentManager.ValidationBatchData[]\",\"name\":\"batches\",\"type\":\"tuple[]\"}],\"name\":\"validateIntentsBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IntentManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IntentManagerMetaData.ABI instead.
var IntentManagerABI = IntentManagerMetaData.ABI

// IntentManager is an auto generated Go binding around an Ethereum contract.
type IntentManager struct {
	IntentManagerCaller     // Read-only binding to the contract
	IntentManagerTransactor // Write-only binding to the contract
	IntentManagerFilterer   // Log filterer for contract events
}

// IntentManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IntentManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IntentManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IntentManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntentManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IntentManagerSession struct {
	Contract     *IntentManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IntentManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IntentManagerCallerSession struct {
	Contract *IntentManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IntentManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IntentManagerTransactorSession struct {
	Contract     *IntentManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IntentManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IntentManagerRaw struct {
	Contract *IntentManager // Generic contract binding to access the raw methods on
}

// IntentManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IntentManagerCallerRaw struct {
	Contract *IntentManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IntentManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IntentManagerTransactorRaw struct {
	Contract *IntentManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIntentManager creates a new instance of IntentManager, bound to a specific deployed contract.
func NewIntentManager(address common.Address, backend bind.ContractBackend) (*IntentManager, error) {
	contract, err := bindIntentManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IntentManager{IntentManagerCaller: IntentManagerCaller{contract: contract}, IntentManagerTransactor: IntentManagerTransactor{contract: contract}, IntentManagerFilterer: IntentManagerFilterer{contract: contract}}, nil
}

// NewIntentManagerCaller creates a new read-only instance of IntentManager, bound to a specific deployed contract.
func NewIntentManagerCaller(address common.Address, caller bind.ContractCaller) (*IntentManagerCaller, error) {
	contract, err := bindIntentManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IntentManagerCaller{contract: contract}, nil
}

// NewIntentManagerTransactor creates a new write-only instance of IntentManager, bound to a specific deployed contract.
func NewIntentManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IntentManagerTransactor, error) {
	contract, err := bindIntentManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IntentManagerTransactor{contract: contract}, nil
}

// NewIntentManagerFilterer creates a new log filterer instance of IntentManager, bound to a specific deployed contract.
func NewIntentManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IntentManagerFilterer, error) {
	contract, err := bindIntentManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IntentManagerFilterer{contract: contract}, nil
}

// bindIntentManager binds a generic wrapper to an already deployed contract.
func bindIntentManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IntentManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntentManager *IntentManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntentManager.Contract.IntentManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntentManager *IntentManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentManager.Contract.IntentManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntentManager *IntentManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntentManager.Contract.IntentManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntentManager *IntentManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IntentManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntentManager *IntentManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntentManager *IntentManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IntentManager.Contract.contract.Transact(opts, method, params...)
}

// ASSIGNMENTTYPEHASH is a free data retrieval call binding the contract method 0x60d4f5a3.
//
// Solidity: function ASSIGNMENT_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) ASSIGNMENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "ASSIGNMENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ASSIGNMENTTYPEHASH is a free data retrieval call binding the contract method 0x60d4f5a3.
//
// Solidity: function ASSIGNMENT_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerSession) ASSIGNMENTTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.ASSIGNMENTTYPEHASH(&_IntentManager.CallOpts)
}

// ASSIGNMENTTYPEHASH is a free data retrieval call binding the contract method 0x60d4f5a3.
//
// Solidity: function ASSIGNMENT_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) ASSIGNMENTTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.ASSIGNMENTTYPEHASH(&_IntentManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IntentManager.Contract.DEFAULTADMINROLE(&_IntentManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IntentManager.Contract.DEFAULTADMINROLE(&_IntentManager.CallOpts)
}

// DEFAULTMAXDURATION is a free data retrieval call binding the contract method 0xd9423faf.
//
// Solidity: function DEFAULT_MAX_DURATION() view returns(uint256)
func (_IntentManager *IntentManagerCaller) DEFAULTMAXDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "DEFAULT_MAX_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTMAXDURATION is a free data retrieval call binding the contract method 0xd9423faf.
//
// Solidity: function DEFAULT_MAX_DURATION() view returns(uint256)
func (_IntentManager *IntentManagerSession) DEFAULTMAXDURATION() (*big.Int, error) {
	return _IntentManager.Contract.DEFAULTMAXDURATION(&_IntentManager.CallOpts)
}

// DEFAULTMAXDURATION is a free data retrieval call binding the contract method 0xd9423faf.
//
// Solidity: function DEFAULT_MAX_DURATION() view returns(uint256)
func (_IntentManager *IntentManagerCallerSession) DEFAULTMAXDURATION() (*big.Int, error) {
	return _IntentManager.Contract.DEFAULTMAXDURATION(&_IntentManager.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _IntentManager.Contract.GOVERNANCEROLE(&_IntentManager.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _IntentManager.Contract.GOVERNANCEROLE(&_IntentManager.CallOpts)
}

// INTENTTYPEHASH is a free data retrieval call binding the contract method 0xb082a274.
//
// Solidity: function INTENT_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) INTENTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "INTENT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// INTENTTYPEHASH is a free data retrieval call binding the contract method 0xb082a274.
//
// Solidity: function INTENT_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerSession) INTENTTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.INTENTTYPEHASH(&_IntentManager.CallOpts)
}

// INTENTTYPEHASH is a free data retrieval call binding the contract method 0xb082a274.
//
// Solidity: function INTENT_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) INTENTTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.INTENTTYPEHASH(&_IntentManager.CallOpts)
}

// LEADERVALIDATORROLE is a free data retrieval call binding the contract method 0x9b883110.
//
// Solidity: function LEADER_VALIDATOR_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) LEADERVALIDATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "LEADER_VALIDATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LEADERVALIDATORROLE is a free data retrieval call binding the contract method 0x9b883110.
//
// Solidity: function LEADER_VALIDATOR_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerSession) LEADERVALIDATORROLE() ([32]byte, error) {
	return _IntentManager.Contract.LEADERVALIDATORROLE(&_IntentManager.CallOpts)
}

// LEADERVALIDATORROLE is a free data retrieval call binding the contract method 0x9b883110.
//
// Solidity: function LEADER_VALIDATOR_ROLE() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) LEADERVALIDATORROLE() ([32]byte, error) {
	return _IntentManager.Contract.LEADERVALIDATORROLE(&_IntentManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IntentManager *IntentManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IntentManager *IntentManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _IntentManager.Contract.UPGRADEINTERFACEVERSION(&_IntentManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_IntentManager *IntentManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _IntentManager.Contract.UPGRADEINTERFACEVERSION(&_IntentManager.CallOpts)
}

// VALIDATIONBATCHTYPEHASH is a free data retrieval call binding the contract method 0x80f69fe9.
//
// Solidity: function VALIDATION_BATCH_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) VALIDATIONBATCHTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "VALIDATION_BATCH_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATIONBATCHTYPEHASH is a free data retrieval call binding the contract method 0x80f69fe9.
//
// Solidity: function VALIDATION_BATCH_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerSession) VALIDATIONBATCHTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.VALIDATIONBATCHTYPEHASH(&_IntentManager.CallOpts)
}

// VALIDATIONBATCHTYPEHASH is a free data retrieval call binding the contract method 0x80f69fe9.
//
// Solidity: function VALIDATION_BATCH_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) VALIDATIONBATCHTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.VALIDATIONBATCHTYPEHASH(&_IntentManager.CallOpts)
}

// VALIDATIONTYPEHASH is a free data retrieval call binding the contract method 0x2bec6e43.
//
// Solidity: function VALIDATION_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) VALIDATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "VALIDATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VALIDATIONTYPEHASH is a free data retrieval call binding the contract method 0x2bec6e43.
//
// Solidity: function VALIDATION_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerSession) VALIDATIONTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.VALIDATIONTYPEHASH(&_IntentManager.CallOpts)
}

// VALIDATIONTYPEHASH is a free data retrieval call binding the contract method 0x2bec6e43.
//
// Solidity: function VALIDATION_TYPEHASH() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) VALIDATIONTYPEHASH() ([32]byte, error) {
	return _IntentManager.Contract.VALIDATIONTYPEHASH(&_IntentManager.CallOpts)
}

// BatchGetIntentInfo is a free data retrieval call binding the contract method 0xb6d1124e.
//
// Solidity: function batchGetIntentInfo(bytes32[] intent_ids) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8,address,uint256,bool,uint256)[] infos)
func (_IntentManager *IntentManagerCaller) BatchGetIntentInfo(opts *bind.CallOpts, intent_ids [][32]byte) ([]DataStructuresIntentInfo, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "batchGetIntentInfo", intent_ids)

	if err != nil {
		return *new([]DataStructuresIntentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructuresIntentInfo)).(*[]DataStructuresIntentInfo)

	return out0, err

}

// BatchGetIntentInfo is a free data retrieval call binding the contract method 0xb6d1124e.
//
// Solidity: function batchGetIntentInfo(bytes32[] intent_ids) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8,address,uint256,bool,uint256)[] infos)
func (_IntentManager *IntentManagerSession) BatchGetIntentInfo(intent_ids [][32]byte) ([]DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.BatchGetIntentInfo(&_IntentManager.CallOpts, intent_ids)
}

// BatchGetIntentInfo is a free data retrieval call binding the contract method 0xb6d1124e.
//
// Solidity: function batchGetIntentInfo(bytes32[] intent_ids) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8,address,uint256,bool,uint256)[] infos)
func (_IntentManager *IntentManagerCallerSession) BatchGetIntentInfo(intent_ids [][32]byte) ([]DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.BatchGetIntentInfo(&_IntentManager.CallOpts, intent_ids)
}

// DirectIntentChallengePeriod is a free data retrieval call binding the contract method 0xe91da9df.
//
// Solidity: function directIntentChallengePeriod() view returns(uint256)
func (_IntentManager *IntentManagerCaller) DirectIntentChallengePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "directIntentChallengePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DirectIntentChallengePeriod is a free data retrieval call binding the contract method 0xe91da9df.
//
// Solidity: function directIntentChallengePeriod() view returns(uint256)
func (_IntentManager *IntentManagerSession) DirectIntentChallengePeriod() (*big.Int, error) {
	return _IntentManager.Contract.DirectIntentChallengePeriod(&_IntentManager.CallOpts)
}

// DirectIntentChallengePeriod is a free data retrieval call binding the contract method 0xe91da9df.
//
// Solidity: function directIntentChallengePeriod() view returns(uint256)
func (_IntentManager *IntentManagerCallerSession) DirectIntentChallengePeriod() (*big.Int, error) {
	return _IntentManager.Contract.DirectIntentChallengePeriod(&_IntentManager.CallOpts)
}

// DirectIntentExecutionTimeout is a free data retrieval call binding the contract method 0x429c8055.
//
// Solidity: function directIntentExecutionTimeout() view returns(uint256)
func (_IntentManager *IntentManagerCaller) DirectIntentExecutionTimeout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "directIntentExecutionTimeout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DirectIntentExecutionTimeout is a free data retrieval call binding the contract method 0x429c8055.
//
// Solidity: function directIntentExecutionTimeout() view returns(uint256)
func (_IntentManager *IntentManagerSession) DirectIntentExecutionTimeout() (*big.Int, error) {
	return _IntentManager.Contract.DirectIntentExecutionTimeout(&_IntentManager.CallOpts)
}

// DirectIntentExecutionTimeout is a free data retrieval call binding the contract method 0x429c8055.
//
// Solidity: function directIntentExecutionTimeout() view returns(uint256)
func (_IntentManager *IntentManagerCallerSession) DirectIntentExecutionTimeout() (*big.Int, error) {
	return _IntentManager.Contract.DirectIntentExecutionTimeout(&_IntentManager.CallOpts)
}

// DisputeDepositRate is a free data retrieval call binding the contract method 0x13394d9d.
//
// Solidity: function disputeDepositRate() view returns(uint256)
func (_IntentManager *IntentManagerCaller) DisputeDepositRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "disputeDepositRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeDepositRate is a free data retrieval call binding the contract method 0x13394d9d.
//
// Solidity: function disputeDepositRate() view returns(uint256)
func (_IntentManager *IntentManagerSession) DisputeDepositRate() (*big.Int, error) {
	return _IntentManager.Contract.DisputeDepositRate(&_IntentManager.CallOpts)
}

// DisputeDepositRate is a free data retrieval call binding the contract method 0x13394d9d.
//
// Solidity: function disputeDepositRate() view returns(uint256)
func (_IntentManager *IntentManagerCallerSession) DisputeDepositRate() (*big.Int, error) {
	return _IntentManager.Contract.DisputeDepositRate(&_IntentManager.CallOpts)
}

// GetClaimableBalance is a free data retrieval call binding the contract method 0x30ff3043.
//
// Solidity: function getClaimableBalance(address user, address token) view returns(uint256 balance)
func (_IntentManager *IntentManagerCaller) GetClaimableBalance(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getClaimableBalance", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableBalance is a free data retrieval call binding the contract method 0x30ff3043.
//
// Solidity: function getClaimableBalance(address user, address token) view returns(uint256 balance)
func (_IntentManager *IntentManagerSession) GetClaimableBalance(user common.Address, token common.Address) (*big.Int, error) {
	return _IntentManager.Contract.GetClaimableBalance(&_IntentManager.CallOpts, user, token)
}

// GetClaimableBalance is a free data retrieval call binding the contract method 0x30ff3043.
//
// Solidity: function getClaimableBalance(address user, address token) view returns(uint256 balance)
func (_IntentManager *IntentManagerCallerSession) GetClaimableBalance(user common.Address, token common.Address) (*big.Int, error) {
	return _IntentManager.Contract.GetClaimableBalance(&_IntentManager.CallOpts, user, token)
}

// GetIntentInfo is a free data retrieval call binding the contract method 0x283ea6bb.
//
// Solidity: function getIntentInfo(bytes32 intent_id) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8,address,uint256,bool,uint256) info)
func (_IntentManager *IntentManagerCaller) GetIntentInfo(opts *bind.CallOpts, intent_id [32]byte) (DataStructuresIntentInfo, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getIntentInfo", intent_id)

	if err != nil {
		return *new(DataStructuresIntentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresIntentInfo)).(*DataStructuresIntentInfo)

	return out0, err

}

// GetIntentInfo is a free data retrieval call binding the contract method 0x283ea6bb.
//
// Solidity: function getIntentInfo(bytes32 intent_id) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8,address,uint256,bool,uint256) info)
func (_IntentManager *IntentManagerSession) GetIntentInfo(intent_id [32]byte) (DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.GetIntentInfo(&_IntentManager.CallOpts, intent_id)
}

// GetIntentInfo is a free data retrieval call binding the contract method 0x283ea6bb.
//
// Solidity: function getIntentInfo(bytes32 intent_id) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8,address,uint256,bool,uint256) info)
func (_IntentManager *IntentManagerCallerSession) GetIntentInfo(intent_id [32]byte) (DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.GetIntentInfo(&_IntentManager.CallOpts, intent_id)
}

// GetMaxIntentDuration is a free data retrieval call binding the contract method 0x12fcd0ec.
//
// Solidity: function getMaxIntentDuration() view returns(uint256 duration)
func (_IntentManager *IntentManagerCaller) GetMaxIntentDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getMaxIntentDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxIntentDuration is a free data retrieval call binding the contract method 0x12fcd0ec.
//
// Solidity: function getMaxIntentDuration() view returns(uint256 duration)
func (_IntentManager *IntentManagerSession) GetMaxIntentDuration() (*big.Int, error) {
	return _IntentManager.Contract.GetMaxIntentDuration(&_IntentManager.CallOpts)
}

// GetMaxIntentDuration is a free data retrieval call binding the contract method 0x12fcd0ec.
//
// Solidity: function getMaxIntentDuration() view returns(uint256 duration)
func (_IntentManager *IntentManagerCallerSession) GetMaxIntentDuration() (*big.Int, error) {
	return _IntentManager.Contract.GetMaxIntentDuration(&_IntentManager.CallOpts)
}

// GetRequiredValidatorCount is a free data retrieval call binding the contract method 0x9a240f43.
//
// Solidity: function getRequiredValidatorCount(bytes32 subnet_id) view returns(uint256 required_validators)
func (_IntentManager *IntentManagerCaller) GetRequiredValidatorCount(opts *bind.CallOpts, subnet_id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getRequiredValidatorCount", subnet_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredValidatorCount is a free data retrieval call binding the contract method 0x9a240f43.
//
// Solidity: function getRequiredValidatorCount(bytes32 subnet_id) view returns(uint256 required_validators)
func (_IntentManager *IntentManagerSession) GetRequiredValidatorCount(subnet_id [32]byte) (*big.Int, error) {
	return _IntentManager.Contract.GetRequiredValidatorCount(&_IntentManager.CallOpts, subnet_id)
}

// GetRequiredValidatorCount is a free data retrieval call binding the contract method 0x9a240f43.
//
// Solidity: function getRequiredValidatorCount(bytes32 subnet_id) view returns(uint256 required_validators)
func (_IntentManager *IntentManagerCallerSession) GetRequiredValidatorCount(subnet_id [32]byte) (*big.Int, error) {
	return _IntentManager.Contract.GetRequiredValidatorCount(&_IntentManager.CallOpts, subnet_id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IntentManager *IntentManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IntentManager *IntentManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IntentManager.Contract.GetRoleAdmin(&_IntentManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IntentManager.Contract.GetRoleAdmin(&_IntentManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IntentManager *IntentManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IntentManager *IntentManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IntentManager.Contract.HasRole(&_IntentManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IntentManager *IntentManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IntentManager.Contract.HasRole(&_IntentManager.CallOpts, role, account)
}

// IntentExists is a free data retrieval call binding the contract method 0x9e028794.
//
// Solidity: function intentExists(bytes32 intent_id) view returns(bool exists)
func (_IntentManager *IntentManagerCaller) IntentExists(opts *bind.CallOpts, intent_id [32]byte) (bool, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "intentExists", intent_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IntentExists is a free data retrieval call binding the contract method 0x9e028794.
//
// Solidity: function intentExists(bytes32 intent_id) view returns(bool exists)
func (_IntentManager *IntentManagerSession) IntentExists(intent_id [32]byte) (bool, error) {
	return _IntentManager.Contract.IntentExists(&_IntentManager.CallOpts, intent_id)
}

// IntentExists is a free data retrieval call binding the contract method 0x9e028794.
//
// Solidity: function intentExists(bytes32 intent_id) view returns(bool exists)
func (_IntentManager *IntentManagerCallerSession) IntentExists(intent_id [32]byte) (bool, error) {
	return _IntentManager.Contract.IntentExists(&_IntentManager.CallOpts, intent_id)
}

// IsIntentExpired is a free data retrieval call binding the contract method 0xf83e0941.
//
// Solidity: function isIntentExpired(bytes32 intent_id) view returns(bool expired)
func (_IntentManager *IntentManagerCaller) IsIntentExpired(opts *bind.CallOpts, intent_id [32]byte) (bool, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "isIntentExpired", intent_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIntentExpired is a free data retrieval call binding the contract method 0xf83e0941.
//
// Solidity: function isIntentExpired(bytes32 intent_id) view returns(bool expired)
func (_IntentManager *IntentManagerSession) IsIntentExpired(intent_id [32]byte) (bool, error) {
	return _IntentManager.Contract.IsIntentExpired(&_IntentManager.CallOpts, intent_id)
}

// IsIntentExpired is a free data retrieval call binding the contract method 0xf83e0941.
//
// Solidity: function isIntentExpired(bytes32 intent_id) view returns(bool expired)
func (_IntentManager *IntentManagerCallerSession) IsIntentExpired(intent_id [32]byte) (bool, error) {
	return _IntentManager.Contract.IsIntentExpired(&_IntentManager.CallOpts, intent_id)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IntentManager *IntentManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IntentManager *IntentManagerSession) Paused() (bool, error) {
	return _IntentManager.Contract.Paused(&_IntentManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IntentManager *IntentManagerCallerSession) Paused() (bool, error) {
	return _IntentManager.Contract.Paused(&_IntentManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IntentManager *IntentManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IntentManager *IntentManagerSession) ProxiableUUID() ([32]byte, error) {
	return _IntentManager.Contract.ProxiableUUID(&_IntentManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IntentManager *IntentManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IntentManager.Contract.ProxiableUUID(&_IntentManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IntentManager *IntentManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IntentManager *IntentManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IntentManager.Contract.SupportsInterface(&_IntentManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IntentManager *IntentManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IntentManager.Contract.SupportsInterface(&_IntentManager.CallOpts, interfaceId)
}

// AssignIntentsBySignatures is a paid mutator transaction binding the contract method 0xc34f2f22.
//
// Solidity: function assignIntentsBySignatures((bytes32,bytes32,bytes32,address,uint8,address)[] assignments, bytes[] signatures) returns(bytes32[] assignment_ids)
func (_IntentManager *IntentManagerTransactor) AssignIntentsBySignatures(opts *bind.TransactOpts, assignments []IIntentManagerAssignmentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "assignIntentsBySignatures", assignments, signatures)
}

// AssignIntentsBySignatures is a paid mutator transaction binding the contract method 0xc34f2f22.
//
// Solidity: function assignIntentsBySignatures((bytes32,bytes32,bytes32,address,uint8,address)[] assignments, bytes[] signatures) returns(bytes32[] assignment_ids)
func (_IntentManager *IntentManagerSession) AssignIntentsBySignatures(assignments []IIntentManagerAssignmentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.AssignIntentsBySignatures(&_IntentManager.TransactOpts, assignments, signatures)
}

// AssignIntentsBySignatures is a paid mutator transaction binding the contract method 0xc34f2f22.
//
// Solidity: function assignIntentsBySignatures((bytes32,bytes32,bytes32,address,uint8,address)[] assignments, bytes[] signatures) returns(bytes32[] assignment_ids)
func (_IntentManager *IntentManagerTransactorSession) AssignIntentsBySignatures(assignments []IIntentManagerAssignmentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.AssignIntentsBySignatures(&_IntentManager.TransactOpts, assignments, signatures)
}

// BatchProcessExpiredIntents is a paid mutator transaction binding the contract method 0x38a8455d.
//
// Solidity: function batchProcessExpiredIntents(bytes32[] intent_ids) returns()
func (_IntentManager *IntentManagerTransactor) BatchProcessExpiredIntents(opts *bind.TransactOpts, intent_ids [][32]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "batchProcessExpiredIntents", intent_ids)
}

// BatchProcessExpiredIntents is a paid mutator transaction binding the contract method 0x38a8455d.
//
// Solidity: function batchProcessExpiredIntents(bytes32[] intent_ids) returns()
func (_IntentManager *IntentManagerSession) BatchProcessExpiredIntents(intent_ids [][32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.BatchProcessExpiredIntents(&_IntentManager.TransactOpts, intent_ids)
}

// BatchProcessExpiredIntents is a paid mutator transaction binding the contract method 0x38a8455d.
//
// Solidity: function batchProcessExpiredIntents(bytes32[] intent_ids) returns()
func (_IntentManager *IntentManagerTransactorSession) BatchProcessExpiredIntents(intent_ids [][32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.BatchProcessExpiredIntents(&_IntentManager.TransactOpts, intent_ids)
}

// ClaimDirectIntent is a paid mutator transaction binding the contract method 0x6b21448c.
//
// Solidity: function claimDirectIntent(bytes32 intent_id) returns()
func (_IntentManager *IntentManagerTransactor) ClaimDirectIntent(opts *bind.TransactOpts, intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "claimDirectIntent", intent_id)
}

// ClaimDirectIntent is a paid mutator transaction binding the contract method 0x6b21448c.
//
// Solidity: function claimDirectIntent(bytes32 intent_id) returns()
func (_IntentManager *IntentManagerSession) ClaimDirectIntent(intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.ClaimDirectIntent(&_IntentManager.TransactOpts, intent_id)
}

// ClaimDirectIntent is a paid mutator transaction binding the contract method 0x6b21448c.
//
// Solidity: function claimDirectIntent(bytes32 intent_id) returns()
func (_IntentManager *IntentManagerTransactorSession) ClaimDirectIntent(intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.ClaimDirectIntent(&_IntentManager.TransactOpts, intent_id)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address token) returns()
func (_IntentManager *IntentManagerTransactor) ClaimRefund(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "claimRefund", token)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address token) returns()
func (_IntentManager *IntentManagerSession) ClaimRefund(token common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.ClaimRefund(&_IntentManager.TransactOpts, token)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0xbffa55d5.
//
// Solidity: function claimRefund(address token) returns()
func (_IntentManager *IntentManagerTransactorSession) ClaimRefund(token common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.ClaimRefund(&_IntentManager.TransactOpts, token)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_IntentManager *IntentManagerTransactor) EmergencyPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "emergencyPause")
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_IntentManager *IntentManagerSession) EmergencyPause() (*types.Transaction, error) {
	return _IntentManager.Contract.EmergencyPause(&_IntentManager.TransactOpts)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_IntentManager *IntentManagerTransactorSession) EmergencyPause() (*types.Transaction, error) {
	return _IntentManager.Contract.EmergencyPause(&_IntentManager.TransactOpts)
}

// EmergencyRefundBatch is a paid mutator transaction binding the contract method 0xa0f548c3.
//
// Solidity: function emergencyRefundBatch(bytes32[] intent_ids, string reason) returns()
func (_IntentManager *IntentManagerTransactor) EmergencyRefundBatch(opts *bind.TransactOpts, intent_ids [][32]byte, reason string) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "emergencyRefundBatch", intent_ids, reason)
}

// EmergencyRefundBatch is a paid mutator transaction binding the contract method 0xa0f548c3.
//
// Solidity: function emergencyRefundBatch(bytes32[] intent_ids, string reason) returns()
func (_IntentManager *IntentManagerSession) EmergencyRefundBatch(intent_ids [][32]byte, reason string) (*types.Transaction, error) {
	return _IntentManager.Contract.EmergencyRefundBatch(&_IntentManager.TransactOpts, intent_ids, reason)
}

// EmergencyRefundBatch is a paid mutator transaction binding the contract method 0xa0f548c3.
//
// Solidity: function emergencyRefundBatch(bytes32[] intent_ids, string reason) returns()
func (_IntentManager *IntentManagerTransactorSession) EmergencyRefundBatch(intent_ids [][32]byte, reason string) (*types.Transaction, error) {
	return _IntentManager.Contract.EmergencyRefundBatch(&_IntentManager.TransactOpts, intent_ids, reason)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_IntentManager *IntentManagerTransactor) EmergencyUnpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "emergencyUnpause")
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_IntentManager *IntentManagerSession) EmergencyUnpause() (*types.Transaction, error) {
	return _IntentManager.Contract.EmergencyUnpause(&_IntentManager.TransactOpts)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_IntentManager *IntentManagerTransactorSession) EmergencyUnpause() (*types.Transaction, error) {
	return _IntentManager.Contract.EmergencyUnpause(&_IntentManager.TransactOpts)
}

// FailIntent is a paid mutator transaction binding the contract method 0xff4d0173.
//
// Solidity: function failIntent(bytes32 intent_id, string reason, address[] validators, bytes[] signatures) returns()
func (_IntentManager *IntentManagerTransactor) FailIntent(opts *bind.TransactOpts, intent_id [32]byte, reason string, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "failIntent", intent_id, reason, validators, signatures)
}

// FailIntent is a paid mutator transaction binding the contract method 0xff4d0173.
//
// Solidity: function failIntent(bytes32 intent_id, string reason, address[] validators, bytes[] signatures) returns()
func (_IntentManager *IntentManagerSession) FailIntent(intent_id [32]byte, reason string, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.FailIntent(&_IntentManager.TransactOpts, intent_id, reason, validators, signatures)
}

// FailIntent is a paid mutator transaction binding the contract method 0xff4d0173.
//
// Solidity: function failIntent(bytes32 intent_id, string reason, address[] validators, bytes[] signatures) returns()
func (_IntentManager *IntentManagerTransactorSession) FailIntent(intent_id [32]byte, reason string, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.FailIntent(&_IntentManager.TransactOpts, intent_id, reason, validators, signatures)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IntentManager *IntentManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IntentManager *IntentManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.GrantRole(&_IntentManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IntentManager *IntentManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.GrantRole(&_IntentManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address subnet_factory) returns()
func (_IntentManager *IntentManagerTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, subnet_factory common.Address) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "initialize", admin, subnet_factory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address subnet_factory) returns()
func (_IntentManager *IntentManagerSession) Initialize(admin common.Address, subnet_factory common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.Initialize(&_IntentManager.TransactOpts, admin, subnet_factory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address subnet_factory) returns()
func (_IntentManager *IntentManagerTransactorSession) Initialize(admin common.Address, subnet_factory common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.Initialize(&_IntentManager.TransactOpts, admin, subnet_factory)
}

// InitiateDispute is a paid mutator transaction binding the contract method 0x7b2e941e.
//
// Solidity: function initiateDispute(bytes32 intent_id) payable returns()
func (_IntentManager *IntentManagerTransactor) InitiateDispute(opts *bind.TransactOpts, intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "initiateDispute", intent_id)
}

// InitiateDispute is a paid mutator transaction binding the contract method 0x7b2e941e.
//
// Solidity: function initiateDispute(bytes32 intent_id) payable returns()
func (_IntentManager *IntentManagerSession) InitiateDispute(intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.InitiateDispute(&_IntentManager.TransactOpts, intent_id)
}

// InitiateDispute is a paid mutator transaction binding the contract method 0x7b2e941e.
//
// Solidity: function initiateDispute(bytes32 intent_id) payable returns()
func (_IntentManager *IntentManagerTransactorSession) InitiateDispute(intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.InitiateDispute(&_IntentManager.TransactOpts, intent_id)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IntentManager *IntentManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IntentManager *IntentManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.RenounceRole(&_IntentManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_IntentManager *IntentManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.RenounceRole(&_IntentManager.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IntentManager *IntentManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IntentManager *IntentManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.RevokeRole(&_IntentManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IntentManager *IntentManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IntentManager.Contract.RevokeRole(&_IntentManager.TransactOpts, role, account)
}

// SetMaxIntentDuration is a paid mutator transaction binding the contract method 0x7ba19837.
//
// Solidity: function setMaxIntentDuration(uint256 max_duration) returns()
func (_IntentManager *IntentManagerTransactor) SetMaxIntentDuration(opts *bind.TransactOpts, max_duration *big.Int) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "setMaxIntentDuration", max_duration)
}

// SetMaxIntentDuration is a paid mutator transaction binding the contract method 0x7ba19837.
//
// Solidity: function setMaxIntentDuration(uint256 max_duration) returns()
func (_IntentManager *IntentManagerSession) SetMaxIntentDuration(max_duration *big.Int) (*types.Transaction, error) {
	return _IntentManager.Contract.SetMaxIntentDuration(&_IntentManager.TransactOpts, max_duration)
}

// SetMaxIntentDuration is a paid mutator transaction binding the contract method 0x7ba19837.
//
// Solidity: function setMaxIntentDuration(uint256 max_duration) returns()
func (_IntentManager *IntentManagerTransactorSession) SetMaxIntentDuration(max_duration *big.Int) (*types.Transaction, error) {
	return _IntentManager.Contract.SetMaxIntentDuration(&_IntentManager.TransactOpts, max_duration)
}

// SubmitDirectIntentsBySignatures is a paid mutator transaction binding the contract method 0xa72c0c2a.
//
// Solidity: function submitDirectIntentsBySignatures((bytes32,bytes32,address,string,bytes32,uint256,address,uint256,address)[] intents, bytes[] signatures) payable returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerTransactor) SubmitDirectIntentsBySignatures(opts *bind.TransactOpts, intents []IIntentManagerDirectIntentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "submitDirectIntentsBySignatures", intents, signatures)
}

// SubmitDirectIntentsBySignatures is a paid mutator transaction binding the contract method 0xa72c0c2a.
//
// Solidity: function submitDirectIntentsBySignatures((bytes32,bytes32,address,string,bytes32,uint256,address,uint256,address)[] intents, bytes[] signatures) payable returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerSession) SubmitDirectIntentsBySignatures(intents []IIntentManagerDirectIntentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.SubmitDirectIntentsBySignatures(&_IntentManager.TransactOpts, intents, signatures)
}

// SubmitDirectIntentsBySignatures is a paid mutator transaction binding the contract method 0xa72c0c2a.
//
// Solidity: function submitDirectIntentsBySignatures((bytes32,bytes32,address,string,bytes32,uint256,address,uint256,address)[] intents, bytes[] signatures) payable returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerTransactorSession) SubmitDirectIntentsBySignatures(intents []IIntentManagerDirectIntentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.SubmitDirectIntentsBySignatures(&_IntentManager.TransactOpts, intents, signatures)
}

// SubmitIntentsBySignatures is a paid mutator transaction binding the contract method 0x931280e3.
//
// Solidity: function submitIntentsBySignatures((bytes32,bytes32,address,string,bytes32,uint256,address,uint256)[] intents, bytes[] signatures) payable returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerTransactor) SubmitIntentsBySignatures(opts *bind.TransactOpts, intents []IIntentManagerIntentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "submitIntentsBySignatures", intents, signatures)
}

// SubmitIntentsBySignatures is a paid mutator transaction binding the contract method 0x931280e3.
//
// Solidity: function submitIntentsBySignatures((bytes32,bytes32,address,string,bytes32,uint256,address,uint256)[] intents, bytes[] signatures) payable returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerSession) SubmitIntentsBySignatures(intents []IIntentManagerIntentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.SubmitIntentsBySignatures(&_IntentManager.TransactOpts, intents, signatures)
}

// SubmitIntentsBySignatures is a paid mutator transaction binding the contract method 0x931280e3.
//
// Solidity: function submitIntentsBySignatures((bytes32,bytes32,address,string,bytes32,uint256,address,uint256)[] intents, bytes[] signatures) payable returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerTransactorSession) SubmitIntentsBySignatures(intents []IIntentManagerIntentData, signatures [][]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.SubmitIntentsBySignatures(&_IntentManager.TransactOpts, intents, signatures)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IntentManager *IntentManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IntentManager *IntentManagerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IntentManager.Contract.UpgradeToAndCall(&_IntentManager.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_IntentManager *IntentManagerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _IntentManager.Contract.UpgradeToAndCall(&_IntentManager.TransactOpts, newImplementation, data)
}

// ValidateIntentBySignature is a paid mutator transaction binding the contract method 0x14f07288.
//
// Solidity: function validateIntentBySignature((bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address[],bytes[]) validation) returns()
func (_IntentManager *IntentManagerTransactor) ValidateIntentBySignature(opts *bind.TransactOpts, validation IIntentManagerValidationBundleData) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "validateIntentBySignature", validation)
}

// ValidateIntentBySignature is a paid mutator transaction binding the contract method 0x14f07288.
//
// Solidity: function validateIntentBySignature((bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address[],bytes[]) validation) returns()
func (_IntentManager *IntentManagerSession) ValidateIntentBySignature(validation IIntentManagerValidationBundleData) (*types.Transaction, error) {
	return _IntentManager.Contract.ValidateIntentBySignature(&_IntentManager.TransactOpts, validation)
}

// ValidateIntentBySignature is a paid mutator transaction binding the contract method 0x14f07288.
//
// Solidity: function validateIntentBySignature((bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address[],bytes[]) validation) returns()
func (_IntentManager *IntentManagerTransactorSession) ValidateIntentBySignature(validation IIntentManagerValidationBundleData) (*types.Transaction, error) {
	return _IntentManager.Contract.ValidateIntentBySignature(&_IntentManager.TransactOpts, validation)
}

// ValidateIntentsBySignatures is a paid mutator transaction binding the contract method 0xbdf22507.
//
// Solidity: function validateIntentsBySignatures((bytes32,bytes32,uint64,bytes32,(bytes32,bytes32,address,bytes32,bytes32)[],address[],bytes[])[] batches) returns()
func (_IntentManager *IntentManagerTransactor) ValidateIntentsBySignatures(opts *bind.TransactOpts, batches []IIntentManagerValidationBatchData) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "validateIntentsBySignatures", batches)
}

// ValidateIntentsBySignatures is a paid mutator transaction binding the contract method 0xbdf22507.
//
// Solidity: function validateIntentsBySignatures((bytes32,bytes32,uint64,bytes32,(bytes32,bytes32,address,bytes32,bytes32)[],address[],bytes[])[] batches) returns()
func (_IntentManager *IntentManagerSession) ValidateIntentsBySignatures(batches []IIntentManagerValidationBatchData) (*types.Transaction, error) {
	return _IntentManager.Contract.ValidateIntentsBySignatures(&_IntentManager.TransactOpts, batches)
}

// ValidateIntentsBySignatures is a paid mutator transaction binding the contract method 0xbdf22507.
//
// Solidity: function validateIntentsBySignatures((bytes32,bytes32,uint64,bytes32,(bytes32,bytes32,address,bytes32,bytes32)[],address[],bytes[])[] batches) returns()
func (_IntentManager *IntentManagerTransactorSession) ValidateIntentsBySignatures(batches []IIntentManagerValidationBatchData) (*types.Transaction, error) {
	return _IntentManager.Contract.ValidateIntentsBySignatures(&_IntentManager.TransactOpts, batches)
}

// IntentManagerDirectIntentClaimedIterator is returned from FilterDirectIntentClaimed and is used to iterate over the raw logs and unpacked data for DirectIntentClaimed events raised by the IntentManager contract.
type IntentManagerDirectIntentClaimedIterator struct {
	Event *IntentManagerDirectIntentClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerDirectIntentClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerDirectIntentClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerDirectIntentClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerDirectIntentClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerDirectIntentClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerDirectIntentClaimed represents a DirectIntentClaimed event raised by the IntentManager contract.
type IntentManagerDirectIntentClaimed struct {
	IntentId [32]byte
	Agent    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDirectIntentClaimed is a free log retrieval operation binding the contract event 0x7b2c4cb5d82a9999a13da98b25d2585ce0d630197b696e0cb9c2c25a13640cc4.
//
// Solidity: event DirectIntentClaimed(bytes32 indexed intent_id, address indexed agent, uint256 amount)
func (_IntentManager *IntentManagerFilterer) FilterDirectIntentClaimed(opts *bind.FilterOpts, intent_id [][32]byte, agent []common.Address) (*IntentManagerDirectIntentClaimedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "DirectIntentClaimed", intent_idRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerDirectIntentClaimedIterator{contract: _IntentManager.contract, event: "DirectIntentClaimed", logs: logs, sub: sub}, nil
}

// WatchDirectIntentClaimed is a free log subscription operation binding the contract event 0x7b2c4cb5d82a9999a13da98b25d2585ce0d630197b696e0cb9c2c25a13640cc4.
//
// Solidity: event DirectIntentClaimed(bytes32 indexed intent_id, address indexed agent, uint256 amount)
func (_IntentManager *IntentManagerFilterer) WatchDirectIntentClaimed(opts *bind.WatchOpts, sink chan<- *IntentManagerDirectIntentClaimed, intent_id [][32]byte, agent []common.Address) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "DirectIntentClaimed", intent_idRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerDirectIntentClaimed)
				if err := _IntentManager.contract.UnpackLog(event, "DirectIntentClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDirectIntentClaimed is a log parse operation binding the contract event 0x7b2c4cb5d82a9999a13da98b25d2585ce0d630197b696e0cb9c2c25a13640cc4.
//
// Solidity: event DirectIntentClaimed(bytes32 indexed intent_id, address indexed agent, uint256 amount)
func (_IntentManager *IntentManagerFilterer) ParseDirectIntentClaimed(log types.Log) (*IntentManagerDirectIntentClaimed, error) {
	event := new(IntentManagerDirectIntentClaimed)
	if err := _IntentManager.contract.UnpackLog(event, "DirectIntentClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerDirectIntentSubmittedIterator is returned from FilterDirectIntentSubmitted and is used to iterate over the raw logs and unpacked data for DirectIntentSubmitted events raised by the IntentManager contract.
type IntentManagerDirectIntentSubmittedIterator struct {
	Event *IntentManagerDirectIntentSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerDirectIntentSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerDirectIntentSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerDirectIntentSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerDirectIntentSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerDirectIntentSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerDirectIntentSubmitted represents a DirectIntentSubmitted event raised by the IntentManager contract.
type IntentManagerDirectIntentSubmitted struct {
	IntentId              [32]byte
	Requester             common.Address
	TargetAgent           common.Address
	Budget                *big.Int
	ChallengeEndTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterDirectIntentSubmitted is a free log retrieval operation binding the contract event 0xfd4ccc27b7fb58fc3ac1e19232e06334c737845c9d15c24442ed56a8ea783ec7.
//
// Solidity: event DirectIntentSubmitted(bytes32 indexed intent_id, address indexed requester, address indexed target_agent, uint256 budget, uint256 challenge_end_timestamp)
func (_IntentManager *IntentManagerFilterer) FilterDirectIntentSubmitted(opts *bind.FilterOpts, intent_id [][32]byte, requester []common.Address, target_agent []common.Address) (*IntentManagerDirectIntentSubmittedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var target_agentRule []interface{}
	for _, target_agentItem := range target_agent {
		target_agentRule = append(target_agentRule, target_agentItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "DirectIntentSubmitted", intent_idRule, requesterRule, target_agentRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerDirectIntentSubmittedIterator{contract: _IntentManager.contract, event: "DirectIntentSubmitted", logs: logs, sub: sub}, nil
}

// WatchDirectIntentSubmitted is a free log subscription operation binding the contract event 0xfd4ccc27b7fb58fc3ac1e19232e06334c737845c9d15c24442ed56a8ea783ec7.
//
// Solidity: event DirectIntentSubmitted(bytes32 indexed intent_id, address indexed requester, address indexed target_agent, uint256 budget, uint256 challenge_end_timestamp)
func (_IntentManager *IntentManagerFilterer) WatchDirectIntentSubmitted(opts *bind.WatchOpts, sink chan<- *IntentManagerDirectIntentSubmitted, intent_id [][32]byte, requester []common.Address, target_agent []common.Address) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var target_agentRule []interface{}
	for _, target_agentItem := range target_agent {
		target_agentRule = append(target_agentRule, target_agentItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "DirectIntentSubmitted", intent_idRule, requesterRule, target_agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerDirectIntentSubmitted)
				if err := _IntentManager.contract.UnpackLog(event, "DirectIntentSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDirectIntentSubmitted is a log parse operation binding the contract event 0xfd4ccc27b7fb58fc3ac1e19232e06334c737845c9d15c24442ed56a8ea783ec7.
//
// Solidity: event DirectIntentSubmitted(bytes32 indexed intent_id, address indexed requester, address indexed target_agent, uint256 budget, uint256 challenge_end_timestamp)
func (_IntentManager *IntentManagerFilterer) ParseDirectIntentSubmitted(log types.Log) (*IntentManagerDirectIntentSubmitted, error) {
	event := new(IntentManagerDirectIntentSubmitted)
	if err := _IntentManager.contract.UnpackLog(event, "DirectIntentSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerDisputeInitiatedIterator is returned from FilterDisputeInitiated and is used to iterate over the raw logs and unpacked data for DisputeInitiated events raised by the IntentManager contract.
type IntentManagerDisputeInitiatedIterator struct {
	Event *IntentManagerDisputeInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerDisputeInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerDisputeInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerDisputeInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerDisputeInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerDisputeInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerDisputeInitiated represents a DisputeInitiated event raised by the IntentManager contract.
type IntentManagerDisputeInitiated struct {
	IntentId      [32]byte
	Requester     common.Address
	DepositAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDisputeInitiated is a free log retrieval operation binding the contract event 0x7a84f95f99d156e1e59639d2b45873711c36b1dd9ded51cb8d6b4e6accb9fb06.
//
// Solidity: event DisputeInitiated(bytes32 indexed intent_id, address indexed requester, uint256 deposit_amount)
func (_IntentManager *IntentManagerFilterer) FilterDisputeInitiated(opts *bind.FilterOpts, intent_id [][32]byte, requester []common.Address) (*IntentManagerDisputeInitiatedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "DisputeInitiated", intent_idRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerDisputeInitiatedIterator{contract: _IntentManager.contract, event: "DisputeInitiated", logs: logs, sub: sub}, nil
}

// WatchDisputeInitiated is a free log subscription operation binding the contract event 0x7a84f95f99d156e1e59639d2b45873711c36b1dd9ded51cb8d6b4e6accb9fb06.
//
// Solidity: event DisputeInitiated(bytes32 indexed intent_id, address indexed requester, uint256 deposit_amount)
func (_IntentManager *IntentManagerFilterer) WatchDisputeInitiated(opts *bind.WatchOpts, sink chan<- *IntentManagerDisputeInitiated, intent_id [][32]byte, requester []common.Address) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "DisputeInitiated", intent_idRule, requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerDisputeInitiated)
				if err := _IntentManager.contract.UnpackLog(event, "DisputeInitiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeInitiated is a log parse operation binding the contract event 0x7a84f95f99d156e1e59639d2b45873711c36b1dd9ded51cb8d6b4e6accb9fb06.
//
// Solidity: event DisputeInitiated(bytes32 indexed intent_id, address indexed requester, uint256 deposit_amount)
func (_IntentManager *IntentManagerFilterer) ParseDisputeInitiated(log types.Log) (*IntentManagerDisputeInitiated, error) {
	event := new(IntentManagerDisputeInitiated)
	if err := _IntentManager.contract.UnpackLog(event, "DisputeInitiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerDisputeResolvedIterator is returned from FilterDisputeResolved and is used to iterate over the raw logs and unpacked data for DisputeResolved events raised by the IntentManager contract.
type IntentManagerDisputeResolvedIterator struct {
	Event *IntentManagerDisputeResolved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerDisputeResolvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerDisputeResolved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerDisputeResolved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerDisputeResolvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerDisputeResolvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerDisputeResolved represents a DisputeResolved event raised by the IntentManager contract.
type IntentManagerDisputeResolved struct {
	IntentId     [32]byte
	AgentCorrect bool
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDisputeResolved is a free log retrieval operation binding the contract event 0x5a5ac8b853b549bf2485e24863e0628003e891a2b72367623dd0b1c324eba8f9.
//
// Solidity: event DisputeResolved(bytes32 indexed intent_id, bool agent_correct, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) FilterDisputeResolved(opts *bind.FilterOpts, intent_id [][32]byte) (*IntentManagerDisputeResolvedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "DisputeResolved", intent_idRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerDisputeResolvedIterator{contract: _IntentManager.contract, event: "DisputeResolved", logs: logs, sub: sub}, nil
}

// WatchDisputeResolved is a free log subscription operation binding the contract event 0x5a5ac8b853b549bf2485e24863e0628003e891a2b72367623dd0b1c324eba8f9.
//
// Solidity: event DisputeResolved(bytes32 indexed intent_id, bool agent_correct, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) WatchDisputeResolved(opts *bind.WatchOpts, sink chan<- *IntentManagerDisputeResolved, intent_id [][32]byte) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "DisputeResolved", intent_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerDisputeResolved)
				if err := _IntentManager.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDisputeResolved is a log parse operation binding the contract event 0x5a5ac8b853b549bf2485e24863e0628003e891a2b72367623dd0b1c324eba8f9.
//
// Solidity: event DisputeResolved(bytes32 indexed intent_id, bool agent_correct, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) ParseDisputeResolved(log types.Log) (*IntentManagerDisputeResolved, error) {
	event := new(IntentManagerDisputeResolved)
	if err := _IntentManager.contract.UnpackLog(event, "DisputeResolved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the IntentManager contract.
type IntentManagerInitializedIterator struct {
	Event *IntentManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerInitialized represents a Initialized event raised by the IntentManager contract.
type IntentManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IntentManager *IntentManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*IntentManagerInitializedIterator, error) {

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IntentManagerInitializedIterator{contract: _IntentManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IntentManager *IntentManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IntentManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerInitialized)
				if err := _IntentManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_IntentManager *IntentManagerFilterer) ParseInitialized(log types.Log) (*IntentManagerInitialized, error) {
	event := new(IntentManagerInitialized)
	if err := _IntentManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerIntentAssignedIterator is returned from FilterIntentAssigned and is used to iterate over the raw logs and unpacked data for IntentAssigned events raised by the IntentManager contract.
type IntentManagerIntentAssignedIterator struct {
	Event *IntentManagerIntentAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerIntentAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerIntentAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerIntentAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerIntentAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerIntentAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerIntentAssigned represents a IntentAssigned event raised by the IntentManager contract.
type IntentManagerIntentAssigned struct {
	AssignmentId [32]byte
	IntentId     [32]byte
	Matcher      common.Address
	Agent        common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIntentAssigned is a free log retrieval operation binding the contract event 0xc43578d43d59b01fa3097859d9bb1cbbf2725e102df510e419b9794554894665.
//
// Solidity: event IntentAssigned(bytes32 indexed assignment_id, bytes32 indexed intent_id, address indexed matcher, address agent)
func (_IntentManager *IntentManagerFilterer) FilterIntentAssigned(opts *bind.FilterOpts, assignment_id [][32]byte, intent_id [][32]byte, matcher []common.Address) (*IntentManagerIntentAssignedIterator, error) {

	var assignment_idRule []interface{}
	for _, assignment_idItem := range assignment_id {
		assignment_idRule = append(assignment_idRule, assignment_idItem)
	}
	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var matcherRule []interface{}
	for _, matcherItem := range matcher {
		matcherRule = append(matcherRule, matcherItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "IntentAssigned", assignment_idRule, intent_idRule, matcherRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerIntentAssignedIterator{contract: _IntentManager.contract, event: "IntentAssigned", logs: logs, sub: sub}, nil
}

// WatchIntentAssigned is a free log subscription operation binding the contract event 0xc43578d43d59b01fa3097859d9bb1cbbf2725e102df510e419b9794554894665.
//
// Solidity: event IntentAssigned(bytes32 indexed assignment_id, bytes32 indexed intent_id, address indexed matcher, address agent)
func (_IntentManager *IntentManagerFilterer) WatchIntentAssigned(opts *bind.WatchOpts, sink chan<- *IntentManagerIntentAssigned, assignment_id [][32]byte, intent_id [][32]byte, matcher []common.Address) (event.Subscription, error) {

	var assignment_idRule []interface{}
	for _, assignment_idItem := range assignment_id {
		assignment_idRule = append(assignment_idRule, assignment_idItem)
	}
	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var matcherRule []interface{}
	for _, matcherItem := range matcher {
		matcherRule = append(matcherRule, matcherItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "IntentAssigned", assignment_idRule, intent_idRule, matcherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerIntentAssigned)
				if err := _IntentManager.contract.UnpackLog(event, "IntentAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIntentAssigned is a log parse operation binding the contract event 0xc43578d43d59b01fa3097859d9bb1cbbf2725e102df510e419b9794554894665.
//
// Solidity: event IntentAssigned(bytes32 indexed assignment_id, bytes32 indexed intent_id, address indexed matcher, address agent)
func (_IntentManager *IntentManagerFilterer) ParseIntentAssigned(log types.Log) (*IntentManagerIntentAssigned, error) {
	event := new(IntentManagerIntentAssigned)
	if err := _IntentManager.contract.UnpackLog(event, "IntentAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerIntentCompletedIterator is returned from FilterIntentCompleted and is used to iterate over the raw logs and unpacked data for IntentCompleted events raised by the IntentManager contract.
type IntentManagerIntentCompletedIterator struct {
	Event *IntentManagerIntentCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerIntentCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerIntentCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerIntentCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerIntentCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerIntentCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerIntentCompleted represents a IntentCompleted event raised by the IntentManager contract.
type IntentManagerIntentCompleted struct {
	IntentId   [32]byte
	ResultHash [32]byte
	Budget     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterIntentCompleted is a free log retrieval operation binding the contract event 0x6337eb6af1128c06eedb6c409b7873f58f3afdf8ee2442615a5aeb0c0771fa72.
//
// Solidity: event IntentCompleted(bytes32 indexed intent_id, bytes32 result_hash, uint256 budget)
func (_IntentManager *IntentManagerFilterer) FilterIntentCompleted(opts *bind.FilterOpts, intent_id [][32]byte) (*IntentManagerIntentCompletedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "IntentCompleted", intent_idRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerIntentCompletedIterator{contract: _IntentManager.contract, event: "IntentCompleted", logs: logs, sub: sub}, nil
}

// WatchIntentCompleted is a free log subscription operation binding the contract event 0x6337eb6af1128c06eedb6c409b7873f58f3afdf8ee2442615a5aeb0c0771fa72.
//
// Solidity: event IntentCompleted(bytes32 indexed intent_id, bytes32 result_hash, uint256 budget)
func (_IntentManager *IntentManagerFilterer) WatchIntentCompleted(opts *bind.WatchOpts, sink chan<- *IntentManagerIntentCompleted, intent_id [][32]byte) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "IntentCompleted", intent_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerIntentCompleted)
				if err := _IntentManager.contract.UnpackLog(event, "IntentCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIntentCompleted is a log parse operation binding the contract event 0x6337eb6af1128c06eedb6c409b7873f58f3afdf8ee2442615a5aeb0c0771fa72.
//
// Solidity: event IntentCompleted(bytes32 indexed intent_id, bytes32 result_hash, uint256 budget)
func (_IntentManager *IntentManagerFilterer) ParseIntentCompleted(log types.Log) (*IntentManagerIntentCompleted, error) {
	event := new(IntentManagerIntentCompleted)
	if err := _IntentManager.contract.UnpackLog(event, "IntentCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerIntentExpiredIterator is returned from FilterIntentExpired and is used to iterate over the raw logs and unpacked data for IntentExpired events raised by the IntentManager contract.
type IntentManagerIntentExpiredIterator struct {
	Event *IntentManagerIntentExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerIntentExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerIntentExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerIntentExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerIntentExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerIntentExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerIntentExpired represents a IntentExpired event raised by the IntentManager contract.
type IntentManagerIntentExpired struct {
	IntentId     [32]byte
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIntentExpired is a free log retrieval operation binding the contract event 0xf5daf52b2e562ec522fee04d65ed2e6335332dc3fcc8634e2de1ced482b41cfd.
//
// Solidity: event IntentExpired(bytes32 indexed intent_id, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) FilterIntentExpired(opts *bind.FilterOpts, intent_id [][32]byte) (*IntentManagerIntentExpiredIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "IntentExpired", intent_idRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerIntentExpiredIterator{contract: _IntentManager.contract, event: "IntentExpired", logs: logs, sub: sub}, nil
}

// WatchIntentExpired is a free log subscription operation binding the contract event 0xf5daf52b2e562ec522fee04d65ed2e6335332dc3fcc8634e2de1ced482b41cfd.
//
// Solidity: event IntentExpired(bytes32 indexed intent_id, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) WatchIntentExpired(opts *bind.WatchOpts, sink chan<- *IntentManagerIntentExpired, intent_id [][32]byte) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "IntentExpired", intent_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerIntentExpired)
				if err := _IntentManager.contract.UnpackLog(event, "IntentExpired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIntentExpired is a log parse operation binding the contract event 0xf5daf52b2e562ec522fee04d65ed2e6335332dc3fcc8634e2de1ced482b41cfd.
//
// Solidity: event IntentExpired(bytes32 indexed intent_id, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) ParseIntentExpired(log types.Log) (*IntentManagerIntentExpired, error) {
	event := new(IntentManagerIntentExpired)
	if err := _IntentManager.contract.UnpackLog(event, "IntentExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerIntentFailedIterator is returned from FilterIntentFailed and is used to iterate over the raw logs and unpacked data for IntentFailed events raised by the IntentManager contract.
type IntentManagerIntentFailedIterator struct {
	Event *IntentManagerIntentFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerIntentFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerIntentFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerIntentFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerIntentFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerIntentFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerIntentFailed represents a IntentFailed event raised by the IntentManager contract.
type IntentManagerIntentFailed struct {
	IntentId     [32]byte
	Reason       string
	RefundAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterIntentFailed is a free log retrieval operation binding the contract event 0xd6414e0e326510fe0fa0469f2640d922d86657b6847fb93acf969a40b1d5c098.
//
// Solidity: event IntentFailed(bytes32 indexed intent_id, string reason, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) FilterIntentFailed(opts *bind.FilterOpts, intent_id [][32]byte) (*IntentManagerIntentFailedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "IntentFailed", intent_idRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerIntentFailedIterator{contract: _IntentManager.contract, event: "IntentFailed", logs: logs, sub: sub}, nil
}

// WatchIntentFailed is a free log subscription operation binding the contract event 0xd6414e0e326510fe0fa0469f2640d922d86657b6847fb93acf969a40b1d5c098.
//
// Solidity: event IntentFailed(bytes32 indexed intent_id, string reason, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) WatchIntentFailed(opts *bind.WatchOpts, sink chan<- *IntentManagerIntentFailed, intent_id [][32]byte) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "IntentFailed", intent_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerIntentFailed)
				if err := _IntentManager.contract.UnpackLog(event, "IntentFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIntentFailed is a log parse operation binding the contract event 0xd6414e0e326510fe0fa0469f2640d922d86657b6847fb93acf969a40b1d5c098.
//
// Solidity: event IntentFailed(bytes32 indexed intent_id, string reason, uint256 refund_amount)
func (_IntentManager *IntentManagerFilterer) ParseIntentFailed(log types.Log) (*IntentManagerIntentFailed, error) {
	event := new(IntentManagerIntentFailed)
	if err := _IntentManager.contract.UnpackLog(event, "IntentFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerIntentStatusUpdatedIterator is returned from FilterIntentStatusUpdated and is used to iterate over the raw logs and unpacked data for IntentStatusUpdated events raised by the IntentManager contract.
type IntentManagerIntentStatusUpdatedIterator struct {
	Event *IntentManagerIntentStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerIntentStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerIntentStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerIntentStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerIntentStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerIntentStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerIntentStatusUpdated represents a IntentStatusUpdated event raised by the IntentManager contract.
type IntentManagerIntentStatusUpdated struct {
	IntentId  [32]byte
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntentStatusUpdated is a free log retrieval operation binding the contract event 0x9273b77aa79b875a72314eda04e588cae40d4a1aedf5ca90861b34a9b8fdebc0.
//
// Solidity: event IntentStatusUpdated(bytes32 indexed intent_id, uint8 old_status, uint8 new_status)
func (_IntentManager *IntentManagerFilterer) FilterIntentStatusUpdated(opts *bind.FilterOpts, intent_id [][32]byte) (*IntentManagerIntentStatusUpdatedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "IntentStatusUpdated", intent_idRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerIntentStatusUpdatedIterator{contract: _IntentManager.contract, event: "IntentStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchIntentStatusUpdated is a free log subscription operation binding the contract event 0x9273b77aa79b875a72314eda04e588cae40d4a1aedf5ca90861b34a9b8fdebc0.
//
// Solidity: event IntentStatusUpdated(bytes32 indexed intent_id, uint8 old_status, uint8 new_status)
func (_IntentManager *IntentManagerFilterer) WatchIntentStatusUpdated(opts *bind.WatchOpts, sink chan<- *IntentManagerIntentStatusUpdated, intent_id [][32]byte) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "IntentStatusUpdated", intent_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerIntentStatusUpdated)
				if err := _IntentManager.contract.UnpackLog(event, "IntentStatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIntentStatusUpdated is a log parse operation binding the contract event 0x9273b77aa79b875a72314eda04e588cae40d4a1aedf5ca90861b34a9b8fdebc0.
//
// Solidity: event IntentStatusUpdated(bytes32 indexed intent_id, uint8 old_status, uint8 new_status)
func (_IntentManager *IntentManagerFilterer) ParseIntentStatusUpdated(log types.Log) (*IntentManagerIntentStatusUpdated, error) {
	event := new(IntentManagerIntentStatusUpdated)
	if err := _IntentManager.contract.UnpackLog(event, "IntentStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerIntentSubmittedIterator is returned from FilterIntentSubmitted and is used to iterate over the raw logs and unpacked data for IntentSubmitted events raised by the IntentManager contract.
type IntentManagerIntentSubmittedIterator struct {
	Event *IntentManagerIntentSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerIntentSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerIntentSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerIntentSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerIntentSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerIntentSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerIntentSubmitted represents a IntentSubmitted event raised by the IntentManager contract.
type IntentManagerIntentSubmitted struct {
	IntentId  [32]byte
	Requester common.Address
	SubnetId  [32]byte
	Budget    *big.Int
	Deadline  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntentSubmitted is a free log retrieval operation binding the contract event 0xcc1464db8c5ff6ec1c2d9dffc7fb39e45fdcde48ac1e01b96689be38745da504.
//
// Solidity: event IntentSubmitted(bytes32 indexed intent_id, address indexed requester, bytes32 indexed subnet_id, uint256 budget, uint256 deadline)
func (_IntentManager *IntentManagerFilterer) FilterIntentSubmitted(opts *bind.FilterOpts, intent_id [][32]byte, requester []common.Address, subnet_id [][32]byte) (*IntentManagerIntentSubmittedIterator, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "IntentSubmitted", intent_idRule, requesterRule, subnet_idRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerIntentSubmittedIterator{contract: _IntentManager.contract, event: "IntentSubmitted", logs: logs, sub: sub}, nil
}

// WatchIntentSubmitted is a free log subscription operation binding the contract event 0xcc1464db8c5ff6ec1c2d9dffc7fb39e45fdcde48ac1e01b96689be38745da504.
//
// Solidity: event IntentSubmitted(bytes32 indexed intent_id, address indexed requester, bytes32 indexed subnet_id, uint256 budget, uint256 deadline)
func (_IntentManager *IntentManagerFilterer) WatchIntentSubmitted(opts *bind.WatchOpts, sink chan<- *IntentManagerIntentSubmitted, intent_id [][32]byte, requester []common.Address, subnet_id [][32]byte) (event.Subscription, error) {

	var intent_idRule []interface{}
	for _, intent_idItem := range intent_id {
		intent_idRule = append(intent_idRule, intent_idItem)
	}
	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}
	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "IntentSubmitted", intent_idRule, requesterRule, subnet_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerIntentSubmitted)
				if err := _IntentManager.contract.UnpackLog(event, "IntentSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIntentSubmitted is a log parse operation binding the contract event 0xcc1464db8c5ff6ec1c2d9dffc7fb39e45fdcde48ac1e01b96689be38745da504.
//
// Solidity: event IntentSubmitted(bytes32 indexed intent_id, address indexed requester, bytes32 indexed subnet_id, uint256 budget, uint256 deadline)
func (_IntentManager *IntentManagerFilterer) ParseIntentSubmitted(log types.Log) (*IntentManagerIntentSubmitted, error) {
	event := new(IntentManagerIntentSubmitted)
	if err := _IntentManager.contract.UnpackLog(event, "IntentSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerMaxIntentDurationUpdatedIterator is returned from FilterMaxIntentDurationUpdated and is used to iterate over the raw logs and unpacked data for MaxIntentDurationUpdated events raised by the IntentManager contract.
type IntentManagerMaxIntentDurationUpdatedIterator struct {
	Event *IntentManagerMaxIntentDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerMaxIntentDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerMaxIntentDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerMaxIntentDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerMaxIntentDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerMaxIntentDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerMaxIntentDurationUpdated represents a MaxIntentDurationUpdated event raised by the IntentManager contract.
type IntentManagerMaxIntentDurationUpdated struct {
	MaxDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxIntentDurationUpdated is a free log retrieval operation binding the contract event 0xb8cdf117f99c9c1f14ad2f8db63ef887aa2444734cf5d51aaea00e46ce898ad6.
//
// Solidity: event MaxIntentDurationUpdated(uint256 max_duration)
func (_IntentManager *IntentManagerFilterer) FilterMaxIntentDurationUpdated(opts *bind.FilterOpts) (*IntentManagerMaxIntentDurationUpdatedIterator, error) {

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "MaxIntentDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &IntentManagerMaxIntentDurationUpdatedIterator{contract: _IntentManager.contract, event: "MaxIntentDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxIntentDurationUpdated is a free log subscription operation binding the contract event 0xb8cdf117f99c9c1f14ad2f8db63ef887aa2444734cf5d51aaea00e46ce898ad6.
//
// Solidity: event MaxIntentDurationUpdated(uint256 max_duration)
func (_IntentManager *IntentManagerFilterer) WatchMaxIntentDurationUpdated(opts *bind.WatchOpts, sink chan<- *IntentManagerMaxIntentDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "MaxIntentDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerMaxIntentDurationUpdated)
				if err := _IntentManager.contract.UnpackLog(event, "MaxIntentDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxIntentDurationUpdated is a log parse operation binding the contract event 0xb8cdf117f99c9c1f14ad2f8db63ef887aa2444734cf5d51aaea00e46ce898ad6.
//
// Solidity: event MaxIntentDurationUpdated(uint256 max_duration)
func (_IntentManager *IntentManagerFilterer) ParseMaxIntentDurationUpdated(log types.Log) (*IntentManagerMaxIntentDurationUpdated, error) {
	event := new(IntentManagerMaxIntentDurationUpdated)
	if err := _IntentManager.contract.UnpackLog(event, "MaxIntentDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the IntentManager contract.
type IntentManagerPausedIterator struct {
	Event *IntentManagerPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerPaused represents a Paused event raised by the IntentManager contract.
type IntentManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IntentManager *IntentManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*IntentManagerPausedIterator, error) {

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &IntentManagerPausedIterator{contract: _IntentManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IntentManager *IntentManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *IntentManagerPaused) (event.Subscription, error) {

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerPaused)
				if err := _IntentManager.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_IntentManager *IntentManagerFilterer) ParsePaused(log types.Log) (*IntentManagerPaused, error) {
	event := new(IntentManagerPaused)
	if err := _IntentManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerRefundClaimableIterator is returned from FilterRefundClaimable and is used to iterate over the raw logs and unpacked data for RefundClaimable events raised by the IntentManager contract.
type IntentManagerRefundClaimableIterator struct {
	Event *IntentManagerRefundClaimable // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerRefundClaimableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerRefundClaimable)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerRefundClaimable)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerRefundClaimableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerRefundClaimableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerRefundClaimable represents a RefundClaimable event raised by the IntentManager contract.
type IntentManagerRefundClaimable struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundClaimable is a free log retrieval operation binding the contract event 0x3749ee386837d1a3e4a5d6221cca5c78ce0ef45a2e7207d50b56cdaba913d463.
//
// Solidity: event RefundClaimable(address indexed user, address indexed token, uint256 amount)
func (_IntentManager *IntentManagerFilterer) FilterRefundClaimable(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*IntentManagerRefundClaimableIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "RefundClaimable", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerRefundClaimableIterator{contract: _IntentManager.contract, event: "RefundClaimable", logs: logs, sub: sub}, nil
}

// WatchRefundClaimable is a free log subscription operation binding the contract event 0x3749ee386837d1a3e4a5d6221cca5c78ce0ef45a2e7207d50b56cdaba913d463.
//
// Solidity: event RefundClaimable(address indexed user, address indexed token, uint256 amount)
func (_IntentManager *IntentManagerFilterer) WatchRefundClaimable(opts *bind.WatchOpts, sink chan<- *IntentManagerRefundClaimable, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "RefundClaimable", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerRefundClaimable)
				if err := _IntentManager.contract.UnpackLog(event, "RefundClaimable", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundClaimable is a log parse operation binding the contract event 0x3749ee386837d1a3e4a5d6221cca5c78ce0ef45a2e7207d50b56cdaba913d463.
//
// Solidity: event RefundClaimable(address indexed user, address indexed token, uint256 amount)
func (_IntentManager *IntentManagerFilterer) ParseRefundClaimable(log types.Log) (*IntentManagerRefundClaimable, error) {
	event := new(IntentManagerRefundClaimable)
	if err := _IntentManager.contract.UnpackLog(event, "RefundClaimable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerRefundClaimedIterator is returned from FilterRefundClaimed and is used to iterate over the raw logs and unpacked data for RefundClaimed events raised by the IntentManager contract.
type IntentManagerRefundClaimedIterator struct {
	Event *IntentManagerRefundClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerRefundClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerRefundClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerRefundClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerRefundClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerRefundClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerRefundClaimed represents a RefundClaimed event raised by the IntentManager contract.
type IntentManagerRefundClaimed struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundClaimed is a free log retrieval operation binding the contract event 0x39bed68a008a68cbf907d7ff6bc3629912af6516cb837cfa3f871ad9f2b8a944.
//
// Solidity: event RefundClaimed(address indexed user, address indexed token, uint256 amount)
func (_IntentManager *IntentManagerFilterer) FilterRefundClaimed(opts *bind.FilterOpts, user []common.Address, token []common.Address) (*IntentManagerRefundClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "RefundClaimed", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerRefundClaimedIterator{contract: _IntentManager.contract, event: "RefundClaimed", logs: logs, sub: sub}, nil
}

// WatchRefundClaimed is a free log subscription operation binding the contract event 0x39bed68a008a68cbf907d7ff6bc3629912af6516cb837cfa3f871ad9f2b8a944.
//
// Solidity: event RefundClaimed(address indexed user, address indexed token, uint256 amount)
func (_IntentManager *IntentManagerFilterer) WatchRefundClaimed(opts *bind.WatchOpts, sink chan<- *IntentManagerRefundClaimed, user []common.Address, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "RefundClaimed", userRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerRefundClaimed)
				if err := _IntentManager.contract.UnpackLog(event, "RefundClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefundClaimed is a log parse operation binding the contract event 0x39bed68a008a68cbf907d7ff6bc3629912af6516cb837cfa3f871ad9f2b8a944.
//
// Solidity: event RefundClaimed(address indexed user, address indexed token, uint256 amount)
func (_IntentManager *IntentManagerFilterer) ParseRefundClaimed(log types.Log) (*IntentManagerRefundClaimed, error) {
	event := new(IntentManagerRefundClaimed)
	if err := _IntentManager.contract.UnpackLog(event, "RefundClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the IntentManager contract.
type IntentManagerRoleAdminChangedIterator struct {
	Event *IntentManagerRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerRoleAdminChanged represents a RoleAdminChanged event raised by the IntentManager contract.
type IntentManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IntentManager *IntentManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*IntentManagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerRoleAdminChangedIterator{contract: _IntentManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IntentManager *IntentManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *IntentManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerRoleAdminChanged)
				if err := _IntentManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_IntentManager *IntentManagerFilterer) ParseRoleAdminChanged(log types.Log) (*IntentManagerRoleAdminChanged, error) {
	event := new(IntentManagerRoleAdminChanged)
	if err := _IntentManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the IntentManager contract.
type IntentManagerRoleGrantedIterator struct {
	Event *IntentManagerRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerRoleGranted represents a RoleGranted event raised by the IntentManager contract.
type IntentManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IntentManager *IntentManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IntentManagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerRoleGrantedIterator{contract: _IntentManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IntentManager *IntentManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *IntentManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerRoleGranted)
				if err := _IntentManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_IntentManager *IntentManagerFilterer) ParseRoleGranted(log types.Log) (*IntentManagerRoleGranted, error) {
	event := new(IntentManagerRoleGranted)
	if err := _IntentManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the IntentManager contract.
type IntentManagerRoleRevokedIterator struct {
	Event *IntentManagerRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerRoleRevoked represents a RoleRevoked event raised by the IntentManager contract.
type IntentManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IntentManager *IntentManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*IntentManagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerRoleRevokedIterator{contract: _IntentManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IntentManager *IntentManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *IntentManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerRoleRevoked)
				if err := _IntentManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_IntentManager *IntentManagerFilterer) ParseRoleRevoked(log types.Log) (*IntentManagerRoleRevoked, error) {
	event := new(IntentManagerRoleRevoked)
	if err := _IntentManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the IntentManager contract.
type IntentManagerUnpausedIterator struct {
	Event *IntentManagerUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerUnpaused represents a Unpaused event raised by the IntentManager contract.
type IntentManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IntentManager *IntentManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*IntentManagerUnpausedIterator, error) {

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &IntentManagerUnpausedIterator{contract: _IntentManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IntentManager *IntentManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *IntentManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerUnpaused)
				if err := _IntentManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_IntentManager *IntentManagerFilterer) ParseUnpaused(log types.Log) (*IntentManagerUnpaused, error) {
	event := new(IntentManagerUnpaused)
	if err := _IntentManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IntentManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the IntentManager contract.
type IntentManagerUpgradedIterator struct {
	Event *IntentManagerUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IntentManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IntentManagerUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IntentManagerUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IntentManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IntentManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IntentManagerUpgraded represents a Upgraded event raised by the IntentManager contract.
type IntentManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IntentManager *IntentManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*IntentManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IntentManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &IntentManagerUpgradedIterator{contract: _IntentManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IntentManager *IntentManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *IntentManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _IntentManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IntentManagerUpgraded)
				if err := _IntentManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_IntentManager *IntentManagerFilterer) ParseUpgraded(log types.Log) (*IntentManagerUpgraded, error) {
	event := new(IntentManagerUpgraded)
	if err := _IntentManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
