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
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"feedback_id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to_agent\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.FeedbackType\",\"name\":\"feedback_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"rating\",\"type\":\"uint8\"}],\"name\":\"FeedbackSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"}],\"name\":\"ParticipantActivityUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participantType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ParticipantApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"ParticipantRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participantType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rejector\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"ParticipantRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participantType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"}],\"name\":\"ParticipantStaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"ParticipantStatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"params_hash\",\"type\":\"bytes32\"}],\"name\":\"SubnetConfigUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"activeParticipantCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"approveParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getActiveParticipantCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidFrequencyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCheckpointPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getMinimumStakeByType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getParticipantCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getParticipantInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint80\",\"name\":\"reputation_score\",\"type\":\"uint80\"},{\"internalType\":\"uint128\",\"name\":\"registered_at\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"reserved\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"internalType\":\"structDataStructures.ParticipantInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"getParticipantStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestAmount\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequireKYC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignatureThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSlashingRates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakeGovernanceConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSubnetInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalParticipantCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnstakeLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"hasAvailableSlots\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"subnet_id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"canonical_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"da_kind\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sig_scheme\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"autoApprove\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireKYC\",\"type\":\"bool\"},{\"internalType\":\"enumDataStructures.SubnetStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"challenge_window\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"min_epoch_interval\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"max_epoch_interval\",\"type\":\"uint64\"}],\"internalType\":\"structDataStructures.CheckpointPolicy\",\"name\":\"cp_policy\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"threshold_numerator\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"threshold_denominator\",\"type\":\"uint32\"}],\"internalType\":\"structDataStructures.SignatureThreshold\",\"name\":\"sig_threshold\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"stake_cfg\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updated_at\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidFrequencyLimit\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.SubnetInfo\",\"name\":\"_subnetInfo\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stakingManager_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"isActiveParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"listActiveParticipants\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint80\",\"name\":\"reputation_score\",\"type\":\"uint80\"},{\"internalType\":\"uint128\",\"name\":\"registered_at\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"reserved\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"internalType\":\"structDataStructures.ParticipantInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participantListByType\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"enumDataStructures.ParticipantStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint80\",\"name\":\"reputation_score\",\"type\":\"uint80\"},{\"internalType\":\"uint128\",\"name\":\"registered_at\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"last_active\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"reserved\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"name\":\"registerParticipant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadata_uri\",\"type\":\"string\"}],\"name\":\"registerParticipantERC20\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"rejectParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant_addr\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"resumeParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minValidatorStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAgentStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minMatcherStake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAgents\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxMatchers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"slashingRates\",\"type\":\"uint256[]\"}],\"internalType\":\"structDataStructures.StakeGovernanceConfig\",\"name\":\"_stakeConfig\",\"type\":\"tuple\"}],\"name\":\"setStakeConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"suspendParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"participant_type\",\"type\":\"uint8\"}],\"name\":\"updateParticipantActivity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b5061499c8061001f6000396000f3fe6080604052600436106101e15760003560e01c806301ffc9a7146101e65780630fcc0c281461021b5780631da65cc314610232578063248a9ca3146102525780632c5c1222146102805780632cae8f81146102ad5780632f2ff15d146102cf5780633451479e146102ef57806336568abe1461030f5780633d1dba561461032f5780633f4ba83a1461034f57806340b33161146103645780634aee0b92146103795780634b2875661461038e578063509e8745146103a35780635c975abb146103c55780635caecd9c146103da5780636ee88b1614610407578063708b73a214610427578063735d23e81461044757806373a3e8711461047457806373a6d7c51461049157806375b238fc146104b15780638456cb59146104d35780638941382a146104e857806391d14854146104fb57806396611afd1461051b5780639b0a82f61461053b5780639f9106d11461055b578063a217fddf1461057d578063ac6c2a4514610592578063ac921890146105b2578063b8e90b67146105d2578063bd352a2b146105f2578063c0a6367114610612578063ca8cf96614610632578063ce89f1de14610668578063d15442ad1461068a578063d2b4353f146106ac578063d547741f146106ce578063e758b8ae146106ee578063fa7366001461070e575b600080fd5b3480156101f257600080fd5b506102066102013660046138fe565b610750565b60405190151581526020015b60405180910390f35b34801561022757600080fd5b50610230610787565b005b34801561023e57600080fd5b5061023061024d36600461393c565b61083a565b34801561025e57600080fd5b5061027261026d366004613959565b6109f6565b604051908152602001610212565b34801561028c57600080fd5b5061027261029b36600461393c565b60156020526000908152604090205481565b3480156102b957600080fd5b506102c2610a16565b604051610212919061398a565b3480156102db57600080fd5b506102306102ea3660046139ad565b610a47565b3480156102fb57600080fd5b5061023061030a3660046139dd565b610a69565b34801561031b57600080fd5b5061023061032a3660046139ad565b610bd7565b34801561033b57600080fd5b5061027261034a36600461393c565b610c0f565b34801561035b57600080fd5b50610230610c4e565b34801561037057600080fd5b50610272610d2b565b34801561038557600080fd5b50601354610272565b34801561039a57600080fd5b50600e54610272565b3480156103af57600080fd5b506103b8610dd5565b6040516102129190613a0b565b3480156103d157600080fd5b50610206610e2c565b3480156103e657600080fd5b506103fa6103f536600461393c565b610e41565b6040516102129190613bca565b34801561041357600080fd5b50610230610422366004613c77565b61135e565b34801561043357600080fd5b50610230610442366004613cdb565b611543565b34801561045357600080fd5b506104676104623660046139dd565b611720565b6040516102129190613d44565b34801561048057600080fd5b50600554610100900460ff16610206565b34801561049d57600080fd5b506102066104ac366004613d57565b611a02565b3480156104bd57600080fd5b5061027260008051602061494783398151915281565b3480156104df57600080fd5b50610230611a3a565b6102066104f6366004613e16565b611b0a565b34801561050757600080fd5b506102066105163660046139ad565b611b41565b34801561052757600080fd5b50610230610536366004613ecc565b611b77565b34801561054757600080fd5b5061027261055636600461393c565b611d91565b34801561056757600080fd5b50610570611d9c565b6040516102129190613f20565b34801561058957600080fd5b50610272600081565b34801561059e57600080fd5b506105706105ad366004613f34565b611e0f565b3480156105be57600080fd5b506102066105cd3660046139dd565b611e47565b3480156105de57600080fd5b506102306105ed3660046139dd565b611eba565b3480156105fe57600080fd5b5061020661060d36600461393c565b612083565b34801561061e57600080fd5b5061027261062d36600461393c565b6120e3565b34801561063e57600080fd5b5061065261064d3660046139dd565b612123565b6040516102129a99989796959493929190613f60565b34801561067457600080fd5b5061067d612345565b6040516102129190614096565b34801561069657600080fd5b5061069f6123eb565b60405161021291906140da565b3480156106b857600080fd5b506106c1612809565b6040516102129190614233565b3480156106da57600080fd5b506102306106e93660046139ad565b612854565b3480156106fa57600080fd5b50610230610709366004614241565b612870565b34801561071a57600080fd5b5061072e6107293660046139dd565b612895565b6040805182518152602080840151908201529181015190820152606001610212565b60006001600160e01b03198216637965db0b60e01b148061078157506301ffc9a760e01b6001600160e01b03198316145b92915050565b60008051602061494783398151915261079f81612941565b60006003600582015462010000900460ff1660038111156107c2576107c2613a5b565b036108095760405162461bcd60e51b8152602060048201526012602482015271105b1c9958591e4819195c1c9958d85d195960721b60448201526064015b60405180910390fd5b60058101805462ff000019166203000017905542601282015561082a610e2c565b6108365761083661294e565b5050565b3360008181526014602052604081208391908183600381111561085f5761085f613a5b565b600381111561087057610870613a5b565b81526020810191909152604001600020546001600160a01b0316036108a75760405162461bcd60e51b81526004016108009061427c565b33600081815260146020526040812085918391908360038111156108cd576108cd613a5b565b60038111156108de576108de613a5b565b81526020810191909152604001600020546001600160a01b03161461093d5760405162461bcd60e51b81526020600482015260156024820152742737ba103830b93a34b1b4b830b73a1037bbb732b960591b6044820152606401610800565b336000908152601460205260408120429187600381111561096057610960613a5b565b600381111561097157610971613a5b565b8152602081019190915260400160002060010180546001600160801b03928316600160801b0292169190911790558460038111156109b1576109b1613a5b565b6040516001600160801b034216815233907f6b8db42cee7260f00718e255640423e4d8dd1b15b9ce881aaee80e5a45e7d0779060200160405180910390a35050505050565b600080610a016129ad565b60009384526020525050604090206001015490565b610a1e61372e565b506040805180820190915260075463ffffffff8082168352600160201b90910416602082015290565b610a50826109f6565b610a5981612941565b610a6383836129d1565b50505050565b600080516020614947833981519152610a8181612941565b6001600160a01b038316600090815260146020526040812081846003811115610aac57610aac613a5b565b6003811115610abd57610abd613a5b565b8152602081019190915260400160002080549091506001600160a01b0316610af75760405162461bcd60e51b81526004016108009061427c565b60028154600160a81b900460ff166002811115610b1657610b16613a5b565b14610b335760405162461bcd60e51b8152600401610800906142ac565b805460ff60a81b1916815560156000846003811115610b5457610b54613a5b565b6003811115610b6557610b65613a5b565b81526020019081526020016000206000815480929190610b84906142e7565b90915550339050836003811115610b9d57610b9d613a5b565b6040516001600160a01b038716907f824961050c68118de8b5a2c2d827932d4a9efc756d8252776df33ca28b0e775490600090a450505050565b6001600160a01b0381163314610c005760405163334bd91960e11b815260040160405180910390fd5b610c0a8282612a72565b505050565b600060156000836003811115610c2757610c27613a5b565b6003811115610c3857610c38613a5b565b8152602001908152602001600020549050919050565b600080516020614947833981519152610c6681612941565b60006003600582015462010000900460ff166003811115610c8957610c89613a5b565b14158015610cb657506001600582015462010000900460ff166003811115610cb357610cb3613a5b565b14155b610cfe5760405162461bcd60e51b815260206004820152601960248201527814dd589b995d0818d85b9b9bdd081899481d5b9c185d5cd959603a1b6044820152606401610800565b610d06612aea565b6005810180546001919062ff0000191662010000835b02179055504260129091015550565b600080805260166020527f0263c2b778d062355049effc2dece97bc6547ff8a88a3258daa512061c2153dd548190610d639082614300565b600160005260166020527f4c4dc693d7db52f85fe052106f4b4b920e78e8ef37dee82878a60ab8585faf4954909150610d9c9082614300565b600260005260166020527fcaff291fe014adc6b72a172705750b4cabe8f8667664d2924a166caab2885648549091506107819082614300565b600f80546040805160208084028201810190925282815260609390929091830182828015610e2257602002820191906000526020600020905b815481526020019060010190808311610e0e575b5050505050905090565b600080610e37612b2d565b5460ff1692915050565b6060600060166000846003811115610e5b57610e5b613a5b565b6003811115610e6c57610e6c613a5b565b81526020019081526020016000209050600081805490506001600160401b03811115610e9a57610e9a614313565b604051908082528060200260200182016040528015610ec3578160200160208202803683370190505b5090506000805b8354811015610fd457600060146000868481548110610eeb57610eeb614329565b60009182526020808320909101546001600160a01b03168352820192909252604001812090886003811115610f2257610f22613a5b565b6003811115610f3357610f33613a5b565b8152602081019190915260400160002054600160a81b900460ff166002811115610f5f57610f5f613a5b565b03610fcc57838181548110610f7657610f76614329565b9060005260206000200160009054906101000a90046001600160a01b0316838381518110610fa657610fa6614329565b6001600160a01b039092166020928302919091019091015281610fc8816142e7565b9250505b600101610eca565b506000816001600160401b03811115610fef57610fef614313565b60405190808252806020026020018201604052801561102857816020015b611015613745565b81526020019060019003908161100d5790505b50905060005b82811015611354576014600085838151811061104c5761104c614329565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600088600381111561108857611088613a5b565b600381111561109957611099613a5b565b815260208082019290925260409081016000208151610140810190925280546001600160a01b03811683529192909190830190600160a01b900460ff1660038111156110e7576110e7613a5b565b60038111156110f8576110f8613a5b565b81528154602090910190600160a81b900460ff16600281111561111d5761111d613a5b565b600281111561112e5761112e613a5b565b81528154600160b01b90046001600160501b0316602082015260018201546001600160801b038082166040840152600160801b9091041660608201526002820154608082015260038201805460a0909201916111899061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546111b59061433f565b80156112025780601f106111d757610100808354040283529160200191611202565b820191906000526020600020905b8154815290600101906020018083116111e557829003601f168201915b5050505050815260200160048201805461121b9061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546112479061433f565b80156112945780601f1061126957610100808354040283529160200191611294565b820191906000526020600020905b81548152906001019060200180831161127757829003601f168201915b505050505081526020016005820180546112ad9061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546112d99061433f565b80156113265780601f106112fb57610100808354040283529160200191611326565b820191906000526020600020905b81548152906001019060200180831161130957829003601f168201915b50505050508152505082828151811061134157611341614329565b602090810291909101015260010161102e565b5095945050505050565b60008051602061494783398151915261137681612941565b6001600160a01b0385166000908152601460205260408120818660038111156113a1576113a1613a5b565b60038111156113b2576113b2613a5b565b8152602081019190915260400160002080549091506001600160a01b03166113ec5760405162461bcd60e51b81526004016108009061427c565b60028154600160a81b900460ff16600281111561140b5761140b613a5b565b146114285760405162461bcd60e51b8152600401610800906142ac565b6017546001600160a01b03166114505760405162461bcd60e51b815260040161080090614379565b6001600160a01b03861660009081526014602052604081209086600381111561147b5761147b613a5b565b600381111561148c5761148c613a5b565b8152602081019190915260400160009081208181556001810182905560028101829055906114bd60038301826137b4565b6114cb6004830160006137b4565b6114d96005830160006137b4565b50506114e58686612b51565b338560038111156114f8576114f8613a5b565b876001600160a01b03167f911be13cdd33376f43fe83b43f0d8e53f2ae9f538271671647804ccf4d055f2287876040516115339291906143aa565b60405180910390a4505050505050565b600061154d612c89565b805490915060ff600160401b82041615906001600160401b03166000811580156115745750825b90506000826001600160401b031660011480156115905750303b155b90508115801561159e575080155b156115bc5760405163f92ee8a960e01b815260040160405180910390fd5b84546001600160401b031916600117855583156115e557845460ff60401b1916600160401b1785555b6001600160a01b0387161580159061160557506001600160a01b03861615155b61163d5760405162461bcd60e51b815260206004820152600960248201526806164647265737320360bc1b6044820152606401610800565b611645612cb2565b61164d612cb2565b611655612cbc565b61165e88612ccc565b611678600061167360608b0160408c016143d9565b6129d1565b5061169b60008051602061494783398151915261167360608b0160408c016143d9565b506116b4600080516020614947833981519152886129d1565b50601780546001600160a01b0319166001600160a01b038816179055831561171657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b611728613745565b6001600160a01b03831660009081526014602052604081209083600381111561175357611753613a5b565b600381111561176457611764613a5b565b815260208082019290925260409081016000208151610140810190925280546001600160a01b03811683529192909190830190600160a01b900460ff1660038111156117b2576117b2613a5b565b60038111156117c3576117c3613a5b565b81528154602090910190600160a81b900460ff1660028111156117e8576117e8613a5b565b60028111156117f9576117f9613a5b565b81528154600160b01b90046001600160501b0316602082015260018201546001600160801b038082166040840152600160801b9091041660608201526002820154608082015260038201805460a0909201916118549061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546118809061433f565b80156118cd5780601f106118a2576101008083540402835291602001916118cd565b820191906000526020600020905b8154815290600101906020018083116118b057829003601f168201915b505050505081526020016004820180546118e69061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546119129061433f565b801561195f5780601f106119345761010080835404028352916020019161195f565b820191906000526020600020905b81548152906001019060200180831161194257829003601f168201915b505050505081526020016005820180546119789061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546119a49061433f565b80156119f15780601f106119c6576101008083540402835291602001916119f1565b820191906000526020600020905b8154815290600101906020018083116119d457829003601f168201915b505050505081525050905092915050565b6000611a0c612ea7565b611a14612edd565b611a24898888888888888f612f03565b9050611a2e6132d7565b98975050505050505050565b600080516020614947833981519152611a5281612941565b60006003600582015462010000900460ff166003811115611a7557611a75613a5b565b14158015611aa257506002600582015462010000900460ff166003811115611a9f57611a9f613a5b565b14155b611ae85760405162461bcd60e51b815260206004820152601760248201527614dd589b995d0818d85b9b9bdd081899481c185d5cd959604a1b6044820152606401610800565b611af061294e565b6005810180546002919062ff000019166201000083610d1c565b6000611b14612ea7565b611b1c612edd565b611b2c8888888888888834612f03565b9050611b366132d7565b979650505050505050565b600080611b4c6129ad565b6000948552602090815260408086206001600160a01b03959095168652939052505090205460ff1690565b33600081815260146020526040812085919081836003811115611b9c57611b9c613a5b565b6003811115611bad57611bad613a5b565b81526020810191909152604001600020546001600160a01b031603611be45760405162461bcd60e51b81526004016108009061427c565b33600090815260146020526040812081876003811115611c0657611c06613a5b565b6003811115611c1757611c17613a5b565b81526020810191909152604001600090812091508154600160a81b900460ff166002811115611c4857611c48613a5b565b14611c825760405162461bcd60e51b815260206004820152600a6024820152694e6f742061637469766560b01b6044820152606401610800565b611c9a60008051602061494783398151915233611b41565b80611cae575080546001600160a01b031633145b611ceb5760405162461bcd60e51b815260206004820152600e60248201526d139bdd08185d5d1a1bdc9a5e995960921b6044820152606401610800565b805460ff60a81b1916600160a81b17815560156000876003811115611d1257611d12613a5b565b6003811115611d2357611d23613a5b565b81526020019081526020016000206000815480929190611d42906143f6565b9190505550856003811115611d5957611d59613a5b565b336001600160a01b03166000805160206149278339815191526001604051611d81919061440d565b60405180910390a3505050505050565b6000610781826132e8565b60175460408051639f9106d160e01b815290516000926001600160a01b031691639f9106d19160048083019260209291908290030181865afa158015611de6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e0a919061441b565b905090565b60166020528160005260406000208181548110611e2b57600080fd5b6000918252602090912001546001600160a01b03169150829050565b6000806001600160a01b038416600090815260146020526040812090846003811115611e7557611e75613a5b565b6003811115611e8657611e86613a5b565b8152602081019190915260400160002054600160a81b900460ff166002811115611eb257611eb2613a5b565b149392505050565b6001600160a01b03821660009081526014602052604081208391839181836003811115611ee957611ee9613a5b565b6003811115611efa57611efa613a5b565b81526020810191909152604001600020546001600160a01b031603611f315760405162461bcd60e51b81526004016108009061427c565b600080516020614947833981519152611f4981612941565b6001600160a01b038516600090815260146020526040812081866003811115611f7457611f74613a5b565b6003811115611f8557611f85613a5b565b81526020810191909152604001600020905060018154600160a81b900460ff166002811115611fb657611fb6613a5b565b14611ff35760405162461bcd60e51b815260206004820152600d60248201526c139bdd081cdd5cdc195b991959609a1b6044820152606401610800565b805460ff60a81b191681556015600086600381111561201457612014613a5b565b600381111561202557612025613a5b565b81526020019081526020016000206000815480929190612044906142e7565b919050555084600381111561205b5761205b613a5b565b866001600160a01b03166000805160206149278339815191526000604051611d81919061440d565b60008061208f83613354565b9050806000036120a25750600192915050565b80601660008560038111156120b9576120b9613a5b565b60038111156120ca576120ca613a5b565b8152602081019190915260400160002054109392505050565b6000601660008360038111156120fb576120fb613a5b565b600381111561210c5761210c613a5b565b815260208101919091526040016000205492915050565b601460209081526000928352604080842090915290825290208054600182015460028301546003840180546001600160a01b03851695600160a01b860460ff90811696600160a81b810490911695600160b01b9091046001600160501b0316946001600160801b0380831695600160801b9093041693909291906121a69061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546121d29061433f565b801561221f5780601f106121f45761010080835404028352916020019161221f565b820191906000526020600020905b81548152906001019060200180831161220257829003601f168201915b5050505050908060040180546122349061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546122609061433f565b80156122ad5780601f10612282576101008083540402835291602001916122ad565b820191906000526020600020905b81548152906001019060200180831161229057829003601f168201915b5050505050908060050180546122c29061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546122ee9061433f565b801561233b5780601f106123105761010080835404028352916020019161233b565b820191906000526020600020905b81548152906001019060200180831161231e57829003601f168201915b505050505090508a565b61234d6137ee565b6040805161010081018252600880548252600954602080840191909152600a5483850152600b546060840152600c546080840152600d5460a0840152600e5460c0840152600f80548551818402810184019096528086529394929360e0860193928301828280156123dd57602002820191906000526020600020905b8154815260200190600101908083116123c9575b505050505081525050905090565b6123f3613833565b604080516102008101909152600080548252600180546020840191906124189061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546124449061433f565b80156124915780601f1061246657610100808354040283529160200191612491565b820191906000526020600020905b81548152906001019060200180831161247457829003601f168201915b505050918352505060028201546001600160a01b0381166020830152600160a01b900461ffff1660408201526003820180546060909201916124d29061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546124fe9061433f565b801561254b5780601f106125205761010080835404028352916020019161254b565b820191906000526020600020905b81548152906001019060200180831161252e57829003601f168201915b505050505081526020016004820180546125649061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546125909061433f565b80156125dd5780601f106125b2576101008083540402835291602001916125dd565b820191906000526020600020905b8154815290600101906020018083116125c057829003601f168201915b5050509183525050600582015460ff8082161515602084015261010082048116151560408401526060909201916201000090910416600381111561262357612623613a5b565b600381111561263457612634613a5b565b8152604080516060808201835260068501546001600160801b03811683526001600160401b03600160801b82048116602085810191909152600160c01b90920416838501528085019290925282518084018452600786015463ffffffff8082168352600160201b9091041681840152838501528251610100810184526008860180548252600987015482850152600a87015482860152600b87015482840152600c8701546080830152600d87015460a0830152600e87015460c0830152600f8701805486518187028101870190975280875293909601959194909360e08601939192919083018282801561274757602002820191906000526020600020905b815481526020019060010190808311612733575b50505050508152505081526020016010820180546127649061433f565b80601f01602080910402602001604051908101604052809291908181526020018280546127909061433f565b80156127dd5780601f106127b2576101008083540402835291602001916127dd565b820191906000526020600020905b8154815290600101906020018083116127c057829003601f168201915b505050505081526020016011820154815260200160128201548152602001601382015481525050905090565b6128116138c5565b50604080516060810182526006546001600160801b03811682526001600160401b03600160801b820481166020840152600160c01b909104169181019190915290565b61285d826109f6565b61286681612941565b610a638383612a72565b60008051602061494783398151915261288881612941565b816008610a6382826144d0565b6128b960405180606001604052806000815260200160008152602001600081525090565b60006128c3611d9c565b6017546040516322151bbb60e11b81529192506001600160a01b03169063442a3776906128f890879085908890600401614569565b606060405180830381865afa158015612915573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129399190614598565b949350505050565b61294b81336133b8565b50565b612956612edd565b6000612960612b2d565b805460ff1916600117815590507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586129953390565b6040516129a29190613f20565b60405180910390a150565b7f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680090565b6000806129dc6129ad565b90506129e88484611b41565b612a68576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055612a1e3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610781565b6000915050610781565b600080612a7d6129ad565b9050612a898484611b41565b15612a68576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610781565b612af26133e3565b6000612afc612b2d565b805460ff1916815590507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa33612995565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330090565b600060166000836003811115612b6957612b69613a5b565b6003811115612b7a57612b7a613a5b565b8152602001908152602001600020905060005b8154811015610a6357836001600160a01b0316828281548110612bb257612bb2614329565b6000918252602090912001546001600160a01b031603612c815781548290612bdc906001906145f2565b81548110612bec57612bec614329565b9060005260206000200160009054906101000a90046001600160a01b0316828281548110612c1c57612c1c614329565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555081805480612c5a57612c5a614605565b600082815260209020810160001990810180546001600160a01b0319169055019055610a63565b600101612b8d565b6000807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610781565b612cba613408565b565b612cc4613408565b612cba61342d565b80356000908155612ce0602083018361461b565b6001830191612cf09190836146a1565b50612d0160608301604084016143d9565b6002820180546001600160a01b0319166001600160a01b0392909216919091179055612d336080830160608401614760565b60028201805461ffff92909216600160a01b0261ffff60a01b19909216919091179055612d63608083018361461b565b6003830191612d739190836146a1565b50612d8160a083018361461b565b6004830191612d919190836146a1565b50610120820160068201612da5828261479d565b5050610180820160078201612dba8282614828565b50612dcd905060e0830160c08401614880565b60058201805460ff1916911515919091179055612dee6101c083018361489d565b60088201612dfc82826144d0565b50612e0d90506101e083018361461b565b6010830191612e1d9190836146a1565b5061020082013560118201556102208201356012820155612e466101208301610100840161393c565b60058201805462ff0000191662010000836003811115612e6857612e68613a5b565b02179055506102408201356013820155612e89610100830160e08401614880565b600590910180549115156101000261ff001990921691909117905550565b6000612eb1613435565b805490915060011901612ed757604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b612ee5610e2c565b15612cba5760405163d93c066560e01b815260040160405180910390fd5b3360009081526014602052604081208190818b6003811115612f2757612f27613a5b565b6003811115612f3857612f38613a5b565b81526020810191909152604001600020546001600160a01b031614612f9f5760405162461bcd60e51b815260206004820181905260248201527f416c7265616479207265676973746572656420666f72207468697320726f6c656044820152606401610800565b84612fdd5760405162461bcd60e51b815260206004820152600e60248201526d115b5c1d1e48195b991c1bda5b9d60921b6044820152606401610800565b6000612fe88a613354565b905080601660008c600381111561300157613001613a5b565b600381111561301257613012613a5b565b81526020810191909152604001600020541061306b5760405162461bcd60e51b815260206004820152601860248201527713585e081c185c9d1a58da5c185b9d1cc81c995858da195960421b6044820152606401610800565b60006130768b6132e8565b905080156130ce57808410156130c35760405162461bcd60e51b8152602060048201526012602482015271496e73756666696369656e74207374616b6560701b6044820152606401610800565b6130ce338c86613459565b336000908152601460205260408120818d60038111156130f0576130f0613a5b565b600381111561310157613101613a5b565b8152602081019190915260400160002080549091508c90829060ff60a01b1916600160a01b83600381111561313857613138613a5b565b021790555060055460ff16156131a357805460ff60a81b19168155601560008d600381111561316957613169613a5b565b600381111561317a5761317a613a5b565b81526020019081526020016000206000815480929190613199906142e7565b91905055506131b5565b805460ff60a81b1916600160a91b1781555b805461ffff60a01b163317607d60b21b17815560006002820155600160801b6001600160801b034216908102176001820155600381016131f6898b836146a1565b50600481016132068b8d836146a1565b50600581016132168789836146a1565b50601660008d600381111561322d5761322d613a5b565b600381111561323e5761323e613a5b565b8152602080820192909252604001600090812080546001810182559082529190200180546001600160a01b031916331790558b600381111561328257613282613a5b565b336001600160a01b03167fe3564c470dc79d88d220de1a64bc335ec0a928450ce631293691bba84dd9a6118d8d6040516132bd9291906143aa565b60405180910390a35060019b9a5050505050505050505050565b60006132e1613435565b6001905550565b6000808260038111156132fd576132fd613a5b565b0361330a57505060085490565b600182600381111561331e5761331e613a5b565b0361332b57505060095490565b600282600381111561333f5761333f613a5b565b0361334c575050600a5490565b506000919050565b60008082600381111561336957613369613a5b565b03613376575050600b5490565b600182600381111561338a5761338a613a5b565b03613397575050600c5490565b60028260038111156133ab576133ab613a5b565b0361334c575050600d5490565b6133c28282611b41565b61083657808260405163e2517d3f60e01b81526004016108009291906148bd565b6133eb610e2c565b612cba57604051638dfc202b60e01b815260040160405180910390fd5b613410613714565b612cba57604051631afcd79f60e31b815260040160405180910390fd5b6132d7613408565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0090565b6017546001600160a01b03166134815760405162461bcd60e51b815260040161080090614379565b8060000361348e57505050565b6000613498611d9c565b90506001600160a01b03811661355e57813410156134ee5760405162461bcd60e51b815260206004820152601360248201527208aa89040c2dadeeadce840dad2e6dac2e8c6d606b1b6044820152606401610800565b601754604051630b51be0f60e21b81526001600160a01b0390911690632d46f83c908490613527908890600090899085906004016148d6565b6000604051808303818588803b15801561354057600080fd5b505af1158015613554573d6000803e3d6000fd5b50505050506136b9565b6040516323b872dd60e01b81526001600160a01b038581166004830152306024830152604482018490528216906323b872dd906064016020604051808303816000875af11580156135b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135d79190614909565b5060175460405163095ea7b360e01b81526001600160a01b038381169263095ea7b39261360c929091169086906004016148bd565b6020604051808303816000875af115801561362b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061364f9190614909565b50601754604051630b51be0f60e21b81526001600160a01b0390911690632d46f83c906136869087908590889088906004016148d6565b600060405180830381600087803b1580156136a057600080fd5b505af11580156136b4573d6000803e3d6000fd5b505050505b8260038111156136cb576136cb613a5b565b846001600160a01b03167f1304e75cb7d7b5caf45a5614e9da12e06f2e5ce43fe410b18986ef8347bbd9ee8460405161370691815260200190565b60405180910390a350505050565b600061371e612c89565b54600160401b900460ff16919050565b604080518082019091526000808252602082015290565b6040805161014081019091526000808252602082019081526020016000815260200160006001600160501b0316815260200160006001600160801b0316815260200160006001600160801b03168152602001600081526020016060815260200160608152602001606081525090565b5080546137c09061433f565b6000825580601f106137d0575050565b601f01602090049060005260206000209081019061294b91906138e5565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b604080516102008101825260008082526060602083018190529282018190528282018190526080820183905260a082019290925260c0810182905260e081018290529061010082019081526020016138896138c5565b815260200161389661372e565b81526020016138a36137ee565b8152602001606081526020016000815260200160008152602001600081525090565b604080516060810182526000808252602082018190529181019190915290565b5b808211156138fa57600081556001016138e6565b5090565b60006020828403121561391057600080fd5b81356001600160e01b03198116811461392857600080fd5b9392505050565b6004811061294b57600080fd5b60006020828403121561394e57600080fd5b81356139288161392f565b60006020828403121561396b57600080fd5b5035919050565b805163ffffffff908116835260209182015116910152565b604081016107818284613972565b6001600160a01b038116811461294b57600080fd5b600080604083850312156139c057600080fd5b8235915060208301356139d281613998565b809150509250929050565b600080604083850312156139f057600080fd5b82356139fb81613998565b915060208301356139d28161392f565b602080825282518282018190526000918401906040840190835b81811015613a43578351835260209384019390920191600101613a25565b509095945050505050565b6001600160a01b03169052565b634e487b7160e01b600052602160045260246000fd5b6004811061294b5761294b613a5b565b613a8a81613a71565b9052565b60038110613a8a57613a8a613a5b565b6001600160801b03169052565b6000815180845260005b81811015613ad157602081850181015186830182015201613ab5565b506000602082860101526020601f19601f83011685010191505092915050565b613afc828251613a4e565b60006020820151613b106020850182613a81565b506040820151613b236040850182613a8e565b506060820151613b3e60608501826001600160501b03169052565b506080820151613b516080850182613a9e565b5060a0820151613b6460a0850182613a9e565b5060c082015160c084015260e082015161014060e0850152613b8a610140850182613aab565b9050610100830151848203610100860152613ba58282613aab565b915050610120830151848203610120860152613bc18282613aab565b95945050505050565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b82811015613c2357603f19878603018452613c0e858351613af1565b94506020938401939190910190600101613bf2565b50929695505050505050565b60008083601f840112613c4157600080fd5b5081356001600160401b03811115613c5857600080fd5b602083019150836020828501011115613c7057600080fd5b9250929050565b60008060008060608587031215613c8d57600080fd5b8435613c9881613998565b93506020850135613ca88161392f565b925060408501356001600160401b03811115613cc357600080fd5b613ccf87828801613c2f565b95989497509550505050565b600080600060608486031215613cf057600080fd5b83356001600160401b03811115613d0657600080fd5b84016102608187031215613d1957600080fd5b92506020840135613d2981613998565b91506040840135613d3981613998565b809150509250925092565b6020815260006139286020830184613af1565b60008060008060008060008060a0898b031215613d7357600080fd5b8835613d7e8161392f565b97506020890135965060408901356001600160401b03811115613da057600080fd5b613dac8b828c01613c2f565b90975095505060608901356001600160401b03811115613dcb57600080fd5b613dd78b828c01613c2f565b90955093505060808901356001600160401b03811115613df657600080fd5b613e028b828c01613c2f565b999c989b5096995094979396929594505050565b60008060008060008060006080888a031215613e3157600080fd5b8735613e3c8161392f565b965060208801356001600160401b03811115613e5757600080fd5b613e638a828b01613c2f565b90975095505060408801356001600160401b03811115613e8257600080fd5b613e8e8a828b01613c2f565b90955093505060608801356001600160401b03811115613ead57600080fd5b613eb98a828b01613c2f565b989b979a50959850939692959293505050565b600080600060408486031215613ee157600080fd5b8335613eec8161392f565b925060208401356001600160401b03811115613f0757600080fd5b613f1386828701613c2f565b9497909650939450505050565b6001600160a01b0391909116815260200190565b60008060408385031215613f4757600080fd5b8235613f528161392f565b946020939093013593505050565b6001600160a01b038b168152613f758a613a71565b896020820152613f88604082018a613a8e565b6001600160501b03881660608201526001600160801b038781166080830152861660a082015260c0810185905261014060e08201819052600090613fce90830186613aab565b828103610100840152613fe18186613aab565b9050828103610120840152613ff68185613aab565b9d9c50505050505050505050505050565b60006101008301825184526020830151602085015260408301516040850152606083015160608501526080830151608085015260a083015160a085015260c083015160c085015260e083015161010060e086015281815180845261012087019150602083019350600092505b808310156113545783518252602082019150602084019350600183019250614073565b6020815260006139286020830184614007565b80516001600160801b031682526020808201516001600160401b039081169184019190915260409182015116910152565b6020815281516020820152600060208301516102606040840152614102610280840182613aab565b905060408401516141166060850182613a4e565b50606084015161ffff81166080850152506080840151838203601f190160a08501526141428282613aab565b91505060a0840151601f198483030160c08501526141608282613aab565b91505060c084015161417660e085018215159052565b5060e08401518015156101008501525061010084015161419a610120850182613a81565b506101208401516141af6101408501826140a9565b506101408401516141c46101a0850182613972565b50610160840151838203601f19016101e08501526141e28282614007565b915050610180840151601f19848303016102008501526142028282613aab565b9150506101a08401516102208401526101c08401516102408401526101e08401516102608401528091505092915050565b6060810161078182846140a9565b60006020828403121561425357600080fd5b81356001600160401b0381111561426957600080fd5b8201610100818503121561392857600080fd5b6020808252601690820152755061727469636970616e74206e6f742065786973747360501b604082015260600190565b6020808252600b908201526a4e6f742070656e64696e6760a81b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000600182016142f9576142f96142d1565b5060010190565b80820180821115610781576107816142d1565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600181811c9082168061435357607f821691505b60208210810361437357634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526017908201527614dd185ada5b99c81b585b9859d95c881b9bdd081cd95d604a1b604082015260600190565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b6000602082840312156143eb57600080fd5b813561392881613998565b600081614405576144056142d1565b506000190190565b602081016107818284613a8e565b60006020828403121561442d57600080fd5b815161392881613998565b5b818110156108365760008155600101614439565b6001600160401b0383111561446457614464614313565b600160401b83111561447857614478614313565b80548382558084101561449e5781600052602060002061449c828201868301614438565b505b508181600052602060002060005b858110156144c8578235828201556020909201916001016144ac565b505050505050565b813581556020820135600182015560408201356002820155606082013560038201556080820135600482015560a0820135600582015560c0820135600682015560008060e084013536859003601e1901811261452a578283fd5b8401803591506001600160401b03821115614543578283fd5b6020019150600581901b360382131561455b57600080fd5b610a6381836007860161444d565b6001600160a01b038481168252831660208201526060810161458a83613a71565b826040830152949350505050565b600060608284031280156145ab57600080fd5b50604051606081016001600160401b03811182821017156145ce576145ce614313565b60409081528351825260208085015190830152928301519281019290925250919050565b81810381811115610781576107816142d1565b634e487b7160e01b600052603160045260246000fd5b6000808335601e1984360301811261463257600080fd5b8301803591506001600160401b0382111561464c57600080fd5b602001915036819003821315613c7057600080fd5b601f821115610c0a57806000526020600020601f840160051c810160208510156146885750805b61469a601f850160051c830182614438565b5050505050565b6001600160401b038311156146b8576146b8614313565b6146cc836146c6835461433f565b83614661565b6000601f84116001811461470057600085156146e85750838201355b600019600387901b1c1916600186901b17835561469a565b600083815260209020601f19861690835b828110156147315786850135825560209485019460019092019101614711565b508682101561474e5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b60006020828403121561477257600080fd5b813561ffff8116811461392857600080fd5b600081356001600160401b038116811461078157600080fd5b81356001600160801b0381168082146147b557600080fd5b82546001600160801b03198116821784559150600160801b600160c01b036147df60208601614784565b60801b166001600160c01b031983811683178217855561480160408701614784565b60c01b168282171784555050505050565b6000813563ffffffff8116811461078157600080fd5b63ffffffff61483683614812565b1681548163ffffffff19821617835563ffffffff60201b61485960208601614812565b60201b168260018060401b031983161717835550505050565b801515811461294b57600080fd5b60006020828403121561489257600080fd5b813561392881614872565b6000823560fe198336030181126148b357600080fd5b9190910192915050565b6001600160a01b03929092168252602082015260400190565b6001600160a01b03858116825284166020820152608081016148f784613a71565b60408201939093526060015292915050565b60006020828403121561491b57600080fd5b81516139288161487256fe0a4b355e037ef9ee17aaae57aea8f4bc674d63b8b77b33523e3d0b488c02c215a49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775a2646970667358221220106deda742673736263f12b1064bd56780ce5792ae55869bdfdd015700f762ff64736f6c634300081b0033",
}

// SubnetABI is the input ABI used to generate the binding from.
// Deprecated: Use SubnetMetaData.ABI instead.
var SubnetABI = SubnetMetaData.ABI

// SubnetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SubnetMetaData.Bin instead.
var SubnetBin = SubnetMetaData.Bin

// DeploySubnet deploys a new Ethereum contract, binding an instance of Subnet to it.
func DeploySubnet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Subnet, error) {
	parsed, err := SubnetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SubnetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Subnet{SubnetCaller: SubnetCaller{contract: contract}, SubnetTransactor: SubnetTransactor{contract: contract}, SubnetFilterer: SubnetFilterer{contract: contract}}, nil
}

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

// GetParticipantInfo is a free data retrieval call binding the contract method 0x735d23e8.
//
// Solidity: function getParticipantInfo(address participant_addr, uint8 participant_type) view returns((address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string))
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
// Solidity: function getParticipantInfo(address participant_addr, uint8 participant_type) view returns((address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string))
func (_Subnet *SubnetSession) GetParticipantInfo(participant_addr common.Address, participant_type uint8) (DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.GetParticipantInfo(&_Subnet.CallOpts, participant_addr, participant_type)
}

// GetParticipantInfo is a free data retrieval call binding the contract method 0x735d23e8.
//
// Solidity: function getParticipantInfo(address participant_addr, uint8 participant_type) view returns((address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string))
func (_Subnet *SubnetCallerSession) GetParticipantInfo(participant_addr common.Address, participant_type uint8) (DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.GetParticipantInfo(&_Subnet.CallOpts, participant_addr, participant_type)
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
// Solidity: function listActiveParticipants(uint8 participant_type) view returns((address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[])
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
// Solidity: function listActiveParticipants(uint8 participant_type) view returns((address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[])
func (_Subnet *SubnetSession) ListActiveParticipants(participant_type uint8) ([]DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.ListActiveParticipants(&_Subnet.CallOpts, participant_type)
}

// ListActiveParticipants is a free data retrieval call binding the contract method 0x5caecd9c.
//
// Solidity: function listActiveParticipants(uint8 participant_type) view returns((address,uint8,uint8,uint80,uint128,uint128,uint256,string,string,string)[])
func (_Subnet *SubnetCallerSession) ListActiveParticipants(participant_type uint8) ([]DataStructuresParticipantInfo, error) {
	return _Subnet.Contract.ListActiveParticipants(&_Subnet.CallOpts, participant_type)
}

// ParticipantListByType is a free data retrieval call binding the contract method 0xac6c2a45.
//
// Solidity: function participantListByType(uint8 , uint256 ) view returns(address)
func (_Subnet *SubnetCaller) ParticipantListByType(opts *bind.CallOpts, arg0 uint8, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "participantListByType", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParticipantListByType is a free data retrieval call binding the contract method 0xac6c2a45.
//
// Solidity: function participantListByType(uint8 , uint256 ) view returns(address)
func (_Subnet *SubnetSession) ParticipantListByType(arg0 uint8, arg1 *big.Int) (common.Address, error) {
	return _Subnet.Contract.ParticipantListByType(&_Subnet.CallOpts, arg0, arg1)
}

// ParticipantListByType is a free data retrieval call binding the contract method 0xac6c2a45.
//
// Solidity: function participantListByType(uint8 , uint256 ) view returns(address)
func (_Subnet *SubnetCallerSession) ParticipantListByType(arg0 uint8, arg1 *big.Int) (common.Address, error) {
	return _Subnet.Contract.ParticipantListByType(&_Subnet.CallOpts, arg0, arg1)
}

// Participants is a free data retrieval call binding the contract method 0xca8cf966.
//
// Solidity: function participants(address , uint8 ) view returns(address owner, uint8 participant_type, uint8 status, uint80 reputation_score, uint128 registered_at, uint128 last_active, uint256 reserved, string endpoint, string domain, string metadata_uri)
func (_Subnet *SubnetCaller) Participants(opts *bind.CallOpts, arg0 common.Address, arg1 uint8) (struct {
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
}, error) {
	var out []interface{}
	err := _Subnet.contract.Call(opts, &out, "participants", arg0, arg1)

	outstruct := new(struct {
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
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ParticipantType = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ReputationScore = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RegisteredAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LastActive = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Reserved = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Endpoint = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.Domain = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.MetadataUri = *abi.ConvertType(out[9], new(string)).(*string)

	return *outstruct, err

}

// Participants is a free data retrieval call binding the contract method 0xca8cf966.
//
// Solidity: function participants(address , uint8 ) view returns(address owner, uint8 participant_type, uint8 status, uint80 reputation_score, uint128 registered_at, uint128 last_active, uint256 reserved, string endpoint, string domain, string metadata_uri)
func (_Subnet *SubnetSession) Participants(arg0 common.Address, arg1 uint8) (struct {
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
}, error) {
	return _Subnet.Contract.Participants(&_Subnet.CallOpts, arg0, arg1)
}

// Participants is a free data retrieval call binding the contract method 0xca8cf966.
//
// Solidity: function participants(address , uint8 ) view returns(address owner, uint8 participant_type, uint8 status, uint80 reputation_score, uint128 registered_at, uint128 last_active, uint256 reserved, string endpoint, string domain, string metadata_uri)
func (_Subnet *SubnetCallerSession) Participants(arg0 common.Address, arg1 uint8) (struct {
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

// Initialize is a paid mutator transaction binding the contract method 0x708b73a2.
//
// Solidity: function initialize((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) _subnetInfo, address _factory, address stakingManager_) returns()
func (_Subnet *SubnetTransactor) Initialize(opts *bind.TransactOpts, _subnetInfo DataStructuresSubnetInfo, _factory common.Address, stakingManager_ common.Address) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "initialize", _subnetInfo, _factory, stakingManager_)
}

// Initialize is a paid mutator transaction binding the contract method 0x708b73a2.
//
// Solidity: function initialize((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) _subnetInfo, address _factory, address stakingManager_) returns()
func (_Subnet *SubnetSession) Initialize(_subnetInfo DataStructuresSubnetInfo, _factory common.Address, stakingManager_ common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.Initialize(&_Subnet.TransactOpts, _subnetInfo, _factory, stakingManager_)
}

// Initialize is a paid mutator transaction binding the contract method 0x708b73a2.
//
// Solidity: function initialize((bytes32,string,address,uint16,string,string,bool,bool,uint8,(uint128,uint64,uint64),(uint32,uint32),(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256[]),string,uint256,uint256,uint256) _subnetInfo, address _factory, address stakingManager_) returns()
func (_Subnet *SubnetTransactorSession) Initialize(_subnetInfo DataStructuresSubnetInfo, _factory common.Address, stakingManager_ common.Address) (*types.Transaction, error) {
	return _Subnet.Contract.Initialize(&_Subnet.TransactOpts, _subnetInfo, _factory, stakingManager_)
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

// RegisterParticipant is a paid mutator transaction binding the contract method 0x8941382a.
//
// Solidity: function registerParticipant(uint8 participant_type, string domain, string endpoint, string metadata_uri) payable returns(bool)
func (_Subnet *SubnetTransactor) RegisterParticipant(opts *bind.TransactOpts, participant_type uint8, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "registerParticipant", participant_type, domain, endpoint, metadata_uri)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x8941382a.
//
// Solidity: function registerParticipant(uint8 participant_type, string domain, string endpoint, string metadata_uri) payable returns(bool)
func (_Subnet *SubnetSession) RegisterParticipant(participant_type uint8, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipant(&_Subnet.TransactOpts, participant_type, domain, endpoint, metadata_uri)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x8941382a.
//
// Solidity: function registerParticipant(uint8 participant_type, string domain, string endpoint, string metadata_uri) payable returns(bool)
func (_Subnet *SubnetTransactorSession) RegisterParticipant(participant_type uint8, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipant(&_Subnet.TransactOpts, participant_type, domain, endpoint, metadata_uri)
}

// RegisterParticipantERC20 is a paid mutator transaction binding the contract method 0x73a6d7c5.
//
// Solidity: function registerParticipantERC20(uint8 participant_type, uint256 amount, string domain, string endpoint, string metadata_uri) returns(bool)
func (_Subnet *SubnetTransactor) RegisterParticipantERC20(opts *bind.TransactOpts, participant_type uint8, amount *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "registerParticipantERC20", participant_type, amount, domain, endpoint, metadata_uri)
}

// RegisterParticipantERC20 is a paid mutator transaction binding the contract method 0x73a6d7c5.
//
// Solidity: function registerParticipantERC20(uint8 participant_type, uint256 amount, string domain, string endpoint, string metadata_uri) returns(bool)
func (_Subnet *SubnetSession) RegisterParticipantERC20(participant_type uint8, amount *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipantERC20(&_Subnet.TransactOpts, participant_type, amount, domain, endpoint, metadata_uri)
}

// RegisterParticipantERC20 is a paid mutator transaction binding the contract method 0x73a6d7c5.
//
// Solidity: function registerParticipantERC20(uint8 participant_type, uint256 amount, string domain, string endpoint, string metadata_uri) returns(bool)
func (_Subnet *SubnetTransactorSession) RegisterParticipantERC20(participant_type uint8, amount *big.Int, domain string, endpoint string, metadata_uri string) (*types.Transaction, error) {
	return _Subnet.Contract.RegisterParticipantERC20(&_Subnet.TransactOpts, participant_type, amount, domain, endpoint, metadata_uri)
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

// SuspendParticipant is a paid mutator transaction binding the contract method 0x96611afd.
//
// Solidity: function suspendParticipant(uint8 participant_type, string ) returns()
func (_Subnet *SubnetTransactor) SuspendParticipant(opts *bind.TransactOpts, participant_type uint8, arg1 string) (*types.Transaction, error) {
	return _Subnet.contract.Transact(opts, "suspendParticipant", participant_type, arg1)
}

// SuspendParticipant is a paid mutator transaction binding the contract method 0x96611afd.
//
// Solidity: function suspendParticipant(uint8 participant_type, string ) returns()
func (_Subnet *SubnetSession) SuspendParticipant(participant_type uint8, arg1 string) (*types.Transaction, error) {
	return _Subnet.Contract.SuspendParticipant(&_Subnet.TransactOpts, participant_type, arg1)
}

// SuspendParticipant is a paid mutator transaction binding the contract method 0x96611afd.
//
// Solidity: function suspendParticipant(uint8 participant_type, string ) returns()
func (_Subnet *SubnetTransactorSession) SuspendParticipant(participant_type uint8, arg1 string) (*types.Transaction, error) {
	return _Subnet.Contract.SuspendParticipant(&_Subnet.TransactOpts, participant_type, arg1)
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
