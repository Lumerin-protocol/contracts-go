// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validatorregistry

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

// ValidatorRegistryValidator is an auto generated low-level Go binding around an user-defined struct.
type ValidatorRegistryValidator struct {
	Stake          *big.Int
	Addr           common.Address
	PubKeyYparity  bool
	LastComplainer common.Address
	Complains      uint8
	Host           string
	PubKeyX        [32]byte
}

// ValidatorregistryMetaData contains all meta data concerning the Validatorregistry contract.
var ValidatorregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyComplained\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HostTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"complainer\",\"type\":\"address\"}],\"name\":\"ValidatorComplain\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorPunished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRegisteredUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeValidatorsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"forceUpdateActive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"limit\",\"type\":\"uint8\"}],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pubKeyYparity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"lastComplainer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"complains\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorRegistry.Validator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getValidatorV2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pubKeyYparity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"lastComplainer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"complains\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"}],\"internalType\":\"structValidatorRegistry.Validator\",\"name\":\"validator\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"limit\",\"type\":\"uint8\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeMinimun\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeRegister\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_punishAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_punishThreshold\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punishAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punishThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setPunishAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"val\",\"type\":\"uint8\"}],\"name\":\"setPunishThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setStakeMinimum\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setStakeRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeMinimum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRegister\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"validatorComplain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorDeregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pubKeyYparity\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"}],\"name\":\"validatorRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"pubKeyYparity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"lastComplainer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"complains\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"host\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"pubKeyX\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ValidatorregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorregistryMetaData.ABI instead.
var ValidatorregistryABI = ValidatorregistryMetaData.ABI

// Validatorregistry is an auto generated Go binding around an Ethereum contract.
type Validatorregistry struct {
	ValidatorregistryCaller     // Read-only binding to the contract
	ValidatorregistryTransactor // Write-only binding to the contract
	ValidatorregistryFilterer   // Log filterer for contract events
}

// ValidatorregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorregistrySession struct {
	Contract     *Validatorregistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ValidatorregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorregistryCallerSession struct {
	Contract *ValidatorregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ValidatorregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorregistryTransactorSession struct {
	Contract     *ValidatorregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ValidatorregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorregistryRaw struct {
	Contract *Validatorregistry // Generic contract binding to access the raw methods on
}

// ValidatorregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorregistryCallerRaw struct {
	Contract *ValidatorregistryCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorregistryTransactorRaw struct {
	Contract *ValidatorregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorregistry creates a new instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistry(address common.Address, backend bind.ContractBackend) (*Validatorregistry, error) {
	contract, err := bindValidatorregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validatorregistry{ValidatorregistryCaller: ValidatorregistryCaller{contract: contract}, ValidatorregistryTransactor: ValidatorregistryTransactor{contract: contract}, ValidatorregistryFilterer: ValidatorregistryFilterer{contract: contract}}, nil
}

// NewValidatorregistryCaller creates a new read-only instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryCaller(address common.Address, caller bind.ContractCaller) (*ValidatorregistryCaller, error) {
	contract, err := bindValidatorregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryCaller{contract: contract}, nil
}

// NewValidatorregistryTransactor creates a new write-only instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorregistryTransactor, error) {
	contract, err := bindValidatorregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryTransactor{contract: contract}, nil
}

// NewValidatorregistryFilterer creates a new log filterer instance of Validatorregistry, bound to a specific deployed contract.
func NewValidatorregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorregistryFilterer, error) {
	contract, err := bindValidatorregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryFilterer{contract: contract}, nil
}

// bindValidatorregistry binds a generic wrapper to an already deployed contract.
func bindValidatorregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorregistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorregistry *ValidatorregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorregistry.Contract.ValidatorregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorregistry *ValidatorregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorregistry *ValidatorregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validatorregistry *ValidatorregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validatorregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validatorregistry *ValidatorregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validatorregistry *ValidatorregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validatorregistry.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Validatorregistry *ValidatorregistryCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Validatorregistry *ValidatorregistrySession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Validatorregistry.Contract.UPGRADEINTERFACEVERSION(&_Validatorregistry.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Validatorregistry *ValidatorregistryCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Validatorregistry.Contract.UPGRADEINTERFACEVERSION(&_Validatorregistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Validatorregistry *ValidatorregistryCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Validatorregistry *ValidatorregistrySession) VERSION() (string, error) {
	return _Validatorregistry.Contract.VERSION(&_Validatorregistry.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Validatorregistry *ValidatorregistryCallerSession) VERSION() (string, error) {
	return _Validatorregistry.Contract.VERSION(&_Validatorregistry.CallOpts)
}

// ActiveValidatorsLength is a free data retrieval call binding the contract method 0x75bac430.
//
// Solidity: function activeValidatorsLength() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) ActiveValidatorsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "activeValidatorsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveValidatorsLength is a free data retrieval call binding the contract method 0x75bac430.
//
// Solidity: function activeValidatorsLength() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) ActiveValidatorsLength() (*big.Int, error) {
	return _Validatorregistry.Contract.ActiveValidatorsLength(&_Validatorregistry.CallOpts)
}

// ActiveValidatorsLength is a free data retrieval call binding the contract method 0x75bac430.
//
// Solidity: function activeValidatorsLength() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) ActiveValidatorsLength() (*big.Int, error) {
	return _Validatorregistry.Contract.ActiveValidatorsLength(&_Validatorregistry.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x290d263f.
//
// Solidity: function getActiveValidators(uint256 offset, uint8 limit) view returns(address[])
func (_Validatorregistry *ValidatorregistryCaller) GetActiveValidators(opts *bind.CallOpts, offset *big.Int, limit uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getActiveValidators", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x290d263f.
//
// Solidity: function getActiveValidators(uint256 offset, uint8 limit) view returns(address[])
func (_Validatorregistry *ValidatorregistrySession) GetActiveValidators(offset *big.Int, limit uint8) ([]common.Address, error) {
	return _Validatorregistry.Contract.GetActiveValidators(&_Validatorregistry.CallOpts, offset, limit)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x290d263f.
//
// Solidity: function getActiveValidators(uint256 offset, uint8 limit) view returns(address[])
func (_Validatorregistry *ValidatorregistryCallerSession) GetActiveValidators(offset *big.Int, limit uint8) ([]common.Address, error) {
	return _Validatorregistry.Contract.GetActiveValidators(&_Validatorregistry.CallOpts, offset, limit)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address addr) view returns((uint256,address,bool,address,uint8,string,bytes32))
func (_Validatorregistry *ValidatorregistryCaller) GetValidator(opts *bind.CallOpts, addr common.Address) (ValidatorRegistryValidator, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getValidator", addr)

	if err != nil {
		return *new(ValidatorRegistryValidator), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorRegistryValidator)).(*ValidatorRegistryValidator)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address addr) view returns((uint256,address,bool,address,uint8,string,bytes32))
func (_Validatorregistry *ValidatorregistrySession) GetValidator(addr common.Address) (ValidatorRegistryValidator, error) {
	return _Validatorregistry.Contract.GetValidator(&_Validatorregistry.CallOpts, addr)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address addr) view returns((uint256,address,bool,address,uint8,string,bytes32))
func (_Validatorregistry *ValidatorregistryCallerSession) GetValidator(addr common.Address) (ValidatorRegistryValidator, error) {
	return _Validatorregistry.Contract.GetValidator(&_Validatorregistry.CallOpts, addr)
}

// GetValidatorV2 is a free data retrieval call binding the contract method 0xc9d8d888.
//
// Solidity: function getValidatorV2(address addr) view returns((uint256,address,bool,address,uint8,string,bytes32) validator, bool isActive, bool isRegistered)
func (_Validatorregistry *ValidatorregistryCaller) GetValidatorV2(opts *bind.CallOpts, addr common.Address) (struct {
	Validator    ValidatorRegistryValidator
	IsActive     bool
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getValidatorV2", addr)

	outstruct := new(struct {
		Validator    ValidatorRegistryValidator
		IsActive     bool
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validator = *abi.ConvertType(out[0], new(ValidatorRegistryValidator)).(*ValidatorRegistryValidator)
	outstruct.IsActive = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.IsRegistered = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidatorV2 is a free data retrieval call binding the contract method 0xc9d8d888.
//
// Solidity: function getValidatorV2(address addr) view returns((uint256,address,bool,address,uint8,string,bytes32) validator, bool isActive, bool isRegistered)
func (_Validatorregistry *ValidatorregistrySession) GetValidatorV2(addr common.Address) (struct {
	Validator    ValidatorRegistryValidator
	IsActive     bool
	IsRegistered bool
}, error) {
	return _Validatorregistry.Contract.GetValidatorV2(&_Validatorregistry.CallOpts, addr)
}

// GetValidatorV2 is a free data retrieval call binding the contract method 0xc9d8d888.
//
// Solidity: function getValidatorV2(address addr) view returns((uint256,address,bool,address,uint8,string,bytes32) validator, bool isActive, bool isRegistered)
func (_Validatorregistry *ValidatorregistryCallerSession) GetValidatorV2(addr common.Address) (struct {
	Validator    ValidatorRegistryValidator
	IsActive     bool
	IsRegistered bool
}, error) {
	return _Validatorregistry.Contract.GetValidatorV2(&_Validatorregistry.CallOpts, addr)
}

// GetValidators is a free data retrieval call binding the contract method 0xafe3749d.
//
// Solidity: function getValidators(uint256 offset, uint8 limit) view returns(address[])
func (_Validatorregistry *ValidatorregistryCaller) GetValidators(opts *bind.CallOpts, offset *big.Int, limit uint8) ([]common.Address, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "getValidators", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xafe3749d.
//
// Solidity: function getValidators(uint256 offset, uint8 limit) view returns(address[])
func (_Validatorregistry *ValidatorregistrySession) GetValidators(offset *big.Int, limit uint8) ([]common.Address, error) {
	return _Validatorregistry.Contract.GetValidators(&_Validatorregistry.CallOpts, offset, limit)
}

// GetValidators is a free data retrieval call binding the contract method 0xafe3749d.
//
// Solidity: function getValidators(uint256 offset, uint8 limit) view returns(address[])
func (_Validatorregistry *ValidatorregistryCallerSession) GetValidators(offset *big.Int, limit uint8) ([]common.Address, error) {
	return _Validatorregistry.Contract.GetValidators(&_Validatorregistry.CallOpts, offset, limit)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validatorregistry *ValidatorregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validatorregistry *ValidatorregistrySession) Owner() (common.Address, error) {
	return _Validatorregistry.Contract.Owner(&_Validatorregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Validatorregistry *ValidatorregistryCallerSession) Owner() (common.Address, error) {
	return _Validatorregistry.Contract.Owner(&_Validatorregistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Validatorregistry *ValidatorregistryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Validatorregistry *ValidatorregistrySession) ProxiableUUID() ([32]byte, error) {
	return _Validatorregistry.Contract.ProxiableUUID(&_Validatorregistry.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Validatorregistry *ValidatorregistryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Validatorregistry.Contract.ProxiableUUID(&_Validatorregistry.CallOpts)
}

// PunishAmount is a free data retrieval call binding the contract method 0x3fa1b5bc.
//
// Solidity: function punishAmount() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) PunishAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "punishAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishAmount is a free data retrieval call binding the contract method 0x3fa1b5bc.
//
// Solidity: function punishAmount() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) PunishAmount() (*big.Int, error) {
	return _Validatorregistry.Contract.PunishAmount(&_Validatorregistry.CallOpts)
}

// PunishAmount is a free data retrieval call binding the contract method 0x3fa1b5bc.
//
// Solidity: function punishAmount() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) PunishAmount() (*big.Int, error) {
	return _Validatorregistry.Contract.PunishAmount(&_Validatorregistry.CallOpts)
}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint8)
func (_Validatorregistry *ValidatorregistryCaller) PunishThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "punishThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint8)
func (_Validatorregistry *ValidatorregistrySession) PunishThreshold() (uint8, error) {
	return _Validatorregistry.Contract.PunishThreshold(&_Validatorregistry.CallOpts)
}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint8)
func (_Validatorregistry *ValidatorregistryCallerSession) PunishThreshold() (uint8, error) {
	return _Validatorregistry.Contract.PunishThreshold(&_Validatorregistry.CallOpts)
}

// StakeMinimum is a free data retrieval call binding the contract method 0x4b690429.
//
// Solidity: function stakeMinimum() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) StakeMinimum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "stakeMinimum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeMinimum is a free data retrieval call binding the contract method 0x4b690429.
//
// Solidity: function stakeMinimum() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) StakeMinimum() (*big.Int, error) {
	return _Validatorregistry.Contract.StakeMinimum(&_Validatorregistry.CallOpts)
}

// StakeMinimum is a free data retrieval call binding the contract method 0x4b690429.
//
// Solidity: function stakeMinimum() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) StakeMinimum() (*big.Int, error) {
	return _Validatorregistry.Contract.StakeMinimum(&_Validatorregistry.CallOpts)
}

// StakeRegister is a free data retrieval call binding the contract method 0x1aa25cda.
//
// Solidity: function stakeRegister() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) StakeRegister(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "stakeRegister")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeRegister is a free data retrieval call binding the contract method 0x1aa25cda.
//
// Solidity: function stakeRegister() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) StakeRegister() (*big.Int, error) {
	return _Validatorregistry.Contract.StakeRegister(&_Validatorregistry.CallOpts)
}

// StakeRegister is a free data retrieval call binding the contract method 0x1aa25cda.
//
// Solidity: function stakeRegister() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) StakeRegister() (*big.Int, error) {
	return _Validatorregistry.Contract.StakeRegister(&_Validatorregistry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Validatorregistry *ValidatorregistryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Validatorregistry *ValidatorregistrySession) Token() (common.Address, error) {
	return _Validatorregistry.Contract.Token(&_Validatorregistry.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Validatorregistry *ValidatorregistryCallerSession) Token() (common.Address, error) {
	return _Validatorregistry.Contract.Token(&_Validatorregistry.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) TotalStake() (*big.Int, error) {
	return _Validatorregistry.Contract.TotalStake(&_Validatorregistry.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) TotalStake() (*big.Int, error) {
	return _Validatorregistry.Contract.TotalStake(&_Validatorregistry.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, address addr, bool pubKeyYparity, address lastComplainer, uint8 complains, string host, bytes32 pubKeyX)
func (_Validatorregistry *ValidatorregistryCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Stake          *big.Int
	Addr           common.Address
	PubKeyYparity  bool
	LastComplainer common.Address
	Complains      uint8
	Host           string
	PubKeyX        [32]byte
}, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Stake          *big.Int
		Addr           common.Address
		PubKeyYparity  bool
		LastComplainer common.Address
		Complains      uint8
		Host           string
		PubKeyX        [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stake = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Addr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PubKeyYparity = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.LastComplainer = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Complains = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Host = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.PubKeyX = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, address addr, bool pubKeyYparity, address lastComplainer, uint8 complains, string host, bytes32 pubKeyX)
func (_Validatorregistry *ValidatorregistrySession) Validators(arg0 common.Address) (struct {
	Stake          *big.Int
	Addr           common.Address
	PubKeyYparity  bool
	LastComplainer common.Address
	Complains      uint8
	Host           string
	PubKeyX        [32]byte
}, error) {
	return _Validatorregistry.Contract.Validators(&_Validatorregistry.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(uint256 stake, address addr, bool pubKeyYparity, address lastComplainer, uint8 complains, string host, bytes32 pubKeyX)
func (_Validatorregistry *ValidatorregistryCallerSession) Validators(arg0 common.Address) (struct {
	Stake          *big.Int
	Addr           common.Address
	PubKeyYparity  bool
	LastComplainer common.Address
	Complains      uint8
	Host           string
	PubKeyX        [32]byte
}, error) {
	return _Validatorregistry.Contract.Validators(&_Validatorregistry.CallOpts, arg0)
}

// ValidatorsLength is a free data retrieval call binding the contract method 0xf1105a7e.
//
// Solidity: function validatorsLength() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCaller) ValidatorsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validatorregistry.contract.Call(opts, &out, "validatorsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorsLength is a free data retrieval call binding the contract method 0xf1105a7e.
//
// Solidity: function validatorsLength() view returns(uint256)
func (_Validatorregistry *ValidatorregistrySession) ValidatorsLength() (*big.Int, error) {
	return _Validatorregistry.Contract.ValidatorsLength(&_Validatorregistry.CallOpts)
}

// ValidatorsLength is a free data retrieval call binding the contract method 0xf1105a7e.
//
// Solidity: function validatorsLength() view returns(uint256)
func (_Validatorregistry *ValidatorregistryCallerSession) ValidatorsLength() (*big.Int, error) {
	return _Validatorregistry.Contract.ValidatorsLength(&_Validatorregistry.CallOpts)
}

// ForceUpdateActive is a paid mutator transaction binding the contract method 0xd864b499.
//
// Solidity: function forceUpdateActive(address validator) returns()
func (_Validatorregistry *ValidatorregistryTransactor) ForceUpdateActive(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "forceUpdateActive", validator)
}

// ForceUpdateActive is a paid mutator transaction binding the contract method 0xd864b499.
//
// Solidity: function forceUpdateActive(address validator) returns()
func (_Validatorregistry *ValidatorregistrySession) ForceUpdateActive(validator common.Address) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ForceUpdateActive(&_Validatorregistry.TransactOpts, validator)
}

// ForceUpdateActive is a paid mutator transaction binding the contract method 0xd864b499.
//
// Solidity: function forceUpdateActive(address validator) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) ForceUpdateActive(validator common.Address) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ForceUpdateActive(&_Validatorregistry.TransactOpts, validator)
}

// Initialize is a paid mutator transaction binding the contract method 0x79926848.
//
// Solidity: function initialize(address _token, uint256 _stakeMinimun, uint256 _stakeRegister, uint256 _punishAmount, uint8 _punishThreshold) returns()
func (_Validatorregistry *ValidatorregistryTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _stakeMinimun *big.Int, _stakeRegister *big.Int, _punishAmount *big.Int, _punishThreshold uint8) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "initialize", _token, _stakeMinimun, _stakeRegister, _punishAmount, _punishThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x79926848.
//
// Solidity: function initialize(address _token, uint256 _stakeMinimun, uint256 _stakeRegister, uint256 _punishAmount, uint8 _punishThreshold) returns()
func (_Validatorregistry *ValidatorregistrySession) Initialize(_token common.Address, _stakeMinimun *big.Int, _stakeRegister *big.Int, _punishAmount *big.Int, _punishThreshold uint8) (*types.Transaction, error) {
	return _Validatorregistry.Contract.Initialize(&_Validatorregistry.TransactOpts, _token, _stakeMinimun, _stakeRegister, _punishAmount, _punishThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x79926848.
//
// Solidity: function initialize(address _token, uint256 _stakeMinimun, uint256 _stakeRegister, uint256 _punishAmount, uint8 _punishThreshold) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) Initialize(_token common.Address, _stakeMinimun *big.Int, _stakeRegister *big.Int, _punishAmount *big.Int, _punishThreshold uint8) (*types.Transaction, error) {
	return _Validatorregistry.Contract.Initialize(&_Validatorregistry.TransactOpts, _token, _stakeMinimun, _stakeRegister, _punishAmount, _punishThreshold)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validatorregistry *ValidatorregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validatorregistry *ValidatorregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Validatorregistry.Contract.RenounceOwnership(&_Validatorregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validatorregistry.Contract.RenounceOwnership(&_Validatorregistry.TransactOpts)
}

// SetPunishAmount is a paid mutator transaction binding the contract method 0x82bdad8f.
//
// Solidity: function setPunishAmount(uint256 val) returns()
func (_Validatorregistry *ValidatorregistryTransactor) SetPunishAmount(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "setPunishAmount", val)
}

// SetPunishAmount is a paid mutator transaction binding the contract method 0x82bdad8f.
//
// Solidity: function setPunishAmount(uint256 val) returns()
func (_Validatorregistry *ValidatorregistrySession) SetPunishAmount(val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetPunishAmount(&_Validatorregistry.TransactOpts, val)
}

// SetPunishAmount is a paid mutator transaction binding the contract method 0x82bdad8f.
//
// Solidity: function setPunishAmount(uint256 val) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) SetPunishAmount(val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetPunishAmount(&_Validatorregistry.TransactOpts, val)
}

// SetPunishThreshold is a paid mutator transaction binding the contract method 0x58ea7e60.
//
// Solidity: function setPunishThreshold(uint8 val) returns()
func (_Validatorregistry *ValidatorregistryTransactor) SetPunishThreshold(opts *bind.TransactOpts, val uint8) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "setPunishThreshold", val)
}

// SetPunishThreshold is a paid mutator transaction binding the contract method 0x58ea7e60.
//
// Solidity: function setPunishThreshold(uint8 val) returns()
func (_Validatorregistry *ValidatorregistrySession) SetPunishThreshold(val uint8) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetPunishThreshold(&_Validatorregistry.TransactOpts, val)
}

// SetPunishThreshold is a paid mutator transaction binding the contract method 0x58ea7e60.
//
// Solidity: function setPunishThreshold(uint8 val) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) SetPunishThreshold(val uint8) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetPunishThreshold(&_Validatorregistry.TransactOpts, val)
}

// SetStakeMinimum is a paid mutator transaction binding the contract method 0x941d172e.
//
// Solidity: function setStakeMinimum(uint256 val) returns()
func (_Validatorregistry *ValidatorregistryTransactor) SetStakeMinimum(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "setStakeMinimum", val)
}

// SetStakeMinimum is a paid mutator transaction binding the contract method 0x941d172e.
//
// Solidity: function setStakeMinimum(uint256 val) returns()
func (_Validatorregistry *ValidatorregistrySession) SetStakeMinimum(val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetStakeMinimum(&_Validatorregistry.TransactOpts, val)
}

// SetStakeMinimum is a paid mutator transaction binding the contract method 0x941d172e.
//
// Solidity: function setStakeMinimum(uint256 val) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) SetStakeMinimum(val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetStakeMinimum(&_Validatorregistry.TransactOpts, val)
}

// SetStakeRegister is a paid mutator transaction binding the contract method 0xcbdc17cb.
//
// Solidity: function setStakeRegister(uint256 val) returns()
func (_Validatorregistry *ValidatorregistryTransactor) SetStakeRegister(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "setStakeRegister", val)
}

// SetStakeRegister is a paid mutator transaction binding the contract method 0xcbdc17cb.
//
// Solidity: function setStakeRegister(uint256 val) returns()
func (_Validatorregistry *ValidatorregistrySession) SetStakeRegister(val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetStakeRegister(&_Validatorregistry.TransactOpts, val)
}

// SetStakeRegister is a paid mutator transaction binding the contract method 0xcbdc17cb.
//
// Solidity: function setStakeRegister(uint256 val) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) SetStakeRegister(val *big.Int) (*types.Transaction, error) {
	return _Validatorregistry.Contract.SetStakeRegister(&_Validatorregistry.TransactOpts, val)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validatorregistry *ValidatorregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validatorregistry *ValidatorregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validatorregistry.Contract.TransferOwnership(&_Validatorregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validatorregistry.Contract.TransferOwnership(&_Validatorregistry.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Validatorregistry *ValidatorregistryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Validatorregistry *ValidatorregistrySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Validatorregistry.Contract.UpgradeToAndCall(&_Validatorregistry.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Validatorregistry.Contract.UpgradeToAndCall(&_Validatorregistry.TransactOpts, newImplementation, data)
}

// ValidatorComplain is a paid mutator transaction binding the contract method 0x3b69abbf.
//
// Solidity: function validatorComplain(address addr) returns()
func (_Validatorregistry *ValidatorregistryTransactor) ValidatorComplain(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "validatorComplain", addr)
}

// ValidatorComplain is a paid mutator transaction binding the contract method 0x3b69abbf.
//
// Solidity: function validatorComplain(address addr) returns()
func (_Validatorregistry *ValidatorregistrySession) ValidatorComplain(addr common.Address) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorComplain(&_Validatorregistry.TransactOpts, addr)
}

// ValidatorComplain is a paid mutator transaction binding the contract method 0x3b69abbf.
//
// Solidity: function validatorComplain(address addr) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) ValidatorComplain(addr common.Address) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorComplain(&_Validatorregistry.TransactOpts, addr)
}

// ValidatorDeregister is a paid mutator transaction binding the contract method 0x1f99c4d6.
//
// Solidity: function validatorDeregister() returns()
func (_Validatorregistry *ValidatorregistryTransactor) ValidatorDeregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "validatorDeregister")
}

// ValidatorDeregister is a paid mutator transaction binding the contract method 0x1f99c4d6.
//
// Solidity: function validatorDeregister() returns()
func (_Validatorregistry *ValidatorregistrySession) ValidatorDeregister() (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorDeregister(&_Validatorregistry.TransactOpts)
}

// ValidatorDeregister is a paid mutator transaction binding the contract method 0x1f99c4d6.
//
// Solidity: function validatorDeregister() returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) ValidatorDeregister() (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorDeregister(&_Validatorregistry.TransactOpts)
}

// ValidatorRegister is a paid mutator transaction binding the contract method 0xd3a9cfce.
//
// Solidity: function validatorRegister(uint256 stake, bool pubKeyYparity, bytes32 pubKeyX, string host) returns()
func (_Validatorregistry *ValidatorregistryTransactor) ValidatorRegister(opts *bind.TransactOpts, stake *big.Int, pubKeyYparity bool, pubKeyX [32]byte, host string) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "validatorRegister", stake, pubKeyYparity, pubKeyX, host)
}

// ValidatorRegister is a paid mutator transaction binding the contract method 0xd3a9cfce.
//
// Solidity: function validatorRegister(uint256 stake, bool pubKeyYparity, bytes32 pubKeyX, string host) returns()
func (_Validatorregistry *ValidatorregistrySession) ValidatorRegister(stake *big.Int, pubKeyYparity bool, pubKeyX [32]byte, host string) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorRegister(&_Validatorregistry.TransactOpts, stake, pubKeyYparity, pubKeyX, host)
}

// ValidatorRegister is a paid mutator transaction binding the contract method 0xd3a9cfce.
//
// Solidity: function validatorRegister(uint256 stake, bool pubKeyYparity, bytes32 pubKeyX, string host) returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) ValidatorRegister(stake *big.Int, pubKeyYparity bool, pubKeyX [32]byte, host string) (*types.Transaction, error) {
	return _Validatorregistry.Contract.ValidatorRegister(&_Validatorregistry.TransactOpts, stake, pubKeyYparity, pubKeyX, host)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Validatorregistry *ValidatorregistryTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validatorregistry.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Validatorregistry *ValidatorregistrySession) Withdraw() (*types.Transaction, error) {
	return _Validatorregistry.Contract.Withdraw(&_Validatorregistry.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Validatorregistry *ValidatorregistryTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Validatorregistry.Contract.Withdraw(&_Validatorregistry.TransactOpts)
}

// ValidatorregistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Validatorregistry contract.
type ValidatorregistryInitializedIterator struct {
	Event *ValidatorregistryInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryInitialized)
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
		it.Event = new(ValidatorregistryInitialized)
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
func (it *ValidatorregistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryInitialized represents a Initialized event raised by the Validatorregistry contract.
type ValidatorregistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Validatorregistry *ValidatorregistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorregistryInitializedIterator, error) {

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryInitializedIterator{contract: _Validatorregistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Validatorregistry *ValidatorregistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorregistryInitialized) (event.Subscription, error) {

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryInitialized)
				if err := _Validatorregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Validatorregistry *ValidatorregistryFilterer) ParseInitialized(log types.Log) (*ValidatorregistryInitialized, error) {
	event := new(ValidatorregistryInitialized)
	if err := _Validatorregistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Validatorregistry contract.
type ValidatorregistryOwnershipTransferredIterator struct {
	Event *ValidatorregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryOwnershipTransferred)
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
		it.Event = new(ValidatorregistryOwnershipTransferred)
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
func (it *ValidatorregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Validatorregistry contract.
type ValidatorregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validatorregistry *ValidatorregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryOwnershipTransferredIterator{contract: _Validatorregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validatorregistry *ValidatorregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryOwnershipTransferred)
				if err := _Validatorregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Validatorregistry *ValidatorregistryFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorregistryOwnershipTransferred, error) {
	event := new(ValidatorregistryOwnershipTransferred)
	if err := _Validatorregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorregistryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Validatorregistry contract.
type ValidatorregistryUpgradedIterator struct {
	Event *ValidatorregistryUpgraded // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryUpgraded)
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
		it.Event = new(ValidatorregistryUpgraded)
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
func (it *ValidatorregistryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryUpgraded represents a Upgraded event raised by the Validatorregistry contract.
type ValidatorregistryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Validatorregistry *ValidatorregistryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ValidatorregistryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryUpgradedIterator{contract: _Validatorregistry.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Validatorregistry *ValidatorregistryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ValidatorregistryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryUpgraded)
				if err := _Validatorregistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Validatorregistry *ValidatorregistryFilterer) ParseUpgraded(log types.Log) (*ValidatorregistryUpgraded, error) {
	event := new(ValidatorregistryUpgraded)
	if err := _Validatorregistry.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorregistryValidatorComplainIterator is returned from FilterValidatorComplain and is used to iterate over the raw logs and unpacked data for ValidatorComplain events raised by the Validatorregistry contract.
type ValidatorregistryValidatorComplainIterator struct {
	Event *ValidatorregistryValidatorComplain // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryValidatorComplainIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryValidatorComplain)
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
		it.Event = new(ValidatorregistryValidatorComplain)
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
func (it *ValidatorregistryValidatorComplainIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryValidatorComplainIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryValidatorComplain represents a ValidatorComplain event raised by the Validatorregistry contract.
type ValidatorregistryValidatorComplain struct {
	Validator  common.Address
	Complainer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorComplain is a free log retrieval operation binding the contract event 0x7f26ac01ce8cd2b3ef849663458bd3d70d7ac99342f366677ec965fe4d4617e6.
//
// Solidity: event ValidatorComplain(address indexed validator, address indexed complainer)
func (_Validatorregistry *ValidatorregistryFilterer) FilterValidatorComplain(opts *bind.FilterOpts, validator []common.Address, complainer []common.Address) (*ValidatorregistryValidatorComplainIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var complainerRule []interface{}
	for _, complainerItem := range complainer {
		complainerRule = append(complainerRule, complainerItem)
	}

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "ValidatorComplain", validatorRule, complainerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryValidatorComplainIterator{contract: _Validatorregistry.contract, event: "ValidatorComplain", logs: logs, sub: sub}, nil
}

// WatchValidatorComplain is a free log subscription operation binding the contract event 0x7f26ac01ce8cd2b3ef849663458bd3d70d7ac99342f366677ec965fe4d4617e6.
//
// Solidity: event ValidatorComplain(address indexed validator, address indexed complainer)
func (_Validatorregistry *ValidatorregistryFilterer) WatchValidatorComplain(opts *bind.WatchOpts, sink chan<- *ValidatorregistryValidatorComplain, validator []common.Address, complainer []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var complainerRule []interface{}
	for _, complainerItem := range complainer {
		complainerRule = append(complainerRule, complainerItem)
	}

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "ValidatorComplain", validatorRule, complainerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryValidatorComplain)
				if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorComplain", log); err != nil {
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

// ParseValidatorComplain is a log parse operation binding the contract event 0x7f26ac01ce8cd2b3ef849663458bd3d70d7ac99342f366677ec965fe4d4617e6.
//
// Solidity: event ValidatorComplain(address indexed validator, address indexed complainer)
func (_Validatorregistry *ValidatorregistryFilterer) ParseValidatorComplain(log types.Log) (*ValidatorregistryValidatorComplain, error) {
	event := new(ValidatorregistryValidatorComplain)
	if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorComplain", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorregistryValidatorDeregisteredIterator is returned from FilterValidatorDeregistered and is used to iterate over the raw logs and unpacked data for ValidatorDeregistered events raised by the Validatorregistry contract.
type ValidatorregistryValidatorDeregisteredIterator struct {
	Event *ValidatorregistryValidatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryValidatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryValidatorDeregistered)
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
		it.Event = new(ValidatorregistryValidatorDeregistered)
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
func (it *ValidatorregistryValidatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryValidatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryValidatorDeregistered represents a ValidatorDeregistered event raised by the Validatorregistry contract.
type ValidatorregistryValidatorDeregistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeregistered is a free log retrieval operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) FilterValidatorDeregistered(opts *bind.FilterOpts, validator []common.Address) (*ValidatorregistryValidatorDeregisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "ValidatorDeregistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryValidatorDeregisteredIterator{contract: _Validatorregistry.contract, event: "ValidatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchValidatorDeregistered is a free log subscription operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) WatchValidatorDeregistered(opts *bind.WatchOpts, sink chan<- *ValidatorregistryValidatorDeregistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "ValidatorDeregistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryValidatorDeregistered)
				if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
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

// ParseValidatorDeregistered is a log parse operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) ParseValidatorDeregistered(log types.Log) (*ValidatorregistryValidatorDeregistered, error) {
	event := new(ValidatorregistryValidatorDeregistered)
	if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorregistryValidatorPunishedIterator is returned from FilterValidatorPunished and is used to iterate over the raw logs and unpacked data for ValidatorPunished events raised by the Validatorregistry contract.
type ValidatorregistryValidatorPunishedIterator struct {
	Event *ValidatorregistryValidatorPunished // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryValidatorPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryValidatorPunished)
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
		it.Event = new(ValidatorregistryValidatorPunished)
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
func (it *ValidatorregistryValidatorPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryValidatorPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryValidatorPunished represents a ValidatorPunished event raised by the Validatorregistry contract.
type ValidatorregistryValidatorPunished struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorPunished is a free log retrieval operation binding the contract event 0x8d76dace3eea3177ede619176fca83c0b96fb43118cd8c5423ff5c84fcc5d3ac.
//
// Solidity: event ValidatorPunished(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) FilterValidatorPunished(opts *bind.FilterOpts, validator []common.Address) (*ValidatorregistryValidatorPunishedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "ValidatorPunished", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryValidatorPunishedIterator{contract: _Validatorregistry.contract, event: "ValidatorPunished", logs: logs, sub: sub}, nil
}

// WatchValidatorPunished is a free log subscription operation binding the contract event 0x8d76dace3eea3177ede619176fca83c0b96fb43118cd8c5423ff5c84fcc5d3ac.
//
// Solidity: event ValidatorPunished(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) WatchValidatorPunished(opts *bind.WatchOpts, sink chan<- *ValidatorregistryValidatorPunished, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "ValidatorPunished", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryValidatorPunished)
				if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
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

// ParseValidatorPunished is a log parse operation binding the contract event 0x8d76dace3eea3177ede619176fca83c0b96fb43118cd8c5423ff5c84fcc5d3ac.
//
// Solidity: event ValidatorPunished(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) ParseValidatorPunished(log types.Log) (*ValidatorregistryValidatorPunished, error) {
	event := new(ValidatorregistryValidatorPunished)
	if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorregistryValidatorRegisteredUpdatedIterator is returned from FilterValidatorRegisteredUpdated and is used to iterate over the raw logs and unpacked data for ValidatorRegisteredUpdated events raised by the Validatorregistry contract.
type ValidatorregistryValidatorRegisteredUpdatedIterator struct {
	Event *ValidatorregistryValidatorRegisteredUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorregistryValidatorRegisteredUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorregistryValidatorRegisteredUpdated)
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
		it.Event = new(ValidatorregistryValidatorRegisteredUpdated)
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
func (it *ValidatorregistryValidatorRegisteredUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorregistryValidatorRegisteredUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorregistryValidatorRegisteredUpdated represents a ValidatorRegisteredUpdated event raised by the Validatorregistry contract.
type ValidatorregistryValidatorRegisteredUpdated struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegisteredUpdated is a free log retrieval operation binding the contract event 0xfa8eaa9a4538bcd66427267278e2c8671fb3145b6ca6b6ac87cf3bd4ae2e0662.
//
// Solidity: event ValidatorRegisteredUpdated(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) FilterValidatorRegisteredUpdated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorregistryValidatorRegisteredUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorregistry.contract.FilterLogs(opts, "ValidatorRegisteredUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorregistryValidatorRegisteredUpdatedIterator{contract: _Validatorregistry.contract, event: "ValidatorRegisteredUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorRegisteredUpdated is a free log subscription operation binding the contract event 0xfa8eaa9a4538bcd66427267278e2c8671fb3145b6ca6b6ac87cf3bd4ae2e0662.
//
// Solidity: event ValidatorRegisteredUpdated(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) WatchValidatorRegisteredUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorregistryValidatorRegisteredUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validatorregistry.contract.WatchLogs(opts, "ValidatorRegisteredUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorregistryValidatorRegisteredUpdated)
				if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorRegisteredUpdated", log); err != nil {
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

// ParseValidatorRegisteredUpdated is a log parse operation binding the contract event 0xfa8eaa9a4538bcd66427267278e2c8671fb3145b6ca6b6ac87cf3bd4ae2e0662.
//
// Solidity: event ValidatorRegisteredUpdated(address indexed validator)
func (_Validatorregistry *ValidatorregistryFilterer) ParseValidatorRegisteredUpdated(log types.Log) (*ValidatorregistryValidatorRegisteredUpdated, error) {
	event := new(ValidatorregistryValidatorRegisteredUpdated)
	if err := _Validatorregistry.contract.UnpackLog(event, "ValidatorRegisteredUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
