// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package clonefactory

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

// ClonefactoryMetaData contains all meta data concerning the Clonefactory contract.
var ClonefactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"clonefactoryContractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"contractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"}],\"name\":\"contractDeleteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"purchaseInfoUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"baseImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"checkWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lumerin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isContractDead\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lumerin\",\"outputs\":[{\"internalType\":\"contractLumerin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noMoreWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payMarketplaceFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rentalContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"}],\"name\":\"setContractDeleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pubKey\",\"type\":\"string\"}],\"name\":\"setCreateNewRentalContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pubKey\",\"type\":\"string\"}],\"name\":\"setCreateNewRentalContractV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setDisableWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"setMarketplaceFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_cipherText\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"termsVersion\",\"type\":\"uint32\"}],\"name\":\"setPurchaseRentalContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"termsVersion\",\"type\":\"uint32\"}],\"name\":\"setPurchaseRentalContractV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRemoveFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setUpdateContractInformation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"name\":\"setUpdateContractInformationV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ClonefactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ClonefactoryMetaData.ABI instead.
var ClonefactoryABI = ClonefactoryMetaData.ABI

// Clonefactory is an auto generated Go binding around an Ethereum contract.
type Clonefactory struct {
	ClonefactoryCaller     // Read-only binding to the contract
	ClonefactoryTransactor // Write-only binding to the contract
	ClonefactoryFilterer   // Log filterer for contract events
}

// ClonefactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonefactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonefactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonefactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonefactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonefactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonefactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonefactorySession struct {
	Contract     *Clonefactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonefactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonefactoryCallerSession struct {
	Contract *ClonefactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ClonefactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonefactoryTransactorSession struct {
	Contract     *ClonefactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ClonefactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonefactoryRaw struct {
	Contract *Clonefactory // Generic contract binding to access the raw methods on
}

// ClonefactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonefactoryCallerRaw struct {
	Contract *ClonefactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ClonefactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonefactoryTransactorRaw struct {
	Contract *ClonefactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClonefactory creates a new instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactory(address common.Address, backend bind.ContractBackend) (*Clonefactory, error) {
	contract, err := bindClonefactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clonefactory{ClonefactoryCaller: ClonefactoryCaller{contract: contract}, ClonefactoryTransactor: ClonefactoryTransactor{contract: contract}, ClonefactoryFilterer: ClonefactoryFilterer{contract: contract}}, nil
}

// NewClonefactoryCaller creates a new read-only instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactoryCaller(address common.Address, caller bind.ContractCaller) (*ClonefactoryCaller, error) {
	contract, err := bindClonefactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryCaller{contract: contract}, nil
}

// NewClonefactoryTransactor creates a new write-only instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonefactoryTransactor, error) {
	contract, err := bindClonefactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryTransactor{contract: contract}, nil
}

// NewClonefactoryFilterer creates a new log filterer instance of Clonefactory, bound to a specific deployed contract.
func NewClonefactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonefactoryFilterer, error) {
	contract, err := bindClonefactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryFilterer{contract: contract}, nil
}

// bindClonefactory binds a generic wrapper to an already deployed contract.
func bindClonefactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClonefactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clonefactory *ClonefactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clonefactory.Contract.ClonefactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clonefactory *ClonefactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.Contract.ClonefactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clonefactory *ClonefactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clonefactory.Contract.ClonefactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clonefactory *ClonefactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clonefactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clonefactory *ClonefactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clonefactory *ClonefactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clonefactory.Contract.contract.Transact(opts, method, params...)
}

// BaseImplementation is a free data retrieval call binding the contract method 0xa2ebe749.
//
// Solidity: function baseImplementation() view returns(address)
func (_Clonefactory *ClonefactoryCaller) BaseImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "baseImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BaseImplementation is a free data retrieval call binding the contract method 0xa2ebe749.
//
// Solidity: function baseImplementation() view returns(address)
func (_Clonefactory *ClonefactorySession) BaseImplementation() (common.Address, error) {
	return _Clonefactory.Contract.BaseImplementation(&_Clonefactory.CallOpts)
}

// BaseImplementation is a free data retrieval call binding the contract method 0xa2ebe749.
//
// Solidity: function baseImplementation() view returns(address)
func (_Clonefactory *ClonefactoryCallerSession) BaseImplementation() (common.Address, error) {
	return _Clonefactory.Contract.BaseImplementation(&_Clonefactory.CallOpts)
}

// CheckWhitelist is a free data retrieval call binding the contract method 0x1950c218.
//
// Solidity: function checkWhitelist(address _address) view returns(bool)
func (_Clonefactory *ClonefactoryCaller) CheckWhitelist(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "checkWhitelist", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckWhitelist is a free data retrieval call binding the contract method 0x1950c218.
//
// Solidity: function checkWhitelist(address _address) view returns(bool)
func (_Clonefactory *ClonefactorySession) CheckWhitelist(_address common.Address) (bool, error) {
	return _Clonefactory.Contract.CheckWhitelist(&_Clonefactory.CallOpts, _address)
}

// CheckWhitelist is a free data retrieval call binding the contract method 0x1950c218.
//
// Solidity: function checkWhitelist(address _address) view returns(bool)
func (_Clonefactory *ClonefactoryCallerSession) CheckWhitelist(_address common.Address) (bool, error) {
	return _Clonefactory.Contract.CheckWhitelist(&_Clonefactory.CallOpts, _address)
}

// GetContractList is a free data retrieval call binding the contract method 0x99acac8c.
//
// Solidity: function getContractList() view returns(address[])
func (_Clonefactory *ClonefactoryCaller) GetContractList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "getContractList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetContractList is a free data retrieval call binding the contract method 0x99acac8c.
//
// Solidity: function getContractList() view returns(address[])
func (_Clonefactory *ClonefactorySession) GetContractList() ([]common.Address, error) {
	return _Clonefactory.Contract.GetContractList(&_Clonefactory.CallOpts)
}

// GetContractList is a free data retrieval call binding the contract method 0x99acac8c.
//
// Solidity: function getContractList() view returns(address[])
func (_Clonefactory *ClonefactoryCallerSession) GetContractList() ([]common.Address, error) {
	return _Clonefactory.Contract.GetContractList(&_Clonefactory.CallOpts)
}

// IsContractDead is a free data retrieval call binding the contract method 0xbd80a1c9.
//
// Solidity: function isContractDead(address ) view returns(bool)
func (_Clonefactory *ClonefactoryCaller) IsContractDead(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "isContractDead", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContractDead is a free data retrieval call binding the contract method 0xbd80a1c9.
//
// Solidity: function isContractDead(address ) view returns(bool)
func (_Clonefactory *ClonefactorySession) IsContractDead(arg0 common.Address) (bool, error) {
	return _Clonefactory.Contract.IsContractDead(&_Clonefactory.CallOpts, arg0)
}

// IsContractDead is a free data retrieval call binding the contract method 0xbd80a1c9.
//
// Solidity: function isContractDead(address ) view returns(bool)
func (_Clonefactory *ClonefactoryCallerSession) IsContractDead(arg0 common.Address) (bool, error) {
	return _Clonefactory.Contract.IsContractDead(&_Clonefactory.CallOpts, arg0)
}

// Lumerin is a free data retrieval call binding the contract method 0xe6108971.
//
// Solidity: function lumerin() view returns(address)
func (_Clonefactory *ClonefactoryCaller) Lumerin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "lumerin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lumerin is a free data retrieval call binding the contract method 0xe6108971.
//
// Solidity: function lumerin() view returns(address)
func (_Clonefactory *ClonefactorySession) Lumerin() (common.Address, error) {
	return _Clonefactory.Contract.Lumerin(&_Clonefactory.CallOpts)
}

// Lumerin is a free data retrieval call binding the contract method 0xe6108971.
//
// Solidity: function lumerin() view returns(address)
func (_Clonefactory *ClonefactoryCallerSession) Lumerin() (common.Address, error) {
	return _Clonefactory.Contract.Lumerin(&_Clonefactory.CallOpts)
}

// MarketplaceFee is a free data retrieval call binding the contract method 0x6a166964.
//
// Solidity: function marketplaceFee() view returns(uint256)
func (_Clonefactory *ClonefactoryCaller) MarketplaceFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "marketplaceFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketplaceFee is a free data retrieval call binding the contract method 0x6a166964.
//
// Solidity: function marketplaceFee() view returns(uint256)
func (_Clonefactory *ClonefactorySession) MarketplaceFee() (*big.Int, error) {
	return _Clonefactory.Contract.MarketplaceFee(&_Clonefactory.CallOpts)
}

// MarketplaceFee is a free data retrieval call binding the contract method 0x6a166964.
//
// Solidity: function marketplaceFee() view returns(uint256)
func (_Clonefactory *ClonefactoryCallerSession) MarketplaceFee() (*big.Int, error) {
	return _Clonefactory.Contract.MarketplaceFee(&_Clonefactory.CallOpts)
}

// NoMoreWhitelist is a free data retrieval call binding the contract method 0xe32e2450.
//
// Solidity: function noMoreWhitelist() view returns(bool)
func (_Clonefactory *ClonefactoryCaller) NoMoreWhitelist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "noMoreWhitelist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NoMoreWhitelist is a free data retrieval call binding the contract method 0xe32e2450.
//
// Solidity: function noMoreWhitelist() view returns(bool)
func (_Clonefactory *ClonefactorySession) NoMoreWhitelist() (bool, error) {
	return _Clonefactory.Contract.NoMoreWhitelist(&_Clonefactory.CallOpts)
}

// NoMoreWhitelist is a free data retrieval call binding the contract method 0xe32e2450.
//
// Solidity: function noMoreWhitelist() view returns(bool)
func (_Clonefactory *ClonefactoryCallerSession) NoMoreWhitelist() (bool, error) {
	return _Clonefactory.Contract.NoMoreWhitelist(&_Clonefactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clonefactory *ClonefactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clonefactory *ClonefactorySession) Owner() (common.Address, error) {
	return _Clonefactory.Contract.Owner(&_Clonefactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Clonefactory *ClonefactoryCallerSession) Owner() (common.Address, error) {
	return _Clonefactory.Contract.Owner(&_Clonefactory.CallOpts)
}

// RentalContracts is a free data retrieval call binding the contract method 0x53da0206.
//
// Solidity: function rentalContracts(uint256 ) view returns(address)
func (_Clonefactory *ClonefactoryCaller) RentalContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "rentalContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RentalContracts is a free data retrieval call binding the contract method 0x53da0206.
//
// Solidity: function rentalContracts(uint256 ) view returns(address)
func (_Clonefactory *ClonefactorySession) RentalContracts(arg0 *big.Int) (common.Address, error) {
	return _Clonefactory.Contract.RentalContracts(&_Clonefactory.CallOpts, arg0)
}

// RentalContracts is a free data retrieval call binding the contract method 0x53da0206.
//
// Solidity: function rentalContracts(uint256 ) view returns(address)
func (_Clonefactory *ClonefactoryCallerSession) RentalContracts(arg0 *big.Int) (common.Address, error) {
	return _Clonefactory.Contract.RentalContracts(&_Clonefactory.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Clonefactory *ClonefactoryCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Clonefactory *ClonefactorySession) Whitelist(arg0 common.Address) (bool, error) {
	return _Clonefactory.Contract.Whitelist(&_Clonefactory.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Clonefactory *ClonefactoryCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Clonefactory.Contract.Whitelist(&_Clonefactory.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _baseImplementation, address _lumerin, address _feeRecipient) returns()
func (_Clonefactory *ClonefactoryTransactor) Initialize(opts *bind.TransactOpts, _baseImplementation common.Address, _lumerin common.Address, _feeRecipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "initialize", _baseImplementation, _lumerin, _feeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _baseImplementation, address _lumerin, address _feeRecipient) returns()
func (_Clonefactory *ClonefactorySession) Initialize(_baseImplementation common.Address, _lumerin common.Address, _feeRecipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.Initialize(&_Clonefactory.TransactOpts, _baseImplementation, _lumerin, _feeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _baseImplementation, address _lumerin, address _feeRecipient) returns()
func (_Clonefactory *ClonefactoryTransactorSession) Initialize(_baseImplementation common.Address, _lumerin common.Address, _feeRecipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.Initialize(&_Clonefactory.TransactOpts, _baseImplementation, _lumerin, _feeRecipient)
}

// PayMarketplaceFee is a paid mutator transaction binding the contract method 0x1088f93f.
//
// Solidity: function payMarketplaceFee() payable returns(bool)
func (_Clonefactory *ClonefactoryTransactor) PayMarketplaceFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "payMarketplaceFee")
}

// PayMarketplaceFee is a paid mutator transaction binding the contract method 0x1088f93f.
//
// Solidity: function payMarketplaceFee() payable returns(bool)
func (_Clonefactory *ClonefactorySession) PayMarketplaceFee() (*types.Transaction, error) {
	return _Clonefactory.Contract.PayMarketplaceFee(&_Clonefactory.TransactOpts)
}

// PayMarketplaceFee is a paid mutator transaction binding the contract method 0x1088f93f.
//
// Solidity: function payMarketplaceFee() payable returns(bool)
func (_Clonefactory *ClonefactoryTransactorSession) PayMarketplaceFee() (*types.Transaction, error) {
	return _Clonefactory.Contract.PayMarketplaceFee(&_Clonefactory.TransactOpts)
}

// SetAddToWhitelist is a paid mutator transaction binding the contract method 0xa762f7f1.
//
// Solidity: function setAddToWhitelist(address _address) returns()
func (_Clonefactory *ClonefactoryTransactor) SetAddToWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setAddToWhitelist", _address)
}

// SetAddToWhitelist is a paid mutator transaction binding the contract method 0xa762f7f1.
//
// Solidity: function setAddToWhitelist(address _address) returns()
func (_Clonefactory *ClonefactorySession) SetAddToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetAddToWhitelist(&_Clonefactory.TransactOpts, _address)
}

// SetAddToWhitelist is a paid mutator transaction binding the contract method 0xa762f7f1.
//
// Solidity: function setAddToWhitelist(address _address) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetAddToWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetAddToWhitelist(&_Clonefactory.TransactOpts, _address)
}

// SetContractDeleted is a paid mutator transaction binding the contract method 0x8664aa51.
//
// Solidity: function setContractDeleted(address _contractAddress, bool _isDeleted) returns()
func (_Clonefactory *ClonefactoryTransactor) SetContractDeleted(opts *bind.TransactOpts, _contractAddress common.Address, _isDeleted bool) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setContractDeleted", _contractAddress, _isDeleted)
}

// SetContractDeleted is a paid mutator transaction binding the contract method 0x8664aa51.
//
// Solidity: function setContractDeleted(address _contractAddress, bool _isDeleted) returns()
func (_Clonefactory *ClonefactorySession) SetContractDeleted(_contractAddress common.Address, _isDeleted bool) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractDeleted(&_Clonefactory.TransactOpts, _contractAddress, _isDeleted)
}

// SetContractDeleted is a paid mutator transaction binding the contract method 0x8664aa51.
//
// Solidity: function setContractDeleted(address _contractAddress, bool _isDeleted) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetContractDeleted(_contractAddress common.Address, _isDeleted bool) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractDeleted(&_Clonefactory.TransactOpts, _contractAddress, _isDeleted)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x86712686.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _validator, string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactoryTransactor) SetCreateNewRentalContract(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setCreateNewRentalContract", _price, _limit, _speed, _length, _validator, _pubKey)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x86712686.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _validator, string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactorySession) SetCreateNewRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContract(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _validator, _pubKey)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x86712686.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _validator, string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactoryTransactorSession) SetCreateNewRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContract(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _validator, _pubKey)
}

// SetCreateNewRentalContractV2 is a paid mutator transaction binding the contract method 0x2e1e8a93.
//
// Solidity: function setCreateNewRentalContractV2(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget, address _validator, string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactoryTransactor) SetCreateNewRentalContractV2(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setCreateNewRentalContractV2", _price, _limit, _speed, _length, _profitTarget, _validator, _pubKey)
}

// SetCreateNewRentalContractV2 is a paid mutator transaction binding the contract method 0x2e1e8a93.
//
// Solidity: function setCreateNewRentalContractV2(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget, address _validator, string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactorySession) SetCreateNewRentalContractV2(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContractV2(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _profitTarget, _validator, _pubKey)
}

// SetCreateNewRentalContractV2 is a paid mutator transaction binding the contract method 0x2e1e8a93.
//
// Solidity: function setCreateNewRentalContractV2(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget, address _validator, string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactoryTransactorSession) SetCreateNewRentalContractV2(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContractV2(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _profitTarget, _validator, _pubKey)
}

// SetDisableWhitelist is a paid mutator transaction binding the contract method 0xb312b94b.
//
// Solidity: function setDisableWhitelist() returns()
func (_Clonefactory *ClonefactoryTransactor) SetDisableWhitelist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setDisableWhitelist")
}

// SetDisableWhitelist is a paid mutator transaction binding the contract method 0xb312b94b.
//
// Solidity: function setDisableWhitelist() returns()
func (_Clonefactory *ClonefactorySession) SetDisableWhitelist() (*types.Transaction, error) {
	return _Clonefactory.Contract.SetDisableWhitelist(&_Clonefactory.TransactOpts)
}

// SetDisableWhitelist is a paid mutator transaction binding the contract method 0xb312b94b.
//
// Solidity: function setDisableWhitelist() returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetDisableWhitelist() (*types.Transaction, error) {
	return _Clonefactory.Contract.SetDisableWhitelist(&_Clonefactory.TransactOpts)
}

// SetMarketplaceFeeRecipient is a paid mutator transaction binding the contract method 0xe80b04b4.
//
// Solidity: function setMarketplaceFeeRecipient(uint256 fee, address recipient) returns()
func (_Clonefactory *ClonefactoryTransactor) SetMarketplaceFeeRecipient(opts *bind.TransactOpts, fee *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setMarketplaceFeeRecipient", fee, recipient)
}

// SetMarketplaceFeeRecipient is a paid mutator transaction binding the contract method 0xe80b04b4.
//
// Solidity: function setMarketplaceFeeRecipient(uint256 fee, address recipient) returns()
func (_Clonefactory *ClonefactorySession) SetMarketplaceFeeRecipient(fee *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetMarketplaceFeeRecipient(&_Clonefactory.TransactOpts, fee, recipient)
}

// SetMarketplaceFeeRecipient is a paid mutator transaction binding the contract method 0xe80b04b4.
//
// Solidity: function setMarketplaceFeeRecipient(uint256 fee, address recipient) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetMarketplaceFeeRecipient(fee *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetMarketplaceFeeRecipient(&_Clonefactory.TransactOpts, fee, recipient)
}

// SetPurchaseRentalContract is a paid mutator transaction binding the contract method 0xee878a4a.
//
// Solidity: function setPurchaseRentalContract(address _contractAddress, string _cipherText, uint32 termsVersion) payable returns()
func (_Clonefactory *ClonefactoryTransactor) SetPurchaseRentalContract(opts *bind.TransactOpts, _contractAddress common.Address, _cipherText string, termsVersion uint32) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setPurchaseRentalContract", _contractAddress, _cipherText, termsVersion)
}

// SetPurchaseRentalContract is a paid mutator transaction binding the contract method 0xee878a4a.
//
// Solidity: function setPurchaseRentalContract(address _contractAddress, string _cipherText, uint32 termsVersion) payable returns()
func (_Clonefactory *ClonefactorySession) SetPurchaseRentalContract(_contractAddress common.Address, _cipherText string, termsVersion uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetPurchaseRentalContract(&_Clonefactory.TransactOpts, _contractAddress, _cipherText, termsVersion)
}

// SetPurchaseRentalContract is a paid mutator transaction binding the contract method 0xee878a4a.
//
// Solidity: function setPurchaseRentalContract(address _contractAddress, string _cipherText, uint32 termsVersion) payable returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetPurchaseRentalContract(_contractAddress common.Address, _cipherText string, termsVersion uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetPurchaseRentalContract(&_Clonefactory.TransactOpts, _contractAddress, _cipherText, termsVersion)
}

// SetPurchaseRentalContractV2 is a paid mutator transaction binding the contract method 0x3af4158f.
//
// Solidity: function setPurchaseRentalContractV2(address _contractAddress, address _validatorAddress, string _encrValidatorURL, string _encrDestURL, uint32 termsVersion) payable returns()
func (_Clonefactory *ClonefactoryTransactor) SetPurchaseRentalContractV2(opts *bind.TransactOpts, _contractAddress common.Address, _validatorAddress common.Address, _encrValidatorURL string, _encrDestURL string, termsVersion uint32) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setPurchaseRentalContractV2", _contractAddress, _validatorAddress, _encrValidatorURL, _encrDestURL, termsVersion)
}

// SetPurchaseRentalContractV2 is a paid mutator transaction binding the contract method 0x3af4158f.
//
// Solidity: function setPurchaseRentalContractV2(address _contractAddress, address _validatorAddress, string _encrValidatorURL, string _encrDestURL, uint32 termsVersion) payable returns()
func (_Clonefactory *ClonefactorySession) SetPurchaseRentalContractV2(_contractAddress common.Address, _validatorAddress common.Address, _encrValidatorURL string, _encrDestURL string, termsVersion uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetPurchaseRentalContractV2(&_Clonefactory.TransactOpts, _contractAddress, _validatorAddress, _encrValidatorURL, _encrDestURL, termsVersion)
}

// SetPurchaseRentalContractV2 is a paid mutator transaction binding the contract method 0x3af4158f.
//
// Solidity: function setPurchaseRentalContractV2(address _contractAddress, address _validatorAddress, string _encrValidatorURL, string _encrDestURL, uint32 termsVersion) payable returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetPurchaseRentalContractV2(_contractAddress common.Address, _validatorAddress common.Address, _encrValidatorURL string, _encrDestURL string, termsVersion uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetPurchaseRentalContractV2(&_Clonefactory.TransactOpts, _contractAddress, _validatorAddress, _encrValidatorURL, _encrDestURL, termsVersion)
}

// SetRemoveFromWhitelist is a paid mutator transaction binding the contract method 0x164c1976.
//
// Solidity: function setRemoveFromWhitelist(address _address) returns()
func (_Clonefactory *ClonefactoryTransactor) SetRemoveFromWhitelist(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setRemoveFromWhitelist", _address)
}

// SetRemoveFromWhitelist is a paid mutator transaction binding the contract method 0x164c1976.
//
// Solidity: function setRemoveFromWhitelist(address _address) returns()
func (_Clonefactory *ClonefactorySession) SetRemoveFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetRemoveFromWhitelist(&_Clonefactory.TransactOpts, _address)
}

// SetRemoveFromWhitelist is a paid mutator transaction binding the contract method 0x164c1976.
//
// Solidity: function setRemoveFromWhitelist(address _address) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetRemoveFromWhitelist(_address common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetRemoveFromWhitelist(&_Clonefactory.TransactOpts, _address)
}

// SetUpdateContractInformation is a paid mutator transaction binding the contract method 0xaafcc02a.
//
// Solidity: function setUpdateContractInformation(address _contractAddress, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) payable returns()
func (_Clonefactory *ClonefactoryTransactor) SetUpdateContractInformation(opts *bind.TransactOpts, _contractAddress common.Address, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setUpdateContractInformation", _contractAddress, _price, _limit, _speed, _length)
}

// SetUpdateContractInformation is a paid mutator transaction binding the contract method 0xaafcc02a.
//
// Solidity: function setUpdateContractInformation(address _contractAddress, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) payable returns()
func (_Clonefactory *ClonefactorySession) SetUpdateContractInformation(_contractAddress common.Address, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetUpdateContractInformation(&_Clonefactory.TransactOpts, _contractAddress, _price, _limit, _speed, _length)
}

// SetUpdateContractInformation is a paid mutator transaction binding the contract method 0xaafcc02a.
//
// Solidity: function setUpdateContractInformation(address _contractAddress, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length) payable returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetUpdateContractInformation(_contractAddress common.Address, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetUpdateContractInformation(&_Clonefactory.TransactOpts, _contractAddress, _price, _limit, _speed, _length)
}

// SetUpdateContractInformationV2 is a paid mutator transaction binding the contract method 0x3a6738b2.
//
// Solidity: function setUpdateContractInformationV2(address _contractAddress, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Clonefactory *ClonefactoryTransactor) SetUpdateContractInformationV2(opts *bind.TransactOpts, _contractAddress common.Address, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setUpdateContractInformationV2", _contractAddress, _price, _limit, _speed, _length, _profitTarget)
}

// SetUpdateContractInformationV2 is a paid mutator transaction binding the contract method 0x3a6738b2.
//
// Solidity: function setUpdateContractInformationV2(address _contractAddress, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Clonefactory *ClonefactorySession) SetUpdateContractInformationV2(_contractAddress common.Address, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetUpdateContractInformationV2(&_Clonefactory.TransactOpts, _contractAddress, _price, _limit, _speed, _length, _profitTarget)
}

// SetUpdateContractInformationV2 is a paid mutator transaction binding the contract method 0x3a6738b2.
//
// Solidity: function setUpdateContractInformationV2(address _contractAddress, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetUpdateContractInformationV2(_contractAddress common.Address, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetUpdateContractInformationV2(&_Clonefactory.TransactOpts, _contractAddress, _price, _limit, _speed, _length, _profitTarget)
}

// ClonefactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Clonefactory contract.
type ClonefactoryInitializedIterator struct {
	Event *ClonefactoryInitialized // Event containing the contract specifics and raw log

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
func (it *ClonefactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryInitialized)
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
		it.Event = new(ClonefactoryInitialized)
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
func (it *ClonefactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryInitialized represents a Initialized event raised by the Clonefactory contract.
type ClonefactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Clonefactory *ClonefactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ClonefactoryInitializedIterator, error) {

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ClonefactoryInitializedIterator{contract: _Clonefactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Clonefactory *ClonefactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ClonefactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryInitialized)
				if err := _Clonefactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Clonefactory *ClonefactoryFilterer) ParseInitialized(log types.Log) (*ClonefactoryInitialized, error) {
	event := new(ClonefactoryInitialized)
	if err := _Clonefactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonefactoryClonefactoryContractPurchasedIterator is returned from FilterClonefactoryContractPurchased and is used to iterate over the raw logs and unpacked data for ClonefactoryContractPurchased events raised by the Clonefactory contract.
type ClonefactoryClonefactoryContractPurchasedIterator struct {
	Event *ClonefactoryClonefactoryContractPurchased // Event containing the contract specifics and raw log

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
func (it *ClonefactoryClonefactoryContractPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryClonefactoryContractPurchased)
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
		it.Event = new(ClonefactoryClonefactoryContractPurchased)
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
func (it *ClonefactoryClonefactoryContractPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryClonefactoryContractPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryClonefactoryContractPurchased represents a ClonefactoryContractPurchased event raised by the Clonefactory contract.
type ClonefactoryClonefactoryContractPurchased struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterClonefactoryContractPurchased is a free log retrieval operation binding the contract event 0xbf1df41b401a1bb8d9bd03fb6fe59b6ced0e61a76cdd3d3d511b4d06eb2cdebe.
//
// Solidity: event clonefactoryContractPurchased(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) FilterClonefactoryContractPurchased(opts *bind.FilterOpts, _address []common.Address) (*ClonefactoryClonefactoryContractPurchasedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "clonefactoryContractPurchased", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryClonefactoryContractPurchasedIterator{contract: _Clonefactory.contract, event: "clonefactoryContractPurchased", logs: logs, sub: sub}, nil
}

// WatchClonefactoryContractPurchased is a free log subscription operation binding the contract event 0xbf1df41b401a1bb8d9bd03fb6fe59b6ced0e61a76cdd3d3d511b4d06eb2cdebe.
//
// Solidity: event clonefactoryContractPurchased(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) WatchClonefactoryContractPurchased(opts *bind.WatchOpts, sink chan<- *ClonefactoryClonefactoryContractPurchased, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "clonefactoryContractPurchased", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryClonefactoryContractPurchased)
				if err := _Clonefactory.contract.UnpackLog(event, "clonefactoryContractPurchased", log); err != nil {
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

// ParseClonefactoryContractPurchased is a log parse operation binding the contract event 0xbf1df41b401a1bb8d9bd03fb6fe59b6ced0e61a76cdd3d3d511b4d06eb2cdebe.
//
// Solidity: event clonefactoryContractPurchased(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) ParseClonefactoryContractPurchased(log types.Log) (*ClonefactoryClonefactoryContractPurchased, error) {
	event := new(ClonefactoryClonefactoryContractPurchased)
	if err := _Clonefactory.contract.UnpackLog(event, "clonefactoryContractPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonefactoryContractCreatedIterator is returned from FilterContractCreated and is used to iterate over the raw logs and unpacked data for ContractCreated events raised by the Clonefactory contract.
type ClonefactoryContractCreatedIterator struct {
	Event *ClonefactoryContractCreated // Event containing the contract specifics and raw log

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
func (it *ClonefactoryContractCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryContractCreated)
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
		it.Event = new(ClonefactoryContractCreated)
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
func (it *ClonefactoryContractCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryContractCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryContractCreated represents a ContractCreated event raised by the Clonefactory contract.
type ClonefactoryContractCreated struct {
	Address common.Address
	Pubkey  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractCreated is a free log retrieval operation binding the contract event 0x1b08e1646439b7521399d47f93ab6b1ebc92803e155d0b2f2ad2d4702a050804.
//
// Solidity: event contractCreated(address indexed _address, string _pubkey)
func (_Clonefactory *ClonefactoryFilterer) FilterContractCreated(opts *bind.FilterOpts, _address []common.Address) (*ClonefactoryContractCreatedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "contractCreated", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryContractCreatedIterator{contract: _Clonefactory.contract, event: "contractCreated", logs: logs, sub: sub}, nil
}

// WatchContractCreated is a free log subscription operation binding the contract event 0x1b08e1646439b7521399d47f93ab6b1ebc92803e155d0b2f2ad2d4702a050804.
//
// Solidity: event contractCreated(address indexed _address, string _pubkey)
func (_Clonefactory *ClonefactoryFilterer) WatchContractCreated(opts *bind.WatchOpts, sink chan<- *ClonefactoryContractCreated, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "contractCreated", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryContractCreated)
				if err := _Clonefactory.contract.UnpackLog(event, "contractCreated", log); err != nil {
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

// ParseContractCreated is a log parse operation binding the contract event 0x1b08e1646439b7521399d47f93ab6b1ebc92803e155d0b2f2ad2d4702a050804.
//
// Solidity: event contractCreated(address indexed _address, string _pubkey)
func (_Clonefactory *ClonefactoryFilterer) ParseContractCreated(log types.Log) (*ClonefactoryContractCreated, error) {
	event := new(ClonefactoryContractCreated)
	if err := _Clonefactory.contract.UnpackLog(event, "contractCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonefactoryContractDeleteUpdatedIterator is returned from FilterContractDeleteUpdated and is used to iterate over the raw logs and unpacked data for ContractDeleteUpdated events raised by the Clonefactory contract.
type ClonefactoryContractDeleteUpdatedIterator struct {
	Event *ClonefactoryContractDeleteUpdated // Event containing the contract specifics and raw log

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
func (it *ClonefactoryContractDeleteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryContractDeleteUpdated)
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
		it.Event = new(ClonefactoryContractDeleteUpdated)
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
func (it *ClonefactoryContractDeleteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryContractDeleteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryContractDeleteUpdated represents a ContractDeleteUpdated event raised by the Clonefactory contract.
type ClonefactoryContractDeleteUpdated struct {
	Address   common.Address
	IsDeleted bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterContractDeleteUpdated is a free log retrieval operation binding the contract event 0x94706a90834905eb13df4134bd7c40e90c8218fbe9ded8631b8062c7bd4a5a06.
//
// Solidity: event contractDeleteUpdated(address _address, bool _isDeleted)
func (_Clonefactory *ClonefactoryFilterer) FilterContractDeleteUpdated(opts *bind.FilterOpts) (*ClonefactoryContractDeleteUpdatedIterator, error) {

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "contractDeleteUpdated")
	if err != nil {
		return nil, err
	}
	return &ClonefactoryContractDeleteUpdatedIterator{contract: _Clonefactory.contract, event: "contractDeleteUpdated", logs: logs, sub: sub}, nil
}

// WatchContractDeleteUpdated is a free log subscription operation binding the contract event 0x94706a90834905eb13df4134bd7c40e90c8218fbe9ded8631b8062c7bd4a5a06.
//
// Solidity: event contractDeleteUpdated(address _address, bool _isDeleted)
func (_Clonefactory *ClonefactoryFilterer) WatchContractDeleteUpdated(opts *bind.WatchOpts, sink chan<- *ClonefactoryContractDeleteUpdated) (event.Subscription, error) {

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "contractDeleteUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryContractDeleteUpdated)
				if err := _Clonefactory.contract.UnpackLog(event, "contractDeleteUpdated", log); err != nil {
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

// ParseContractDeleteUpdated is a log parse operation binding the contract event 0x94706a90834905eb13df4134bd7c40e90c8218fbe9ded8631b8062c7bd4a5a06.
//
// Solidity: event contractDeleteUpdated(address _address, bool _isDeleted)
func (_Clonefactory *ClonefactoryFilterer) ParseContractDeleteUpdated(log types.Log) (*ClonefactoryContractDeleteUpdated, error) {
	event := new(ClonefactoryContractDeleteUpdated)
	if err := _Clonefactory.contract.UnpackLog(event, "contractDeleteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonefactoryPurchaseInfoUpdatedIterator is returned from FilterPurchaseInfoUpdated and is used to iterate over the raw logs and unpacked data for PurchaseInfoUpdated events raised by the Clonefactory contract.
type ClonefactoryPurchaseInfoUpdatedIterator struct {
	Event *ClonefactoryPurchaseInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ClonefactoryPurchaseInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryPurchaseInfoUpdated)
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
		it.Event = new(ClonefactoryPurchaseInfoUpdated)
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
func (it *ClonefactoryPurchaseInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryPurchaseInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryPurchaseInfoUpdated represents a PurchaseInfoUpdated event raised by the Clonefactory contract.
type ClonefactoryPurchaseInfoUpdated struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPurchaseInfoUpdated is a free log retrieval operation binding the contract event 0x7865b2bbe0087425b223d9cf674d8d5328e31546b364da98c36db21193c17d55.
//
// Solidity: event purchaseInfoUpdated(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) FilterPurchaseInfoUpdated(opts *bind.FilterOpts, _address []common.Address) (*ClonefactoryPurchaseInfoUpdatedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "purchaseInfoUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryPurchaseInfoUpdatedIterator{contract: _Clonefactory.contract, event: "purchaseInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPurchaseInfoUpdated is a free log subscription operation binding the contract event 0x7865b2bbe0087425b223d9cf674d8d5328e31546b364da98c36db21193c17d55.
//
// Solidity: event purchaseInfoUpdated(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) WatchPurchaseInfoUpdated(opts *bind.WatchOpts, sink chan<- *ClonefactoryPurchaseInfoUpdated, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "purchaseInfoUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryPurchaseInfoUpdated)
				if err := _Clonefactory.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
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

// ParsePurchaseInfoUpdated is a log parse operation binding the contract event 0x7865b2bbe0087425b223d9cf674d8d5328e31546b364da98c36db21193c17d55.
//
// Solidity: event purchaseInfoUpdated(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) ParsePurchaseInfoUpdated(log types.Log) (*ClonefactoryPurchaseInfoUpdated, error) {
	event := new(ClonefactoryPurchaseInfoUpdated)
	if err := _Clonefactory.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
