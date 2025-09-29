// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingmanager

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

// DataStructuresStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type DataStructuresStakeInfo struct {
	StakedAmount         *big.Int
	UnstakeRequestTime   *big.Int
	UnstakeRequestAmount *big.Int
}

// StakingManagerMetaData contains all meta data concerning the StakingManager contract.
var StakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"availableAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentTime\",\"type\":\"uint256\"}],\"name\":\"InitialLockNotPassed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InsufficientWithdrawable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"InvalidSlashAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"StakeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedNativeValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"UnsupportedToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AuthorizedCallerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"InitialLockPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"StakeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotal\",\"type\":\"uint256\"}],\"name\":\"StakeToppedUp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"UnlockPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"name\":\"UnstakeRequested\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOVERNANCE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SLASHER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositStakeFor\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"getStakeInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stakedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakeRequestAmount\",\"type\":\"uint256\"}],\"internalType\":\"structDataStructures.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUnlockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"requestAmount\",\"type\":\"uint256\"}],\"name\":\"requestUnstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"setGovernanceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"}],\"name\":\"setSlasherRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setUnlockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumDataStructures.ParticipantType\",\"name\":\"role\",\"type\":\"uint8\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingManagerMetaData.ABI instead.
var StakingManagerABI = StakingManagerMetaData.ABI

// StakingManager is an auto generated Go binding around an Ethereum contract.
type StakingManager struct {
	StakingManagerCaller     // Read-only binding to the contract
	StakingManagerTransactor // Write-only binding to the contract
	StakingManagerFilterer   // Log filterer for contract events
}

// StakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingManagerSession struct {
	Contract     *StakingManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingManagerCallerSession struct {
	Contract *StakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingManagerTransactorSession struct {
	Contract     *StakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingManagerRaw struct {
	Contract *StakingManager // Generic contract binding to access the raw methods on
}

// StakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingManagerCallerRaw struct {
	Contract *StakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingManagerTransactorRaw struct {
	Contract *StakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingManager creates a new instance of StakingManager, bound to a specific deployed contract.
func NewStakingManager(address common.Address, backend bind.ContractBackend) (*StakingManager, error) {
	contract, err := bindStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingManager{StakingManagerCaller: StakingManagerCaller{contract: contract}, StakingManagerTransactor: StakingManagerTransactor{contract: contract}, StakingManagerFilterer: StakingManagerFilterer{contract: contract}}, nil
}

// NewStakingManagerCaller creates a new read-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*StakingManagerCaller, error) {
	contract, err := bindStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerCaller{contract: contract}, nil
}

// NewStakingManagerTransactor creates a new write-only instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingManagerTransactor, error) {
	contract, err := bindStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingManagerTransactor{contract: contract}, nil
}

// NewStakingManagerFilterer creates a new log filterer instance of StakingManager, bound to a specific deployed contract.
func NewStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingManagerFilterer, error) {
	contract, err := bindStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingManagerFilterer{contract: contract}, nil
}

// bindStakingManager binds a generic wrapper to an already deployed contract.
func bindStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.StakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.StakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingManager *StakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingManager *StakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingManager *StakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingManager.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.DEFAULTADMINROLE(&_StakingManager.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) GOVERNANCEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "GOVERNANCE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _StakingManager.Contract.GOVERNANCEROLE(&_StakingManager.CallOpts)
}

// GOVERNANCEROLE is a free data retrieval call binding the contract method 0xf36c8f5c.
//
// Solidity: function GOVERNANCE_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) GOVERNANCEROLE() ([32]byte, error) {
	return _StakingManager.Contract.GOVERNANCEROLE(&_StakingManager.CallOpts)
}

// SLASHERROLE is a free data retrieval call binding the contract method 0x5095af64.
//
// Solidity: function SLASHER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) SLASHERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "SLASHER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SLASHERROLE is a free data retrieval call binding the contract method 0x5095af64.
//
// Solidity: function SLASHER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) SLASHERROLE() ([32]byte, error) {
	return _StakingManager.Contract.SLASHERROLE(&_StakingManager.CallOpts)
}

// SLASHERROLE is a free data retrieval call binding the contract method 0x5095af64.
//
// Solidity: function SLASHER_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) SLASHERROLE() ([32]byte, error) {
	return _StakingManager.Contract.SLASHERROLE(&_StakingManager.CallOpts)
}

// STAKINGADMINROLE is a free data retrieval call binding the contract method 0x04f67aa1.
//
// Solidity: function STAKING_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCaller) STAKINGADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "STAKING_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGADMINROLE is a free data retrieval call binding the contract method 0x04f67aa1.
//
// Solidity: function STAKING_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerSession) STAKINGADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGADMINROLE(&_StakingManager.CallOpts)
}

// STAKINGADMINROLE is a free data retrieval call binding the contract method 0x04f67aa1.
//
// Solidity: function STAKING_ADMIN_ROLE() view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) STAKINGADMINROLE() ([32]byte, error) {
	return _StakingManager.Contract.STAKINGADMINROLE(&_StakingManager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StakingManager *StakingManagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StakingManager.Contract.GetRoleAdmin(&_StakingManager.CallOpts, role)
}

// GetStakeInfo is a free data retrieval call binding the contract method 0x442a3776.
//
// Solidity: function getStakeInfo(address user, address token, uint8 role) view returns((uint256,uint256,uint256))
func (_StakingManager *StakingManagerCaller) GetStakeInfo(opts *bind.CallOpts, user common.Address, token common.Address, role uint8) (DataStructuresStakeInfo, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakeInfo", user, token, role)

	if err != nil {
		return *new(DataStructuresStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DataStructuresStakeInfo)).(*DataStructuresStakeInfo)

	return out0, err

}

// GetStakeInfo is a free data retrieval call binding the contract method 0x442a3776.
//
// Solidity: function getStakeInfo(address user, address token, uint8 role) view returns((uint256,uint256,uint256))
func (_StakingManager *StakingManagerSession) GetStakeInfo(user common.Address, token common.Address, role uint8) (DataStructuresStakeInfo, error) {
	return _StakingManager.Contract.GetStakeInfo(&_StakingManager.CallOpts, user, token, role)
}

// GetStakeInfo is a free data retrieval call binding the contract method 0x442a3776.
//
// Solidity: function getStakeInfo(address user, address token, uint8 role) view returns((uint256,uint256,uint256))
func (_StakingManager *StakingManagerCallerSession) GetStakeInfo(user common.Address, token common.Address, role uint8) (DataStructuresStakeInfo, error) {
	return _StakingManager.Contract.GetStakeInfo(&_StakingManager.CallOpts, user, token, role)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_StakingManager *StakingManagerCaller) GetStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_StakingManager *StakingManagerSession) GetStakingToken() (common.Address, error) {
	return _StakingManager.Contract.GetStakingToken(&_StakingManager.CallOpts)
}

// GetStakingToken is a free data retrieval call binding the contract method 0x9f9106d1.
//
// Solidity: function getStakingToken() view returns(address)
func (_StakingManager *StakingManagerCallerSession) GetStakingToken() (common.Address, error) {
	return _StakingManager.Contract.GetStakingToken(&_StakingManager.CallOpts)
}

// GetUnlockPeriod is a free data retrieval call binding the contract method 0x8e43509c.
//
// Solidity: function getUnlockPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCaller) GetUnlockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "getUnlockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnlockPeriod is a free data retrieval call binding the contract method 0x8e43509c.
//
// Solidity: function getUnlockPeriod() view returns(uint256)
func (_StakingManager *StakingManagerSession) GetUnlockPeriod() (*big.Int, error) {
	return _StakingManager.Contract.GetUnlockPeriod(&_StakingManager.CallOpts)
}

// GetUnlockPeriod is a free data retrieval call binding the contract method 0x8e43509c.
//
// Solidity: function getUnlockPeriod() view returns(uint256)
func (_StakingManager *StakingManagerCallerSession) GetUnlockPeriod() (*big.Int, error) {
	return _StakingManager.Contract.GetUnlockPeriod(&_StakingManager.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StakingManager.Contract.HasRole(&_StakingManager.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerSession) Paused() (bool, error) {
	return _StakingManager.Contract.Paused(&_StakingManager.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StakingManager *StakingManagerCallerSession) Paused() (bool, error) {
	return _StakingManager.Contract.Paused(&_StakingManager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakingManager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakingManager *StakingManagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakingManager.Contract.SupportsInterface(&_StakingManager.CallOpts, interfaceId)
}

// DepositStakeFor is a paid mutator transaction binding the contract method 0x2d46f83c.
//
// Solidity: function depositStakeFor(address user, address token, uint8 role, uint256 amount) payable returns()
func (_StakingManager *StakingManagerTransactor) DepositStakeFor(opts *bind.TransactOpts, user common.Address, token common.Address, role uint8, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "depositStakeFor", user, token, role, amount)
}

// DepositStakeFor is a paid mutator transaction binding the contract method 0x2d46f83c.
//
// Solidity: function depositStakeFor(address user, address token, uint8 role, uint256 amount) payable returns()
func (_StakingManager *StakingManagerSession) DepositStakeFor(user common.Address, token common.Address, role uint8, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DepositStakeFor(&_StakingManager.TransactOpts, user, token, role, amount)
}

// DepositStakeFor is a paid mutator transaction binding the contract method 0x2d46f83c.
//
// Solidity: function depositStakeFor(address user, address token, uint8 role, uint256 amount) payable returns()
func (_StakingManager *StakingManagerTransactorSession) DepositStakeFor(user common.Address, token common.Address, role uint8, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.DepositStakeFor(&_StakingManager.TransactOpts, user, token, role, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address recipient, address token, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactor) EmergencyWithdraw(opts *bind.TransactOpts, recipient common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "emergencyWithdraw", recipient, token, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address recipient, address token, uint256 amount) returns()
func (_StakingManager *StakingManagerSession) EmergencyWithdraw(recipient common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.EmergencyWithdraw(&_StakingManager.TransactOpts, recipient, token, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address recipient, address token, uint256 amount) returns()
func (_StakingManager *StakingManagerTransactorSession) EmergencyWithdraw(recipient common.Address, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.EmergencyWithdraw(&_StakingManager.TransactOpts, recipient, token, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.GrantRole(&_StakingManager.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address governance) returns()
func (_StakingManager *StakingManagerTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, governance common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "initialize", admin, governance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address governance) returns()
func (_StakingManager *StakingManagerSession) Initialize(admin common.Address, governance common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, admin, governance)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address governance) returns()
func (_StakingManager *StakingManagerTransactorSession) Initialize(admin common.Address, governance common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.Initialize(&_StakingManager.TransactOpts, admin, governance)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingManager *StakingManagerTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingManager *StakingManagerSession) Pause() (*types.Transaction, error) {
	return _StakingManager.Contract.Pause(&_StakingManager.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StakingManager *StakingManagerTransactorSession) Pause() (*types.Transaction, error) {
	return _StakingManager.Contract.Pause(&_StakingManager.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StakingManager *StakingManagerTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RenounceRole(&_StakingManager.TransactOpts, role, callerConfirmation)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0x40147a5b.
//
// Solidity: function requestUnstake(address token, uint8 role, uint256 requestAmount) returns()
func (_StakingManager *StakingManagerTransactor) RequestUnstake(opts *bind.TransactOpts, token common.Address, role uint8, requestAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "requestUnstake", token, role, requestAmount)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0x40147a5b.
//
// Solidity: function requestUnstake(address token, uint8 role, uint256 requestAmount) returns()
func (_StakingManager *StakingManagerSession) RequestUnstake(token common.Address, role uint8, requestAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnstake(&_StakingManager.TransactOpts, token, role, requestAmount)
}

// RequestUnstake is a paid mutator transaction binding the contract method 0x40147a5b.
//
// Solidity: function requestUnstake(address token, uint8 role, uint256 requestAmount) returns()
func (_StakingManager *StakingManagerTransactorSession) RequestUnstake(token common.Address, role uint8, requestAmount *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.RequestUnstake(&_StakingManager.TransactOpts, token, role, requestAmount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StakingManager *StakingManagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.RevokeRole(&_StakingManager.TransactOpts, role, account)
}

// SetGovernanceRole is a paid mutator transaction binding the contract method 0xe1ba6889.
//
// Solidity: function setGovernanceRole(address governance) returns()
func (_StakingManager *StakingManagerTransactor) SetGovernanceRole(opts *bind.TransactOpts, governance common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setGovernanceRole", governance)
}

// SetGovernanceRole is a paid mutator transaction binding the contract method 0xe1ba6889.
//
// Solidity: function setGovernanceRole(address governance) returns()
func (_StakingManager *StakingManagerSession) SetGovernanceRole(governance common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetGovernanceRole(&_StakingManager.TransactOpts, governance)
}

// SetGovernanceRole is a paid mutator transaction binding the contract method 0xe1ba6889.
//
// Solidity: function setGovernanceRole(address governance) returns()
func (_StakingManager *StakingManagerTransactorSession) SetGovernanceRole(governance common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetGovernanceRole(&_StakingManager.TransactOpts, governance)
}

// SetSlasherRole is a paid mutator transaction binding the contract method 0x44009962.
//
// Solidity: function setSlasherRole(address slasher) returns()
func (_StakingManager *StakingManagerTransactor) SetSlasherRole(opts *bind.TransactOpts, slasher common.Address) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setSlasherRole", slasher)
}

// SetSlasherRole is a paid mutator transaction binding the contract method 0x44009962.
//
// Solidity: function setSlasherRole(address slasher) returns()
func (_StakingManager *StakingManagerSession) SetSlasherRole(slasher common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlasherRole(&_StakingManager.TransactOpts, slasher)
}

// SetSlasherRole is a paid mutator transaction binding the contract method 0x44009962.
//
// Solidity: function setSlasherRole(address slasher) returns()
func (_StakingManager *StakingManagerTransactorSession) SetSlasherRole(slasher common.Address) (*types.Transaction, error) {
	return _StakingManager.Contract.SetSlasherRole(&_StakingManager.TransactOpts, slasher)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0x3d0ddf84.
//
// Solidity: function setUnlockPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactor) SetUnlockPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "setUnlockPeriod", period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0x3d0ddf84.
//
// Solidity: function setUnlockPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerSession) SetUnlockPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetUnlockPeriod(&_StakingManager.TransactOpts, period)
}

// SetUnlockPeriod is a paid mutator transaction binding the contract method 0x3d0ddf84.
//
// Solidity: function setUnlockPeriod(uint256 period) returns()
func (_StakingManager *StakingManagerTransactorSession) SetUnlockPeriod(period *big.Int) (*types.Transaction, error) {
	return _StakingManager.Contract.SetUnlockPeriod(&_StakingManager.TransactOpts, period)
}

// Slash is a paid mutator transaction binding the contract method 0x5bee5f3a.
//
// Solidity: function slash(address user, address token, uint8 role, uint256 amount, string reason) returns()
func (_StakingManager *StakingManagerTransactor) Slash(opts *bind.TransactOpts, user common.Address, token common.Address, role uint8, amount *big.Int, reason string) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "slash", user, token, role, amount, reason)
}

// Slash is a paid mutator transaction binding the contract method 0x5bee5f3a.
//
// Solidity: function slash(address user, address token, uint8 role, uint256 amount, string reason) returns()
func (_StakingManager *StakingManagerSession) Slash(user common.Address, token common.Address, role uint8, amount *big.Int, reason string) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, user, token, role, amount, reason)
}

// Slash is a paid mutator transaction binding the contract method 0x5bee5f3a.
//
// Solidity: function slash(address user, address token, uint8 role, uint256 amount, string reason) returns()
func (_StakingManager *StakingManagerTransactorSession) Slash(user common.Address, token common.Address, role uint8, amount *big.Int, reason string) (*types.Transaction, error) {
	return _StakingManager.Contract.Slash(&_StakingManager.TransactOpts, user, token, role, amount, reason)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingManager *StakingManagerTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingManager *StakingManagerSession) Unpause() (*types.Transaction, error) {
	return _StakingManager.Contract.Unpause(&_StakingManager.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StakingManager *StakingManagerTransactorSession) Unpause() (*types.Transaction, error) {
	return _StakingManager.Contract.Unpause(&_StakingManager.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcc5a02cb.
//
// Solidity: function withdraw(address token, uint8 role) returns()
func (_StakingManager *StakingManagerTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, role uint8) (*types.Transaction, error) {
	return _StakingManager.contract.Transact(opts, "withdraw", token, role)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcc5a02cb.
//
// Solidity: function withdraw(address token, uint8 role) returns()
func (_StakingManager *StakingManagerSession) Withdraw(token common.Address, role uint8) (*types.Transaction, error) {
	return _StakingManager.Contract.Withdraw(&_StakingManager.TransactOpts, token, role)
}

// Withdraw is a paid mutator transaction binding the contract method 0xcc5a02cb.
//
// Solidity: function withdraw(address token, uint8 role) returns()
func (_StakingManager *StakingManagerTransactorSession) Withdraw(token common.Address, role uint8) (*types.Transaction, error) {
	return _StakingManager.Contract.Withdraw(&_StakingManager.TransactOpts, token, role)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_StakingManager *StakingManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _StakingManager.Contract.Receive(&_StakingManager.TransactOpts)
}

// StakingManagerAuthorizedCallerUpdatedIterator is returned from FilterAuthorizedCallerUpdated and is used to iterate over the raw logs and unpacked data for AuthorizedCallerUpdated events raised by the StakingManager contract.
type StakingManagerAuthorizedCallerUpdatedIterator struct {
	Event *StakingManagerAuthorizedCallerUpdated // Event containing the contract specifics and raw log

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
func (it *StakingManagerAuthorizedCallerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerAuthorizedCallerUpdated)
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
		it.Event = new(StakingManagerAuthorizedCallerUpdated)
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
func (it *StakingManagerAuthorizedCallerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerAuthorizedCallerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerAuthorizedCallerUpdated represents a AuthorizedCallerUpdated event raised by the StakingManager contract.
type StakingManagerAuthorizedCallerUpdated struct {
	Caller  common.Address
	Allowed bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuthorizedCallerUpdated is a free log retrieval operation binding the contract event 0xad857fa38c9319cb80848f1ef2f924383b90297eb5d71755738ff037d100faa1.
//
// Solidity: event AuthorizedCallerUpdated(address indexed caller, bool allowed)
func (_StakingManager *StakingManagerFilterer) FilterAuthorizedCallerUpdated(opts *bind.FilterOpts, caller []common.Address) (*StakingManagerAuthorizedCallerUpdatedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "AuthorizedCallerUpdated", callerRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerAuthorizedCallerUpdatedIterator{contract: _StakingManager.contract, event: "AuthorizedCallerUpdated", logs: logs, sub: sub}, nil
}

// WatchAuthorizedCallerUpdated is a free log subscription operation binding the contract event 0xad857fa38c9319cb80848f1ef2f924383b90297eb5d71755738ff037d100faa1.
//
// Solidity: event AuthorizedCallerUpdated(address indexed caller, bool allowed)
func (_StakingManager *StakingManagerFilterer) WatchAuthorizedCallerUpdated(opts *bind.WatchOpts, sink chan<- *StakingManagerAuthorizedCallerUpdated, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "AuthorizedCallerUpdated", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerAuthorizedCallerUpdated)
				if err := _StakingManager.contract.UnpackLog(event, "AuthorizedCallerUpdated", log); err != nil {
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

// ParseAuthorizedCallerUpdated is a log parse operation binding the contract event 0xad857fa38c9319cb80848f1ef2f924383b90297eb5d71755738ff037d100faa1.
//
// Solidity: event AuthorizedCallerUpdated(address indexed caller, bool allowed)
func (_StakingManager *StakingManagerFilterer) ParseAuthorizedCallerUpdated(log types.Log) (*StakingManagerAuthorizedCallerUpdated, error) {
	event := new(StakingManagerAuthorizedCallerUpdated)
	if err := _StakingManager.contract.UnpackLog(event, "AuthorizedCallerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerInitialLockPeriodUpdatedIterator is returned from FilterInitialLockPeriodUpdated and is used to iterate over the raw logs and unpacked data for InitialLockPeriodUpdated events raised by the StakingManager contract.
type StakingManagerInitialLockPeriodUpdatedIterator struct {
	Event *StakingManagerInitialLockPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *StakingManagerInitialLockPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerInitialLockPeriodUpdated)
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
		it.Event = new(StakingManagerInitialLockPeriodUpdated)
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
func (it *StakingManagerInitialLockPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerInitialLockPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerInitialLockPeriodUpdated represents a InitialLockPeriodUpdated event raised by the StakingManager contract.
type StakingManagerInitialLockPeriodUpdated struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitialLockPeriodUpdated is a free log retrieval operation binding the contract event 0x6f509e09f509e987131a2abbbbe5535b2dc71f3f9b35d47c4f6a28f9f2541699.
//
// Solidity: event InitialLockPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_StakingManager *StakingManagerFilterer) FilterInitialLockPeriodUpdated(opts *bind.FilterOpts) (*StakingManagerInitialLockPeriodUpdatedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "InitialLockPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingManagerInitialLockPeriodUpdatedIterator{contract: _StakingManager.contract, event: "InitialLockPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchInitialLockPeriodUpdated is a free log subscription operation binding the contract event 0x6f509e09f509e987131a2abbbbe5535b2dc71f3f9b35d47c4f6a28f9f2541699.
//
// Solidity: event InitialLockPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_StakingManager *StakingManagerFilterer) WatchInitialLockPeriodUpdated(opts *bind.WatchOpts, sink chan<- *StakingManagerInitialLockPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "InitialLockPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerInitialLockPeriodUpdated)
				if err := _StakingManager.contract.UnpackLog(event, "InitialLockPeriodUpdated", log); err != nil {
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

// ParseInitialLockPeriodUpdated is a log parse operation binding the contract event 0x6f509e09f509e987131a2abbbbe5535b2dc71f3f9b35d47c4f6a28f9f2541699.
//
// Solidity: event InitialLockPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_StakingManager *StakingManagerFilterer) ParseInitialLockPeriodUpdated(log types.Log) (*StakingManagerInitialLockPeriodUpdated, error) {
	event := new(StakingManagerInitialLockPeriodUpdated)
	if err := _StakingManager.contract.UnpackLog(event, "InitialLockPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StakingManager contract.
type StakingManagerInitializedIterator struct {
	Event *StakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *StakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerInitialized)
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
		it.Event = new(StakingManagerInitialized)
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
func (it *StakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerInitialized represents a Initialized event raised by the StakingManager contract.
type StakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingManagerInitializedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingManagerInitializedIterator{contract: _StakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StakingManager *StakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerInitialized)
				if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseInitialized(log types.Log) (*StakingManagerInitialized, error) {
	event := new(StakingManagerInitialized)
	if err := _StakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StakingManager contract.
type StakingManagerPausedIterator struct {
	Event *StakingManagerPaused // Event containing the contract specifics and raw log

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
func (it *StakingManagerPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerPaused)
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
		it.Event = new(StakingManagerPaused)
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
func (it *StakingManagerPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerPaused represents a Paused event raised by the StakingManager contract.
type StakingManagerPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingManager *StakingManagerFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingManagerPausedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingManagerPausedIterator{contract: _StakingManager.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StakingManager *StakingManagerFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingManagerPaused) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerPaused)
				if err := _StakingManager.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParsePaused(log types.Log) (*StakingManagerPaused, error) {
	event := new(StakingManagerPaused)
	if err := _StakingManager.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StakingManager contract.
type StakingManagerRoleAdminChangedIterator struct {
	Event *StakingManagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleAdminChanged)
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
		it.Event = new(StakingManagerRoleAdminChanged)
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
func (it *StakingManagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleAdminChanged represents a RoleAdminChanged event raised by the StakingManager contract.
type StakingManagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingManagerRoleAdminChangedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleAdminChangedIterator{contract: _StakingManager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StakingManager *StakingManagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleAdminChanged)
				if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleAdminChanged(log types.Log) (*StakingManagerRoleAdminChanged, error) {
	event := new(StakingManagerRoleAdminChanged)
	if err := _StakingManager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StakingManager contract.
type StakingManagerRoleGrantedIterator struct {
	Event *StakingManagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleGranted)
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
		it.Event = new(StakingManagerRoleGranted)
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
func (it *StakingManagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleGranted represents a RoleGranted event raised by the StakingManager contract.
type StakingManagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleGrantedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleGrantedIterator{contract: _StakingManager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleGranted)
				if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleGranted(log types.Log) (*StakingManagerRoleGranted, error) {
	event := new(StakingManagerRoleGranted)
	if err := _StakingManager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StakingManager contract.
type StakingManagerRoleRevokedIterator struct {
	Event *StakingManagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingManagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerRoleRevoked)
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
		it.Event = new(StakingManagerRoleRevoked)
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
func (it *StakingManagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerRoleRevoked represents a RoleRevoked event raised by the StakingManager contract.
type StakingManagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingManagerRoleRevokedIterator, error) {

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

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerRoleRevokedIterator{contract: _StakingManager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StakingManager *StakingManagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingManagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerRoleRevoked)
				if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseRoleRevoked(log types.Log) (*StakingManagerRoleRevoked, error) {
	event := new(StakingManagerRoleRevoked)
	if err := _StakingManager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeDepositedIterator is returned from FilterStakeDeposited and is used to iterate over the raw logs and unpacked data for StakeDeposited events raised by the StakingManager contract.
type StakingManagerStakeDepositedIterator struct {
	Event *StakingManagerStakeDeposited // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeDeposited)
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
		it.Event = new(StakingManagerStakeDeposited)
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
func (it *StakingManagerStakeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeDeposited represents a StakeDeposited event raised by the StakingManager contract.
type StakingManagerStakeDeposited struct {
	User   common.Address
	Token  common.Address
	Role   uint8
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeDeposited is a free log retrieval operation binding the contract event 0x6b99a8d22aef57b5be1ce24c59a97744e05a886a823d056b501114c654510316.
//
// Solidity: event StakeDeposited(address indexed user, address indexed token, uint8 indexed role, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterStakeDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address, role []uint8) (*StakingManagerStakeDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeDeposited", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeDepositedIterator{contract: _StakingManager.contract, event: "StakeDeposited", logs: logs, sub: sub}, nil
}

// WatchStakeDeposited is a free log subscription operation binding the contract event 0x6b99a8d22aef57b5be1ce24c59a97744e05a886a823d056b501114c654510316.
//
// Solidity: event StakeDeposited(address indexed user, address indexed token, uint8 indexed role, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchStakeDeposited(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeDeposited, user []common.Address, token []common.Address, role []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeDeposited", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeDeposited)
				if err := _StakingManager.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
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

// ParseStakeDeposited is a log parse operation binding the contract event 0x6b99a8d22aef57b5be1ce24c59a97744e05a886a823d056b501114c654510316.
//
// Solidity: event StakeDeposited(address indexed user, address indexed token, uint8 indexed role, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseStakeDeposited(log types.Log) (*StakingManagerStakeDeposited, error) {
	event := new(StakingManagerStakeDeposited)
	if err := _StakingManager.contract.UnpackLog(event, "StakeDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeSlashedIterator is returned from FilterStakeSlashed and is used to iterate over the raw logs and unpacked data for StakeSlashed events raised by the StakingManager contract.
type StakingManagerStakeSlashedIterator struct {
	Event *StakingManagerStakeSlashed // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeSlashed)
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
		it.Event = new(StakingManagerStakeSlashed)
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
func (it *StakingManagerStakeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeSlashed represents a StakeSlashed event raised by the StakingManager contract.
type StakingManagerStakeSlashed struct {
	User   common.Address
	Token  common.Address
	Role   uint8
	Amount *big.Int
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeSlashed is a free log retrieval operation binding the contract event 0x36673588113c15d7cff24e3b3f02a2a7e1254ad3336b915a9e3d89309621a180.
//
// Solidity: event StakeSlashed(address indexed user, address indexed token, uint8 indexed role, uint256 amount, string reason)
func (_StakingManager *StakingManagerFilterer) FilterStakeSlashed(opts *bind.FilterOpts, user []common.Address, token []common.Address, role []uint8) (*StakingManagerStakeSlashedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeSlashed", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeSlashedIterator{contract: _StakingManager.contract, event: "StakeSlashed", logs: logs, sub: sub}, nil
}

// WatchStakeSlashed is a free log subscription operation binding the contract event 0x36673588113c15d7cff24e3b3f02a2a7e1254ad3336b915a9e3d89309621a180.
//
// Solidity: event StakeSlashed(address indexed user, address indexed token, uint8 indexed role, uint256 amount, string reason)
func (_StakingManager *StakingManagerFilterer) WatchStakeSlashed(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeSlashed, user []common.Address, token []common.Address, role []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeSlashed", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeSlashed)
				if err := _StakingManager.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
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

// ParseStakeSlashed is a log parse operation binding the contract event 0x36673588113c15d7cff24e3b3f02a2a7e1254ad3336b915a9e3d89309621a180.
//
// Solidity: event StakeSlashed(address indexed user, address indexed token, uint8 indexed role, uint256 amount, string reason)
func (_StakingManager *StakingManagerFilterer) ParseStakeSlashed(log types.Log) (*StakingManagerStakeSlashed, error) {
	event := new(StakingManagerStakeSlashed)
	if err := _StakingManager.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeToppedUpIterator is returned from FilterStakeToppedUp and is used to iterate over the raw logs and unpacked data for StakeToppedUp events raised by the StakingManager contract.
type StakingManagerStakeToppedUpIterator struct {
	Event *StakingManagerStakeToppedUp // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeToppedUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeToppedUp)
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
		it.Event = new(StakingManagerStakeToppedUp)
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
func (it *StakingManagerStakeToppedUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeToppedUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeToppedUp represents a StakeToppedUp event raised by the StakingManager contract.
type StakingManagerStakeToppedUp struct {
	User     common.Address
	Token    common.Address
	Role     uint8
	Amount   *big.Int
	NewTotal *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStakeToppedUp is a free log retrieval operation binding the contract event 0x2f163c03de81cc10616878005e54a0ec9f4a95238568d8bab7197f341edddff0.
//
// Solidity: event StakeToppedUp(address indexed user, address indexed token, uint8 indexed role, uint256 amount, uint256 newTotal)
func (_StakingManager *StakingManagerFilterer) FilterStakeToppedUp(opts *bind.FilterOpts, user []common.Address, token []common.Address, role []uint8) (*StakingManagerStakeToppedUpIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeToppedUp", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeToppedUpIterator{contract: _StakingManager.contract, event: "StakeToppedUp", logs: logs, sub: sub}, nil
}

// WatchStakeToppedUp is a free log subscription operation binding the contract event 0x2f163c03de81cc10616878005e54a0ec9f4a95238568d8bab7197f341edddff0.
//
// Solidity: event StakeToppedUp(address indexed user, address indexed token, uint8 indexed role, uint256 amount, uint256 newTotal)
func (_StakingManager *StakingManagerFilterer) WatchStakeToppedUp(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeToppedUp, user []common.Address, token []common.Address, role []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeToppedUp", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeToppedUp)
				if err := _StakingManager.contract.UnpackLog(event, "StakeToppedUp", log); err != nil {
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

// ParseStakeToppedUp is a log parse operation binding the contract event 0x2f163c03de81cc10616878005e54a0ec9f4a95238568d8bab7197f341edddff0.
//
// Solidity: event StakeToppedUp(address indexed user, address indexed token, uint8 indexed role, uint256 amount, uint256 newTotal)
func (_StakingManager *StakingManagerFilterer) ParseStakeToppedUp(log types.Log) (*StakingManagerStakeToppedUp, error) {
	event := new(StakingManagerStakeToppedUp)
	if err := _StakingManager.contract.UnpackLog(event, "StakeToppedUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the StakingManager contract.
type StakingManagerStakeWithdrawnIterator struct {
	Event *StakingManagerStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingManagerStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerStakeWithdrawn)
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
		it.Event = new(StakingManagerStakeWithdrawn)
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
func (it *StakingManagerStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerStakeWithdrawn represents a StakeWithdrawn event raised by the StakingManager contract.
type StakingManagerStakeWithdrawn struct {
	User   common.Address
	Token  common.Address
	Role   uint8
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x454e8e02a030acc47b8222ce0377a03830511b834d41e4410f67592ba1ea7d73.
//
// Solidity: event StakeWithdrawn(address indexed user, address indexed token, uint8 indexed role, uint256 amount)
func (_StakingManager *StakingManagerFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, user []common.Address, token []common.Address, role []uint8) (*StakingManagerStakeWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "StakeWithdrawn", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerStakeWithdrawnIterator{contract: _StakingManager.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x454e8e02a030acc47b8222ce0377a03830511b834d41e4410f67592ba1ea7d73.
//
// Solidity: event StakeWithdrawn(address indexed user, address indexed token, uint8 indexed role, uint256 amount)
func (_StakingManager *StakingManagerFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingManagerStakeWithdrawn, user []common.Address, token []common.Address, role []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "StakeWithdrawn", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerStakeWithdrawn)
				if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x454e8e02a030acc47b8222ce0377a03830511b834d41e4410f67592ba1ea7d73.
//
// Solidity: event StakeWithdrawn(address indexed user, address indexed token, uint8 indexed role, uint256 amount)
func (_StakingManager *StakingManagerFilterer) ParseStakeWithdrawn(log types.Log) (*StakingManagerStakeWithdrawn, error) {
	event := new(StakingManagerStakeWithdrawn)
	if err := _StakingManager.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnlockPeriodUpdatedIterator is returned from FilterUnlockPeriodUpdated and is used to iterate over the raw logs and unpacked data for UnlockPeriodUpdated events raised by the StakingManager contract.
type StakingManagerUnlockPeriodUpdatedIterator struct {
	Event *StakingManagerUnlockPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnlockPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnlockPeriodUpdated)
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
		it.Event = new(StakingManagerUnlockPeriodUpdated)
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
func (it *StakingManagerUnlockPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnlockPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnlockPeriodUpdated represents a UnlockPeriodUpdated event raised by the StakingManager contract.
type StakingManagerUnlockPeriodUpdated struct {
	OldPeriod *big.Int
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockPeriodUpdated is a free log retrieval operation binding the contract event 0xaa1e10941c3aafb56ad74dad40ea5ec52cf44d83495362e44c775124edb040f5.
//
// Solidity: event UnlockPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_StakingManager *StakingManagerFilterer) FilterUnlockPeriodUpdated(opts *bind.FilterOpts) (*StakingManagerUnlockPeriodUpdatedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnlockPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnlockPeriodUpdatedIterator{contract: _StakingManager.contract, event: "UnlockPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchUnlockPeriodUpdated is a free log subscription operation binding the contract event 0xaa1e10941c3aafb56ad74dad40ea5ec52cf44d83495362e44c775124edb040f5.
//
// Solidity: event UnlockPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_StakingManager *StakingManagerFilterer) WatchUnlockPeriodUpdated(opts *bind.WatchOpts, sink chan<- *StakingManagerUnlockPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnlockPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnlockPeriodUpdated)
				if err := _StakingManager.contract.UnpackLog(event, "UnlockPeriodUpdated", log); err != nil {
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

// ParseUnlockPeriodUpdated is a log parse operation binding the contract event 0xaa1e10941c3aafb56ad74dad40ea5ec52cf44d83495362e44c775124edb040f5.
//
// Solidity: event UnlockPeriodUpdated(uint256 oldPeriod, uint256 newPeriod)
func (_StakingManager *StakingManagerFilterer) ParseUnlockPeriodUpdated(log types.Log) (*StakingManagerUnlockPeriodUpdated, error) {
	event := new(StakingManagerUnlockPeriodUpdated)
	if err := _StakingManager.contract.UnpackLog(event, "UnlockPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StakingManager contract.
type StakingManagerUnpausedIterator struct {
	Event *StakingManagerUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnpaused)
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
		it.Event = new(StakingManagerUnpaused)
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
func (it *StakingManagerUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnpaused represents a Unpaused event raised by the StakingManager contract.
type StakingManagerUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingManager *StakingManagerFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingManagerUnpausedIterator, error) {

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnpausedIterator{contract: _StakingManager.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StakingManager *StakingManagerFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingManagerUnpaused) (event.Subscription, error) {

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnpaused)
				if err := _StakingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_StakingManager *StakingManagerFilterer) ParseUnpaused(log types.Log) (*StakingManagerUnpaused, error) {
	event := new(StakingManagerUnpaused)
	if err := _StakingManager.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingManagerUnstakeRequestedIterator is returned from FilterUnstakeRequested and is used to iterate over the raw logs and unpacked data for UnstakeRequested events raised by the StakingManager contract.
type StakingManagerUnstakeRequestedIterator struct {
	Event *StakingManagerUnstakeRequested // Event containing the contract specifics and raw log

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
func (it *StakingManagerUnstakeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingManagerUnstakeRequested)
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
		it.Event = new(StakingManagerUnstakeRequested)
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
func (it *StakingManagerUnstakeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingManagerUnstakeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingManagerUnstakeRequested represents a UnstakeRequested event raised by the StakingManager contract.
type StakingManagerUnstakeRequested struct {
	User       common.Address
	Token      common.Address
	Role       uint8
	Amount     *big.Int
	UnlockTime *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUnstakeRequested is a free log retrieval operation binding the contract event 0x330db065fe0f6c59ef4d8ea2e564ed016c0fdf26944e2ea69e4be64e80967d49.
//
// Solidity: event UnstakeRequested(address indexed user, address indexed token, uint8 indexed role, uint256 amount, uint256 unlockTime)
func (_StakingManager *StakingManagerFilterer) FilterUnstakeRequested(opts *bind.FilterOpts, user []common.Address, token []common.Address, role []uint8) (*StakingManagerUnstakeRequestedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.FilterLogs(opts, "UnstakeRequested", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &StakingManagerUnstakeRequestedIterator{contract: _StakingManager.contract, event: "UnstakeRequested", logs: logs, sub: sub}, nil
}

// WatchUnstakeRequested is a free log subscription operation binding the contract event 0x330db065fe0f6c59ef4d8ea2e564ed016c0fdf26944e2ea69e4be64e80967d49.
//
// Solidity: event UnstakeRequested(address indexed user, address indexed token, uint8 indexed role, uint256 amount, uint256 unlockTime)
func (_StakingManager *StakingManagerFilterer) WatchUnstakeRequested(opts *bind.WatchOpts, sink chan<- *StakingManagerUnstakeRequested, user []common.Address, token []common.Address, role []uint8) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _StakingManager.contract.WatchLogs(opts, "UnstakeRequested", userRule, tokenRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingManagerUnstakeRequested)
				if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
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

// ParseUnstakeRequested is a log parse operation binding the contract event 0x330db065fe0f6c59ef4d8ea2e564ed016c0fdf26944e2ea69e4be64e80967d49.
//
// Solidity: event UnstakeRequested(address indexed user, address indexed token, uint8 indexed role, uint256 amount, uint256 unlockTime)
func (_StakingManager *StakingManagerFilterer) ParseUnstakeRequested(log types.Log) (*StakingManagerUnstakeRequested, error) {
	event := new(StakingManagerUnstakeRequested)
	if err := _StakingManager.contract.UnpackLog(event, "UnstakeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
