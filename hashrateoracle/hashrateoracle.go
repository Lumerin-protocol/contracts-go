// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hashrateoracle

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

// HashrateoracleMetaData contains all meta data concerning the Hashrateoracle contract.
var HashrateoracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_btcTokenOracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_tokenDecimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueCannotBeZero\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBlockReward\",\"type\":\"uint256\"}],\"name\":\"BlockRewardUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDifficulty\",\"type\":\"uint256\"}],\"name\":\"DifficultyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"btcTokenOracle\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBlockReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDifficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHashesForBTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHashesforToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBlockReward\",\"type\":\"uint256\"}],\"name\":\"setBlockReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newDifficulty\",\"type\":\"uint256\"}],\"name\":\"setDifficulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// HashrateoracleABI is the input ABI used to generate the binding from.
// Deprecated: Use HashrateoracleMetaData.ABI instead.
var HashrateoracleABI = HashrateoracleMetaData.ABI

// Hashrateoracle is an auto generated Go binding around an Ethereum contract.
type Hashrateoracle struct {
	HashrateoracleCaller     // Read-only binding to the contract
	HashrateoracleTransactor // Write-only binding to the contract
	HashrateoracleFilterer   // Log filterer for contract events
}

// HashrateoracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashrateoracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashrateoracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashrateoracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashrateoracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashrateoracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashrateoracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashrateoracleSession struct {
	Contract     *Hashrateoracle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashrateoracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashrateoracleCallerSession struct {
	Contract *HashrateoracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HashrateoracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashrateoracleTransactorSession struct {
	Contract     *HashrateoracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HashrateoracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashrateoracleRaw struct {
	Contract *Hashrateoracle // Generic contract binding to access the raw methods on
}

// HashrateoracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashrateoracleCallerRaw struct {
	Contract *HashrateoracleCaller // Generic read-only contract binding to access the raw methods on
}

// HashrateoracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashrateoracleTransactorRaw struct {
	Contract *HashrateoracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashrateoracle creates a new instance of Hashrateoracle, bound to a specific deployed contract.
func NewHashrateoracle(address common.Address, backend bind.ContractBackend) (*Hashrateoracle, error) {
	contract, err := bindHashrateoracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hashrateoracle{HashrateoracleCaller: HashrateoracleCaller{contract: contract}, HashrateoracleTransactor: HashrateoracleTransactor{contract: contract}, HashrateoracleFilterer: HashrateoracleFilterer{contract: contract}}, nil
}

// NewHashrateoracleCaller creates a new read-only instance of Hashrateoracle, bound to a specific deployed contract.
func NewHashrateoracleCaller(address common.Address, caller bind.ContractCaller) (*HashrateoracleCaller, error) {
	contract, err := bindHashrateoracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashrateoracleCaller{contract: contract}, nil
}

// NewHashrateoracleTransactor creates a new write-only instance of Hashrateoracle, bound to a specific deployed contract.
func NewHashrateoracleTransactor(address common.Address, transactor bind.ContractTransactor) (*HashrateoracleTransactor, error) {
	contract, err := bindHashrateoracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashrateoracleTransactor{contract: contract}, nil
}

// NewHashrateoracleFilterer creates a new log filterer instance of Hashrateoracle, bound to a specific deployed contract.
func NewHashrateoracleFilterer(address common.Address, filterer bind.ContractFilterer) (*HashrateoracleFilterer, error) {
	contract, err := bindHashrateoracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashrateoracleFilterer{contract: contract}, nil
}

// bindHashrateoracle binds a generic wrapper to an already deployed contract.
func bindHashrateoracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashrateoracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashrateoracle *HashrateoracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashrateoracle.Contract.HashrateoracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashrateoracle *HashrateoracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.HashrateoracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashrateoracle *HashrateoracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.HashrateoracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hashrateoracle *HashrateoracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hashrateoracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hashrateoracle *HashrateoracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hashrateoracle *HashrateoracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Hashrateoracle *HashrateoracleCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Hashrateoracle *HashrateoracleSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Hashrateoracle.Contract.UPGRADEINTERFACEVERSION(&_Hashrateoracle.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Hashrateoracle *HashrateoracleCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Hashrateoracle.Contract.UPGRADEINTERFACEVERSION(&_Hashrateoracle.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Hashrateoracle *HashrateoracleCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Hashrateoracle *HashrateoracleSession) VERSION() (string, error) {
	return _Hashrateoracle.Contract.VERSION(&_Hashrateoracle.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Hashrateoracle *HashrateoracleCallerSession) VERSION() (string, error) {
	return _Hashrateoracle.Contract.VERSION(&_Hashrateoracle.CallOpts)
}

// BtcTokenOracle is a free data retrieval call binding the contract method 0xeb94400c.
//
// Solidity: function btcTokenOracle() view returns(address)
func (_Hashrateoracle *HashrateoracleCaller) BtcTokenOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "btcTokenOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BtcTokenOracle is a free data retrieval call binding the contract method 0xeb94400c.
//
// Solidity: function btcTokenOracle() view returns(address)
func (_Hashrateoracle *HashrateoracleSession) BtcTokenOracle() (common.Address, error) {
	return _Hashrateoracle.Contract.BtcTokenOracle(&_Hashrateoracle.CallOpts)
}

// BtcTokenOracle is a free data retrieval call binding the contract method 0xeb94400c.
//
// Solidity: function btcTokenOracle() view returns(address)
func (_Hashrateoracle *HashrateoracleCallerSession) BtcTokenOracle() (common.Address, error) {
	return _Hashrateoracle.Contract.BtcTokenOracle(&_Hashrateoracle.CallOpts)
}

// GetBlockReward is a free data retrieval call binding the contract method 0xf89d4086.
//
// Solidity: function getBlockReward() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCaller) GetBlockReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "getBlockReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockReward is a free data retrieval call binding the contract method 0xf89d4086.
//
// Solidity: function getBlockReward() view returns(uint256)
func (_Hashrateoracle *HashrateoracleSession) GetBlockReward() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetBlockReward(&_Hashrateoracle.CallOpts)
}

// GetBlockReward is a free data retrieval call binding the contract method 0xf89d4086.
//
// Solidity: function getBlockReward() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCallerSession) GetBlockReward() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetBlockReward(&_Hashrateoracle.CallOpts)
}

// GetDifficulty is a free data retrieval call binding the contract method 0xb6baffe3.
//
// Solidity: function getDifficulty() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCaller) GetDifficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "getDifficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDifficulty is a free data retrieval call binding the contract method 0xb6baffe3.
//
// Solidity: function getDifficulty() view returns(uint256)
func (_Hashrateoracle *HashrateoracleSession) GetDifficulty() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetDifficulty(&_Hashrateoracle.CallOpts)
}

// GetDifficulty is a free data retrieval call binding the contract method 0xb6baffe3.
//
// Solidity: function getDifficulty() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCallerSession) GetDifficulty() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetDifficulty(&_Hashrateoracle.CallOpts)
}

// GetHashesForBTC is a free data retrieval call binding the contract method 0x19e26291.
//
// Solidity: function getHashesForBTC() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCaller) GetHashesForBTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "getHashesForBTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHashesForBTC is a free data retrieval call binding the contract method 0x19e26291.
//
// Solidity: function getHashesForBTC() view returns(uint256)
func (_Hashrateoracle *HashrateoracleSession) GetHashesForBTC() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetHashesForBTC(&_Hashrateoracle.CallOpts)
}

// GetHashesForBTC is a free data retrieval call binding the contract method 0x19e26291.
//
// Solidity: function getHashesForBTC() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCallerSession) GetHashesForBTC() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetHashesForBTC(&_Hashrateoracle.CallOpts)
}

// GetHashesforToken is a free data retrieval call binding the contract method 0x3256ce75.
//
// Solidity: function getHashesforToken() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCaller) GetHashesforToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "getHashesforToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHashesforToken is a free data retrieval call binding the contract method 0x3256ce75.
//
// Solidity: function getHashesforToken() view returns(uint256)
func (_Hashrateoracle *HashrateoracleSession) GetHashesforToken() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetHashesforToken(&_Hashrateoracle.CallOpts)
}

// GetHashesforToken is a free data retrieval call binding the contract method 0x3256ce75.
//
// Solidity: function getHashesforToken() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCallerSession) GetHashesforToken() (*big.Int, error) {
	return _Hashrateoracle.Contract.GetHashesforToken(&_Hashrateoracle.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCaller) OracleDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "oracleDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_Hashrateoracle *HashrateoracleSession) OracleDecimals() (*big.Int, error) {
	return _Hashrateoracle.Contract.OracleDecimals(&_Hashrateoracle.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCallerSession) OracleDecimals() (*big.Int, error) {
	return _Hashrateoracle.Contract.OracleDecimals(&_Hashrateoracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hashrateoracle *HashrateoracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hashrateoracle *HashrateoracleSession) Owner() (common.Address, error) {
	return _Hashrateoracle.Contract.Owner(&_Hashrateoracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hashrateoracle *HashrateoracleCallerSession) Owner() (common.Address, error) {
	return _Hashrateoracle.Contract.Owner(&_Hashrateoracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Hashrateoracle *HashrateoracleCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Hashrateoracle *HashrateoracleSession) ProxiableUUID() ([32]byte, error) {
	return _Hashrateoracle.Contract.ProxiableUUID(&_Hashrateoracle.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Hashrateoracle *HashrateoracleCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Hashrateoracle.Contract.ProxiableUUID(&_Hashrateoracle.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCaller) TokenDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hashrateoracle.contract.Call(opts, &out, "tokenDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint256)
func (_Hashrateoracle *HashrateoracleSession) TokenDecimals() (*big.Int, error) {
	return _Hashrateoracle.Contract.TokenDecimals(&_Hashrateoracle.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x3b97e856.
//
// Solidity: function tokenDecimals() view returns(uint256)
func (_Hashrateoracle *HashrateoracleCallerSession) TokenDecimals() (*big.Int, error) {
	return _Hashrateoracle.Contract.TokenDecimals(&_Hashrateoracle.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Hashrateoracle *HashrateoracleTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashrateoracle.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Hashrateoracle *HashrateoracleSession) Initialize() (*types.Transaction, error) {
	return _Hashrateoracle.Contract.Initialize(&_Hashrateoracle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Hashrateoracle *HashrateoracleTransactorSession) Initialize() (*types.Transaction, error) {
	return _Hashrateoracle.Contract.Initialize(&_Hashrateoracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hashrateoracle *HashrateoracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hashrateoracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hashrateoracle *HashrateoracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hashrateoracle.Contract.RenounceOwnership(&_Hashrateoracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hashrateoracle *HashrateoracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hashrateoracle.Contract.RenounceOwnership(&_Hashrateoracle.TransactOpts)
}

// SetBlockReward is a paid mutator transaction binding the contract method 0x1a18e707.
//
// Solidity: function setBlockReward(uint256 newBlockReward) returns()
func (_Hashrateoracle *HashrateoracleTransactor) SetBlockReward(opts *bind.TransactOpts, newBlockReward *big.Int) (*types.Transaction, error) {
	return _Hashrateoracle.contract.Transact(opts, "setBlockReward", newBlockReward)
}

// SetBlockReward is a paid mutator transaction binding the contract method 0x1a18e707.
//
// Solidity: function setBlockReward(uint256 newBlockReward) returns()
func (_Hashrateoracle *HashrateoracleSession) SetBlockReward(newBlockReward *big.Int) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.SetBlockReward(&_Hashrateoracle.TransactOpts, newBlockReward)
}

// SetBlockReward is a paid mutator transaction binding the contract method 0x1a18e707.
//
// Solidity: function setBlockReward(uint256 newBlockReward) returns()
func (_Hashrateoracle *HashrateoracleTransactorSession) SetBlockReward(newBlockReward *big.Int) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.SetBlockReward(&_Hashrateoracle.TransactOpts, newBlockReward)
}

// SetDifficulty is a paid mutator transaction binding the contract method 0x602512e1.
//
// Solidity: function setDifficulty(uint256 newDifficulty) returns()
func (_Hashrateoracle *HashrateoracleTransactor) SetDifficulty(opts *bind.TransactOpts, newDifficulty *big.Int) (*types.Transaction, error) {
	return _Hashrateoracle.contract.Transact(opts, "setDifficulty", newDifficulty)
}

// SetDifficulty is a paid mutator transaction binding the contract method 0x602512e1.
//
// Solidity: function setDifficulty(uint256 newDifficulty) returns()
func (_Hashrateoracle *HashrateoracleSession) SetDifficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.SetDifficulty(&_Hashrateoracle.TransactOpts, newDifficulty)
}

// SetDifficulty is a paid mutator transaction binding the contract method 0x602512e1.
//
// Solidity: function setDifficulty(uint256 newDifficulty) returns()
func (_Hashrateoracle *HashrateoracleTransactorSession) SetDifficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.SetDifficulty(&_Hashrateoracle.TransactOpts, newDifficulty)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hashrateoracle *HashrateoracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hashrateoracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hashrateoracle *HashrateoracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.TransferOwnership(&_Hashrateoracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hashrateoracle *HashrateoracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.TransferOwnership(&_Hashrateoracle.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Hashrateoracle *HashrateoracleTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Hashrateoracle.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Hashrateoracle *HashrateoracleSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.UpgradeToAndCall(&_Hashrateoracle.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Hashrateoracle *HashrateoracleTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Hashrateoracle.Contract.UpgradeToAndCall(&_Hashrateoracle.TransactOpts, newImplementation, data)
}

// HashrateoracleBlockRewardUpdatedIterator is returned from FilterBlockRewardUpdated and is used to iterate over the raw logs and unpacked data for BlockRewardUpdated events raised by the Hashrateoracle contract.
type HashrateoracleBlockRewardUpdatedIterator struct {
	Event *HashrateoracleBlockRewardUpdated // Event containing the contract specifics and raw log

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
func (it *HashrateoracleBlockRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashrateoracleBlockRewardUpdated)
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
		it.Event = new(HashrateoracleBlockRewardUpdated)
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
func (it *HashrateoracleBlockRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashrateoracleBlockRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashrateoracleBlockRewardUpdated represents a BlockRewardUpdated event raised by the Hashrateoracle contract.
type HashrateoracleBlockRewardUpdated struct {
	NewBlockReward *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardUpdated is a free log retrieval operation binding the contract event 0x36ea59653b93d089d49c74d420b7884cd51117c3ff1ed62ed8b98cf43835d888.
//
// Solidity: event BlockRewardUpdated(uint256 newBlockReward)
func (_Hashrateoracle *HashrateoracleFilterer) FilterBlockRewardUpdated(opts *bind.FilterOpts) (*HashrateoracleBlockRewardUpdatedIterator, error) {

	logs, sub, err := _Hashrateoracle.contract.FilterLogs(opts, "BlockRewardUpdated")
	if err != nil {
		return nil, err
	}
	return &HashrateoracleBlockRewardUpdatedIterator{contract: _Hashrateoracle.contract, event: "BlockRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockRewardUpdated is a free log subscription operation binding the contract event 0x36ea59653b93d089d49c74d420b7884cd51117c3ff1ed62ed8b98cf43835d888.
//
// Solidity: event BlockRewardUpdated(uint256 newBlockReward)
func (_Hashrateoracle *HashrateoracleFilterer) WatchBlockRewardUpdated(opts *bind.WatchOpts, sink chan<- *HashrateoracleBlockRewardUpdated) (event.Subscription, error) {

	logs, sub, err := _Hashrateoracle.contract.WatchLogs(opts, "BlockRewardUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashrateoracleBlockRewardUpdated)
				if err := _Hashrateoracle.contract.UnpackLog(event, "BlockRewardUpdated", log); err != nil {
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

// ParseBlockRewardUpdated is a log parse operation binding the contract event 0x36ea59653b93d089d49c74d420b7884cd51117c3ff1ed62ed8b98cf43835d888.
//
// Solidity: event BlockRewardUpdated(uint256 newBlockReward)
func (_Hashrateoracle *HashrateoracleFilterer) ParseBlockRewardUpdated(log types.Log) (*HashrateoracleBlockRewardUpdated, error) {
	event := new(HashrateoracleBlockRewardUpdated)
	if err := _Hashrateoracle.contract.UnpackLog(event, "BlockRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashrateoracleDifficultyUpdatedIterator is returned from FilterDifficultyUpdated and is used to iterate over the raw logs and unpacked data for DifficultyUpdated events raised by the Hashrateoracle contract.
type HashrateoracleDifficultyUpdatedIterator struct {
	Event *HashrateoracleDifficultyUpdated // Event containing the contract specifics and raw log

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
func (it *HashrateoracleDifficultyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashrateoracleDifficultyUpdated)
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
		it.Event = new(HashrateoracleDifficultyUpdated)
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
func (it *HashrateoracleDifficultyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashrateoracleDifficultyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashrateoracleDifficultyUpdated represents a DifficultyUpdated event raised by the Hashrateoracle contract.
type HashrateoracleDifficultyUpdated struct {
	NewDifficulty *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDifficultyUpdated is a free log retrieval operation binding the contract event 0x5a790c48cbebdceff3f1fcd445afd12d57302b7196738d61c60dcd491bf3efba.
//
// Solidity: event DifficultyUpdated(uint256 newDifficulty)
func (_Hashrateoracle *HashrateoracleFilterer) FilterDifficultyUpdated(opts *bind.FilterOpts) (*HashrateoracleDifficultyUpdatedIterator, error) {

	logs, sub, err := _Hashrateoracle.contract.FilterLogs(opts, "DifficultyUpdated")
	if err != nil {
		return nil, err
	}
	return &HashrateoracleDifficultyUpdatedIterator{contract: _Hashrateoracle.contract, event: "DifficultyUpdated", logs: logs, sub: sub}, nil
}

// WatchDifficultyUpdated is a free log subscription operation binding the contract event 0x5a790c48cbebdceff3f1fcd445afd12d57302b7196738d61c60dcd491bf3efba.
//
// Solidity: event DifficultyUpdated(uint256 newDifficulty)
func (_Hashrateoracle *HashrateoracleFilterer) WatchDifficultyUpdated(opts *bind.WatchOpts, sink chan<- *HashrateoracleDifficultyUpdated) (event.Subscription, error) {

	logs, sub, err := _Hashrateoracle.contract.WatchLogs(opts, "DifficultyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashrateoracleDifficultyUpdated)
				if err := _Hashrateoracle.contract.UnpackLog(event, "DifficultyUpdated", log); err != nil {
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

// ParseDifficultyUpdated is a log parse operation binding the contract event 0x5a790c48cbebdceff3f1fcd445afd12d57302b7196738d61c60dcd491bf3efba.
//
// Solidity: event DifficultyUpdated(uint256 newDifficulty)
func (_Hashrateoracle *HashrateoracleFilterer) ParseDifficultyUpdated(log types.Log) (*HashrateoracleDifficultyUpdated, error) {
	event := new(HashrateoracleDifficultyUpdated)
	if err := _Hashrateoracle.contract.UnpackLog(event, "DifficultyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashrateoracleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Hashrateoracle contract.
type HashrateoracleInitializedIterator struct {
	Event *HashrateoracleInitialized // Event containing the contract specifics and raw log

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
func (it *HashrateoracleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashrateoracleInitialized)
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
		it.Event = new(HashrateoracleInitialized)
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
func (it *HashrateoracleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashrateoracleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashrateoracleInitialized represents a Initialized event raised by the Hashrateoracle contract.
type HashrateoracleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Hashrateoracle *HashrateoracleFilterer) FilterInitialized(opts *bind.FilterOpts) (*HashrateoracleInitializedIterator, error) {

	logs, sub, err := _Hashrateoracle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HashrateoracleInitializedIterator{contract: _Hashrateoracle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Hashrateoracle *HashrateoracleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HashrateoracleInitialized) (event.Subscription, error) {

	logs, sub, err := _Hashrateoracle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashrateoracleInitialized)
				if err := _Hashrateoracle.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Hashrateoracle *HashrateoracleFilterer) ParseInitialized(log types.Log) (*HashrateoracleInitialized, error) {
	event := new(HashrateoracleInitialized)
	if err := _Hashrateoracle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashrateoracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hashrateoracle contract.
type HashrateoracleOwnershipTransferredIterator struct {
	Event *HashrateoracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HashrateoracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashrateoracleOwnershipTransferred)
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
		it.Event = new(HashrateoracleOwnershipTransferred)
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
func (it *HashrateoracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashrateoracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashrateoracleOwnershipTransferred represents a OwnershipTransferred event raised by the Hashrateoracle contract.
type HashrateoracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hashrateoracle *HashrateoracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HashrateoracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hashrateoracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HashrateoracleOwnershipTransferredIterator{contract: _Hashrateoracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hashrateoracle *HashrateoracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HashrateoracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hashrateoracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashrateoracleOwnershipTransferred)
				if err := _Hashrateoracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hashrateoracle *HashrateoracleFilterer) ParseOwnershipTransferred(log types.Log) (*HashrateoracleOwnershipTransferred, error) {
	event := new(HashrateoracleOwnershipTransferred)
	if err := _Hashrateoracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HashrateoracleUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Hashrateoracle contract.
type HashrateoracleUpgradedIterator struct {
	Event *HashrateoracleUpgraded // Event containing the contract specifics and raw log

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
func (it *HashrateoracleUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HashrateoracleUpgraded)
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
		it.Event = new(HashrateoracleUpgraded)
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
func (it *HashrateoracleUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HashrateoracleUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HashrateoracleUpgraded represents a Upgraded event raised by the Hashrateoracle contract.
type HashrateoracleUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Hashrateoracle *HashrateoracleFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*HashrateoracleUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Hashrateoracle.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &HashrateoracleUpgradedIterator{contract: _Hashrateoracle.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Hashrateoracle *HashrateoracleFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *HashrateoracleUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Hashrateoracle.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HashrateoracleUpgraded)
				if err := _Hashrateoracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Hashrateoracle *HashrateoracleFilterer) ParseUpgraded(log types.Log) (*HashrateoracleUpgraded, error) {
	event := new(HashrateoracleUpgraded)
	if err := _Hashrateoracle.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
