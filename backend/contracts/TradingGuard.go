// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TradingGuardMetaData contains all meta data concerning the TradingGuard contract.
var TradingGuardMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bot\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BOT_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resumeTrading\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"triggerCircuitBreaker\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EmergencyStopTriggered\",\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TradingResumed\",\"inputs\":[{\"name\":\"initiator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]}]",
}

// TradingGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use TradingGuardMetaData.ABI instead.
var TradingGuardABI = TradingGuardMetaData.ABI

// TradingGuard is an auto generated Go binding around an Ethereum contract.
type TradingGuard struct {
	TradingGuardCaller     // Read-only binding to the contract
	TradingGuardTransactor // Write-only binding to the contract
	TradingGuardFilterer   // Log filterer for contract events
}

// TradingGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradingGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradingGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradingGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradingGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradingGuardSession struct {
	Contract     *TradingGuard     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradingGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradingGuardCallerSession struct {
	Contract *TradingGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TradingGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradingGuardTransactorSession struct {
	Contract     *TradingGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TradingGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradingGuardRaw struct {
	Contract *TradingGuard // Generic contract binding to access the raw methods on
}

// TradingGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradingGuardCallerRaw struct {
	Contract *TradingGuardCaller // Generic read-only contract binding to access the raw methods on
}

// TradingGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradingGuardTransactorRaw struct {
	Contract *TradingGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradingGuard creates a new instance of TradingGuard, bound to a specific deployed contract.
func NewTradingGuard(address common.Address, backend bind.ContractBackend) (*TradingGuard, error) {
	contract, err := bindTradingGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradingGuard{TradingGuardCaller: TradingGuardCaller{contract: contract}, TradingGuardTransactor: TradingGuardTransactor{contract: contract}, TradingGuardFilterer: TradingGuardFilterer{contract: contract}}, nil
}

// NewTradingGuardCaller creates a new read-only instance of TradingGuard, bound to a specific deployed contract.
func NewTradingGuardCaller(address common.Address, caller bind.ContractCaller) (*TradingGuardCaller, error) {
	contract, err := bindTradingGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradingGuardCaller{contract: contract}, nil
}

// NewTradingGuardTransactor creates a new write-only instance of TradingGuard, bound to a specific deployed contract.
func NewTradingGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*TradingGuardTransactor, error) {
	contract, err := bindTradingGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradingGuardTransactor{contract: contract}, nil
}

// NewTradingGuardFilterer creates a new log filterer instance of TradingGuard, bound to a specific deployed contract.
func NewTradingGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*TradingGuardFilterer, error) {
	contract, err := bindTradingGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradingGuardFilterer{contract: contract}, nil
}

// bindTradingGuard binds a generic wrapper to an already deployed contract.
func bindTradingGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TradingGuardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingGuard *TradingGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingGuard.Contract.TradingGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingGuard *TradingGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingGuard.Contract.TradingGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingGuard *TradingGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingGuard.Contract.TradingGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradingGuard *TradingGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradingGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradingGuard *TradingGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradingGuard *TradingGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradingGuard.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingGuard.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardSession) ADMINROLE() ([32]byte, error) {
	return _TradingGuard.Contract.ADMINROLE(&_TradingGuard.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardCallerSession) ADMINROLE() ([32]byte, error) {
	return _TradingGuard.Contract.ADMINROLE(&_TradingGuard.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardCaller) BOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingGuard.contract.Call(opts, &out, "BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardSession) BOTROLE() ([32]byte, error) {
	return _TradingGuard.Contract.BOTROLE(&_TradingGuard.CallOpts)
}

// BOTROLE is a free data retrieval call binding the contract method 0xb1503774.
//
// Solidity: function BOT_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardCallerSession) BOTROLE() ([32]byte, error) {
	return _TradingGuard.Contract.BOTROLE(&_TradingGuard.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TradingGuard.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TradingGuard.Contract.DEFAULTADMINROLE(&_TradingGuard.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TradingGuard *TradingGuardCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TradingGuard.Contract.DEFAULTADMINROLE(&_TradingGuard.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingGuard *TradingGuardCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TradingGuard.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingGuard *TradingGuardSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TradingGuard.Contract.GetRoleAdmin(&_TradingGuard.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TradingGuard *TradingGuardCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TradingGuard.Contract.GetRoleAdmin(&_TradingGuard.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingGuard *TradingGuardCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TradingGuard.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingGuard *TradingGuardSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TradingGuard.Contract.HasRole(&_TradingGuard.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TradingGuard *TradingGuardCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TradingGuard.Contract.HasRole(&_TradingGuard.CallOpts, role, account)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TradingGuard *TradingGuardCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TradingGuard.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TradingGuard *TradingGuardSession) Paused() (bool, error) {
	return _TradingGuard.Contract.Paused(&_TradingGuard.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TradingGuard *TradingGuardCallerSession) Paused() (bool, error) {
	return _TradingGuard.Contract.Paused(&_TradingGuard.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingGuard *TradingGuardCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TradingGuard.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingGuard *TradingGuardSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TradingGuard.Contract.SupportsInterface(&_TradingGuard.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TradingGuard *TradingGuardCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TradingGuard.Contract.SupportsInterface(&_TradingGuard.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingGuard *TradingGuardTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingGuard.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingGuard *TradingGuardSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingGuard.Contract.GrantRole(&_TradingGuard.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TradingGuard *TradingGuardTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingGuard.Contract.GrantRole(&_TradingGuard.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingGuard *TradingGuardTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingGuard.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingGuard *TradingGuardSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingGuard.Contract.RenounceRole(&_TradingGuard.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_TradingGuard *TradingGuardTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _TradingGuard.Contract.RenounceRole(&_TradingGuard.TransactOpts, role, callerConfirmation)
}

// ResumeTrading is a paid mutator transaction binding the contract method 0x0694db1e.
//
// Solidity: function resumeTrading() returns()
func (_TradingGuard *TradingGuardTransactor) ResumeTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradingGuard.contract.Transact(opts, "resumeTrading")
}

// ResumeTrading is a paid mutator transaction binding the contract method 0x0694db1e.
//
// Solidity: function resumeTrading() returns()
func (_TradingGuard *TradingGuardSession) ResumeTrading() (*types.Transaction, error) {
	return _TradingGuard.Contract.ResumeTrading(&_TradingGuard.TransactOpts)
}

// ResumeTrading is a paid mutator transaction binding the contract method 0x0694db1e.
//
// Solidity: function resumeTrading() returns()
func (_TradingGuard *TradingGuardTransactorSession) ResumeTrading() (*types.Transaction, error) {
	return _TradingGuard.Contract.ResumeTrading(&_TradingGuard.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingGuard *TradingGuardTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingGuard.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingGuard *TradingGuardSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingGuard.Contract.RevokeRole(&_TradingGuard.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TradingGuard *TradingGuardTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TradingGuard.Contract.RevokeRole(&_TradingGuard.TransactOpts, role, account)
}

// TriggerCircuitBreaker is a paid mutator transaction binding the contract method 0x61b242c5.
//
// Solidity: function triggerCircuitBreaker(string reason) returns()
func (_TradingGuard *TradingGuardTransactor) TriggerCircuitBreaker(opts *bind.TransactOpts, reason string) (*types.Transaction, error) {
	return _TradingGuard.contract.Transact(opts, "triggerCircuitBreaker", reason)
}

// TriggerCircuitBreaker is a paid mutator transaction binding the contract method 0x61b242c5.
//
// Solidity: function triggerCircuitBreaker(string reason) returns()
func (_TradingGuard *TradingGuardSession) TriggerCircuitBreaker(reason string) (*types.Transaction, error) {
	return _TradingGuard.Contract.TriggerCircuitBreaker(&_TradingGuard.TransactOpts, reason)
}

// TriggerCircuitBreaker is a paid mutator transaction binding the contract method 0x61b242c5.
//
// Solidity: function triggerCircuitBreaker(string reason) returns()
func (_TradingGuard *TradingGuardTransactorSession) TriggerCircuitBreaker(reason string) (*types.Transaction, error) {
	return _TradingGuard.Contract.TriggerCircuitBreaker(&_TradingGuard.TransactOpts, reason)
}

// TradingGuardEmergencyStopTriggeredIterator is returned from FilterEmergencyStopTriggered and is used to iterate over the raw logs and unpacked data for EmergencyStopTriggered events raised by the TradingGuard contract.
type TradingGuardEmergencyStopTriggeredIterator struct {
	Event *TradingGuardEmergencyStopTriggered // Event containing the contract specifics and raw log

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
func (it *TradingGuardEmergencyStopTriggeredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingGuardEmergencyStopTriggered)
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
		it.Event = new(TradingGuardEmergencyStopTriggered)
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
func (it *TradingGuardEmergencyStopTriggeredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingGuardEmergencyStopTriggeredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingGuardEmergencyStopTriggered represents a EmergencyStopTriggered event raised by the TradingGuard contract.
type TradingGuardEmergencyStopTriggered struct {
	Initiator common.Address
	Reason    string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEmergencyStopTriggered is a free log retrieval operation binding the contract event 0xa51c2899420ebe4053bea840e365e915f284333379d46edccab38f5661fe00c0.
//
// Solidity: event EmergencyStopTriggered(address indexed initiator, string reason)
func (_TradingGuard *TradingGuardFilterer) FilterEmergencyStopTriggered(opts *bind.FilterOpts, initiator []common.Address) (*TradingGuardEmergencyStopTriggeredIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _TradingGuard.contract.FilterLogs(opts, "EmergencyStopTriggered", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &TradingGuardEmergencyStopTriggeredIterator{contract: _TradingGuard.contract, event: "EmergencyStopTriggered", logs: logs, sub: sub}, nil
}

// WatchEmergencyStopTriggered is a free log subscription operation binding the contract event 0xa51c2899420ebe4053bea840e365e915f284333379d46edccab38f5661fe00c0.
//
// Solidity: event EmergencyStopTriggered(address indexed initiator, string reason)
func (_TradingGuard *TradingGuardFilterer) WatchEmergencyStopTriggered(opts *bind.WatchOpts, sink chan<- *TradingGuardEmergencyStopTriggered, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _TradingGuard.contract.WatchLogs(opts, "EmergencyStopTriggered", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingGuardEmergencyStopTriggered)
				if err := _TradingGuard.contract.UnpackLog(event, "EmergencyStopTriggered", log); err != nil {
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

// ParseEmergencyStopTriggered is a log parse operation binding the contract event 0xa51c2899420ebe4053bea840e365e915f284333379d46edccab38f5661fe00c0.
//
// Solidity: event EmergencyStopTriggered(address indexed initiator, string reason)
func (_TradingGuard *TradingGuardFilterer) ParseEmergencyStopTriggered(log types.Log) (*TradingGuardEmergencyStopTriggered, error) {
	event := new(TradingGuardEmergencyStopTriggered)
	if err := _TradingGuard.contract.UnpackLog(event, "EmergencyStopTriggered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingGuardPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TradingGuard contract.
type TradingGuardPausedIterator struct {
	Event *TradingGuardPaused // Event containing the contract specifics and raw log

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
func (it *TradingGuardPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingGuardPaused)
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
		it.Event = new(TradingGuardPaused)
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
func (it *TradingGuardPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingGuardPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingGuardPaused represents a Paused event raised by the TradingGuard contract.
type TradingGuardPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TradingGuard *TradingGuardFilterer) FilterPaused(opts *bind.FilterOpts) (*TradingGuardPausedIterator, error) {

	logs, sub, err := _TradingGuard.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TradingGuardPausedIterator{contract: _TradingGuard.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TradingGuard *TradingGuardFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TradingGuardPaused) (event.Subscription, error) {

	logs, sub, err := _TradingGuard.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingGuardPaused)
				if err := _TradingGuard.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TradingGuard *TradingGuardFilterer) ParsePaused(log types.Log) (*TradingGuardPaused, error) {
	event := new(TradingGuardPaused)
	if err := _TradingGuard.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingGuardRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TradingGuard contract.
type TradingGuardRoleAdminChangedIterator struct {
	Event *TradingGuardRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TradingGuardRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingGuardRoleAdminChanged)
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
		it.Event = new(TradingGuardRoleAdminChanged)
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
func (it *TradingGuardRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingGuardRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingGuardRoleAdminChanged represents a RoleAdminChanged event raised by the TradingGuard contract.
type TradingGuardRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TradingGuard *TradingGuardFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TradingGuardRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TradingGuard.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TradingGuardRoleAdminChangedIterator{contract: _TradingGuard.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TradingGuard *TradingGuardFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TradingGuardRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TradingGuard.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingGuardRoleAdminChanged)
				if err := _TradingGuard.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TradingGuard *TradingGuardFilterer) ParseRoleAdminChanged(log types.Log) (*TradingGuardRoleAdminChanged, error) {
	event := new(TradingGuardRoleAdminChanged)
	if err := _TradingGuard.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingGuardRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TradingGuard contract.
type TradingGuardRoleGrantedIterator struct {
	Event *TradingGuardRoleGranted // Event containing the contract specifics and raw log

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
func (it *TradingGuardRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingGuardRoleGranted)
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
		it.Event = new(TradingGuardRoleGranted)
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
func (it *TradingGuardRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingGuardRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingGuardRoleGranted represents a RoleGranted event raised by the TradingGuard contract.
type TradingGuardRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingGuard *TradingGuardFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TradingGuardRoleGrantedIterator, error) {

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

	logs, sub, err := _TradingGuard.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TradingGuardRoleGrantedIterator{contract: _TradingGuard.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingGuard *TradingGuardFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TradingGuardRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TradingGuard.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingGuardRoleGranted)
				if err := _TradingGuard.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TradingGuard *TradingGuardFilterer) ParseRoleGranted(log types.Log) (*TradingGuardRoleGranted, error) {
	event := new(TradingGuardRoleGranted)
	if err := _TradingGuard.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingGuardRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TradingGuard contract.
type TradingGuardRoleRevokedIterator struct {
	Event *TradingGuardRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TradingGuardRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingGuardRoleRevoked)
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
		it.Event = new(TradingGuardRoleRevoked)
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
func (it *TradingGuardRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingGuardRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingGuardRoleRevoked represents a RoleRevoked event raised by the TradingGuard contract.
type TradingGuardRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingGuard *TradingGuardFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TradingGuardRoleRevokedIterator, error) {

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

	logs, sub, err := _TradingGuard.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TradingGuardRoleRevokedIterator{contract: _TradingGuard.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TradingGuard *TradingGuardFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TradingGuardRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TradingGuard.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingGuardRoleRevoked)
				if err := _TradingGuard.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TradingGuard *TradingGuardFilterer) ParseRoleRevoked(log types.Log) (*TradingGuardRoleRevoked, error) {
	event := new(TradingGuardRoleRevoked)
	if err := _TradingGuard.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingGuardTradingResumedIterator is returned from FilterTradingResumed and is used to iterate over the raw logs and unpacked data for TradingResumed events raised by the TradingGuard contract.
type TradingGuardTradingResumedIterator struct {
	Event *TradingGuardTradingResumed // Event containing the contract specifics and raw log

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
func (it *TradingGuardTradingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingGuardTradingResumed)
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
		it.Event = new(TradingGuardTradingResumed)
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
func (it *TradingGuardTradingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingGuardTradingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingGuardTradingResumed represents a TradingResumed event raised by the TradingGuard contract.
type TradingGuardTradingResumed struct {
	Initiator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTradingResumed is a free log retrieval operation binding the contract event 0xd0e147a2c5ae727118c4a9359445a77d77d81945ccc12a8feec087160dae00a5.
//
// Solidity: event TradingResumed(address indexed initiator)
func (_TradingGuard *TradingGuardFilterer) FilterTradingResumed(opts *bind.FilterOpts, initiator []common.Address) (*TradingGuardTradingResumedIterator, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _TradingGuard.contract.FilterLogs(opts, "TradingResumed", initiatorRule)
	if err != nil {
		return nil, err
	}
	return &TradingGuardTradingResumedIterator{contract: _TradingGuard.contract, event: "TradingResumed", logs: logs, sub: sub}, nil
}

// WatchTradingResumed is a free log subscription operation binding the contract event 0xd0e147a2c5ae727118c4a9359445a77d77d81945ccc12a8feec087160dae00a5.
//
// Solidity: event TradingResumed(address indexed initiator)
func (_TradingGuard *TradingGuardFilterer) WatchTradingResumed(opts *bind.WatchOpts, sink chan<- *TradingGuardTradingResumed, initiator []common.Address) (event.Subscription, error) {

	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}

	logs, sub, err := _TradingGuard.contract.WatchLogs(opts, "TradingResumed", initiatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingGuardTradingResumed)
				if err := _TradingGuard.contract.UnpackLog(event, "TradingResumed", log); err != nil {
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

// ParseTradingResumed is a log parse operation binding the contract event 0xd0e147a2c5ae727118c4a9359445a77d77d81945ccc12a8feec087160dae00a5.
//
// Solidity: event TradingResumed(address indexed initiator)
func (_TradingGuard *TradingGuardFilterer) ParseTradingResumed(log types.Log) (*TradingGuardTradingResumed, error) {
	event := new(TradingGuardTradingResumed)
	if err := _TradingGuard.contract.UnpackLog(event, "TradingResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TradingGuardUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TradingGuard contract.
type TradingGuardUnpausedIterator struct {
	Event *TradingGuardUnpaused // Event containing the contract specifics and raw log

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
func (it *TradingGuardUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TradingGuardUnpaused)
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
		it.Event = new(TradingGuardUnpaused)
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
func (it *TradingGuardUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TradingGuardUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TradingGuardUnpaused represents a Unpaused event raised by the TradingGuard contract.
type TradingGuardUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TradingGuard *TradingGuardFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TradingGuardUnpausedIterator, error) {

	logs, sub, err := _TradingGuard.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TradingGuardUnpausedIterator{contract: _TradingGuard.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TradingGuard *TradingGuardFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TradingGuardUnpaused) (event.Subscription, error) {

	logs, sub, err := _TradingGuard.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TradingGuardUnpaused)
				if err := _TradingGuard.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TradingGuard *TradingGuardFilterer) ParseUnpaused(log types.Log) (*TradingGuardUnpaused, error) {
	event := new(TradingGuardUnpaused)
	if err := _TradingGuard.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
