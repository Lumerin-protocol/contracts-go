// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package implementation

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

// ImplementationHistoryEntry is an auto generated low-level Go binding around an user-defined struct.
type ImplementationHistoryEntry struct {
	GoodCloseout bool
	PurchaseTime *big.Int
	EndTime      *big.Int
	Price        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Buyer        common.Address
}

// ImplementationTerms is an auto generated low-level Go binding around an user-defined struct.
type ImplementationTerms struct {
	Price        *big.Int
	Limit        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Version      uint32
	ProfitTarget int8
}

// ImplementationMetaData contains all meta data concerning the Implementation contract.
var ImplementationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newCipherText\",\"type\":\"string\"}],\"name\":\"cipherTextUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumImplementation.CloseReason\",\"name\":\"_reason\",\"type\":\"uint8\"}],\"name\":\"closedEarly\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValidatorURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDestURL\",\"type\":\"string\"}],\"name\":\"destinationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"fundsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"}],\"name\":\"fundsClaimedValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"purchaseInfoUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATOR_FEE_MULT\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFundsValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cloneFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumImplementation.CloseReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"closeEarly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"encrDestURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"encrValidatorURL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureTerms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"_goodCloseout\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_purchaseTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"internalType\":\"structImplementation.HistoryEntry[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicVariables\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"_state\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingBlockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_encryptedPoolData\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_hasFutureTerms\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicVariablesV2\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"_state\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"internalType\":\"structImplementation.Terms\",\"name\":\"_terms\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_startingBlockTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_encryptedPoolData\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_hasFutureTerms\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStats\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_successCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_failCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"history\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_goodCloseout\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_purchaseTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lmrAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_cloneFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pubKey\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLastValidatorNotPaid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lumerin\",\"outputs\":[{\"internalType\":\"contractLumerin\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"closeOutType\",\"type\":\"uint256\"}],\"name\":\"setContractCloseOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"}],\"name\":\"setContractDeleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"}],\"name\":\"setDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"_validatorFeeRateScaled\",\"type\":\"uint16\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newEncryptedPoolData\",\"type\":\"string\"}],\"name\":\"setUpdateMiningInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"name\":\"setUpdatePurchaseInformation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startingBlockTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"terms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorFeeRateScaled\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ImplementationABI is the input ABI used to generate the binding from.
// Deprecated: Use ImplementationMetaData.ABI instead.
var ImplementationABI = ImplementationMetaData.ABI

// Implementation is an auto generated Go binding around an Ethereum contract.
type Implementation struct {
	ImplementationCaller     // Read-only binding to the contract
	ImplementationTransactor // Write-only binding to the contract
	ImplementationFilterer   // Log filterer for contract events
}

// ImplementationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ImplementationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ImplementationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ImplementationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ImplementationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ImplementationSession struct {
	Contract     *Implementation   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ImplementationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ImplementationCallerSession struct {
	Contract *ImplementationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ImplementationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ImplementationTransactorSession struct {
	Contract     *ImplementationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ImplementationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ImplementationRaw struct {
	Contract *Implementation // Generic contract binding to access the raw methods on
}

// ImplementationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ImplementationCallerRaw struct {
	Contract *ImplementationCaller // Generic read-only contract binding to access the raw methods on
}

// ImplementationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ImplementationTransactorRaw struct {
	Contract *ImplementationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewImplementation creates a new instance of Implementation, bound to a specific deployed contract.
func NewImplementation(address common.Address, backend bind.ContractBackend) (*Implementation, error) {
	contract, err := bindImplementation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Implementation{ImplementationCaller: ImplementationCaller{contract: contract}, ImplementationTransactor: ImplementationTransactor{contract: contract}, ImplementationFilterer: ImplementationFilterer{contract: contract}}, nil
}

// NewImplementationCaller creates a new read-only instance of Implementation, bound to a specific deployed contract.
func NewImplementationCaller(address common.Address, caller bind.ContractCaller) (*ImplementationCaller, error) {
	contract, err := bindImplementation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ImplementationCaller{contract: contract}, nil
}

// NewImplementationTransactor creates a new write-only instance of Implementation, bound to a specific deployed contract.
func NewImplementationTransactor(address common.Address, transactor bind.ContractTransactor) (*ImplementationTransactor, error) {
	contract, err := bindImplementation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ImplementationTransactor{contract: contract}, nil
}

// NewImplementationFilterer creates a new log filterer instance of Implementation, bound to a specific deployed contract.
func NewImplementationFilterer(address common.Address, filterer bind.ContractFilterer) (*ImplementationFilterer, error) {
	contract, err := bindImplementation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ImplementationFilterer{contract: contract}, nil
}

// bindImplementation binds a generic wrapper to an already deployed contract.
func bindImplementation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ImplementationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Implementation *ImplementationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Implementation.Contract.ImplementationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Implementation *ImplementationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.Contract.ImplementationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Implementation *ImplementationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Implementation.Contract.ImplementationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Implementation *ImplementationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Implementation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Implementation *ImplementationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Implementation *ImplementationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Implementation.Contract.contract.Transact(opts, method, params...)
}

// VALIDATORFEEMULT is a free data retrieval call binding the contract method 0x444b5c73.
//
// Solidity: function VALIDATOR_FEE_MULT() view returns(uint16)
func (_Implementation *ImplementationCaller) VALIDATORFEEMULT(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "VALIDATOR_FEE_MULT")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// VALIDATORFEEMULT is a free data retrieval call binding the contract method 0x444b5c73.
//
// Solidity: function VALIDATOR_FEE_MULT() view returns(uint16)
func (_Implementation *ImplementationSession) VALIDATORFEEMULT() (uint16, error) {
	return _Implementation.Contract.VALIDATORFEEMULT(&_Implementation.CallOpts)
}

// VALIDATORFEEMULT is a free data retrieval call binding the contract method 0x444b5c73.
//
// Solidity: function VALIDATOR_FEE_MULT() view returns(uint16)
func (_Implementation *ImplementationCallerSession) VALIDATORFEEMULT() (uint16, error) {
	return _Implementation.Contract.VALIDATORFEEMULT(&_Implementation.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationCaller) Buyer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "buyer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationSession) Buyer() (common.Address, error) {
	return _Implementation.Contract.Buyer(&_Implementation.CallOpts)
}

// Buyer is a free data retrieval call binding the contract method 0x7150d8ae.
//
// Solidity: function buyer() view returns(address)
func (_Implementation *ImplementationCallerSession) Buyer() (common.Address, error) {
	return _Implementation.Contract.Buyer(&_Implementation.CallOpts)
}

// CloneFactory is a free data retrieval call binding the contract method 0xa064f9a1.
//
// Solidity: function cloneFactory() view returns(address)
func (_Implementation *ImplementationCaller) CloneFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "cloneFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CloneFactory is a free data retrieval call binding the contract method 0xa064f9a1.
//
// Solidity: function cloneFactory() view returns(address)
func (_Implementation *ImplementationSession) CloneFactory() (common.Address, error) {
	return _Implementation.Contract.CloneFactory(&_Implementation.CallOpts)
}

// CloneFactory is a free data retrieval call binding the contract method 0xa064f9a1.
//
// Solidity: function cloneFactory() view returns(address)
func (_Implementation *ImplementationCallerSession) CloneFactory() (common.Address, error) {
	return _Implementation.Contract.CloneFactory(&_Implementation.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationCaller) ContractState(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "contractState")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationSession) ContractState() (uint8, error) {
	return _Implementation.Contract.ContractState(&_Implementation.CallOpts)
}

// ContractState is a free data retrieval call binding the contract method 0x85209ee0.
//
// Solidity: function contractState() view returns(uint8)
func (_Implementation *ImplementationCallerSession) ContractState() (uint8, error) {
	return _Implementation.Contract.ContractState(&_Implementation.CallOpts)
}

// EncrDestURL is a free data retrieval call binding the contract method 0x628eda24.
//
// Solidity: function encrDestURL() view returns(string)
func (_Implementation *ImplementationCaller) EncrDestURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "encrDestURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncrDestURL is a free data retrieval call binding the contract method 0x628eda24.
//
// Solidity: function encrDestURL() view returns(string)
func (_Implementation *ImplementationSession) EncrDestURL() (string, error) {
	return _Implementation.Contract.EncrDestURL(&_Implementation.CallOpts)
}

// EncrDestURL is a free data retrieval call binding the contract method 0x628eda24.
//
// Solidity: function encrDestURL() view returns(string)
func (_Implementation *ImplementationCallerSession) EncrDestURL() (string, error) {
	return _Implementation.Contract.EncrDestURL(&_Implementation.CallOpts)
}

// EncrValidatorURL is a free data retrieval call binding the contract method 0x30bad9e1.
//
// Solidity: function encrValidatorURL() view returns(string)
func (_Implementation *ImplementationCaller) EncrValidatorURL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "encrValidatorURL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EncrValidatorURL is a free data retrieval call binding the contract method 0x30bad9e1.
//
// Solidity: function encrValidatorURL() view returns(string)
func (_Implementation *ImplementationSession) EncrValidatorURL() (string, error) {
	return _Implementation.Contract.EncrValidatorURL(&_Implementation.CallOpts)
}

// EncrValidatorURL is a free data retrieval call binding the contract method 0x30bad9e1.
//
// Solidity: function encrValidatorURL() view returns(string)
func (_Implementation *ImplementationCallerSession) EncrValidatorURL() (string, error) {
	return _Implementation.Contract.EncrValidatorURL(&_Implementation.CallOpts)
}

// FutureTerms is a free data retrieval call binding the contract method 0xa4e27566.
//
// Solidity: function futureTerms() view returns(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint32 _version, int8 _profitTarget)
func (_Implementation *ImplementationCaller) FutureTerms(opts *bind.CallOpts) (struct {
	Price        *big.Int
	Limit        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Version      uint32
	ProfitTarget int8
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "futureTerms")

	outstruct := new(struct {
		Price        *big.Int
		Limit        *big.Int
		Speed        *big.Int
		Length       *big.Int
		Version      uint32
		ProfitTarget int8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Limit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Speed = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ProfitTarget = *abi.ConvertType(out[5], new(int8)).(*int8)

	return *outstruct, err

}

// FutureTerms is a free data retrieval call binding the contract method 0xa4e27566.
//
// Solidity: function futureTerms() view returns(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint32 _version, int8 _profitTarget)
func (_Implementation *ImplementationSession) FutureTerms() (struct {
	Price        *big.Int
	Limit        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Version      uint32
	ProfitTarget int8
}, error) {
	return _Implementation.Contract.FutureTerms(&_Implementation.CallOpts)
}

// FutureTerms is a free data retrieval call binding the contract method 0xa4e27566.
//
// Solidity: function futureTerms() view returns(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint32 _version, int8 _profitTarget)
func (_Implementation *ImplementationCallerSession) FutureTerms() (struct {
	Price        *big.Int
	Limit        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Version      uint32
	ProfitTarget int8
}, error) {
	return _Implementation.Contract.FutureTerms(&_Implementation.CallOpts)
}

// GetHistory is a free data retrieval call binding the contract method 0x6906679b.
//
// Solidity: function getHistory(uint256 _offset, uint256 _limit) view returns((bool,uint256,uint256,uint256,uint256,uint256,address)[])
func (_Implementation *ImplementationCaller) GetHistory(opts *bind.CallOpts, _offset *big.Int, _limit *big.Int) ([]ImplementationHistoryEntry, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "getHistory", _offset, _limit)

	if err != nil {
		return *new([]ImplementationHistoryEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]ImplementationHistoryEntry)).(*[]ImplementationHistoryEntry)

	return out0, err

}

// GetHistory is a free data retrieval call binding the contract method 0x6906679b.
//
// Solidity: function getHistory(uint256 _offset, uint256 _limit) view returns((bool,uint256,uint256,uint256,uint256,uint256,address)[])
func (_Implementation *ImplementationSession) GetHistory(_offset *big.Int, _limit *big.Int) ([]ImplementationHistoryEntry, error) {
	return _Implementation.Contract.GetHistory(&_Implementation.CallOpts, _offset, _limit)
}

// GetHistory is a free data retrieval call binding the contract method 0x6906679b.
//
// Solidity: function getHistory(uint256 _offset, uint256 _limit) view returns((bool,uint256,uint256,uint256,uint256,uint256,address)[])
func (_Implementation *ImplementationCallerSession) GetHistory(_offset *big.Int, _limit *big.Int) ([]ImplementationHistoryEntry, error) {
	return _Implementation.Contract.GetHistory(&_Implementation.CallOpts, _offset, _limit)
}

// GetPublicVariables is a free data retrieval call binding the contract method 0xca3225fa.
//
// Solidity: function getPublicVariables() view returns(uint8 _state, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _startingBlockTimestamp, address _buyer, address _seller, string _encryptedPoolData, bool _isDeleted, uint256 _balance, bool _hasFutureTerms, uint32 _version)
func (_Implementation *ImplementationCaller) GetPublicVariables(opts *bind.CallOpts) (struct {
	State                  uint8
	Price                  *big.Int
	Limit                  *big.Int
	Speed                  *big.Int
	Length                 *big.Int
	StartingBlockTimestamp *big.Int
	Buyer                  common.Address
	Seller                 common.Address
	EncryptedPoolData      string
	IsDeleted              bool
	Balance                *big.Int
	HasFutureTerms         bool
	Version                uint32
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "getPublicVariables")

	outstruct := new(struct {
		State                  uint8
		Price                  *big.Int
		Limit                  *big.Int
		Speed                  *big.Int
		Length                 *big.Int
		StartingBlockTimestamp *big.Int
		Buyer                  common.Address
		Seller                 common.Address
		EncryptedPoolData      string
		IsDeleted              bool
		Balance                *big.Int
		HasFutureTerms         bool
		Version                uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Limit = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Speed = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StartingBlockTimestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Seller = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.EncryptedPoolData = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.IsDeleted = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.Balance = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.HasFutureTerms = *abi.ConvertType(out[11], new(bool)).(*bool)
	outstruct.Version = *abi.ConvertType(out[12], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetPublicVariables is a free data retrieval call binding the contract method 0xca3225fa.
//
// Solidity: function getPublicVariables() view returns(uint8 _state, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _startingBlockTimestamp, address _buyer, address _seller, string _encryptedPoolData, bool _isDeleted, uint256 _balance, bool _hasFutureTerms, uint32 _version)
func (_Implementation *ImplementationSession) GetPublicVariables() (struct {
	State                  uint8
	Price                  *big.Int
	Limit                  *big.Int
	Speed                  *big.Int
	Length                 *big.Int
	StartingBlockTimestamp *big.Int
	Buyer                  common.Address
	Seller                 common.Address
	EncryptedPoolData      string
	IsDeleted              bool
	Balance                *big.Int
	HasFutureTerms         bool
	Version                uint32
}, error) {
	return _Implementation.Contract.GetPublicVariables(&_Implementation.CallOpts)
}

// GetPublicVariables is a free data retrieval call binding the contract method 0xca3225fa.
//
// Solidity: function getPublicVariables() view returns(uint8 _state, uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint256 _startingBlockTimestamp, address _buyer, address _seller, string _encryptedPoolData, bool _isDeleted, uint256 _balance, bool _hasFutureTerms, uint32 _version)
func (_Implementation *ImplementationCallerSession) GetPublicVariables() (struct {
	State                  uint8
	Price                  *big.Int
	Limit                  *big.Int
	Speed                  *big.Int
	Length                 *big.Int
	StartingBlockTimestamp *big.Int
	Buyer                  common.Address
	Seller                 common.Address
	EncryptedPoolData      string
	IsDeleted              bool
	Balance                *big.Int
	HasFutureTerms         bool
	Version                uint32
}, error) {
	return _Implementation.Contract.GetPublicVariables(&_Implementation.CallOpts)
}

// GetPublicVariablesV2 is a free data retrieval call binding the contract method 0x896187fb.
//
// Solidity: function getPublicVariablesV2() view returns(uint8 _state, (uint256,uint256,uint256,uint256,uint32,int8) _terms, uint256 _startingBlockTimestamp, address _buyer, address _seller, string _encryptedPoolData, bool _isDeleted, uint256 _balance, bool _hasFutureTerms)
func (_Implementation *ImplementationCaller) GetPublicVariablesV2(opts *bind.CallOpts) (struct {
	State                  uint8
	Terms                  ImplementationTerms
	StartingBlockTimestamp *big.Int
	Buyer                  common.Address
	Seller                 common.Address
	EncryptedPoolData      string
	IsDeleted              bool
	Balance                *big.Int
	HasFutureTerms         bool
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "getPublicVariablesV2")

	outstruct := new(struct {
		State                  uint8
		Terms                  ImplementationTerms
		StartingBlockTimestamp *big.Int
		Buyer                  common.Address
		Seller                 common.Address
		EncryptedPoolData      string
		IsDeleted              bool
		Balance                *big.Int
		HasFutureTerms         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Terms = *abi.ConvertType(out[1], new(ImplementationTerms)).(*ImplementationTerms)
	outstruct.StartingBlockTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Seller = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.EncryptedPoolData = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.IsDeleted = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.Balance = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.HasFutureTerms = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetPublicVariablesV2 is a free data retrieval call binding the contract method 0x896187fb.
//
// Solidity: function getPublicVariablesV2() view returns(uint8 _state, (uint256,uint256,uint256,uint256,uint32,int8) _terms, uint256 _startingBlockTimestamp, address _buyer, address _seller, string _encryptedPoolData, bool _isDeleted, uint256 _balance, bool _hasFutureTerms)
func (_Implementation *ImplementationSession) GetPublicVariablesV2() (struct {
	State                  uint8
	Terms                  ImplementationTerms
	StartingBlockTimestamp *big.Int
	Buyer                  common.Address
	Seller                 common.Address
	EncryptedPoolData      string
	IsDeleted              bool
	Balance                *big.Int
	HasFutureTerms         bool
}, error) {
	return _Implementation.Contract.GetPublicVariablesV2(&_Implementation.CallOpts)
}

// GetPublicVariablesV2 is a free data retrieval call binding the contract method 0x896187fb.
//
// Solidity: function getPublicVariablesV2() view returns(uint8 _state, (uint256,uint256,uint256,uint256,uint32,int8) _terms, uint256 _startingBlockTimestamp, address _buyer, address _seller, string _encryptedPoolData, bool _isDeleted, uint256 _balance, bool _hasFutureTerms)
func (_Implementation *ImplementationCallerSession) GetPublicVariablesV2() (struct {
	State                  uint8
	Terms                  ImplementationTerms
	StartingBlockTimestamp *big.Int
	Buyer                  common.Address
	Seller                 common.Address
	EncryptedPoolData      string
	IsDeleted              bool
	Balance                *big.Int
	HasFutureTerms         bool
}, error) {
	return _Implementation.Contract.GetPublicVariablesV2(&_Implementation.CallOpts)
}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns(uint256 _successCount, uint256 _failCount)
func (_Implementation *ImplementationCaller) GetStats(opts *bind.CallOpts) (struct {
	SuccessCount *big.Int
	FailCount    *big.Int
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "getStats")

	outstruct := new(struct {
		SuccessCount *big.Int
		FailCount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SuccessCount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FailCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns(uint256 _successCount, uint256 _failCount)
func (_Implementation *ImplementationSession) GetStats() (struct {
	SuccessCount *big.Int
	FailCount    *big.Int
}, error) {
	return _Implementation.Contract.GetStats(&_Implementation.CallOpts)
}

// GetStats is a free data retrieval call binding the contract method 0xc59d4847.
//
// Solidity: function getStats() view returns(uint256 _successCount, uint256 _failCount)
func (_Implementation *ImplementationCallerSession) GetStats() (struct {
	SuccessCount *big.Int
	FailCount    *big.Int
}, error) {
	return _Implementation.Contract.GetStats(&_Implementation.CallOpts)
}

// History is a free data retrieval call binding the contract method 0xa7a38f0b.
//
// Solidity: function history(uint256 ) view returns(bool _goodCloseout, uint256 _purchaseTime, uint256 _endTime, uint256 _price, uint256 _speed, uint256 _length, address _buyer)
func (_Implementation *ImplementationCaller) History(opts *bind.CallOpts, arg0 *big.Int) (struct {
	GoodCloseout bool
	PurchaseTime *big.Int
	EndTime      *big.Int
	Price        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Buyer        common.Address
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "history", arg0)

	outstruct := new(struct {
		GoodCloseout bool
		PurchaseTime *big.Int
		EndTime      *big.Int
		Price        *big.Int
		Speed        *big.Int
		Length       *big.Int
		Buyer        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GoodCloseout = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PurchaseTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Speed = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// History is a free data retrieval call binding the contract method 0xa7a38f0b.
//
// Solidity: function history(uint256 ) view returns(bool _goodCloseout, uint256 _purchaseTime, uint256 _endTime, uint256 _price, uint256 _speed, uint256 _length, address _buyer)
func (_Implementation *ImplementationSession) History(arg0 *big.Int) (struct {
	GoodCloseout bool
	PurchaseTime *big.Int
	EndTime      *big.Int
	Price        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Buyer        common.Address
}, error) {
	return _Implementation.Contract.History(&_Implementation.CallOpts, arg0)
}

// History is a free data retrieval call binding the contract method 0xa7a38f0b.
//
// Solidity: function history(uint256 ) view returns(bool _goodCloseout, uint256 _purchaseTime, uint256 _endTime, uint256 _price, uint256 _speed, uint256 _length, address _buyer)
func (_Implementation *ImplementationCallerSession) History(arg0 *big.Int) (struct {
	GoodCloseout bool
	PurchaseTime *big.Int
	EndTime      *big.Int
	Price        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Buyer        common.Address
}, error) {
	return _Implementation.Contract.History(&_Implementation.CallOpts, arg0)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_Implementation *ImplementationCaller) IsDeleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "isDeleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_Implementation *ImplementationSession) IsDeleted() (bool, error) {
	return _Implementation.Contract.IsDeleted(&_Implementation.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_Implementation *ImplementationCallerSession) IsDeleted() (bool, error) {
	return _Implementation.Contract.IsDeleted(&_Implementation.CallOpts)
}

// IsLastValidatorNotPaid is a free data retrieval call binding the contract method 0x3fd2fd65.
//
// Solidity: function isLastValidatorNotPaid() view returns(bool)
func (_Implementation *ImplementationCaller) IsLastValidatorNotPaid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "isLastValidatorNotPaid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLastValidatorNotPaid is a free data retrieval call binding the contract method 0x3fd2fd65.
//
// Solidity: function isLastValidatorNotPaid() view returns(bool)
func (_Implementation *ImplementationSession) IsLastValidatorNotPaid() (bool, error) {
	return _Implementation.Contract.IsLastValidatorNotPaid(&_Implementation.CallOpts)
}

// IsLastValidatorNotPaid is a free data retrieval call binding the contract method 0x3fd2fd65.
//
// Solidity: function isLastValidatorNotPaid() view returns(bool)
func (_Implementation *ImplementationCallerSession) IsLastValidatorNotPaid() (bool, error) {
	return _Implementation.Contract.IsLastValidatorNotPaid(&_Implementation.CallOpts)
}

// Lumerin is a free data retrieval call binding the contract method 0xe6108971.
//
// Solidity: function lumerin() view returns(address)
func (_Implementation *ImplementationCaller) Lumerin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "lumerin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lumerin is a free data retrieval call binding the contract method 0xe6108971.
//
// Solidity: function lumerin() view returns(address)
func (_Implementation *ImplementationSession) Lumerin() (common.Address, error) {
	return _Implementation.Contract.Lumerin(&_Implementation.CallOpts)
}

// Lumerin is a free data retrieval call binding the contract method 0xe6108971.
//
// Solidity: function lumerin() view returns(address)
func (_Implementation *ImplementationCallerSession) Lumerin() (common.Address, error) {
	return _Implementation.Contract.Lumerin(&_Implementation.CallOpts)
}

// PubKey is a free data retrieval call binding the contract method 0xac2a5dfd.
//
// Solidity: function pubKey() view returns(string)
func (_Implementation *ImplementationCaller) PubKey(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "pubKey")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PubKey is a free data retrieval call binding the contract method 0xac2a5dfd.
//
// Solidity: function pubKey() view returns(string)
func (_Implementation *ImplementationSession) PubKey() (string, error) {
	return _Implementation.Contract.PubKey(&_Implementation.CallOpts)
}

// PubKey is a free data retrieval call binding the contract method 0xac2a5dfd.
//
// Solidity: function pubKey() view returns(string)
func (_Implementation *ImplementationCallerSession) PubKey() (string, error) {
	return _Implementation.Contract.PubKey(&_Implementation.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationCaller) Seller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "seller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationSession) Seller() (common.Address, error) {
	return _Implementation.Contract.Seller(&_Implementation.CallOpts)
}

// Seller is a free data retrieval call binding the contract method 0x08551a53.
//
// Solidity: function seller() view returns(address)
func (_Implementation *ImplementationCallerSession) Seller() (common.Address, error) {
	return _Implementation.Contract.Seller(&_Implementation.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationCaller) StartingBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "startingBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationSession) StartingBlockTimestamp() (*big.Int, error) {
	return _Implementation.Contract.StartingBlockTimestamp(&_Implementation.CallOpts)
}

// StartingBlockTimestamp is a free data retrieval call binding the contract method 0xc5095d68.
//
// Solidity: function startingBlockTimestamp() view returns(uint256)
func (_Implementation *ImplementationCallerSession) StartingBlockTimestamp() (*big.Int, error) {
	return _Implementation.Contract.StartingBlockTimestamp(&_Implementation.CallOpts)
}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint32 _version, int8 _profitTarget)
func (_Implementation *ImplementationCaller) Terms(opts *bind.CallOpts) (struct {
	Price        *big.Int
	Limit        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Version      uint32
	ProfitTarget int8
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "terms")

	outstruct := new(struct {
		Price        *big.Int
		Limit        *big.Int
		Speed        *big.Int
		Length       *big.Int
		Version      uint32
		ProfitTarget int8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Limit = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Speed = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.ProfitTarget = *abi.ConvertType(out[5], new(int8)).(*int8)

	return *outstruct, err

}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint32 _version, int8 _profitTarget)
func (_Implementation *ImplementationSession) Terms() (struct {
	Price        *big.Int
	Limit        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Version      uint32
	ProfitTarget int8
}, error) {
	return _Implementation.Contract.Terms(&_Implementation.CallOpts)
}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, uint32 _version, int8 _profitTarget)
func (_Implementation *ImplementationCallerSession) Terms() (struct {
	Price        *big.Int
	Limit        *big.Int
	Speed        *big.Int
	Length       *big.Int
	Version      uint32
	ProfitTarget int8
}, error) {
	return _Implementation.Contract.Terms(&_Implementation.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Implementation *ImplementationCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Implementation *ImplementationSession) Validator() (common.Address, error) {
	return _Implementation.Contract.Validator(&_Implementation.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Implementation *ImplementationCallerSession) Validator() (common.Address, error) {
	return _Implementation.Contract.Validator(&_Implementation.CallOpts)
}

// ValidatorFeeRateScaled is a free data retrieval call binding the contract method 0x9e3a9bb8.
//
// Solidity: function validatorFeeRateScaled() view returns(uint16)
func (_Implementation *ImplementationCaller) ValidatorFeeRateScaled(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "validatorFeeRateScaled")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ValidatorFeeRateScaled is a free data retrieval call binding the contract method 0x9e3a9bb8.
//
// Solidity: function validatorFeeRateScaled() view returns(uint16)
func (_Implementation *ImplementationSession) ValidatorFeeRateScaled() (uint16, error) {
	return _Implementation.Contract.ValidatorFeeRateScaled(&_Implementation.CallOpts)
}

// ValidatorFeeRateScaled is a free data retrieval call binding the contract method 0x9e3a9bb8.
//
// Solidity: function validatorFeeRateScaled() view returns(uint16)
func (_Implementation *ImplementationCallerSession) ValidatorFeeRateScaled() (uint16, error) {
	return _Implementation.Contract.ValidatorFeeRateScaled(&_Implementation.CallOpts)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0xac307773.
//
// Solidity: function claimFunds() payable returns()
func (_Implementation *ImplementationTransactor) ClaimFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "claimFunds")
}

// ClaimFunds is a paid mutator transaction binding the contract method 0xac307773.
//
// Solidity: function claimFunds() payable returns()
func (_Implementation *ImplementationSession) ClaimFunds() (*types.Transaction, error) {
	return _Implementation.Contract.ClaimFunds(&_Implementation.TransactOpts)
}

// ClaimFunds is a paid mutator transaction binding the contract method 0xac307773.
//
// Solidity: function claimFunds() payable returns()
func (_Implementation *ImplementationTransactorSession) ClaimFunds() (*types.Transaction, error) {
	return _Implementation.Contract.ClaimFunds(&_Implementation.TransactOpts)
}

// ClaimFundsValidator is a paid mutator transaction binding the contract method 0xc5a62342.
//
// Solidity: function claimFundsValidator() returns()
func (_Implementation *ImplementationTransactor) ClaimFundsValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "claimFundsValidator")
}

// ClaimFundsValidator is a paid mutator transaction binding the contract method 0xc5a62342.
//
// Solidity: function claimFundsValidator() returns()
func (_Implementation *ImplementationSession) ClaimFundsValidator() (*types.Transaction, error) {
	return _Implementation.Contract.ClaimFundsValidator(&_Implementation.TransactOpts)
}

// ClaimFundsValidator is a paid mutator transaction binding the contract method 0xc5a62342.
//
// Solidity: function claimFundsValidator() returns()
func (_Implementation *ImplementationTransactorSession) ClaimFundsValidator() (*types.Transaction, error) {
	return _Implementation.Contract.ClaimFundsValidator(&_Implementation.TransactOpts)
}

// CloseEarly is a paid mutator transaction binding the contract method 0x6a5caece.
//
// Solidity: function closeEarly(uint8 reason) returns()
func (_Implementation *ImplementationTransactor) CloseEarly(opts *bind.TransactOpts, reason uint8) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "closeEarly", reason)
}

// CloseEarly is a paid mutator transaction binding the contract method 0x6a5caece.
//
// Solidity: function closeEarly(uint8 reason) returns()
func (_Implementation *ImplementationSession) CloseEarly(reason uint8) (*types.Transaction, error) {
	return _Implementation.Contract.CloseEarly(&_Implementation.TransactOpts, reason)
}

// CloseEarly is a paid mutator transaction binding the contract method 0x6a5caece.
//
// Solidity: function closeEarly(uint8 reason) returns()
func (_Implementation *ImplementationTransactorSession) CloseEarly(reason uint8) (*types.Transaction, error) {
	return _Implementation.Contract.CloseEarly(&_Implementation.TransactOpts, reason)
}

// Initialize is a paid mutator transaction binding the contract method 0x2ffe5545.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget, address _seller, address _lmrAddress, address _cloneFactory, address _validator, string _pubKey) returns()
func (_Implementation *ImplementationTransactor) Initialize(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, _seller common.Address, _lmrAddress common.Address, _cloneFactory common.Address, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "initialize", _price, _limit, _speed, _length, _profitTarget, _seller, _lmrAddress, _cloneFactory, _validator, _pubKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x2ffe5545.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget, address _seller, address _lmrAddress, address _cloneFactory, address _validator, string _pubKey) returns()
func (_Implementation *ImplementationSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, _seller common.Address, _lmrAddress common.Address, _cloneFactory common.Address, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _profitTarget, _seller, _lmrAddress, _cloneFactory, _validator, _pubKey)
}

// Initialize is a paid mutator transaction binding the contract method 0x2ffe5545.
//
// Solidity: function initialize(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget, address _seller, address _lmrAddress, address _cloneFactory, address _validator, string _pubKey) returns()
func (_Implementation *ImplementationTransactorSession) Initialize(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8, _seller common.Address, _lmrAddress common.Address, _cloneFactory common.Address, _validator common.Address, _pubKey string) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _profitTarget, _seller, _lmrAddress, _cloneFactory, _validator, _pubKey)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) payable returns()
func (_Implementation *ImplementationTransactor) SetContractCloseOut(opts *bind.TransactOpts, closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setContractCloseOut", closeOutType)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) payable returns()
func (_Implementation *ImplementationSession) SetContractCloseOut(closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts, closeOutType)
}

// SetContractCloseOut is a paid mutator transaction binding the contract method 0x8e2e6d5d.
//
// Solidity: function setContractCloseOut(uint256 closeOutType) payable returns()
func (_Implementation *ImplementationTransactorSession) SetContractCloseOut(closeOutType *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractCloseOut(&_Implementation.TransactOpts, closeOutType)
}

// SetContractDeleted is a paid mutator transaction binding the contract method 0x567d91f0.
//
// Solidity: function setContractDeleted(bool _isDeleted) returns()
func (_Implementation *ImplementationTransactor) SetContractDeleted(opts *bind.TransactOpts, _isDeleted bool) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setContractDeleted", _isDeleted)
}

// SetContractDeleted is a paid mutator transaction binding the contract method 0x567d91f0.
//
// Solidity: function setContractDeleted(bool _isDeleted) returns()
func (_Implementation *ImplementationSession) SetContractDeleted(_isDeleted bool) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractDeleted(&_Implementation.TransactOpts, _isDeleted)
}

// SetContractDeleted is a paid mutator transaction binding the contract method 0x567d91f0.
//
// Solidity: function setContractDeleted(bool _isDeleted) returns()
func (_Implementation *ImplementationTransactorSession) SetContractDeleted(_isDeleted bool) (*types.Transaction, error) {
	return _Implementation.Contract.SetContractDeleted(&_Implementation.TransactOpts, _isDeleted)
}

// SetDestination is a paid mutator transaction binding the contract method 0xc6a10091.
//
// Solidity: function setDestination(string _encrValidatorURL, string _encrDestURL) returns()
func (_Implementation *ImplementationTransactor) SetDestination(opts *bind.TransactOpts, _encrValidatorURL string, _encrDestURL string) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setDestination", _encrValidatorURL, _encrDestURL)
}

// SetDestination is a paid mutator transaction binding the contract method 0xc6a10091.
//
// Solidity: function setDestination(string _encrValidatorURL, string _encrDestURL) returns()
func (_Implementation *ImplementationSession) SetDestination(_encrValidatorURL string, _encrDestURL string) (*types.Transaction, error) {
	return _Implementation.Contract.SetDestination(&_Implementation.TransactOpts, _encrValidatorURL, _encrDestURL)
}

// SetDestination is a paid mutator transaction binding the contract method 0xc6a10091.
//
// Solidity: function setDestination(string _encrValidatorURL, string _encrDestURL) returns()
func (_Implementation *ImplementationTransactorSession) SetDestination(_encrValidatorURL string, _encrDestURL string) (*types.Transaction, error) {
	return _Implementation.Contract.SetDestination(&_Implementation.TransactOpts, _encrValidatorURL, _encrDestURL)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x829a805a.
//
// Solidity: function setPurchaseContract(string _encrValidatorURL, string _encrDestURL, address _buyer, address _validator, uint16 _validatorFeeRateScaled) returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _encrValidatorURL string, _encrDestURL string, _buyer common.Address, _validator common.Address, _validatorFeeRateScaled uint16) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _encrValidatorURL, _encrDestURL, _buyer, _validator, _validatorFeeRateScaled)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x829a805a.
//
// Solidity: function setPurchaseContract(string _encrValidatorURL, string _encrDestURL, address _buyer, address _validator, uint16 _validatorFeeRateScaled) returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_encrValidatorURL string, _encrDestURL string, _buyer common.Address, _validator common.Address, _validatorFeeRateScaled uint16) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encrValidatorURL, _encrDestURL, _buyer, _validator, _validatorFeeRateScaled)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x829a805a.
//
// Solidity: function setPurchaseContract(string _encrValidatorURL, string _encrDestURL, address _buyer, address _validator, uint16 _validatorFeeRateScaled) returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_encrValidatorURL string, _encrDestURL string, _buyer common.Address, _validator common.Address, _validatorFeeRateScaled uint16) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encrValidatorURL, _encrDestURL, _buyer, _validator, _validatorFeeRateScaled)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationTransactor) SetUpdateMiningInformation(opts *bind.TransactOpts, _newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setUpdateMiningInformation", _newEncryptedPoolData)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationSession) SetUpdateMiningInformation(_newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdateMiningInformation(&_Implementation.TransactOpts, _newEncryptedPoolData)
}

// SetUpdateMiningInformation is a paid mutator transaction binding the contract method 0x719e6b5b.
//
// Solidity: function setUpdateMiningInformation(string _newEncryptedPoolData) returns()
func (_Implementation *ImplementationTransactorSession) SetUpdateMiningInformation(_newEncryptedPoolData string) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdateMiningInformation(&_Implementation.TransactOpts, _newEncryptedPoolData)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xa811af1e.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Implementation *ImplementationTransactor) SetUpdatePurchaseInformation(opts *bind.TransactOpts, _price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setUpdatePurchaseInformation", _price, _limit, _speed, _length, _profitTarget)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xa811af1e.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Implementation *ImplementationSession) SetUpdatePurchaseInformation(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdatePurchaseInformation(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _profitTarget)
}

// SetUpdatePurchaseInformation is a paid mutator transaction binding the contract method 0xa811af1e.
//
// Solidity: function setUpdatePurchaseInformation(uint256 _price, uint256 _limit, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Implementation *ImplementationTransactorSession) SetUpdatePurchaseInformation(_price *big.Int, _limit *big.Int, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Implementation.Contract.SetUpdatePurchaseInformation(&_Implementation.TransactOpts, _price, _limit, _speed, _length, _profitTarget)
}

// ImplementationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Implementation contract.
type ImplementationInitializedIterator struct {
	Event *ImplementationInitialized // Event containing the contract specifics and raw log

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
func (it *ImplementationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationInitialized)
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
		it.Event = new(ImplementationInitialized)
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
func (it *ImplementationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationInitialized represents a Initialized event raised by the Implementation contract.
type ImplementationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Implementation *ImplementationFilterer) FilterInitialized(opts *bind.FilterOpts) (*ImplementationInitializedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ImplementationInitializedIterator{contract: _Implementation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Implementation *ImplementationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ImplementationInitialized) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationInitialized)
				if err := _Implementation.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Implementation *ImplementationFilterer) ParseInitialized(log types.Log) (*ImplementationInitialized, error) {
	event := new(ImplementationInitialized)
	if err := _Implementation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationCipherTextUpdatedIterator is returned from FilterCipherTextUpdated and is used to iterate over the raw logs and unpacked data for CipherTextUpdated events raised by the Implementation contract.
type ImplementationCipherTextUpdatedIterator struct {
	Event *ImplementationCipherTextUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationCipherTextUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationCipherTextUpdated)
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
		it.Event = new(ImplementationCipherTextUpdated)
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
func (it *ImplementationCipherTextUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationCipherTextUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationCipherTextUpdated represents a CipherTextUpdated event raised by the Implementation contract.
type ImplementationCipherTextUpdated struct {
	NewCipherText string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCipherTextUpdated is a free log retrieval operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) FilterCipherTextUpdated(opts *bind.FilterOpts) (*ImplementationCipherTextUpdatedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "cipherTextUpdated")
	if err != nil {
		return nil, err
	}
	return &ImplementationCipherTextUpdatedIterator{contract: _Implementation.contract, event: "cipherTextUpdated", logs: logs, sub: sub}, nil
}

// WatchCipherTextUpdated is a free log subscription operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) WatchCipherTextUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationCipherTextUpdated) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "cipherTextUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationCipherTextUpdated)
				if err := _Implementation.contract.UnpackLog(event, "cipherTextUpdated", log); err != nil {
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

// ParseCipherTextUpdated is a log parse operation binding the contract event 0x2301ef7d9f42b857543faf9e285b5807e028d4ae99810ea7fe0aadda3a717e9d.
//
// Solidity: event cipherTextUpdated(string newCipherText)
func (_Implementation *ImplementationFilterer) ParseCipherTextUpdated(log types.Log) (*ImplementationCipherTextUpdated, error) {
	event := new(ImplementationCipherTextUpdated)
	if err := _Implementation.contract.UnpackLog(event, "cipherTextUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationClosedEarlyIterator is returned from FilterClosedEarly and is used to iterate over the raw logs and unpacked data for ClosedEarly events raised by the Implementation contract.
type ImplementationClosedEarlyIterator struct {
	Event *ImplementationClosedEarly // Event containing the contract specifics and raw log

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
func (it *ImplementationClosedEarlyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationClosedEarly)
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
		it.Event = new(ImplementationClosedEarly)
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
func (it *ImplementationClosedEarlyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationClosedEarlyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationClosedEarly represents a ClosedEarly event raised by the Implementation contract.
type ImplementationClosedEarly struct {
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClosedEarly is a free log retrieval operation binding the contract event 0x55ba7028426dc073c5d2688d6c531f1aeeb2c89b28f341f0ed254266a0c856fd.
//
// Solidity: event closedEarly(uint8 _reason)
func (_Implementation *ImplementationFilterer) FilterClosedEarly(opts *bind.FilterOpts) (*ImplementationClosedEarlyIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "closedEarly")
	if err != nil {
		return nil, err
	}
	return &ImplementationClosedEarlyIterator{contract: _Implementation.contract, event: "closedEarly", logs: logs, sub: sub}, nil
}

// WatchClosedEarly is a free log subscription operation binding the contract event 0x55ba7028426dc073c5d2688d6c531f1aeeb2c89b28f341f0ed254266a0c856fd.
//
// Solidity: event closedEarly(uint8 _reason)
func (_Implementation *ImplementationFilterer) WatchClosedEarly(opts *bind.WatchOpts, sink chan<- *ImplementationClosedEarly) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "closedEarly")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationClosedEarly)
				if err := _Implementation.contract.UnpackLog(event, "closedEarly", log); err != nil {
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

// ParseClosedEarly is a log parse operation binding the contract event 0x55ba7028426dc073c5d2688d6c531f1aeeb2c89b28f341f0ed254266a0c856fd.
//
// Solidity: event closedEarly(uint8 _reason)
func (_Implementation *ImplementationFilterer) ParseClosedEarly(log types.Log) (*ImplementationClosedEarly, error) {
	event := new(ImplementationClosedEarly)
	if err := _Implementation.contract.UnpackLog(event, "closedEarly", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractClosedIterator is returned from FilterContractClosed and is used to iterate over the raw logs and unpacked data for ContractClosed events raised by the Implementation contract.
type ImplementationContractClosedIterator struct {
	Event *ImplementationContractClosed // Event containing the contract specifics and raw log

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
func (it *ImplementationContractClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractClosed)
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
		it.Event = new(ImplementationContractClosed)
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
func (it *ImplementationContractClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractClosed represents a ContractClosed event raised by the Implementation contract.
type ImplementationContractClosed struct {
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractClosed is a free log retrieval operation binding the contract event 0xaadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512.
//
// Solidity: event contractClosed(address indexed _buyer)
func (_Implementation *ImplementationFilterer) FilterContractClosed(opts *bind.FilterOpts, _buyer []common.Address) (*ImplementationContractClosedIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractClosed", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationContractClosedIterator{contract: _Implementation.contract, event: "contractClosed", logs: logs, sub: sub}, nil
}

// WatchContractClosed is a free log subscription operation binding the contract event 0xaadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512.
//
// Solidity: event contractClosed(address indexed _buyer)
func (_Implementation *ImplementationFilterer) WatchContractClosed(opts *bind.WatchOpts, sink chan<- *ImplementationContractClosed, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractClosed", _buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractClosed)
				if err := _Implementation.contract.UnpackLog(event, "contractClosed", log); err != nil {
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

// ParseContractClosed is a log parse operation binding the contract event 0xaadd128c35976a01ffffa9dfb8d363b3358597ce6b30248bcf25e80bd3af4512.
//
// Solidity: event contractClosed(address indexed _buyer)
func (_Implementation *ImplementationFilterer) ParseContractClosed(log types.Log) (*ImplementationContractClosed, error) {
	event := new(ImplementationContractClosed)
	if err := _Implementation.contract.UnpackLog(event, "contractClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractPurchasedIterator is returned from FilterContractPurchased and is used to iterate over the raw logs and unpacked data for ContractPurchased events raised by the Implementation contract.
type ImplementationContractPurchasedIterator struct {
	Event *ImplementationContractPurchased // Event containing the contract specifics and raw log

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
func (it *ImplementationContractPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractPurchased)
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
		it.Event = new(ImplementationContractPurchased)
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
func (it *ImplementationContractPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractPurchased represents a ContractPurchased event raised by the Implementation contract.
type ImplementationContractPurchased struct {
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterContractPurchased is a free log retrieval operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) FilterContractPurchased(opts *bind.FilterOpts, _buyer []common.Address) (*ImplementationContractPurchasedIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractPurchased", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationContractPurchasedIterator{contract: _Implementation.contract, event: "contractPurchased", logs: logs, sub: sub}, nil
}

// WatchContractPurchased is a free log subscription operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) WatchContractPurchased(opts *bind.WatchOpts, sink chan<- *ImplementationContractPurchased, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractPurchased", _buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractPurchased)
				if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
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

// ParseContractPurchased is a log parse operation binding the contract event 0x0c00d1d6cea0bd55f7d3b6e92ef60237b117b050185fc2816c708fd45f45e5bb.
//
// Solidity: event contractPurchased(address indexed _buyer)
func (_Implementation *ImplementationFilterer) ParseContractPurchased(log types.Log) (*ImplementationContractPurchased, error) {
	event := new(ImplementationContractPurchased)
	if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationDestinationUpdatedIterator is returned from FilterDestinationUpdated and is used to iterate over the raw logs and unpacked data for DestinationUpdated events raised by the Implementation contract.
type ImplementationDestinationUpdatedIterator struct {
	Event *ImplementationDestinationUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationDestinationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationDestinationUpdated)
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
		it.Event = new(ImplementationDestinationUpdated)
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
func (it *ImplementationDestinationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationDestinationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationDestinationUpdated represents a DestinationUpdated event raised by the Implementation contract.
type ImplementationDestinationUpdated struct {
	NewValidatorURL string
	NewDestURL      string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestinationUpdated is a free log retrieval operation binding the contract event 0x74e9d8ab61381d8415dfd359c6404d06a0f357a5fda5f29a22fe9512e9277637.
//
// Solidity: event destinationUpdated(string newValidatorURL, string newDestURL)
func (_Implementation *ImplementationFilterer) FilterDestinationUpdated(opts *bind.FilterOpts) (*ImplementationDestinationUpdatedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "destinationUpdated")
	if err != nil {
		return nil, err
	}
	return &ImplementationDestinationUpdatedIterator{contract: _Implementation.contract, event: "destinationUpdated", logs: logs, sub: sub}, nil
}

// WatchDestinationUpdated is a free log subscription operation binding the contract event 0x74e9d8ab61381d8415dfd359c6404d06a0f357a5fda5f29a22fe9512e9277637.
//
// Solidity: event destinationUpdated(string newValidatorURL, string newDestURL)
func (_Implementation *ImplementationFilterer) WatchDestinationUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationDestinationUpdated) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "destinationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationDestinationUpdated)
				if err := _Implementation.contract.UnpackLog(event, "destinationUpdated", log); err != nil {
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

// ParseDestinationUpdated is a log parse operation binding the contract event 0x74e9d8ab61381d8415dfd359c6404d06a0f357a5fda5f29a22fe9512e9277637.
//
// Solidity: event destinationUpdated(string newValidatorURL, string newDestURL)
func (_Implementation *ImplementationFilterer) ParseDestinationUpdated(log types.Log) (*ImplementationDestinationUpdated, error) {
	event := new(ImplementationDestinationUpdated)
	if err := _Implementation.contract.UnpackLog(event, "destinationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationFundsClaimedIterator is returned from FilterFundsClaimed and is used to iterate over the raw logs and unpacked data for FundsClaimed events raised by the Implementation contract.
type ImplementationFundsClaimedIterator struct {
	Event *ImplementationFundsClaimed // Event containing the contract specifics and raw log

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
func (it *ImplementationFundsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationFundsClaimed)
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
		it.Event = new(ImplementationFundsClaimed)
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
func (it *ImplementationFundsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationFundsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationFundsClaimed represents a FundsClaimed event raised by the Implementation contract.
type ImplementationFundsClaimed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterFundsClaimed is a free log retrieval operation binding the contract event 0x5c6a0246cde15dcaffe4c495b5be4c3e0525a2ba9c3f03d6f32095fe118c376a.
//
// Solidity: event fundsClaimed()
func (_Implementation *ImplementationFilterer) FilterFundsClaimed(opts *bind.FilterOpts) (*ImplementationFundsClaimedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "fundsClaimed")
	if err != nil {
		return nil, err
	}
	return &ImplementationFundsClaimedIterator{contract: _Implementation.contract, event: "fundsClaimed", logs: logs, sub: sub}, nil
}

// WatchFundsClaimed is a free log subscription operation binding the contract event 0x5c6a0246cde15dcaffe4c495b5be4c3e0525a2ba9c3f03d6f32095fe118c376a.
//
// Solidity: event fundsClaimed()
func (_Implementation *ImplementationFilterer) WatchFundsClaimed(opts *bind.WatchOpts, sink chan<- *ImplementationFundsClaimed) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "fundsClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationFundsClaimed)
				if err := _Implementation.contract.UnpackLog(event, "fundsClaimed", log); err != nil {
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

// ParseFundsClaimed is a log parse operation binding the contract event 0x5c6a0246cde15dcaffe4c495b5be4c3e0525a2ba9c3f03d6f32095fe118c376a.
//
// Solidity: event fundsClaimed()
func (_Implementation *ImplementationFilterer) ParseFundsClaimed(log types.Log) (*ImplementationFundsClaimed, error) {
	event := new(ImplementationFundsClaimed)
	if err := _Implementation.contract.UnpackLog(event, "fundsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationFundsClaimedValidatorIterator is returned from FilterFundsClaimedValidator and is used to iterate over the raw logs and unpacked data for FundsClaimedValidator events raised by the Implementation contract.
type ImplementationFundsClaimedValidatorIterator struct {
	Event *ImplementationFundsClaimedValidator // Event containing the contract specifics and raw log

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
func (it *ImplementationFundsClaimedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationFundsClaimedValidator)
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
		it.Event = new(ImplementationFundsClaimedValidator)
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
func (it *ImplementationFundsClaimedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationFundsClaimedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationFundsClaimedValidator represents a FundsClaimedValidator event raised by the Implementation contract.
type ImplementationFundsClaimedValidator struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundsClaimedValidator is a free log retrieval operation binding the contract event 0x97e6caeb0ee93285bb3391ae7ecf34d0f5a555cca282be55288f16a76686efe6.
//
// Solidity: event fundsClaimedValidator(address indexed _validator)
func (_Implementation *ImplementationFilterer) FilterFundsClaimedValidator(opts *bind.FilterOpts, _validator []common.Address) (*ImplementationFundsClaimedValidatorIterator, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "fundsClaimedValidator", _validatorRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationFundsClaimedValidatorIterator{contract: _Implementation.contract, event: "fundsClaimedValidator", logs: logs, sub: sub}, nil
}

// WatchFundsClaimedValidator is a free log subscription operation binding the contract event 0x97e6caeb0ee93285bb3391ae7ecf34d0f5a555cca282be55288f16a76686efe6.
//
// Solidity: event fundsClaimedValidator(address indexed _validator)
func (_Implementation *ImplementationFilterer) WatchFundsClaimedValidator(opts *bind.WatchOpts, sink chan<- *ImplementationFundsClaimedValidator, _validator []common.Address) (event.Subscription, error) {

	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "fundsClaimedValidator", _validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationFundsClaimedValidator)
				if err := _Implementation.contract.UnpackLog(event, "fundsClaimedValidator", log); err != nil {
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

// ParseFundsClaimedValidator is a log parse operation binding the contract event 0x97e6caeb0ee93285bb3391ae7ecf34d0f5a555cca282be55288f16a76686efe6.
//
// Solidity: event fundsClaimedValidator(address indexed _validator)
func (_Implementation *ImplementationFilterer) ParseFundsClaimedValidator(log types.Log) (*ImplementationFundsClaimedValidator, error) {
	event := new(ImplementationFundsClaimedValidator)
	if err := _Implementation.contract.UnpackLog(event, "fundsClaimedValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationPurchaseInfoUpdatedIterator is returned from FilterPurchaseInfoUpdated and is used to iterate over the raw logs and unpacked data for PurchaseInfoUpdated events raised by the Implementation contract.
type ImplementationPurchaseInfoUpdatedIterator struct {
	Event *ImplementationPurchaseInfoUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationPurchaseInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationPurchaseInfoUpdated)
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
		it.Event = new(ImplementationPurchaseInfoUpdated)
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
func (it *ImplementationPurchaseInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationPurchaseInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationPurchaseInfoUpdated represents a PurchaseInfoUpdated event raised by the Implementation contract.
type ImplementationPurchaseInfoUpdated struct {
	Address common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPurchaseInfoUpdated is a free log retrieval operation binding the contract event 0x7865b2bbe0087425b223d9cf674d8d5328e31546b364da98c36db21193c17d55.
//
// Solidity: event purchaseInfoUpdated(address indexed _address)
func (_Implementation *ImplementationFilterer) FilterPurchaseInfoUpdated(opts *bind.FilterOpts, _address []common.Address) (*ImplementationPurchaseInfoUpdatedIterator, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "purchaseInfoUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationPurchaseInfoUpdatedIterator{contract: _Implementation.contract, event: "purchaseInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPurchaseInfoUpdated is a free log subscription operation binding the contract event 0x7865b2bbe0087425b223d9cf674d8d5328e31546b364da98c36db21193c17d55.
//
// Solidity: event purchaseInfoUpdated(address indexed _address)
func (_Implementation *ImplementationFilterer) WatchPurchaseInfoUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationPurchaseInfoUpdated, _address []common.Address) (event.Subscription, error) {

	var _addressRule []interface{}
	for _, _addressItem := range _address {
		_addressRule = append(_addressRule, _addressItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "purchaseInfoUpdated", _addressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationPurchaseInfoUpdated)
				if err := _Implementation.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
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
func (_Implementation *ImplementationFilterer) ParsePurchaseInfoUpdated(log types.Log) (*ImplementationPurchaseInfoUpdated, error) {
	event := new(ImplementationPurchaseInfoUpdated)
	if err := _Implementation.contract.UnpackLog(event, "purchaseInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
