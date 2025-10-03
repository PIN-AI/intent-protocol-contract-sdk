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

// SubnetFactoryMetaData contains all meta data concerning the SubnetFactory contract.
var SubnetFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"}],\"name\":\"GlobalParamsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MinStakeCreateSubnetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"subnet_contract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"}],\"name\":\"SubnetCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"SubnetDeprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"SubnetPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"SubnetResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARDIAN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allSubnets\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"}],\"internalType\":\"structDataStructures.CreateSubnetInfo\",\"name\":\"create_info\",\"type\":\"tuple\"}],\"name\":\"createSubnet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"deprecateSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"}],\"name\":\"findSubnetByName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveSubnetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllSubnetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo[]\",\"name\":\"infos\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"validator_counts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"agent_counts\",\"type\":\"uint256[]\"},{\"internalType\":\"enumDataStructures.SubnetStatus[]\",\"name\":\"statuses\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCreationStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"total_created\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStakeCreateSubnet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_stakingManager\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"subnet_contract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubnetImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"validator_count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agent_count\",\"type\":\"uint256\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"subnet_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetStakeConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"getSubnetThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"threshold\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getSubnetsByOwner\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"subnet_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"getSubnetsByStatus\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"subnet_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSubnetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guardian\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_subnetImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_stakingManager\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"isSubnetActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listAllSubnets\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"subnet_ids\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"pauseSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"predictSubnetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"predicted_address\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"resumeSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMinStakeCreateSubnet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingManager\",\"type\":\"address\"}],\"name\":\"setStakingManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_implementation\",\"type\":\"address\"}],\"name\":\"setSubnetImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingManager\",\"outputs\":[{\"internalType\":\"contractIStakingManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subnetContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"}],\"name\":\"subnetExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subnetImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50613caa8061001f6000396000f3fe6080604052600436106101e15760003560e01c806301ffc9a7146101e65780630ffdeb6a1461021b578063110bc3c01461024857806315866eaf1461025d5780631b00326c1461027e5780631f039689146102a057806322828cc2146102bf578063248a9ca3146102ec57806324ea54f41461030c5780632f2ff15d1461032e57806336568abe1461034e57806341c5d3d01461036e57806348f0d0231461038e5780634a4e3bd5146103bb5780634b3d2f01146103d057806351858e27146103f057806354eae4ca146104055780635c975abb1461043257806381f9fb191461044757806382e97eef1461047d578063844e0acd1461049d57806391d14854146104b35780639587feb2146104d35780639b567b221461050b5780639dba09761461052b578063a217fddf14610550578063a2379c9114610565578063a3923d8314610585578063aabf2e0d146105b5578063b00bba6a146105d5578063c0c53b8b146105f5578063c410d56214610615578063c6a1a01a14610635578063d547741f1461064a578063d5a84a111461066a578063df43dc0f146106a0578063e8175c4a146106c0578063ea869cfd146106de578063ee7941d9146106f3578063f3e434f614610711578063f5caee8114610731578063f6f1b76a14610746575b600080fd5b3480156101f257600080fd5b50610206610201366004612c8c565b610766565b60405190151581526020015b60405180910390f35b34801561022757600080fd5b5061023b610236366004612cc3565b61079d565b6040516102129190612ce0565b34801561025457600080fd5b5061023b6109a6565b61027061026b366004612d23565b6109fe565b604051610212929190612d5e565b34801561028a57600080fd5b5061029e610299366004612d8c565b610d86565b005b3480156102ac57600080fd5b506002545b604051908152602001610212565b3480156102cb57600080fd5b506005546102df906001600160a01b031681565b6040516102129190612da9565b3480156102f857600080fd5b506102b1610307366004612dbd565b610e10565b34801561031857600080fd5b506102b1600080516020613c5583398151915281565b34801561033a57600080fd5b5061029e610349366004612dd6565b610e30565b34801561035a57600080fd5b5061029e610369366004612dd6565b610e52565b34801561037a57600080fd5b50610206610389366004612dbd565b610e8a565b34801561039a57600080fd5b506103ae6103a9366004612dbd565b610f3c565b6040516102129190612ea7565b3480156103c757600080fd5b5061029e610fe6565b3480156103dc57600080fd5b5061029e6103eb366004612dbd565b611009565b3480156103fc57600080fd5b5061029e611067565b34801561041157600080fd5b50610425610420366004612dbd565b611087565b6040516102129190612ed2565b34801561043e57600080fd5b50610206611139565b34801561045357600080fd5b506102df610462366004612dbd565b6001602052600090815260409020546001600160a01b031681565b34801561048957600080fd5b5061023b610498366004612d8c565b61114e565b3480156104a957600080fd5b506102b160035481565b3480156104bf57600080fd5b506102066104ce366004612dd6565b61133d565b3480156104df57600080fd5b506102066104ee366004612dbd565b6000908152600160205260409020546001600160a01b0316151590565b34801561051757600080fd5b5061029e610526366004612f28565b611373565b34801561053757600080fd5b50610540611512565b60405161021294939291906131bc565b34801561055c57600080fd5b506102b1600081565b34801561057157600080fd5b506102df610580366004612dbd565b61188c565b34801561059157600080fd5b506105a56105a0366004612dbd565b6118a3565b604051610212949392919061325e565b3480156105c157600080fd5b506000546102df906001600160a01b031681565b3480156105e157600080fd5b5061029e6105f0366004612d8c565b611a37565b34801561060157600080fd5b5061029e610610366004613297565b611a98565b34801561062157600080fd5b506102b1610630366004612dbd565b611cce565b34801561064157600080fd5b506102b1611cef565b34801561065657600080fd5b5061029e610665366004612dd6565b611df0565b34801561067657600080fd5b506102df610685366004612dbd565b6000908152600160205260409020546001600160a01b031690565b3480156106ac57600080fd5b506102b16106bb3660046132e2565b611e0c565b3480156106cc57600080fd5b506000546001600160a01b03166102df565b3480156106ea57600080fd5b506003546102b1565b3480156106ff57600080fd5b506005546001600160a01b03166102df565b34801561071d57600080fd5b5061029e61072c366004612f28565b611f24565b34801561073d57600080fd5b506006546102b1565b34801561075257600080fd5b5061029e610761366004612dbd565b6120bb565b60006001600160e01b03198216637965db0b60e01b148061079757506301ffc9a760e01b6001600160e01b03198316145b92915050565b6002546060906000816001600160401b038111156107bd576107bd613323565b6040519080825280602002602001820160405280156107e6578160200160208202803683370190505b5090506000805b8381101561090e5760006002828154811061080a5761080a613339565b600091825260208083209091015480835260019091526040909120549091506001600160a01b03168061083e575050610906565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa15801561087e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526108a691908101906136ae565b90508860038111156108ba576108ba612fc3565b81610100015160038111156108d1576108d1612fc3565b0361090257828686815181106108e9576108e9613339565b6020908102919091010152846108fe8161387b565b9550505b5050505b6001016107ed565b50806001600160401b0381111561092757610927613323565b604051908082528060200260200182016040528015610950578160200160208202803683370190505b50935060005b8181101561099d5782818151811061097057610970613339565b602002602001015185828151811061098a5761098a613339565b6020908102919091010152600101610956565b50505050919050565b606060028054806020026020016040519081016040528092919081815260200182805480156109f457602002820191906000526020600020905b8154815260200190600101908083116109e0575b5050505050905090565b600080610a09612249565b610a1161227f565b600080610a1d856122a7565b6005549091506001600160a01b0316610a765760405162461bcd60e51b815260206004820152601660248201527514dd185ada5b99d3585b9859d95c881b9bdd081cd95d60521b60448201526064015b60405180910390fd5b600054610a8c906001600160a01b0316826123c9565b91506000610a9a86836123d7565b600554604051633845b9d160e11b81529192506001600160a01b038086169263708b73a292610ad39286923092909116906004016138a2565b600060405180830381600087803b158015610aed57600080fd5b505af1158015610b01573d6000803e3d6000fd5b505060065460055460408051639f9106d160e01b81529051929550600094506001600160a01b039091169250639f9106d19160048083019260209291908290030181865afa158015610b57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b7b91906138d5565b6005549091506001600160a01b03168215610d5c576001600160a01b038216610c605782341015610be55760405162461bcd60e51b8152602060048201526014602482015273125b9d985b1a59081cdd185ad948185b5bdd5b9d60621b6044820152606401610a6d565b6001600160a01b038116632d46f83c84610c0560408c0160208d01612d8c565b60006003886040518663ffffffff1660e01b8152600401610c2994939291906138f2565b6000604051808303818588803b158015610c4257600080fd5b505af1158015610c56573d6000803e3d6000fd5b5050505050610d5c565b610c756001600160a01b0383163330866125f2565b60405163095ea7b360e01b81526001600160a01b0383169063095ea7b390610ca39084908790600401612d5e565b6020604051808303816000875af1158015610cc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ce69190613925565b506001600160a01b038116632d46f83c610d0660408b0160208c01612d8c565b846003876040518563ffffffff1660e01b8152600401610d2994939291906138f2565b600060405180830381600087803b158015610d4357600080fd5b505af1158015610d57573d6000803e3d6000fd5b505050505b610d708486610d6b8b80613942565b61264c565b509294509092505050610d8161270c565b915091565b600080516020613c55833981519152610d9e8161271d565b6001600160a01b038216610ded5760405162461bcd60e51b815260206004820152601660248201527524b73b30b634b21034b6b83632b6b2b73a30ba34b7b760511b6044820152606401610a6d565b50600080546001600160a01b0319166001600160a01b0392909216919091179055565b600080610e1b612727565b60009384526020525050604090206001015490565b610e3982610e10565b610e428161271d565b610e4c838361274b565b50505050565b6001600160a01b0381163314610e7b5760405163334bd91960e11b815260040160405180910390fd5b610e8582826127ec565b505050565b6000818152600160205260408120546001600160a01b031680610eb05750600092915050565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015610ef0573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f1891908101906136ae565b905060018161010001516003811115610f3357610f33612fc3565b14949350505050565b610f44612b8e565b6000828152600160205260409020546001600160a01b031680610f795760405162461bcd60e51b8152600401610a6d90613988565b806001600160a01b031663ce89f1de6040518163ffffffff1660e01b8152600401600060405180830381865afa158015610fb7573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610fdf91908101906139b3565b9392505050565b600080516020613c55833981519152610ffe8161271d565b611006612864565b50565b600080516020613c558339815191526110218161271d565b600680549083905560408051828152602081018590527fc0cfb9ca495321c534e9d87f33634ed5e47ee716e5c306e619124b1155f1acb3910160405180910390a1505050565b600080516020613c5583398151915261107f8161271d565b6110066128bb565b61108f612bd3565b6000828152600160205260409020546001600160a01b0316806110c45760405162461bcd60e51b8152600401610a6d90613988565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611104573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261112c91908101906136ae565b6101400151949350505050565b600080611144612902565b5460ff1692915050565b6002546060906000816001600160401b0381111561116e5761116e613323565b604051908082528060200260200182016040528015611197578160200160208202803683370190505b5090506000805b838110156112ae576000600282815481106111bb576111bb613339565b600091825260208083209091015480835260019091526040909120549091506001600160a01b0316806111ef5750506112a6565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa15801561122f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261125791908101906136ae565b9050886001600160a01b031681604001516001600160a01b0316036112a2578286868151811061128957611289613339565b60209081029190910101528461129e8161387b565b9550505b5050505b60010161119e565b50806001600160401b038111156112c7576112c7613323565b6040519080825280602002602001820160405280156112f0578160200160208202803683370190505b50935060005b8181101561099d5782818151811061131057611310613339565b602002602001015185828151811061132a5761132a613339565b60209081029190910101526001016112f6565b600080611348612727565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b600080516020613c5583398151915261138b8161271d565b6000848152600160205260409020546001600160a01b0316806113c05760405162461bcd60e51b8152600401610a6d90613988565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611400573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261142891908101906136ae565b90506001816101000151600381111561144357611443612fc3565b1461147d5760405162461bcd60e51b815260206004820152600a6024820152694e6f742061637469766560b01b6044820152606401610a6d565b816001600160a01b0316638456cb596040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156114b857600080fd5b505af11580156114cc573d6000803e3d6000fd5b50505050857fc488d48cb0aa0278ef9ba6c630ace872f1bc50eeb5e7d9446213455571598a7786866040516115029291906139e7565b60405180910390a2505050505050565b600254606090819081908190806001600160401b0381111561153657611536613323565b60405190808252806020026020018201604052801561156f57816020015b61155c612bea565b8152602001906001900390816115545790505b509450806001600160401b0381111561158a5761158a613323565b6040519080825280602002602001820160405280156115b3578160200160208202803683370190505b509350806001600160401b038111156115ce576115ce613323565b6040519080825280602002602001820160405280156115f7578160200160208202803683370190505b509250806001600160401b0381111561161257611612613323565b60405190808252806020026020018201604052801561163b578160200160208202803683370190505b50915060005b818110156118845760006002828154811061165e5761165e613339565b600091825260208083209091015480835260019091526040909120549091506001600160a01b0316801561187a5760008190506000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa1580156116d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526116f991908101906136ae565b9050808a868151811061170e5761170e613339565b602002602001018190525080610100015187868151811061173157611731613339565b6020026020010190600381111561174a5761174a612fc3565b9081600381111561175d5761175d612fc3565b90525060405163c0a6367160e01b81526001600160a01b0383169063c0a636719061178d90600090600401613a16565b602060405180830381865afa1580156117aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ce9190613a29565b8986815181106117e0576117e0613339565b602090810291909101015260405163c0a6367160e01b81526001600160a01b0383169063c0a636719061181890600190600401613a16565b602060405180830381865afa158015611835573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118599190613a29565b88868151811061186b5761186b613339565b60200260200101818152505050505b5050600101611641565b505090919293565b60008054610797906001600160a01b031683612926565b6118ab612bea565b600082815260016020526040812054819081906001600160a01b03168015611a2a576000819050806001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611910573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261193891908101906136ae565b60405163c0a6367160e01b81529096506001600160a01b0382169063c0a636719061196890600090600401613a16565b602060405180830381865afa158015611985573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119a99190613a29565b60405163c0a6367160e01b81529095506001600160a01b0382169063c0a63671906119d990600190600401613a16565b602060405180830381865afa1580156119f6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a1a9190613a29565b9350856101000151925050611a2f565b600091505b509193509193565b600080516020613c55833981519152611a4f8161271d565b6001600160a01b038216611a755760405162461bcd60e51b8152600401610a6d90613a42565b50600580546001600160a01b0319166001600160a01b0392909216919091179055565b6000611aa261298f565b805490915060ff600160401b82041615906001600160401b0316600081158015611ac95750825b90506000826001600160401b03166001148015611ae55750303b155b905081158015611af3575080155b15611b115760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b03191660011785558315611b3a57845460ff60401b1916600160401b1785555b6001600160a01b038716611b905760405162461bcd60e51b815260206004820152601d60248201527f496e76616c6964207375626e657420696d706c656d656e746174696f6e0000006044820152606401610a6d565b6001600160a01b038816611be15760405162461bcd60e51b8152602060048201526018602482015277496e76616c696420677561726469616e206164647265737360401b6044820152606401610a6d565b6001600160a01b038616611c075760405162461bcd60e51b8152600401610a6d90613a42565b611c0f6129b8565b611c176129b8565b611c1f6129c0565b611c2a60008961274b565b50611c43600080516020613c558339815191528961274b565b50600080546001600160a01b03808a166001600160a01b0319928316179092556005805492891692909116919091179055655af3107a40006006558315611cc457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b60028181548110611cde57600080fd5b600091825260209091200154905081565b600254600090815b81811015611deb5760006001600060028481548110611d1857611d18613339565b600091825260208083209091015483528201929092526040019020546001600160a01b0316905080611d4a5750611de3565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611d8a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611db291908101906136ae565b905060018161010001516003811115611dcd57611dcd612fc3565b03611de05784611ddc8161387b565b9550505b50505b600101611cf7565b505090565b611df982610e10565b611e028161271d565b610e4c83836127ec565b6000808383604051611e1f929190613a73565b60405190819003902060025490915060005b81811015611f1857600060028281548110611e4e57611e4e613339565b600091825260208083209091015480835260019091526040909120549091506001600160a01b031680611e82575050611f10565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611ec2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611eea91908101906136ae565b90508581602001518051906020012003611f0c57829650505050505050610797565b5050505b600101611e31565b50600095945050505050565b600080516020613c55833981519152611f3c8161271d565b6000848152600160205260409020546001600160a01b031680611f715760405162461bcd60e51b8152600401610a6d90613988565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015611fb1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611fd991908101906136ae565b905060038161010001516003811115611ff457611ff4612fc3565b036120365760405162461bcd60e51b8152602060048201526012602482015271105b1c9958591e4819195c1c9958d85d195960721b6044820152606401610a6d565b816001600160a01b0316630fcc0c286040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561207157600080fd5b505af1158015612085573d6000803e3d6000fd5b50505050857fc7c42a68a787ba94abb3dab5ab0a94a4f15eb4a9083e4649f0cbaa66dcc5816e86866040516115029291906139e7565b600080516020613c558339815191526120d38161271d565b6000828152600160205260409020546001600160a01b0316806121085760405162461bcd60e51b8152600401610a6d90613988565b6000816001600160a01b031663d15442ad6040518163ffffffff1660e01b8152600401600060405180830381865afa158015612148573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261217091908101906136ae565b90506002816101000151600381111561218b5761218b612fc3565b146121c55760405162461bcd60e51b815260206004820152600a602482015269139bdd081c185d5cd95960b21b6044820152606401610a6d565b816001600160a01b0316633f4ba83a6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561220057600080fd5b505af1158015612214573d6000803e3d6000fd5b50506040518692507faca29e7fd1ca4a247b79568dcc8b3baf96de82f9258b10e76832cd381ddf28f89150600090a250505050565b60006122536129d0565b80549091506001190161227957604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b612287611139565b156122a55760405163d93c066560e01b815260040160405180910390fd5b565b6000806122b48380613942565b9050116122f75760405162461bcd60e51b8152602060048201526011602482015270456d707479207375626e6574206e616d6560781b6044820152606401610a6d565b336123086040840160208501612d8c565b6001600160a01b03161461234e5760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b21037bbb732b960991b6044820152606401610a6d565b6004805490600061235e8361387b565b90915550506004546000818152600160205260409020549091506001600160a01b0316156123c45760405162461bcd60e51b815260206004820152601360248201527229bab13732ba1024a21031b7b63634b9b4b7b760691b6044820152606401610a6d565b919050565b6000610fdf838360006129f4565b6123df612bea565b8181526123ec8380613942565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506020808301919091526124399060408501908501612d8c565b6001600160a01b031660408083019190915261245b9060608501908501613a83565b61ffff1660608083019190915261247490840184613942565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050506080808301919091526124bd90840184613942565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060a08083019190915261250c90368590038501908501613aa0565b6101208201526125253684900384016101008501613af6565b6101408083019190915261253b90840184613b39565b61254490613bb7565b6101608083019190915261255a90840184613942565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050610180820152426101a082018190526101c082015260016101008201819052506101808301356101e08201526125cc6101c084016101a08501613c37565b151560e08201526125e56101e084016101c08501613c37565b151560c082015292915050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610e4c908590612a8b565b6000848152600160208190526040822080546001600160a01b0319166001600160a01b03871617905560028054918201815582527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0185905560038054916126b38361387b565b9190505550336001600160a01b0316836001600160a01b0316857f3a0eaad0a112681d755b277092b10d60c0c2f186622548d539b6d778b087657085856040516126fe9291906139e7565b60405180910390a450505050565b60006127166129d0565b6001905550565b6110068133612af3565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b600080612756612727565b9050612762848461133d565b6127e2576000848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556127983390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610797565b6000915050610797565b6000806127f7612727565b9050612803848461133d565b156127e2576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610797565b61286c612b22565b6000612876612902565b805460ff1916815590507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516128b09190612da9565b60405180910390a150565b6128c361227f565b60006128cd612902565b805460ff1916600117815590507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586128a33390565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330090565b6040513060388201526f5af43d82803e903d91602b57fd5bf3ff602482015260148101839052733d602d80600a3d3981f3363d3d373d3d3d363d738152605881018290526037600c820120607882015260556043909101206000906001600160a01b0316610fdf565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610797565b6122a5612b47565b6129c8612b47565b6122a5612b6c565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b600081471015612a205760405163cf47918160e01b815247600482015260248101839052604401610a6d565b763d602d80600a3d3981f3363d3d373d3d3d363d730000008460601b60e81c176000526e5af43d82803e903d91602b57fd5bf38460781b17602052826037600984f590506001600160a01b038116610fdf5760405163b06ebf3d60e01b815260040160405180910390fd5b600080602060008451602086016000885af180612aae576040513d6000823e3d81fd5b50506000513d91508115612ac6578060011415612ad3565b6001600160a01b0384163b155b15610e4c5783604051635274afe760e01b8152600401610a6d9190612da9565b612afd828261133d565b612b1e57808260405163e2517d3f60e01b8152600401610a6d929190612d5e565b5050565b612b2a611139565b6122a557604051638dfc202b60e01b815260040160405180910390fd5b612b4f612b74565b6122a557604051631afcd79f60e31b815260040160405180910390fd5b61270c612b47565b6000612b7e61298f565b54600160401b900460ff16919050565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b604080518082019091526000808252602082015290565b60408051610200810182526000808252606060208084018290528385018390528184018390526080840182905260a0840182905260c0840183905260e084018390526101008401839052845191820185528282528101829052928301526101208101919091526101408101612c5d612bd3565b8152602001612c6a612b8e565b8152602001606081526020016000815260200160008152602001600081525090565b600060208284031215612c9e57600080fd5b81356001600160e01b031981168114610fdf57600080fd5b6004811061100657600080fd5b600060208284031215612cd557600080fd5b8135610fdf81612cb6565b602080825282518282018190526000918401906040840190835b81811015612d18578351835260209384019390920191600101612cfa565b509095945050505050565b600060208284031215612d3557600080fd5b81356001600160401b03811115612d4b57600080fd5b82016101e08185031215610fdf57600080fd5b6001600160a01b03929092168252602082015260400190565b6001600160a01b038116811461100657600080fd5b600060208284031215612d9e57600080fd5b8135610fdf81612d77565b6001600160a01b0391909116815260200190565b600060208284031215612dcf57600080fd5b5035919050565b60008060408385031215612de957600080fd5b823591506020830135612dfb81612d77565b809150509250929050565b600081518084526020840193506020830160005b82811015612e38578151865260209586019590910190600101612e1a565b5093949350505050565b805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c0830152600060e082015161010060e0850152612e9f610100850182612e06565b949350505050565b602081526000610fdf6020830184612e42565b805163ffffffff908116835260209182015116910152565b604081016107978284612eba565b60008083601f840112612ef257600080fd5b5081356001600160401b03811115612f0957600080fd5b602083019150836020828501011115612f2157600080fd5b9250929050565b600080600060408486031215612f3d57600080fd5b8335925060208401356001600160401b03811115612f5a57600080fd5b612f6686828701612ee0565b9497909650939450505050565b60005b83811015612f8e578181015183820152602001612f76565b50506000910152565b60008151808452612faf816020860160208601612f73565b601f01601f19169290920160200192915050565b634e487b7160e01b600052602160045260246000fd5b6004811061100657634e487b7160e01b600052602160045260246000fd5b61300081612fd9565b9052565b80518252600060208201516102606020850152613025610260850182612f97565b9050604083015161304160408601826001600160a01b03169052565b506060830151613057606086018261ffff169052565b506080830151848203608086015261306f8282612f97565b91505060a083015184820360a08601526130898282612f97565b91505060c083015161309f60c086018215159052565b5060e08301516130b360e086018215159052565b506101008301516130c8610100860182612ff7565b506101208381015180516001600160801b03169186019190915260208101516001600160401b0390811661014087015260408201511661016086015250610140830151613119610180860182612eba565b506101608301518482036101c08601526131338282612e42565b9150506101808301518482036101e086015261314f8282612f97565b9150506101a08301516102008501526101c08301516102208501526101e08301516102408501528091505092915050565b600081518084526020840193506020830160005b82811015612e385781516131a781612fd9565b86526020958601959190910190600101613194565b6000608082016080835280875180835260a08501915060a08160051b86010192506020890160005b8281101561321557609f19878603018452613200858351613004565b945060209384019391909101906001016131e4565b50505050828103602084015261322b8187612e06565b9050828103604084015261323f8186612e06565b905082810360608401526132538185613180565b979650505050505050565b6080815260006132716080830187613004565b905084602083015283604083015261328883612fd9565b82606083015295945050505050565b6000806000606084860312156132ac57600080fd5b83356132b781612d77565b925060208401356132c781612d77565b915060408401356132d781612d77565b809150509250925092565b600080602083850312156132f557600080fd5b82356001600160401b0381111561330b57600080fd5b61331785828601612ee0565b90969095509350505050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b604051606081016001600160401b038111828210171561337157613371613323565b60405290565b604080519081016001600160401b038111828210171561337157613371613323565b60405161010081016001600160401b038111828210171561337157613371613323565b60405161020081016001600160401b038111828210171561337157613371613323565b604051601f8201601f191681016001600160401b038111828210171561340757613407613323565b604052919050565b600082601f83011261342057600080fd5b81516001600160401b0381111561343957613439613323565b61344c601f8201601f19166020016133df565b81815284602083860101111561346157600080fd5b612e9f826020830160208701612f73565b80516123c481612d77565b61ffff8116811461100657600080fd5b80516123c48161347d565b801515811461100657600080fd5b80516123c481613498565b80516123c481612cb6565b6001600160801b038116811461100657600080fd5b6001600160401b038116811461100657600080fd5b6000606082840312156134f857600080fd5b61350061334f565b9050815161350d816134bc565b8152602082015161351d816134d1565b60208201526040820151613530816134d1565b604082015292915050565b63ffffffff8116811461100657600080fd5b60006040828403121561355f57600080fd5b613567613377565b905081516135748161353b565b815260208201516135848161353b565b602082015292915050565b60006001600160401b038211156135a8576135a8613323565b5060051b60200190565b600082601f8301126135c357600080fd5b81516135d66135d18261358f565b6133df565b8082825260208201915060208360051b8601019250858311156135f857600080fd5b602085015b838110156136155780518352602092830192016135fd565b5095945050505050565b6000610100828403121561363257600080fd5b61363a613399565b825181526020808401519082015260408084015190820152606080840151908201526080808401519082015260a0808401519082015260c0808401519082015260e08301519091506001600160401b0381111561369657600080fd5b6136a2848285016135b2565b60e08301525092915050565b6000602082840312156136c057600080fd5b81516001600160401b038111156136d657600080fd5b820161026081850312156136e957600080fd5b6136f16133bc565b8151815260208201516001600160401b0381111561370e57600080fd5b61371a8682850161340f565b60208301525061372c60408301613472565b604082015261373d6060830161348d565b606082015260808201516001600160401b0381111561375b57600080fd5b6137678682850161340f565b60808301525060a08201516001600160401b0381111561378657600080fd5b6137928682850161340f565b60a0830152506137a460c083016134a6565b60c08201526137b560e083016134a6565b60e08201526137c761010083016134b1565b6101008201526137db8561012084016134e6565b6101208201526137ef85610180840161354d565b6101408201526101c08201516001600160401b0381111561380f57600080fd5b61381b8682850161361f565b610160830152506101e08201516001600160401b0381111561383c57600080fd5b6138488682850161340f565b610180830152506102008201516101a08201526102208201516101c0820152610240909101516101e08201529392505050565b60006001820161389b57634e487b7160e01b600052601160045260246000fd5b5060010190565b6060815260006138b56060830186613004565b6001600160a01b0394851660208401529290931660409091015292915050565b6000602082840312156138e757600080fd5b8151610fdf81612d77565b6001600160a01b038581168252841660208201526080810161391384612fd9565b60408201939093526060015292915050565b60006020828403121561393757600080fd5b8151610fdf81613498565b6000808335601e1984360301811261395957600080fd5b8301803591506001600160401b0382111561397357600080fd5b602001915036819003821315612f2157600080fd5b6020808252601190820152705375626e6574206e6f742065786973747360781b604082015260600190565b6000602082840312156139c557600080fd5b81516001600160401b038111156139db57600080fd5b612e9f8482850161361f565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60208101613a2383612fd9565b91905290565b600060208284031215613a3b57600080fd5b5051919050565b60208082526017908201527624b73b30b634b21039ba30b5b4b7339036b0b730b3b2b960491b604082015260600190565b8183823760009101908152919050565b600060208284031215613a9557600080fd5b8135610fdf8161347d565b60006060828403128015613ab357600080fd5b50613abc61334f565b8235613ac7816134bc565b81526020830135613ad7816134d1565b60208201526040830135613aea816134d1565b60408201529392505050565b60006040828403128015613b0957600080fd5b50613b12613377565b8235613b1d8161353b565b81526020830135613b2d8161353b565b60208201529392505050565b6000823560fe19833603018112613b4f57600080fd5b9190910192915050565b600082601f830112613b6a57600080fd5b8135613b786135d18261358f565b8082825260208201915060208360051b860101925085831115613b9a57600080fd5b602085015b83811015613615578035835260209283019201613b9f565b60006101008236031215613bca57600080fd5b613bd2613399565b823581526020808401359082015260408084013590820152606080840135908201526080808401359082015260a0808401359082015260c0808401359082015260e08301356001600160401b03811115613c2b57600080fd5b6136a236828601613b59565b600060208284031215613c4957600080fd5b8135610fdf8161349856fe55435dd261a4b9b3364963f7738a7a662ad9c84396d64be3365284bb7f0a5041a264697066735822122057b4f2bf3acbfb67bf36e6d249fbaf7ea0143cd4507cb426fd6b9deebeda7cb264736f6c634300081b0033",
}

// SubnetFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use SubnetFactoryMetaData.ABI instead.
var SubnetFactoryABI = SubnetFactoryMetaData.ABI

// SubnetFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubnetFactoryMetaData.Bin instead.
var SubnetFactoryBin = SubnetFactoryMetaData.Bin

// DeploySubnetFactory deploys a new Ethereum contract, binding an instance of SubnetFactory to it.
func DeploySubnetFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SubnetFactory, error) {
	parsed, err := SubnetFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubnetFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SubnetFactory{SubnetFactoryCaller: SubnetFactoryCaller{contract: contract}, SubnetFactoryTransactor: SubnetFactoryTransactor{contract: contract}, SubnetFactoryFilterer: SubnetFactoryFilterer{contract: contract}}, nil
}

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
