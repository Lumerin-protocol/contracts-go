// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package faucet

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

// FaucetMetaData contains all meta data concerning the Faucet contract.
var FaucetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmr\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyLimitCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emptyGeth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gethAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTransferLumerin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cooldownPeriod\",\"type\":\"uint256\"}],\"name\":\"setUpdateCooldownPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gwei\",\"type\":\"uint256\"}],\"name\":\"setUpdateGWEIAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmr\",\"type\":\"address\"}],\"name\":\"setUpdateLumerin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setUpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_txAmount\",\"type\":\"uint256\"}],\"name\":\"setUpdateTxAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startOfDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_claiment\",\"type\":\"address\"}],\"name\":\"supervisedClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"txAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// FaucetABI is the input ABI used to generate the binding from.
// Deprecated: Use FaucetMetaData.ABI instead.
var FaucetABI = FaucetMetaData.ABI

// Faucet is an auto generated Go binding around an Ethereum contract.
type Faucet struct {
	FaucetCaller     // Read-only binding to the contract
	FaucetTransactor // Write-only binding to the contract
	FaucetFilterer   // Log filterer for contract events
}

// FaucetCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaucetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaucetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaucetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaucetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaucetSession struct {
	Contract     *Faucet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaucetCallerSession struct {
	Contract *FaucetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FaucetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaucetTransactorSession struct {
	Contract     *FaucetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaucetRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaucetRaw struct {
	Contract *Faucet // Generic contract binding to access the raw methods on
}

// FaucetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaucetCallerRaw struct {
	Contract *FaucetCaller // Generic read-only contract binding to access the raw methods on
}

// FaucetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaucetTransactorRaw struct {
	Contract *FaucetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaucet creates a new instance of Faucet, bound to a specific deployed contract.
func NewFaucet(address common.Address, backend bind.ContractBackend) (*Faucet, error) {
	contract, err := bindFaucet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Faucet{FaucetCaller: FaucetCaller{contract: contract}, FaucetTransactor: FaucetTransactor{contract: contract}, FaucetFilterer: FaucetFilterer{contract: contract}}, nil
}

// NewFaucetCaller creates a new read-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetCaller(address common.Address, caller bind.ContractCaller) (*FaucetCaller, error) {
	contract, err := bindFaucet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetCaller{contract: contract}, nil
}

// NewFaucetTransactor creates a new write-only instance of Faucet, bound to a specific deployed contract.
func NewFaucetTransactor(address common.Address, transactor bind.ContractTransactor) (*FaucetTransactor, error) {
	contract, err := bindFaucet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaucetTransactor{contract: contract}, nil
}

// NewFaucetFilterer creates a new log filterer instance of Faucet, bound to a specific deployed contract.
func NewFaucetFilterer(address common.Address, filterer bind.ContractFilterer) (*FaucetFilterer, error) {
	contract, err := bindFaucet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaucetFilterer{contract: contract}, nil
}

// bindFaucet binds a generic wrapper to an already deployed contract.
func bindFaucet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FaucetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.FaucetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.FaucetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Faucet *FaucetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Faucet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Faucet *FaucetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Faucet *FaucetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Faucet.Contract.contract.Transact(opts, method, params...)
}

// CooldownPeriod is a free data retrieval call binding the contract method 0x04646a49.
//
// Solidity: function cooldownPeriod() view returns(uint256)
func (_Faucet *FaucetCaller) CooldownPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "cooldownPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownPeriod is a free data retrieval call binding the contract method 0x04646a49.
//
// Solidity: function cooldownPeriod() view returns(uint256)
func (_Faucet *FaucetSession) CooldownPeriod() (*big.Int, error) {
	return _Faucet.Contract.CooldownPeriod(&_Faucet.CallOpts)
}

// CooldownPeriod is a free data retrieval call binding the contract method 0x04646a49.
//
// Solidity: function cooldownPeriod() view returns(uint256)
func (_Faucet *FaucetCallerSession) CooldownPeriod() (*big.Int, error) {
	return _Faucet.Contract.CooldownPeriod(&_Faucet.CallOpts)
}

// DailyLimitCount is a free data retrieval call binding the contract method 0x12b3c5b4.
//
// Solidity: function dailyLimitCount() view returns(uint256)
func (_Faucet *FaucetCaller) DailyLimitCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "dailyLimitCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyLimitCount is a free data retrieval call binding the contract method 0x12b3c5b4.
//
// Solidity: function dailyLimitCount() view returns(uint256)
func (_Faucet *FaucetSession) DailyLimitCount() (*big.Int, error) {
	return _Faucet.Contract.DailyLimitCount(&_Faucet.CallOpts)
}

// DailyLimitCount is a free data retrieval call binding the contract method 0x12b3c5b4.
//
// Solidity: function dailyLimitCount() view returns(uint256)
func (_Faucet *FaucetCallerSession) DailyLimitCount() (*big.Int, error) {
	return _Faucet.Contract.DailyLimitCount(&_Faucet.CallOpts)
}

// GethAmount is a free data retrieval call binding the contract method 0xda294632.
//
// Solidity: function gethAmount() view returns(uint256)
func (_Faucet *FaucetCaller) GethAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "gethAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GethAmount is a free data retrieval call binding the contract method 0xda294632.
//
// Solidity: function gethAmount() view returns(uint256)
func (_Faucet *FaucetSession) GethAmount() (*big.Int, error) {
	return _Faucet.Contract.GethAmount(&_Faucet.CallOpts)
}

// GethAmount is a free data retrieval call binding the contract method 0xda294632.
//
// Solidity: function gethAmount() view returns(uint256)
func (_Faucet *FaucetCallerSession) GethAmount() (*big.Int, error) {
	return _Faucet.Contract.GethAmount(&_Faucet.CallOpts)
}

// StartOfDay is a free data retrieval call binding the contract method 0x57f84324.
//
// Solidity: function startOfDay() view returns(uint256)
func (_Faucet *FaucetCaller) StartOfDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "startOfDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartOfDay is a free data retrieval call binding the contract method 0x57f84324.
//
// Solidity: function startOfDay() view returns(uint256)
func (_Faucet *FaucetSession) StartOfDay() (*big.Int, error) {
	return _Faucet.Contract.StartOfDay(&_Faucet.CallOpts)
}

// StartOfDay is a free data retrieval call binding the contract method 0x57f84324.
//
// Solidity: function startOfDay() view returns(uint256)
func (_Faucet *FaucetCallerSession) StartOfDay() (*big.Int, error) {
	return _Faucet.Contract.StartOfDay(&_Faucet.CallOpts)
}

// TxAmount is a free data retrieval call binding the contract method 0xa0c10c2d.
//
// Solidity: function txAmount() view returns(uint256)
func (_Faucet *FaucetCaller) TxAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "txAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TxAmount is a free data retrieval call binding the contract method 0xa0c10c2d.
//
// Solidity: function txAmount() view returns(uint256)
func (_Faucet *FaucetSession) TxAmount() (*big.Int, error) {
	return _Faucet.Contract.TxAmount(&_Faucet.CallOpts)
}

// TxAmount is a free data retrieval call binding the contract method 0xa0c10c2d.
//
// Solidity: function txAmount() view returns(uint256)
func (_Faucet *FaucetCallerSession) TxAmount() (*big.Int, error) {
	return _Faucet.Contract.TxAmount(&_Faucet.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Faucet *FaucetTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Faucet *FaucetSession) Claim() (*types.Transaction, error) {
	return _Faucet.Contract.Claim(&_Faucet.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Faucet *FaucetTransactorSession) Claim() (*types.Transaction, error) {
	return _Faucet.Contract.Claim(&_Faucet.TransactOpts)
}

// EmptyGeth is a paid mutator transaction binding the contract method 0xb2f4e93d.
//
// Solidity: function emptyGeth() returns()
func (_Faucet *FaucetTransactor) EmptyGeth(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "emptyGeth")
}

// EmptyGeth is a paid mutator transaction binding the contract method 0xb2f4e93d.
//
// Solidity: function emptyGeth() returns()
func (_Faucet *FaucetSession) EmptyGeth() (*types.Transaction, error) {
	return _Faucet.Contract.EmptyGeth(&_Faucet.TransactOpts)
}

// EmptyGeth is a paid mutator transaction binding the contract method 0xb2f4e93d.
//
// Solidity: function emptyGeth() returns()
func (_Faucet *FaucetTransactorSession) EmptyGeth() (*types.Transaction, error) {
	return _Faucet.Contract.EmptyGeth(&_Faucet.TransactOpts)
}

// SetTransferLumerin is a paid mutator transaction binding the contract method 0xc612f0ab.
//
// Solidity: function setTransferLumerin(address _to, uint256 _amount) returns()
func (_Faucet *FaucetTransactor) SetTransferLumerin(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setTransferLumerin", _to, _amount)
}

// SetTransferLumerin is a paid mutator transaction binding the contract method 0xc612f0ab.
//
// Solidity: function setTransferLumerin(address _to, uint256 _amount) returns()
func (_Faucet *FaucetSession) SetTransferLumerin(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetTransferLumerin(&_Faucet.TransactOpts, _to, _amount)
}

// SetTransferLumerin is a paid mutator transaction binding the contract method 0xc612f0ab.
//
// Solidity: function setTransferLumerin(address _to, uint256 _amount) returns()
func (_Faucet *FaucetTransactorSession) SetTransferLumerin(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetTransferLumerin(&_Faucet.TransactOpts, _to, _amount)
}

// SetUpdateCooldownPeriod is a paid mutator transaction binding the contract method 0xe8f25899.
//
// Solidity: function setUpdateCooldownPeriod(uint256 _cooldownPeriod) returns()
func (_Faucet *FaucetTransactor) SetUpdateCooldownPeriod(opts *bind.TransactOpts, _cooldownPeriod *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdateCooldownPeriod", _cooldownPeriod)
}

// SetUpdateCooldownPeriod is a paid mutator transaction binding the contract method 0xe8f25899.
//
// Solidity: function setUpdateCooldownPeriod(uint256 _cooldownPeriod) returns()
func (_Faucet *FaucetSession) SetUpdateCooldownPeriod(_cooldownPeriod *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateCooldownPeriod(&_Faucet.TransactOpts, _cooldownPeriod)
}

// SetUpdateCooldownPeriod is a paid mutator transaction binding the contract method 0xe8f25899.
//
// Solidity: function setUpdateCooldownPeriod(uint256 _cooldownPeriod) returns()
func (_Faucet *FaucetTransactorSession) SetUpdateCooldownPeriod(_cooldownPeriod *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateCooldownPeriod(&_Faucet.TransactOpts, _cooldownPeriod)
}

// SetUpdateGWEIAmount is a paid mutator transaction binding the contract method 0xf9c167bf.
//
// Solidity: function setUpdateGWEIAmount(uint256 _gwei) returns()
func (_Faucet *FaucetTransactor) SetUpdateGWEIAmount(opts *bind.TransactOpts, _gwei *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdateGWEIAmount", _gwei)
}

// SetUpdateGWEIAmount is a paid mutator transaction binding the contract method 0xf9c167bf.
//
// Solidity: function setUpdateGWEIAmount(uint256 _gwei) returns()
func (_Faucet *FaucetSession) SetUpdateGWEIAmount(_gwei *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateGWEIAmount(&_Faucet.TransactOpts, _gwei)
}

// SetUpdateGWEIAmount is a paid mutator transaction binding the contract method 0xf9c167bf.
//
// Solidity: function setUpdateGWEIAmount(uint256 _gwei) returns()
func (_Faucet *FaucetTransactorSession) SetUpdateGWEIAmount(_gwei *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateGWEIAmount(&_Faucet.TransactOpts, _gwei)
}

// SetUpdateLumerin is a paid mutator transaction binding the contract method 0x26fd8f62.
//
// Solidity: function setUpdateLumerin(address _lmr) returns()
func (_Faucet *FaucetTransactor) SetUpdateLumerin(opts *bind.TransactOpts, _lmr common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdateLumerin", _lmr)
}

// SetUpdateLumerin is a paid mutator transaction binding the contract method 0x26fd8f62.
//
// Solidity: function setUpdateLumerin(address _lmr) returns()
func (_Faucet *FaucetSession) SetUpdateLumerin(_lmr common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateLumerin(&_Faucet.TransactOpts, _lmr)
}

// SetUpdateLumerin is a paid mutator transaction binding the contract method 0x26fd8f62.
//
// Solidity: function setUpdateLumerin(address _lmr) returns()
func (_Faucet *FaucetTransactorSession) SetUpdateLumerin(_lmr common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateLumerin(&_Faucet.TransactOpts, _lmr)
}

// SetUpdateOwner is a paid mutator transaction binding the contract method 0x49ba1071.
//
// Solidity: function setUpdateOwner(address _newOwner) returns()
func (_Faucet *FaucetTransactor) SetUpdateOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdateOwner", _newOwner)
}

// SetUpdateOwner is a paid mutator transaction binding the contract method 0x49ba1071.
//
// Solidity: function setUpdateOwner(address _newOwner) returns()
func (_Faucet *FaucetSession) SetUpdateOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateOwner(&_Faucet.TransactOpts, _newOwner)
}

// SetUpdateOwner is a paid mutator transaction binding the contract method 0x49ba1071.
//
// Solidity: function setUpdateOwner(address _newOwner) returns()
func (_Faucet *FaucetTransactorSession) SetUpdateOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateOwner(&_Faucet.TransactOpts, _newOwner)
}

// SetUpdateTxAmount is a paid mutator transaction binding the contract method 0x0a3a4f7e.
//
// Solidity: function setUpdateTxAmount(uint256 _txAmount) returns()
func (_Faucet *FaucetTransactor) SetUpdateTxAmount(opts *bind.TransactOpts, _txAmount *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdateTxAmount", _txAmount)
}

// SetUpdateTxAmount is a paid mutator transaction binding the contract method 0x0a3a4f7e.
//
// Solidity: function setUpdateTxAmount(uint256 _txAmount) returns()
func (_Faucet *FaucetSession) SetUpdateTxAmount(_txAmount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateTxAmount(&_Faucet.TransactOpts, _txAmount)
}

// SetUpdateTxAmount is a paid mutator transaction binding the contract method 0x0a3a4f7e.
//
// Solidity: function setUpdateTxAmount(uint256 _txAmount) returns()
func (_Faucet *FaucetTransactorSession) SetUpdateTxAmount(_txAmount *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateTxAmount(&_Faucet.TransactOpts, _txAmount)
}

// SupervisedClaim is a paid mutator transaction binding the contract method 0x8885173a.
//
// Solidity: function supervisedClaim(address _claiment) returns()
func (_Faucet *FaucetTransactor) SupervisedClaim(opts *bind.TransactOpts, _claiment common.Address) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "supervisedClaim", _claiment)
}

// SupervisedClaim is a paid mutator transaction binding the contract method 0x8885173a.
//
// Solidity: function supervisedClaim(address _claiment) returns()
func (_Faucet *FaucetSession) SupervisedClaim(_claiment common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SupervisedClaim(&_Faucet.TransactOpts, _claiment)
}

// SupervisedClaim is a paid mutator transaction binding the contract method 0x8885173a.
//
// Solidity: function supervisedClaim(address _claiment) returns()
func (_Faucet *FaucetTransactorSession) SupervisedClaim(_claiment common.Address) (*types.Transaction, error) {
	return _Faucet.Contract.SupervisedClaim(&_Faucet.TransactOpts, _claiment)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetSession) Receive() (*types.Transaction, error) {
	return _Faucet.Contract.Receive(&_Faucet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Faucet *FaucetTransactorSession) Receive() (*types.Transaction, error) {
	return _Faucet.Contract.Receive(&_Faucet.TransactOpts)
}
