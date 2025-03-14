// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package com

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

// AutomationCompatibleMetaData contains all meta data concerning the AutomationCompatible contract.
var AutomationCompatibleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"}],\"name\":\"checkUpkeep\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"upkeepNeeded\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"performData\",\"type\":\"bytes\"}],\"name\":\"performUpkeep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AutomationCompatibleABI is the input ABI used to generate the binding from.
// Deprecated: Use AutomationCompatibleMetaData.ABI instead.
var AutomationCompatibleABI = AutomationCompatibleMetaData.ABI

// AutomationCompatible is an auto generated Go binding around an Ethereum contract.
type AutomationCompatible struct {
	AutomationCompatibleCaller     // Read-only binding to the contract
	AutomationCompatibleTransactor // Write-only binding to the contract
	AutomationCompatibleFilterer   // Log filterer for contract events
}

// AutomationCompatibleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AutomationCompatibleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationCompatibleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AutomationCompatibleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationCompatibleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AutomationCompatibleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationCompatibleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AutomationCompatibleSession struct {
	Contract     *AutomationCompatible // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AutomationCompatibleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AutomationCompatibleCallerSession struct {
	Contract *AutomationCompatibleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AutomationCompatibleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AutomationCompatibleTransactorSession struct {
	Contract     *AutomationCompatibleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AutomationCompatibleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AutomationCompatibleRaw struct {
	Contract *AutomationCompatible // Generic contract binding to access the raw methods on
}

// AutomationCompatibleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AutomationCompatibleCallerRaw struct {
	Contract *AutomationCompatibleCaller // Generic read-only contract binding to access the raw methods on
}

// AutomationCompatibleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AutomationCompatibleTransactorRaw struct {
	Contract *AutomationCompatibleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAutomationCompatible creates a new instance of AutomationCompatible, bound to a specific deployed contract.
func NewAutomationCompatible(address common.Address, backend bind.ContractBackend) (*AutomationCompatible, error) {
	contract, err := bindAutomationCompatible(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AutomationCompatible{AutomationCompatibleCaller: AutomationCompatibleCaller{contract: contract}, AutomationCompatibleTransactor: AutomationCompatibleTransactor{contract: contract}, AutomationCompatibleFilterer: AutomationCompatibleFilterer{contract: contract}}, nil
}

// NewAutomationCompatibleCaller creates a new read-only instance of AutomationCompatible, bound to a specific deployed contract.
func NewAutomationCompatibleCaller(address common.Address, caller bind.ContractCaller) (*AutomationCompatibleCaller, error) {
	contract, err := bindAutomationCompatible(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationCompatibleCaller{contract: contract}, nil
}

// NewAutomationCompatibleTransactor creates a new write-only instance of AutomationCompatible, bound to a specific deployed contract.
func NewAutomationCompatibleTransactor(address common.Address, transactor bind.ContractTransactor) (*AutomationCompatibleTransactor, error) {
	contract, err := bindAutomationCompatible(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationCompatibleTransactor{contract: contract}, nil
}

// NewAutomationCompatibleFilterer creates a new log filterer instance of AutomationCompatible, bound to a specific deployed contract.
func NewAutomationCompatibleFilterer(address common.Address, filterer bind.ContractFilterer) (*AutomationCompatibleFilterer, error) {
	contract, err := bindAutomationCompatible(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutomationCompatibleFilterer{contract: contract}, nil
}

// bindAutomationCompatible binds a generic wrapper to an already deployed contract.
func bindAutomationCompatible(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AutomationCompatibleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutomationCompatible *AutomationCompatibleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationCompatible.Contract.AutomationCompatibleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutomationCompatible *AutomationCompatibleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.AutomationCompatibleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutomationCompatible *AutomationCompatibleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.AutomationCompatibleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutomationCompatible *AutomationCompatibleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationCompatible.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutomationCompatible *AutomationCompatibleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutomationCompatible *AutomationCompatibleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.contract.Transact(opts, method, params...)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) returns(bool upkeepNeeded, bytes performData)
func (_AutomationCompatible *AutomationCompatibleTransactor) CheckUpkeep(opts *bind.TransactOpts, checkData []byte) (*types.Transaction, error) {
	return _AutomationCompatible.contract.Transact(opts, "checkUpkeep", checkData)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) returns(bool upkeepNeeded, bytes performData)
func (_AutomationCompatible *AutomationCompatibleSession) CheckUpkeep(checkData []byte) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.CheckUpkeep(&_AutomationCompatible.TransactOpts, checkData)
}

// CheckUpkeep is a paid mutator transaction binding the contract method 0x6e04ff0d.
//
// Solidity: function checkUpkeep(bytes checkData) returns(bool upkeepNeeded, bytes performData)
func (_AutomationCompatible *AutomationCompatibleTransactorSession) CheckUpkeep(checkData []byte) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.CheckUpkeep(&_AutomationCompatible.TransactOpts, checkData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_AutomationCompatible *AutomationCompatibleTransactor) PerformUpkeep(opts *bind.TransactOpts, performData []byte) (*types.Transaction, error) {
	return _AutomationCompatible.contract.Transact(opts, "performUpkeep", performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_AutomationCompatible *AutomationCompatibleSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.PerformUpkeep(&_AutomationCompatible.TransactOpts, performData)
}

// PerformUpkeep is a paid mutator transaction binding the contract method 0x4585e33b.
//
// Solidity: function performUpkeep(bytes performData) returns()
func (_AutomationCompatible *AutomationCompatibleTransactorSession) PerformUpkeep(performData []byte) (*types.Transaction, error) {
	return _AutomationCompatible.Contract.PerformUpkeep(&_AutomationCompatible.TransactOpts, performData)
}
