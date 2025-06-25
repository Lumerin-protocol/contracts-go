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

// CloneFactorySeller is an auto generated low-level Go binding around an user-defined struct.
type CloneFactorySeller struct {
	Stake *big.Int
}

// ClonefactoryMetaData contains all meta data concerning the Clonefactory contract.
var ClonefactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"clonefactoryContractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_pubkey\",\"type\":\"string\"}],\"name\":\"contractCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"}],\"name\":\"contractDeleteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"contractHardDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"purchaseInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_validatorFeeRateScaled\",\"type\":\"uint256\"}],\"name\":\"validatorFeeRateUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VALIDATOR_FEE_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"contractHardDelete\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractDurationInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_limit\",\"type\":\"uint8\"}],\"name\":\"getSellers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hashrateOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_baseImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hashrateOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_validatorFeeRateScaled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minSellerStake\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_minContractDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxContractDuration\",\"type\":\"uint32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSellerStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rentalContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"sellerByAddress\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"internalType\":\"structCloneFactory.Seller\",\"name\":\"seller\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellerDeregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stake\",\"type\":\"uint256\"}],\"name\":\"sellerRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_min\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_max\",\"type\":\"uint32\"}],\"name\":\"setContractDurationInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"}],\"name\":\"setContractsDeleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pubKey\",\"type\":\"string\"}],\"name\":\"setCreateNewRentalContractV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_hashrateOracle\",\"type\":\"address\"}],\"name\":\"setHashrateOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minSellerStake\",\"type\":\"uint256\"}],\"name\":\"setMinSellerStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"termsVersion\",\"type\":\"uint32\"}],\"name\":\"setPurchaseRentalContractV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"name\":\"setUpdateContractInformationV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_validatorFeeRateScaled\",\"type\":\"uint256\"}],\"name\":\"setValidatorFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorFeeRateScaled\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Clonefactory *ClonefactoryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Clonefactory *ClonefactorySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Clonefactory.Contract.UPGRADEINTERFACEVERSION(&_Clonefactory.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Clonefactory *ClonefactoryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Clonefactory.Contract.UPGRADEINTERFACEVERSION(&_Clonefactory.CallOpts)
}

// VALIDATORFEEDECIMALS is a free data retrieval call binding the contract method 0x7cbcd50d.
//
// Solidity: function VALIDATOR_FEE_DECIMALS() view returns(uint8)
func (_Clonefactory *ClonefactoryCaller) VALIDATORFEEDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "VALIDATOR_FEE_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORFEEDECIMALS is a free data retrieval call binding the contract method 0x7cbcd50d.
//
// Solidity: function VALIDATOR_FEE_DECIMALS() view returns(uint8)
func (_Clonefactory *ClonefactorySession) VALIDATORFEEDECIMALS() (uint8, error) {
	return _Clonefactory.Contract.VALIDATORFEEDECIMALS(&_Clonefactory.CallOpts)
}

// VALIDATORFEEDECIMALS is a free data retrieval call binding the contract method 0x7cbcd50d.
//
// Solidity: function VALIDATOR_FEE_DECIMALS() view returns(uint8)
func (_Clonefactory *ClonefactoryCallerSession) VALIDATORFEEDECIMALS() (uint8, error) {
	return _Clonefactory.Contract.VALIDATORFEEDECIMALS(&_Clonefactory.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Clonefactory *ClonefactoryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Clonefactory *ClonefactorySession) VERSION() (string, error) {
	return _Clonefactory.Contract.VERSION(&_Clonefactory.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Clonefactory *ClonefactoryCallerSession) VERSION() (string, error) {
	return _Clonefactory.Contract.VERSION(&_Clonefactory.CallOpts)
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

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Clonefactory *ClonefactoryCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Clonefactory *ClonefactorySession) FeeToken() (common.Address, error) {
	return _Clonefactory.Contract.FeeToken(&_Clonefactory.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Clonefactory *ClonefactoryCallerSession) FeeToken() (common.Address, error) {
	return _Clonefactory.Contract.FeeToken(&_Clonefactory.CallOpts)
}

// GetContractDurationInterval is a free data retrieval call binding the contract method 0x531ae12d.
//
// Solidity: function getContractDurationInterval() view returns(uint32, uint32)
func (_Clonefactory *ClonefactoryCaller) GetContractDurationInterval(opts *bind.CallOpts) (uint32, uint32, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "getContractDurationInterval")

	if err != nil {
		return *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetContractDurationInterval is a free data retrieval call binding the contract method 0x531ae12d.
//
// Solidity: function getContractDurationInterval() view returns(uint32, uint32)
func (_Clonefactory *ClonefactorySession) GetContractDurationInterval() (uint32, uint32, error) {
	return _Clonefactory.Contract.GetContractDurationInterval(&_Clonefactory.CallOpts)
}

// GetContractDurationInterval is a free data retrieval call binding the contract method 0x531ae12d.
//
// Solidity: function getContractDurationInterval() view returns(uint32, uint32)
func (_Clonefactory *ClonefactoryCallerSession) GetContractDurationInterval() (uint32, uint32, error) {
	return _Clonefactory.Contract.GetContractDurationInterval(&_Clonefactory.CallOpts)
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

// GetSellers is a free data retrieval call binding the contract method 0xf6d7b097.
//
// Solidity: function getSellers(uint256 _offset, uint8 _limit) view returns(address[])
func (_Clonefactory *ClonefactoryCaller) GetSellers(opts *bind.CallOpts, _offset *big.Int, _limit uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "getSellers", _offset, _limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSellers is a free data retrieval call binding the contract method 0xf6d7b097.
//
// Solidity: function getSellers(uint256 _offset, uint8 _limit) view returns(address[])
func (_Clonefactory *ClonefactorySession) GetSellers(_offset *big.Int, _limit uint8) ([]common.Address, error) {
	return _Clonefactory.Contract.GetSellers(&_Clonefactory.CallOpts, _offset, _limit)
}

// GetSellers is a free data retrieval call binding the contract method 0xf6d7b097.
//
// Solidity: function getSellers(uint256 _offset, uint8 _limit) view returns(address[])
func (_Clonefactory *ClonefactoryCallerSession) GetSellers(_offset *big.Int, _limit uint8) ([]common.Address, error) {
	return _Clonefactory.Contract.GetSellers(&_Clonefactory.CallOpts, _offset, _limit)
}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Clonefactory *ClonefactoryCaller) HashrateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "hashrateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Clonefactory *ClonefactorySession) HashrateOracle() (common.Address, error) {
	return _Clonefactory.Contract.HashrateOracle(&_Clonefactory.CallOpts)
}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Clonefactory *ClonefactoryCallerSession) HashrateOracle() (common.Address, error) {
	return _Clonefactory.Contract.HashrateOracle(&_Clonefactory.CallOpts)
}

// MinSellerStake is a free data retrieval call binding the contract method 0xb12d23c0.
//
// Solidity: function minSellerStake() view returns(uint256)
func (_Clonefactory *ClonefactoryCaller) MinSellerStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "minSellerStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSellerStake is a free data retrieval call binding the contract method 0xb12d23c0.
//
// Solidity: function minSellerStake() view returns(uint256)
func (_Clonefactory *ClonefactorySession) MinSellerStake() (*big.Int, error) {
	return _Clonefactory.Contract.MinSellerStake(&_Clonefactory.CallOpts)
}

// MinSellerStake is a free data retrieval call binding the contract method 0xb12d23c0.
//
// Solidity: function minSellerStake() view returns(uint256)
func (_Clonefactory *ClonefactoryCallerSession) MinSellerStake() (*big.Int, error) {
	return _Clonefactory.Contract.MinSellerStake(&_Clonefactory.CallOpts)
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

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Clonefactory *ClonefactoryCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Clonefactory *ClonefactorySession) PaymentToken() (common.Address, error) {
	return _Clonefactory.Contract.PaymentToken(&_Clonefactory.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Clonefactory *ClonefactoryCallerSession) PaymentToken() (common.Address, error) {
	return _Clonefactory.Contract.PaymentToken(&_Clonefactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Clonefactory *ClonefactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Clonefactory *ClonefactorySession) ProxiableUUID() ([32]byte, error) {
	return _Clonefactory.Contract.ProxiableUUID(&_Clonefactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Clonefactory *ClonefactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Clonefactory.Contract.ProxiableUUID(&_Clonefactory.CallOpts)
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

// SellerByAddress is a free data retrieval call binding the contract method 0x57364d44.
//
// Solidity: function sellerByAddress(address _seller) view returns((uint256) seller, bool isActive, bool isRegistered)
func (_Clonefactory *ClonefactoryCaller) SellerByAddress(opts *bind.CallOpts, _seller common.Address) (struct {
	Seller       CloneFactorySeller
	IsActive     bool
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "sellerByAddress", _seller)

	outstruct := new(struct {
		Seller       CloneFactorySeller
		IsActive     bool
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(CloneFactorySeller)).(*CloneFactorySeller)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.IsRegistered = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// SellerByAddress is a free data retrieval call binding the contract method 0x57364d44.
//
// Solidity: function sellerByAddress(address _seller) view returns((uint256) seller, bool isActive, bool isRegistered)
func (_Clonefactory *ClonefactorySession) SellerByAddress(_seller common.Address) (struct {
	Seller       CloneFactorySeller
	IsActive     bool
	IsRegistered bool
}, error) {
	return _Clonefactory.Contract.SellerByAddress(&_Clonefactory.CallOpts, _seller)
}

// SellerByAddress is a free data retrieval call binding the contract method 0x57364d44.
//
// Solidity: function sellerByAddress(address _seller) view returns((uint256) seller, bool isActive, bool isRegistered)
func (_Clonefactory *ClonefactoryCallerSession) SellerByAddress(_seller common.Address) (struct {
	Seller       CloneFactorySeller
	IsActive     bool
	IsRegistered bool
}, error) {
	return _Clonefactory.Contract.SellerByAddress(&_Clonefactory.CallOpts, _seller)
}

// ValidatorFeeRateScaled is a free data retrieval call binding the contract method 0x9e3a9bb8.
//
// Solidity: function validatorFeeRateScaled() view returns(uint256)
func (_Clonefactory *ClonefactoryCaller) ValidatorFeeRateScaled(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Clonefactory.contract.Call(opts, &out, "validatorFeeRateScaled")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorFeeRateScaled is a free data retrieval call binding the contract method 0x9e3a9bb8.
//
// Solidity: function validatorFeeRateScaled() view returns(uint256)
func (_Clonefactory *ClonefactorySession) ValidatorFeeRateScaled() (*big.Int, error) {
	return _Clonefactory.Contract.ValidatorFeeRateScaled(&_Clonefactory.CallOpts)
}

// ValidatorFeeRateScaled is a free data retrieval call binding the contract method 0x9e3a9bb8.
//
// Solidity: function validatorFeeRateScaled() view returns(uint256)
func (_Clonefactory *ClonefactoryCallerSession) ValidatorFeeRateScaled() (*big.Int, error) {
	return _Clonefactory.Contract.ValidatorFeeRateScaled(&_Clonefactory.CallOpts)
}

// ContractHardDelete is a paid mutator transaction binding the contract method 0xdd10fbd8.
//
// Solidity: function contractHardDelete(uint256 _index, address _address) returns()
func (_Clonefactory *ClonefactoryTransactor) ContractHardDelete(opts *bind.TransactOpts, _index *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "contractHardDelete", _index, _address)
}

// ContractHardDelete is a paid mutator transaction binding the contract method 0xdd10fbd8.
//
// Solidity: function contractHardDelete(uint256 _index, address _address) returns()
func (_Clonefactory *ClonefactorySession) ContractHardDelete(_index *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.ContractHardDelete(&_Clonefactory.TransactOpts, _index, _address)
}

// ContractHardDelete is a paid mutator transaction binding the contract method 0xdd10fbd8.
//
// Solidity: function contractHardDelete(uint256 _index, address _address) returns()
func (_Clonefactory *ClonefactoryTransactorSession) ContractHardDelete(_index *big.Int, _address common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.ContractHardDelete(&_Clonefactory.TransactOpts, _index, _address)
}

// Initialize is a paid mutator transaction binding the contract method 0x6beb1835.
//
// Solidity: function initialize(address _baseImplementation, address _hashrateOracle, address _paymentToken, address _feeToken, uint256 _validatorFeeRateScaled, uint256 _minSellerStake, uint32 _minContractDuration, uint32 _maxContractDuration) returns()
func (_Clonefactory *ClonefactoryTransactor) Initialize(opts *bind.TransactOpts, _baseImplementation common.Address, _hashrateOracle common.Address, _paymentToken common.Address, _feeToken common.Address, _validatorFeeRateScaled *big.Int, _minSellerStake *big.Int, _minContractDuration uint32, _maxContractDuration uint32) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "initialize", _baseImplementation, _hashrateOracle, _paymentToken, _feeToken, _validatorFeeRateScaled, _minSellerStake, _minContractDuration, _maxContractDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x6beb1835.
//
// Solidity: function initialize(address _baseImplementation, address _hashrateOracle, address _paymentToken, address _feeToken, uint256 _validatorFeeRateScaled, uint256 _minSellerStake, uint32 _minContractDuration, uint32 _maxContractDuration) returns()
func (_Clonefactory *ClonefactorySession) Initialize(_baseImplementation common.Address, _hashrateOracle common.Address, _paymentToken common.Address, _feeToken common.Address, _validatorFeeRateScaled *big.Int, _minSellerStake *big.Int, _minContractDuration uint32, _maxContractDuration uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.Initialize(&_Clonefactory.TransactOpts, _baseImplementation, _hashrateOracle, _paymentToken, _feeToken, _validatorFeeRateScaled, _minSellerStake, _minContractDuration, _maxContractDuration)
}

// Initialize is a paid mutator transaction binding the contract method 0x6beb1835.
//
// Solidity: function initialize(address _baseImplementation, address _hashrateOracle, address _paymentToken, address _feeToken, uint256 _validatorFeeRateScaled, uint256 _minSellerStake, uint32 _minContractDuration, uint32 _maxContractDuration) returns()
func (_Clonefactory *ClonefactoryTransactorSession) Initialize(_baseImplementation common.Address, _hashrateOracle common.Address, _paymentToken common.Address, _feeToken common.Address, _validatorFeeRateScaled *big.Int, _minSellerStake *big.Int, _minContractDuration uint32, _maxContractDuration uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.Initialize(&_Clonefactory.TransactOpts, _baseImplementation, _hashrateOracle, _paymentToken, _feeToken, _validatorFeeRateScaled, _minSellerStake, _minContractDuration, _maxContractDuration)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clonefactory *ClonefactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clonefactory *ClonefactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Clonefactory.Contract.RenounceOwnership(&_Clonefactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Clonefactory *ClonefactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Clonefactory.Contract.RenounceOwnership(&_Clonefactory.TransactOpts)
}

// SellerDeregister is a paid mutator transaction binding the contract method 0x4afa526f.
//
// Solidity: function sellerDeregister() returns()
func (_Clonefactory *ClonefactoryTransactor) SellerDeregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "sellerDeregister")
}

// SellerDeregister is a paid mutator transaction binding the contract method 0x4afa526f.
//
// Solidity: function sellerDeregister() returns()
func (_Clonefactory *ClonefactorySession) SellerDeregister() (*types.Transaction, error) {
	return _Clonefactory.Contract.SellerDeregister(&_Clonefactory.TransactOpts)
}

// SellerDeregister is a paid mutator transaction binding the contract method 0x4afa526f.
//
// Solidity: function sellerDeregister() returns()
func (_Clonefactory *ClonefactoryTransactorSession) SellerDeregister() (*types.Transaction, error) {
	return _Clonefactory.Contract.SellerDeregister(&_Clonefactory.TransactOpts)
}

// SellerRegister is a paid mutator transaction binding the contract method 0xfc6743c4.
//
// Solidity: function sellerRegister(uint256 _stake) returns()
func (_Clonefactory *ClonefactoryTransactor) SellerRegister(opts *bind.TransactOpts, _stake *big.Int) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "sellerRegister", _stake)
}

// SellerRegister is a paid mutator transaction binding the contract method 0xfc6743c4.
//
// Solidity: function sellerRegister(uint256 _stake) returns()
func (_Clonefactory *ClonefactorySession) SellerRegister(_stake *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SellerRegister(&_Clonefactory.TransactOpts, _stake)
}

// SellerRegister is a paid mutator transaction binding the contract method 0xfc6743c4.
//
// Solidity: function sellerRegister(uint256 _stake) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SellerRegister(_stake *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SellerRegister(&_Clonefactory.TransactOpts, _stake)
}

// SetContractDurationInterval is a paid mutator transaction binding the contract method 0x95e6b78d.
//
// Solidity: function setContractDurationInterval(uint32 _min, uint32 _max) returns()
func (_Clonefactory *ClonefactoryTransactor) SetContractDurationInterval(opts *bind.TransactOpts, _min uint32, _max uint32) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setContractDurationInterval", _min, _max)
}

// SetContractDurationInterval is a paid mutator transaction binding the contract method 0x95e6b78d.
//
// Solidity: function setContractDurationInterval(uint32 _min, uint32 _max) returns()
func (_Clonefactory *ClonefactorySession) SetContractDurationInterval(_min uint32, _max uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractDurationInterval(&_Clonefactory.TransactOpts, _min, _max)
}

// SetContractDurationInterval is a paid mutator transaction binding the contract method 0x95e6b78d.
//
// Solidity: function setContractDurationInterval(uint32 _min, uint32 _max) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetContractDurationInterval(_min uint32, _max uint32) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractDurationInterval(&_Clonefactory.TransactOpts, _min, _max)
}

// SetContractsDeleted is a paid mutator transaction binding the contract method 0x08727ec0.
//
// Solidity: function setContractsDeleted(address[] _contractAddresses, bool _isDeleted) returns()
func (_Clonefactory *ClonefactoryTransactor) SetContractsDeleted(opts *bind.TransactOpts, _contractAddresses []common.Address, _isDeleted bool) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setContractsDeleted", _contractAddresses, _isDeleted)
}

// SetContractsDeleted is a paid mutator transaction binding the contract method 0x08727ec0.
//
// Solidity: function setContractsDeleted(address[] _contractAddresses, bool _isDeleted) returns()
func (_Clonefactory *ClonefactorySession) SetContractsDeleted(_contractAddresses []common.Address, _isDeleted bool) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractsDeleted(&_Clonefactory.TransactOpts, _contractAddresses, _isDeleted)
}

// SetContractsDeleted is a paid mutator transaction binding the contract method 0x08727ec0.
//
// Solidity: function setContractsDeleted(address[] _contractAddresses, bool _isDeleted) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetContractsDeleted(_contractAddresses []common.Address, _isDeleted bool) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetContractsDeleted(&_Clonefactory.TransactOpts, _contractAddresses, _isDeleted)
}

// SetCreateNewRentalContractV2 is a paid mutator transaction binding the contract method 0x2e1e8a93.
//
// Solidity: function setCreateNewRentalContractV2(uint256 , uint256 , uint256 _speed, uint256 _length, int8 _profitTarget, address , string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactoryTransactor) SetCreateNewRentalContractV2(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, arg5 common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setCreateNewRentalContractV2", arg0, arg1, _speed, _length, _profitTarget, arg5, _pubKey)
}

// SetCreateNewRentalContractV2 is a paid mutator transaction binding the contract method 0x2e1e8a93.
//
// Solidity: function setCreateNewRentalContractV2(uint256 , uint256 , uint256 _speed, uint256 _length, int8 _profitTarget, address , string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactorySession) SetCreateNewRentalContractV2(arg0 *big.Int, arg1 *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, arg5 common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContractV2(&_Clonefactory.TransactOpts, arg0, arg1, _speed, _length, _profitTarget, arg5, _pubKey)
}

// SetCreateNewRentalContractV2 is a paid mutator transaction binding the contract method 0x2e1e8a93.
//
// Solidity: function setCreateNewRentalContractV2(uint256 , uint256 , uint256 _speed, uint256 _length, int8 _profitTarget, address , string _pubKey) payable returns(address)
func (_Clonefactory *ClonefactoryTransactorSession) SetCreateNewRentalContractV2(arg0 *big.Int, arg1 *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, arg5 common.Address, _pubKey string) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetCreateNewRentalContractV2(&_Clonefactory.TransactOpts, arg0, arg1, _speed, _length, _profitTarget, arg5, _pubKey)
}

// SetHashrateOracle is a paid mutator transaction binding the contract method 0xf0f55462.
//
// Solidity: function setHashrateOracle(address _hashrateOracle) returns()
func (_Clonefactory *ClonefactoryTransactor) SetHashrateOracle(opts *bind.TransactOpts, _hashrateOracle common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setHashrateOracle", _hashrateOracle)
}

// SetHashrateOracle is a paid mutator transaction binding the contract method 0xf0f55462.
//
// Solidity: function setHashrateOracle(address _hashrateOracle) returns()
func (_Clonefactory *ClonefactorySession) SetHashrateOracle(_hashrateOracle common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetHashrateOracle(&_Clonefactory.TransactOpts, _hashrateOracle)
}

// SetHashrateOracle is a paid mutator transaction binding the contract method 0xf0f55462.
//
// Solidity: function setHashrateOracle(address _hashrateOracle) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetHashrateOracle(_hashrateOracle common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetHashrateOracle(&_Clonefactory.TransactOpts, _hashrateOracle)
}

// SetMinSellerStake is a paid mutator transaction binding the contract method 0x8d3728a2.
//
// Solidity: function setMinSellerStake(uint256 _minSellerStake) returns()
func (_Clonefactory *ClonefactoryTransactor) SetMinSellerStake(opts *bind.TransactOpts, _minSellerStake *big.Int) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setMinSellerStake", _minSellerStake)
}

// SetMinSellerStake is a paid mutator transaction binding the contract method 0x8d3728a2.
//
// Solidity: function setMinSellerStake(uint256 _minSellerStake) returns()
func (_Clonefactory *ClonefactorySession) SetMinSellerStake(_minSellerStake *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetMinSellerStake(&_Clonefactory.TransactOpts, _minSellerStake)
}

// SetMinSellerStake is a paid mutator transaction binding the contract method 0x8d3728a2.
//
// Solidity: function setMinSellerStake(uint256 _minSellerStake) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetMinSellerStake(_minSellerStake *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetMinSellerStake(&_Clonefactory.TransactOpts, _minSellerStake)
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

// SetUpdateContractInformationV2 is a paid mutator transaction binding the contract method 0x3a6738b2.
//
// Solidity: function setUpdateContractInformationV2(address _contractAddress, uint256 , uint256 , uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Clonefactory *ClonefactoryTransactor) SetUpdateContractInformationV2(opts *bind.TransactOpts, _contractAddress common.Address, arg1 *big.Int, arg2 *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setUpdateContractInformationV2", _contractAddress, arg1, arg2, _speed, _length, _profitTarget)
}

// SetUpdateContractInformationV2 is a paid mutator transaction binding the contract method 0x3a6738b2.
//
// Solidity: function setUpdateContractInformationV2(address _contractAddress, uint256 , uint256 , uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Clonefactory *ClonefactorySession) SetUpdateContractInformationV2(_contractAddress common.Address, arg1 *big.Int, arg2 *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetUpdateContractInformationV2(&_Clonefactory.TransactOpts, _contractAddress, arg1, arg2, _speed, _length, _profitTarget)
}

// SetUpdateContractInformationV2 is a paid mutator transaction binding the contract method 0x3a6738b2.
//
// Solidity: function setUpdateContractInformationV2(address _contractAddress, uint256 , uint256 , uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetUpdateContractInformationV2(_contractAddress common.Address, arg1 *big.Int, arg2 *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetUpdateContractInformationV2(&_Clonefactory.TransactOpts, _contractAddress, arg1, arg2, _speed, _length, _profitTarget)
}

// SetValidatorFeeRate is a paid mutator transaction binding the contract method 0x93feb7bd.
//
// Solidity: function setValidatorFeeRate(uint256 _validatorFeeRateScaled) returns()
func (_Clonefactory *ClonefactoryTransactor) SetValidatorFeeRate(opts *bind.TransactOpts, _validatorFeeRateScaled *big.Int) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "setValidatorFeeRate", _validatorFeeRateScaled)
}

// SetValidatorFeeRate is a paid mutator transaction binding the contract method 0x93feb7bd.
//
// Solidity: function setValidatorFeeRate(uint256 _validatorFeeRateScaled) returns()
func (_Clonefactory *ClonefactorySession) SetValidatorFeeRate(_validatorFeeRateScaled *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetValidatorFeeRate(&_Clonefactory.TransactOpts, _validatorFeeRateScaled)
}

// SetValidatorFeeRate is a paid mutator transaction binding the contract method 0x93feb7bd.
//
// Solidity: function setValidatorFeeRate(uint256 _validatorFeeRateScaled) returns()
func (_Clonefactory *ClonefactoryTransactorSession) SetValidatorFeeRate(_validatorFeeRateScaled *big.Int) (*types.Transaction, error) {
	return _Clonefactory.Contract.SetValidatorFeeRate(&_Clonefactory.TransactOpts, _validatorFeeRateScaled)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clonefactory *ClonefactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clonefactory *ClonefactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.TransferOwnership(&_Clonefactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Clonefactory *ClonefactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Clonefactory.Contract.TransferOwnership(&_Clonefactory.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Clonefactory *ClonefactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Clonefactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Clonefactory *ClonefactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Clonefactory.Contract.UpgradeToAndCall(&_Clonefactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Clonefactory *ClonefactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Clonefactory.Contract.UpgradeToAndCall(&_Clonefactory.TransactOpts, newImplementation, data)
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
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Clonefactory *ClonefactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ClonefactoryInitializedIterator, error) {

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ClonefactoryInitializedIterator{contract: _Clonefactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Clonefactory *ClonefactoryFilterer) ParseInitialized(log types.Log) (*ClonefactoryInitialized, error) {
	event := new(ClonefactoryInitialized)
	if err := _Clonefactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonefactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Clonefactory contract.
type ClonefactoryOwnershipTransferredIterator struct {
	Event *ClonefactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClonefactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryOwnershipTransferred)
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
		it.Event = new(ClonefactoryOwnershipTransferred)
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
func (it *ClonefactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Clonefactory contract.
type ClonefactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clonefactory *ClonefactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClonefactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryOwnershipTransferredIterator{contract: _Clonefactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Clonefactory *ClonefactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClonefactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryOwnershipTransferred)
				if err := _Clonefactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Clonefactory *ClonefactoryFilterer) ParseOwnershipTransferred(log types.Log) (*ClonefactoryOwnershipTransferred, error) {
	event := new(ClonefactoryOwnershipTransferred)
	if err := _Clonefactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonefactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Clonefactory contract.
type ClonefactoryUpgradedIterator struct {
	Event *ClonefactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *ClonefactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryUpgraded)
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
		it.Event = new(ClonefactoryUpgraded)
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
func (it *ClonefactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryUpgraded represents a Upgraded event raised by the Clonefactory contract.
type ClonefactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Clonefactory *ClonefactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ClonefactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryUpgradedIterator{contract: _Clonefactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Clonefactory *ClonefactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ClonefactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryUpgraded)
				if err := _Clonefactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Clonefactory *ClonefactoryFilterer) ParseUpgraded(log types.Log) (*ClonefactoryUpgraded, error) {
	event := new(ClonefactoryUpgraded)
	if err := _Clonefactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
	Address   common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClonefactoryContractPurchased is a free log retrieval operation binding the contract event 0xff1dec24325264c678c99ba82c9b203272dc942010bfb7941b47a72917884753.
//
// Solidity: event clonefactoryContractPurchased(address indexed _address, address indexed _validator)
func (_Clonefactory *ClonefactoryFilterer) FilterClonefactoryContractPurchased(opts *bind.FilterOpts, _address []common.Address, _validator []common.Address) (*ClonefactoryClonefactoryContractPurchasedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}
	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "clonefactoryContractPurchased", _addressRule, _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryClonefactoryContractPurchasedIterator{contract: _Clonefactory.contract, event: "clonefactoryContractPurchased", logs: logs, sub: sub}, nil
}

// WatchClonefactoryContractPurchased is a free log subscription operation binding the contract event 0xff1dec24325264c678c99ba82c9b203272dc942010bfb7941b47a72917884753.
//
// Solidity: event clonefactoryContractPurchased(address indexed _address, address indexed _validator)
func (_Clonefactory *ClonefactoryFilterer) WatchClonefactoryContractPurchased(opts *bind.WatchOpts, sink chan<- *ClonefactoryClonefactoryContractPurchased, _address []common.Address, _validator []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}
	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "clonefactoryContractPurchased", _addressRule, _validatorRule)
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

// ParseClonefactoryContractPurchased is a log parse operation binding the contract event 0xff1dec24325264c678c99ba82c9b203272dc942010bfb7941b47a72917884753.
//
// Solidity: event clonefactoryContractPurchased(address indexed _address, address indexed _validator)
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

// ClonefactoryContractHardDeletedIterator is returned from FilterContractHardDeleted and is used to iterate over the raw logs and unpacked data for ContractHardDeleted events raised by the Clonefactory contract.
type ClonefactoryContractHardDeletedIterator struct {
	Event *ClonefactoryContractHardDeleted // Event containing the contract specifics and raw log

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
func (it *ClonefactoryContractHardDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryContractHardDeleted)
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
		it.Event = new(ClonefactoryContractHardDeleted)
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
func (it *ClonefactoryContractHardDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryContractHardDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryContractHardDeleted represents a ContractHardDeleted event raised by the Clonefactory contract.
type ClonefactoryContractHardDeleted struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractHardDeleted is a free log retrieval operation binding the contract event 0x4aa35bed19905c9f8456f7cde36306d10516bf490a71f22b11421e94c9794b81.
//
// Solidity: event contractHardDeleted(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) FilterContractHardDeleted(opts *bind.FilterOpts, _address []common.Address) (*ClonefactoryContractHardDeletedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "contractHardDeleted", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ClonefactoryContractHardDeletedIterator{contract: _Clonefactory.contract, event: "contractHardDeleted", logs: logs, sub: sub}, nil
}

// WatchContractHardDeleted is a free log subscription operation binding the contract event 0x4aa35bed19905c9f8456f7cde36306d10516bf490a71f22b11421e94c9794b81.
//
// Solidity: event contractHardDeleted(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) WatchContractHardDeleted(opts *bind.WatchOpts, sink chan<- *ClonefactoryContractHardDeleted, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "contractHardDeleted", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryContractHardDeleted)
				if err := _Clonefactory.contract.UnpackLog(event, "contractHardDeleted", log); err != nil {
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

// ParseContractHardDeleted is a log parse operation binding the contract event 0x4aa35bed19905c9f8456f7cde36306d10516bf490a71f22b11421e94c9794b81.
//
// Solidity: event contractHardDeleted(address indexed _address)
func (_Clonefactory *ClonefactoryFilterer) ParseContractHardDeleted(log types.Log) (*ClonefactoryContractHardDeleted, error) {
	event := new(ClonefactoryContractHardDeleted)
	if err := _Clonefactory.contract.UnpackLog(event, "contractHardDeleted", log); err != nil {
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

// ClonefactoryValidatorFeeRateUpdatedIterator is returned from FilterValidatorFeeRateUpdated and is used to iterate over the raw logs and unpacked data for ValidatorFeeRateUpdated events raised by the Clonefactory contract.
type ClonefactoryValidatorFeeRateUpdatedIterator struct {
	Event *ClonefactoryValidatorFeeRateUpdated // Event containing the contract specifics and raw log

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
func (it *ClonefactoryValidatorFeeRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClonefactoryValidatorFeeRateUpdated)
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
		it.Event = new(ClonefactoryValidatorFeeRateUpdated)
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
func (it *ClonefactoryValidatorFeeRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClonefactoryValidatorFeeRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClonefactoryValidatorFeeRateUpdated represents a ValidatorFeeRateUpdated event raised by the Clonefactory contract.
type ClonefactoryValidatorFeeRateUpdated struct {
	ValidatorFeeRateScaled *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterValidatorFeeRateUpdated is a free log retrieval operation binding the contract event 0xa765cde4b738d0ceba8998592a4ba3c70577da605d79b49decf1067cfb6bf08f.
//
// Solidity: event validatorFeeRateUpdated(uint256 _validatorFeeRateScaled)
func (_Clonefactory *ClonefactoryFilterer) FilterValidatorFeeRateUpdated(opts *bind.FilterOpts) (*ClonefactoryValidatorFeeRateUpdatedIterator, error) {

	logs, sub, err := _Clonefactory.contract.FilterLogs(opts, "validatorFeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return &ClonefactoryValidatorFeeRateUpdatedIterator{contract: _Clonefactory.contract, event: "validatorFeeRateUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorFeeRateUpdated is a free log subscription operation binding the contract event 0xa765cde4b738d0ceba8998592a4ba3c70577da605d79b49decf1067cfb6bf08f.
//
// Solidity: event validatorFeeRateUpdated(uint256 _validatorFeeRateScaled)
func (_Clonefactory *ClonefactoryFilterer) WatchValidatorFeeRateUpdated(opts *bind.WatchOpts, sink chan<- *ClonefactoryValidatorFeeRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Clonefactory.contract.WatchLogs(opts, "validatorFeeRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClonefactoryValidatorFeeRateUpdated)
				if err := _Clonefactory.contract.UnpackLog(event, "validatorFeeRateUpdated", log); err != nil {
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

// ParseValidatorFeeRateUpdated is a log parse operation binding the contract event 0xa765cde4b738d0ceba8998592a4ba3c70577da605d79b49decf1067cfb6bf08f.
//
// Solidity: event validatorFeeRateUpdated(uint256 _validatorFeeRateScaled)
func (_Clonefactory *ClonefactoryFilterer) ParseValidatorFeeRateUpdated(log types.Log) (*ClonefactoryValidatorFeeRateUpdated, error) {
	event := new(ClonefactoryValidatorFeeRateUpdated)
	if err := _Clonefactory.contract.UnpackLog(event, "validatorFeeRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
