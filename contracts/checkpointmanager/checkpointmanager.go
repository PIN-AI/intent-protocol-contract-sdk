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
	Bin: "0x6080604052348015600f57600080fd5b50615c0a8061001f6000396000f3fe608060405234801561001057600080fd5b506004361061011d5760003560e01c806301ffc9a714610122578063042b21f81461014a578063159cedb91461015f578063181b4e6514610180578063248a9ca3146101a05780632b37c0ba146101b35780632f2ff15d146101c65780633135a81e146101d957806336568abe146101ec578063485cc955146101ff5780634cf0adf6146102125780635c975abb146102255780636d9949ab1461022d57806380333bb91461024e57806387087e691461026e57806390b974381461028157806391d14854146102945780639dec4a95146102a7578063a217fddf146102ba578063d547741f146102c2578063da493740146102d5578063e1ba6889146102f6578063f36c8f5c14610309575b600080fd5b610135610130366004613e05565b61031e565b60405190151581526020015b60405180910390f35b61015d610158366004613e54565b610355565b005b61017261016d366004613e9d565b610503565b604051908152602001610141565b61019361018e366004613ed1565b610519565b60405161014191906141a0565b6101726101ae366004613ed1565b6109e2565b6101356101c1366004614217565b610a02565b61015d6101d4366004614272565b610a1e565b6101356101e7366004613e54565b610a40565b61015d6101fa366004614272565b610ae9565b61015d61020d366004614297565b610b21565b610193610220366004613e54565b610ca5565b610135611169565b61024061023b366004613e54565b611191565b6040516101419291906142c5565b61026161025c366004613e54565b6111ad565b60405161014191906142e0565b61019361027c366004613ed1565b611678565b61015d61028f3660046142f3565b611990565b6101356102a2366004614272565b611a66565b6101726102b53660046143c1565b611a9c565b610172600081565b61015d6102d0366004614272565b611be2565b6102de6104b081565b6040516001600160801b039091168152602001610141565b61015d61030436600461445f565b611bfe565b610172600080516020615bb583398151915281565b60006001600160e01b03198216637965db0b60e01b148061034f57506301ffc9a760e01b6001600160e01b03198316145b92915050565b61035d611c21565b6000828152602081815260408083206001600160401b0385168452909152902054829082906103ac57818160405163034ad14960e51b81526004016103a392919061447c565b60405180910390fd5b6000848152602081815260408083206001600160401b038716845290915290206001601582015460ff1660048111156103e7576103e7614176565b14610409578484604051634617d04160e01b81526004016103a392919061447c565b600061041486611c49565b9050806001600160801b0316826016015461042f91906144a9565b42101561045257604051632a4ef15560e21b8152600060048201526024016103a3565b60158201805460ff19166002179055426017830155600086815260016020908152604080832080546001600160401b038a166001600160401b0319909116179055516104be916104a491869101614646565b604051602081830303815290604052805160209091012090565b9050867f0c667fed5b7cb98ff28e05e7b6f653e237cbf3e8b76d3ff481af7bd6f7e4a4ac87836040516104f2929190614789565b60405180910390a250505050505050565b600061034f826040516020016104a491906149f8565b610521613cef565b6000828152600260209081526040808320548383528184206001600160401b0390911680855292529091205461056e57828160405163034ad14960e51b81526004016103a392919061447c565b6000838152602081815260408083206001600160401b038086168552908352818420825161028081018452815460c0808301918252600184015490941660e0808401919091526002840154610100808501919091526003850154610120850152600485015463ffffffff166101408501526005850154610160850152865190810187526006850154815260078501548189015260088501548188015260098501546060820152600a8501546080820152600b85015460a0820152600c85015495810195909552600d84015490850152610180820193909352600e8201805485518188028101880190965280865291969295879587946101a088019491939284015b8282101561083157838290600052602060002090600502016040518060a00160405290816000820180546106a2906144bc565b80601f01602080910402602001604051908101604052809291908181526020018280546106ce906144bc565b801561071b5780601f106106f05761010080835404028352916020019161071b565b820191906000526020600020905b8154815290600101906020018083116106fe57829003601f168201915b50505050508152602001600182018054610734906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054610760906144bc565b80156107ad5780601f10610782576101008083540402835291602001916107ad565b820191906000526020600020905b81548152906001019060200180831161079057829003601f168201915b50505050508152602001600282015481526020016003820180548060200260200160405190810160405280929190818152602001828054801561080f57602002820191906000526020600020905b8154815260200190600101908083116107fb575b505050505081526020016004820154815250508152602001906001019061066f565b50505090825250604080518082018252600f8401546001600160401b038082168352600160401b9091041660208281019190915283015260108301805491909201919061087d906144bc565b80601f01602080910402602001604051908101604052809291908181526020018280546108a9906144bc565b80156108f65780601f106108cb576101008083540402835291602001916108f6565b820191906000526020600020905b8154815290600101906020018083116108d957829003601f168201915b5050505050815260200160118201548152602001601282015481526020016013820154815260200160148201548152505081526020016015820160009054906101000a900460ff16600481111561094f5761094f614176565b600481111561096057610960614176565b81526020016016820154815260200160178201548152602001601882018054806020026020016040519081016040528092919081815260200182805480156109c757602002820191906000526020600020905b8154815260200190600101908083116109b3575b50505050508152602001601982015481525050915050919050565b6000806109ed611d5d565b60009384526020525050604090206001015490565b6000610a178184610a1285614ef6565b611d81565b9392505050565b610a27826109e2565b610a3081612276565b610a3a8383612283565b50505050565b6000828152602081815260408083206001600160401b0385168452909152812054610a6d5750600061034f565b6000838152602081815260408083206001600160401b038616845290915290206001601582015460ff166004811115610aa857610aa8614176565b14610ab757600091505061034f565b6000610ac285611c49565b9050806001600160801b03168260160154610add91906144a9565b42101595945050505050565b6001600160a01b0381163314610b125760405163334bd91960e11b815260040160405180910390fd5b610b1c8282612324565b505050565b6000610b2b61239c565b805490915060ff600160401b82041615906001600160401b0316600081158015610b525750825b90506000826001600160401b03166001148015610b6e5750303b155b905081158015610b7c575080155b15610b9a5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315610bc357845460ff60401b1916600160401b1785555b6001600160a01b0387161580610be057506001600160a01b038616155b15610bfe5760405163d92e233d60e01b815260040160405180910390fd5b610c066123c5565b610c0e6123c5565b610c166123cd565b610c21600088612283565b50610c3a600080516020615bb583398151915288612283565b50600380546001600160a01b0319166001600160a01b0388161790558315610c9c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b610cad613cef565b6000838152602081815260408083206001600160401b038616845290915290205483908390610cf357818160405163034ad14960e51b81526004016103a392919061447c565b6000858152602081815260408083206001600160401b038089168552908352818420825161028081018452815460c0808301918252600184015490941660e0808401919091526002840154610100808501919091526003850154610120850152600485015463ffffffff166101408501526005850154610160850152865190810187526006850154815260078501548189015260088501548188015260098501546060820152600a8501546080820152600b85015460a0820152600c85015495810195909552600d84015490850152610180820193909352600e8201805485518188028101880190965280865291969295879587946101a088019491939284015b82821015610fb657838290600052602060002090600502016040518060a0016040529081600082018054610e27906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054610e53906144bc565b8015610ea05780601f10610e7557610100808354040283529160200191610ea0565b820191906000526020600020905b815481529060010190602001808311610e8357829003601f168201915b50505050508152602001600182018054610eb9906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054610ee5906144bc565b8015610f325780601f10610f0757610100808354040283529160200191610f32565b820191906000526020600020905b815481529060010190602001808311610f1557829003601f168201915b505050505081526020016002820154815260200160038201805480602002602001604051908101604052809291908181526020018280548015610f9457602002820191906000526020600020905b815481526020019060010190808311610f80575b5050505050815260200160048201548152505081526020019060010190610df4565b50505090825250604080518082018252600f8401546001600160401b038082168352600160401b90910416602082810191909152830152601083018054919092019190611002906144bc565b80601f016020809104026020016040519081016040528092919081815260200182805461102e906144bc565b801561107b5780601f106110505761010080835404028352916020019161107b565b820191906000526020600020905b81548152906001019060200180831161105e57829003601f168201915b5050505050815260200160118201548152602001601282015481526020016013820154815260200160148201548152505081526020016015820160009054906101000a900460ff1660048111156110d4576110d4614176565b60048111156110e5576110e5614176565b815260200160168201548152602001601782015481526020016018820180548060200260200160405190810160405280929190818152602001828054801561114c57602002820191906000526020600020905b815481526020019060010190808311611138575b505050505081526020016019820154815250509250505092915050565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1690565b600060606111a1600085856123dd565b915091505b9250929050565b6000828152602081815260408083206001600160401b0385168452909152902054606090839083906111f657818160405163034ad14960e51b81526004016103a392919061447c565b6000858152602081815260408083206001600160401b038089168552908352818420825161028081018452815460c0808301918252600184015490941660e0808401919091526002840154610100808501919091526003850154610120850152600485015463ffffffff166101408501526005850154610160850152865190810187526006850154815260078501548189015260088501548188015260098501546060820152600a8501546080820152600b85015460a0820152600c85015495810195909552600d84015490850152610180820193909352600e820180548551818802810188019096528086529195929486949386936101a0870193918a9084015b828210156114ba57838290600052602060002090600502016040518060a001604052908160008201805461132b906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054611357906144bc565b80156113a45780601f10611379576101008083540402835291602001916113a4565b820191906000526020600020905b81548152906001019060200180831161138757829003601f168201915b505050505081526020016001820180546113bd906144bc565b80601f01602080910402602001604051908101604052809291908181526020018280546113e9906144bc565b80156114365780601f1061140b57610100808354040283529160200191611436565b820191906000526020600020905b81548152906001019060200180831161141957829003601f168201915b50505050508152602001600282015481526020016003820180548060200260200160405190810160405280929190818152602001828054801561149857602002820191906000526020600020905b815481526020019060010190808311611484575b50505050508152602001600482015481525050815260200190600101906112f8565b50505090825250604080518082018252600f8401546001600160401b038082168352600160401b90910416602082810191909152830152601083018054919092019190611506906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054611532906144bc565b801561157f5780601f106115545761010080835404028352916020019161157f565b820191906000526020600020905b81548152906001019060200180831161156257829003601f168201915b5050505050815260200160118201548152602001601282015481526020016013820154815260200160148201548152505081526020016015820160009054906101000a900460ff1660048111156115d8576115d8614176565b60048111156115e9576115e9614176565b815260200160168201548152602001601782015481526020016018820180548060200260200160405190810160405280929190818152602001828054801561165057602002820191906000526020600020905b81548152602001906001019080831161163c575b50505050508152602001601982015481525050905061166e81612958565b9695505050505050565b611680613cef565b6000828152600160209081526040808320548383528184206001600160401b039091168085529252909120546116cd57828160405163034ad14960e51b81526004016103a392919061447c565b6000838152602081815260408083206001600160401b038086168552908352818420825161028081018452815460c0808301918252600184015490941660e0808401919091526002840154610100808501919091526003850154610120850152600485015463ffffffff166101408501526005850154610160850152865190810187526006850154815260078501548189015260088501548188015260098501546060820152600a8501546080820152600b85015460a0820152600c85015495810195909552600d84015490850152610180820193909352600e8201805485518188028101880190965280865291969295879587946101a088019491939284015b8282101561083157838290600052602060002090600502016040518060a0016040529081600082018054611801906144bc565b80601f016020809104026020016040519081016040528092919081815260200182805461182d906144bc565b801561187a5780601f1061184f5761010080835404028352916020019161187a565b820191906000526020600020905b81548152906001019060200180831161185d57829003601f168201915b50505050508152602001600182018054611893906144bc565b80601f01602080910402602001604051908101604052809291908181526020018280546118bf906144bc565b801561190c5780601f106118e15761010080835404028352916020019161190c565b820191906000526020600020905b8154815290600101906020018083116118ef57829003601f168201915b50505050508152602001600282015481526020016003820180548060200260200160405190810160405280929190818152602001828054801561196e57602002820191906000526020600020905b81548152602001906001019080831161195a575b50505050508152602001600482015481525050815260200190600101906117ce565b600080516020615bb58339815191526119a881612276565b6000858152602081815260408083206001600160401b0388168452909152902054859085906119ee57818160405163034ad14960e51b81526004016103a392919061447c565b6000878152602081815260408083206001600160401b038a1684529091529081902060158101805460ff19166004179055905188907f9207c8713e9d56fa646c111dd7918bf65a826668f24cc55d424c700598e6558b90611a54908a908a908a90615009565b60405180910390a25050505050505050565b600080611a71611d5d565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b6000611aa6611c21565b611aae612997565b611abb86868686866129cd565b60035460405163d5a84a1160e01b8152873560048201526000916001600160a01b03169063d5a84a1190602401602060405180830381865afa158015611b05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b299190615039565b90506001600160a01b038116611b7c57604051636337886960e11b81526020600482015260176024820152765375626e657420636f6e7472616374206d697373696e6760481b60448201526064016103a3565b6000611b8782612aed565b9050611b968835888884612b6a565b611b9f88612d0a565b611ba888612e86565b6000611bb389613020565b9050611bc381848a8a8a8a61321c565b611bcc896133fc565b9350505050611bd9613530565b95945050505050565b611beb826109e2565b611bf481612276565b610a3a8383612324565b6000611c0981612276565b610b1c600080516020615bb583398151915283612283565b611c29611169565b15611c475760405163d93c066560e01b815260040160405180910390fd5b565b60035460405163d5a84a1160e01b81526004810183905260009182916001600160a01b039091169063d5a84a1190602401602060405180830381865afa158015611c97573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cbb9190615039565b90506001600160a01b038116611cd557506104b092915050565b6000816001600160a01b031663d2b4353f6040518163ffffffff1660e01b8152600401606060405180830381865afa158015611d15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d39919061506d565b8051935090506001600160801b038316600003611d56576104b092505b5050919050565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b600081602001516001600160401b0316600003611da057506001610a17565b6000838152602085815260408220908401518290611dc0906001906150de565b6001600160401b03908116825260208083019390935260409182016000908120835161028081018552815460c0808301918252600184015490951660e0808401919091526002840154610100808501919091526003850154610120850152600485015463ffffffff16610140850152600585015461016085015287519081018852600685015481526007850154818a015260088501548189015260098501546060820152600a8501546080820152600b85015460a0820152600c85015496810196909652600d84015490860152610180820194909452600e820180548651818902810189019097528087529196929587959487946101a0880194929392919084015b8282101561208457838290600052602060002090600502016040518060a0016040529081600082018054611ef5906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054611f21906144bc565b8015611f6e5780601f10611f4357610100808354040283529160200191611f6e565b820191906000526020600020905b815481529060010190602001808311611f5157829003601f168201915b50505050508152602001600182018054611f87906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054611fb3906144bc565b80156120005780601f10611fd557610100808354040283529160200191612000565b820191906000526020600020905b815481529060010190602001808311611fe357829003601f168201915b50505050508152602001600282015481526020016003820180548060200260200160405190810160405280929190818152602001828054801561206257602002820191906000526020600020905b81548152602001906001019080831161204e575b5050505050815260200160048201548152505081526020019060010190611ec2565b50505090825250604080518082018252600f8401546001600160401b038082168352600160401b909104166020828101919091528301526010830180549190920191906120d0906144bc565b80601f01602080910402602001604051908101604052809291908181526020018280546120fc906144bc565b80156121495780601f1061211e57610100808354040283529160200191612149565b820191906000526020600020905b81548152906001019060200180831161212c57829003601f168201915b5050505050815260200160118201548152602001601282015481526020016013820154815260200160148201548152505081526020016015820160009054906101000a900460ff1660048111156121a2576121a2614176565b60048111156121b3576121b3614176565b815260200160168201548152602001601782015481526020016018820180548060200260200160405190810160405280929190818152602001828054801561221a57602002820191906000526020600020905b815481526020019060010190808311612206575b505050918352505060199190910154602090910152805151909150612243576000915050610a17565b8051602001516122549060016150fd565b6001600160401b031683602001516001600160401b0316149150509392505050565b6122808133613541565b50565b60008061228e611d5d565b905061229a8484611a66565b61231a576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556122d03390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061034f565b600091505061034f565b60008061232f611d5d565b905061233b8484611a66565b1561231a576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061034f565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061034f565b611c4761357e565b6123d561357e565b611c476135a3565b6000828152602084815260408083206001600160401b038516845290915281205460609061243a57505060408051808201909152601481527310da1958dadc1bda5b9d081b9bdd08199bdd5b9960621b6020820152600090612950565b6000848152602086815260408083206001600160401b038088168552908352818420825161028081018452815460c0808301918252600184015490941660e0808401919091526002840154610100808501919091526003850154610120850152600485015463ffffffff166101408501526005850154610160850152865190810187526006850154815260078501548189015260088501548188015260098501546060820152600a8501546080820152600b85015460a0820152600c85015495810195909552600d84015490850152610180820193909352600e820180548551818802810188019096528086529195929486949386936101a0870193918a9084015b828210156126fe57838290600052602060002090600502016040518060a001604052908160008201805461256f906144bc565b80601f016020809104026020016040519081016040528092919081815260200182805461259b906144bc565b80156125e85780601f106125bd576101008083540402835291602001916125e8565b820191906000526020600020905b8154815290600101906020018083116125cb57829003601f168201915b50505050508152602001600182018054612601906144bc565b80601f016020809104026020016040519081016040528092919081815260200182805461262d906144bc565b801561267a5780601f1061264f5761010080835404028352916020019161267a565b820191906000526020600020905b81548152906001019060200180831161265d57829003601f168201915b5050505050815260200160028201548152602001600382018054806020026020016040519081016040528092919081815260200182805480156126dc57602002820191906000526020600020905b8154815260200190600101908083116126c8575b505050505081526020016004820154815250508152602001906001019061253c565b50505090825250604080518082018252600f8401546001600160401b038082168352600160401b9091041660208281019190915283015260108301805491909201919061274a906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054612776906144bc565b80156127c35780601f10612798576101008083540402835291602001916127c3565b820191906000526020600020905b8154815290600101906020018083116127a657829003601f168201915b5050505050815260200160118201548152602001601282015481526020016013820154815260200160148201548152505081526020016015820160009054906101000a900460ff16600481111561281c5761281c614176565b600481111561282d5761282d614176565b815260200160168201548152602001601782015481526020016018820180548060200260200160405190810160405280929190818152602001828054801561289457602002820191906000526020600020905b815481526020019060010190808311612880575b50505091835250506019919091015460209091015280515190915085146128e95750506040805180820190915260128152710a6eac4dccae840928840dad2e6dac2e8c6d60731b602082015260009150612950565b836001600160401b03168160000151602001516001600160401b03161461293a57505060408051808201909152600e81526d08ae0dec6d040dad2e6dac2e8c6d60931b602082015260009150612950565b5050604080516020810190915260008152600191505b935093915050565b60608160000151826020015183604001518460600151604051602001612981949392919061511c565b6040516020818303038152906040529050919050565b60006129a16135ab565b8054909150600119016129c757604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b8281146129ed5760405163512509d360e11b815260040160405180910390fd5b6000839003612a3857604051636337886960e11b8152602060048201526016602482015275139bc81d985b1a59185d1bdc9cc81c1c9bdd9a59195960521b60448201526064016103a3565b60035460405163041c5d3d60e41b8152863560048201526001600160a01b03909116906341c5d3d090602401602060405180830381865afa158015612a81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612aa59190615150565b612ae657604051636337886960e11b81526020600482015260116024820152705375626e6574206e6f742061637469766560781b60448201526064016103a3565b5050505050565b600080826001600160a01b0316635caecd9c60006040518263ffffffff1660e01b8152600401612b1d9190615182565b600060405180830381865afa158015612b3a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052612b62919081019061520a565b519392505050565b60035460405163d5a84a1160e01b8152600481018690526000916001600160a01b03169063d5a84a1190602401602060405180830381865afa158015612bb4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bd89190615039565b90506000816001600160a01b0316632cae8f816040518163ffffffff1660e01b81526004016040805180830381865afa158015612c19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c3d91906153cb565b905082600003612c9057604051636337886960e11b815260206004820152601e60248201527f496e73756666696369656e74206163746976652076616c696461746f7273000060448201526064016103a3565b6000816020015163ffffffff16826000015163ffffffff1685612cb3919061540e565b612cbd9190615425565b905080851015610c9c57604051636337886960e11b8152602060048201526017602482015276496e73756666696369656e74207369676e61747572657360481b60448201526064016103a3565b80356000908152600260209081526040808320548383528184206001600160401b039091168085529252909120541580159081612d5d5750612d526040840160208501615447565b6001600160401b0316155b15612d6757505050565b612d728260016150fd565b6001600160401b0316612d8b6040850160208601615447565b6001600160401b031614612de057612da48260016150fd565b612db46040850160208601615447565b6040516375174c0b60e01b81526001600160401b039283166004820152911660248201526044016103a3565b8235600090815260208181526040808320918391612e0391908801908801615447565b6001600160401b0316815260208101919091526040016000205414612e4f578235612e346040850160208601615447565b604051631c0a5e0f60e21b81526004016103a392919061447c565b612e5f60008435610a1286614ef6565b610b1c576040516340e7e82160e11b815260206004820152600060248201526044016103a3565b612e966040820160208301615447565b6001600160401b0316600003612f0f5760408101351561228057604051636337886960e11b815260206004820152602d60248201527f47656e6573697320636865636b706f696e74206d7573742068617665207a657260448201526c0de40e0c2e4cadce840d0c2e6d609b1b60648201526084016103a3565b60006001612f236040840160208501615447565b612f2d91906150de565b82356000908152602081815260408083206001600160401b03851684529091529020805491925090612fa057604051636337886960e11b815260206004820152601b60248201527a14185c995b9d0818da1958dadc1bda5b9d081b9bdd08199bdd5b99602a1b60448201526064016103a3565b604051600090612fb4908390602001614646565b60405160208183030381529060405280519060200120905080846040013514610a3a57604051636337886960e11b815260206004820152601f60248201527f506172656e7420636865636b706f696e742068617368206d69736d617463680060448201526064016103a3565b6000807f9a81d924239394aaac4d18a0776281a7f1c26f8fd357ba04caedf4f93fa3840583356130566040860160208701615447565b6040860135606087013561307060a0890160808a01615464565b8860a001358960c0016040516020016130899190615481565b60408051601f1981840301815291905280516020909101206130af6101c08c018c615490565b6040516020016130c09291906154d9565b60408051601f1981840301815282825280516020918201209083019a909a528101979097526001600160401b039095166060870152608086019390935260a085019190915263ffffffff1660c084015260e08301526101008201526101208101919091526101400160405160208183030381529060405290506000836101e00160405160200161315091906154ed565b60408051601f1981840301815291905280516020909101206131766102208601866154fb565b604051613184929190615541565b6040805191829003822060208301939093528101919091526102408501356060820152610260850135608082015261028085013560a08201526102a085013560c08201523060e08201524661010082015261012001604051602081830303815290604052905081816040516020016131fd929190615551565b6040516020818303038152906040528051906020012092505050919050565b60005b83811015610c9c576000866001600160a01b031663ac92189087878581811061324a5761324a615580565b905060200201602081019061325f919061445f565b60006040518363ffffffff1660e01b815260040161327e929190615596565b602060405180830381865afa15801561329b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132bf9190615150565b90508061331b5760405163151a7bff60e11b815260206004820152602360248201527f56616c696461746f72206e6f7420617574686f72697a656420666f72207375626044820152621b995d60ea1b60648201526084016103a3565b6133a68885858581811061333157613331615580565b905060200281019061334391906154fb565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508b92508a915087905081811061338c5761338c615580565b90506020020160208101906133a1919061445f565b6135cf565b6133f35760405163151a7bff60e11b815260206004820181905260248201527f496e76616c6964207369676e61747572652066726f6d2076616c696461746f7260448201526064016103a3565b5060010161321f565b6000613412826040516020016104a491906149f8565b82356000908152602081815260408083209394509192918391613439918701908701615447565b6001600160401b031681526020810191909152604001600020905082816134608282615a7d565b505060158101805460ff191660011790554260168201556000601782015561349883356134936040860160208701615447565b613623565b60198201556134ad6040840160208501615447565b833560008181526002602090815260409182902080546001600160401b0319166001600160401b03959095169490941790935590917f53cae6135bd446f0fcf621e412aabe4254b325eeab58936647f9a4dd6099c41a9161351391908701908701615447565b84604051613522929190614789565b60405180910390a250919050565b600061353a6135ab565b6001905550565b61354b8282611a66565b61357a5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016103a3565b5050565b613586613ae0565b611c4757604051631afcd79f60e31b815260040160405180910390fd5b61353061357e565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b7b0ca2ba3432b932bab69029b4b3b732b21026b2b9b9b0b3b29d05199960211b6000908152601c849052603c812060006136098286613afa565b6001600160a01b0390811690851614925050509392505050565b6000816001600160401b031660000361363e5750600161034f565b6000838152602081905260408120816136586001866150de565b6001600160401b03908116825260208083019390935260409182016000908120835161028081018552815460c0808301918252600184015490951660e0808401919091526002840154610100808501919091526003850154610120850152600485015463ffffffff16610140850152600585015461016085015287519081018852600685015481526007850154818a015260088501548189015260098501546060820152600a8501546080820152600b85015460a0820152600c85015496810196909652600d84015490860152610180820194909452600e820180548651818902810189019097528087529196929587959487946101a0880194929392919084015b8282101561391c57838290600052602060002090600502016040518060a001604052908160008201805461378d906144bc565b80601f01602080910402602001604051908101604052809291908181526020018280546137b9906144bc565b80156138065780601f106137db57610100808354040283529160200191613806565b820191906000526020600020905b8154815290600101906020018083116137e957829003601f168201915b5050505050815260200160018201805461381f906144bc565b80601f016020809104026020016040519081016040528092919081815260200182805461384b906144bc565b80156138985780601f1061386d57610100808354040283529160200191613898565b820191906000526020600020905b81548152906001019060200180831161387b57829003601f168201915b5050505050815260200160028201548152602001600382018054806020026020016040519081016040528092919081815260200182805480156138fa57602002820191906000526020600020905b8154815260200190600101908083116138e6575b505050505081526020016004820154815250508152602001906001019061375a565b50505090825250604080518082018252600f8401546001600160401b038082168352600160401b90910416602082810191909152830152601083018054919092019190613968906144bc565b80601f0160208091040260200160405190810160405280929190818152602001828054613994906144bc565b80156139e15780601f106139b6576101008083540402835291602001916139e1565b820191906000526020600020905b8154815290600101906020018083116139c457829003601f168201915b5050505050815260200160118201548152602001601282015481526020016013820154815260200160148201548152505081526020016015820160009054906101000a900460ff166004811115613a3a57613a3a614176565b6004811115613a4b57613a4b614176565b8152602001601682015481526020016017820154815260200160188201805480602002602001604051908101604052809291908181526020018280548015613ab257602002820191906000526020600020905b815481526020019060010190808311613a9e575b5050505050815260200160198201548152505090508060a001516001613ad891906144a9565b949350505050565b6000613aea61239c565b54600160401b900460ff16919050565b600080600080613b0a8686613b24565b925092509250613b1a8282613b71565b5090949350505050565b60008060008351604103613b5e5760208401516040850151606086015160001a613b5088828585613c2a565b955095509550505050613b6a565b50508151600091506002905b9250925092565b6000826003811115613b8557613b85614176565b03613b8e575050565b6001826003811115613ba257613ba2614176565b03613bc05760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613bd457613bd4614176565b03613bf55760405163fce698f760e01b8152600481018290526024016103a3565b6003826003811115613c0957613c09614176565b0361357a576040516335e2f38360e21b8152600481018290526024016103a3565b600080806fa2a8918ca85bafe22016d0b997e4df60600160ff1b03841115613c5b5750600091506003905082613ce5565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613caf573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613cdb57506000925060019150829050613ce5565b9250600091508190505b9450945094915050565b6040518060c00160405280613d02613d2b565b815260200160008152602001600081526020016000815260200160608152602001600081525090565b604080516101c081018252600080825260208083018290528284018290526060808401839052608080850184905260a0808601859052865161010081018852858152938401859052958301849052908201839052810182905292830181905260c083810182905260e0840191909152909190820190815260200160608152602001613ddb604051806040016040528060006001600160401b0316815260200160006001600160401b031681525090565b81526060602082018190526000604083018190529082018190526080820181905260a09091015290565b600060208284031215613e1757600080fd5b81356001600160e01b031981168114610a1757600080fd5b6001600160401b038116811461228057600080fd5b8035613e4f81613e2f565b919050565b60008060408385031215613e6757600080fd5b823591506020830135613e7981613e2f565b809150509250929050565b60006102c08284031215613e9757600080fd5b50919050565b600060208284031215613eaf57600080fd5b81356001600160401b03811115613ec557600080fd5b613ad884828501613e84565b600060208284031215613ee357600080fd5b5035919050565b60005b83811015613f05578181015183820152602001613eed565b50506000910152565b60008151808452613f26816020860160208601613eea565b601f01601f19169290920160200192915050565b600081518084526020840193506020830160005b82811015613f6c578151865260209586019590910190600101613f4e565b5093949350505050565b600082825180855260208501945060208160051b8301016020850160005b8381101561401c57601f198584030188528151805160a08552613fba60a0860182613f0e565b905060208201518582036020870152613fd38282613f0e565b9150506040820151604086015260608201518582036060870152613ff78282613f3a565b6080938401519690930195909552506020988901989093509190910190600101613f94565b50909695505050505050565b805182526000602082015161404860208501826001600160401b03169052565b5060408201516040840152606082015160608401526080820151614074608085018263ffffffff169052565b5060a082015160a084015260c08201516140da60c0850182805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b5060e08201516102c06101c08501526140f76102c0850182613f76565b61010084015180516001600160401b039081166101e0880152602082015116610200870152909150506101208301518482036102208601526141398282613f0e565b9150506101408301516102408501526101608301516102608501526101808301516102808501526101a08301516102a08501528091505092915050565b634e487b7160e01b600052602160045260246000fd5b6005811061419c5761419c614176565b9052565b602081526000825160c060208401526141bc60e0840182614028565b905060208401516141d0604085018261418c565b5060408401516060840152606084015160808401526080840151601f198483030160a08501526142008282613f3a565b91505060a084015160c08401528091505092915050565b6000806040838503121561422a57600080fd5b8235915060208301356001600160401b0381111561424757600080fd5b61425385828601613e84565b9150509250929050565b6001600160a01b038116811461228057600080fd5b6000806040838503121561428557600080fd5b823591506020830135613e798161425d565b600080604083850312156142aa57600080fd5b82356142b58161425d565b91506020830135613e798161425d565b8215158152604060208201526000613ad86040830184613f0e565b602081526000610a176020830184613f0e565b6000806000806060858703121561430957600080fd5b84359350602085013561431b81613e2f565b925060408501356001600160401b0381111561433657600080fd5b8501601f8101871361434757600080fd5b80356001600160401b0381111561435d57600080fd5b87602082840101111561436f57600080fd5b949793965060200194505050565b60008083601f84011261438f57600080fd5b5081356001600160401b038111156143a657600080fd5b6020830191508360208260051b85010111156111a657600080fd5b6000806000806000606086880312156143d957600080fd5b85356001600160401b038111156143ef57600080fd5b6143fb88828901613e84565b95505060208601356001600160401b0381111561441757600080fd5b6144238882890161437d565b90955093505060408601356001600160401b0381111561444257600080fd5b61444e8882890161437d565b969995985093965092949392505050565b60006020828403121561447157600080fd5b8135610a178161425d565b9182526001600160401b0316602082015260400190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561034f5761034f614493565b600181811c908216806144d057607f821691505b602082108103613e9757634e487b7160e01b600052602260045260246000fd5b600081546144fd816144bc565b80855260018216801561451757600181146145335761456a565b60ff1983166020870152602082151560051b870101935061456a565b84600052602060002060005b838110156145615781546020828a01015260018201915060208101905061453f565b87016020019450505b50505092915050565b600082825480855260208501945060208160051b83010184600052602060002060005b8381101561401c57601f1985840301885260a083526145b860a08401836144f0565b83810360208501526145cd81600185016144f0565b600284015460408601528481036060860152600384018054808352600091825260208083209450919291909101905b8083101561461f57835482526020820191506001840193506001830192506145fc565b50600485015460809690960195909552505060209790970196600590910190600101614596565b6020815281546020820152600061466760018401546001600160401b031690565b6001600160401b031660408301526002830154606083015260038301546080830152600483015463ffffffff1660a0830152600583015460c0830152600683015460e0830152600783015461010083015260088301546101208301526009830154610140830152600a830154610160830152600b830154610180830152600c8301546101a0830152600d8301546101c08301526102c06101e08301526147146102e08301600e8501614573565b600f8401546001600160401b0380821661020086015260409190911c16610220840152828103601f190161024084015261475181601086016144f0565b90506011840154610260840152601284015461028084015260138401546102a084015260148401546102c08401528091505092915050565b6001600160401b03929092168252602082015260400190565b63ffffffff8116811461228057600080fd5b8035613e4f816147a2565b803582526020808201359083015260408082013590830152606080820135908301526080808201359083015260a0808201359083015260c0808201359083015260e090810135910152565b6000808335601e1984360301811261482157600080fd5b83016020810192503590506001600160401b0381111561484057600080fd5b8060051b36038213156111a657600080fd5b6000808335601e1984360301811261486957600080fd5b83016020810192503590506001600160401b0381111561488857600080fd5b8036038213156111a657600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008235609e198336030181126148d657600080fd5b90910192915050565b60008383855260208501945060208460051b8201018360005b8681101561401c57838303601f1901885261491382876148c0565b61491d8182614852565b60a0865261492f60a087018284614897565b91505061493f6020830183614852565b8683036020880152614952838284614897565b60408581013590890152925061496e915050606083018361480a565b86830360608801528083526001600160fb1b0381111561498d57600080fd5b60051b8082602085013760809384013596909301959095526020998a019991018101935091909101906001016148f8565b80356149c981613e2f565b6001600160401b0316825260208101356149e281613e2f565b6001600160401b03166020929092019190915250565b6020808252823582820152600090614a11908401613e44565b6001600160401b0381166040840152506040830135606083810191909152830135608080840191909152614a469084016147b4565b63ffffffff811660a084015250600060a084013590508060c084015250614a7360e0830160c085016147bf565b614a816101c084018461480a565b6102c06101e0850152614a996102e0850182846148df565b915050614aae61020084016101e086016149be565b614abc610220850185614852565b848303601f1901610240860152614ad4838284614897565b61024087013561026087810191909152870135610280808801919091528701356102a080880191909152909601356102c09095019490945250929392505050565b634e487b7160e01b600052604160045260246000fd5b60405161010081016001600160401b0381118282101715614b4e57614b4e614b15565b60405290565b60405160a081016001600160401b0381118282101715614b4e57614b4e614b15565b604080519081016001600160401b0381118282101715614b4e57614b4e614b15565b6040516101c081016001600160401b0381118282101715614b4e57614b4e614b15565b60405161014081016001600160401b0381118282101715614b4e57614b4e614b15565b604051601f8201601f191681016001600160401b0381118282101715614c0657614c06614b15565b604052919050565b60006101008284031215614c2157600080fd5b614c29614b2b565b823581526020808401359082015260408084013590820152606080840135908201526080808401359082015260a0808401359082015260c0808401359082015260e0928301359281019290925250919050565b60006001600160401b03821115614c9557614c95614b15565b5060051b60200190565b60006001600160401b03821115614cb857614cb8614b15565b50601f01601f191660200190565b600082601f830112614cd757600080fd5b8135602083016000614cf0614ceb84614c9f565b614bde565b9050828152858383011115614d0457600080fd5b82826020830137600092810160200192909252509392505050565b600082601f830112614d3057600080fd5b8135614d3e614ceb82614c7c565b8082825260208201915060208360051b860101925085831115614d6057600080fd5b602085015b83811015614d7d578035835260209283019201614d65565b5095945050505050565b600082601f830112614d9857600080fd5b8135614da6614ceb82614c7c565b8082825260208201915060208360051b860101925085831115614dc857600080fd5b602085015b83811015614d7d5780356001600160401b03811115614deb57600080fd5b860160a0818903601f19011215614e0157600080fd5b614e09614b54565b60208201356001600160401b03811115614e2257600080fd5b614e318a602083860101614cc6565b82525060408201356001600160401b03811115614e4d57600080fd5b614e5c8a602083860101614cc6565b6020830152506060820135604082015260808201356001600160401b03811115614e8557600080fd5b614e948a602083860101614d1f565b60608301525060a091909101356080820152835260209283019201614dcd565b600060408284031215614ec657600080fd5b614ece614b76565b90508135614edb81613e2f565b81526020820135614eeb81613e2f565b602082015292915050565b60006102c08236031215614f0957600080fd5b614f11614b98565b82358152614f2160208401613e44565b60208201526040838101359082015260608084013590820152614f46608084016147b4565b608082015260a08381013590820152614f623660c08501614c0e565b60c08201526101c08301356001600160401b03811115614f8157600080fd5b614f8d36828601614d87565b60e083015250614fa1366101e08501614eb4565b6101008201526102208301356001600160401b03811115614fc157600080fd5b614fcd36828601614cc6565b610120830152506102408301356101408201526102608301356101608201526102808301356101808201526102a0909201356101a08301525090565b6001600160401b0384168152604060208201819052600090611bd99083018486614897565b8051613e4f8161425d565b60006020828403121561504b57600080fd5b8151610a178161425d565b80516001600160801b0381168114613e4f57600080fd5b6000606082840312801561508057600080fd5b50604051606081016001600160401b03811182821017156150a3576150a3614b15565b6040526150af83615056565b815260208301516150bf81613e2f565b602082015260408301516150d281613e2f565b60408201529392505050565b6001600160401b03828116828216039081111561034f5761034f614493565b6001600160401b03818116838216019081111561034f5761034f614493565b60808152600061512f6080830187614028565b905061513e602083018661418c565b60408201939093526060015292915050565b60006020828403121561516257600080fd5b81518015158114610a1757600080fd5b6004811061419c5761419c614176565b6020810161034f8284615172565b805160048110613e4f57600080fd5b805160038110613e4f57600080fd5b80516001600160501b0381168114613e4f57600080fd5b600082601f8301126151d657600080fd5b81516151e4614ceb82614c9f565b8181528460208386010111156151f957600080fd5b613ad8826020830160208701613eea565b60006020828403121561521c57600080fd5b81516001600160401b0381111561523257600080fd5b8201601f8101841361524357600080fd5b8051615251614ceb82614c7c565b8082825260208201915060208360051b85010192508683111561527357600080fd5b602084015b838110156153c05780516001600160401b0381111561529657600080fd5b8501610140818a03601f190112156152ad57600080fd5b6152b5614bbb565b6152c16020830161502e565b81526152cf60408301615190565b60208201526152e06060830161519f565b60408201526152f1608083016151ae565b606082015261530260a08301615056565b608082015261531360c08301615056565b60a082015260e082015160c08201526101008201516001600160401b0381111561533c57600080fd5b61534b8b6020838601016151c5565b60e0830152506101208201516001600160401b0381111561536b57600080fd5b61537a8b6020838601016151c5565b610100830152506101408201516001600160401b0381111561539b57600080fd5b6153aa8b6020838601016151c5565b6101208301525084525060209283019201615278565b509695505050505050565b600060408284031280156153de57600080fd5b506153e7614b76565b82516153f2816147a2565b81526020830151615402816147a2565b60208201529392505050565b808202811582820484141761034f5761034f614493565b60008261544257634e487b7160e01b600052601260045260246000fd5b500490565b60006020828403121561545957600080fd5b8135610a1781613e2f565b60006020828403121561547657600080fd5b8135610a17816147a2565b610100810161034f82846147bf565b6000808335601e198436030181126154a757600080fd5b8301803591506001600160401b038211156154c157600080fd5b6020019150600581901b36038213156111a657600080fd5b602081526000613ad86020830184866148df565b6040810161034f82846149be565b6000808335601e1984360301811261551257600080fd5b8301803591506001600160401b0382111561552c57600080fd5b6020019150368190038213156111a657600080fd5b8183823760009101908152919050565b60008351615563818460208801613eea565b835190830190615577818360208801613eea565b01949350505050565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b038316815260408101610a176020830184615172565b6000813561034f81613e2f565b80546001600160401b0319166001600160401b0392909216919091179055565b6000813561034f816147a2565b60008235609e1983360301811261560357600080fd5b9190910192915050565b5b8181101561357a576000815560010161560e565b600019600383901b1c191660019190911b1790565b61564181546144bc565b801561357a57601f81116001811461565b57505060009055565b600083815260209020615679601f840160051c82016001830161560d565b5060008381526020812081855555505050565b805460008255801561357a578160005260206000208181015b80821015610a3a57600082556001820191506156a5565b601f821115610b1c57806000526020600020601f840160051c810160208510156156e35750805b612ae6601f850160051c83018261560d565b6001600160401b0383111561570c5761570c614b15565b6157208361571a83546144bc565b836156bc565b6000601f84116001811461574e576000851561573c5750838201355b6157468682615622565b845550612ae6565b600083815260209020601f19861690835b8281101561577f578685013582556020948501946001909201910161575f565b508682101561579c5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b6001600160401b038311156157c5576157c5614b15565b600160401b8311156157d9576157d9614b15565b805483825580841015615810576000828152602090208481019082015b8082101561580d57600082556001820191506157f6565b50505b508181600052602060002060005b8581101561583a5782358282015560209092019160010161581e565b505050505050565b61584c82836154fb565b6001600160401b0381111561586357615863614b15565b6158778161587185546144bc565b856156bc565b6000601f8211600181146158a557600083156158935750838201355b61589d8482615622565b8655506158ff565b600085815260209020601f19841690835b828110156158d657868501358255602094850194600190920191016158b6565b50848210156158f35760001960f88660031b161c19848701351681555b505060018360011b0185555b5050505061591060208301836154fb565b61591e8183600186016156f5565b5050604082013560028201556159376060830183615490565b6159458183600386016157ae565b505060809190910135600490910155565b600160401b83111561596a5761596a614b15565b8054838255808410156159f4578060050260058104821461598d5761598d614493565b846005026005810486146159a3576159a3614493565b60008481526020902091820191015b818110156159f1576159c381615637565b6159cf60018201615637565b600060028201556159e26003820161568c565b600060048201556005016159b2565b50505b5060008181526020812083915b8581101561583a57615a1c615a1684876155ed565b83615842565b6020929092019160059190910190600101615a01565b8135615a3d81613e2f565b615a4781836155c0565b506020820135615a5681613e2f565b8154600160401b600160801b03191660409190911b600160401b600160801b031617905550565b81358155615a99615a90602084016155b3565b600183016155c0565b6040820135600282015560608201356003820155615ad7615abc608084016155e0565b6004830163ffffffff821663ffffffff198254161781555050565b60a0820135600582015560c0820135600682015560e0820135600782015561010082013560088201556101208201356009820155610140820135600a820155610160820135600b820155610180820135600c8201556101a0820135600d820155615b456101c0830183615490565b615b538183600e8601615956565b5050615b666101e08301600f8301615a32565b615b746102208301836154fb565b615b828183601086016156f5565b50506102408201356011820155610260820135601282015561028082013560138201556102a09091013560149091015556fe71840dc4906352362b0cdaf79870196c8e42acafade72d5d5a6d59291253ceb1a26469706673582212207eef8341438ec6dcdd643b534f51ef91e1ed501911341d37fffdb2479038821764736f6c634300081b0033",
}

// CheckpointManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use CheckpointManagerMetaData.ABI instead.
var CheckpointManagerABI = CheckpointManagerMetaData.ABI

// CheckpointManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CheckpointManagerMetaData.Bin instead.
var CheckpointManagerBin = CheckpointManagerMetaData.Bin

// DeployCheckpointManager deploys a new Ethereum contract, binding an instance of CheckpointManager to it.
func DeployCheckpointManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CheckpointManager, error) {
	parsed, err := CheckpointManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CheckpointManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CheckpointManager{CheckpointManagerCaller: CheckpointManagerCaller{contract: contract}, CheckpointManagerTransactor: CheckpointManagerTransactor{contract: contract}, CheckpointManagerFilterer: CheckpointManagerFilterer{contract: contract}}, nil
}

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
