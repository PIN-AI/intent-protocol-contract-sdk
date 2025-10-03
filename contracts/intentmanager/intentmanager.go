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
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssignmentIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateDigest\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyArray\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InsufficientBudget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"IntentAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"current_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"IntentNotExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"IntentNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"current_status\",\"type\":\"uint8\"}],\"name\":\"IntentNotPending\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"}],\"name\":\"InvalidBudget\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"InvalidDeadline\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIntentStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"InvalidSubnet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"InvalidValidator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"InvalidValidatorCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"SignatureVerificationFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"IntentAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"}],\"name\":\"IntentCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund_amount\",\"type\":\"uint256\"}],\"name\":\"IntentExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund_amount\",\"type\":\"uint256\"}],\"name\":\"IntentFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"old_status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"new_status\",\"type\":\"uint8\"}],\"name\":\"IntentStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"IntentSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_MAX_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_MIN_BUDGET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"bid_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.AssignmentStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matcher\",\"type\":\"address\"}],\"internalType\":\"structIIntentManager.AssignmentData[]\",\"name\":\"assignments\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"assignIntentsBySignatures\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"assignment_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"name\":\"batchProcessExpiredIntents\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"emergencyRefundBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"failIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"getIntentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structDataStructures.IntentInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.IntentStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"getIntentsByStatus\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxIntentDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinIntentBudget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"budget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getPendingIntentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getRequiredValidatorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"required_validators\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetIntents\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTotalEscrowBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalIntentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subnet_factory\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"intentExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"isIntentExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"expired\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"}],\"name\":\"processExpiredIntent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max_duration\",\"type\":\"uint256\"}],\"name\":\"setMaxIntentDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min_budget\",\"type\":\"uint256\"}],\"name\":\"setMinIntentBudget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"submitIntent\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"intent_type\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"payment_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIIntentManager.IntentData[]\",\"name\":\"intents\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"submitIntentsBySignatures\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"intent_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"intent_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignment_id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"result_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"proof_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"root_height\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"root_hash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structIIntentManager.ValidationBundleData[]\",\"name\":\"validations\",\"type\":\"tuple[]\"}],\"name\":\"validateIntentsBySignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061436f8061001f6000396000f3fe6080604052600436106101935760003560e01c8062161b7c1461019857806301ffc9a7146101cb578063072c9176146101fb57806312fcd0ec14610210578063248a9ca314610225578063283ea6bb146102455780632f2ff15d1461027257806336568abe1461029457806338a8455d146102b457806340846e54146102d4578063485cc955146102e95780634a4e3bd5146103095780634cea1bbf1461031e57806351858e271461034b5780635c975abb1461036057806360e6276914610375578063707e9ea21461039557806377a1386e146103b55780637ba19837146103eb5780638e1a703c1461040b57806391d148541461041e578063931280e31461043e5780639a240f43146104515780639e028794146104715780639fd71f6114610491578063a0f548c3146104b1578063a217fddf146104d1578063c34f2f22146104e6578063d3f80e0614610506578063d547741f14610526578063d9423faf14610546578063da6290a1146104d1578063f36c8f5c1461055d578063f83e09411461057f578063ff4d01731461059f575b600080fd5b3480156101a457600080fd5b506101b86101b33660046134bc565b6105bf565b6040519081526020015b60405180910390f35b3480156101d757600080fd5b506101eb6101e63660046134d5565b6106c1565b60405190151581526020016101c2565b34801561020757600080fd5b50600a546101b8565b34801561021c57600080fd5b506009546101b8565b34801561023157600080fd5b506101b86102403660046134bc565b6106f8565b34801561025157600080fd5b506102656102603660046134bc565b610718565b6040516101c29190613586565b34801561027e57600080fd5b5061029261028d36600461363c565b6108ba565b005b3480156102a057600080fd5b506102926102af36600461363c565b6108dc565b3480156102c057600080fd5b506102926102cf3660046136b7565b610914565b3480156102e057600080fd5b506101b8610a4a565b3480156102f557600080fd5b506102926103043660046136f8565b610b1d565b34801561031557600080fd5b50610292610c72565b34801561032a57600080fd5b5061033e6103393660046134bc565b610c95565b6040516101c29190613726565b34801561035757600080fd5b50610292610cf7565b34801561036c57600080fd5b506101eb610d17565b34801561038157600080fd5b506102926103903660046136b7565b610d2c565b3480156103a157600080fd5b506102926103b03660046134bc565b611043565b3480156103c157600080fd5b506101b86103d0366004613769565b6001600160a01b031660009081526003602052604090205490565b3480156103f757600080fd5b506102926104063660046134bc565b61118d565b6101b86104193660046137c7565b6111cf565b34801561042a57600080fd5b506101eb61043936600461363c565b611383565b61033e61044c36600461384d565b6113b9565b34801561045d57600080fd5b506101b861046c3660046134bc565b6117ef565b34801561047d57600080fd5b506101eb61048c3660046134bc565b6119d2565b34801561049d57600080fd5b5061033e6104ac3660046138bc565b6119e9565b3480156104bd57600080fd5b506102926104cc3660046138dd565b611a74565b3480156104dd57600080fd5b506101b8600081565b3480156104f257600080fd5b5061033e610501366004613940565b611bad565b34801561051257600080fd5b506102926105213660046134bc565b611ffe565b34801561053257600080fd5b5061029261054136600461363c565b61201c565b34801561055257600080fd5b506101b862093a8081565b34801561056957600080fd5b506101b86000805160206142ba83398151915281565b34801561058b57600080fd5b506101eb61059a3660046134bc565b612038565b3480156105ab57600080fd5b506102926105ba3660046139c8565b61219f565b6000816105e45750506000805260026020526000805160206142fa8339815191525490565b60008281526001602090815260408083208054825181850281018501909352808352919290919083018282801561063a57602002820191906000526020600020905b815481526020019060010190808311610626575b505050505090506000915060005b81518110156106ba576000808084848151811061066757610667613a75565b6020026020010151815260200190815260200160002060080160149054906101000a900460ff16600481111561069f5761069f61355c565b036106b257826106ae81613aa1565b9350505b600101610648565b5050919050565b60006001600160e01b03198216637965db0b60e01b14806106f257506301ffc9a760e01b6001600160e01b03198316145b92915050565b6000806107036122e4565b60009384526020525050604090206001015490565b61076c60408051610140810182526000808252602082018190529181018290526060808201526080810182905260a0810182905260c0810182905260e0810182905261010081018290529061012082015290565b600082815260208181526040918290208251610140810184528154815260018201549281019290925260028101546001600160a01b0316928201929092526003820180549192916060840191906107c290613aba565b80601f01602080910402602001604051908101604052809291908181526020018280546107ee90613aba565b801561083b5780601f106108105761010080835404028352916020019161083b565b820191906000526020600020905b81548152906001019060200180831161081e57829003601f168201915b5050509183525050600482810154602083015260058301546040830152600683015460608301526007830154608083015260088301546001600160a01b03811660a084015260c090920191600160a01b900460ff16908111156108a0576108a061355c565b60048111156108b1576108b161355c565b90525092915050565b6108c3826106f8565b6108cc81612308565b6108d68383612312565b50505050565b6001600160a01b03811633146109055760405163334bd91960e11b815260040160405180910390fd5b61090f82826123b3565b505050565b61091c61242b565b610924612453565b60008190036109465760405163521299a960e01b815260040160405180910390fd5b60005b81811015610a3d57600083838381811061096557610965613a75565b905060200201359050600080600083815260200190815260200160002090508060040154600014806109b7575060006008820154600160a01b900460ff1660048111156109b4576109b461355c565b14155b156109c3575050610a35565b806005015442116109d5575050610a35565b6109e0826003612489565b600281015460078201546008830154610a06926001600160a01b03908116929116612566565b816000805160206142da8339815191528260070154604051610a2a91815260200190565b60405180910390a250505b600101610949565b50610a4661259f565b5050565b60026020527fee60d0579bcffd98e668647d59fec1ff86a7fb340ce572e844f234ae73a6918f547f88601476d11616a71c5be67555bd1dff4b1cbf21533d2669b768b61518cfe1c3547f679795a0195a1b76cdebb7c51d74e058aee92919b8c3389af86ef24535e8a28c547fe90b7bceb6e7df5418fb78d8ee546e97c83a08bbccc01a0644d599ccd2a7c2e05460008080526000805160206142fa833981519152549094939291610afa91613af4565b610b049190613af4565b610b0e9190613af4565b610b189190613af4565b905090565b6000610b276125b0565b805490915060ff600160401b82041615906001600160401b0316600081158015610b4e5750825b90506000826001600160401b03166001148015610b6a5750303b155b905081158015610b78575080155b15610b965760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315610bbf57845460ff60401b1916600160401b1785555b610bc76125d9565b610bcf6125d9565b610bd76125e1565b610be2600088612312565b50610bfb6000805160206142ba83398151915288612312565b50600b80546001600160a01b0319166001600160a01b03881617905562093a806009556000600a558315610c6957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b6000805160206142ba833981519152610c8a81612308565b610c926125f1565b50565b600081815260016020908152604091829020805483518184028101840190945280845260609392830182828015610ceb57602002820191906000526020600020905b815481526020019060010190808311610cd7575b50505050509050919050565b6000805160206142ba833981519152610d0f81612308565b610c92612648565b600080610d2261268f565b5460ff1692915050565b610d3461242b565b610d3c612453565b6000819003610d5e5760405163521299a960e01b815260040160405180910390fd5b60005b81811015610a3d5736838383818110610d7c57610d7c613a75565b9050602002810190610d8e9190613b07565b9050610d9a81356119d2565b610dbf57604051639481f8b960e01b8152813560048201526024015b60405180910390fd5b8035600090815260208190526040902060016008820154600160a01b900460ff166004811115610df157610df161355c565b14610e0f576040516301cc40e960e11b815260040160405180910390fd5b8060050154421115610e7d57610e2782356003612489565b600281015460078201546008830154610e4d926001600160a01b03908116929116612566565b60078101546040519081528235906000805160206142da8339815191529060200160405180910390a2505061103b565b8135600090815260086020908152604090912054908301358114610eb457604051630c4a772960e41b815260040160405180910390fd5b6000818152600760205260409020610ed26080850160608601613769565b60038201546001600160a01b03908116911614610f0257604051633b30df9b60e01b815260040160405180910390fd5b6000610f0d856126b3565b60008181526006602052604090205490915060ff1615610f405760405163e5fcfa6160e01b815260040160405180910390fd5b6000818152600660205260409020805460ff19166001179055610f638582612793565b610f83576040516359d58b1560e11b815285356004820152602401610db6565b610f8f85356002612489565b600784015460088501546001600160a01b031660009081526003602052604081208054909190610fc0908490613b28565b90915550506008840154610ff1906001600160a01b0316610fe76080880160608901613769565b8660070154612a0f565b60078401546040518635917f6337eb6af1128c06eedb6c409b7873f58f3afdf8ee2442615a5aeb0c0771fa729161102d9160808a013591613b3b565b60405180910390a250505050505b600101610d61565b61104b61242b565b611053612453565b60008181526020819052604081206004015482910361108857604051639481f8b960e01b815260048101829052602401610db6565b6000828152602081905260408120600801548391600160a01b90910460ff16908160048111156110ba576110ba61355c565b146110dc578181604051631d2592d360e01b8152600401610db6929190613b49565b6000848152602081905260409020600581015442116111245760058101546040516313fbeee160e31b8152600481018790524260248201526044810191909152606401610db6565b61112f856003612489565b600281015460078201546008830154611155926001600160a01b03908116929116612566565b846000805160206142da833981519152826007015460405161117991815260200190565b60405180910390a250505050610c9261259f565b6000805160206142ba8339815191526111a581612308565b816000036111c9576040516313b783af60e21b815260048101839052602401610db6565b50600955565b60006111d961242b565b6111e1612453565b600b54604051634ac3ff5960e11b8152600481018a905289916001600160a01b031690639587feb290602401602060405180830381865afa15801561122a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124e9190613b5d565b15806112c45750600b5460405163041c5d3d60e41b8152600481018390526001600160a01b03909116906341c5d3d090602401602060405180830381865afa15801561129e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c29190613b5d565b155b156112e55760405163337d291560e11b815260048101829052602401610db6565b6112ee8a6119d2565b1561130f57604051635eaf4c6960e01b8152600481018b9052602401610db6565b61131885612b21565b60006001600160a01b03851661134357503461133381612b5c565b61133e600082612b84565b611359565b508261134e81612b5c565b611359853383612bc1565b61136a8b8b338c8c8b8d888d612c08565b8a9250505061137761259f565b98975050505050505050565b60008061138e6122e4565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b60606113c361242b565b6113cb612453565b60008490036113ed5760405163521299a960e01b815260040160405180910390fd5b83821461140d5760405163512509d360e11b815260040160405180910390fd5b836001600160401b0381111561142557611425613b7f565b60405190808252806020026020018201604052801561144e578160200160208202803683370190505b5090506000805b858110156117a8573687878381811061147057611470613a75565b90506020028101906114829190613b95565b905061148e81356119d2565b156114af57604051635eaf4c6960e01b815281356004820152602401610db6565b6114bc8160200135612d3a565b6114c98160a00135612b21565b6114d68160e00135612b5c565b60006114e860e0830160c08401613769565b6001600160a01b0316036115075761150460e082013584613af4565b92505b60007f9f2f6d331bdf969f68be0bf76d7a65cadf5bed3811195f9f46c4899b15c627af823560208401356115416060860160408701613769565b61154e6060870187613bab565b60405161155c929190613bf1565b604051908190039020608087013560a088013561157f60e08a0160c08b01613769565b60408051602081019990995288019690965260608701949094526001600160a01b03928316608087015260a086019190915260c085015260e08481019290925290911661010083015283013561012082015230610140820152466101608201526101800160408051601f1981840301815291815281516020928301206000818152600490935291205490915060ff161561162f576040516359d58b1560e11b815260006004820152602401610db6565b6000818152600460205260408120805460ff191660011790556116ba8289898781811061165e5761165e613a75565b90506020028101906116709190613bab565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116b5925050506060870160408801613769565b612dc7565b9050806116dd576040516359d58b1560e11b815260006004820152602401610db6565b60006116ef60e0850160c08601613769565b6001600160a01b03161461172a5761172a61171060e0850160c08601613769565b6117206060860160408701613769565b8560e00135612bc1565b826000013586858151811061174157611741613a75565b60200260200101818152505061179d8360000135846020013585604001602081019061176d9190613769565b61177a6060880188613bab565b60a089013560808a013560e08b018035906117989060c08e01613769565b612c08565b505050600101611455565b5080156117de578034146117d3573481604051630a9a711d60e41b8152600401610db6929190613b3b565b6117de600082612b84565b506117e761259f565b949350505050565b600b5460405163d5a84a1160e01b81526004810183905260009182916001600160a01b039091169063d5a84a1190602401602060405180830381865afa15801561183d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118619190613c11565b90506001600160a01b03811661188d5760405163337d291560e11b815260048101849052602401610db6565b600b54604051632a75726560e11b8152600481018590526000916001600160a01b0316906354eae4ca906024016040805180830381865afa1580156118d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118fa9190613c9b565b90506000826001600160a01b0316635caecd9c60006040518263ffffffff1660e01b815260040161192b9190613d07565b600060405180830381865afa158015611948573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526119709190810190613dcd565b8051909150600081900361198a5750600095945050505050565b6020830151835163ffffffff9182169160019183916119aa911685613fa8565b6119b49190613af4565b6119be9190613b28565b6119c89190613fbf565b9695505050505050565b600090815260208190526040902060040154151590565b606060026000836004811115611a0157611a0161355c565b6004811115611a1257611a1261355c565b8152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610ceb5760200282019190600052602060002090815481526020019060010190808311610cd75750505050509050919050565b6000805160206142ba833981519152611a8c81612308565b611a94612453565b6000849003611ab65760405163521299a960e01b815260040160405180910390fd5b60005b84811015611b9d576000868683818110611ad557611ad5613a75565b90506020020135905060008060008381526020019081526020016000209050806004015460001480611b27575060006008820154600160a01b900460ff166004811115611b2457611b2461355c565b14155b15611b33575050611b95565b611b3e826004612489565b600281015460078201546008830154611b64926001600160a01b03908116929116612566565b8160008051602061431a83398151915287878460070154604051611b8a93929190613fe1565b60405180910390a250505b600101611ab9565b50611ba661259f565b5050505050565b6060611bb761242b565b611bbf612453565b6000849003611be15760405163521299a960e01b815260040160405180910390fd5b838214611c015760405163512509d360e11b815260040160405180910390fd5b836001600160401b03811115611c1957611c19613b7f565b604051908082528060200260200182016040528015611c42578160200160208202803683370190505b50905060005b848110156117de5736868683818110611c6357611c63613a75565b905060c002019050611c7881602001356119d2565b611c9b57604051639481f8b960e01b815260208201356004820152602401610db6565b60208082013560009081529081905260408120906008820154600160a01b900460ff166004811115611ccf57611ccf61355c565b14611d065781602001358160080160149054906101000a900460ff16604051631d2592d360e01b8152600401610db6929190613b49565b60007fdf81f35bd6b981a96c1aa947b697bbe032c52587f803483d420a8b43e4d80b11833560208501356040860135611d456080880160608901613769565b611d5560a0890160808a0161401a565b611d6560c08a0160a08b01613769565b3046604051602001611d7f99989796959493929190614037565b60408051601f1981840301815291815281516020928301206000818152600590935291205490915060ff1615611dc85760405163e5fcfa6160e01b815260040160405180910390fd5b6000818152600560205260408120805460ff19166001179055611e4e82898988818110611df757611df7613a75565b9050602002810190611e099190613bab565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506116b59250505060c0880160a08901613769565b905080611e71576040516359d58b1560e11b815260006004820152602401610db6565b611e8084602001356001612489565b8360000135868681518110611e9757611e97613a75565b6020908102919091018101919091528435600081815260078352604090819020918255918601356001820155908501356002820155611edc6080860160608701613769565b6003820180546001600160a01b0319166001600160a01b0392909216919091179055611f0e60a086016080870161401a565b60038201805460ff60a01b1916600160a01b836002811115611f3257611f3261355c565b0217905550611f4760c0860160a08701613769565b6004820180546001600160a01b0319166001600160a01b0392909216919091179055426005820155602085810135600090815260089091526040902085359055611f9760c0860160a08701613769565b6001600160a01b0316602086013586357fc43578d43d59b01fa3097859d9bb1cbbf2725e102df510e419b9794554894665611fd860808a0160608b01613769565b604051611fe591906140a2565b60405180910390a4505060019093019250611c48915050565b6000805160206142ba83398151915261201681612308565b50600a55565b612025826106f8565b61202e81612308565b6108d683836123b3565b6000818152602081815260408083208151610140810183528154815260018201549381019390935260028101546001600160a01b03169183019190915260038101805484939291606084019161208d90613aba565b80601f01602080910402602001604051908101604052809291908181526020018280546120b990613aba565b80156121065780601f106120db57610100808354040283529160200191612106565b820191906000526020600020905b8154815290600101906020018083116120e957829003601f168201915b5050509183525050600482810154602083015260058301546040830152600683015460608301526007830154608083015260088301546001600160a01b03811660a084015260c090920191600160a01b900460ff169081111561216b5761216b61355c565b600481111561217c5761217c61355c565b90525060808101519091501580159061219857508060a0015142115b9392505050565b6121a761242b565b6121af612453565b6000878152602081905260408120600401548891036121e457604051639481f8b960e01b815260048101829052602401610db6565b6000888152602081905260408120600801548991600160a01b90910460ff16908160048111156122165761221661355c565b14612238578181604051631d2592d360e01b8152600401610db6929190613b49565b60008a81526020819052604090206001810154612259908c8a8a8a8a612e1a565b612279576040516359d58b1560e11b8152600481018c9052602401610db6565b6122848b6004612489565b6002810154600782015460088301546122aa926001600160a01b03908116929116612566565b8a60008051602061431a8339815191528b8b84600701546040516122d093929190613fe1565b60405180910390a250505050610c6961259f565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b610c928133613064565b60008061231d6122e4565b90506123298484611383565b6123a9576000848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561235f3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506106f2565b60009150506106f2565b6000806123be6122e4565b90506123ca8484611383565b156123a9576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506106f2565b612433610d17565b156124515760405163d93c066560e01b815260040160405180910390fd5b565b600061245d61308f565b80549091506001190161248357604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b60008281526020819052604090206008018054600160a01b80820460ff16928492909160ff60a01b19909116908360048111156124c8576124c861355c565b02179055506124d783826130b3565b600260008360048111156124ed576124ed61355c565b60048111156124fe576124fe61355c565b81526020808201929092526040908101600090812080546001810182559082529290209091018490555183907f9273b77aa79b875a72314eda04e588cae40d4a1aedf5ca90861b34a9b8fdebc09061255990849086906140b6565b60405180910390a2505050565b6001600160a01b0381166000908152600360205260408120805484929061258e908490613b28565b9091555061090f9050818484612a0f565b60006125a961308f565b6001905550565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006106f2565b612451613199565b6125e9613199565b6124516131be565b6125f96131c6565b600061260361268f565b805460ff1916815590507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405161263d91906140a2565b60405180910390a150565b61265061242b565b600061265a61268f565b805460ff1916600117815590507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586126303390565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330090565b60007fb3893b731985b5aec1eea59d0df50d6b756e6b081496142b2c5a7bfe7659476d8235602084013560408501356126f26080870160608801613769565b608087013560a088013561270c60e08a0160c08b016140d1565b604080516020810199909952880196909652606087019490945260808601929092526001600160a01b031660a085015260c084015260e0838101919091526001600160401b03909116610100830152830135610120820152306101408201524661016082015261018001604051602081830303815290604052805190602001209050919050565b60006127a36101208401846140fa565b90506127b36101008501856140fa565b9050146127c2575060006106f2565b6127d06101008401846140fa565b90506000036127e1575060006106f2565b60006127f084604001356117ef565b9050806128016101008601866140fa565b905010156128135760009150506106f2565b600b546040805163d5a84a1160e01b81529086013560048201526000916001600160a01b03169063d5a84a1190602401602060405180830381865afa158015612860573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128849190613c11565b90506001600160a01b03811661289f576000925050506106f2565b60005b6128b06101008701876140fa565b9050811015612a035760006128c96101008801886140fa565b838181106128d9576128d9613a75565b90506020020160208101906128ee9190613769565b90506000836001600160a01b031663ac9218908360006040518363ffffffff1660e01b8152600401612921929190614143565b602060405180830381865afa15801561293e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129629190613b5d565b905080612977576000955050505050506106f2565b6129e7876129896101208b018b6140fa565b8681811061299957612999613a75565b90506020028101906129ab9190613bab565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250612dc7915050565b6129f9576000955050505050506106f2565b50506001016128a2565b50600195945050505050565b80600003612a1c57505050565b60006001600160a01b038416612a85576040516001600160a01b038416908390600081818185875af1925050503d8060008114612a75576040519150601f19603f3d011682016040523d82523d6000602084013e612a7a565b606091505b505080915050612afd565b60405163a9059cbb60e01b81526001600160a01b0385169063a9059cbb90612ab39086908690600401614160565b6020604051808303816000875af1925050508015612aee575060408051601f3d908101601f19168201909252612aeb91810190613b5d565b60015b612afa57506000612afd565b90505b806108d6578383836040516317e3057d60e31b8152600401610db693929190614179565b4281111580612b3b5750600954612b389042613af4565b81115b15610c92576040516309decb5d60e21b815260048101829052602401610db6565b600a54811015610c9257600a54604051630a9a711d60e41b8152610db6918391600401613b3b565b80600003612b90575050565b6001600160a01b03821660009081526003602052604081208054839290612bb8908490613af4565b90915550505050565b612bd66001600160a01b0384168330846131eb565b6001600160a01b03831660009081526003602052604081208054839290612bfe908490613af4565b9091555050505050565b6000898152602081905260409020898155600181018990556002810180546001600160a01b0319166001600160a01b038a1617905560038101612c4c8789836141e4565b504260048201556005810185905560068101849055600781018390556008810180546001600160a81b0319166001600160a01b038481169190911790915560008a815260016020818152604080842080548085018255908552828520018f905560029091526000805160206142fa833981519152805492830181559092527f7d2944a272ac5bae96b5bd2f67b6c13276d541dc09eb1cf414d96b19a09e1c2f018c9055518a918a16908c907fcc1464db8c5ff6ec1c2d9dffc7fb39e45fdcde48ac1e01b96689be38745da50490612d269088908b90613b3b565b60405180910390a450505050505050505050565b600b5460405163041c5d3d60e41b8152600481018390526001600160a01b03909116906341c5d3d090602401602060405180830381865afa158015612d83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612da79190613b5d565b610c925760405163337d291560e11b815260048101829052602401610db6565b7b0ca2ba3432b932bab69029b4b3b732b21026b2b9b9b0b3b29d05199960211b6000908152601c849052603c81206000612e018286613245565b6001600160a01b03858116911614925050509392505050565b6000838214612e3c5760405163512509d360e11b815260040160405180910390fd5b6000612e47886117ef565b905080851015612e6e57604051635d0354e760e01b8152610db69086908390600401613b3b565b6040517134b73a32b73a2fb1b7b6b83632ba34b7b71d60711b60208201526032810188905246605282015260009060720160408051808303601f19018152908290528051602090910120600b5463d5a84a1160e01b8352600483018c90529092506000916001600160a01b039091169063d5a84a1190602401602060405180830381865afa158015612f04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f289190613c11565b90506001600160a01b038116612f545760405163337d291560e11b8152600481018b9052602401610db6565b60005b87811015613053576000898983818110612f7357612f73613a75565b9050602002016020810190612f889190613769565b90506000836001600160a01b031663ac9218908360006040518363ffffffff1660e01b8152600401612fbb929190614143565b602060405180830381865afa158015612fd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612ffc9190613b5d565b90508061302057818d60405163107d624560e01b8152600401610db6929190614160565b613036858a8a8681811061299957612999613a75565b61304957600096505050505050506119c8565b5050600101612f57565b5060019a9950505050505050505050565b61306e8282611383565b610a4657808260405163e2517d3f60e01b8152600401610db6929190614160565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b6000600260008360048111156130cb576130cb61355c565b60048111156130dc576130dc61355c565b8152602001908152602001600020905060005b81548110156108d6578382828154811061310b5761310b613a75565b906000526020600020015403613191578154829061312b90600190613b28565b8154811061313b5761313b613a75565b906000526020600020015482828154811061315857613158613a75565b906000526020600020018190555081805480613176576131766142a3565b600190038181906000526020600020016000905590556108d6565b6001016130ef565b6131a161326f565b61245157604051631afcd79f60e31b815260040160405180910390fd5b61259f613199565b6131ce610d17565b61245157604051638dfc202b60e01b815260040160405180910390fd5b6108d684856001600160a01b03166323b872dd86868660405160240161321393929190614179565b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613289565b60008060008061325586866132f1565b925092509250613265828261333e565b5090949350505050565b60006132796125b0565b54600160401b900460ff16919050565b600080602060008451602086016000885af1806132ac576040513d6000823e3d81fd5b50506000513d915081156132c45780600114156132d1565b6001600160a01b0384163b155b156108d65783604051635274afe760e01b8152600401610db691906140a2565b6000806000835160410361332b5760208401516040850151606086015160001a61331d888285856133f7565b955095509550505050613337565b50508151600091506002905b9250925092565b60008260038111156133525761335261355c565b0361335b575050565b600182600381111561336f5761336f61355c565b0361338d5760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156133a1576133a161355c565b036133c25760405163fce698f760e01b815260048101829052602401610db6565b60038260038111156133d6576133d661355c565b03610a46576040516335e2f38360e21b815260048101829052602401610db6565b600080806fa2a8918ca85bafe22016d0b997e4df60600160ff1b0384111561342857506000915060039050826134b2565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561347c573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b0381166134a8575060009250600191508290506134b2565b9250600091508190505b9450945094915050565b6000602082840312156134ce57600080fd5b5035919050565b6000602082840312156134e757600080fd5b81356001600160e01b03198116811461219857600080fd5b6001600160a01b03169052565b60005b8381101561352757818101518382015260200161350f565b50506000910152565b6000815180845261354881602086016020860161350c565b601f01601f19169290920160200192915050565b634e487b7160e01b600052602160045260246000fd5b600581106135825761358261355c565b9052565b602081528151602082015260208201516040820152600060408301516135af60608401826134ff565b50606083015161014060808401526135cb610160840182613530565b9050608084015160a084015260a084015160c084015260c084015160e084015260e084015161010084015261010084015161360a6101208501826134ff565b5061012084015161361f610140850182613572565b509392505050565b6001600160a01b0381168114610c9257600080fd5b6000806040838503121561364f57600080fd5b82359150602083013561366181613627565b809150509250929050565b60008083601f84011261367e57600080fd5b5081356001600160401b0381111561369557600080fd5b6020830191508360208260051b85010111156136b057600080fd5b9250929050565b600080602083850312156136ca57600080fd5b82356001600160401b038111156136e057600080fd5b6136ec8582860161366c565b90969095509350505050565b6000806040838503121561370b57600080fd5b823561371681613627565b9150602083013561366181613627565b602080825282518282018190526000918401906040840190835b8181101561375e578351835260209384019390920191600101613740565b509095945050505050565b60006020828403121561377b57600080fd5b813561219881613627565b60008083601f84011261379857600080fd5b5081356001600160401b038111156137af57600080fd5b6020830191508360208285010111156136b057600080fd5b60008060008060008060008060e0898b0312156137e357600080fd5b883597506020890135965060408901356001600160401b0381111561380757600080fd5b6138138b828c01613786565b909750955050606089013593506080890135925060a089013561383581613627565b979a969950949793969295919450919260c001359150565b6000806000806040858703121561386357600080fd5b84356001600160401b0381111561387957600080fd5b6138858782880161366c565b90955093505060208501356001600160401b038111156138a457600080fd5b6138b08782880161366c565b95989497509550505050565b6000602082840312156138ce57600080fd5b81356005811061219857600080fd5b600080600080604085870312156138f357600080fd5b84356001600160401b0381111561390957600080fd5b6139158782880161366c565b90955093505060208501356001600160401b0381111561393457600080fd5b6138b087828801613786565b6000806000806040858703121561395657600080fd5b84356001600160401b0381111561396c57600080fd5b8501601f8101871361397d57600080fd5b80356001600160401b0381111561399357600080fd5b87602060c0830284010111156139a857600080fd5b6020918201955093508501356001600160401b038111156138a457600080fd5b60008060008060008060006080888a0312156139e357600080fd5b8735965060208801356001600160401b03811115613a0057600080fd5b613a0c8a828b01613786565b90975095505060408801356001600160401b03811115613a2b57600080fd5b613a378a828b0161366c565b90955093505060608801356001600160401b03811115613a5657600080fd5b613a628a828b0161366c565b989b979a50959850939692959293505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201613ab357613ab3613a8b565b5060010190565b600181811c90821680613ace57607f821691505b602082108103613aee57634e487b7160e01b600052602260045260246000fd5b50919050565b808201808211156106f2576106f2613a8b565b6000823561013e19833603018112613b1e57600080fd5b9190910192915050565b818103818111156106f2576106f2613a8b565b918252602082015260400190565b828152604081016121986020830184613572565b600060208284031215613b6f57600080fd5b8151801515811461219857600080fd5b634e487b7160e01b600052604160045260246000fd5b6000823560fe19833603018112613b1e57600080fd5b6000808335601e19843603018112613bc257600080fd5b8301803591506001600160401b03821115613bdc57600080fd5b6020019150368190038213156136b057600080fd5b8183823760009101908152919050565b8051613c0c81613627565b919050565b600060208284031215613c2357600080fd5b815161219881613627565b60405161014081016001600160401b0381118282101715613c5157613c51613b7f565b60405290565b604051601f8201601f191681016001600160401b0381118282101715613c7f57613c7f613b7f565b604052919050565b805163ffffffff81168114613c0c57600080fd5b60006040828403128015613cae57600080fd5b50604080519081016001600160401b0381118282101715613cd157613cd1613b7f565b604052613cdd83613c87565b8152613ceb60208401613c87565b60208201529392505050565b600481106135825761358261355c565b602081016106f28284613cf7565b805160048110613c0c57600080fd5b60038110610c9257600080fd5b8051613c0c81613d24565b80516001600160501b0381168114613c0c57600080fd5b80516001600160801b0381168114613c0c57600080fd5b600082601f830112613d7b57600080fd5b81516001600160401b03811115613d9457613d94613b7f565b613da7601f8201601f1916602001613c57565b818152846020838601011115613dbc57600080fd5b6117e782602083016020870161350c565b600060208284031215613ddf57600080fd5b81516001600160401b03811115613df557600080fd5b8201601f81018413613e0657600080fd5b80516001600160401b03811115613e1f57613e1f613b7f565b8060051b613e2f60208201613c57565b91825260208184018101929081019087841115613e4b57600080fd5b6020850192505b83831015613f9d5782516001600160401b03811115613e7057600080fd5b8501610140818a03601f19011215613e8757600080fd5b613e8f613c2e565b613e9b60208301613c01565b8152613ea960408301613d15565b6020820152613eba60608301613d31565b6040820152613ecb60808301613d3c565b6060820152613edc60a08301613d53565b6080820152613eed60c08301613d53565b60a082015260e082015160c08201526101008201516001600160401b03811115613f1657600080fd5b613f258b602083860101613d6a565b60e0830152506101208201516001600160401b03811115613f4557600080fd5b613f548b602083860101613d6a565b610100830152506101408201516001600160401b03811115613f7557600080fd5b613f848b602083860101613d6a565b6101208301525083525060209283019290910190613e52565b979650505050505050565b80820281158282048414176106f2576106f2613a8b565b600082613fdc57634e487b7160e01b600052601260045260246000fd5b500490565b604081528260408201528284606083013760006060848301015260006060601f19601f8601168301019050826020830152949350505050565b60006020828403121561402c57600080fd5b813561219881613d24565b8981526020810189905260408101889052606081018790526001600160a01b03861660808201526101208101600386106140735761407361355c565b60a08201959095526001600160a01b0393841660c08201529190921660e0820152610100015295945050505050565b6001600160a01b0391909116815260200190565b604081016140c48285613572565b6121986020830184613572565b6000602082840312156140e357600080fd5b81356001600160401b038116811461219857600080fd5b6000808335601e1984360301811261411157600080fd5b8301803591506001600160401b0382111561412b57600080fd5b6020019150600581901b36038213156136b057600080fd5b6001600160a01b0383168152604081016121986020830184613cf7565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b601f82111561090f57806000526020600020601f840160051c810160208510156141c45750805b601f840160051c820191505b81811015611ba657600081556001016141d0565b6001600160401b038311156141fb576141fb613b7f565b61420f836142098354613aba565b8361419d565b6000601f841160018114614243576000851561422b5750838201355b600019600387901b1c1916600186901b178355611ba6565b600083815260209020601f19861690835b828110156142745786850135825560209485019460019092019101614254565b50868210156142915760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b634e487b7160e01b600052603160045260246000fdfe71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb1f5daf52b2e562ec522fee04d65ed2e6335332dc3fcc8634e2de1ced482b41cfdac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077bd6414e0e326510fe0fa0469f2640d922d86657b6847fb93acf969a40b1d5c098a264697066735822122086844b8bf2c94f450be221f7ebdcc3db603fb7f31150e27dfc7637fb8c46d87b64736f6c634300081b0033",
}

// IntentManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IntentManagerMetaData.ABI instead.
var IntentManagerABI = IntentManagerMetaData.ABI

// IntentManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IntentManagerMetaData.Bin instead.
var IntentManagerBin = IntentManagerMetaData.Bin

// DeployIntentManager deploys a new Ethereum contract, binding an instance of IntentManager to it.
func DeployIntentManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IntentManager, error) {
	parsed, err := IntentManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IntentManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IntentManager{IntentManagerCaller: IntentManagerCaller{contract: contract}, IntentManagerTransactor: IntentManagerTransactor{contract: contract}, IntentManagerFilterer: IntentManagerFilterer{contract: contract}}, nil
}

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
