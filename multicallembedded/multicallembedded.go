// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multicallembedded

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

// MulticallembeddedMetaData contains all meta data concerning the Multicallembedded contract.
var MulticallembeddedMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MulticallembeddedABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallembeddedMetaData.ABI instead.
var MulticallembeddedABI = MulticallembeddedMetaData.ABI

// Multicallembedded is an auto generated Go binding around an Ethereum contract.
type Multicallembedded struct {
	MulticallembeddedCaller     // Read-only binding to the contract
	MulticallembeddedTransactor // Write-only binding to the contract
	MulticallembeddedFilterer   // Log filterer for contract events
}

// MulticallembeddedCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallembeddedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallembeddedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallembeddedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallembeddedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallembeddedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallembeddedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallembeddedSession struct {
	Contract     *Multicallembedded // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MulticallembeddedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallembeddedCallerSession struct {
	Contract *MulticallembeddedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MulticallembeddedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallembeddedTransactorSession struct {
	Contract     *MulticallembeddedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MulticallembeddedRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallembeddedRaw struct {
	Contract *Multicallembedded // Generic contract binding to access the raw methods on
}

// MulticallembeddedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallembeddedCallerRaw struct {
	Contract *MulticallembeddedCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallembeddedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallembeddedTransactorRaw struct {
	Contract *MulticallembeddedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallembedded creates a new instance of Multicallembedded, bound to a specific deployed contract.
func NewMulticallembedded(address common.Address, backend bind.ContractBackend) (*Multicallembedded, error) {
	contract, err := bindMulticallembedded(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicallembedded{MulticallembeddedCaller: MulticallembeddedCaller{contract: contract}, MulticallembeddedTransactor: MulticallembeddedTransactor{contract: contract}, MulticallembeddedFilterer: MulticallembeddedFilterer{contract: contract}}, nil
}

// NewMulticallembeddedCaller creates a new read-only instance of Multicallembedded, bound to a specific deployed contract.
func NewMulticallembeddedCaller(address common.Address, caller bind.ContractCaller) (*MulticallembeddedCaller, error) {
	contract, err := bindMulticallembedded(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallembeddedCaller{contract: contract}, nil
}

// NewMulticallembeddedTransactor creates a new write-only instance of Multicallembedded, bound to a specific deployed contract.
func NewMulticallembeddedTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallembeddedTransactor, error) {
	contract, err := bindMulticallembedded(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallembeddedTransactor{contract: contract}, nil
}

// NewMulticallembeddedFilterer creates a new log filterer instance of Multicallembedded, bound to a specific deployed contract.
func NewMulticallembeddedFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallembeddedFilterer, error) {
	contract, err := bindMulticallembedded(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallembeddedFilterer{contract: contract}, nil
}

// bindMulticallembedded binds a generic wrapper to an already deployed contract.
func bindMulticallembedded(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallembeddedMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicallembedded *MulticallembeddedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicallembedded.Contract.MulticallembeddedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicallembedded *MulticallembeddedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicallembedded.Contract.MulticallembeddedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicallembedded *MulticallembeddedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicallembedded.Contract.MulticallembeddedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicallembedded *MulticallembeddedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicallembedded.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicallembedded *MulticallembeddedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicallembedded.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicallembedded *MulticallembeddedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicallembedded.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multicallembedded *MulticallembeddedTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Multicallembedded.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multicallembedded *MulticallembeddedSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicallembedded.Contract.Multicall(&_Multicallembedded.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multicallembedded *MulticallembeddedTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multicallembedded.Contract.Multicall(&_Multicallembedded.TransactOpts, data)
}
