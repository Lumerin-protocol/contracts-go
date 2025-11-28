// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregatorv3interface

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

// Aggregatorv3interfaceMetaData contains all meta data concerning the Aggregatorv3interface contract.
var Aggregatorv3interfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Aggregatorv3interfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use Aggregatorv3interfaceMetaData.ABI instead.
var Aggregatorv3interfaceABI = Aggregatorv3interfaceMetaData.ABI

// Aggregatorv3interface is an auto generated Go binding around an Ethereum contract.
type Aggregatorv3interface struct {
	Aggregatorv3interfaceCaller     // Read-only binding to the contract
	Aggregatorv3interfaceTransactor // Write-only binding to the contract
	Aggregatorv3interfaceFilterer   // Log filterer for contract events
}

// Aggregatorv3interfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type Aggregatorv3interfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aggregatorv3interfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Aggregatorv3interfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aggregatorv3interfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Aggregatorv3interfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Aggregatorv3interfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Aggregatorv3interfaceSession struct {
	Contract     *Aggregatorv3interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// Aggregatorv3interfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Aggregatorv3interfaceCallerSession struct {
	Contract *Aggregatorv3interfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// Aggregatorv3interfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Aggregatorv3interfaceTransactorSession struct {
	Contract     *Aggregatorv3interfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// Aggregatorv3interfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type Aggregatorv3interfaceRaw struct {
	Contract *Aggregatorv3interface // Generic contract binding to access the raw methods on
}

// Aggregatorv3interfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Aggregatorv3interfaceCallerRaw struct {
	Contract *Aggregatorv3interfaceCaller // Generic read-only contract binding to access the raw methods on
}

// Aggregatorv3interfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Aggregatorv3interfaceTransactorRaw struct {
	Contract *Aggregatorv3interfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatorv3interface creates a new instance of Aggregatorv3interface, bound to a specific deployed contract.
func NewAggregatorv3interface(address common.Address, backend bind.ContractBackend) (*Aggregatorv3interface, error) {
	contract, err := bindAggregatorv3interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggregatorv3interface{Aggregatorv3interfaceCaller: Aggregatorv3interfaceCaller{contract: contract}, Aggregatorv3interfaceTransactor: Aggregatorv3interfaceTransactor{contract: contract}, Aggregatorv3interfaceFilterer: Aggregatorv3interfaceFilterer{contract: contract}}, nil
}

// NewAggregatorv3interfaceCaller creates a new read-only instance of Aggregatorv3interface, bound to a specific deployed contract.
func NewAggregatorv3interfaceCaller(address common.Address, caller bind.ContractCaller) (*Aggregatorv3interfaceCaller, error) {
	contract, err := bindAggregatorv3interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Aggregatorv3interfaceCaller{contract: contract}, nil
}

// NewAggregatorv3interfaceTransactor creates a new write-only instance of Aggregatorv3interface, bound to a specific deployed contract.
func NewAggregatorv3interfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*Aggregatorv3interfaceTransactor, error) {
	contract, err := bindAggregatorv3interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Aggregatorv3interfaceTransactor{contract: contract}, nil
}

// NewAggregatorv3interfaceFilterer creates a new log filterer instance of Aggregatorv3interface, bound to a specific deployed contract.
func NewAggregatorv3interfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*Aggregatorv3interfaceFilterer, error) {
	contract, err := bindAggregatorv3interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Aggregatorv3interfaceFilterer{contract: contract}, nil
}

// bindAggregatorv3interface binds a generic wrapper to an already deployed contract.
func bindAggregatorv3interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Aggregatorv3interfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregatorv3interface *Aggregatorv3interfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregatorv3interface.Contract.Aggregatorv3interfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregatorv3interface *Aggregatorv3interfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregatorv3interface.Contract.Aggregatorv3interfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregatorv3interface *Aggregatorv3interfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregatorv3interface.Contract.Aggregatorv3interfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregatorv3interface *Aggregatorv3interfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregatorv3interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregatorv3interface *Aggregatorv3interfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregatorv3interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregatorv3interface *Aggregatorv3interfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregatorv3interface.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Aggregatorv3interface *Aggregatorv3interfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Aggregatorv3interface.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Aggregatorv3interface *Aggregatorv3interfaceSession) Decimals() (uint8, error) {
	return _Aggregatorv3interface.Contract.Decimals(&_Aggregatorv3interface.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Aggregatorv3interface *Aggregatorv3interfaceCallerSession) Decimals() (uint8, error) {
	return _Aggregatorv3interface.Contract.Decimals(&_Aggregatorv3interface.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Aggregatorv3interface *Aggregatorv3interfaceCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggregatorv3interface.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Aggregatorv3interface *Aggregatorv3interfaceSession) Description() (string, error) {
	return _Aggregatorv3interface.Contract.Description(&_Aggregatorv3interface.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Aggregatorv3interface *Aggregatorv3interfaceCallerSession) Description() (string, error) {
	return _Aggregatorv3interface.Contract.Description(&_Aggregatorv3interface.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregatorv3interface *Aggregatorv3interfaceCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Aggregatorv3interface.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregatorv3interface *Aggregatorv3interfaceSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregatorv3interface.Contract.GetRoundData(&_Aggregatorv3interface.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregatorv3interface *Aggregatorv3interfaceCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregatorv3interface.Contract.GetRoundData(&_Aggregatorv3interface.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregatorv3interface *Aggregatorv3interfaceCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Aggregatorv3interface.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregatorv3interface *Aggregatorv3interfaceSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregatorv3interface.Contract.LatestRoundData(&_Aggregatorv3interface.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregatorv3interface *Aggregatorv3interfaceCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregatorv3interface.Contract.LatestRoundData(&_Aggregatorv3interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Aggregatorv3interface *Aggregatorv3interfaceCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregatorv3interface.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Aggregatorv3interface *Aggregatorv3interfaceSession) Version() (*big.Int, error) {
	return _Aggregatorv3interface.Contract.Version(&_Aggregatorv3interface.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Aggregatorv3interface *Aggregatorv3interfaceCallerSession) Version() (*big.Int, error) {
	return _Aggregatorv3interface.Contract.Version(&_Aggregatorv3interface.CallOpts)
}
