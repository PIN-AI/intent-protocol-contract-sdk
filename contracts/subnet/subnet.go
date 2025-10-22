// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package subnet

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

// DataStructuresParticipantInfo is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresParticipantInfo struct {
	Status       uint8
	AgentId      *big.Int
	RegisteredAt uint64
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

// DataStructuresStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresStakeInfo struct {
	StakedAmount         *big.Int
	UnstakeRequestTime   *big.Int
	UnstakeRequestAmount *big.Int
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

// SubnetMetaData contains all meta data concerning the Subnet contract.
var SubnetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AgentIdUsedByOtherParticipant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDeprecatedSubnet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyRegisteredForRole\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArrayLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidIdentityRegistry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParticipant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxParticipantsReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAgentNFTOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRegisteredInSubnet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSuspended\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ParticipantNotExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingManagerNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubnetCannotBePaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SubnetCannotBeUnpaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"AgentIdentityRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedback_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to_agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.FeedbackType\",\"name\":\"feedback_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rating\",\"type\":\"uint8\"}],\"name\":\"FeedbackSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"}],\"name\":\"ParticipantActivityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participantType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ParticipantApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"ParticipantIdentityBound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"ParticipantIdentityRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"ParticipantRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participantType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rejector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ParticipantRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participantType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"ParticipantStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ParticipantStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"}],\"name\":\"SubnetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"activeParticipantCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"approveParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getActiveParticipantCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAgentIdentityRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidFrequencyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCheckpointPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getMinimumStakeByType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"}],\"name\":\"getParticipantAgentIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"}],\"name\":\"getParticipantByAgentId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getParticipantCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"}],\"name\":\"getParticipantDomain\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"}],\"name\":\"getParticipantEndpoint\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getParticipantInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"agent_id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"registered_at\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.ParticipantInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"}],\"name\":\"getParticipantMetadataUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getParticipantStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestAmount\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequireKYC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignatureThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlashingRates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeGovernanceConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubnetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalParticipantCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnstakeLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"hasAvailableSlots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo\",\"name\":\"_subnetInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"identityRegistry_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"isActiveParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"listActiveParticipants\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"agent_id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"registered_at\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.ParticipantInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"agent_id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"registered_at\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"name\":\"registerParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"name\":\"registerParticipantERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"rejectParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"participant_addrs\",\"type\":\"address[]\"},{\"internalType\":\"enumDataStructures.ParticipantType[]\",\"name\":\"participant_types\",\"type\":\"uint8[]\"}],\"name\":\"requireActiveParticipants\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"resumeParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"_stakeConfig\",\"type\":\"tuple\"}],\"name\":\"setStakeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"suspendParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"updateParticipantActivity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SubnetABI is the input ABI used to generate the binding from.
// Deprecated: Use SubnetMetaData.ABI instead.
var SubnetABI = SubnetMetaData.ABI

// Subnet is an auto generated Go binding around an Ethereum contract.
type Subnet struct {
	SubnetCaller     // Read-only binding to the contract
	SubnetTransactor // Write-only binding to the contract
	SubnetFilterer   // Log filterer for contract events
}

// SubnetCaller is an auto generated read-only Go binding around an Ethereum contract.
type SubnetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SubnetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SubnetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SubnetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SubnetSession struct {
	Contract     *Subnet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubnetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SubnetCallerSession struct {
	Contract *SubnetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SubnetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SubnetTransactorSession struct {
	Contract     *SubnetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SubnetRaw is an auto generated low-level Go binding around an Ethereum contract.
type SubnetRaw struct {
	Contract *Subnet // Generic contract binding to access the raw methods on
}

// SubnetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SubnetCallerRaw struct {
	Contract *SubnetCaller // Generic read-only contract binding to access the raw methods on
}

// SubnetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SubnetTransactorRaw struct {
	Contract *SubnetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSubnet creates a new instance of Subnet, bound to a specific deployed contract.
func NewSubnet(address common.Address, backend bind.ContractBackend) (*Subnet, error) {
	contract, err := bindSubnet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Subnet{SubnetCaller: SubnetCaller{contract: contract}, SubnetTransactor: SubnetTransactor{contract: contract}, SubnetFilterer: SubnetFilterer{contract: contract}}, nil
}

// NewSubnetCaller creates a new read-only instance of Subnet, bound to a specific deployed contract.
func NewSubnetCaller(address common.Address, caller bind.ContractCaller) (*SubnetCaller, error) {
	contract, err := bindSubnet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SubnetCaller{contract: contract}, nil
}

// NewSubnetTransactor creates a new write-only instance of Subnet, bound to a specific deployed contract.
func NewSubnetTransactor(address common.Address, transactor bind.ContractTransactor) (*SubnetTransactor, error) {
	contract, err := bindSubnet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SubnetTransactor{contract: contract}, nil
}

// NewSubnetFilterer creates a new log filterer instance of Subnet, bound to a specific deployed contract.
func NewSubnetFilterer(address common.Address, filterer bind.ContractFilterer) (*SubnetFilterer, error) {
	contract, err := bindSubnet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SubnetFilterer{contract: contract}, nil
}

// bindSubnet binds a generic wrapper to an already deployed contract.
func bindSubnet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SubnetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subnet *SubnetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subnet.Contract.SubnetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subnet *SubnetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subnet.Contract.SubnetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subnet *SubnetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subnet.Contract.SubnetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Subnet *SubnetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Subnet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Subnet *SubnetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subnet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Subnet *SubnetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Subnet.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Subnet *SubnetCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Subnet *SubnetSession) ADMINROLE() ([32]byte, error) {
	return _Subnet.Contract.ADMINROLE(&_Subnet.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Subnet *SubnetCallerSession) ADMINROLE() ([32]byte, error) {
	return _Subnet.Contract.ADMINROLE(&_Subnet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Subnet *SubnetCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Subnet *SubnetSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Subnet.Contract.DEFAULTADMINROLE(&_Subnet.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Subnet *SubnetCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Subnet.Contract.DEFAULTADMINROLE(&_Subnet.CallOpts)
}

// ActiveParticipantCounts is a free data retrieval call binding the contract method 0x2c5c1222.
//
// Solidity: function activeParticipantCounts(uint8 ) view returns(uint256)
func (_Subnet *SubnetCaller) ActiveParticipantCounts(opts *bind.CallOpts, arg0 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "activeParticipantCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveParticipantCounts is a free data retrieval call binding the contract method 0x2c5c1222.
//
// Solidity: function activeParticipantCounts(uint8 ) view returns(uint256)
func (_Subnet *SubnetSession) ActiveParticipantCounts(arg0 uint8) (*big.Int, error) {
	return _Subnet.Contract.ActiveParticipantCounts(&_Subnet.CallOpts, arg0)
}

// ActiveParticipantCounts is a free data retrieval call binding the contract method 0x2c5c1222.
//
// Solidity: function activeParticipantCounts(uint8 ) view returns(uint256)
func (_Subnet *SubnetCallerSession) ActiveParticipantCounts(arg0 uint8) (*big.Int, error) {
	return _Subnet.Contract.ActiveParticipantCounts(&_Subnet.CallOpts, arg0)
}

// GetActiveParticipantCount is a free data retrieval call binding the contract method 0x3d1dba56.
//
// Solidity: function getActiveParticipantCount(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetCaller) GetActiveParticipantCount(opts *bind.CallOpts, participant_type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getActiveParticipantCount", participant_type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveParticipantCount is a free data retrieval call binding the contract method 0x3d1dba56.
//
// Solidity: function getActiveParticipantCount(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetSession) GetActiveParticipantCount(participant_type uint8) (*big.Int, error) {
	return _Subnet.Contract.GetActiveParticipantCount(&_Subnet.CallOpts, participant_type)
}

// GetActiveParticipantCount is a free data retrieval call binding the contract method 0x3d1dba56.
//
// Solidity: function getActiveParticipantCount(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetCallerSession) GetActiveParticipantCount(participant_type uint8) (*big.Int, error) {
	return _Subnet.Contract.GetActiveParticipantCount(&_Subnet.CallOpts, participant_type)
}

// GetAgentIdentityRegistry is a free data retrieval call binding the contract method 0x03b9f732.
//
// Solidity: function getAgentIdentityRegistry() view returns(address)
func (_Subnet *SubnetCaller) GetAgentIdentityRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getAgentIdentityRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAgentIdentityRegistry is a free data retrieval call binding the contract method 0x03b9f732.
//
// Solidity: function getAgentIdentityRegistry() view returns(address)
func (_Subnet *SubnetSession) GetAgentIdentityRegistry() (common.Address, error) {
	return _Subnet.Contract.GetAgentIdentityRegistry(&_Subnet.CallOpts)
}

// GetAgentIdentityRegistry is a free data retrieval call binding the contract method 0x03b9f732.
//
// Solidity: function getAgentIdentityRegistry() view returns(address)
func (_Subnet *SubnetCallerSession) GetAgentIdentityRegistry() (common.Address, error) {
	return _Subnet.Contract.GetAgentIdentityRegistry(&_Subnet.CallOpts)
}

// GetBidFrequencyLimit is a free data retrieval call binding the contract method 0x4aee0b92.
//
// Solidity: function getBidFrequencyLimit() view returns(uint256)
func (_Subnet *SubnetCaller) GetBidFrequencyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getBidFrequencyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidFrequencyLimit is a free data retrieval call binding the contract method 0x4aee0b92.
//
// Solidity: function getBidFrequencyLimit() view returns(uint256)
func (_Subnet *SubnetSession) GetBidFrequencyLimit() (*big.Int, error) {
	return _Subnet.Contract.GetBidFrequencyLimit(&_Subnet.CallOpts)
}

// GetBidFrequencyLimit is a free data retrieval call binding the contract method 0x4aee0b92.
//
// Solidity: function getBidFrequencyLimit() view returns(uint256)
func (_Subnet *SubnetCallerSession) GetBidFrequencyLimit() (*big.Int, error) {
	return _Subnet.Contract.GetBidFrequencyLimit(&_Subnet.CallOpts)
}

// GetCheckpointPolicy is a free data retrieval call binding the contract method 0xd2b4353f.
//
// Solidity: function getCheckpointPolicy() view returns((uint128,uint64,uint64))
func (_Subnet *SubnetCaller) GetCheckpointPolicy(opts *bind.CallOpts) (DataStructuresCheckpointPolicy, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getCheckpointPolicy")

	if err != nil {
		return *new(DataStructuresCheckpointPolicy), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresCheckpointPolicy)).(*DataStructuresCheckpointPolicy)

	return out0, err

}

// GetCheckpointPolicy is a free data retrieval call binding the contract method 0xd2b4353f.
//
// Solidity: function getCheckpointPolicy() view returns((uint128,uint64,uint64))
func (_Subnet *SubnetSession) GetCheckpointPolicy() (DataStructuresCheckpointPolicy, error) {
	return _Subnet.Contract.GetCheckpointPolicy(&_Subnet.CallOpts)
}

// GetCheckpointPolicy is a free data retrieval call binding the contract method 0xd2b4353f.
//
// Solidity: function getCheckpointPolicy() view returns((uint128,uint64,uint64))
func (_Subnet *SubnetCallerSession) GetCheckpointPolicy() (DataStructuresCheckpointPolicy, error) {
	return _Subnet.Contract.GetCheckpointPolicy(&_Subnet.CallOpts)
}

// GetMinimumStakeByType is a free data retrieval call binding the contract method 0x9b0a82f6.
//
// Solidity: function getMinimumStakeByType(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetCaller) GetMinimumStakeByType(opts *bind.CallOpts, participant_type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getMinimumStakeByType", participant_type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinimumStakeByType is a free data retrieval call binding the contract method 0x9b0a82f6.
//
// Solidity: function getMinimumStakeByType(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetSession) GetMinimumStakeByType(participant_type uint8) (*big.Int, error) {
	return _Subnet.Contract.GetMinimumStakeByType(&_Subnet.CallOpts, participant_type)
}

// GetMinimumStakeByType is a free data retrieval call binding the contract method 0x9b0a82f6.
//
// Solidity: function getMinimumStakeByType(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetCallerSession) GetMinimumStakeByType(participant_type uint8) (*big.Int, error) {
	return _Subnet.Contract.GetMinimumStakeByType(&_Subnet.CallOpts, participant_type)
}

// GetParticipantAgentIds is a free data retrieval call binding the contract method 0x4064422e.
//
// Solidity: function getParticipantAgentIds(address participant_addr) view returns(uint256[])
func (_Subnet *SubnetCaller) GetParticipantAgentIds(opts *bind.CallOpts, participant_addr common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantAgentIds", participant_addr)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetParticipantAgentIds is a free data retrieval call binding the contract method 0x4064422e.
//
// Solidity: function getParticipantAgentIds(address participant_addr) view returns(uint256[])
func (_Subnet *SubnetSession) GetParticipantAgentIds(participant_addr common.Address) ([]*big.Int, error) {
	return _Subnet.Contract.GetParticipantAgentIds(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantAgentIds is a free data retrieval call binding the contract method 0x4064422e.
//
// Solidity: function getParticipantAgentIds(address participant_addr) view returns(uint256[])
func (_Subnet *SubnetCallerSession) GetParticipantAgentIds(participant_addr common.Address) ([]*big.Int, error) {
	return _Subnet.Contract.GetParticipantAgentIds(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantByAgentId is a free data retrieval call binding the contract method 0x78dc8e62.
//
// Solidity: function getParticipantByAgentId(uint256 agentId) view returns(address)
func (_Subnet *SubnetCaller) GetParticipantByAgentId(opts *bind.CallOpts, agentId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantByAgentId", agentId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetParticipantByAgentId is a free data retrieval call binding the contract method 0x78dc8e62.
//
// Solidity: function getParticipantByAgentId(uint256 agentId) view returns(address)
func (_Subnet *SubnetSession) GetParticipantByAgentId(agentId *big.Int) (common.Address, error) {
	return _Subnet.Contract.GetParticipantByAgentId(&_Subnet.CallOpts, agentId)
}

// GetParticipantByAgentId is a free data retrieval call binding the contract method 0x78dc8e62.
//
// Solidity: function getParticipantByAgentId(uint256 agentId) view returns(address)
func (_Subnet *SubnetCallerSession) GetParticipantByAgentId(agentId *big.Int) (common.Address, error) {
	return _Subnet.Contract.GetParticipantByAgentId(&_Subnet.CallOpts, agentId)
}

// GetParticipantCount is a free data retrieval call binding the contract method 0xc0a63671.
//
// Solidity: function getParticipantCount(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetCaller) GetParticipantCount(opts *bind.CallOpts, participant_type uint8) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantCount", participant_type)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParticipantCount is a free data retrieval call binding the contract method 0xc0a63671.
//
// Solidity: function getParticipantCount(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetSession) GetParticipantCount(participant_type uint8) (*big.Int, error) {
	return _Subnet.Contract.GetParticipantCount(&_Subnet.CallOpts, participant_type)
}

// GetParticipantCount is a free data retrieval call binding the contract method 0xc0a63671.
//
// Solidity: function getParticipantCount(uint8 participant_type) view returns(uint256)
func (_Subnet *SubnetCallerSession) GetParticipantCount(participant_type uint8) (*big.Int, error) {
	return _Subnet.Contract.GetParticipantCount(&_Subnet.CallOpts, participant_type)
}

// GetParticipantDomain is a free data retrieval call binding the contract method 0x96c586c2.
//
// Solidity: function getParticipantDomain(address participant_addr) view returns(string)
func (_Subnet *SubnetCaller) GetParticipantDomain(opts *bind.CallOpts, participant_addr common.Address) (string, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantDomain", participant_addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetParticipantDomain is a free data retrieval call binding the contract method 0x96c586c2.
//
// Solidity: function getParticipantDomain(address participant_addr) view returns(string)
func (_Subnet *SubnetSession) GetParticipantDomain(participant_addr common.Address) (string, error) {
	return _Subnet.Contract.GetParticipantDomain(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantDomain is a free data retrieval call binding the contract method 0x96c586c2.
//
// Solidity: function getParticipantDomain(address participant_addr) view returns(string)
func (_Subnet *SubnetCallerSession) GetParticipantDomain(participant_addr common.Address) (string, error) {
	return _Subnet.Contract.GetParticipantDomain(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantEndpoint is a free data retrieval call binding the contract method 0xce1245c4.
//
// Solidity: function getParticipantEndpoint(address participant_addr) view returns(string)
func (_Subnet *SubnetCaller) GetParticipantEndpoint(opts *bind.CallOpts, participant_addr common.Address) (string, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantEndpoint", participant_addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetParticipantEndpoint is a free data retrieval call binding the contract method 0xce1245c4.
//
// Solidity: function getParticipantEndpoint(address participant_addr) view returns(string)
func (_Subnet *SubnetSession) GetParticipantEndpoint(participant_addr common.Address) (string, error) {
	return _Subnet.Contract.GetParticipantEndpoint(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantEndpoint is a free data retrieval call binding the contract method 0xce1245c4.
//
// Solidity: function getParticipantEndpoint(address participant_addr) view returns(string)
func (_Subnet *SubnetCallerSession) GetParticipantEndpoint(participant_addr common.Address) (string, error) {
	return _Subnet.Contract.GetParticipantEndpoint(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantInfo is a free data retrieval call binding the contract method 0x735d23e8.
//
// Solidity: function getParticipantInfo(address participant_addr, uint8 participant_type) view returns((uint8,uint256,uint64))
func (_Subnet *SubnetCaller) GetParticipantInfo(opts *bind.CallOpts, participant_addr common.Address, participant_type uint8) (DataStructuresParticipantInfo, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantInfo", participant_addr, participant_type)

	if err != nil {
		return *new(DataStructuresParticipantInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresParticipantInfo)).(*DataStructuresParticipantInfo)

	return out0, err

}

// GetParticipantInfo is a free data retrieval call binding the contract method 0x735d23e8.
//
// Solidity: function getParticipantInfo(address participant_addr, uint8 participant_type) view returns((uint8,uint256,uint64))
func (_Subnet *SubnetSession) GetParticipantInfo(participant_addr common.Address, participant_type uint8) (DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.GetParticipantInfo(&_Subnet.CallOpts, participant_addr, participant_type)
}

// GetParticipantInfo is a free data retrieval call binding the contract method 0x735d23e8.
//
// Solidity: function getParticipantInfo(address participant_addr, uint8 participant_type) view returns((uint8,uint256,uint64))
func (_Subnet *SubnetCallerSession) GetParticipantInfo(participant_addr common.Address, participant_type uint8) (DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.GetParticipantInfo(&_Subnet.CallOpts, participant_addr, participant_type)
}

// GetParticipantMetadataUri is a free data retrieval call binding the contract method 0x75c6bbe2.
//
// Solidity: function getParticipantMetadataUri(address participant_addr) view returns(string)
func (_Subnet *SubnetCaller) GetParticipantMetadataUri(opts *bind.CallOpts, participant_addr common.Address) (string, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantMetadataUri", participant_addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetParticipantMetadataUri is a free data retrieval call binding the contract method 0x75c6bbe2.
//
// Solidity: function getParticipantMetadataUri(address participant_addr) view returns(string)
func (_Subnet *SubnetSession) GetParticipantMetadataUri(participant_addr common.Address) (string, error) {
	return _Subnet.Contract.GetParticipantMetadataUri(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantMetadataUri is a free data retrieval call binding the contract method 0x75c6bbe2.
//
// Solidity: function getParticipantMetadataUri(address participant_addr) view returns(string)
func (_Subnet *SubnetCallerSession) GetParticipantMetadataUri(participant_addr common.Address) (string, error) {
	return _Subnet.Contract.GetParticipantMetadataUri(&_Subnet.CallOpts, participant_addr)
}

// GetParticipantStakeInfo is a free data retrieval call binding the contract method 0xfa736600.
//
// Solidity: function getParticipantStakeInfo(address participant_addr, uint8 participant_type) view returns((uint256,uint256,uint256))
func (_Subnet *SubnetCaller) GetParticipantStakeInfo(opts *bind.CallOpts, participant_addr common.Address, participant_type uint8) (DataStructuresStakeInfo, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getParticipantStakeInfo", participant_addr, participant_type)

	if err != nil {
		return *new(DataStructuresStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresStakeInfo)).(*DataStructuresStakeInfo)

	return out0, err

}

// GetParticipantStakeInfo is a free data retrieval call binding the contract method 0xfa736600.
//
// Solidity: function getParticipantStakeInfo(address participant_addr, uint8 participant_type) view returns((uint256,uint256,uint256))
func (_Subnet *SubnetSession) GetParticipantStakeInfo(participant_addr common.Address, participant_type uint8) (DataStructuresStakeInfo, error) {
	return _Subnet.Contract.GetParticipantStakeInfo(&_Subnet.CallOpts, participant_addr, participant_type)
}

// GetParticipantStakeInfo is a free data retrieval call binding the contract method 0xfa736600.
//
// Solidity: function getParticipantStakeInfo(address participant_addr, uint8 participant_type) view returns((uint256,uint256,uint256))
func (_Subnet *SubnetCallerSession) GetParticipantStakeInfo(participant_addr common.Address, participant_type uint8) (DataStructuresStakeInfo, error) {
	return _Subnet.Contract.GetParticipantStakeInfo(&_Subnet.CallOpts, participant_addr, participant_type)
}

// GetRequireKYC is a free data retrieval call binding the contract method 0x73a3e871.
//
// Solidity: function getRequireKYC() view returns(bool)
func (_Subnet *SubnetCaller) GetRequireKYC(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getRequireKYC")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetRequireKYC is a free data retrieval call binding the contract method 0x73a3e871.
//
// Solidity: function getRequireKYC() view returns(bool)
func (_Subnet *SubnetSession) GetRequireKYC() (bool, error) {
	return _Subnet.Contract.GetRequireKYC(&_Subnet.CallOpts)
}

// GetRequireKYC is a free data retrieval call binding the contract method 0x73a3e871.
//
// Solidity: function getRequireKYC() view returns(bool)
func (_Subnet *SubnetCallerSession) GetRequireKYC() (bool, error) {
	return _Subnet.Contract.GetRequireKYC(&_Subnet.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Subnet *SubnetCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Subnet *SubnetSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Subnet.Contract.GetRoleAdmin(&_Subnet.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Subnet *SubnetCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Subnet.Contract.GetRoleAdmin(&_Subnet.CallOpts, role)
}

// GetSignatureThreshold is a free data retrieval call binding the contract method 0x2cae8f81.
//
// Solidity: function getSignatureThreshold() view returns((uint32,uint32))
func (_Subnet *SubnetCaller) GetSignatureThreshold(opts *bind.CallOpts) (DataStructuresSignatureThreshold, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getSignatureThreshold")

	if err != nil {
		return *new(DataStructuresSignatureThreshold), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresSignatureThreshold)).(*DataStructuresSignatureThreshold)

	return out0, err

}

// GetSignatureThreshold is a free data retrieval call binding the contract method 0x2cae8f81.
//
// Solidity: function getSignatureThreshold() view returns((uint32,uint32))
func (_Subnet *SubnetSession) GetSignatureThreshold() (DataStructuresSignatureThreshold, error) {
	return _Subnet.Contract.GetSignatureThreshold(&_Subnet.CallOpts)
}

// GetSignatureThreshold is a free data retrieval call binding the contract method 0x2cae8f81.
//
// Solidity: function getSignatureThreshold() view returns((uint32,uint32))
func (_Subnet *SubnetCallerSession) GetSignatureThreshold() (DataStructuresSignatureThreshold, error) {
	return _Subnet.Contract.GetSignatureThreshold(&_Subnet.CallOpts)
}

// GetSlashingRates is a free data retrieval call binding the contract method 0x509e8745.
//
// Solidity: function getSlashingRates() view returns(uint256[])
func (_Subnet *SubnetCaller) GetSlashingRates(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getSlashingRates")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSlashingRates is a free data retrieval call binding the contract method 0x509e8745.
//
// Solidity: function getSlashingRates() view returns(uint256[])
func (_Subnet *SubnetSession) GetSlashingRates() ([]*big.Int, error) {
	return _Subnet.Contract.GetSlashingRates(&_Subnet.CallOpts)
}

// GetSlashingRates is a free data retrieval call binding the contract method 0x509e8745.
//
// Solidity: function getSlashingRates() view returns(uint256[])
func (_Subnet *SubnetCallerSession) GetSlashingRates() ([]*big.Int, error) {
	return _Subnet.Contract.GetSlashingRates(&_Subnet.CallOpts)
}

// GetStakeGovernanceConfig is a free data retrieval call binding the contract method 0xce89f1de.
//
// Solidity: function getStakeGovernanceConfig() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_Subnet *SubnetCaller) GetStakeGovernanceConfig(opts *bind.CallOpts) (DataStructuresStakeGovernanceConfig, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getStakeGovernanceConfig")

	if err != nil {
		return *new(DataStructuresStakeGovernanceConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresStakeGovernanceConfig)).(*DataStructuresStakeGovernanceConfig)

	return out0, err

}

// GetStakeGovernanceConfig is a free data retrieval call binding the contract method 0xce89f1de.
//
// Solidity: function getStakeGovernanceConfig() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_Subnet *SubnetSession) GetStakeGovernanceConfig() (DataStructuresStakeGovernanceConfig, error) {
	return _Subnet.Contract.GetStakeGovernanceConfig(&_Subnet.CallOpts)
}

// GetStakeGovernanceConfig is a free data retrieval call binding the contract method 0xce89f1de.
//
// Solidity: function getStakeGovernanceConfig() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]))
func (_Subnet *SubnetCallerSession) GetStakeGovernanceConfig() (DataStructuresStakeGovernanceConfig, error) {
	return _Subnet.Contract.GetStakeGovernanceConfig(&_Subnet.CallOpts)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Subnet *SubnetCaller) GetStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Subnet *SubnetSession) GetStakingToken() (common.Address, error) {
	return _Subnet.Contract.GetStakingToken(&_Subnet.CallOpts)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_Subnet *SubnetCallerSession) GetStakingToken() (common.Address, error) {
	return _Subnet.Contract.GetStakingToken(&_Subnet.CallOpts)
}

// GetSubnetInfo is a free data retrieval call binding the contract method 0xd15442ad.
//
// Solidity: function getSubnetInfo() view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256))
func (_Subnet *SubnetCaller) GetSubnetInfo(opts *bind.CallOpts) (DataStructuresSubnetInfo, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getSubnetInfo")

	if err != nil {
		return *new(DataStructuresSubnetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresSubnetInfo)).(*DataStructuresSubnetInfo)

	return out0, err

}

// GetSubnetInfo is a free data retrieval call binding the contract method 0xd15442ad.
//
// Solidity: function getSubnetInfo() view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256))
func (_Subnet *SubnetSession) GetSubnetInfo() (DataStructuresSubnetInfo, error) {
	return _Subnet.Contract.GetSubnetInfo(&_Subnet.CallOpts)
}

// GetSubnetInfo is a free data retrieval call binding the contract method 0xd15442ad.
//
// Solidity: function getSubnetInfo() view returns((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256))
func (_Subnet *SubnetCallerSession) GetSubnetInfo() (DataStructuresSubnetInfo, error) {
	return _Subnet.Contract.GetSubnetInfo(&_Subnet.CallOpts)
}

// GetTotalParticipantCount is a free data retrieval call binding the contract method 0x40b33161.
//
// Solidity: function getTotalParticipantCount() view returns(uint256)
func (_Subnet *SubnetCaller) GetTotalParticipantCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getTotalParticipantCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalParticipantCount is a free data retrieval call binding the contract method 0x40b33161.
//
// Solidity: function getTotalParticipantCount() view returns(uint256)
func (_Subnet *SubnetSession) GetTotalParticipantCount() (*big.Int, error) {
	return _Subnet.Contract.GetTotalParticipantCount(&_Subnet.CallOpts)
}

// GetTotalParticipantCount is a free data retrieval call binding the contract method 0x40b33161.
//
// Solidity: function getTotalParticipantCount() view returns(uint256)
func (_Subnet *SubnetCallerSession) GetTotalParticipantCount() (*big.Int, error) {
	return _Subnet.Contract.GetTotalParticipantCount(&_Subnet.CallOpts)
}

// GetUnstakeLockPeriod is a free data retrieval call binding the contract method 0x4b287566.
//
// Solidity: function getUnstakeLockPeriod() view returns(uint256)
func (_Subnet *SubnetCaller) GetUnstakeLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "getUnstakeLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnstakeLockPeriod is a free data retrieval call binding the contract method 0x4b287566.
//
// Solidity: function getUnstakeLockPeriod() view returns(uint256)
func (_Subnet *SubnetSession) GetUnstakeLockPeriod() (*big.Int, error) {
	return _Subnet.Contract.GetUnstakeLockPeriod(&_Subnet.CallOpts)
}

// GetUnstakeLockPeriod is a free data retrieval call binding the contract method 0x4b287566.
//
// Solidity: function getUnstakeLockPeriod() view returns(uint256)
func (_Subnet *SubnetCallerSession) GetUnstakeLockPeriod() (*big.Int, error) {
	return _Subnet.Contract.GetUnstakeLockPeriod(&_Subnet.CallOpts)
}

// HasAvailableSlots is a free data retrieval call binding the contract method 0xbd352a2b.
//
// Solidity: function hasAvailableSlots(uint8 participant_type) view returns(bool)
func (_Subnet *SubnetCaller) HasAvailableSlots(opts *bind.CallOpts, participant_type uint8) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "hasAvailableSlots", participant_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAvailableSlots is a free data retrieval call binding the contract method 0xbd352a2b.
//
// Solidity: function hasAvailableSlots(uint8 participant_type) view returns(bool)
func (_Subnet *SubnetSession) HasAvailableSlots(participant_type uint8) (bool, error) {
	return _Subnet.Contract.HasAvailableSlots(&_Subnet.CallOpts, participant_type)
}

// HasAvailableSlots is a free data retrieval call binding the contract method 0xbd352a2b.
//
// Solidity: function hasAvailableSlots(uint8 participant_type) view returns(bool)
func (_Subnet *SubnetCallerSession) HasAvailableSlots(participant_type uint8) (bool, error) {
	return _Subnet.Contract.HasAvailableSlots(&_Subnet.CallOpts, participant_type)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Subnet *SubnetCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Subnet *SubnetSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Subnet.Contract.HasRole(&_Subnet.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Subnet *SubnetCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Subnet.Contract.HasRole(&_Subnet.CallOpts, role, account)
}

// IsActiveParticipant is a free data retrieval call binding the contract method 0xac921890.
//
// Solidity: function isActiveParticipant(address participant_addr, uint8 participant_type) view returns(bool)
func (_Subnet *SubnetCaller) IsActiveParticipant(opts *bind.CallOpts, participant_addr common.Address, participant_type uint8) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "isActiveParticipant", participant_addr, participant_type)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveParticipant is a free data retrieval call binding the contract method 0xac921890.
//
// Solidity: function isActiveParticipant(address participant_addr, uint8 participant_type) view returns(bool)
func (_Subnet *SubnetSession) IsActiveParticipant(participant_addr common.Address, participant_type uint8) (bool, error) {
	return _Subnet.Contract.IsActiveParticipant(&_Subnet.CallOpts, participant_addr, participant_type)
}

// IsActiveParticipant is a free data retrieval call binding the contract method 0xac921890.
//
// Solidity: function isActiveParticipant(address participant_addr, uint8 participant_type) view returns(bool)
func (_Subnet *SubnetCallerSession) IsActiveParticipant(participant_addr common.Address, participant_type uint8) (bool, error) {
	return _Subnet.Contract.IsActiveParticipant(&_Subnet.CallOpts, participant_addr, participant_type)
}

// ListActiveParticipants is a free data retrieval call binding the contract method 0x5caecd9c.
//
// Solidity: function listActiveParticipants(uint8 participant_type) view returns((uint8,uint256,uint64)[])
func (_Subnet *SubnetCaller) ListActiveParticipants(opts *bind.CallOpts, participant_type uint8) ([]DataStructuresParticipantInfo, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "listActiveParticipants", participant_type)

	if err != nil {
		return *new([]DataStructuresParticipantInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DataStructuresParticipantInfo)).(*[]DataStructuresParticipantInfo)

	return out0, err

}

// ListActiveParticipants is a free data retrieval call binding the contract method 0x5caecd9c.
//
// Solidity: function listActiveParticipants(uint8 participant_type) view returns((uint8,uint256,uint64)[])
func (_Subnet *SubnetSession) ListActiveParticipants(participant_type uint8) ([]DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.ListActiveParticipants(&_Subnet.CallOpts, participant_type)
}

// ListActiveParticipants is a free data retrieval call binding the contract method 0x5caecd9c.
//
// Solidity: function listActiveParticipants(uint8 participant_type) view returns((uint8,uint256,uint64)[])
func (_Subnet *SubnetCallerSession) ListActiveParticipants(participant_type uint8) ([]DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.ListActiveParticipants(&_Subnet.CallOpts, participant_type)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Subnet *SubnetCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Subnet *SubnetSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Subnet.Contract.OnERC721Received(&_Subnet.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Subnet *SubnetCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Subnet.Contract.OnERC721Received(&_Subnet.CallOpts, arg0, arg1, arg2, arg3)
}

// Participants is a free data retrieval call binding the contract method 0xca8cf966.
//
// Solidity: function participants(address , uint8 ) view returns(uint8 status, uint256 agent_id, uint64 registered_at)
func (_Subnet *SubnetCaller) Participants(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (struct {
	Status       uint8
	AgentId      *big.Int
	RegisteredAt uint64
}, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "participants", arg0, arg1)

	outstruct := new(struct {
		Status       uint8
		AgentId      *big.Int
		RegisteredAt uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.AgentId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RegisteredAt = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Participants is a free data retrieval call binding the contract method 0xca8cf966.
//
// Solidity: function participants(address , uint8 ) view returns(uint8 status, uint256 agent_id, uint64 registered_at)
func (_Subnet *SubnetSession) Participants(arg0 common.Address, arg1 uint8) (struct {
	Status       uint8
	AgentId      *big.Int
	RegisteredAt uint64
}, error) {
	return _Subnet.Contract.Participants(&_Subnet.CallOpts, arg0, arg1)
}

// Participants is a free data retrieval call binding the contract method 0xca8cf966.
//
// Solidity: function participants(address , uint8 ) view returns(uint8 status, uint256 agent_id, uint64 registered_at)
func (_Subnet *SubnetCallerSession) Participants(arg0 common.Address, arg1 uint8) (struct {
	Status       uint8
	AgentId      *big.Int
	RegisteredAt uint64
}, error) {
	return _Subnet.Contract.Participants(&_Subnet.CallOpts, arg0, arg1)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Subnet *SubnetCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Subnet *SubnetSession) Paused() (bool, error) {
	return _Subnet.Contract.Paused(&_Subnet.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Subnet *SubnetCallerSession) Paused() (bool, error) {
	return _Subnet.Contract.Paused(&_Subnet.CallOpts)
}

// RequireActiveParticipants is a free data retrieval call binding the contract method 0xd12630c1.
//
// Solidity: function requireActiveParticipants(address[] participant_addrs, uint8[] participant_types) view returns()
func (_Subnet *SubnetCaller) RequireActiveParticipants(opts *bind.CallOpts, participant_addrs []common.Address, participant_types []uint8) error {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "requireActiveParticipants", participant_addrs, participant_types)

	if err != nil {
		return err
	}

	return err

}

// RequireActiveParticipants is a free data retrieval call binding the contract method 0xd12630c1.
//
// Solidity: function requireActiveParticipants(address[] participant_addrs, uint8[] participant_types) view returns()
func (_Subnet *SubnetSession) RequireActiveParticipants(participant_addrs []common.Address, participant_types []uint8) error {
	return _Subnet.Contract.RequireActiveParticipants(&_Subnet.CallOpts, participant_addrs, participant_types)
}

// RequireActiveParticipants is a free data retrieval call binding the contract method 0xd12630c1.
//
// Solidity: function requireActiveParticipants(address[] participant_addrs, uint8[] participant_types) view returns()
func (_Subnet *SubnetCallerSession) RequireActiveParticipants(participant_addrs []common.Address, participant_types []uint8) error {
	return _Subnet.Contract.RequireActiveParticipants(&_Subnet.CallOpts, participant_addrs, participant_types)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subnet *SubnetCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subnet *SubnetSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Subnet.Contract.SupportsInterface(&_Subnet.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Subnet *SubnetCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Subnet.Contract.SupportsInterface(&_Subnet.CallOpts, interfaceId)
}

// ApproveParticipant is a paid mutator transaction binding the contract method 0x3451479e.
//
// Solidity: function approveParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetTransactor) ApproveParticipant(opts *bind.TransactOpts, participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "approveParticipant", participant_addr, participant_type)
}

// ApproveParticipant is a paid mutator transaction binding the contract method 0x3451479e.
//
// Solidity: function approveParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetSession) ApproveParticipant(participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.ApproveParticipant(&_Subnet.TransactOpts, participant_addr, participant_type)
}

// ApproveParticipant is a paid mutator transaction binding the contract method 0x3451479e.
//
// Solidity: function approveParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetTransactorSession) ApproveParticipant(participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.ApproveParticipant(&_Subnet.TransactOpts, participant_addr, participant_type)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_Subnet *SubnetTransactor) Deprecate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "deprecate")
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_Subnet *SubnetSession) Deprecate() (*types.Transaction, error) {
	return _Subnet.Contract.Deprecate(&_Subnet.TransactOpts)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_Subnet *SubnetTransactorSession) Deprecate() (*types.Transaction, error) {
	return _Subnet.Contract.Deprecate(&_Subnet.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Subnet *SubnetTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Subnet *SubnetSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.GrantRole(&_Subnet.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Subnet *SubnetTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.GrantRole(&_Subnet.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8729ff46.
//
// Solidity: function initialize((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) _subnetInfo, address _factory, address stakingManager_, address identityRegistry_) returns()
func (_Subnet *SubnetTransactor) Initialize(opts *bind.TransactOpts, _subnetInfo DataStructuresSubnetInfo, _factory common.Address, stakingManager_ common.Address, identityRegistry_ common.Address) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "initialize", _subnetInfo, _factory, stakingManager_, identityRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8729ff46.
//
// Solidity: function initialize((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) _subnetInfo, address _factory, address stakingManager_, address identityRegistry_) returns()
func (_Subnet *SubnetSession) Initialize(_subnetInfo DataStructuresSubnetInfo, _factory common.Address, stakingManager_ common.Address, identityRegistry_ common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.Initialize(&_Subnet.TransactOpts, _subnetInfo, _factory, stakingManager_, identityRegistry_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8729ff46.
//
// Solidity: function initialize((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) _subnetInfo, address _factory, address stakingManager_, address identityRegistry_) returns()
func (_Subnet *SubnetTransactorSession) Initialize(_subnetInfo DataStructuresSubnetInfo, _factory common.Address, stakingManager_ common.Address, identityRegistry_ common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.Initialize(&_Subnet.TransactOpts, _subnetInfo, _factory, stakingManager_, identityRegistry_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Subnet *SubnetTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Subnet *SubnetSession) Pause() (*types.Transaction, error) {
	return _Subnet.Contract.Pause(&_Subnet.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Subnet *SubnetTransactorSession) Pause() (*types.Transaction, error) {
	return _Subnet.Contract.Pause(&_Subnet.TransactOpts)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x6636fd2e.
//
// Solidity: function registerParticipant(uint8 participant_type, uint256 agentId, string domain, string endpoint, string metadata_uri) payable returns(bool)
func (_Subnet *SubnetTransactor) RegisterParticipant(opts *bind.TransactOpts, participant_type uint8, agentId *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "registerParticipant", participant_type, agentId, domain, endpoint, metadata_uri)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x6636fd2e.
//
// Solidity: function registerParticipant(uint8 participant_type, uint256 agentId, string domain, string endpoint, string metadata_uri) payable returns(bool)
func (_Subnet *SubnetSession) RegisterParticipant(participant_type uint8, agentId *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipant(&_Subnet.TransactOpts, participant_type, agentId, domain, endpoint, metadata_uri)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x6636fd2e.
//
// Solidity: function registerParticipant(uint8 participant_type, uint256 agentId, string domain, string endpoint, string metadata_uri) payable returns(bool)
func (_Subnet *SubnetTransactorSession) RegisterParticipant(participant_type uint8, agentId *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipant(&_Subnet.TransactOpts, participant_type, agentId, domain, endpoint, metadata_uri)
}

// RegisterParticipantERC20 is a paid mutator transaction binding the contract method 0x765bf529.
//
// Solidity: function registerParticipantERC20(uint8 participant_type, uint256 agentId, uint256 amount, string domain, string endpoint, string metadata_uri) returns(bool)
func (_Subnet *SubnetTransactor) RegisterParticipantERC20(opts *bind.TransactOpts, participant_type uint8, agentId *big.Int, amount *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "registerParticipantERC20", participant_type, agentId, amount, domain, endpoint, metadata_uri)
}

// RegisterParticipantERC20 is a paid mutator transaction binding the contract method 0x765bf529.
//
// Solidity: function registerParticipantERC20(uint8 participant_type, uint256 agentId, uint256 amount, string domain, string endpoint, string metadata_uri) returns(bool)
func (_Subnet *SubnetSession) RegisterParticipantERC20(participant_type uint8, agentId *big.Int, amount *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipantERC20(&_Subnet.TransactOpts, participant_type, agentId, amount, domain, endpoint, metadata_uri)
}

// RegisterParticipantERC20 is a paid mutator transaction binding the contract method 0x765bf529.
//
// Solidity: function registerParticipantERC20(uint8 participant_type, uint256 agentId, uint256 amount, string domain, string endpoint, string metadata_uri) returns(bool)
func (_Subnet *SubnetTransactorSession) RegisterParticipantERC20(participant_type uint8, agentId *big.Int, amount *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipantERC20(&_Subnet.TransactOpts, participant_type, agentId, amount, domain, endpoint, metadata_uri)
}

// RejectParticipant is a paid mutator transaction binding the contract method 0x6ee88b16.
//
// Solidity: function rejectParticipant(address participant_addr, uint8 participant_type, string reason) returns()
func (_Subnet *SubnetTransactor) RejectParticipant(opts *bind.TransactOpts, participant_addr common.Address, participant_type uint8, reason string) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "rejectParticipant", participant_addr, participant_type, reason)
}

// RejectParticipant is a paid mutator transaction binding the contract method 0x6ee88b16.
//
// Solidity: function rejectParticipant(address participant_addr, uint8 participant_type, string reason) returns()
func (_Subnet *SubnetSession) RejectParticipant(participant_addr common.Address, participant_type uint8, reason string) (*types.Transaction, error) {
	return _Subnet.Contract.RejectParticipant(&_Subnet.TransactOpts, participant_addr, participant_type, reason)
}

// RejectParticipant is a paid mutator transaction binding the contract method 0x6ee88b16.
//
// Solidity: function rejectParticipant(address participant_addr, uint8 participant_type, string reason) returns()
func (_Subnet *SubnetTransactorSession) RejectParticipant(participant_addr common.Address, participant_type uint8, reason string) (*types.Transaction, error) {
	return _Subnet.Contract.RejectParticipant(&_Subnet.TransactOpts, participant_addr, participant_type, reason)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Subnet *SubnetTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Subnet *SubnetSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.RenounceRole(&_Subnet.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Subnet *SubnetTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.RenounceRole(&_Subnet.TransactOpts, role, callerConfirmation)
}

// ResumeParticipant is a paid mutator transaction binding the contract method 0xb8e90b67.
//
// Solidity: function resumeParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetTransactor) ResumeParticipant(opts *bind.TransactOpts, participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "resumeParticipant", participant_addr, participant_type)
}

// ResumeParticipant is a paid mutator transaction binding the contract method 0xb8e90b67.
//
// Solidity: function resumeParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetSession) ResumeParticipant(participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.ResumeParticipant(&_Subnet.TransactOpts, participant_addr, participant_type)
}

// ResumeParticipant is a paid mutator transaction binding the contract method 0xb8e90b67.
//
// Solidity: function resumeParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetTransactorSession) ResumeParticipant(participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.ResumeParticipant(&_Subnet.TransactOpts, participant_addr, participant_type)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Subnet *SubnetTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Subnet *SubnetSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.RevokeRole(&_Subnet.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Subnet *SubnetTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.RevokeRole(&_Subnet.TransactOpts, role, account)
}

// SetStakeConfig is a paid mutator transaction binding the contract method 0xe758b8ae.
//
// Solidity: function setStakeConfig((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) _stakeConfig) returns()
func (_Subnet *SubnetTransactor) SetStakeConfig(opts *bind.TransactOpts, _stakeConfig DataStructuresStakeGovernanceConfig) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "setStakeConfig", _stakeConfig)
}

// SetStakeConfig is a paid mutator transaction binding the contract method 0xe758b8ae.
//
// Solidity: function setStakeConfig((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) _stakeConfig) returns()
func (_Subnet *SubnetSession) SetStakeConfig(_stakeConfig DataStructuresStakeGovernanceConfig) (*types.Transaction, error) {
	return _Subnet.Contract.SetStakeConfig(&_Subnet.TransactOpts, _stakeConfig)
}

// SetStakeConfig is a paid mutator transaction binding the contract method 0xe758b8ae.
//
// Solidity: function setStakeConfig((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]) _stakeConfig) returns()
func (_Subnet *SubnetTransactorSession) SetStakeConfig(_stakeConfig DataStructuresStakeGovernanceConfig) (*types.Transaction, error) {
	return _Subnet.Contract.SetStakeConfig(&_Subnet.TransactOpts, _stakeConfig)
}

// SuspendParticipant is a paid mutator transaction binding the contract method 0x932b4a5a.
//
// Solidity: function suspendParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetTransactor) SuspendParticipant(opts *bind.TransactOpts, participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "suspendParticipant", participant_addr, participant_type)
}

// SuspendParticipant is a paid mutator transaction binding the contract method 0x932b4a5a.
//
// Solidity: function suspendParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetSession) SuspendParticipant(participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.SuspendParticipant(&_Subnet.TransactOpts, participant_addr, participant_type)
}

// SuspendParticipant is a paid mutator transaction binding the contract method 0x932b4a5a.
//
// Solidity: function suspendParticipant(address participant_addr, uint8 participant_type) returns()
func (_Subnet *SubnetTransactorSession) SuspendParticipant(participant_addr common.Address, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.SuspendParticipant(&_Subnet.TransactOpts, participant_addr, participant_type)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Subnet *SubnetTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Subnet *SubnetSession) Unpause() (*types.Transaction, error) {
	return _Subnet.Contract.Unpause(&_Subnet.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Subnet *SubnetTransactorSession) Unpause() (*types.Transaction, error) {
	return _Subnet.Contract.Unpause(&_Subnet.TransactOpts)
}

// UpdateParticipantActivity is a paid mutator transaction binding the contract method 0x1da65cc3.
//
// Solidity: function updateParticipantActivity(uint8 participant_type) returns()
func (_Subnet *SubnetTransactor) UpdateParticipantActivity(opts *bind.TransactOpts, participant_type uint8) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "updateParticipantActivity", participant_type)
}

// UpdateParticipantActivity is a paid mutator transaction binding the contract method 0x1da65cc3.
//
// Solidity: function updateParticipantActivity(uint8 participant_type) returns()
func (_Subnet *SubnetSession) UpdateParticipantActivity(participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.UpdateParticipantActivity(&_Subnet.TransactOpts, participant_type)
}

// UpdateParticipantActivity is a paid mutator transaction binding the contract method 0x1da65cc3.
//
// Solidity: function updateParticipantActivity(uint8 participant_type) returns()
func (_Subnet *SubnetTransactorSession) UpdateParticipantActivity(participant_type uint8) (*types.Transaction, error) {
	return _Subnet.Contract.UpdateParticipantActivity(&_Subnet.TransactOpts, participant_type)
}

// SubnetAgentIdentityRegistrySetIterator is returned from FilterAgentIdentityRegistrySet and is used to iterate over the raw logs and unpacked data for AgentIdentityRegistrySet events raised by the Subnet contract.
type SubnetAgentIdentityRegistrySetIterator struct {
	Event *SubnetAgentIdentityRegistrySet // Event containing the contract specifics and raw log

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
func (it *SubnetAgentIdentityRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetAgentIdentityRegistrySet)
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
		it.Event = new(SubnetAgentIdentityRegistrySet)
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
func (it *SubnetAgentIdentityRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetAgentIdentityRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetAgentIdentityRegistrySet represents a AgentIdentityRegistrySet event raised by the Subnet contract.
type SubnetAgentIdentityRegistrySet struct {
	Registry common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentIdentityRegistrySet is a free log retrieval operation binding the contract event 0x320f157b4b262e9289d2b57973872e6e058cb44ba465f5271d57287afdcc1adc.
//
// Solidity: event AgentIdentityRegistrySet(address indexed registry)
func (_Subnet *SubnetFilterer) FilterAgentIdentityRegistrySet(opts *bind.FilterOpts, registry []common.Address) (*SubnetAgentIdentityRegistrySetIterator, error) {

	var registryRule []interface{}
	for _, registryItem := range registry {
		registryRule = append(registryRule, registryItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "AgentIdentityRegistrySet", registryRule)
	if err != nil {
		return nil, err
	}
	return &SubnetAgentIdentityRegistrySetIterator{contract: _Subnet.contract, event: "AgentIdentityRegistrySet", logs: logs, sub: sub}, nil
}

// WatchAgentIdentityRegistrySet is a free log subscription operation binding the contract event 0x320f157b4b262e9289d2b57973872e6e058cb44ba465f5271d57287afdcc1adc.
//
// Solidity: event AgentIdentityRegistrySet(address indexed registry)
func (_Subnet *SubnetFilterer) WatchAgentIdentityRegistrySet(opts *bind.WatchOpts, sink chan<- *SubnetAgentIdentityRegistrySet, registry []common.Address) (event.Subscription, error) {

	var registryRule []interface{}
	for _, registryItem := range registry {
		registryRule = append(registryRule, registryItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "AgentIdentityRegistrySet", registryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetAgentIdentityRegistrySet)
				if err := _Subnet.contract.UnpackLog(event, "AgentIdentityRegistrySet", log); err != nil {
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

// ParseAgentIdentityRegistrySet is a log parse operation binding the contract event 0x320f157b4b262e9289d2b57973872e6e058cb44ba465f5271d57287afdcc1adc.
//
// Solidity: event AgentIdentityRegistrySet(address indexed registry)
func (_Subnet *SubnetFilterer) ParseAgentIdentityRegistrySet(log types.Log) (*SubnetAgentIdentityRegistrySet, error) {
	event := new(SubnetAgentIdentityRegistrySet)
	if err := _Subnet.contract.UnpackLog(event, "AgentIdentityRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetFeedbackSubmittedIterator is returned from FilterFeedbackSubmitted and is used to iterate over the raw logs and unpacked data for FeedbackSubmitted events raised by the Subnet contract.
type SubnetFeedbackSubmittedIterator struct {
	Event *SubnetFeedbackSubmitted // Event containing the contract specifics and raw log

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
func (it *SubnetFeedbackSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetFeedbackSubmitted)
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
		it.Event = new(SubnetFeedbackSubmitted)
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
func (it *SubnetFeedbackSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetFeedbackSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetFeedbackSubmitted represents a FeedbackSubmitted event raised by the Subnet contract.
type SubnetFeedbackSubmitted struct {
	FeedbackId   [32]byte
	FromAddr     common.Address
	ToAgent      common.Address
	FeedbackType uint8
	Rating       uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeedbackSubmitted is a free log retrieval operation binding the contract event 0xdaa80b0d543897a8b95ff35c0e62fce70e3af8fdb5ef622dbaae789ccea60754.
//
// Solidity: event FeedbackSubmitted(bytes32 indexed feedback_id, address indexed from_addr, address indexed to_agent, uint8 feedback_type, uint8 rating)
func (_Subnet *SubnetFilterer) FilterFeedbackSubmitted(opts *bind.FilterOpts, feedback_id [][32]byte, from_addr []common.Address, to_agent []common.Address) (*SubnetFeedbackSubmittedIterator, error) {

	var feedback_idRule []interface{}
	for _, feedback_idItem := range feedback_id {
		feedback_idRule = append(feedback_idRule, feedback_idItem)
	}
	var from_addrRule []interface{}
	for _, from_addrItem := range from_addr {
		from_addrRule = append(from_addrRule, from_addrItem)
	}
	var to_agentRule []interface{}
	for _, to_agentItem := range to_agent {
		to_agentRule = append(to_agentRule, to_agentItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "FeedbackSubmitted", feedback_idRule, from_addrRule, to_agentRule)
	if err != nil {
		return nil, err
	}
	return &SubnetFeedbackSubmittedIterator{contract: _Subnet.contract, event: "FeedbackSubmitted", logs: logs, sub: sub}, nil
}

// WatchFeedbackSubmitted is a free log subscription operation binding the contract event 0xdaa80b0d543897a8b95ff35c0e62fce70e3af8fdb5ef622dbaae789ccea60754.
//
// Solidity: event FeedbackSubmitted(bytes32 indexed feedback_id, address indexed from_addr, address indexed to_agent, uint8 feedback_type, uint8 rating)
func (_Subnet *SubnetFilterer) WatchFeedbackSubmitted(opts *bind.WatchOpts, sink chan<- *SubnetFeedbackSubmitted, feedback_id [][32]byte, from_addr []common.Address, to_agent []common.Address) (event.Subscription, error) {

	var feedback_idRule []interface{}
	for _, feedback_idItem := range feedback_id {
		feedback_idRule = append(feedback_idRule, feedback_idItem)
	}
	var from_addrRule []interface{}
	for _, from_addrItem := range from_addr {
		from_addrRule = append(from_addrRule, from_addrItem)
	}
	var to_agentRule []interface{}
	for _, to_agentItem := range to_agent {
		to_agentRule = append(to_agentRule, to_agentItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "FeedbackSubmitted", feedback_idRule, from_addrRule, to_agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetFeedbackSubmitted)
				if err := _Subnet.contract.UnpackLog(event, "FeedbackSubmitted", log); err != nil {
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

// ParseFeedbackSubmitted is a log parse operation binding the contract event 0xdaa80b0d543897a8b95ff35c0e62fce70e3af8fdb5ef622dbaae789ccea60754.
//
// Solidity: event FeedbackSubmitted(bytes32 indexed feedback_id, address indexed from_addr, address indexed to_agent, uint8 feedback_type, uint8 rating)
func (_Subnet *SubnetFilterer) ParseFeedbackSubmitted(log types.Log) (*SubnetFeedbackSubmitted, error) {
	event := new(SubnetFeedbackSubmitted)
	if err := _Subnet.contract.UnpackLog(event, "FeedbackSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Subnet contract.
type SubnetInitializedIterator struct {
	Event *SubnetInitialized // Event containing the contract specifics and raw log

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
func (it *SubnetInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetInitialized)
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
		it.Event = new(SubnetInitialized)
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
func (it *SubnetInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetInitialized represents a Initialized event raised by the Subnet contract.
type SubnetInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Subnet *SubnetFilterer) FilterInitialized(opts *bind.FilterOpts) (*SubnetInitializedIterator, error) {

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SubnetInitializedIterator{contract: _Subnet.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Subnet *SubnetFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SubnetInitialized) (event.Subscription, error) {

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetInitialized)
				if err := _Subnet.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Subnet *SubnetFilterer) ParseInitialized(log types.Log) (*SubnetInitialized, error) {
	event := new(SubnetInitialized)
	if err := _Subnet.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantActivityUpdatedIterator is returned from FilterParticipantActivityUpdated and is used to iterate over the raw logs and unpacked data for ParticipantActivityUpdated events raised by the Subnet contract.
type SubnetParticipantActivityUpdatedIterator struct {
	Event *SubnetParticipantActivityUpdated // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantActivityUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantActivityUpdated)
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
		it.Event = new(SubnetParticipantActivityUpdated)
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
func (it *SubnetParticipantActivityUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantActivityUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantActivityUpdated represents a ParticipantActivityUpdated event raised by the Subnet contract.
type SubnetParticipantActivityUpdated struct {
	Addr            common.Address
	ParticipantType uint8
	LastActive      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterParticipantActivityUpdated is a free log retrieval operation binding the contract event 0x6b8db42cee7260f00718e255640423e4d8dd1b15b9ce881aaee80e5a45e7d077.
//
// Solidity: event ParticipantActivityUpdated(address indexed addr, uint8 indexed participant_type, uint128 last_active)
func (_Subnet *SubnetFilterer) FilterParticipantActivityUpdated(opts *bind.FilterOpts, addr []common.Address, participant_type []uint8) (*SubnetParticipantActivityUpdatedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var participant_typeRule []interface{}
	for _, participant_typeItem := range participant_type {
		participant_typeRule = append(participant_typeRule, participant_typeItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantActivityUpdated", addrRule, participant_typeRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantActivityUpdatedIterator{contract: _Subnet.contract, event: "ParticipantActivityUpdated", logs: logs, sub: sub}, nil
}

// WatchParticipantActivityUpdated is a free log subscription operation binding the contract event 0x6b8db42cee7260f00718e255640423e4d8dd1b15b9ce881aaee80e5a45e7d077.
//
// Solidity: event ParticipantActivityUpdated(address indexed addr, uint8 indexed participant_type, uint128 last_active)
func (_Subnet *SubnetFilterer) WatchParticipantActivityUpdated(opts *bind.WatchOpts, sink chan<- *SubnetParticipantActivityUpdated, addr []common.Address, participant_type []uint8) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var participant_typeRule []interface{}
	for _, participant_typeItem := range participant_type {
		participant_typeRule = append(participant_typeRule, participant_typeItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantActivityUpdated", addrRule, participant_typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantActivityUpdated)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantActivityUpdated", log); err != nil {
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

// ParseParticipantActivityUpdated is a log parse operation binding the contract event 0x6b8db42cee7260f00718e255640423e4d8dd1b15b9ce881aaee80e5a45e7d077.
//
// Solidity: event ParticipantActivityUpdated(address indexed addr, uint8 indexed participant_type, uint128 last_active)
func (_Subnet *SubnetFilterer) ParseParticipantActivityUpdated(log types.Log) (*SubnetParticipantActivityUpdated, error) {
	event := new(SubnetParticipantActivityUpdated)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantActivityUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantApprovedIterator is returned from FilterParticipantApproved and is used to iterate over the raw logs and unpacked data for ParticipantApproved events raised by the Subnet contract.
type SubnetParticipantApprovedIterator struct {
	Event *SubnetParticipantApproved // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantApproved)
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
		it.Event = new(SubnetParticipantApproved)
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
func (it *SubnetParticipantApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantApproved represents a ParticipantApproved event raised by the Subnet contract.
type SubnetParticipantApproved struct {
	Participant     common.Address
	ParticipantType uint8
	Approver        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterParticipantApproved is a free log retrieval operation binding the contract event 0x824961050c68118de8b5a2c2d827932d4a9efc756d8252776df33ca28b0e7754.
//
// Solidity: event ParticipantApproved(address indexed participant, uint8 indexed participantType, address indexed approver)
func (_Subnet *SubnetFilterer) FilterParticipantApproved(opts *bind.FilterOpts, participant []common.Address, participantType []uint8, approver []common.Address) (*SubnetParticipantApprovedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var participantTypeRule []interface{}
	for _, participantTypeItem := range participantType {
		participantTypeRule = append(participantTypeRule, participantTypeItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantApproved", participantRule, participantTypeRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantApprovedIterator{contract: _Subnet.contract, event: "ParticipantApproved", logs: logs, sub: sub}, nil
}

// WatchParticipantApproved is a free log subscription operation binding the contract event 0x824961050c68118de8b5a2c2d827932d4a9efc756d8252776df33ca28b0e7754.
//
// Solidity: event ParticipantApproved(address indexed participant, uint8 indexed participantType, address indexed approver)
func (_Subnet *SubnetFilterer) WatchParticipantApproved(opts *bind.WatchOpts, sink chan<- *SubnetParticipantApproved, participant []common.Address, participantType []uint8, approver []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var participantTypeRule []interface{}
	for _, participantTypeItem := range participantType {
		participantTypeRule = append(participantTypeRule, participantTypeItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantApproved", participantRule, participantTypeRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantApproved)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantApproved", log); err != nil {
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

// ParseParticipantApproved is a log parse operation binding the contract event 0x824961050c68118de8b5a2c2d827932d4a9efc756d8252776df33ca28b0e7754.
//
// Solidity: event ParticipantApproved(address indexed participant, uint8 indexed participantType, address indexed approver)
func (_Subnet *SubnetFilterer) ParseParticipantApproved(log types.Log) (*SubnetParticipantApproved, error) {
	event := new(SubnetParticipantApproved)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantIdentityBoundIterator is returned from FilterParticipantIdentityBound and is used to iterate over the raw logs and unpacked data for ParticipantIdentityBound events raised by the Subnet contract.
type SubnetParticipantIdentityBoundIterator struct {
	Event *SubnetParticipantIdentityBound // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantIdentityBoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantIdentityBound)
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
		it.Event = new(SubnetParticipantIdentityBound)
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
func (it *SubnetParticipantIdentityBoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantIdentityBoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantIdentityBound represents a ParticipantIdentityBound event raised by the Subnet contract.
type SubnetParticipantIdentityBound struct {
	Participant common.Address
	AgentId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantIdentityBound is a free log retrieval operation binding the contract event 0xd54ee5b46503fad8a00f45ed58660ef98a860ed5ace41d9b5795f91b3d0dc6d5.
//
// Solidity: event ParticipantIdentityBound(address indexed participant, uint256 indexed agentId)
func (_Subnet *SubnetFilterer) FilterParticipantIdentityBound(opts *bind.FilterOpts, participant []common.Address, agentId []*big.Int) (*SubnetParticipantIdentityBoundIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantIdentityBound", participantRule, agentIdRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantIdentityBoundIterator{contract: _Subnet.contract, event: "ParticipantIdentityBound", logs: logs, sub: sub}, nil
}

// WatchParticipantIdentityBound is a free log subscription operation binding the contract event 0xd54ee5b46503fad8a00f45ed58660ef98a860ed5ace41d9b5795f91b3d0dc6d5.
//
// Solidity: event ParticipantIdentityBound(address indexed participant, uint256 indexed agentId)
func (_Subnet *SubnetFilterer) WatchParticipantIdentityBound(opts *bind.WatchOpts, sink chan<- *SubnetParticipantIdentityBound, participant []common.Address, agentId []*big.Int) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantIdentityBound", participantRule, agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantIdentityBound)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantIdentityBound", log); err != nil {
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

// ParseParticipantIdentityBound is a log parse operation binding the contract event 0xd54ee5b46503fad8a00f45ed58660ef98a860ed5ace41d9b5795f91b3d0dc6d5.
//
// Solidity: event ParticipantIdentityBound(address indexed participant, uint256 indexed agentId)
func (_Subnet *SubnetFilterer) ParseParticipantIdentityBound(log types.Log) (*SubnetParticipantIdentityBound, error) {
	event := new(SubnetParticipantIdentityBound)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantIdentityBound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantIdentityRegisteredIterator is returned from FilterParticipantIdentityRegistered and is used to iterate over the raw logs and unpacked data for ParticipantIdentityRegistered events raised by the Subnet contract.
type SubnetParticipantIdentityRegisteredIterator struct {
	Event *SubnetParticipantIdentityRegistered // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantIdentityRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantIdentityRegistered)
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
		it.Event = new(SubnetParticipantIdentityRegistered)
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
func (it *SubnetParticipantIdentityRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantIdentityRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantIdentityRegistered represents a ParticipantIdentityRegistered event raised by the Subnet contract.
type SubnetParticipantIdentityRegistered struct {
	Participant common.Address
	AgentId     *big.Int
	Endpoint    string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterParticipantIdentityRegistered is a free log retrieval operation binding the contract event 0x1cc884a47420b53b8bef39857ab08f65c5fa32e1b66d04efcbb602f3253492f1.
//
// Solidity: event ParticipantIdentityRegistered(address indexed participant, uint256 indexed agentId, string endpoint)
func (_Subnet *SubnetFilterer) FilterParticipantIdentityRegistered(opts *bind.FilterOpts, participant []common.Address, agentId []*big.Int) (*SubnetParticipantIdentityRegisteredIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantIdentityRegistered", participantRule, agentIdRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantIdentityRegisteredIterator{contract: _Subnet.contract, event: "ParticipantIdentityRegistered", logs: logs, sub: sub}, nil
}

// WatchParticipantIdentityRegistered is a free log subscription operation binding the contract event 0x1cc884a47420b53b8bef39857ab08f65c5fa32e1b66d04efcbb602f3253492f1.
//
// Solidity: event ParticipantIdentityRegistered(address indexed participant, uint256 indexed agentId, string endpoint)
func (_Subnet *SubnetFilterer) WatchParticipantIdentityRegistered(opts *bind.WatchOpts, sink chan<- *SubnetParticipantIdentityRegistered, participant []common.Address, agentId []*big.Int) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var agentIdRule []interface{}
	for _, agentIdItem := range agentId {
		agentIdRule = append(agentIdRule, agentIdItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantIdentityRegistered", participantRule, agentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantIdentityRegistered)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantIdentityRegistered", log); err != nil {
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

// ParseParticipantIdentityRegistered is a log parse operation binding the contract event 0x1cc884a47420b53b8bef39857ab08f65c5fa32e1b66d04efcbb602f3253492f1.
//
// Solidity: event ParticipantIdentityRegistered(address indexed participant, uint256 indexed agentId, string endpoint)
func (_Subnet *SubnetFilterer) ParseParticipantIdentityRegistered(log types.Log) (*SubnetParticipantIdentityRegistered, error) {
	event := new(SubnetParticipantIdentityRegistered)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantIdentityRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantRegisteredIterator is returned from FilterParticipantRegistered and is used to iterate over the raw logs and unpacked data for ParticipantRegistered events raised by the Subnet contract.
type SubnetParticipantRegisteredIterator struct {
	Event *SubnetParticipantRegistered // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantRegistered)
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
		it.Event = new(SubnetParticipantRegistered)
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
func (it *SubnetParticipantRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantRegistered represents a ParticipantRegistered event raised by the Subnet contract.
type SubnetParticipantRegistered struct {
	Addr            common.Address
	ParticipantType uint8
	Domain          string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterParticipantRegistered is a free log retrieval operation binding the contract event 0xe3564c470dc79d88d220de1a64bc335ec0a928450ce631293691bba84dd9a611.
//
// Solidity: event ParticipantRegistered(address indexed addr, uint8 indexed participant_type, string domain)
func (_Subnet *SubnetFilterer) FilterParticipantRegistered(opts *bind.FilterOpts, addr []common.Address, participant_type []uint8) (*SubnetParticipantRegisteredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var participant_typeRule []interface{}
	for _, participant_typeItem := range participant_type {
		participant_typeRule = append(participant_typeRule, participant_typeItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantRegistered", addrRule, participant_typeRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantRegisteredIterator{contract: _Subnet.contract, event: "ParticipantRegistered", logs: logs, sub: sub}, nil
}

// WatchParticipantRegistered is a free log subscription operation binding the contract event 0xe3564c470dc79d88d220de1a64bc335ec0a928450ce631293691bba84dd9a611.
//
// Solidity: event ParticipantRegistered(address indexed addr, uint8 indexed participant_type, string domain)
func (_Subnet *SubnetFilterer) WatchParticipantRegistered(opts *bind.WatchOpts, sink chan<- *SubnetParticipantRegistered, addr []common.Address, participant_type []uint8) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var participant_typeRule []interface{}
	for _, participant_typeItem := range participant_type {
		participant_typeRule = append(participant_typeRule, participant_typeItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantRegistered", addrRule, participant_typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantRegistered)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantRegistered", log); err != nil {
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

// ParseParticipantRegistered is a log parse operation binding the contract event 0xe3564c470dc79d88d220de1a64bc335ec0a928450ce631293691bba84dd9a611.
//
// Solidity: event ParticipantRegistered(address indexed addr, uint8 indexed participant_type, string domain)
func (_Subnet *SubnetFilterer) ParseParticipantRegistered(log types.Log) (*SubnetParticipantRegistered, error) {
	event := new(SubnetParticipantRegistered)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantRejectedIterator is returned from FilterParticipantRejected and is used to iterate over the raw logs and unpacked data for ParticipantRejected events raised by the Subnet contract.
type SubnetParticipantRejectedIterator struct {
	Event *SubnetParticipantRejected // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantRejected)
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
		it.Event = new(SubnetParticipantRejected)
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
func (it *SubnetParticipantRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantRejected represents a ParticipantRejected event raised by the Subnet contract.
type SubnetParticipantRejected struct {
	Participant     common.Address
	ParticipantType uint8
	Rejector        common.Address
	Reason          string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterParticipantRejected is a free log retrieval operation binding the contract event 0x911be13cdd33376f43fe83b43f0d8e53f2ae9f538271671647804ccf4d055f22.
//
// Solidity: event ParticipantRejected(address indexed participant, uint8 indexed participantType, address indexed rejector, string reason)
func (_Subnet *SubnetFilterer) FilterParticipantRejected(opts *bind.FilterOpts, participant []common.Address, participantType []uint8, rejector []common.Address) (*SubnetParticipantRejectedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var participantTypeRule []interface{}
	for _, participantTypeItem := range participantType {
		participantTypeRule = append(participantTypeRule, participantTypeItem)
	}
	var rejectorRule []interface{}
	for _, rejectorItem := range rejector {
		rejectorRule = append(rejectorRule, rejectorItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantRejected", participantRule, participantTypeRule, rejectorRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantRejectedIterator{contract: _Subnet.contract, event: "ParticipantRejected", logs: logs, sub: sub}, nil
}

// WatchParticipantRejected is a free log subscription operation binding the contract event 0x911be13cdd33376f43fe83b43f0d8e53f2ae9f538271671647804ccf4d055f22.
//
// Solidity: event ParticipantRejected(address indexed participant, uint8 indexed participantType, address indexed rejector, string reason)
func (_Subnet *SubnetFilterer) WatchParticipantRejected(opts *bind.WatchOpts, sink chan<- *SubnetParticipantRejected, participant []common.Address, participantType []uint8, rejector []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var participantTypeRule []interface{}
	for _, participantTypeItem := range participantType {
		participantTypeRule = append(participantTypeRule, participantTypeItem)
	}
	var rejectorRule []interface{}
	for _, rejectorItem := range rejector {
		rejectorRule = append(rejectorRule, rejectorItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantRejected", participantRule, participantTypeRule, rejectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantRejected)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantRejected", log); err != nil {
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

// ParseParticipantRejected is a log parse operation binding the contract event 0x911be13cdd33376f43fe83b43f0d8e53f2ae9f538271671647804ccf4d055f22.
//
// Solidity: event ParticipantRejected(address indexed participant, uint8 indexed participantType, address indexed rejector, string reason)
func (_Subnet *SubnetFilterer) ParseParticipantRejected(log types.Log) (*SubnetParticipantRejected, error) {
	event := new(SubnetParticipantRejected)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantStakedIterator is returned from FilterParticipantStaked and is used to iterate over the raw logs and unpacked data for ParticipantStaked events raised by the Subnet contract.
type SubnetParticipantStakedIterator struct {
	Event *SubnetParticipantStaked // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantStaked)
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
		it.Event = new(SubnetParticipantStaked)
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
func (it *SubnetParticipantStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantStaked represents a ParticipantStaked event raised by the Subnet contract.
type SubnetParticipantStaked struct {
	Participant     common.Address
	ParticipantType uint8
	StakedAmount    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterParticipantStaked is a free log retrieval operation binding the contract event 0x1304e75cb7d7b5caf45a5614e9da12e06f2e5ce43fe410b18986ef8347bbd9ee.
//
// Solidity: event ParticipantStaked(address indexed participant, uint8 indexed participantType, uint256 stakedAmount)
func (_Subnet *SubnetFilterer) FilterParticipantStaked(opts *bind.FilterOpts, participant []common.Address, participantType []uint8) (*SubnetParticipantStakedIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var participantTypeRule []interface{}
	for _, participantTypeItem := range participantType {
		participantTypeRule = append(participantTypeRule, participantTypeItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantStaked", participantRule, participantTypeRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantStakedIterator{contract: _Subnet.contract, event: "ParticipantStaked", logs: logs, sub: sub}, nil
}

// WatchParticipantStaked is a free log subscription operation binding the contract event 0x1304e75cb7d7b5caf45a5614e9da12e06f2e5ce43fe410b18986ef8347bbd9ee.
//
// Solidity: event ParticipantStaked(address indexed participant, uint8 indexed participantType, uint256 stakedAmount)
func (_Subnet *SubnetFilterer) WatchParticipantStaked(opts *bind.WatchOpts, sink chan<- *SubnetParticipantStaked, participant []common.Address, participantType []uint8) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var participantTypeRule []interface{}
	for _, participantTypeItem := range participantType {
		participantTypeRule = append(participantTypeRule, participantTypeItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantStaked", participantRule, participantTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantStaked)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantStaked", log); err != nil {
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

// ParseParticipantStaked is a log parse operation binding the contract event 0x1304e75cb7d7b5caf45a5614e9da12e06f2e5ce43fe410b18986ef8347bbd9ee.
//
// Solidity: event ParticipantStaked(address indexed participant, uint8 indexed participantType, uint256 stakedAmount)
func (_Subnet *SubnetFilterer) ParseParticipantStaked(log types.Log) (*SubnetParticipantStaked, error) {
	event := new(SubnetParticipantStaked)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantStaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetParticipantStatusChangedIterator is returned from FilterParticipantStatusChanged and is used to iterate over the raw logs and unpacked data for ParticipantStatusChanged events raised by the Subnet contract.
type SubnetParticipantStatusChangedIterator struct {
	Event *SubnetParticipantStatusChanged // Event containing the contract specifics and raw log

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
func (it *SubnetParticipantStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetParticipantStatusChanged)
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
		it.Event = new(SubnetParticipantStatusChanged)
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
func (it *SubnetParticipantStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetParticipantStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetParticipantStatusChanged represents a ParticipantStatusChanged event raised by the Subnet contract.
type SubnetParticipantStatusChanged struct {
	Addr            common.Address
	ParticipantType uint8
	Status          uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterParticipantStatusChanged is a free log retrieval operation binding the contract event 0x0a4b355e037ef9ee17aaae57aea8f4bc674d63b8b77b33523e3d0b488c02c215.
//
// Solidity: event ParticipantStatusChanged(address indexed addr, uint8 indexed participant_type, uint8 status)
func (_Subnet *SubnetFilterer) FilterParticipantStatusChanged(opts *bind.FilterOpts, addr []common.Address, participant_type []uint8) (*SubnetParticipantStatusChangedIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var participant_typeRule []interface{}
	for _, participant_typeItem := range participant_type {
		participant_typeRule = append(participant_typeRule, participant_typeItem)
	}

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "ParticipantStatusChanged", addrRule, participant_typeRule)
	if err != nil {
		return nil, err
	}
	return &SubnetParticipantStatusChangedIterator{contract: _Subnet.contract, event: "ParticipantStatusChanged", logs: logs, sub: sub}, nil
}

// WatchParticipantStatusChanged is a free log subscription operation binding the contract event 0x0a4b355e037ef9ee17aaae57aea8f4bc674d63b8b77b33523e3d0b488c02c215.
//
// Solidity: event ParticipantStatusChanged(address indexed addr, uint8 indexed participant_type, uint8 status)
func (_Subnet *SubnetFilterer) WatchParticipantStatusChanged(opts *bind.WatchOpts, sink chan<- *SubnetParticipantStatusChanged, addr []common.Address, participant_type []uint8) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var participant_typeRule []interface{}
	for _, participant_typeItem := range participant_type {
		participant_typeRule = append(participant_typeRule, participant_typeItem)
	}

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "ParticipantStatusChanged", addrRule, participant_typeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetParticipantStatusChanged)
				if err := _Subnet.contract.UnpackLog(event, "ParticipantStatusChanged", log); err != nil {
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

// ParseParticipantStatusChanged is a log parse operation binding the contract event 0x0a4b355e037ef9ee17aaae57aea8f4bc674d63b8b77b33523e3d0b488c02c215.
//
// Solidity: event ParticipantStatusChanged(address indexed addr, uint8 indexed participant_type, uint8 status)
func (_Subnet *SubnetFilterer) ParseParticipantStatusChanged(log types.Log) (*SubnetParticipantStatusChanged, error) {
	event := new(SubnetParticipantStatusChanged)
	if err := _Subnet.contract.UnpackLog(event, "ParticipantStatusChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Subnet contract.
type SubnetPausedIterator struct {
	Event *SubnetPaused // Event containing the contract specifics and raw log

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
func (it *SubnetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetPaused)
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
		it.Event = new(SubnetPaused)
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
func (it *SubnetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetPaused represents a Paused event raised by the Subnet contract.
type SubnetPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Subnet *SubnetFilterer) FilterPaused(opts *bind.FilterOpts) (*SubnetPausedIterator, error) {

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SubnetPausedIterator{contract: _Subnet.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Subnet *SubnetFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SubnetPaused) (event.Subscription, error) {

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetPaused)
				if err := _Subnet.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Subnet *SubnetFilterer) ParsePaused(log types.Log) (*SubnetPaused, error) {
	event := new(SubnetPaused)
	if err := _Subnet.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Subnet contract.
type SubnetRoleAdminChangedIterator struct {
	Event *SubnetRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SubnetRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetRoleAdminChanged)
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
		it.Event = new(SubnetRoleAdminChanged)
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
func (it *SubnetRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetRoleAdminChanged represents a RoleAdminChanged event raised by the Subnet contract.
type SubnetRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Subnet *SubnetFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SubnetRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SubnetRoleAdminChangedIterator{contract: _Subnet.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Subnet *SubnetFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SubnetRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetRoleAdminChanged)
				if err := _Subnet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Subnet *SubnetFilterer) ParseRoleAdminChanged(log types.Log) (*SubnetRoleAdminChanged, error) {
	event := new(SubnetRoleAdminChanged)
	if err := _Subnet.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Subnet contract.
type SubnetRoleGrantedIterator struct {
	Event *SubnetRoleGranted // Event containing the contract specifics and raw log

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
func (it *SubnetRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetRoleGranted)
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
		it.Event = new(SubnetRoleGranted)
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
func (it *SubnetRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetRoleGranted represents a RoleGranted event raised by the Subnet contract.
type SubnetRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subnet *SubnetFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubnetRoleGrantedIterator, error) {

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

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubnetRoleGrantedIterator{contract: _Subnet.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subnet *SubnetFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SubnetRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetRoleGranted)
				if err := _Subnet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Subnet *SubnetFilterer) ParseRoleGranted(log types.Log) (*SubnetRoleGranted, error) {
	event := new(SubnetRoleGranted)
	if err := _Subnet.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Subnet contract.
type SubnetRoleRevokedIterator struct {
	Event *SubnetRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SubnetRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetRoleRevoked)
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
		it.Event = new(SubnetRoleRevoked)
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
func (it *SubnetRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetRoleRevoked represents a RoleRevoked event raised by the Subnet contract.
type SubnetRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subnet *SubnetFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SubnetRoleRevokedIterator, error) {

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

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SubnetRoleRevokedIterator{contract: _Subnet.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Subnet *SubnetFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SubnetRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetRoleRevoked)
				if err := _Subnet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Subnet *SubnetFilterer) ParseRoleRevoked(log types.Log) (*SubnetRoleRevoked, error) {
	event := new(SubnetRoleRevoked)
	if err := _Subnet.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetSubnetConfigUpdatedIterator is returned from FilterSubnetConfigUpdated and is used to iterate over the raw logs and unpacked data for SubnetConfigUpdated events raised by the Subnet contract.
type SubnetSubnetConfigUpdatedIterator struct {
	Event *SubnetSubnetConfigUpdated // Event containing the contract specifics and raw log

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
func (it *SubnetSubnetConfigUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetSubnetConfigUpdated)
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
		it.Event = new(SubnetSubnetConfigUpdated)
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
func (it *SubnetSubnetConfigUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetSubnetConfigUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetSubnetConfigUpdated represents a SubnetConfigUpdated event raised by the Subnet contract.
type SubnetSubnetConfigUpdated struct {
	ParamsHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubnetConfigUpdated is a free log retrieval operation binding the contract event 0x043a29de91e100fb7eb5924307f5a54ff0dc2d578a058354170a2f0cca43ccb1.
//
// Solidity: event SubnetConfigUpdated(bytes32 params_hash)
func (_Subnet *SubnetFilterer) FilterSubnetConfigUpdated(opts *bind.FilterOpts) (*SubnetSubnetConfigUpdatedIterator, error) {

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "SubnetConfigUpdated")
	if err != nil {
		return nil, err
	}
	return &SubnetSubnetConfigUpdatedIterator{contract: _Subnet.contract, event: "SubnetConfigUpdated", logs: logs, sub: sub}, nil
}

// WatchSubnetConfigUpdated is a free log subscription operation binding the contract event 0x043a29de91e100fb7eb5924307f5a54ff0dc2d578a058354170a2f0cca43ccb1.
//
// Solidity: event SubnetConfigUpdated(bytes32 params_hash)
func (_Subnet *SubnetFilterer) WatchSubnetConfigUpdated(opts *bind.WatchOpts, sink chan<- *SubnetSubnetConfigUpdated) (event.Subscription, error) {

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "SubnetConfigUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetSubnetConfigUpdated)
				if err := _Subnet.contract.UnpackLog(event, "SubnetConfigUpdated", log); err != nil {
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

// ParseSubnetConfigUpdated is a log parse operation binding the contract event 0x043a29de91e100fb7eb5924307f5a54ff0dc2d578a058354170a2f0cca43ccb1.
//
// Solidity: event SubnetConfigUpdated(bytes32 params_hash)
func (_Subnet *SubnetFilterer) ParseSubnetConfigUpdated(log types.Log) (*SubnetSubnetConfigUpdated, error) {
	event := new(SubnetSubnetConfigUpdated)
	if err := _Subnet.contract.UnpackLog(event, "SubnetConfigUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SubnetUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Subnet contract.
type SubnetUnpausedIterator struct {
	Event *SubnetUnpaused // Event containing the contract specifics and raw log

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
func (it *SubnetUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SubnetUnpaused)
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
		it.Event = new(SubnetUnpaused)
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
func (it *SubnetUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SubnetUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SubnetUnpaused represents a Unpaused event raised by the Subnet contract.
type SubnetUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Subnet *SubnetFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SubnetUnpausedIterator, error) {

	logs, sub, err := _Subnet.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SubnetUnpausedIterator{contract: _Subnet.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Subnet *SubnetFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SubnetUnpaused) (event.Subscription, error) {

	logs, sub, err := _Subnet.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SubnetUnpaused)
				if err := _Subnet.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Subnet *SubnetFilterer) ParseUnpaused(log types.Log) (*SubnetUnpaused, error) {
	event := new(SubnetUnpaused)
	if err := _Subnet.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
