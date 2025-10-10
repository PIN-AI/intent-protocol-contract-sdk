// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subnetfactory

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

// DataStructuresCheckpointPolicy is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresCheckpointPolicy struct {
	ChallengeWindow  *big.Int
	MinEpochInterval uint64
	MaxEpochInterval uint64
}

// DataStructuresCreateSubnetInfo is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresCreateSubnetInfo struct {
	CanonicalName     string
	Owner             common.Address
	Version           uint16
	DaKind            string
	SigScheme         string
	CpPolicy          DataStructuresCheckpointPolicy
	SigThreshold      DataStructuresSignatureThreshold
	StakeCfg          DataStructuresStakeGovernanceConfig
	MetadataUri       string
	BidFrequencyLimit *big.Int
	RequireKYC        bool
	AutoApprove       bool
}

// DataStructuresParticipantInfo is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresParticipantInfo struct {
	Owner           common.Address
	ParticipantType uint8
	Status          uint8
	ReputationScore *big.Int
	RegisteredAt    *big.Int
	LastActive      *big.Int
	Reserved        *big.Int
	Endpoint        string
	Domain          string
	MetadataUri     string
}

// DataStructuresSignatureThreshold is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresSignatureThreshold struct {
	ThresholdNumerator   uint32
	ThresholdDenominator uint32
}

// DataStructuresStakeGovernanceConfig is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresStakeGovernanceConfig struct {
	MinValidatorStake *big.Int
	MinAgentStake     *big.Int
	MinMatcherStake   *big.Int
	MaxValidators     *big.Int
	MaxAgents         *big.Int
	MaxMatchers       *big.Int
	UnstakeLockPeriod *big.Int
	SlashingRates     []*big.Int
}

// DataStructuresSubnetInfo is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresSubnetInfo struct {
	SubnetId          [32]byte
	CanonicalName     string
	Owner             common.Address
	Version           uint16
	DaKind            string
	SigScheme         string
	AutoApprove       bool
	RequireKYC        bool
	Status            uint8
	CpPolicy          DataStructuresCheckpointPolicy
	SigThreshold      DataStructuresSignatureThreshold
	StakeCfg          DataStructuresStakeGovernanceConfig
	MetadataUri       string
	CreatedAt         *big.Int
	UpdatedAt         *big.Int
	BidFrequencyLimit *big.Int
}

// DataStructuresSubnetParticipants is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresSubnetParticipants struct {
	SubnetId   [32]byte
	Validators []DataStructuresParticipantInfo
	Agents     []DataStructuresParticipantInfo
	Matchers   []DataStructuresParticipantInfo
}

// SubnetFactoryMetaData contains all meta data concerning the SubnetFactory contract.
var SubnetFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"}],\"name\":\"GlobalParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MinStakeCreateSubnetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subnet_contract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"}],\"name\":\"SubnetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"SubnetDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"SubnetPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"SubnetResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARDIAN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSubnets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"}],\"internalType\":\"structDataStructures.CreateSubnetInfo\",\"name\":\"create_info\",\"type\":\"tuple\"}],\"name\":\"createSubnet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"deprecateSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"}],\"name\":\"findSubnetByName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"subnet_ids\",\"type\":\"bytes32[]\"}],\"name\":\"getActiveParticipants\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint80\",\"name\":\"reputation_score\",\"type\":\"uint80\"},{\"internalType\":\"uint128\",\"name\":\"registered_at\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"reserved\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"internalType\":\"structDataStructures.ParticipantInfo[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint80\",\"name\":\"reputation_score\",\"type\":\"uint80\"},{\"internalType\":\"uint128\",\"name\":\"registered_at\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"reserved\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"internalType\":\"structDataStructures.ParticipantInfo[]\",\"name\":\"agents\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint80\",\"name\":\"reputation_score\",\"type\":\"uint80\"},{\"internalType\":\"uint128\",\"name\":\"registered_at\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"reserved\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"internalType\":\"structDataStructures.ParticipantInfo[]\",\"name\":\"matchers\",\"type\":\"tuple[]\"}],\"internalType\":\"structDataStructures.SubnetParticipants[]\",\"name\":\"participants\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveSubnetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSubnetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"validator_counts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"agent_counts\",\"type\":\"uint256[]\"},{\"internalType\":\"enumDataStructures.SubnetStatus[]\",\"name\":\"statuses\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreationStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total_created\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStakeCreateSubnet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_stakingManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"subnet_contract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubnetImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"validator_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agent_count\",\"type\":\"uint256\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"subnet_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetStakeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"threshold\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getSubnetsByOwner\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"subnet_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"getSubnetsByStatus\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"subnet_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSubnetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_subnetImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"isSubnetActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listAllSubnets\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"subnet_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"pauseSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"predictSubnetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"predicted_address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"resumeSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeCreateSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingManager\",\"type\":\"address\"}],\"name\":\"setStakingManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_implementation\",\"type\":\"address\"}],\"name\":\"setSubnetImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingManager\",\"outputs\":[{\"internalType\":\"contractIStakingManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subnetContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"subnetExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subnetImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SubnetFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SubnetFactoryMetaData.ABI instead.
var SubnetFactoryABI = SubnetFactoryMetaData.ABI

// SubnetFactory is an auto generated Go binding around an Ethereum contract.
type SubnetFactory struct {
	SubnetFactoryCaller     // Read-only binding to the contract
	SubnetFactoryTransactor // Write-only binding to the contract
	SubnetFactoryFilterer   // Log filterer for contract events
}

// SubnetFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubnetFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubnetFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubnetFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubnetFactorySession struct {
	Contract     *SubnetFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubnetFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubnetFactoryCallerSession struct {
	Contract *SubnetFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SubnetFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubnetFactoryTransactorSession struct {
	Contract     *SubnetFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SubnetFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubnetFactoryRaw struct {
	Contract *SubnetFactory // Generic contract binding to access the raw methods on
}

// SubnetFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubnetFactoryCallerRaw struct {
	Contract *SubnetFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// SubnetFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubnetFactoryTransactorRaw struct {
	Contract *SubnetFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubnetFactory creates a new instance of SubnetFactory, bound to a specific deployed contract.
func NewSubnetFactory(address common.Address, backend bind.ContractBackend) (*SubnetFactory, error) {
	contract, err := bindSubnetFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SubnetFactory{SubnetFactoryCaller: SubnetFactoryCaller{contract: contract}, SubnetFactoryTransactor: SubnetFactoryTransactor{contract: contract}, SubnetFactoryFilterer: SubnetFactoryFilterer{contract: contract}}, nil
}

// NewSubnetFactoryCaller creates a new read-only instance of SubnetFactory, bound to a specific deployed contract.
func NewSubnetFactoryCaller(address common.Address, caller bind.ContractCaller) (*SubnetFactoryCaller, error) {
	contract, err := bindSubnetFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryCaller{contract: contract}, nil
}

// NewSubnetFactoryTransactor creates a new write-only instance of SubnetFactory, bound to a specific deployed contract.
func NewSubnetFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*SubnetFactoryTransactor, error) {
	contract, err := bindSubnetFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryTransactor{contract: contract}, nil
}

// NewSubnetFactoryFilterer creates a new log filterer instance of SubnetFactory, bound to a specific deployed contract.
func NewSubnetFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*SubnetFactoryFilterer, error) {
	contract, err := bindSubnetFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryFilterer{contract: contract}, nil
}

// bindSubnetFactory binds a generic wrapper to an already deployed contract.
func bindSubnetFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SubnetFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubnetFactory *SubnetFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubnetFactory.Contract.SubnetFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubnetFactory *SubnetFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SubnetFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubnetFactory *SubnetFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SubnetFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SubnetFactory *SubnetFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SubnetFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SubnetFactory *SubnetFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubnetFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SubnetFactory *SubnetFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SubnetFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SubnetFactory *SubnetFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SubnetFactory.Contract.DEFAULTADMINROLE(&_SubnetFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SubnetFactory.Contract.DEFAULTADMINROLE(&_SubnetFactory.CallOpts)
}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCaller) GUARDIANROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "GUARDIAN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_SubnetFactory *SubnetFactorySession) GUARDIANROLE() ([32]byte, error) {
	return _SubnetFactory.Contract.GUARDIANROLE(&_SubnetFactory.CallOpts)
}

// GUARDIANROLE is a free data retrieval call binding the contract method 0x24ea54f4.
//
// Solidity: function GUARDIAN_ROLE() view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCallerSession) GUARDIANROLE() ([32]byte, error) {
	return _SubnetFactory.Contract.GUARDIANROLE(&_SubnetFactory.CallOpts)
}

// AllSubnets is a free data retrieval call binding the contract method 0xc410d562.
//
// Solidity: function allSubnets(uint256 ) view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCaller) AllSubnets(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "allSubnets", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AllSubnets is a free data retrieval call binding the contract method 0xc410d562.
//
// Solidity: function allSubnets(uint256 ) view returns(bytes32)
func (_SubnetFactory *SubnetFactorySession) AllSubnets(arg0 *big.Int) ([32]byte, error) {
	return _SubnetFactory.Contract.AllSubnets(&_SubnetFactory.CallOpts, arg0)
}

// AllSubnets is a free data retrieval call binding the contract method 0xc410d562.
//
// Solidity: function allSubnets(uint256 ) view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCallerSession) AllSubnets(arg0 *big.Int) ([32]byte, error) {
	return _SubnetFactory.Contract.AllSubnets(&_SubnetFactory.CallOpts, arg0)
}

// FindSubnetByName is a free data retrieval call binding the contract method 0xdf43dc0f.
//
// Solidity: function findSubnetByName(string canonical_name) view returns(bytes32 subnet_id)
func (_SubnetFactory *SubnetFactoryCaller) FindSubnetByName(opts *bind.CallOpts, canonical_name string) ([32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "findSubnetByName", canonical_name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FindSubnetByName is a free data retrieval call binding the contract method 0xdf43dc0f.
//
// Solidity: function findSubnetByName(string canonical_name) view returns(bytes32 subnet_id)
func (_SubnetFactory *SubnetFactorySession) FindSubnetByName(canonical_name string) ([32]byte, error) {
	return _SubnetFactory.Contract.FindSubnetByName(&_SubnetFactory.CallOpts, canonical_name)
}

// FindSubnetByName is a free data retrieval call binding the contract method 0xdf43dc0f.
//
// Solidity: function findSubnetByName(string canonical_name) view returns(bytes32 subnet_id)
func (_SubnetFactory *SubnetFactoryCallerSession) FindSubnetByName(canonical_name string) ([32]byte, error) {
	return _SubnetFactory.Contract.FindSubnetByName(&_SubnetFactory.CallOpts, canonical_name)
}

// GetActiveParticipants is a free data retrieval call binding the contract method 0xd421495e.
//
// Solidity: function getActiveParticipants(bytes32[] subnet_ids) view returns((bytes32,(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[],(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[],(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[])[] participants)
func (_SubnetFactory *SubnetFactoryCaller) GetActiveParticipants(opts *bind.CallOpts, subnet_ids [][32]byte) ([]DataStructuresSubnetParticipants, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getActiveParticipants", subnet_ids)

	if err != nil {
		return *new([]DataStructuresSubnetParticipants), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructuresSubnetParticipants)).(*[]DataStructuresSubnetParticipants)

	return out0, err

}

// GetActiveParticipants is a free data retrieval call binding the contract method 0xd421495e.
//
// Solidity: function getActiveParticipants(bytes32[] subnet_ids) view returns((bytes32,(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[],(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[],(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[])[] participants)
func (_SubnetFactory *SubnetFactorySession) GetActiveParticipants(subnet_ids [][32]byte) ([]DataStructuresSubnetParticipants, error) {
	return _SubnetFactory.Contract.GetActiveParticipants(&_SubnetFactory.CallOpts, subnet_ids)
}

// GetActiveParticipants is a free data retrieval call binding the contract method 0xd421495e.
//
// Solidity: function getActiveParticipants(bytes32[] subnet_ids) view returns((bytes32,(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[],(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[],(address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[])[] participants)
func (_SubnetFactory *SubnetFactoryCallerSession) GetActiveParticipants(subnet_ids [][32]byte) ([]DataStructuresSubnetParticipants, error) {
	return _SubnetFactory.Contract.GetActiveParticipants(&_SubnetFactory.CallOpts, subnet_ids)
}

// GetActiveSubnetCount is a free data retrieval call binding the contract method 0xc6a1a01a.
//
// Solidity: function getActiveSubnetCount() view returns(uint256 count)
func (_SubnetFactory *SubnetFactoryCaller) GetActiveSubnetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getActiveSubnetCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveSubnetCount is a free data retrieval call binding the contract method 0xc6a1a01a.
//
// Solidity: function getActiveSubnetCount() view returns(uint256 count)
func (_SubnetFactory *SubnetFactorySession) GetActiveSubnetCount() (*big.Int, error) {
	return _SubnetFactory.Contract.GetActiveSubnetCount(&_SubnetFactory.CallOpts)
}

// GetActiveSubnetCount is a free data retrieval call binding the contract method 0xc6a1a01a.
//
// Solidity: function getActiveSubnetCount() view returns(uint256 count)
func (_SubnetFactory *SubnetFactoryCallerSession) GetActiveSubnetCount() (*big.Int, error) {
	return _SubnetFactory.Contract.GetActiveSubnetCount(&_SubnetFactory.CallOpts)
}

// GetAllSubnetInfo is a free data retrieval call binding the contract method 0x9dba0976.
//
// Solidity: function getAllSubnetInfo() view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256)[] infos, uint256[] validator_counts, uint256[] agent_counts, uint8[] statuses)
func (_SubnetFactory *SubnetFactoryCaller) GetAllSubnetInfo(opts *bind.CallOpts) (struct {
	Infos           []DataStructuresSubnetInfo
	ValidatorCounts []*big.Int
	AgentCounts     []*big.Int
	Statuses        []uint8
}, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getAllSubnetInfo")

	outstruct := new(struct {
		Infos           []DataStructuresSubnetInfo
		ValidatorCounts []*big.Int
		AgentCounts     []*big.Int
		Statuses        []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Infos = *abi.ConvertType(out[0], new([]DataStructuresSubnetInfo)).(*[]DataStructuresSubnetInfo)
	outstruct.ValidatorCounts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.AgentCounts = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.Statuses = *abi.ConvertType(out[3], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// GetAllSubnetInfo is a free data retrieval call binding the contract method 0x9dba0976.
//
// Solidity: function getAllSubnetInfo() view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256)[] infos, uint256[] validator_counts, uint256[] agent_counts, uint8[] statuses)
func (_SubnetFactory *SubnetFactorySession) GetAllSubnetInfo() (struct {
	Infos           []DataStructuresSubnetInfo
	ValidatorCounts []*big.Int
	AgentCounts     []*big.Int
	Statuses        []uint8
}, error) {
	return _SubnetFactory.Contract.GetAllSubnetInfo(&_SubnetFactory.CallOpts)
}

// GetAllSubnetInfo is a free data retrieval call binding the contract method 0x9dba0976.
//
// Solidity: function getAllSubnetInfo() view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256)[] infos, uint256[] validator_counts, uint256[] agent_counts, uint8[] statuses)
func (_SubnetFactory *SubnetFactoryCallerSession) GetAllSubnetInfo() (struct {
	Infos           []DataStructuresSubnetInfo
	ValidatorCounts []*big.Int
	AgentCounts     []*big.Int
	Statuses        []uint8
}, error) {
	return _SubnetFactory.Contract.GetAllSubnetInfo(&_SubnetFactory.CallOpts)
}

// GetCreationStats is a free data retrieval call binding the contract method 0xea869cfd.
//
// Solidity: function getCreationStats() view returns(uint256 total_created)
func (_SubnetFactory *SubnetFactoryCaller) GetCreationStats(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getCreationStats")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCreationStats is a free data retrieval call binding the contract method 0xea869cfd.
//
// Solidity: function getCreationStats() view returns(uint256 total_created)
func (_SubnetFactory *SubnetFactorySession) GetCreationStats() (*big.Int, error) {
	return _SubnetFactory.Contract.GetCreationStats(&_SubnetFactory.CallOpts)
}

// GetCreationStats is a free data retrieval call binding the contract method 0xea869cfd.
//
// Solidity: function getCreationStats() view returns(uint256 total_created)
func (_SubnetFactory *SubnetFactoryCallerSession) GetCreationStats() (*big.Int, error) {
	return _SubnetFactory.Contract.GetCreationStats(&_SubnetFactory.CallOpts)
}

// GetMinStakeCreateSubnet is a free data retrieval call binding the contract method 0xf5caee81.
//
// Solidity: function getMinStakeCreateSubnet() view returns(uint256 amount)
func (_SubnetFactory *SubnetFactoryCaller) GetMinStakeCreateSubnet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getMinStakeCreateSubnet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinStakeCreateSubnet is a free data retrieval call binding the contract method 0xf5caee81.
//
// Solidity: function getMinStakeCreateSubnet() view returns(uint256 amount)
func (_SubnetFactory *SubnetFactorySession) GetMinStakeCreateSubnet() (*big.Int, error) {
	return _SubnetFactory.Contract.GetMinStakeCreateSubnet(&_SubnetFactory.CallOpts)
}

// GetMinStakeCreateSubnet is a free data retrieval call binding the contract method 0xf5caee81.
//
// Solidity: function getMinStakeCreateSubnet() view returns(uint256 amount)
func (_SubnetFactory *SubnetFactoryCallerSession) GetMinStakeCreateSubnet() (*big.Int, error) {
	return _SubnetFactory.Contract.GetMinStakeCreateSubnet(&_SubnetFactory.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SubnetFactory *SubnetFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SubnetFactory.Contract.GetRoleAdmin(&_SubnetFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SubnetFactory *SubnetFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SubnetFactory.Contract.GetRoleAdmin(&_SubnetFactory.CallOpts, role)
}

// GetStakingManager is a free data retrieval call binding the contract method 0xee7941d9.
//
// Solidity: function getStakingManager() view returns(address _stakingManager)
func (_SubnetFactory *SubnetFactoryCaller) GetStakingManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getStakingManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingManager is a free data retrieval call binding the contract method 0xee7941d9.
//
// Solidity: function getStakingManager() view returns(address _stakingManager)
func (_SubnetFactory *SubnetFactorySession) GetStakingManager() (common.Address, error) {
	return _SubnetFactory.Contract.GetStakingManager(&_SubnetFactory.CallOpts)
}

// GetStakingManager is a free data retrieval call binding the contract method 0xee7941d9.
//
// Solidity: function getStakingManager() view returns(address _stakingManager)
func (_SubnetFactory *SubnetFactoryCallerSession) GetStakingManager() (common.Address, error) {
	return _SubnetFactory.Contract.GetStakingManager(&_SubnetFactory.CallOpts)
}

// GetSubnetContract is a free data retrieval call binding the contract method 0xd5a84a11.
//
// Solidity: function getSubnetContract(bytes32 subnet_id) view returns(address subnet_contract)
func (_SubnetFactory *SubnetFactoryCaller) GetSubnetContract(opts *bind.CallOpts, subnet_id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getSubnetContract", subnet_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubnetContract is a free data retrieval call binding the contract method 0xd5a84a11.
//
// Solidity: function getSubnetContract(bytes32 subnet_id) view returns(address subnet_contract)
func (_SubnetFactory *SubnetFactorySession) GetSubnetContract(subnet_id [32]byte) (common.Address, error) {
	return _SubnetFactory.Contract.GetSubnetContract(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetContract is a free data retrieval call binding the contract method 0xd5a84a11.
//
// Solidity: function getSubnetContract(bytes32 subnet_id) view returns(address subnet_contract)
func (_SubnetFactory *SubnetFactoryCallerSession) GetSubnetContract(subnet_id [32]byte) (common.Address, error) {
	return _SubnetFactory.Contract.GetSubnetContract(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetImplementation is a free data retrieval call binding the contract method 0xe8175c4a.
//
// Solidity: function getSubnetImplementation() view returns(address implementation)
func (_SubnetFactory *SubnetFactoryCaller) GetSubnetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getSubnetImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSubnetImplementation is a free data retrieval call binding the contract method 0xe8175c4a.
//
// Solidity: function getSubnetImplementation() view returns(address implementation)
func (_SubnetFactory *SubnetFactorySession) GetSubnetImplementation() (common.Address, error) {
	return _SubnetFactory.Contract.GetSubnetImplementation(&_SubnetFactory.CallOpts)
}

// GetSubnetImplementation is a free data retrieval call binding the contract method 0xe8175c4a.
//
// Solidity: function getSubnetImplementation() view returns(address implementation)
func (_SubnetFactory *SubnetFactoryCallerSession) GetSubnetImplementation() (common.Address, error) {
	return _SubnetFactory.Contract.GetSubnetImplementation(&_SubnetFactory.CallOpts)
}

// GetSubnetInfo is a free data retrieval call binding the contract method 0xa3923d83.
//
// Solidity: function getSubnetInfo(bytes32 subnet_id) view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) info, uint256 validator_count, uint256 agent_count, uint8 subnet_status)
func (_SubnetFactory *SubnetFactoryCaller) GetSubnetInfo(opts *bind.CallOpts, subnet_id [32]byte) (struct {
	Info           DataStructuresSubnetInfo
	ValidatorCount *big.Int
	AgentCount     *big.Int
	SubnetStatus   uint8
}, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getSubnetInfo", subnet_id)

	outstruct := new(struct {
		Info           DataStructuresSubnetInfo
		ValidatorCount *big.Int
		AgentCount     *big.Int
		SubnetStatus   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Info = *abi.ConvertType(out[0], new(DataStructuresSubnetInfo)).(*DataStructuresSubnetInfo)
	outstruct.ValidatorCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AgentCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SubnetStatus = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetSubnetInfo is a free data retrieval call binding the contract method 0xa3923d83.
//
// Solidity: function getSubnetInfo(bytes32 subnet_id) view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) info, uint256 validator_count, uint256 agent_count, uint8 subnet_status)
func (_SubnetFactory *SubnetFactorySession) GetSubnetInfo(subnet_id [32]byte) (struct {
	Info           DataStructuresSubnetInfo
	ValidatorCount *big.Int
	AgentCount     *big.Int
	SubnetStatus   uint8
}, error) {
	return _SubnetFactory.Contract.GetSubnetInfo(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetInfo is a free data retrieval call binding the contract method 0xa3923d83.
//
// Solidity: function getSubnetInfo(bytes32 subnet_id) view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) info, uint256 validator_count, uint256 agent_count, uint8 subnet_status)
func (_SubnetFactory *SubnetFactoryCallerSession) GetSubnetInfo(subnet_id [32]byte) (struct {
	Info           DataStructuresSubnetInfo
	ValidatorCount *big.Int
	AgentCount     *big.Int
	SubnetStatus   uint8
}, error) {
	return _SubnetFactory.Contract.GetSubnetInfo(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetStakeConfig is a free data retrieval call binding the contract method 0x48f0d023.
//
// Solidity: function getSubnetStakeConfig(bytes32 subnet_id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) config)
func (_SubnetFactory *SubnetFactoryCaller) GetSubnetStakeConfig(opts *bind.CallOpts, subnet_id [32]byte) (DataStructuresStakeGovernanceConfig, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getSubnetStakeConfig", subnet_id)

	if err != nil {
		return *new(DataStructuresStakeGovernanceConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresStakeGovernanceConfig)).(*DataStructuresStakeGovernanceConfig)

	return out0, err

}

// GetSubnetStakeConfig is a free data retrieval call binding the contract method 0x48f0d023.
//
// Solidity: function getSubnetStakeConfig(bytes32 subnet_id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) config)
func (_SubnetFactory *SubnetFactorySession) GetSubnetStakeConfig(subnet_id [32]byte) (DataStructuresStakeGovernanceConfig, error) {
	return _SubnetFactory.Contract.GetSubnetStakeConfig(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetStakeConfig is a free data retrieval call binding the contract method 0x48f0d023.
//
// Solidity: function getSubnetStakeConfig(bytes32 subnet_id) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) config)
func (_SubnetFactory *SubnetFactoryCallerSession) GetSubnetStakeConfig(subnet_id [32]byte) (DataStructuresStakeGovernanceConfig, error) {
	return _SubnetFactory.Contract.GetSubnetStakeConfig(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetThreshold is a free data retrieval call binding the contract method 0x54eae4ca.
//
// Solidity: function getSubnetThreshold(bytes32 subnet_id) view returns((uint32,uint32) threshold)
func (_SubnetFactory *SubnetFactoryCaller) GetSubnetThreshold(opts *bind.CallOpts, subnet_id [32]byte) (DataStructuresSignatureThreshold, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getSubnetThreshold", subnet_id)

	if err != nil {
		return *new(DataStructuresSignatureThreshold), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresSignatureThreshold)).(*DataStructuresSignatureThreshold)

	return out0, err

}

// GetSubnetThreshold is a free data retrieval call binding the contract method 0x54eae4ca.
//
// Solidity: function getSubnetThreshold(bytes32 subnet_id) view returns((uint32,uint32) threshold)
func (_SubnetFactory *SubnetFactorySession) GetSubnetThreshold(subnet_id [32]byte) (DataStructuresSignatureThreshold, error) {
	return _SubnetFactory.Contract.GetSubnetThreshold(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetThreshold is a free data retrieval call binding the contract method 0x54eae4ca.
//
// Solidity: function getSubnetThreshold(bytes32 subnet_id) view returns((uint32,uint32) threshold)
func (_SubnetFactory *SubnetFactoryCallerSession) GetSubnetThreshold(subnet_id [32]byte) (DataStructuresSignatureThreshold, error) {
	return _SubnetFactory.Contract.GetSubnetThreshold(&_SubnetFactory.CallOpts, subnet_id)
}

// GetSubnetsByOwner is a free data retrieval call binding the contract method 0x82e97eef.
//
// Solidity: function getSubnetsByOwner(address owner) view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactoryCaller) GetSubnetsByOwner(opts *bind.CallOpts, owner common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getSubnetsByOwner", owner)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSubnetsByOwner is a free data retrieval call binding the contract method 0x82e97eef.
//
// Solidity: function getSubnetsByOwner(address owner) view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactorySession) GetSubnetsByOwner(owner common.Address) ([][32]byte, error) {
	return _SubnetFactory.Contract.GetSubnetsByOwner(&_SubnetFactory.CallOpts, owner)
}

// GetSubnetsByOwner is a free data retrieval call binding the contract method 0x82e97eef.
//
// Solidity: function getSubnetsByOwner(address owner) view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactoryCallerSession) GetSubnetsByOwner(owner common.Address) ([][32]byte, error) {
	return _SubnetFactory.Contract.GetSubnetsByOwner(&_SubnetFactory.CallOpts, owner)
}

// GetSubnetsByStatus is a free data retrieval call binding the contract method 0x0ffdeb6a.
//
// Solidity: function getSubnetsByStatus(uint8 status) view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactoryCaller) GetSubnetsByStatus(opts *bind.CallOpts, status uint8) ([][32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getSubnetsByStatus", status)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetSubnetsByStatus is a free data retrieval call binding the contract method 0x0ffdeb6a.
//
// Solidity: function getSubnetsByStatus(uint8 status) view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactorySession) GetSubnetsByStatus(status uint8) ([][32]byte, error) {
	return _SubnetFactory.Contract.GetSubnetsByStatus(&_SubnetFactory.CallOpts, status)
}

// GetSubnetsByStatus is a free data retrieval call binding the contract method 0x0ffdeb6a.
//
// Solidity: function getSubnetsByStatus(uint8 status) view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactoryCallerSession) GetSubnetsByStatus(status uint8) ([][32]byte, error) {
	return _SubnetFactory.Contract.GetSubnetsByStatus(&_SubnetFactory.CallOpts, status)
}

// GetTotalSubnetCount is a free data retrieval call binding the contract method 0x1f039689.
//
// Solidity: function getTotalSubnetCount() view returns(uint256 count)
func (_SubnetFactory *SubnetFactoryCaller) GetTotalSubnetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "getTotalSubnetCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSubnetCount is a free data retrieval call binding the contract method 0x1f039689.
//
// Solidity: function getTotalSubnetCount() view returns(uint256 count)
func (_SubnetFactory *SubnetFactorySession) GetTotalSubnetCount() (*big.Int, error) {
	return _SubnetFactory.Contract.GetTotalSubnetCount(&_SubnetFactory.CallOpts)
}

// GetTotalSubnetCount is a free data retrieval call binding the contract method 0x1f039689.
//
// Solidity: function getTotalSubnetCount() view returns(uint256 count)
func (_SubnetFactory *SubnetFactoryCallerSession) GetTotalSubnetCount() (*big.Int, error) {
	return _SubnetFactory.Contract.GetTotalSubnetCount(&_SubnetFactory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SubnetFactory *SubnetFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SubnetFactory *SubnetFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SubnetFactory.Contract.HasRole(&_SubnetFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SubnetFactory *SubnetFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SubnetFactory.Contract.HasRole(&_SubnetFactory.CallOpts, role, account)
}

// IsSubnetActive is a free data retrieval call binding the contract method 0x41c5d3d0.
//
// Solidity: function isSubnetActive(bytes32 subnet_id) view returns(bool active)
func (_SubnetFactory *SubnetFactoryCaller) IsSubnetActive(opts *bind.CallOpts, subnet_id [32]byte) (bool, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "isSubnetActive", subnet_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSubnetActive is a free data retrieval call binding the contract method 0x41c5d3d0.
//
// Solidity: function isSubnetActive(bytes32 subnet_id) view returns(bool active)
func (_SubnetFactory *SubnetFactorySession) IsSubnetActive(subnet_id [32]byte) (bool, error) {
	return _SubnetFactory.Contract.IsSubnetActive(&_SubnetFactory.CallOpts, subnet_id)
}

// IsSubnetActive is a free data retrieval call binding the contract method 0x41c5d3d0.
//
// Solidity: function isSubnetActive(bytes32 subnet_id) view returns(bool active)
func (_SubnetFactory *SubnetFactoryCallerSession) IsSubnetActive(subnet_id [32]byte) (bool, error) {
	return _SubnetFactory.Contract.IsSubnetActive(&_SubnetFactory.CallOpts, subnet_id)
}

// ListAllSubnets is a free data retrieval call binding the contract method 0x110bc3c0.
//
// Solidity: function listAllSubnets() view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactoryCaller) ListAllSubnets(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "listAllSubnets")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ListAllSubnets is a free data retrieval call binding the contract method 0x110bc3c0.
//
// Solidity: function listAllSubnets() view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactorySession) ListAllSubnets() ([][32]byte, error) {
	return _SubnetFactory.Contract.ListAllSubnets(&_SubnetFactory.CallOpts)
}

// ListAllSubnets is a free data retrieval call binding the contract method 0x110bc3c0.
//
// Solidity: function listAllSubnets() view returns(bytes32[] subnet_ids)
func (_SubnetFactory *SubnetFactoryCallerSession) ListAllSubnets() ([][32]byte, error) {
	return _SubnetFactory.Contract.ListAllSubnets(&_SubnetFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SubnetFactory *SubnetFactoryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SubnetFactory *SubnetFactorySession) Paused() (bool, error) {
	return _SubnetFactory.Contract.Paused(&_SubnetFactory.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SubnetFactory *SubnetFactoryCallerSession) Paused() (bool, error) {
	return _SubnetFactory.Contract.Paused(&_SubnetFactory.CallOpts)
}

// PredictSubnetAddress is a free data retrieval call binding the contract method 0xa2379c91.
//
// Solidity: function predictSubnetAddress(bytes32 subnet_id) view returns(address predicted_address)
func (_SubnetFactory *SubnetFactoryCaller) PredictSubnetAddress(opts *bind.CallOpts, subnet_id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "predictSubnetAddress", subnet_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictSubnetAddress is a free data retrieval call binding the contract method 0xa2379c91.
//
// Solidity: function predictSubnetAddress(bytes32 subnet_id) view returns(address predicted_address)
func (_SubnetFactory *SubnetFactorySession) PredictSubnetAddress(subnet_id [32]byte) (common.Address, error) {
	return _SubnetFactory.Contract.PredictSubnetAddress(&_SubnetFactory.CallOpts, subnet_id)
}

// PredictSubnetAddress is a free data retrieval call binding the contract method 0xa2379c91.
//
// Solidity: function predictSubnetAddress(bytes32 subnet_id) view returns(address predicted_address)
func (_SubnetFactory *SubnetFactoryCallerSession) PredictSubnetAddress(subnet_id [32]byte) (common.Address, error) {
	return _SubnetFactory.Contract.PredictSubnetAddress(&_SubnetFactory.CallOpts, subnet_id)
}

// StakingManager is a free data retrieval call binding the contract method 0x22828cc2.
//
// Solidity: function stakingManager() view returns(address)
func (_SubnetFactory *SubnetFactoryCaller) StakingManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "stakingManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingManager is a free data retrieval call binding the contract method 0x22828cc2.
//
// Solidity: function stakingManager() view returns(address)
func (_SubnetFactory *SubnetFactorySession) StakingManager() (common.Address, error) {
	return _SubnetFactory.Contract.StakingManager(&_SubnetFactory.CallOpts)
}

// StakingManager is a free data retrieval call binding the contract method 0x22828cc2.
//
// Solidity: function stakingManager() view returns(address)
func (_SubnetFactory *SubnetFactoryCallerSession) StakingManager() (common.Address, error) {
	return _SubnetFactory.Contract.StakingManager(&_SubnetFactory.CallOpts)
}

// SubnetContracts is a free data retrieval call binding the contract method 0x81f9fb19.
//
// Solidity: function subnetContracts(bytes32 ) view returns(address)
func (_SubnetFactory *SubnetFactoryCaller) SubnetContracts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "subnetContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubnetContracts is a free data retrieval call binding the contract method 0x81f9fb19.
//
// Solidity: function subnetContracts(bytes32 ) view returns(address)
func (_SubnetFactory *SubnetFactorySession) SubnetContracts(arg0 [32]byte) (common.Address, error) {
	return _SubnetFactory.Contract.SubnetContracts(&_SubnetFactory.CallOpts, arg0)
}

// SubnetContracts is a free data retrieval call binding the contract method 0x81f9fb19.
//
// Solidity: function subnetContracts(bytes32 ) view returns(address)
func (_SubnetFactory *SubnetFactoryCallerSession) SubnetContracts(arg0 [32]byte) (common.Address, error) {
	return _SubnetFactory.Contract.SubnetContracts(&_SubnetFactory.CallOpts, arg0)
}

// SubnetExists is a free data retrieval call binding the contract method 0x9587feb2.
//
// Solidity: function subnetExists(bytes32 subnet_id) view returns(bool exists)
func (_SubnetFactory *SubnetFactoryCaller) SubnetExists(opts *bind.CallOpts, subnet_id [32]byte) (bool, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "subnetExists", subnet_id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SubnetExists is a free data retrieval call binding the contract method 0x9587feb2.
//
// Solidity: function subnetExists(bytes32 subnet_id) view returns(bool exists)
func (_SubnetFactory *SubnetFactorySession) SubnetExists(subnet_id [32]byte) (bool, error) {
	return _SubnetFactory.Contract.SubnetExists(&_SubnetFactory.CallOpts, subnet_id)
}

// SubnetExists is a free data retrieval call binding the contract method 0x9587feb2.
//
// Solidity: function subnetExists(bytes32 subnet_id) view returns(bool exists)
func (_SubnetFactory *SubnetFactoryCallerSession) SubnetExists(subnet_id [32]byte) (bool, error) {
	return _SubnetFactory.Contract.SubnetExists(&_SubnetFactory.CallOpts, subnet_id)
}

// SubnetImplementation is a free data retrieval call binding the contract method 0xaabf2e0d.
//
// Solidity: function subnetImplementation() view returns(address)
func (_SubnetFactory *SubnetFactoryCaller) SubnetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "subnetImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SubnetImplementation is a free data retrieval call binding the contract method 0xaabf2e0d.
//
// Solidity: function subnetImplementation() view returns(address)
func (_SubnetFactory *SubnetFactorySession) SubnetImplementation() (common.Address, error) {
	return _SubnetFactory.Contract.SubnetImplementation(&_SubnetFactory.CallOpts)
}

// SubnetImplementation is a free data retrieval call binding the contract method 0xaabf2e0d.
//
// Solidity: function subnetImplementation() view returns(address)
func (_SubnetFactory *SubnetFactoryCallerSession) SubnetImplementation() (common.Address, error) {
	return _SubnetFactory.Contract.SubnetImplementation(&_SubnetFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SubnetFactory *SubnetFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SubnetFactory *SubnetFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SubnetFactory.Contract.SupportsInterface(&_SubnetFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SubnetFactory *SubnetFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SubnetFactory.Contract.SupportsInterface(&_SubnetFactory.CallOpts, interfaceId)
}

// TotalCreated is a free data retrieval call binding the contract method 0x844e0acd.
//
// Solidity: function totalCreated() view returns(uint256)
func (_SubnetFactory *SubnetFactoryCaller) TotalCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SubnetFactory.contract.Call(opts, &out, "totalCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCreated is a free data retrieval call binding the contract method 0x844e0acd.
//
// Solidity: function totalCreated() view returns(uint256)
func (_SubnetFactory *SubnetFactorySession) TotalCreated() (*big.Int, error) {
	return _SubnetFactory.Contract.TotalCreated(&_SubnetFactory.CallOpts)
}

// TotalCreated is a free data retrieval call binding the contract method 0x844e0acd.
//
// Solidity: function totalCreated() view returns(uint256)
func (_SubnetFactory *SubnetFactoryCallerSession) TotalCreated() (*big.Int, error) {
	return _SubnetFactory.Contract.TotalCreated(&_SubnetFactory.CallOpts)
}

// CreateSubnet is a paid mutator transaction binding the contract method 0x15866eaf.
//
// Solidity: function createSubnet((string,address,uint16,string,string,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,bool,bool) create_info) payable returns(address, bytes32)
func (_SubnetFactory *SubnetFactoryTransactor) CreateSubnet(opts *bind.TransactOpts, create_info DataStructuresCreateSubnetInfo) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "createSubnet", create_info)
}

// CreateSubnet is a paid mutator transaction binding the contract method 0x15866eaf.
//
// Solidity: function createSubnet((string,address,uint16,string,string,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,bool,bool) create_info) payable returns(address, bytes32)
func (_SubnetFactory *SubnetFactorySession) CreateSubnet(create_info DataStructuresCreateSubnetInfo) (*types.Transaction, error) {
	return _SubnetFactory.Contract.CreateSubnet(&_SubnetFactory.TransactOpts, create_info)
}

// CreateSubnet is a paid mutator transaction binding the contract method 0x15866eaf.
//
// Solidity: function createSubnet((string,address,uint16,string,string,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,bool,bool) create_info) payable returns(address, bytes32)
func (_SubnetFactory *SubnetFactoryTransactorSession) CreateSubnet(create_info DataStructuresCreateSubnetInfo) (*types.Transaction, error) {
	return _SubnetFactory.Contract.CreateSubnet(&_SubnetFactory.TransactOpts, create_info)
}

// DeprecateSubnet is a paid mutator transaction binding the contract method 0xf3e434f6.
//
// Solidity: function deprecateSubnet(bytes32 subnet_id, string reason) returns()
func (_SubnetFactory *SubnetFactoryTransactor) DeprecateSubnet(opts *bind.TransactOpts, subnet_id [32]byte, reason string) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "deprecateSubnet", subnet_id, reason)
}

// DeprecateSubnet is a paid mutator transaction binding the contract method 0xf3e434f6.
//
// Solidity: function deprecateSubnet(bytes32 subnet_id, string reason) returns()
func (_SubnetFactory *SubnetFactorySession) DeprecateSubnet(subnet_id [32]byte, reason string) (*types.Transaction, error) {
	return _SubnetFactory.Contract.DeprecateSubnet(&_SubnetFactory.TransactOpts, subnet_id, reason)
}

// DeprecateSubnet is a paid mutator transaction binding the contract method 0xf3e434f6.
//
// Solidity: function deprecateSubnet(bytes32 subnet_id, string reason) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) DeprecateSubnet(subnet_id [32]byte, reason string) (*types.Transaction, error) {
	return _SubnetFactory.Contract.DeprecateSubnet(&_SubnetFactory.TransactOpts, subnet_id, reason)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_SubnetFactory *SubnetFactoryTransactor) EmergencyPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "emergencyPause")
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_SubnetFactory *SubnetFactorySession) EmergencyPause() (*types.Transaction, error) {
	return _SubnetFactory.Contract.EmergencyPause(&_SubnetFactory.TransactOpts)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) EmergencyPause() (*types.Transaction, error) {
	return _SubnetFactory.Contract.EmergencyPause(&_SubnetFactory.TransactOpts)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_SubnetFactory *SubnetFactoryTransactor) EmergencyUnpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "emergencyUnpause")
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_SubnetFactory *SubnetFactorySession) EmergencyUnpause() (*types.Transaction, error) {
	return _SubnetFactory.Contract.EmergencyUnpause(&_SubnetFactory.TransactOpts)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) EmergencyUnpause() (*types.Transaction, error) {
	return _SubnetFactory.Contract.EmergencyUnpause(&_SubnetFactory.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SubnetFactory *SubnetFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SubnetFactory *SubnetFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.GrantRole(&_SubnetFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.GrantRole(&_SubnetFactory.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _guardian, address _subnetImplementation, address _stakingManager) returns()
func (_SubnetFactory *SubnetFactoryTransactor) Initialize(opts *bind.TransactOpts, _guardian common.Address, _subnetImplementation common.Address, _stakingManager common.Address) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "initialize", _guardian, _subnetImplementation, _stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _guardian, address _subnetImplementation, address _stakingManager) returns()
func (_SubnetFactory *SubnetFactorySession) Initialize(_guardian common.Address, _subnetImplementation common.Address, _stakingManager common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.Initialize(&_SubnetFactory.TransactOpts, _guardian, _subnetImplementation, _stakingManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _guardian, address _subnetImplementation, address _stakingManager) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) Initialize(_guardian common.Address, _subnetImplementation common.Address, _stakingManager common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.Initialize(&_SubnetFactory.TransactOpts, _guardian, _subnetImplementation, _stakingManager)
}

// PauseSubnet is a paid mutator transaction binding the contract method 0x9b567b22.
//
// Solidity: function pauseSubnet(bytes32 subnet_id, string reason) returns()
func (_SubnetFactory *SubnetFactoryTransactor) PauseSubnet(opts *bind.TransactOpts, subnet_id [32]byte, reason string) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "pauseSubnet", subnet_id, reason)
}

// PauseSubnet is a paid mutator transaction binding the contract method 0x9b567b22.
//
// Solidity: function pauseSubnet(bytes32 subnet_id, string reason) returns()
func (_SubnetFactory *SubnetFactorySession) PauseSubnet(subnet_id [32]byte, reason string) (*types.Transaction, error) {
	return _SubnetFactory.Contract.PauseSubnet(&_SubnetFactory.TransactOpts, subnet_id, reason)
}

// PauseSubnet is a paid mutator transaction binding the contract method 0x9b567b22.
//
// Solidity: function pauseSubnet(bytes32 subnet_id, string reason) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) PauseSubnet(subnet_id [32]byte, reason string) (*types.Transaction, error) {
	return _SubnetFactory.Contract.PauseSubnet(&_SubnetFactory.TransactOpts, subnet_id, reason)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SubnetFactory *SubnetFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SubnetFactory *SubnetFactorySession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.RenounceRole(&_SubnetFactory.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.RenounceRole(&_SubnetFactory.TransactOpts, role, callerConfirmation)
}

// ResumeSubnet is a paid mutator transaction binding the contract method 0xf6f1b76a.
//
// Solidity: function resumeSubnet(bytes32 subnet_id) returns()
func (_SubnetFactory *SubnetFactoryTransactor) ResumeSubnet(opts *bind.TransactOpts, subnet_id [32]byte) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "resumeSubnet", subnet_id)
}

// ResumeSubnet is a paid mutator transaction binding the contract method 0xf6f1b76a.
//
// Solidity: function resumeSubnet(bytes32 subnet_id) returns()
func (_SubnetFactory *SubnetFactorySession) ResumeSubnet(subnet_id [32]byte) (*types.Transaction, error) {
	return _SubnetFactory.Contract.ResumeSubnet(&_SubnetFactory.TransactOpts, subnet_id)
}

// ResumeSubnet is a paid mutator transaction binding the contract method 0xf6f1b76a.
//
// Solidity: function resumeSubnet(bytes32 subnet_id) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) ResumeSubnet(subnet_id [32]byte) (*types.Transaction, error) {
	return _SubnetFactory.Contract.ResumeSubnet(&_SubnetFactory.TransactOpts, subnet_id)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SubnetFactory *SubnetFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SubnetFactory *SubnetFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.RevokeRole(&_SubnetFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.RevokeRole(&_SubnetFactory.TransactOpts, role, account)
}

// SetMinStakeCreateSubnet is a paid mutator transaction binding the contract method 0x4b3d2f01.
//
// Solidity: function setMinStakeCreateSubnet(uint256 amount) returns()
func (_SubnetFactory *SubnetFactoryTransactor) SetMinStakeCreateSubnet(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "setMinStakeCreateSubnet", amount)
}

// SetMinStakeCreateSubnet is a paid mutator transaction binding the contract method 0x4b3d2f01.
//
// Solidity: function setMinStakeCreateSubnet(uint256 amount) returns()
func (_SubnetFactory *SubnetFactorySession) SetMinStakeCreateSubnet(amount *big.Int) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SetMinStakeCreateSubnet(&_SubnetFactory.TransactOpts, amount)
}

// SetMinStakeCreateSubnet is a paid mutator transaction binding the contract method 0x4b3d2f01.
//
// Solidity: function setMinStakeCreateSubnet(uint256 amount) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) SetMinStakeCreateSubnet(amount *big.Int) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SetMinStakeCreateSubnet(&_SubnetFactory.TransactOpts, amount)
}

// SetStakingManager is a paid mutator transaction binding the contract method 0xb00bba6a.
//
// Solidity: function setStakingManager(address _stakingManager) returns()
func (_SubnetFactory *SubnetFactoryTransactor) SetStakingManager(opts *bind.TransactOpts, _stakingManager common.Address) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "setStakingManager", _stakingManager)
}

// SetStakingManager is a paid mutator transaction binding the contract method 0xb00bba6a.
//
// Solidity: function setStakingManager(address _stakingManager) returns()
func (_SubnetFactory *SubnetFactorySession) SetStakingManager(_stakingManager common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SetStakingManager(&_SubnetFactory.TransactOpts, _stakingManager)
}

// SetStakingManager is a paid mutator transaction binding the contract method 0xb00bba6a.
//
// Solidity: function setStakingManager(address _stakingManager) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) SetStakingManager(_stakingManager common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SetStakingManager(&_SubnetFactory.TransactOpts, _stakingManager)
}

// SetSubnetImplementation is a paid mutator transaction binding the contract method 0x1b00326c.
//
// Solidity: function setSubnetImplementation(address new_implementation) returns()
func (_SubnetFactory *SubnetFactoryTransactor) SetSubnetImplementation(opts *bind.TransactOpts, new_implementation common.Address) (*types.Transaction, error) {
	return _SubnetFactory.contract.Transact(opts, "setSubnetImplementation", new_implementation)
}

// SetSubnetImplementation is a paid mutator transaction binding the contract method 0x1b00326c.
//
// Solidity: function setSubnetImplementation(address new_implementation) returns()
func (_SubnetFactory *SubnetFactorySession) SetSubnetImplementation(new_implementation common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SetSubnetImplementation(&_SubnetFactory.TransactOpts, new_implementation)
}

// SetSubnetImplementation is a paid mutator transaction binding the contract method 0x1b00326c.
//
// Solidity: function setSubnetImplementation(address new_implementation) returns()
func (_SubnetFactory *SubnetFactoryTransactorSession) SetSubnetImplementation(new_implementation common.Address) (*types.Transaction, error) {
	return _SubnetFactory.Contract.SetSubnetImplementation(&_SubnetFactory.TransactOpts, new_implementation)
}

// SubnetFactoryGlobalParamsUpdatedIterator is returned from FilterGlobalParamsUpdated and is used to iterate over the raw logs and unpacked data for GlobalParamsUpdated events raised by the SubnetFactory contract.
type SubnetFactoryGlobalParamsUpdatedIterator struct {
	Event *SubnetFactoryGlobalParamsUpdated // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryGlobalParamsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryGlobalParamsUpdated)
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
		it.Event = new(SubnetFactoryGlobalParamsUpdated)
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
func (it *SubnetFactoryGlobalParamsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryGlobalParamsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryGlobalParamsUpdated represents a GlobalParamsUpdated event raised by the SubnetFactory contract.
type SubnetFactoryGlobalParamsUpdated struct {
	ParamsHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGlobalParamsUpdated is a free log retrieval operation binding the contract event 0xf69f719ff818553170dad167d4c56211535f6efc673ed1231b0383ffbbe0ac0b.
//
// Solidity: event GlobalParamsUpdated(bytes32 params_hash)
func (_SubnetFactory *SubnetFactoryFilterer) FilterGlobalParamsUpdated(opts *bind.FilterOpts) (*SubnetFactoryGlobalParamsUpdatedIterator, error) {

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "GlobalParamsUpdated")
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryGlobalParamsUpdatedIterator{contract: _SubnetFactory.contract, event: "GlobalParamsUpdated", logs: logs, sub: sub}, nil
}

// WatchGlobalParamsUpdated is a free log subscription operation binding the contract event 0xf69f719ff818553170dad167d4c56211535f6efc673ed1231b0383ffbbe0ac0b.
//
// Solidity: event GlobalParamsUpdated(bytes32 params_hash)
func (_SubnetFactory *SubnetFactoryFilterer) WatchGlobalParamsUpdated(opts *bind.WatchOpts, sink chan<- *SubnetFactoryGlobalParamsUpdated) (event.Subscription, error) {

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "GlobalParamsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryGlobalParamsUpdated)
				if err := _SubnetFactory.contract.UnpackLog(event, "GlobalParamsUpdated", log); err != nil {
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

// ParseGlobalParamsUpdated is a log parse operation binding the contract event 0xf69f719ff818553170dad167d4c56211535f6efc673ed1231b0383ffbbe0ac0b.
//
// Solidity: event GlobalParamsUpdated(bytes32 params_hash)
func (_SubnetFactory *SubnetFactoryFilterer) ParseGlobalParamsUpdated(log types.Log) (*SubnetFactoryGlobalParamsUpdated, error) {
	event := new(SubnetFactoryGlobalParamsUpdated)
	if err := _SubnetFactory.contract.UnpackLog(event, "GlobalParamsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SubnetFactory contract.
type SubnetFactoryInitializedIterator struct {
	Event *SubnetFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryInitialized)
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
		it.Event = new(SubnetFactoryInitialized)
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
func (it *SubnetFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryInitialized represents a Initialized event raised by the SubnetFactory contract.
type SubnetFactoryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SubnetFactory *SubnetFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*SubnetFactoryInitializedIterator, error) {

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryInitializedIterator{contract: _SubnetFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SubnetFactory *SubnetFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SubnetFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryInitialized)
				if err := _SubnetFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SubnetFactory *SubnetFactoryFilterer) ParseInitialized(log types.Log) (*SubnetFactoryInitialized, error) {
	event := new(SubnetFactoryInitialized)
	if err := _SubnetFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactoryMinStakeCreateSubnetUpdatedIterator is returned from FilterMinStakeCreateSubnetUpdated and is used to iterate over the raw logs and unpacked data for MinStakeCreateSubnetUpdated events raised by the SubnetFactory contract.
type SubnetFactoryMinStakeCreateSubnetUpdatedIterator struct {
	Event *SubnetFactoryMinStakeCreateSubnetUpdated // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryMinStakeCreateSubnetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryMinStakeCreateSubnetUpdated)
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
		it.Event = new(SubnetFactoryMinStakeCreateSubnetUpdated)
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
func (it *SubnetFactoryMinStakeCreateSubnetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryMinStakeCreateSubnetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryMinStakeCreateSubnetUpdated represents a MinStakeCreateSubnetUpdated event raised by the SubnetFactory contract.
type SubnetFactoryMinStakeCreateSubnetUpdated struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinStakeCreateSubnetUpdated is a free log retrieval operation binding the contract event 0xc0cfb9ca495321c534e9d87f33634ed5e47ee716e5c306e619124b1155f1acb3.
//
// Solidity: event MinStakeCreateSubnetUpdated(uint256 oldAmount, uint256 newAmount)
func (_SubnetFactory *SubnetFactoryFilterer) FilterMinStakeCreateSubnetUpdated(opts *bind.FilterOpts) (*SubnetFactoryMinStakeCreateSubnetUpdatedIterator, error) {

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "MinStakeCreateSubnetUpdated")
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryMinStakeCreateSubnetUpdatedIterator{contract: _SubnetFactory.contract, event: "MinStakeCreateSubnetUpdated", logs: logs, sub: sub}, nil
}

// WatchMinStakeCreateSubnetUpdated is a free log subscription operation binding the contract event 0xc0cfb9ca495321c534e9d87f33634ed5e47ee716e5c306e619124b1155f1acb3.
//
// Solidity: event MinStakeCreateSubnetUpdated(uint256 oldAmount, uint256 newAmount)
func (_SubnetFactory *SubnetFactoryFilterer) WatchMinStakeCreateSubnetUpdated(opts *bind.WatchOpts, sink chan<- *SubnetFactoryMinStakeCreateSubnetUpdated) (event.Subscription, error) {

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "MinStakeCreateSubnetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryMinStakeCreateSubnetUpdated)
				if err := _SubnetFactory.contract.UnpackLog(event, "MinStakeCreateSubnetUpdated", log); err != nil {
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

// ParseMinStakeCreateSubnetUpdated is a log parse operation binding the contract event 0xc0cfb9ca495321c534e9d87f33634ed5e47ee716e5c306e619124b1155f1acb3.
//
// Solidity: event MinStakeCreateSubnetUpdated(uint256 oldAmount, uint256 newAmount)
func (_SubnetFactory *SubnetFactoryFilterer) ParseMinStakeCreateSubnetUpdated(log types.Log) (*SubnetFactoryMinStakeCreateSubnetUpdated, error) {
	event := new(SubnetFactoryMinStakeCreateSubnetUpdated)
	if err := _SubnetFactory.contract.UnpackLog(event, "MinStakeCreateSubnetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactoryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SubnetFactory contract.
type SubnetFactoryPausedIterator struct {
	Event *SubnetFactoryPaused // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryPaused)
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
		it.Event = new(SubnetFactoryPaused)
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
func (it *SubnetFactoryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryPaused represents a Paused event raised by the SubnetFactory contract.
type SubnetFactoryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SubnetFactory *SubnetFactoryFilterer) FilterPaused(opts *bind.FilterOpts) (*SubnetFactoryPausedIterator, error) {

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryPausedIterator{contract: _SubnetFactory.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SubnetFactory *SubnetFactoryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SubnetFactoryPaused) (event.Subscription, error) {

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryPaused)
				if err := _SubnetFactory.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_SubnetFactory *SubnetFactoryFilterer) ParsePaused(log types.Log) (*SubnetFactoryPaused, error) {
	event := new(SubnetFactoryPaused)
	if err := _SubnetFactory.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SubnetFactory contract.
type SubnetFactoryRoleAdminChangedIterator struct {
	Event *SubnetFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryRoleAdminChanged)
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
		it.Event = new(SubnetFactoryRoleAdminChanged)
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
func (it *SubnetFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the SubnetFactory contract.
type SubnetFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SubnetFactory *SubnetFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SubnetFactoryRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryRoleAdminChangedIterator{contract: _SubnetFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SubnetFactory *SubnetFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SubnetFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryRoleAdminChanged)
				if err := _SubnetFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SubnetFactory *SubnetFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*SubnetFactoryRoleAdminChanged, error) {
	event := new(SubnetFactoryRoleAdminChanged)
	if err := _SubnetFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SubnetFactory contract.
type SubnetFactoryRoleGrantedIterator struct {
	Event *SubnetFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryRoleGranted)
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
		it.Event = new(SubnetFactoryRoleGranted)
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
func (it *SubnetFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryRoleGranted represents a RoleGranted event raised by the SubnetFactory contract.
type SubnetFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubnetFactory *SubnetFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubnetFactoryRoleGrantedIterator, error) {

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

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryRoleGrantedIterator{contract: _SubnetFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubnetFactory *SubnetFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SubnetFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryRoleGranted)
				if err := _SubnetFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SubnetFactory *SubnetFactoryFilterer) ParseRoleGranted(log types.Log) (*SubnetFactoryRoleGranted, error) {
	event := new(SubnetFactoryRoleGranted)
	if err := _SubnetFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SubnetFactory contract.
type SubnetFactoryRoleRevokedIterator struct {
	Event *SubnetFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryRoleRevoked)
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
		it.Event = new(SubnetFactoryRoleRevoked)
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
func (it *SubnetFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryRoleRevoked represents a RoleRevoked event raised by the SubnetFactory contract.
type SubnetFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubnetFactory *SubnetFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubnetFactoryRoleRevokedIterator, error) {

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

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryRoleRevokedIterator{contract: _SubnetFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SubnetFactory *SubnetFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SubnetFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryRoleRevoked)
				if err := _SubnetFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SubnetFactory *SubnetFactoryFilterer) ParseRoleRevoked(log types.Log) (*SubnetFactoryRoleRevoked, error) {
	event := new(SubnetFactoryRoleRevoked)
	if err := _SubnetFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactorySubnetCreatedIterator is returned from FilterSubnetCreated and is used to iterate over the raw logs and unpacked data for SubnetCreated events raised by the SubnetFactory contract.
type SubnetFactorySubnetCreatedIterator struct {
	Event *SubnetFactorySubnetCreated // Event containing the contract specifics and raw log

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
func (it *SubnetFactorySubnetCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactorySubnetCreated)
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
		it.Event = new(SubnetFactorySubnetCreated)
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
func (it *SubnetFactorySubnetCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactorySubnetCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactorySubnetCreated represents a SubnetCreated event raised by the SubnetFactory contract.
type SubnetFactorySubnetCreated struct {
	SubnetId       [32]byte
	SubnetContract common.Address
	Owner          common.Address
	CanonicalName  string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSubnetCreated is a free log retrieval operation binding the contract event 0x3a0eaad0a112681d755b277092b10d60c0c2f186622548d539b6d778b0876570.
//
// Solidity: event SubnetCreated(bytes32 indexed subnet_id, address indexed subnet_contract, address indexed owner, string canonical_name)
func (_SubnetFactory *SubnetFactoryFilterer) FilterSubnetCreated(opts *bind.FilterOpts, subnet_id [][32]byte, subnet_contract []common.Address, owner []common.Address) (*SubnetFactorySubnetCreatedIterator, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}
	var subnet_contractRule []interface{}
	for _, subnet_contractItem := range subnet_contract {
		subnet_contractRule = append(subnet_contractRule, subnet_contractItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "SubnetCreated", subnet_idRule, subnet_contractRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFactorySubnetCreatedIterator{contract: _SubnetFactory.contract, event: "SubnetCreated", logs: logs, sub: sub}, nil
}

// WatchSubnetCreated is a free log subscription operation binding the contract event 0x3a0eaad0a112681d755b277092b10d60c0c2f186622548d539b6d778b0876570.
//
// Solidity: event SubnetCreated(bytes32 indexed subnet_id, address indexed subnet_contract, address indexed owner, string canonical_name)
func (_SubnetFactory *SubnetFactoryFilterer) WatchSubnetCreated(opts *bind.WatchOpts, sink chan<- *SubnetFactorySubnetCreated, subnet_id [][32]byte, subnet_contract []common.Address, owner []common.Address) (event.Subscription, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}
	var subnet_contractRule []interface{}
	for _, subnet_contractItem := range subnet_contract {
		subnet_contractRule = append(subnet_contractRule, subnet_contractItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "SubnetCreated", subnet_idRule, subnet_contractRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactorySubnetCreated)
				if err := _SubnetFactory.contract.UnpackLog(event, "SubnetCreated", log); err != nil {
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

// ParseSubnetCreated is a log parse operation binding the contract event 0x3a0eaad0a112681d755b277092b10d60c0c2f186622548d539b6d778b0876570.
//
// Solidity: event SubnetCreated(bytes32 indexed subnet_id, address indexed subnet_contract, address indexed owner, string canonical_name)
func (_SubnetFactory *SubnetFactoryFilterer) ParseSubnetCreated(log types.Log) (*SubnetFactorySubnetCreated, error) {
	event := new(SubnetFactorySubnetCreated)
	if err := _SubnetFactory.contract.UnpackLog(event, "SubnetCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactorySubnetDeprecatedIterator is returned from FilterSubnetDeprecated and is used to iterate over the raw logs and unpacked data for SubnetDeprecated events raised by the SubnetFactory contract.
type SubnetFactorySubnetDeprecatedIterator struct {
	Event *SubnetFactorySubnetDeprecated // Event containing the contract specifics and raw log

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
func (it *SubnetFactorySubnetDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactorySubnetDeprecated)
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
		it.Event = new(SubnetFactorySubnetDeprecated)
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
func (it *SubnetFactorySubnetDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactorySubnetDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactorySubnetDeprecated represents a SubnetDeprecated event raised by the SubnetFactory contract.
type SubnetFactorySubnetDeprecated struct {
	SubnetId [32]byte
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubnetDeprecated is a free log retrieval operation binding the contract event 0xc7c42a68a787ba94abb3dab5ab0a94a4f15eb4a9083e4649f0cbaa66dcc5816e.
//
// Solidity: event SubnetDeprecated(bytes32 indexed subnet_id, string reason)
func (_SubnetFactory *SubnetFactoryFilterer) FilterSubnetDeprecated(opts *bind.FilterOpts, subnet_id [][32]byte) (*SubnetFactorySubnetDeprecatedIterator, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "SubnetDeprecated", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFactorySubnetDeprecatedIterator{contract: _SubnetFactory.contract, event: "SubnetDeprecated", logs: logs, sub: sub}, nil
}

// WatchSubnetDeprecated is a free log subscription operation binding the contract event 0xc7c42a68a787ba94abb3dab5ab0a94a4f15eb4a9083e4649f0cbaa66dcc5816e.
//
// Solidity: event SubnetDeprecated(bytes32 indexed subnet_id, string reason)
func (_SubnetFactory *SubnetFactoryFilterer) WatchSubnetDeprecated(opts *bind.WatchOpts, sink chan<- *SubnetFactorySubnetDeprecated, subnet_id [][32]byte) (event.Subscription, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "SubnetDeprecated", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactorySubnetDeprecated)
				if err := _SubnetFactory.contract.UnpackLog(event, "SubnetDeprecated", log); err != nil {
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

// ParseSubnetDeprecated is a log parse operation binding the contract event 0xc7c42a68a787ba94abb3dab5ab0a94a4f15eb4a9083e4649f0cbaa66dcc5816e.
//
// Solidity: event SubnetDeprecated(bytes32 indexed subnet_id, string reason)
func (_SubnetFactory *SubnetFactoryFilterer) ParseSubnetDeprecated(log types.Log) (*SubnetFactorySubnetDeprecated, error) {
	event := new(SubnetFactorySubnetDeprecated)
	if err := _SubnetFactory.contract.UnpackLog(event, "SubnetDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactorySubnetPausedIterator is returned from FilterSubnetPaused and is used to iterate over the raw logs and unpacked data for SubnetPaused events raised by the SubnetFactory contract.
type SubnetFactorySubnetPausedIterator struct {
	Event *SubnetFactorySubnetPaused // Event containing the contract specifics and raw log

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
func (it *SubnetFactorySubnetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactorySubnetPaused)
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
		it.Event = new(SubnetFactorySubnetPaused)
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
func (it *SubnetFactorySubnetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactorySubnetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactorySubnetPaused represents a SubnetPaused event raised by the SubnetFactory contract.
type SubnetFactorySubnetPaused struct {
	SubnetId [32]byte
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubnetPaused is a free log retrieval operation binding the contract event 0xc488d48cb0aa0278ef9ba6c630ace872f1bc50eeb5e7d9446213455571598a77.
//
// Solidity: event SubnetPaused(bytes32 indexed subnet_id, string reason)
func (_SubnetFactory *SubnetFactoryFilterer) FilterSubnetPaused(opts *bind.FilterOpts, subnet_id [][32]byte) (*SubnetFactorySubnetPausedIterator, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "SubnetPaused", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFactorySubnetPausedIterator{contract: _SubnetFactory.contract, event: "SubnetPaused", logs: logs, sub: sub}, nil
}

// WatchSubnetPaused is a free log subscription operation binding the contract event 0xc488d48cb0aa0278ef9ba6c630ace872f1bc50eeb5e7d9446213455571598a77.
//
// Solidity: event SubnetPaused(bytes32 indexed subnet_id, string reason)
func (_SubnetFactory *SubnetFactoryFilterer) WatchSubnetPaused(opts *bind.WatchOpts, sink chan<- *SubnetFactorySubnetPaused, subnet_id [][32]byte) (event.Subscription, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "SubnetPaused", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactorySubnetPaused)
				if err := _SubnetFactory.contract.UnpackLog(event, "SubnetPaused", log); err != nil {
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

// ParseSubnetPaused is a log parse operation binding the contract event 0xc488d48cb0aa0278ef9ba6c630ace872f1bc50eeb5e7d9446213455571598a77.
//
// Solidity: event SubnetPaused(bytes32 indexed subnet_id, string reason)
func (_SubnetFactory *SubnetFactoryFilterer) ParseSubnetPaused(log types.Log) (*SubnetFactorySubnetPaused, error) {
	event := new(SubnetFactorySubnetPaused)
	if err := _SubnetFactory.contract.UnpackLog(event, "SubnetPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactorySubnetResumedIterator is returned from FilterSubnetResumed and is used to iterate over the raw logs and unpacked data for SubnetResumed events raised by the SubnetFactory contract.
type SubnetFactorySubnetResumedIterator struct {
	Event *SubnetFactorySubnetResumed // Event containing the contract specifics and raw log

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
func (it *SubnetFactorySubnetResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactorySubnetResumed)
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
		it.Event = new(SubnetFactorySubnetResumed)
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
func (it *SubnetFactorySubnetResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactorySubnetResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactorySubnetResumed represents a SubnetResumed event raised by the SubnetFactory contract.
type SubnetFactorySubnetResumed struct {
	SubnetId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubnetResumed is a free log retrieval operation binding the contract event 0xaca29e7fd1ca4a247b79568dcc8b3baf96de82f9258b10e76832cd381ddf28f8.
//
// Solidity: event SubnetResumed(bytes32 indexed subnet_id)
func (_SubnetFactory *SubnetFactoryFilterer) FilterSubnetResumed(opts *bind.FilterOpts, subnet_id [][32]byte) (*SubnetFactorySubnetResumedIterator, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "SubnetResumed", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFactorySubnetResumedIterator{contract: _SubnetFactory.contract, event: "SubnetResumed", logs: logs, sub: sub}, nil
}

// WatchSubnetResumed is a free log subscription operation binding the contract event 0xaca29e7fd1ca4a247b79568dcc8b3baf96de82f9258b10e76832cd381ddf28f8.
//
// Solidity: event SubnetResumed(bytes32 indexed subnet_id)
func (_SubnetFactory *SubnetFactoryFilterer) WatchSubnetResumed(opts *bind.WatchOpts, sink chan<- *SubnetFactorySubnetResumed, subnet_id [][32]byte) (event.Subscription, error) {

	var subnet_idRule []interface{}
	for _, subnet_idItem := range subnet_id {
		subnet_idRule = append(subnet_idRule, subnet_idItem)
	}

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "SubnetResumed", subnet_idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactorySubnetResumed)
				if err := _SubnetFactory.contract.UnpackLog(event, "SubnetResumed", log); err != nil {
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

// ParseSubnetResumed is a log parse operation binding the contract event 0xaca29e7fd1ca4a247b79568dcc8b3baf96de82f9258b10e76832cd381ddf28f8.
//
// Solidity: event SubnetResumed(bytes32 indexed subnet_id)
func (_SubnetFactory *SubnetFactoryFilterer) ParseSubnetResumed(log types.Log) (*SubnetFactorySubnetResumed, error) {
	event := new(SubnetFactorySubnetResumed)
	if err := _SubnetFactory.contract.UnpackLog(event, "SubnetResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFactoryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SubnetFactory contract.
type SubnetFactoryUnpausedIterator struct {
	Event *SubnetFactoryUnpaused // Event containing the contract specifics and raw log

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
func (it *SubnetFactoryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFactoryUnpaused)
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
		it.Event = new(SubnetFactoryUnpaused)
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
func (it *SubnetFactoryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFactoryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFactoryUnpaused represents a Unpaused event raised by the SubnetFactory contract.
type SubnetFactoryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SubnetFactory *SubnetFactoryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SubnetFactoryUnpausedIterator, error) {

	logs, sub, err := _SubnetFactory.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SubnetFactoryUnpausedIterator{contract: _SubnetFactory.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SubnetFactory *SubnetFactoryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SubnetFactoryUnpaused) (event.Subscription, error) {

	logs, sub, err := _SubnetFactory.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFactoryUnpaused)
				if err := _SubnetFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_SubnetFactory *SubnetFactoryFilterer) ParseUnpaused(log types.Log) (*SubnetFactoryUnpaused, error) {
	event := new(SubnetFactoryUnpaused)
	if err := _SubnetFactory.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
