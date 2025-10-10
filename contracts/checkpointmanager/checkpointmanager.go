// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package checkpointmanager

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

// DataStructuresCheckpointHeader is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresCheckpointHeader struct {
	SubnetId             [32]byte
	Epoch                uint64
	ParentCpHash         [32]byte
	Timestamp            *big.Int
	Version              uint32
	ParamsHash           [32]byte
	Roots                DataStructuresCommitmentRoots
	DaCommitments        []DataStructuresDACommitment
	EpochSlot            DataStructuresEpochSlot
	Stats                []byte
	AuxHash              [32]byte
	AssignmentsRoot      [32]byte
	ValidationCommitment [32]byte
	PolicyRoot           [32]byte
}

// DataStructuresCheckpointState is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresCheckpointState struct {
	Header           DataStructuresCheckpointHeader
	Status           uint8
	SubmittedAt      *big.Int
	FinalizedAt      *big.Int
	Disputes         [][32]byte
	CumulativeWeight *big.Int
}

// DataStructuresCommitmentRoots is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresCommitmentRoots struct {
	AgentRoot        [32]byte
	AgentServiceRoot [32]byte
	RankRoot         [32]byte
	MetricsRoot      [32]byte
	DataUsageRoot    [32]byte
	StateRoot        [32]byte
	EventRoot        [32]byte
	CrossSubnetRoot  [32]byte
}

// DataStructuresDACommitment is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresDACommitment struct {
	Kind          string
	Pointer       string
	SizeHint      *big.Int
	SegmentHashes [][32]byte
	Expiry        *big.Int
}

// DataStructuresEpochSlot is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresEpochSlot struct {
	Epoch uint64
	Slot  uint64
}

// CheckpointManagerMetaData contains all meta data concerning the CheckpointManager contract.
var CheckpointManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining_time\",\"type\":\"uint256\"}],\"name\":\"ChallengeWindowActive\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"CheckpointAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"CheckpointNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"CheckpointNotPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"expected\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"provided\",\"type\":\"uint64\"}],\"name\":\"EpochMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidCheckpoint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidContinuity\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"InvalidEpoch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"UnauthorizedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cp_hash\",\"type\":\"bytes32\"}],\"name\":\"CheckpointFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CheckpointReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cp_hash\",\"type\":\"bytes32\"}],\"name\":\"CheckpointSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_CHALLENGE_WINDOW\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"canFinalizeCheckpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"can_finalize\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parent_cp_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"agent_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agent_service_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rank_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metrics_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"data_usage_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"event_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cross_subnet_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CommitmentRoots\",\"name\":\"roots\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size_hint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"segment_hashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.DACommitment[]\",\"name\":\"da_commitments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.EpochSlot\",\"name\":\"epoch_slot\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stats\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"aux_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignments_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validation_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"policy_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CheckpointHeader\",\"name\":\"header\",\"type\":\"tuple\"}],\"name\":\"computeCheckpointHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"finalizeCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"getCheckpoint\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parent_cp_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"agent_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agent_service_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rank_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metrics_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"data_usage_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"event_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cross_subnet_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CommitmentRoots\",\"name\":\"roots\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size_hint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"segment_hashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.DACommitment[]\",\"name\":\"da_commitments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.EpochSlot\",\"name\":\"epoch_slot\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stats\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"aux_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignments_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validation_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"policy_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CheckpointHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"enumDataStructures.CheckpointStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"submitted_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalized_at\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"disputes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"cumulative_weight\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.CheckpointState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"getCheckpointProof\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getLatestCheckpoint\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parent_cp_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"agent_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agent_service_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rank_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metrics_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"data_usage_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"event_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cross_subnet_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CommitmentRoots\",\"name\":\"roots\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size_hint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"segment_hashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.DACommitment[]\",\"name\":\"da_commitments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.EpochSlot\",\"name\":\"epoch_slot\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stats\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"aux_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignments_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validation_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"policy_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CheckpointHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"enumDataStructures.CheckpointStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"submitted_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalized_at\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"disputes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"cumulative_weight\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.CheckpointState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getLatestFinalized\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parent_cp_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"agent_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agent_service_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rank_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metrics_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"data_usage_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"event_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cross_subnet_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CommitmentRoots\",\"name\":\"roots\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size_hint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"segment_hashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.DACommitment[]\",\"name\":\"da_commitments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.EpochSlot\",\"name\":\"epoch_slot\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stats\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"aux_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignments_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validation_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"policy_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CheckpointHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"enumDataStructures.CheckpointStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"submitted_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalized_at\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"disputes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"cumulative_weight\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.CheckpointState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"revertCheckpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"setGovernanceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parent_cp_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"agent_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agent_service_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rank_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metrics_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"data_usage_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"event_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cross_subnet_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CommitmentRoots\",\"name\":\"roots\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size_hint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"segment_hashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.DACommitment[]\",\"name\":\"da_commitments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.EpochSlot\",\"name\":\"epoch_slot\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stats\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"aux_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignments_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validation_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"policy_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CheckpointHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"name\":\"submitCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"cp_hash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"}],\"name\":\"verifyCheckpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"parent_cp_hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"agent_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agent_service_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"rank_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"metrics_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"data_usage_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"state_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"event_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cross_subnet_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CommitmentRoots\",\"name\":\"roots\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pointer\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"size_hint\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"segment_hashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.DACommitment[]\",\"name\":\"da_commitments\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"epoch\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"slot\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.EpochSlot\",\"name\":\"epoch_slot\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"stats\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"aux_hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"assignments_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"validation_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"policy_root\",\"type\":\"bytes32\"}],\"internalType\":\"structDataStructures.CheckpointHeader\",\"name\":\"header\",\"type\":\"tuple\"}],\"name\":\"verifyCheckpointContinuity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CheckpointManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CheckpointManagerMetaData.ABI instead.
var CheckpointManagerABI = CheckpointManagerMetaData.ABI

// CheckpointManager is an auto generated Go binding around an Ethereum contract.
type CheckpointManager struct {
	CheckpointManagerCaller     // Read-only binding to the contract
	CheckpointManagerTransactor // Write-only binding to the contract
	CheckpointManagerFilterer   // Log filterer for contract events
}

// CheckpointManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CheckpointManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckpointManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckpointManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckpointManagerSession struct {
	Contract     *CheckpointManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CheckpointManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckpointManagerCallerSession struct {
	Contract *CheckpointManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CheckpointManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckpointManagerTransactorSession struct {
	Contract     *CheckpointManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CheckpointManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CheckpointManagerRaw struct {
	Contract *CheckpointManager // Generic contract binding to access the raw methods on
}

// CheckpointManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckpointManagerCallerRaw struct {
	Contract *CheckpointManagerCaller // Generic read-only contract binding to access the raw methods on
}

// CheckpointManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CheckpointManagerTransactorRaw struct {
	Contract *CheckpointManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCheckpointManager creates a new instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManager(address common.Address, backend bind.ContractBackend) (*CheckpointManager, error) {
	contract, err := bindCheckpointManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CheckpointManager{CheckpointManagerCaller: CheckpointManagerCaller{contract: contract}, CheckpointManagerTransactor: CheckpointManagerTransactor{contract: contract}, CheckpointManagerFilterer: CheckpointManagerFilterer{contract: contract}}, nil
}

// NewCheckpointManagerCaller creates a new read-only instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManagerCaller(address common.Address, caller bind.ContractCaller) (*CheckpointManagerCaller, error) {
	contract, err := bindCheckpointManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerCaller{contract: contract}, nil
}

// NewCheckpointManagerTransactor creates a new write-only instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*CheckpointManagerTransactor, error) {
	contract, err := bindCheckpointManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerTransactor{contract: contract}, nil
}

// NewCheckpointManagerFilterer creates a new log filterer instance of CheckpointManager, bound to a specific deployed contract.
func NewCheckpointManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*CheckpointManagerFilterer, error) {
	contract, err := bindCheckpointManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerFilterer{contract: contract}, nil
}

// bindCheckpointManager binds a generic wrapper to an already deployed contract.
func bindCheckpointManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CheckpointManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointManager *CheckpointManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointManager.Contract.CheckpointManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointManager *CheckpointManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointManager.Contract.CheckpointManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointManager *CheckpointManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointManager.Contract.CheckpointManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CheckpointManager *CheckpointManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CheckpointManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CheckpointManager *CheckpointManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CheckpointManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CheckpointManager *CheckpointManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CheckpointManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CheckpointManager.Contract.DEFAULTADMINROLE(&_CheckpointManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CheckpointManager.Contract.DEFAULTADMINROLE(&_CheckpointManager.CallOpts)
}

// DEFAULTCHALLENGEWINDOW is a free data retrieval call binding the contract method 0xda493740.
//
// Solidity: function DEFAULT_CHALLENGE_WINDOW() view returns(uint128)
func (_CheckpointManager *CheckpointManagerCaller) DEFAULTCHALLENGEWINDOW(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "DEFAULT_CHALLENGE_WINDOW")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTCHALLENGEWINDOW is a free data retrieval call binding the contract method 0xda493740.
//
// Solidity: function DEFAULT_CHALLENGE_WINDOW() view returns(uint128)
func (_CheckpointManager *CheckpointManagerSession) DEFAULTCHALLENGEWINDOW() (*big.Int, error) {
	return _CheckpointManager.Contract.DEFAULTCHALLENGEWINDOW(&_CheckpointManager.CallOpts)
}

// DEFAULTCHALLENGEWINDOW is a free data retrieval call binding the contract method 0xda493740.
//
// Solidity: function DEFAULT_CHALLENGE_WINDOW() view returns(uint128)
func (_CheckpointManager *CheckpointManagerCallerSession) DEFAULTCHALLENGEWINDOW() (*big.Int, error) {
	return _CheckpointManager.Contract.DEFAULTCHALLENGEWINDOW(&_CheckpointManager.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _CheckpointManager.Contract.GOVERNANCEROLE(&_CheckpointManager.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _CheckpointManager.Contract.GOVERNANCEROLE(&_CheckpointManager.CallOpts)
}

// CanFinalizeCheckpoint is a free data retrieval call binding the contract method 0x3135a81e.
//
// Solidity: function canFinalizeCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(bool can_finalize)
func (_CheckpointManager *CheckpointManagerCaller) CanFinalizeCheckpoint(opts *bind.CallOpts, subnet_id [32]byte, epoch uint64) (bool, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "canFinalizeCheckpoint", subnet_id, epoch)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanFinalizeCheckpoint is a free data retrieval call binding the contract method 0x3135a81e.
//
// Solidity: function canFinalizeCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(bool can_finalize)
func (_CheckpointManager *CheckpointManagerSession) CanFinalizeCheckpoint(subnet_id [32]byte, epoch uint64) (bool, error) {
	return _CheckpointManager.Contract.CanFinalizeCheckpoint(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// CanFinalizeCheckpoint is a free data retrieval call binding the contract method 0x3135a81e.
//
// Solidity: function canFinalizeCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(bool can_finalize)
func (_CheckpointManager *CheckpointManagerCallerSession) CanFinalizeCheckpoint(subnet_id [32]byte, epoch uint64) (bool, error) {
	return _CheckpointManager.Contract.CanFinalizeCheckpoint(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// ComputeCheckpointHash is a free data retrieval call binding the contract method 0x159cedb9.
//
// Solidity: function computeCheckpointHash((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header) pure returns(bytes32 hash)
func (_CheckpointManager *CheckpointManagerCaller) ComputeCheckpointHash(opts *bind.CallOpts, header DataStructuresCheckpointHeader) ([32]byte, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "computeCheckpointHash", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ComputeCheckpointHash is a free data retrieval call binding the contract method 0x159cedb9.
//
// Solidity: function computeCheckpointHash((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header) pure returns(bytes32 hash)
func (_CheckpointManager *CheckpointManagerSession) ComputeCheckpointHash(header DataStructuresCheckpointHeader) ([32]byte, error) {
	return _CheckpointManager.Contract.ComputeCheckpointHash(&_CheckpointManager.CallOpts, header)
}

// ComputeCheckpointHash is a free data retrieval call binding the contract method 0x159cedb9.
//
// Solidity: function computeCheckpointHash((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header) pure returns(bytes32 hash)
func (_CheckpointManager *CheckpointManagerCallerSession) ComputeCheckpointHash(header DataStructuresCheckpointHeader) ([32]byte, error) {
	return _CheckpointManager.Contract.ComputeCheckpointHash(&_CheckpointManager.CallOpts, header)
}

// GetCheckpoint is a free data retrieval call binding the contract method 0x4cf0adf6.
//
// Solidity: function getCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerCaller) GetCheckpoint(opts *bind.CallOpts, subnet_id [32]byte, epoch uint64) (DataStructuresCheckpointState, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getCheckpoint", subnet_id, epoch)

	if err != nil {
		return *new(DataStructuresCheckpointState), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresCheckpointState)).(*DataStructuresCheckpointState)

	return out0, err

}

// GetCheckpoint is a free data retrieval call binding the contract method 0x4cf0adf6.
//
// Solidity: function getCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerSession) GetCheckpoint(subnet_id [32]byte, epoch uint64) (DataStructuresCheckpointState, error) {
	return _CheckpointManager.Contract.GetCheckpoint(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// GetCheckpoint is a free data retrieval call binding the contract method 0x4cf0adf6.
//
// Solidity: function getCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerCallerSession) GetCheckpoint(subnet_id [32]byte, epoch uint64) (DataStructuresCheckpointState, error) {
	return _CheckpointManager.Contract.GetCheckpoint(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// GetCheckpointProof is a free data retrieval call binding the contract method 0x80333bb9.
//
// Solidity: function getCheckpointProof(bytes32 subnet_id, uint64 epoch) view returns(bytes proof)
func (_CheckpointManager *CheckpointManagerCaller) GetCheckpointProof(opts *bind.CallOpts, subnet_id [32]byte, epoch uint64) ([]byte, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getCheckpointProof", subnet_id, epoch)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCheckpointProof is a free data retrieval call binding the contract method 0x80333bb9.
//
// Solidity: function getCheckpointProof(bytes32 subnet_id, uint64 epoch) view returns(bytes proof)
func (_CheckpointManager *CheckpointManagerSession) GetCheckpointProof(subnet_id [32]byte, epoch uint64) ([]byte, error) {
	return _CheckpointManager.Contract.GetCheckpointProof(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// GetCheckpointProof is a free data retrieval call binding the contract method 0x80333bb9.
//
// Solidity: function getCheckpointProof(bytes32 subnet_id, uint64 epoch) view returns(bytes proof)
func (_CheckpointManager *CheckpointManagerCallerSession) GetCheckpointProof(subnet_id [32]byte, epoch uint64) ([]byte, error) {
	return _CheckpointManager.Contract.GetCheckpointProof(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x181b4e65.
//
// Solidity: function getLatestCheckpoint(bytes32 subnet_id) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerCaller) GetLatestCheckpoint(opts *bind.CallOpts, subnet_id [32]byte) (DataStructuresCheckpointState, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getLatestCheckpoint", subnet_id)

	if err != nil {
		return *new(DataStructuresCheckpointState), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresCheckpointState)).(*DataStructuresCheckpointState)

	return out0, err

}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x181b4e65.
//
// Solidity: function getLatestCheckpoint(bytes32 subnet_id) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerSession) GetLatestCheckpoint(subnet_id [32]byte) (DataStructuresCheckpointState, error) {
	return _CheckpointManager.Contract.GetLatestCheckpoint(&_CheckpointManager.CallOpts, subnet_id)
}

// GetLatestCheckpoint is a free data retrieval call binding the contract method 0x181b4e65.
//
// Solidity: function getLatestCheckpoint(bytes32 subnet_id) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerCallerSession) GetLatestCheckpoint(subnet_id [32]byte) (DataStructuresCheckpointState, error) {
	return _CheckpointManager.Contract.GetLatestCheckpoint(&_CheckpointManager.CallOpts, subnet_id)
}

// GetLatestFinalized is a free data retrieval call binding the contract method 0x87087e69.
//
// Solidity: function getLatestFinalized(bytes32 subnet_id) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerCaller) GetLatestFinalized(opts *bind.CallOpts, subnet_id [32]byte) (DataStructuresCheckpointState, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getLatestFinalized", subnet_id)

	if err != nil {
		return *new(DataStructuresCheckpointState), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresCheckpointState)).(*DataStructuresCheckpointState)

	return out0, err

}

// GetLatestFinalized is a free data retrieval call binding the contract method 0x87087e69.
//
// Solidity: function getLatestFinalized(bytes32 subnet_id) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerSession) GetLatestFinalized(subnet_id [32]byte) (DataStructuresCheckpointState, error) {
	return _CheckpointManager.Contract.GetLatestFinalized(&_CheckpointManager.CallOpts, subnet_id)
}

// GetLatestFinalized is a free data retrieval call binding the contract method 0x87087e69.
//
// Solidity: function getLatestFinalized(bytes32 subnet_id) view returns(((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32),uint8,uint256,uint256,bytes32[],uint256) state)
func (_CheckpointManager *CheckpointManagerCallerSession) GetLatestFinalized(subnet_id [32]byte) (DataStructuresCheckpointState, error) {
	return _CheckpointManager.Contract.GetLatestFinalized(&_CheckpointManager.CallOpts, subnet_id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CheckpointManager *CheckpointManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CheckpointManager.Contract.GetRoleAdmin(&_CheckpointManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CheckpointManager *CheckpointManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CheckpointManager.Contract.GetRoleAdmin(&_CheckpointManager.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CheckpointManager *CheckpointManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CheckpointManager *CheckpointManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CheckpointManager.Contract.HasRole(&_CheckpointManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CheckpointManager *CheckpointManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CheckpointManager.Contract.HasRole(&_CheckpointManager.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CheckpointManager *CheckpointManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CheckpointManager *CheckpointManagerSession) Paused() (bool, error) {
	return _CheckpointManager.Contract.Paused(&_CheckpointManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CheckpointManager *CheckpointManagerCallerSession) Paused() (bool, error) {
	return _CheckpointManager.Contract.Paused(&_CheckpointManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CheckpointManager *CheckpointManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CheckpointManager *CheckpointManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CheckpointManager.Contract.SupportsInterface(&_CheckpointManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CheckpointManager *CheckpointManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CheckpointManager.Contract.SupportsInterface(&_CheckpointManager.CallOpts, interfaceId)
}

// VerifyCheckpoint is a free data retrieval call binding the contract method 0x6d9949ab.
//
// Solidity: function verifyCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(bool valid, string reason)
func (_CheckpointManager *CheckpointManagerCaller) VerifyCheckpoint(opts *bind.CallOpts, subnet_id [32]byte, epoch uint64) (struct {
	Valid  bool
	Reason string
}, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "verifyCheckpoint", subnet_id, epoch)

	outstruct := new(struct {
		Valid  bool
		Reason string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Valid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Reason = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// VerifyCheckpoint is a free data retrieval call binding the contract method 0x6d9949ab.
//
// Solidity: function verifyCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(bool valid, string reason)
func (_CheckpointManager *CheckpointManagerSession) VerifyCheckpoint(subnet_id [32]byte, epoch uint64) (struct {
	Valid  bool
	Reason string
}, error) {
	return _CheckpointManager.Contract.VerifyCheckpoint(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// VerifyCheckpoint is a free data retrieval call binding the contract method 0x6d9949ab.
//
// Solidity: function verifyCheckpoint(bytes32 subnet_id, uint64 epoch) view returns(bool valid, string reason)
func (_CheckpointManager *CheckpointManagerCallerSession) VerifyCheckpoint(subnet_id [32]byte, epoch uint64) (struct {
	Valid  bool
	Reason string
}, error) {
	return _CheckpointManager.Contract.VerifyCheckpoint(&_CheckpointManager.CallOpts, subnet_id, epoch)
}

// VerifyCheckpointContinuity is a free data retrieval call binding the contract method 0x2b37c0ba.
//
// Solidity: function verifyCheckpointContinuity(bytes32 subnet_id, (bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header) view returns(bool valid)
func (_CheckpointManager *CheckpointManagerCaller) VerifyCheckpointContinuity(opts *bind.CallOpts, subnet_id [32]byte, header DataStructuresCheckpointHeader) (bool, error) {
	var out []interface{}
	err := _CheckpointManager.contract.Call(opts, &out, "verifyCheckpointContinuity", subnet_id, header)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCheckpointContinuity is a free data retrieval call binding the contract method 0x2b37c0ba.
//
// Solidity: function verifyCheckpointContinuity(bytes32 subnet_id, (bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header) view returns(bool valid)
func (_CheckpointManager *CheckpointManagerSession) VerifyCheckpointContinuity(subnet_id [32]byte, header DataStructuresCheckpointHeader) (bool, error) {
	return _CheckpointManager.Contract.VerifyCheckpointContinuity(&_CheckpointManager.CallOpts, subnet_id, header)
}

// VerifyCheckpointContinuity is a free data retrieval call binding the contract method 0x2b37c0ba.
//
// Solidity: function verifyCheckpointContinuity(bytes32 subnet_id, (bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header) view returns(bool valid)
func (_CheckpointManager *CheckpointManagerCallerSession) VerifyCheckpointContinuity(subnet_id [32]byte, header DataStructuresCheckpointHeader) (bool, error) {
	return _CheckpointManager.Contract.VerifyCheckpointContinuity(&_CheckpointManager.CallOpts, subnet_id, header)
}

// FinalizeCheckpoint is a paid mutator transaction binding the contract method 0x042b21f8.
//
// Solidity: function finalizeCheckpoint(bytes32 subnet_id, uint64 epoch) returns()
func (_CheckpointManager *CheckpointManagerTransactor) FinalizeCheckpoint(opts *bind.TransactOpts, subnet_id [32]byte, epoch uint64) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "finalizeCheckpoint", subnet_id, epoch)
}

// FinalizeCheckpoint is a paid mutator transaction binding the contract method 0x042b21f8.
//
// Solidity: function finalizeCheckpoint(bytes32 subnet_id, uint64 epoch) returns()
func (_CheckpointManager *CheckpointManagerSession) FinalizeCheckpoint(subnet_id [32]byte, epoch uint64) (*types.Transaction, error) {
	return _CheckpointManager.Contract.FinalizeCheckpoint(&_CheckpointManager.TransactOpts, subnet_id, epoch)
}

// FinalizeCheckpoint is a paid mutator transaction binding the contract method 0x042b21f8.
//
// Solidity: function finalizeCheckpoint(bytes32 subnet_id, uint64 epoch) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) FinalizeCheckpoint(subnet_id [32]byte, epoch uint64) (*types.Transaction, error) {
	return _CheckpointManager.Contract.FinalizeCheckpoint(&_CheckpointManager.TransactOpts, subnet_id, epoch)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CheckpointManager *CheckpointManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CheckpointManager *CheckpointManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.GrantRole(&_CheckpointManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.GrantRole(&_CheckpointManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address registry) returns()
func (_CheckpointManager *CheckpointManagerTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, registry common.Address) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "initialize", admin, registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address registry) returns()
func (_CheckpointManager *CheckpointManagerSession) Initialize(admin common.Address, registry common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.Initialize(&_CheckpointManager.TransactOpts, admin, registry)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address registry) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) Initialize(admin common.Address, registry common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.Initialize(&_CheckpointManager.TransactOpts, admin, registry)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CheckpointManager *CheckpointManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CheckpointManager *CheckpointManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.RenounceRole(&_CheckpointManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.RenounceRole(&_CheckpointManager.TransactOpts, role, callerConfirmation)
}

// RevertCheckpoint is a paid mutator transaction binding the contract method 0x90b97438.
//
// Solidity: function revertCheckpoint(bytes32 subnet_id, uint64 epoch, string reason) returns()
func (_CheckpointManager *CheckpointManagerTransactor) RevertCheckpoint(opts *bind.TransactOpts, subnet_id [32]byte, epoch uint64, reason string) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "revertCheckpoint", subnet_id, epoch, reason)
}

// RevertCheckpoint is a paid mutator transaction binding the contract method 0x90b97438.
//
// Solidity: function revertCheckpoint(bytes32 subnet_id, uint64 epoch, string reason) returns()
func (_CheckpointManager *CheckpointManagerSession) RevertCheckpoint(subnet_id [32]byte, epoch uint64, reason string) (*types.Transaction, error) {
	return _CheckpointManager.Contract.RevertCheckpoint(&_CheckpointManager.TransactOpts, subnet_id, epoch, reason)
}

// RevertCheckpoint is a paid mutator transaction binding the contract method 0x90b97438.
//
// Solidity: function revertCheckpoint(bytes32 subnet_id, uint64 epoch, string reason) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) RevertCheckpoint(subnet_id [32]byte, epoch uint64, reason string) (*types.Transaction, error) {
	return _CheckpointManager.Contract.RevertCheckpoint(&_CheckpointManager.TransactOpts, subnet_id, epoch, reason)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CheckpointManager *CheckpointManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CheckpointManager *CheckpointManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.RevokeRole(&_CheckpointManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.RevokeRole(&_CheckpointManager.TransactOpts, role, account)
}

// SetGovernanceRole is a paid mutator transaction binding the contract method 0xe1ba6889.
//
// Solidity: function setGovernanceRole(address governance) returns()
func (_CheckpointManager *CheckpointManagerTransactor) SetGovernanceRole(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "setGovernanceRole", governance)
}

// SetGovernanceRole is a paid mutator transaction binding the contract method 0xe1ba6889.
//
// Solidity: function setGovernanceRole(address governance) returns()
func (_CheckpointManager *CheckpointManagerSession) SetGovernanceRole(governance common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.SetGovernanceRole(&_CheckpointManager.TransactOpts, governance)
}

// SetGovernanceRole is a paid mutator transaction binding the contract method 0xe1ba6889.
//
// Solidity: function setGovernanceRole(address governance) returns()
func (_CheckpointManager *CheckpointManagerTransactorSession) SetGovernanceRole(governance common.Address) (*types.Transaction, error) {
	return _CheckpointManager.Contract.SetGovernanceRole(&_CheckpointManager.TransactOpts, governance)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x9dec4a95.
//
// Solidity: function submitCheckpoint((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header, address[] validators, bytes[] signatures) returns(bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerTransactor) SubmitCheckpoint(opts *bind.TransactOpts, header DataStructuresCheckpointHeader, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _CheckpointManager.contract.Transact(opts, "submitCheckpoint", header, validators, signatures)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x9dec4a95.
//
// Solidity: function submitCheckpoint((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header, address[] validators, bytes[] signatures) returns(bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerSession) SubmitCheckpoint(header DataStructuresCheckpointHeader, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _CheckpointManager.Contract.SubmitCheckpoint(&_CheckpointManager.TransactOpts, header, validators, signatures)
}

// SubmitCheckpoint is a paid mutator transaction binding the contract method 0x9dec4a95.
//
// Solidity: function submitCheckpoint((bytes32,uint64,bytes32,uint256,uint32,bytes32,(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32,bytes32),(string,string,uint256,bytes32[],uint256)[],(uint64,uint64),bytes,bytes32,bytes32,bytes32,bytes32) header, address[] validators, bytes[] signatures) returns(bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerTransactorSession) SubmitCheckpoint(header DataStructuresCheckpointHeader, validators []common.Address, signatures [][]byte) (*types.Transaction, error) {
	return _CheckpointManager.Contract.SubmitCheckpoint(&_CheckpointManager.TransactOpts, header, validators, signatures)
}

// CheckpointManagerCheckpointFinalizedIterator is returned from FilterCheckpointFinalized and is used to iterate over the raw logs and unpacked data for CheckpointFinalized events raised by the CheckpointManager contract.
type CheckpointManagerCheckpointFinalizedIterator struct {
	Event *CheckpointManagerCheckpointFinalized // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerCheckpointFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerCheckpointFinalized)
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
		it.Event = new(CheckpointManagerCheckpointFinalized)
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
func (it *CheckpointManagerCheckpointFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerCheckpointFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerCheckpointFinalized represents a CheckpointFinalized event raised by the CheckpointManager contract.
type CheckpointManagerCheckpointFinalized struct {
	SubnetId [32]byte
	Epoch    uint64
	CpHash   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCheckpointFinalized is a free log retrieval operation binding the contract event 0x0c667fed5b7cb98ff28e05e7b6f653e237cbf3e8b76d3ff481af7bd6f7e4a4ac.
//
// Solidity: event CheckpointFinalized(bytes32 indexed subnet_id, uint64 epoch, bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerFilterer) FilterCheckpointFinalized(opts *bind.FilterOpts, subnet_id [][32]byte) (*CheckpointManagerCheckpointFinalizedIterator, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "CheckpointFinalized", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerCheckpointFinalizedIterator{contract: _CheckpointManager.contract, event: "CheckpointFinalized", logs: logs, sub: sub}, nil
}

// WatchCheckpointFinalized is a free log subscription operation binding the contract event 0x0c667fed5b7cb98ff28e05e7b6f653e237cbf3e8b76d3ff481af7bd6f7e4a4ac.
//
// Solidity: event CheckpointFinalized(bytes32 indexed subnet_id, uint64 epoch, bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerFilterer) WatchCheckpointFinalized(opts *bind.WatchOpts, sink chan<- *CheckpointManagerCheckpointFinalized, subnet_id [][32]byte) (event.Subscription, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "CheckpointFinalized", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerCheckpointFinalized)
				if err := _CheckpointManager.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
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

// ParseCheckpointFinalized is a log parse operation binding the contract event 0x0c667fed5b7cb98ff28e05e7b6f653e237cbf3e8b76d3ff481af7bd6f7e4a4ac.
//
// Solidity: event CheckpointFinalized(bytes32 indexed subnet_id, uint64 epoch, bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerFilterer) ParseCheckpointFinalized(log types.Log) (*CheckpointManagerCheckpointFinalized, error) {
	event := new(CheckpointManagerCheckpointFinalized)
	if err := _CheckpointManager.contract.UnpackLog(event, "CheckpointFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerCheckpointRevertedIterator is returned from FilterCheckpointReverted and is used to iterate over the raw logs and unpacked data for CheckpointReverted events raised by the CheckpointManager contract.
type CheckpointManagerCheckpointRevertedIterator struct {
	Event *CheckpointManagerCheckpointReverted // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerCheckpointRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerCheckpointReverted)
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
		it.Event = new(CheckpointManagerCheckpointReverted)
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
func (it *CheckpointManagerCheckpointRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerCheckpointRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerCheckpointReverted represents a CheckpointReverted event raised by the CheckpointManager contract.
type CheckpointManagerCheckpointReverted struct {
	SubnetId [32]byte
	Epoch    uint64
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCheckpointReverted is a free log retrieval operation binding the contract event 0x9207c8713e9d56fa646c111dd7918bf65a826668f24cc55d424c700598e6558b.
//
// Solidity: event CheckpointReverted(bytes32 indexed subnet_id, uint64 epoch, string reason)
func (_CheckpointManager *CheckpointManagerFilterer) FilterCheckpointReverted(opts *bind.FilterOpts, subnet_id [][32]byte) (*CheckpointManagerCheckpointRevertedIterator, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "CheckpointReverted", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerCheckpointRevertedIterator{contract: _CheckpointManager.contract, event: "CheckpointReverted", logs: logs, sub: sub}, nil
}

// WatchCheckpointReverted is a free log subscription operation binding the contract event 0x9207c8713e9d56fa646c111dd7918bf65a826668f24cc55d424c700598e6558b.
//
// Solidity: event CheckpointReverted(bytes32 indexed subnet_id, uint64 epoch, string reason)
func (_CheckpointManager *CheckpointManagerFilterer) WatchCheckpointReverted(opts *bind.WatchOpts, sink chan<- *CheckpointManagerCheckpointReverted, subnet_id [][32]byte) (event.Subscription, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "CheckpointReverted", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerCheckpointReverted)
				if err := _CheckpointManager.contract.UnpackLog(event, "CheckpointReverted", log); err != nil {
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

// ParseCheckpointReverted is a log parse operation binding the contract event 0x9207c8713e9d56fa646c111dd7918bf65a826668f24cc55d424c700598e6558b.
//
// Solidity: event CheckpointReverted(bytes32 indexed subnet_id, uint64 epoch, string reason)
func (_CheckpointManager *CheckpointManagerFilterer) ParseCheckpointReverted(log types.Log) (*CheckpointManagerCheckpointReverted, error) {
	event := new(CheckpointManagerCheckpointReverted)
	if err := _CheckpointManager.contract.UnpackLog(event, "CheckpointReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerCheckpointSubmittedIterator is returned from FilterCheckpointSubmitted and is used to iterate over the raw logs and unpacked data for CheckpointSubmitted events raised by the CheckpointManager contract.
type CheckpointManagerCheckpointSubmittedIterator struct {
	Event *CheckpointManagerCheckpointSubmitted // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerCheckpointSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerCheckpointSubmitted)
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
		it.Event = new(CheckpointManagerCheckpointSubmitted)
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
func (it *CheckpointManagerCheckpointSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerCheckpointSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerCheckpointSubmitted represents a CheckpointSubmitted event raised by the CheckpointManager contract.
type CheckpointManagerCheckpointSubmitted struct {
	SubnetId [32]byte
	Epoch    uint64
	CpHash   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCheckpointSubmitted is a free log retrieval operation binding the contract event 0x53cae6135bd446f0fcf621e412aabe4254b325eeab58936647f9a4dd6099c41a.
//
// Solidity: event CheckpointSubmitted(bytes32 indexed subnet_id, uint64 epoch, bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerFilterer) FilterCheckpointSubmitted(opts *bind.FilterOpts, subnet_id [][32]byte) (*CheckpointManagerCheckpointSubmittedIterator, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "CheckpointSubmitted", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerCheckpointSubmittedIterator{contract: _CheckpointManager.contract, event: "CheckpointSubmitted", logs: logs, sub: sub}, nil
}

// WatchCheckpointSubmitted is a free log subscription operation binding the contract event 0x53cae6135bd446f0fcf621e412aabe4254b325eeab58936647f9a4dd6099c41a.
//
// Solidity: event CheckpointSubmitted(bytes32 indexed subnet_id, uint64 epoch, bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerFilterer) WatchCheckpointSubmitted(opts *bind.WatchOpts, sink chan<- *CheckpointManagerCheckpointSubmitted, subnet_id [][32]byte) (event.Subscription, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "CheckpointSubmitted", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerCheckpointSubmitted)
				if err := _CheckpointManager.contract.UnpackLog(event, "CheckpointSubmitted", log); err != nil {
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

// ParseCheckpointSubmitted is a log parse operation binding the contract event 0x53cae6135bd446f0fcf621e412aabe4254b325eeab58936647f9a4dd6099c41a.
//
// Solidity: event CheckpointSubmitted(bytes32 indexed subnet_id, uint64 epoch, bytes32 cp_hash)
func (_CheckpointManager *CheckpointManagerFilterer) ParseCheckpointSubmitted(log types.Log) (*CheckpointManagerCheckpointSubmitted, error) {
	event := new(CheckpointManagerCheckpointSubmitted)
	if err := _CheckpointManager.contract.UnpackLog(event, "CheckpointSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CheckpointManager contract.
type CheckpointManagerInitializedIterator struct {
	Event *CheckpointManagerInitialized // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerInitialized)
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
		it.Event = new(CheckpointManagerInitialized)
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
func (it *CheckpointManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerInitialized represents a Initialized event raised by the CheckpointManager contract.
type CheckpointManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CheckpointManager *CheckpointManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*CheckpointManagerInitializedIterator, error) {

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerInitializedIterator{contract: _CheckpointManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CheckpointManager *CheckpointManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CheckpointManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerInitialized)
				if err := _CheckpointManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CheckpointManager *CheckpointManagerFilterer) ParseInitialized(log types.Log) (*CheckpointManagerInitialized, error) {
	event := new(CheckpointManagerInitialized)
	if err := _CheckpointManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CheckpointManager contract.
type CheckpointManagerPausedIterator struct {
	Event *CheckpointManagerPaused // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerPaused)
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
		it.Event = new(CheckpointManagerPaused)
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
func (it *CheckpointManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerPaused represents a Paused event raised by the CheckpointManager contract.
type CheckpointManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CheckpointManager *CheckpointManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*CheckpointManagerPausedIterator, error) {

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerPausedIterator{contract: _CheckpointManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CheckpointManager *CheckpointManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CheckpointManagerPaused) (event.Subscription, error) {

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerPaused)
				if err := _CheckpointManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CheckpointManager *CheckpointManagerFilterer) ParsePaused(log types.Log) (*CheckpointManagerPaused, error) {
	event := new(CheckpointManagerPaused)
	if err := _CheckpointManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CheckpointManager contract.
type CheckpointManagerRoleAdminChangedIterator struct {
	Event *CheckpointManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerRoleAdminChanged)
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
		it.Event = new(CheckpointManagerRoleAdminChanged)
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
func (it *CheckpointManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerRoleAdminChanged represents a RoleAdminChanged event raised by the CheckpointManager contract.
type CheckpointManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CheckpointManager *CheckpointManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CheckpointManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerRoleAdminChangedIterator{contract: _CheckpointManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CheckpointManager *CheckpointManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CheckpointManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerRoleAdminChanged)
				if err := _CheckpointManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CheckpointManager *CheckpointManagerFilterer) ParseRoleAdminChanged(log types.Log) (*CheckpointManagerRoleAdminChanged, error) {
	event := new(CheckpointManagerRoleAdminChanged)
	if err := _CheckpointManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CheckpointManager contract.
type CheckpointManagerRoleGrantedIterator struct {
	Event *CheckpointManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerRoleGranted)
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
		it.Event = new(CheckpointManagerRoleGranted)
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
func (it *CheckpointManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerRoleGranted represents a RoleGranted event raised by the CheckpointManager contract.
type CheckpointManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CheckpointManager *CheckpointManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CheckpointManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerRoleGrantedIterator{contract: _CheckpointManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CheckpointManager *CheckpointManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CheckpointManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerRoleGranted)
				if err := _CheckpointManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CheckpointManager *CheckpointManagerFilterer) ParseRoleGranted(log types.Log) (*CheckpointManagerRoleGranted, error) {
	event := new(CheckpointManagerRoleGranted)
	if err := _CheckpointManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CheckpointManager contract.
type CheckpointManagerRoleRevokedIterator struct {
	Event *CheckpointManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerRoleRevoked)
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
		it.Event = new(CheckpointManagerRoleRevoked)
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
func (it *CheckpointManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerRoleRevoked represents a RoleRevoked event raised by the CheckpointManager contract.
type CheckpointManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CheckpointManager *CheckpointManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CheckpointManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerRoleRevokedIterator{contract: _CheckpointManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CheckpointManager *CheckpointManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CheckpointManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerRoleRevoked)
				if err := _CheckpointManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CheckpointManager *CheckpointManagerFilterer) ParseRoleRevoked(log types.Log) (*CheckpointManagerRoleRevoked, error) {
	event := new(CheckpointManagerRoleRevoked)
	if err := _CheckpointManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CheckpointManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CheckpointManager contract.
type CheckpointManagerUnpausedIterator struct {
	Event *CheckpointManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *CheckpointManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CheckpointManagerUnpaused)
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
		it.Event = new(CheckpointManagerUnpaused)
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
func (it *CheckpointManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CheckpointManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CheckpointManagerUnpaused represents a Unpaused event raised by the CheckpointManager contract.
type CheckpointManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CheckpointManager *CheckpointManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CheckpointManagerUnpausedIterator, error) {

	logs, sub, err := _CheckpointManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CheckpointManagerUnpausedIterator{contract: _CheckpointManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CheckpointManager *CheckpointManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CheckpointManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _CheckpointManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CheckpointManagerUnpaused)
				if err := _CheckpointManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_CheckpointManager *CheckpointManagerFilterer) ParseUnpaused(log types.Log) (*CheckpointManagerUnpaused, error) {
	event := new(CheckpointManagerUnpaused)
	if err := _CheckpointManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
