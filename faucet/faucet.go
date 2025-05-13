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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dailyMaxLmr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lmrPayout\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethPayout\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ipAddress\",\"type\":\"string\"}],\"name\":\"canClaimTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownStartingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentLMRTokenDistribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emptyGeth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethPayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lmrPayout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lmrTokenDistributionMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetDistributedTodayLmr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTransferLumerin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ethPayout\",\"type\":\"uint256\"}],\"name\":\"setUpdateEthPayout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lmrPayout\",\"type\":\"uint256\"}],\"name\":\"setUpdateLmrPayout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmr\",\"type\":\"address\"}],\"name\":\"setUpdateLumerin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setUpdateOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dailyMaxLmr\",\"type\":\"uint256\"}],\"name\":\"setUpdatedailyMaxLmr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_claiment\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_ipAddress\",\"type\":\"string\"}],\"name\":\"supervisedClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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

// CanClaimTokens is a free data retrieval call binding the contract method 0x73e1ec86.
//
// Solidity: function canClaimTokens(address _address, string _ipAddress) view returns(bool)
func (_Faucet *FaucetCaller) CanClaimTokens(opts *bind.CallOpts, _address common.Address, _ipAddress string) (bool, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "canClaimTokens", _address, _ipAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaimTokens is a free data retrieval call binding the contract method 0x73e1ec86.
//
// Solidity: function canClaimTokens(address _address, string _ipAddress) view returns(bool)
func (_Faucet *FaucetSession) CanClaimTokens(_address common.Address, _ipAddress string) (bool, error) {
	return _Faucet.Contract.CanClaimTokens(&_Faucet.CallOpts, _address, _ipAddress)
}

// CanClaimTokens is a free data retrieval call binding the contract method 0x73e1ec86.
//
// Solidity: function canClaimTokens(address _address, string _ipAddress) view returns(bool)
func (_Faucet *FaucetCallerSession) CanClaimTokens(_address common.Address, _ipAddress string) (bool, error) {
	return _Faucet.Contract.CanClaimTokens(&_Faucet.CallOpts, _address, _ipAddress)
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

// CooldownStartingTime is a free data retrieval call binding the contract method 0xd882da79.
//
// Solidity: function cooldownStartingTime() view returns(uint256)
func (_Faucet *FaucetCaller) CooldownStartingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "cooldownStartingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownStartingTime is a free data retrieval call binding the contract method 0xd882da79.
//
// Solidity: function cooldownStartingTime() view returns(uint256)
func (_Faucet *FaucetSession) CooldownStartingTime() (*big.Int, error) {
	return _Faucet.Contract.CooldownStartingTime(&_Faucet.CallOpts)
}

// CooldownStartingTime is a free data retrieval call binding the contract method 0xd882da79.
//
// Solidity: function cooldownStartingTime() view returns(uint256)
func (_Faucet *FaucetCallerSession) CooldownStartingTime() (*big.Int, error) {
	return _Faucet.Contract.CooldownStartingTime(&_Faucet.CallOpts)
}

// CurrentLMRTokenDistribution is a free data retrieval call binding the contract method 0x0539c94f.
//
// Solidity: function currentLMRTokenDistribution() view returns(uint256)
func (_Faucet *FaucetCaller) CurrentLMRTokenDistribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "currentLMRTokenDistribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentLMRTokenDistribution is a free data retrieval call binding the contract method 0x0539c94f.
//
// Solidity: function currentLMRTokenDistribution() view returns(uint256)
func (_Faucet *FaucetSession) CurrentLMRTokenDistribution() (*big.Int, error) {
	return _Faucet.Contract.CurrentLMRTokenDistribution(&_Faucet.CallOpts)
}

// CurrentLMRTokenDistribution is a free data retrieval call binding the contract method 0x0539c94f.
//
// Solidity: function currentLMRTokenDistribution() view returns(uint256)
func (_Faucet *FaucetCallerSession) CurrentLMRTokenDistribution() (*big.Int, error) {
	return _Faucet.Contract.CurrentLMRTokenDistribution(&_Faucet.CallOpts)
}

// EthPayout is a free data retrieval call binding the contract method 0x0a213e04.
//
// Solidity: function ethPayout() view returns(uint256)
func (_Faucet *FaucetCaller) EthPayout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "ethPayout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthPayout is a free data retrieval call binding the contract method 0x0a213e04.
//
// Solidity: function ethPayout() view returns(uint256)
func (_Faucet *FaucetSession) EthPayout() (*big.Int, error) {
	return _Faucet.Contract.EthPayout(&_Faucet.CallOpts)
}

// EthPayout is a free data retrieval call binding the contract method 0x0a213e04.
//
// Solidity: function ethPayout() view returns(uint256)
func (_Faucet *FaucetCallerSession) EthPayout() (*big.Int, error) {
	return _Faucet.Contract.EthPayout(&_Faucet.CallOpts)
}

// LmrPayout is a free data retrieval call binding the contract method 0x5ce5b028.
//
// Solidity: function lmrPayout() view returns(uint256)
func (_Faucet *FaucetCaller) LmrPayout(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "lmrPayout")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LmrPayout is a free data retrieval call binding the contract method 0x5ce5b028.
//
// Solidity: function lmrPayout() view returns(uint256)
func (_Faucet *FaucetSession) LmrPayout() (*big.Int, error) {
	return _Faucet.Contract.LmrPayout(&_Faucet.CallOpts)
}

// LmrPayout is a free data retrieval call binding the contract method 0x5ce5b028.
//
// Solidity: function lmrPayout() view returns(uint256)
func (_Faucet *FaucetCallerSession) LmrPayout() (*big.Int, error) {
	return _Faucet.Contract.LmrPayout(&_Faucet.CallOpts)
}

// LmrTokenDistributionMax is a free data retrieval call binding the contract method 0x5c394f2e.
//
// Solidity: function lmrTokenDistributionMax() view returns(uint256)
func (_Faucet *FaucetCaller) LmrTokenDistributionMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Faucet.contract.Call(opts, &out, "lmrTokenDistributionMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LmrTokenDistributionMax is a free data retrieval call binding the contract method 0x5c394f2e.
//
// Solidity: function lmrTokenDistributionMax() view returns(uint256)
func (_Faucet *FaucetSession) LmrTokenDistributionMax() (*big.Int, error) {
	return _Faucet.Contract.LmrTokenDistributionMax(&_Faucet.CallOpts)
}

// LmrTokenDistributionMax is a free data retrieval call binding the contract method 0x5c394f2e.
//
// Solidity: function lmrTokenDistributionMax() view returns(uint256)
func (_Faucet *FaucetCallerSession) LmrTokenDistributionMax() (*big.Int, error) {
	return _Faucet.Contract.LmrTokenDistributionMax(&_Faucet.CallOpts)
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

// ResetDistributedTodayLmr is a paid mutator transaction binding the contract method 0x2e30835f.
//
// Solidity: function resetDistributedTodayLmr() returns()
func (_Faucet *FaucetTransactor) ResetDistributedTodayLmr(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "resetDistributedTodayLmr")
}

// ResetDistributedTodayLmr is a paid mutator transaction binding the contract method 0x2e30835f.
//
// Solidity: function resetDistributedTodayLmr() returns()
func (_Faucet *FaucetSession) ResetDistributedTodayLmr() (*types.Transaction, error) {
	return _Faucet.Contract.ResetDistributedTodayLmr(&_Faucet.TransactOpts)
}

// ResetDistributedTodayLmr is a paid mutator transaction binding the contract method 0x2e30835f.
//
// Solidity: function resetDistributedTodayLmr() returns()
func (_Faucet *FaucetTransactorSession) ResetDistributedTodayLmr() (*types.Transaction, error) {
	return _Faucet.Contract.ResetDistributedTodayLmr(&_Faucet.TransactOpts)
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

// SetUpdateEthPayout is a paid mutator transaction binding the contract method 0x5022aa97.
//
// Solidity: function setUpdateEthPayout(uint256 _ethPayout) returns()
func (_Faucet *FaucetTransactor) SetUpdateEthPayout(opts *bind.TransactOpts, _ethPayout *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdateEthPayout", _ethPayout)
}

// SetUpdateEthPayout is a paid mutator transaction binding the contract method 0x5022aa97.
//
// Solidity: function setUpdateEthPayout(uint256 _ethPayout) returns()
func (_Faucet *FaucetSession) SetUpdateEthPayout(_ethPayout *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateEthPayout(&_Faucet.TransactOpts, _ethPayout)
}

// SetUpdateEthPayout is a paid mutator transaction binding the contract method 0x5022aa97.
//
// Solidity: function setUpdateEthPayout(uint256 _ethPayout) returns()
func (_Faucet *FaucetTransactorSession) SetUpdateEthPayout(_ethPayout *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateEthPayout(&_Faucet.TransactOpts, _ethPayout)
}

// SetUpdateLmrPayout is a paid mutator transaction binding the contract method 0x198e3885.
//
// Solidity: function setUpdateLmrPayout(uint256 _lmrPayout) returns()
func (_Faucet *FaucetTransactor) SetUpdateLmrPayout(opts *bind.TransactOpts, _lmrPayout *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdateLmrPayout", _lmrPayout)
}

// SetUpdateLmrPayout is a paid mutator transaction binding the contract method 0x198e3885.
//
// Solidity: function setUpdateLmrPayout(uint256 _lmrPayout) returns()
func (_Faucet *FaucetSession) SetUpdateLmrPayout(_lmrPayout *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateLmrPayout(&_Faucet.TransactOpts, _lmrPayout)
}

// SetUpdateLmrPayout is a paid mutator transaction binding the contract method 0x198e3885.
//
// Solidity: function setUpdateLmrPayout(uint256 _lmrPayout) returns()
func (_Faucet *FaucetTransactorSession) SetUpdateLmrPayout(_lmrPayout *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdateLmrPayout(&_Faucet.TransactOpts, _lmrPayout)
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

// SetUpdatedailyMaxLmr is a paid mutator transaction binding the contract method 0xcd05b8aa.
//
// Solidity: function setUpdatedailyMaxLmr(uint256 _dailyMaxLmr) returns()
func (_Faucet *FaucetTransactor) SetUpdatedailyMaxLmr(opts *bind.TransactOpts, _dailyMaxLmr *big.Int) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "setUpdatedailyMaxLmr", _dailyMaxLmr)
}

// SetUpdatedailyMaxLmr is a paid mutator transaction binding the contract method 0xcd05b8aa.
//
// Solidity: function setUpdatedailyMaxLmr(uint256 _dailyMaxLmr) returns()
func (_Faucet *FaucetSession) SetUpdatedailyMaxLmr(_dailyMaxLmr *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdatedailyMaxLmr(&_Faucet.TransactOpts, _dailyMaxLmr)
}

// SetUpdatedailyMaxLmr is a paid mutator transaction binding the contract method 0xcd05b8aa.
//
// Solidity: function setUpdatedailyMaxLmr(uint256 _dailyMaxLmr) returns()
func (_Faucet *FaucetTransactorSession) SetUpdatedailyMaxLmr(_dailyMaxLmr *big.Int) (*types.Transaction, error) {
	return _Faucet.Contract.SetUpdatedailyMaxLmr(&_Faucet.TransactOpts, _dailyMaxLmr)
}

// SupervisedClaim is a paid mutator transaction binding the contract method 0x7253e01b.
//
// Solidity: function supervisedClaim(address _claiment, string _ipAddress) returns()
func (_Faucet *FaucetTransactor) SupervisedClaim(opts *bind.TransactOpts, _claiment common.Address, _ipAddress string) (*types.Transaction, error) {
	return _Faucet.contract.Transact(opts, "supervisedClaim", _claiment, _ipAddress)
}

// SupervisedClaim is a paid mutator transaction binding the contract method 0x7253e01b.
//
// Solidity: function supervisedClaim(address _claiment, string _ipAddress) returns()
func (_Faucet *FaucetSession) SupervisedClaim(_claiment common.Address, _ipAddress string) (*types.Transaction, error) {
	return _Faucet.Contract.SupervisedClaim(&_Faucet.TransactOpts, _claiment, _ipAddress)
}

// SupervisedClaim is a paid mutator transaction binding the contract method 0x7253e01b.
//
// Solidity: function supervisedClaim(address _claiment, string _ipAddress) returns()
func (_Faucet *FaucetTransactorSession) SupervisedClaim(_claiment common.Address, _ipAddress string) (*types.Transaction, error) {
	return _Faucet.Contract.SupervisedClaim(&_Faucet.TransactOpts, _claiment, _ipAddress)
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
