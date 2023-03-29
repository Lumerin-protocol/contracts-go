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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lmn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"clonefactoryContractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"contractCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"buyerFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"checkWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isContractDead\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noMoreWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rentalContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellerFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"setBaseImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setChangeBuyerFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"}],\"name\":\"setChangeMarketplaceRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFee\",\"type\":\"uint256\"}],\"name\":\"setChangeSellerFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"closeout\",\"type\":\"bool\"}],\"name\":\"setContractAsDead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pubKey\",\"type\":\"string\"}],\"name\":\"setCreateNewRentalContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setDisableWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_cipherText\",\"type\":\"string\"}],\"name\":\"setPurchaseRentalContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRemoveFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// BuyerFeeRate is a free data retrieval call binding the contract method 0x53f10c80.
//
// Solidity: function buyerFeeRate() view returns(uint256)
func (_Clonefactory *ClonefactoryCaller) BuyerFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "buyerFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyerFeeRate is a free data retrieval call binding the contract method 0x53f10c80.
//
// Solidity: function buyerFeeRate() view returns(uint256)
func (_Clonefactory *ClonefactorySession) BuyerFeeRate() (*big.Int, error) {
	return _Clonefactory.Contract.BuyerFeeRate(&_Clonefactory.CallOpts)
}

// BuyerFeeRate is a free data retrieval call binding the contract method 0x53f10c80.
//
// Solidity: function buyerFeeRate() view returns(uint256)
func (_Clonefactory *ClonefactoryCallerSession) BuyerFeeRate() (*big.Int, error) {
	return _Clonefactory.Contract.BuyerFeeRate(&_Clonefactory.CallOpts)
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

// SellerFeeRate is a free data retrieval call binding the contract method 0x89f68190.
//
// Solidity: function sellerFeeRate() view returns(uint256)
func (_Clonefactory *ClonefactoryCaller) SellerFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "sellerFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellerFeeRate is a free data retrieval call binding the contract method 0x89f68190.
//
// Solidity: function sellerFeeRate() view returns(uint256)
func (_Clonefactory *ClonefactorySession) SellerFeeRate() (*big.Int, error) {
	return _Clonefactory.Contract.SellerFeeRate(&_Clonefactory.CallOpts)
}

// SellerFeeRate is a free data retrieval call binding the contract method 0x89f68190.
//
// Solidity: function sellerFeeRate() view returns(uint256)
func (_Clonefactory *ClonefactoryCallerSession) SellerFeeRate() (*big.Int, error) {
	return _Clonefactory.Contract.SellerFeeRate(&_Clonefactory.CallOpts)
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

// SetBaseImplementation is a paid mutator transaction binding the contract method 0x311c3f5f.
//
// Solidity: function setBaseImplementation(address _newImplementation) returns()
func (_Clonefactory *ClonefactoryTransactor) SetBaseImplementation(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setBaseImplementation", _newImplementation)
}

// SetBaseImplementation is a paid mutator transaction binding the contract method 0x311c3f5f.
//
// Solidity: function setBaseImplementation(address _newImplementation) returns()
func (_Clonefactory *ClonefactorySession) SetBaseImplementation(_newImplementation common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetBaseImplementation(&_Clonefactory.TransactOpts, _newImplementation)
}

// SetBaseImplementation is a paid mutator transaction binding the contract method 0x311c3f5f.
//
// Solidity: function setBaseImplementation(address _newImplementation) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetBaseImplementation(_newImplementation common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetBaseImplementation(&_Clonefactory.TransactOpts, _newImplementation)
}

// SetChangeBuyerFeeRate is a paid mutator transaction binding the contract method 0x8fc7c9b1.
//
// Solidity: function setChangeBuyerFeeRate(uint256 _newFee) returns()
func (_Clonefactory *ClonefactoryTransactor) SetChangeBuyerFeeRate(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setChangeBuyerFeeRate", _newFee)
}

// SetChangeBuyerFeeRate is a paid mutator transaction binding the contract method 0x8fc7c9b1.
//
// Solidity: function setChangeBuyerFeeRate(uint256 _newFee) returns()
func (_Clonefactory *ClonefactorySession) SetChangeBuyerFeeRate(_newFee *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetChangeBuyerFeeRate(&_Clonefactory.TransactOpts, _newFee)
}

// SetChangeBuyerFeeRate is a paid mutator transaction binding the contract method 0x8fc7c9b1.
//
// Solidity: function setChangeBuyerFeeRate(uint256 _newFee) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetChangeBuyerFeeRate(_newFee *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetChangeBuyerFeeRate(&_Clonefactory.TransactOpts, _newFee)
}

// SetChangeMarketplaceRecipient is a paid mutator transaction binding the contract method 0xa6494cb7.
//
// Solidity: function setChangeMarketplaceRecipient(address _newRecipient) returns()
func (_Clonefactory *ClonefactoryTransactor) SetChangeMarketplaceRecipient(opts *bind.TransactOpts, _newRecipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setChangeMarketplaceRecipient", _newRecipient)
}

// SetChangeMarketplaceRecipient is a paid mutator transaction binding the contract method 0xa6494cb7.
//
// Solidity: function setChangeMarketplaceRecipient(address _newRecipient) returns()
func (_Clonefactory *ClonefactorySession) SetChangeMarketplaceRecipient(_newRecipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetChangeMarketplaceRecipient(&_Clonefactory.TransactOpts, _newRecipient)
}

// SetChangeMarketplaceRecipient is a paid mutator transaction binding the contract method 0xa6494cb7.
//
// Solidity: function setChangeMarketplaceRecipient(address _newRecipient) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetChangeMarketplaceRecipient(_newRecipient common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetChangeMarketplaceRecipient(&_Clonefactory.TransactOpts, _newRecipient)
}

// SetChangeSellerFeeRate is a paid mutator transaction binding the contract method 0x0a17bccf.
//
// Solidity: function setChangeSellerFeeRate(uint256 _newFee) returns()
func (_Clonefactory *ClonefactoryTransactor) SetChangeSellerFeeRate(opts *bind.TransactOpts, _newFee *big.Int) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setChangeSellerFeeRate", _newFee)
}

// SetChangeSellerFeeRate is a paid mutator transaction binding the contract method 0x0a17bccf.
//
// Solidity: function setChangeSellerFeeRate(uint256 _newFee) returns()
func (_Clonefactory *ClonefactorySession) SetChangeSellerFeeRate(_newFee *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetChangeSellerFeeRate(&_Clonefactory.TransactOpts, _newFee)
}

// SetChangeSellerFeeRate is a paid mutator transaction binding the contract method 0x0a17bccf.
//
// Solidity: function setChangeSellerFeeRate(uint256 _newFee) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetChangeSellerFeeRate(_newFee *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetChangeSellerFeeRate(&_Clonefactory.TransactOpts, _newFee)
}

// SetContractAsDead is a paid mutator transaction binding the contract method 0x0bcb4f3b.
//
// Solidity: function setContractAsDead(address _contract, bool closeout) returns()
func (_Clonefactory *ClonefactoryTransactor) SetContractAsDead(opts *bind.TransactOpts, _contract common.Address, closeout bool) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setContractAsDead", _contract, closeout)
}

// SetContractAsDead is a paid mutator transaction binding the contract method 0x0bcb4f3b.
//
// Solidity: function setContractAsDead(address _contract, bool closeout) returns()
func (_Clonefactory *ClonefactorySession) SetContractAsDead(_contract common.Address, closeout bool) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractAsDead(&_Clonefactory.TransactOpts, _contract, closeout)
}

// SetContractAsDead is a paid mutator transaction binding the contract method 0x0bcb4f3b.
//
// Solidity: function setContractAsDead(address _contract, bool closeout) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetContractAsDead(_contract common.Address, closeout bool) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractAsDead(&_Clonefactory.TransactOpts, _contract, closeout)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x86712686.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _validator, string _pubKey) returns(address)
func (_Clonefactory *ClonefactoryTransactor) SetCreateNewRentalContract(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setCreateNewRentalContract", _price, _limit, _speed, _length, _validator, _pubKey)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x86712686.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _validator, string _pubKey) returns(address)
func (_Clonefactory *ClonefactorySession) SetCreateNewRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContract(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _validator, _pubKey)
}

// SetCreateNewRentalContract is a paid mutator transaction binding the contract method 0x86712686.
//
// Solidity: function setCreateNewRentalContract(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, address _validator, string _pubKey) returns(address)
func (_Clonefactory *ClonefactoryTransactorSession) SetCreateNewRentalContract(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContract(&_Clonefactory.TransactOpts, _price, _limit, _speed, _length, _validator, _pubKey)
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

// SetPurchaseRentalContract is a paid mutator transaction binding the contract method 0x739a8353.
//
// Solidity: function setPurchaseRentalContract(address contractAddress, string _cipherText) returns()
func (_Clonefactory *ClonefactoryTransactor) SetPurchaseRentalContract(opts *bind.TransactOpts, contractAddress common.Address, _cipherText string) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setPurchaseRentalContract", contractAddress, _cipherText)
}

// SetPurchaseRentalContract is a paid mutator transaction binding the contract method 0x739a8353.
//
// Solidity: function setPurchaseRentalContract(address contractAddress, string _cipherText) returns()
func (_Clonefactory *ClonefactorySession) SetPurchaseRentalContract(contractAddress common.Address, _cipherText string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetPurchaseRentalContract(&_Clonefactory.TransactOpts, contractAddress, _cipherText)
}

// SetPurchaseRentalContract is a paid mutator transaction binding the contract method 0x739a8353.
//
// Solidity: function setPurchaseRentalContract(address contractAddress, string _cipherText) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetPurchaseRentalContract(contractAddress common.Address, _cipherText string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetPurchaseRentalContract(&_Clonefactory.TransactOpts, contractAddress, _cipherText)
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
