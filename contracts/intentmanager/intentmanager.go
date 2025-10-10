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
	IntentId     [32]byte
	SubnetId     [32]byte
	Requester    common.Address
	IntentType   string
	CreatedAt    *big.Int
	Deadline     *big.Int
	ParamsHash   [32]byte
	Budget       *big.Int
	PaymentToken common.Address
	Status       uint8
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

// IntentManagerMetaData contains all meta data concerning the IntentManager contract.
var IntentManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssignmentIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateDigest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientBudget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"IntentAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"current_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"IntentNotExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"IntentNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"current_status\",\"type\":\"uint8\"}],\"name\":\"IntentNotPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"}],\"name\":\"InvalidBudget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"InvalidDeadline\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIntentStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"InvalidSubnet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InvalidValidatorCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"SignatureVerificationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"IntentAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"}],\"name\":\"IntentCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund_amount\",\"type\":\"uint256\"}],\"name\":\"IntentExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund_amount\",\"type\":\"uint256\"}],\"name\":\"IntentFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"old_status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"new_status\",\"type\":\"uint8\"}],\"name\":\"IntentStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"IntentSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_MAX_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_MIN_BUDGET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bid_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.AssignmentStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"}],\"internalType\":\"structIIntentManager.AssignmentData[]\",\"name\":\"assignments\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"assignIntentsBySignatures\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"assignment_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"name\":\"batchGetIntentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structDataStructures.IntentInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"name\":\"batchProcessExpiredIntents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"emergencyRefundBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"failIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"getIntentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structDataStructures.IntentInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"getIntentsByStatus\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxIntentDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinIntentBudget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getPendingIntentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getRequiredValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"required_validators\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetIntents\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTotalEscrowBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalIntentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subnet_factory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"intentExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"isIntentExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"expired\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"processExpiredIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max_duration\",\"type\":\"uint256\"}],\"name\":\"setMaxIntentDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min_budget\",\"type\":\"uint256\"}],\"name\":\"setMinIntentBudget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitIntent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIIntentManager.IntentData[]\",\"name\":\"intents\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"submitIntentsBySignatures\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"result_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"proof_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"root_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structIIntentManager.ValidationBundleData[]\",\"name\":\"validations\",\"type\":\"tuple[]\"}],\"name\":\"validateIntentsBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DEFAULTMINBUDGET is a free data retrieval call binding the contract method 0xda6290a1.
//
// Solidity: function DEFAULT_MIN_BUDGET() view returns(uint256)
func (_IntentManager *IntentManagerCaller) DEFAULTMINBUDGET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "DEFAULT_MIN_BUDGET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTMINBUDGET is a free data retrieval call binding the contract method 0xda6290a1.
//
// Solidity: function DEFAULT_MIN_BUDGET() view returns(uint256)
func (_IntentManager *IntentManagerSession) DEFAULTMINBUDGET() (*big.Int, error) {
	return _IntentManager.Contract.DEFAULTMINBUDGET(&_IntentManager.CallOpts)
}

// DEFAULTMINBUDGET is a free data retrieval call binding the contract method 0xda6290a1.
//
// Solidity: function DEFAULT_MIN_BUDGET() view returns(uint256)
func (_IntentManager *IntentManagerCallerSession) DEFAULTMINBUDGET() (*big.Int, error) {
	return _IntentManager.Contract.DEFAULTMINBUDGET(&_IntentManager.CallOpts)
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

// BatchGetIntentInfo is a free data retrieval call binding the contract method 0xb6d1124e.
//
// Solidity: function batchGetIntentInfo(bytes32[] intent_ids) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8)[] infos)
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
// Solidity: function batchGetIntentInfo(bytes32[] intent_ids) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8)[] infos)
func (_IntentManager *IntentManagerSession) BatchGetIntentInfo(intent_ids [][32]byte) ([]DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.BatchGetIntentInfo(&_IntentManager.CallOpts, intent_ids)
}

// BatchGetIntentInfo is a free data retrieval call binding the contract method 0xb6d1124e.
//
// Solidity: function batchGetIntentInfo(bytes32[] intent_ids) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8)[] infos)
func (_IntentManager *IntentManagerCallerSession) BatchGetIntentInfo(intent_ids [][32]byte) ([]DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.BatchGetIntentInfo(&_IntentManager.CallOpts, intent_ids)
}

// GetIntentInfo is a free data retrieval call binding the contract method 0x283ea6bb.
//
// Solidity: function getIntentInfo(bytes32 intent_id) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8) info)
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
// Solidity: function getIntentInfo(bytes32 intent_id) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8) info)
func (_IntentManager *IntentManagerSession) GetIntentInfo(intent_id [32]byte) (DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.GetIntentInfo(&_IntentManager.CallOpts, intent_id)
}

// GetIntentInfo is a free data retrieval call binding the contract method 0x283ea6bb.
//
// Solidity: function getIntentInfo(bytes32 intent_id) view returns((bytes32,bytes32,address,string,uint256,uint256,bytes32,uint256,address,uint8) info)
func (_IntentManager *IntentManagerCallerSession) GetIntentInfo(intent_id [32]byte) (DataStructuresIntentInfo, error) {
	return _IntentManager.Contract.GetIntentInfo(&_IntentManager.CallOpts, intent_id)
}

// GetIntentsByStatus is a free data retrieval call binding the contract method 0x9fd71f61.
//
// Solidity: function getIntentsByStatus(uint8 status) view returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerCaller) GetIntentsByStatus(opts *bind.CallOpts, status uint8) ([][32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getIntentsByStatus", status)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetIntentsByStatus is a free data retrieval call binding the contract method 0x9fd71f61.
//
// Solidity: function getIntentsByStatus(uint8 status) view returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerSession) GetIntentsByStatus(status uint8) ([][32]byte, error) {
	return _IntentManager.Contract.GetIntentsByStatus(&_IntentManager.CallOpts, status)
}

// GetIntentsByStatus is a free data retrieval call binding the contract method 0x9fd71f61.
//
// Solidity: function getIntentsByStatus(uint8 status) view returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerCallerSession) GetIntentsByStatus(status uint8) ([][32]byte, error) {
	return _IntentManager.Contract.GetIntentsByStatus(&_IntentManager.CallOpts, status)
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

// GetMinIntentBudget is a free data retrieval call binding the contract method 0x072c9176.
//
// Solidity: function getMinIntentBudget() view returns(uint256 budget)
func (_IntentManager *IntentManagerCaller) GetMinIntentBudget(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getMinIntentBudget")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinIntentBudget is a free data retrieval call binding the contract method 0x072c9176.
//
// Solidity: function getMinIntentBudget() view returns(uint256 budget)
func (_IntentManager *IntentManagerSession) GetMinIntentBudget() (*big.Int, error) {
	return _IntentManager.Contract.GetMinIntentBudget(&_IntentManager.CallOpts)
}

// GetMinIntentBudget is a free data retrieval call binding the contract method 0x072c9176.
//
// Solidity: function getMinIntentBudget() view returns(uint256 budget)
func (_IntentManager *IntentManagerCallerSession) GetMinIntentBudget() (*big.Int, error) {
	return _IntentManager.Contract.GetMinIntentBudget(&_IntentManager.CallOpts)
}

// GetPendingIntentCount is a free data retrieval call binding the contract method 0x00161b7c.
//
// Solidity: function getPendingIntentCount(bytes32 subnet_id) view returns(uint256 count)
func (_IntentManager *IntentManagerCaller) GetPendingIntentCount(opts *bind.CallOpts, subnet_id [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getPendingIntentCount", subnet_id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingIntentCount is a free data retrieval call binding the contract method 0x00161b7c.
//
// Solidity: function getPendingIntentCount(bytes32 subnet_id) view returns(uint256 count)
func (_IntentManager *IntentManagerSession) GetPendingIntentCount(subnet_id [32]byte) (*big.Int, error) {
	return _IntentManager.Contract.GetPendingIntentCount(&_IntentManager.CallOpts, subnet_id)
}

// GetPendingIntentCount is a free data retrieval call binding the contract method 0x00161b7c.
//
// Solidity: function getPendingIntentCount(bytes32 subnet_id) view returns(uint256 count)
func (_IntentManager *IntentManagerCallerSession) GetPendingIntentCount(subnet_id [32]byte) (*big.Int, error) {
	return _IntentManager.Contract.GetPendingIntentCount(&_IntentManager.CallOpts, subnet_id)
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

// GetSubnetIntents is a free data retrieval call binding the contract method 0x4cea1bbf.
//
// Solidity: function getSubnetIntents(bytes32 subnet_id) view returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerCaller) GetSubnetIntents(opts *bind.CallOpts, subnet_id [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getSubnetIntents", subnet_id)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSubnetIntents is a free data retrieval call binding the contract method 0x4cea1bbf.
//
// Solidity: function getSubnetIntents(bytes32 subnet_id) view returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerSession) GetSubnetIntents(subnet_id [32]byte) ([][32]byte, error) {
	return _IntentManager.Contract.GetSubnetIntents(&_IntentManager.CallOpts, subnet_id)
}

// GetSubnetIntents is a free data retrieval call binding the contract method 0x4cea1bbf.
//
// Solidity: function getSubnetIntents(bytes32 subnet_id) view returns(bytes32[] intent_ids)
func (_IntentManager *IntentManagerCallerSession) GetSubnetIntents(subnet_id [32]byte) ([][32]byte, error) {
	return _IntentManager.Contract.GetSubnetIntents(&_IntentManager.CallOpts, subnet_id)
}

// GetTotalEscrowBalance is a free data retrieval call binding the contract method 0x77a1386e.
//
// Solidity: function getTotalEscrowBalance(address token) view returns(uint256 amount)
func (_IntentManager *IntentManagerCaller) GetTotalEscrowBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getTotalEscrowBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalEscrowBalance is a free data retrieval call binding the contract method 0x77a1386e.
//
// Solidity: function getTotalEscrowBalance(address token) view returns(uint256 amount)
func (_IntentManager *IntentManagerSession) GetTotalEscrowBalance(token common.Address) (*big.Int, error) {
	return _IntentManager.Contract.GetTotalEscrowBalance(&_IntentManager.CallOpts, token)
}

// GetTotalEscrowBalance is a free data retrieval call binding the contract method 0x77a1386e.
//
// Solidity: function getTotalEscrowBalance(address token) view returns(uint256 amount)
func (_IntentManager *IntentManagerCallerSession) GetTotalEscrowBalance(token common.Address) (*big.Int, error) {
	return _IntentManager.Contract.GetTotalEscrowBalance(&_IntentManager.CallOpts, token)
}

// GetTotalIntentCount is a free data retrieval call binding the contract method 0x40846e54.
//
// Solidity: function getTotalIntentCount() view returns(uint256 count)
func (_IntentManager *IntentManagerCaller) GetTotalIntentCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IntentManager.contract.Call(opts, &out, "getTotalIntentCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalIntentCount is a free data retrieval call binding the contract method 0x40846e54.
//
// Solidity: function getTotalIntentCount() view returns(uint256 count)
func (_IntentManager *IntentManagerSession) GetTotalIntentCount() (*big.Int, error) {
	return _IntentManager.Contract.GetTotalIntentCount(&_IntentManager.CallOpts)
}

// GetTotalIntentCount is a free data retrieval call binding the contract method 0x40846e54.
//
// Solidity: function getTotalIntentCount() view returns(uint256 count)
func (_IntentManager *IntentManagerCallerSession) GetTotalIntentCount() (*big.Int, error) {
	return _IntentManager.Contract.GetTotalIntentCount(&_IntentManager.CallOpts)
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

// ProcessExpiredIntent is a paid mutator transaction binding the contract method 0x707e9ea2.
//
// Solidity: function processExpiredIntent(bytes32 intent_id) returns()
func (_IntentManager *IntentManagerTransactor) ProcessExpiredIntent(opts *bind.TransactOpts, intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "processExpiredIntent", intent_id)
}

// ProcessExpiredIntent is a paid mutator transaction binding the contract method 0x707e9ea2.
//
// Solidity: function processExpiredIntent(bytes32 intent_id) returns()
func (_IntentManager *IntentManagerSession) ProcessExpiredIntent(intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.ProcessExpiredIntent(&_IntentManager.TransactOpts, intent_id)
}

// ProcessExpiredIntent is a paid mutator transaction binding the contract method 0x707e9ea2.
//
// Solidity: function processExpiredIntent(bytes32 intent_id) returns()
func (_IntentManager *IntentManagerTransactorSession) ProcessExpiredIntent(intent_id [32]byte) (*types.Transaction, error) {
	return _IntentManager.Contract.ProcessExpiredIntent(&_IntentManager.TransactOpts, intent_id)
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

// SetMinIntentBudget is a paid mutator transaction binding the contract method 0xd3f80e06.
//
// Solidity: function setMinIntentBudget(uint256 min_budget) returns()
func (_IntentManager *IntentManagerTransactor) SetMinIntentBudget(opts *bind.TransactOpts, min_budget *big.Int) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "setMinIntentBudget", min_budget)
}

// SetMinIntentBudget is a paid mutator transaction binding the contract method 0xd3f80e06.
//
// Solidity: function setMinIntentBudget(uint256 min_budget) returns()
func (_IntentManager *IntentManagerSession) SetMinIntentBudget(min_budget *big.Int) (*types.Transaction, error) {
	return _IntentManager.Contract.SetMinIntentBudget(&_IntentManager.TransactOpts, min_budget)
}

// SetMinIntentBudget is a paid mutator transaction binding the contract method 0xd3f80e06.
//
// Solidity: function setMinIntentBudget(uint256 min_budget) returns()
func (_IntentManager *IntentManagerTransactorSession) SetMinIntentBudget(min_budget *big.Int) (*types.Transaction, error) {
	return _IntentManager.Contract.SetMinIntentBudget(&_IntentManager.TransactOpts, min_budget)
}

// SubmitIntent is a paid mutator transaction binding the contract method 0x8e1a703c.
//
// Solidity: function submitIntent(bytes32 intent_id, bytes32 subnet_id, string intent_type, bytes32 params_hash, uint256 deadline, address payment_token, uint256 amount) payable returns(bytes32)
func (_IntentManager *IntentManagerTransactor) SubmitIntent(opts *bind.TransactOpts, intent_id [32]byte, subnet_id [32]byte, intent_type string, params_hash [32]byte, deadline *big.Int, payment_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "submitIntent", intent_id, subnet_id, intent_type, params_hash, deadline, payment_token, amount)
}

// SubmitIntent is a paid mutator transaction binding the contract method 0x8e1a703c.
//
// Solidity: function submitIntent(bytes32 intent_id, bytes32 subnet_id, string intent_type, bytes32 params_hash, uint256 deadline, address payment_token, uint256 amount) payable returns(bytes32)
func (_IntentManager *IntentManagerSession) SubmitIntent(intent_id [32]byte, subnet_id [32]byte, intent_type string, params_hash [32]byte, deadline *big.Int, payment_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IntentManager.Contract.SubmitIntent(&_IntentManager.TransactOpts, intent_id, subnet_id, intent_type, params_hash, deadline, payment_token, amount)
}

// SubmitIntent is a paid mutator transaction binding the contract method 0x8e1a703c.
//
// Solidity: function submitIntent(bytes32 intent_id, bytes32 subnet_id, string intent_type, bytes32 params_hash, uint256 deadline, address payment_token, uint256 amount) payable returns(bytes32)
func (_IntentManager *IntentManagerTransactorSession) SubmitIntent(intent_id [32]byte, subnet_id [32]byte, intent_type string, params_hash [32]byte, deadline *big.Int, payment_token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IntentManager.Contract.SubmitIntent(&_IntentManager.TransactOpts, intent_id, subnet_id, intent_type, params_hash, deadline, payment_token, amount)
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

// ValidateIntentsBySignatures is a paid mutator transaction binding the contract method 0x60e62769.
//
// Solidity: function validateIntentsBySignatures((bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address[],bytes[])[] validations) returns()
func (_IntentManager *IntentManagerTransactor) ValidateIntentsBySignatures(opts *bind.TransactOpts, validations []IIntentManagerValidationBundleData) (*types.Transaction, error) {
	return _IntentManager.contract.Transact(opts, "validateIntentsBySignatures", validations)
}

// ValidateIntentsBySignatures is a paid mutator transaction binding the contract method 0x60e62769.
//
// Solidity: function validateIntentsBySignatures((bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address[],bytes[])[] validations) returns()
func (_IntentManager *IntentManagerSession) ValidateIntentsBySignatures(validations []IIntentManagerValidationBundleData) (*types.Transaction, error) {
	return _IntentManager.Contract.ValidateIntentsBySignatures(&_IntentManager.TransactOpts, validations)
}

// ValidateIntentsBySignatures is a paid mutator transaction binding the contract method 0x60e62769.
//
// Solidity: function validateIntentsBySignatures((bytes32,bytes32,bytes32,address,bytes32,bytes32,uint64,bytes32,address[],bytes[])[] validations) returns()
func (_IntentManager *IntentManagerTransactorSession) ValidateIntentsBySignatures(validations []IIntentManagerValidationBundleData) (*types.Transaction, error) {
	return _IntentManager.Contract.ValidateIntentsBySignatures(&_IntentManager.TransactOpts, validations)
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
