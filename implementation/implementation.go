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

// ImplementationResellTerms is an auto generated low-level Go binding around an user-defined struct.
type ImplementationResellTerms struct {
	Buyer                  common.Address
	Validator              common.Address
	Price                  *big.Int
	Fee                    *big.Int
	StartTime              *big.Int
	EncrDestURL            string
	EncrValidatorURL       string
	LastSettlementTime     *big.Int
	Seller                 common.Address
	ResellPrice            *big.Int
	ResellProfitTarget     int8
	IsResellable           bool
	IsResellToDefaultBuyer bool
}

// ResellFlags is an auto generated low-level Go binding around an user-defined struct.
type ResellFlags struct {
	IsResellable           bool
	IsResellToDefaultBuyer bool
}

// ImplementationMetaData contains all meta data concerning the Implementation contract.
var ImplementationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_cloneFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_hashrateOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumImplementation.CloseReason\",\"name\":\"_reason\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isResellable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isResellToDefaultBuyer\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structResellFlags\",\"name\":\"_resellFlags\",\"type\":\"tuple\"}],\"name\":\"contractClosedEarly\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_resellPrice\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isResellable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isResellToDefaultBuyer\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structResellFlags\",\"name\":\"_resellFlags\",\"type\":\"tuple\"}],\"name\":\"contractPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"}],\"name\":\"contractTermsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newValidatorURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newDestURL\",\"type\":\"string\"}],\"name\":\"destinationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"fundsClaimed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"VALIDATOR_FEE_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimFunds\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cloneFactory\",\"outputs\":[{\"internalType\":\"contractCloneFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumImplementation.CloseReason\",\"name\":\"reason\",\"type\":\"uint8\"}],\"name\":\"closeEarly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractState\",\"outputs\":[{\"internalType\":\"enumImplementation.ContractState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureTerms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestResell\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_lastSettlementTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_resellPrice\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_resellProfitTarget\",\"type\":\"int8\"},{\"internalType\":\"bool\",\"name\":\"_isResellable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isResellToDefaultBuyer\",\"type\":\"bool\"}],\"internalType\":\"structImplementation.ResellTerms\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hashrateOracle\",\"outputs\":[{\"internalType\":\"contractHashrateOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_pubKey\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isReselling\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceAndFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"_profitTarget\",\"type\":\"int8\"}],\"name\":\"priceV2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"resellChain\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_lastSettlementTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_resellPrice\",\"type\":\"uint256\"},{\"internalType\":\"int8\",\"name\":\"_resellProfitTarget\",\"type\":\"int8\"},{\"internalType\":\"bool\",\"name\":\"_isResellable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_isResellToDefaultBuyer\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isDeleted\",\"type\":\"bool\"}],\"name\":\"setContractDeleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"}],\"name\":\"setDestination\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_encrValidatorURL\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_encrDestURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validator\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isResellable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isResellToDefaultBuyer\",\"type\":\"bool\"}],\"internalType\":\"structResellFlags\",\"name\":\"_resellFlags\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_resellPrice\",\"type\":\"uint256\"}],\"name\":\"setPurchaseContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setTerms\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"successCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"terms\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_speed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// VALIDATORFEEDECIMALS is a free data retrieval call binding the contract method 0x7cbcd50d.
//
// Solidity: function VALIDATOR_FEE_DECIMALS() view returns(uint8)
func (_Implementation *ImplementationCaller) VALIDATORFEEDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "VALIDATOR_FEE_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VALIDATORFEEDECIMALS is a free data retrieval call binding the contract method 0x7cbcd50d.
//
// Solidity: function VALIDATOR_FEE_DECIMALS() view returns(uint8)
func (_Implementation *ImplementationSession) VALIDATORFEEDECIMALS() (uint8, error) {
	return _Implementation.Contract.VALIDATORFEEDECIMALS(&_Implementation.CallOpts)
}

// VALIDATORFEEDECIMALS is a free data retrieval call binding the contract method 0x7cbcd50d.
//
// Solidity: function VALIDATOR_FEE_DECIMALS() view returns(uint8)
func (_Implementation *ImplementationCallerSession) VALIDATORFEEDECIMALS() (uint8, error) {
	return _Implementation.Contract.VALIDATORFEEDECIMALS(&_Implementation.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Implementation *ImplementationCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Implementation *ImplementationSession) VERSION() (string, error) {
	return _Implementation.Contract.VERSION(&_Implementation.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Implementation *ImplementationCallerSession) VERSION() (string, error) {
	return _Implementation.Contract.VERSION(&_Implementation.CallOpts)
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

// FailCount is a free data retrieval call binding the contract method 0x8cb030bf.
//
// Solidity: function failCount() view returns(uint32)
func (_Implementation *ImplementationCaller) FailCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "failCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// FailCount is a free data retrieval call binding the contract method 0x8cb030bf.
//
// Solidity: function failCount() view returns(uint32)
func (_Implementation *ImplementationSession) FailCount() (uint32, error) {
	return _Implementation.Contract.FailCount(&_Implementation.CallOpts)
}

// FailCount is a free data retrieval call binding the contract method 0x8cb030bf.
//
// Solidity: function failCount() view returns(uint32)
func (_Implementation *ImplementationCallerSession) FailCount() (uint32, error) {
	return _Implementation.Contract.FailCount(&_Implementation.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Implementation *ImplementationCaller) FeeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "feeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Implementation *ImplementationSession) FeeToken() (common.Address, error) {
	return _Implementation.Contract.FeeToken(&_Implementation.CallOpts)
}

// FeeToken is a free data retrieval call binding the contract method 0x647846a5.
//
// Solidity: function feeToken() view returns(address)
func (_Implementation *ImplementationCallerSession) FeeToken() (common.Address, error) {
	return _Implementation.Contract.FeeToken(&_Implementation.CallOpts)
}

// FutureTerms is a free data retrieval call binding the contract method 0xa4e27566.
//
// Solidity: function futureTerms() view returns(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationCaller) FutureTerms(opts *bind.CallOpts) (struct {
	Speed   *big.Int
	Length  *big.Int
	Version uint32
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "futureTerms")

	outstruct := new(struct {
		Speed   *big.Int
		Length  *big.Int
		Version uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Speed = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// FutureTerms is a free data retrieval call binding the contract method 0xa4e27566.
//
// Solidity: function futureTerms() view returns(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationSession) FutureTerms() (struct {
	Speed   *big.Int
	Length  *big.Int
	Version uint32
}, error) {
	return _Implementation.Contract.FutureTerms(&_Implementation.CallOpts)
}

// FutureTerms is a free data retrieval call binding the contract method 0xa4e27566.
//
// Solidity: function futureTerms() view returns(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationCallerSession) FutureTerms() (struct {
	Speed   *big.Int
	Length  *big.Int
	Version uint32
}, error) {
	return _Implementation.Contract.FutureTerms(&_Implementation.CallOpts)
}

// GetLatestResell is a free data retrieval call binding the contract method 0x072b3c50.
//
// Solidity: function getLatestResell() view returns((address,address,uint256,uint256,uint256,string,string,uint256,address,uint256,int8,bool,bool))
func (_Implementation *ImplementationCaller) GetLatestResell(opts *bind.CallOpts) (ImplementationResellTerms, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "getLatestResell")

	if err != nil {
		return *new(ImplementationResellTerms), err
	}

	out0 := *abi.ConvertType(out[0], new(ImplementationResellTerms)).(*ImplementationResellTerms)

	return out0, err

}

// GetLatestResell is a free data retrieval call binding the contract method 0x072b3c50.
//
// Solidity: function getLatestResell() view returns((address,address,uint256,uint256,uint256,string,string,uint256,address,uint256,int8,bool,bool))
func (_Implementation *ImplementationSession) GetLatestResell() (ImplementationResellTerms, error) {
	return _Implementation.Contract.GetLatestResell(&_Implementation.CallOpts)
}

// GetLatestResell is a free data retrieval call binding the contract method 0x072b3c50.
//
// Solidity: function getLatestResell() view returns((address,address,uint256,uint256,uint256,string,string,uint256,address,uint256,int8,bool,bool))
func (_Implementation *ImplementationCallerSession) GetLatestResell() (ImplementationResellTerms, error) {
	return _Implementation.Contract.GetLatestResell(&_Implementation.CallOpts)
}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Implementation *ImplementationCaller) HashrateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "hashrateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Implementation *ImplementationSession) HashrateOracle() (common.Address, error) {
	return _Implementation.Contract.HashrateOracle(&_Implementation.CallOpts)
}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Implementation *ImplementationCallerSession) HashrateOracle() (common.Address, error) {
	return _Implementation.Contract.HashrateOracle(&_Implementation.CallOpts)
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

// IsReselling is a free data retrieval call binding the contract method 0xd52c0453.
//
// Solidity: function isReselling() view returns(bool)
func (_Implementation *ImplementationCaller) IsReselling(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "isReselling")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReselling is a free data retrieval call binding the contract method 0xd52c0453.
//
// Solidity: function isReselling() view returns(bool)
func (_Implementation *ImplementationSession) IsReselling() (bool, error) {
	return _Implementation.Contract.IsReselling(&_Implementation.CallOpts)
}

// IsReselling is a free data retrieval call binding the contract method 0xd52c0453.
//
// Solidity: function isReselling() view returns(bool)
func (_Implementation *ImplementationCallerSession) IsReselling() (bool, error) {
	return _Implementation.Contract.IsReselling(&_Implementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Implementation *ImplementationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Implementation *ImplementationSession) Owner() (common.Address, error) {
	return _Implementation.Contract.Owner(&_Implementation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Implementation *ImplementationCallerSession) Owner() (common.Address, error) {
	return _Implementation.Contract.Owner(&_Implementation.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Implementation *ImplementationCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Implementation *ImplementationSession) PaymentToken() (common.Address, error) {
	return _Implementation.Contract.PaymentToken(&_Implementation.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_Implementation *ImplementationCallerSession) PaymentToken() (common.Address, error) {
	return _Implementation.Contract.PaymentToken(&_Implementation.CallOpts)
}

// PriceAndFee is a free data retrieval call binding the contract method 0x10ad1462.
//
// Solidity: function priceAndFee() view returns(uint256, uint256)
func (_Implementation *ImplementationCaller) PriceAndFee(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "priceAndFee")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// PriceAndFee is a free data retrieval call binding the contract method 0x10ad1462.
//
// Solidity: function priceAndFee() view returns(uint256, uint256)
func (_Implementation *ImplementationSession) PriceAndFee() (*big.Int, *big.Int, error) {
	return _Implementation.Contract.PriceAndFee(&_Implementation.CallOpts)
}

// PriceAndFee is a free data retrieval call binding the contract method 0x10ad1462.
//
// Solidity: function priceAndFee() view returns(uint256, uint256)
func (_Implementation *ImplementationCallerSession) PriceAndFee() (*big.Int, *big.Int, error) {
	return _Implementation.Contract.PriceAndFee(&_Implementation.CallOpts)
}

// PriceV2 is a free data retrieval call binding the contract method 0x9b865926.
//
// Solidity: function priceV2(int8 _profitTarget) view returns(uint256)
func (_Implementation *ImplementationCaller) PriceV2(opts *bind.CallOpts, _profitTarget int8) (*big.Int, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "priceV2", _profitTarget)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceV2 is a free data retrieval call binding the contract method 0x9b865926.
//
// Solidity: function priceV2(int8 _profitTarget) view returns(uint256)
func (_Implementation *ImplementationSession) PriceV2(_profitTarget int8) (*big.Int, error) {
	return _Implementation.Contract.PriceV2(&_Implementation.CallOpts, _profitTarget)
}

// PriceV2 is a free data retrieval call binding the contract method 0x9b865926.
//
// Solidity: function priceV2(int8 _profitTarget) view returns(uint256)
func (_Implementation *ImplementationCallerSession) PriceV2(_profitTarget int8) (*big.Int, error) {
	return _Implementation.Contract.PriceV2(&_Implementation.CallOpts, _profitTarget)
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

// ResellChain is a free data retrieval call binding the contract method 0xd02a1c40.
//
// Solidity: function resellChain(uint256 ) view returns(address _buyer, address _validator, uint256 _price, uint256 _fee, uint256 _startTime, string _encrDestURL, string _encrValidatorURL, uint256 _lastSettlementTime, address _seller, uint256 _resellPrice, int8 _resellProfitTarget, bool _isResellable, bool _isResellToDefaultBuyer)
func (_Implementation *ImplementationCaller) ResellChain(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Buyer                  common.Address
	Validator              common.Address
	Price                  *big.Int
	Fee                    *big.Int
	StartTime              *big.Int
	EncrDestURL            string
	EncrValidatorURL       string
	LastSettlementTime     *big.Int
	Seller                 common.Address
	ResellPrice            *big.Int
	ResellProfitTarget     int8
	IsResellable           bool
	IsResellToDefaultBuyer bool
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "resellChain", arg0)

	outstruct := new(struct {
		Buyer                  common.Address
		Validator              common.Address
		Price                  *big.Int
		Fee                    *big.Int
		StartTime              *big.Int
		EncrDestURL            string
		EncrValidatorURL       string
		LastSettlementTime     *big.Int
		Seller                 common.Address
		ResellPrice            *big.Int
		ResellProfitTarget     int8
		IsResellable           bool
		IsResellToDefaultBuyer bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Buyer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Validator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EncrDestURL = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.EncrValidatorURL = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.LastSettlementTime = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.ResellPrice = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.ResellProfitTarget = *abi.ConvertType(out[10], new(int8)).(*int8)
	outstruct.IsResellable = *abi.ConvertType(out[11], new(bool)).(*bool)
	outstruct.IsResellToDefaultBuyer = *abi.ConvertType(out[12], new(bool)).(*bool)

	return *outstruct, err

}

// ResellChain is a free data retrieval call binding the contract method 0xd02a1c40.
//
// Solidity: function resellChain(uint256 ) view returns(address _buyer, address _validator, uint256 _price, uint256 _fee, uint256 _startTime, string _encrDestURL, string _encrValidatorURL, uint256 _lastSettlementTime, address _seller, uint256 _resellPrice, int8 _resellProfitTarget, bool _isResellable, bool _isResellToDefaultBuyer)
func (_Implementation *ImplementationSession) ResellChain(arg0 *big.Int) (struct {
	Buyer                  common.Address
	Validator              common.Address
	Price                  *big.Int
	Fee                    *big.Int
	StartTime              *big.Int
	EncrDestURL            string
	EncrValidatorURL       string
	LastSettlementTime     *big.Int
	Seller                 common.Address
	ResellPrice            *big.Int
	ResellProfitTarget     int8
	IsResellable           bool
	IsResellToDefaultBuyer bool
}, error) {
	return _Implementation.Contract.ResellChain(&_Implementation.CallOpts, arg0)
}

// ResellChain is a free data retrieval call binding the contract method 0xd02a1c40.
//
// Solidity: function resellChain(uint256 ) view returns(address _buyer, address _validator, uint256 _price, uint256 _fee, uint256 _startTime, string _encrDestURL, string _encrValidatorURL, uint256 _lastSettlementTime, address _seller, uint256 _resellPrice, int8 _resellProfitTarget, bool _isResellable, bool _isResellToDefaultBuyer)
func (_Implementation *ImplementationCallerSession) ResellChain(arg0 *big.Int) (struct {
	Buyer                  common.Address
	Validator              common.Address
	Price                  *big.Int
	Fee                    *big.Int
	StartTime              *big.Int
	EncrDestURL            string
	EncrValidatorURL       string
	LastSettlementTime     *big.Int
	Seller                 common.Address
	ResellPrice            *big.Int
	ResellProfitTarget     int8
	IsResellable           bool
	IsResellToDefaultBuyer bool
}, error) {
	return _Implementation.Contract.ResellChain(&_Implementation.CallOpts, arg0)
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

// SuccessCount is a free data retrieval call binding the contract method 0x8cf086ad.
//
// Solidity: function successCount() view returns(uint32)
func (_Implementation *ImplementationCaller) SuccessCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "successCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SuccessCount is a free data retrieval call binding the contract method 0x8cf086ad.
//
// Solidity: function successCount() view returns(uint32)
func (_Implementation *ImplementationSession) SuccessCount() (uint32, error) {
	return _Implementation.Contract.SuccessCount(&_Implementation.CallOpts)
}

// SuccessCount is a free data retrieval call binding the contract method 0x8cf086ad.
//
// Solidity: function successCount() view returns(uint32)
func (_Implementation *ImplementationCallerSession) SuccessCount() (uint32, error) {
	return _Implementation.Contract.SuccessCount(&_Implementation.CallOpts)
}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationCaller) Terms(opts *bind.CallOpts) (struct {
	Speed   *big.Int
	Length  *big.Int
	Version uint32
}, error) {
	var out []interface{}
	err := _Implementation.contract.Call(opts, &out, "terms")

	outstruct := new(struct {
		Speed   *big.Int
		Length  *big.Int
		Version uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Speed = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Length = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Version = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationSession) Terms() (struct {
	Speed   *big.Int
	Length  *big.Int
	Version uint32
}, error) {
	return _Implementation.Contract.Terms(&_Implementation.CallOpts)
}

// Terms is a free data retrieval call binding the contract method 0xd5025625.
//
// Solidity: function terms() view returns(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationCallerSession) Terms() (struct {
	Speed   *big.Int
	Length  *big.Int
	Version uint32
}, error) {
	return _Implementation.Contract.Terms(&_Implementation.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x677ccdbc.
//
// Solidity: function initialize(address _seller, string _pubKey, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Implementation *ImplementationTransactor) Initialize(opts *bind.TransactOpts, _seller common.Address, _pubKey string, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "initialize", _seller, _pubKey, _speed, _length, _profitTarget)
}

// Initialize is a paid mutator transaction binding the contract method 0x677ccdbc.
//
// Solidity: function initialize(address _seller, string _pubKey, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Implementation *ImplementationSession) Initialize(_seller common.Address, _pubKey string, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _seller, _pubKey, _speed, _length, _profitTarget)
}

// Initialize is a paid mutator transaction binding the contract method 0x677ccdbc.
//
// Solidity: function initialize(address _seller, string _pubKey, uint256 _speed, uint256 _length, int8 _profitTarget) returns()
func (_Implementation *ImplementationTransactorSession) Initialize(_seller common.Address, _pubKey string, _speed *big.Int, _length *big.Int, _profitTarget int8) (*types.Transaction, error) {
	return _Implementation.Contract.Initialize(&_Implementation.TransactOpts, _seller, _pubKey, _speed, _length, _profitTarget)
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

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x5b880513.
//
// Solidity: function setPurchaseContract(string _encrValidatorURL, string _encrDestURL, uint256 _price, uint256 _fee, address _buyer, address _seller, address _validator, (bool,bool) _resellFlags, uint256 _resellPrice) returns()
func (_Implementation *ImplementationTransactor) SetPurchaseContract(opts *bind.TransactOpts, _encrValidatorURL string, _encrDestURL string, _price *big.Int, _fee *big.Int, _buyer common.Address, _seller common.Address, _validator common.Address, _resellFlags ResellFlags, _resellPrice *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setPurchaseContract", _encrValidatorURL, _encrDestURL, _price, _fee, _buyer, _seller, _validator, _resellFlags, _resellPrice)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x5b880513.
//
// Solidity: function setPurchaseContract(string _encrValidatorURL, string _encrDestURL, uint256 _price, uint256 _fee, address _buyer, address _seller, address _validator, (bool,bool) _resellFlags, uint256 _resellPrice) returns()
func (_Implementation *ImplementationSession) SetPurchaseContract(_encrValidatorURL string, _encrDestURL string, _price *big.Int, _fee *big.Int, _buyer common.Address, _seller common.Address, _validator common.Address, _resellFlags ResellFlags, _resellPrice *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encrValidatorURL, _encrDestURL, _price, _fee, _buyer, _seller, _validator, _resellFlags, _resellPrice)
}

// SetPurchaseContract is a paid mutator transaction binding the contract method 0x5b880513.
//
// Solidity: function setPurchaseContract(string _encrValidatorURL, string _encrDestURL, uint256 _price, uint256 _fee, address _buyer, address _seller, address _validator, (bool,bool) _resellFlags, uint256 _resellPrice) returns()
func (_Implementation *ImplementationTransactorSession) SetPurchaseContract(_encrValidatorURL string, _encrDestURL string, _price *big.Int, _fee *big.Int, _buyer common.Address, _seller common.Address, _validator common.Address, _resellFlags ResellFlags, _resellPrice *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetPurchaseContract(&_Implementation.TransactOpts, _encrValidatorURL, _encrDestURL, _price, _fee, _buyer, _seller, _validator, _resellFlags, _resellPrice)
}

// SetTerms is a paid mutator transaction binding the contract method 0x75abfede.
//
// Solidity: function setTerms(uint256 _speed, uint256 _length) returns()
func (_Implementation *ImplementationTransactor) SetTerms(opts *bind.TransactOpts, _speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Implementation.contract.Transact(opts, "setTerms", _speed, _length)
}

// SetTerms is a paid mutator transaction binding the contract method 0x75abfede.
//
// Solidity: function setTerms(uint256 _speed, uint256 _length) returns()
func (_Implementation *ImplementationSession) SetTerms(_speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetTerms(&_Implementation.TransactOpts, _speed, _length)
}

// SetTerms is a paid mutator transaction binding the contract method 0x75abfede.
//
// Solidity: function setTerms(uint256 _speed, uint256 _length) returns()
func (_Implementation *ImplementationTransactorSession) SetTerms(_speed *big.Int, _length *big.Int) (*types.Transaction, error) {
	return _Implementation.Contract.SetTerms(&_Implementation.TransactOpts, _speed, _length)
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
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Implementation *ImplementationFilterer) FilterInitialized(opts *bind.FilterOpts) (*ImplementationInitializedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ImplementationInitializedIterator{contract: _Implementation.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Implementation *ImplementationFilterer) ParseInitialized(log types.Log) (*ImplementationInitialized, error) {
	event := new(ImplementationInitialized)
	if err := _Implementation.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractClosedEarlyIterator is returned from FilterContractClosedEarly and is used to iterate over the raw logs and unpacked data for ContractClosedEarly events raised by the Implementation contract.
type ImplementationContractClosedEarlyIterator struct {
	Event *ImplementationContractClosedEarly // Event containing the contract specifics and raw log

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
func (it *ImplementationContractClosedEarlyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractClosedEarly)
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
		it.Event = new(ImplementationContractClosedEarly)
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
func (it *ImplementationContractClosedEarlyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractClosedEarlyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractClosedEarly represents a ContractClosedEarly event raised by the Implementation contract.
type ImplementationContractClosedEarly struct {
	Buyer       common.Address
	Validator   common.Address
	Seller      common.Address
	Reason      uint8
	ResellFlags ResellFlags
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractClosedEarly is a free log retrieval operation binding the contract event 0x55a8740d375819a3225051dae7576ea02b2c47d00d756b80fcb63eed40bc4154.
//
// Solidity: event contractClosedEarly(address indexed _buyer, address indexed _validator, address indexed _seller, uint8 _reason, (bool,bool) _resellFlags)
func (_Implementation *ImplementationFilterer) FilterContractClosedEarly(opts *bind.FilterOpts, _buyer []common.Address, _validator []common.Address, _seller []common.Address) (*ImplementationContractClosedEarlyIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractClosedEarly", _buyerRule, _validatorRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationContractClosedEarlyIterator{contract: _Implementation.contract, event: "contractClosedEarly", logs: logs, sub: sub}, nil
}

// WatchContractClosedEarly is a free log subscription operation binding the contract event 0x55a8740d375819a3225051dae7576ea02b2c47d00d756b80fcb63eed40bc4154.
//
// Solidity: event contractClosedEarly(address indexed _buyer, address indexed _validator, address indexed _seller, uint8 _reason, (bool,bool) _resellFlags)
func (_Implementation *ImplementationFilterer) WatchContractClosedEarly(opts *bind.WatchOpts, sink chan<- *ImplementationContractClosedEarly, _buyer []common.Address, _validator []common.Address, _seller []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractClosedEarly", _buyerRule, _validatorRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractClosedEarly)
				if err := _Implementation.contract.UnpackLog(event, "contractClosedEarly", log); err != nil {
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

// ParseContractClosedEarly is a log parse operation binding the contract event 0x55a8740d375819a3225051dae7576ea02b2c47d00d756b80fcb63eed40bc4154.
//
// Solidity: event contractClosedEarly(address indexed _buyer, address indexed _validator, address indexed _seller, uint8 _reason, (bool,bool) _resellFlags)
func (_Implementation *ImplementationFilterer) ParseContractClosedEarly(log types.Log) (*ImplementationContractClosedEarly, error) {
	event := new(ImplementationContractClosedEarly)
	if err := _Implementation.contract.UnpackLog(event, "contractClosedEarly", log); err != nil {
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
	Buyer       common.Address
	Validator   common.Address
	Seller      common.Address
	Price       *big.Int
	Fee         *big.Int
	ResellPrice *big.Int
	ResellFlags ResellFlags
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractPurchased is a free log retrieval operation binding the contract event 0x62136c9cc68edb922f2aae580846a8fac61ad079ac1260dfc888d5027109f3c9.
//
// Solidity: event contractPurchased(address indexed _buyer, address indexed _validator, address indexed _seller, uint256 _price, uint256 _fee, uint256 _resellPrice, (bool,bool) _resellFlags)
func (_Implementation *ImplementationFilterer) FilterContractPurchased(opts *bind.FilterOpts, _buyer []common.Address, _validator []common.Address, _seller []common.Address) (*ImplementationContractPurchasedIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractPurchased", _buyerRule, _validatorRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return &ImplementationContractPurchasedIterator{contract: _Implementation.contract, event: "contractPurchased", logs: logs, sub: sub}, nil
}

// WatchContractPurchased is a free log subscription operation binding the contract event 0x62136c9cc68edb922f2aae580846a8fac61ad079ac1260dfc888d5027109f3c9.
//
// Solidity: event contractPurchased(address indexed _buyer, address indexed _validator, address indexed _seller, uint256 _price, uint256 _fee, uint256 _resellPrice, (bool,bool) _resellFlags)
func (_Implementation *ImplementationFilterer) WatchContractPurchased(opts *bind.WatchOpts, sink chan<- *ImplementationContractPurchased, _buyer []common.Address, _validator []common.Address, _seller []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _validatorRule []interface{}
	for _, _validatorItem := range _validator {
		_validatorRule = append(_validatorRule, _validatorItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractPurchased", _buyerRule, _validatorRule, _sellerRule)
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

// ParseContractPurchased is a log parse operation binding the contract event 0x62136c9cc68edb922f2aae580846a8fac61ad079ac1260dfc888d5027109f3c9.
//
// Solidity: event contractPurchased(address indexed _buyer, address indexed _validator, address indexed _seller, uint256 _price, uint256 _fee, uint256 _resellPrice, (bool,bool) _resellFlags)
func (_Implementation *ImplementationFilterer) ParseContractPurchased(log types.Log) (*ImplementationContractPurchased, error) {
	event := new(ImplementationContractPurchased)
	if err := _Implementation.contract.UnpackLog(event, "contractPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ImplementationContractTermsUpdatedIterator is returned from FilterContractTermsUpdated and is used to iterate over the raw logs and unpacked data for ContractTermsUpdated events raised by the Implementation contract.
type ImplementationContractTermsUpdatedIterator struct {
	Event *ImplementationContractTermsUpdated // Event containing the contract specifics and raw log

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
func (it *ImplementationContractTermsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ImplementationContractTermsUpdated)
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
		it.Event = new(ImplementationContractTermsUpdated)
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
func (it *ImplementationContractTermsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ImplementationContractTermsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ImplementationContractTermsUpdated represents a ContractTermsUpdated event raised by the Implementation contract.
type ImplementationContractTermsUpdated struct {
	Speed   *big.Int
	Length  *big.Int
	Version uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractTermsUpdated is a free log retrieval operation binding the contract event 0xfc88033c029a35f78bd341a41b67bfdea6355bfb60b7bc899c76f7f97431c9fc.
//
// Solidity: event contractTermsUpdated(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationFilterer) FilterContractTermsUpdated(opts *bind.FilterOpts) (*ImplementationContractTermsUpdatedIterator, error) {

	logs, sub, err := _Implementation.contract.FilterLogs(opts, "contractTermsUpdated")
	if err != nil {
		return nil, err
	}
	return &ImplementationContractTermsUpdatedIterator{contract: _Implementation.contract, event: "contractTermsUpdated", logs: logs, sub: sub}, nil
}

// WatchContractTermsUpdated is a free log subscription operation binding the contract event 0xfc88033c029a35f78bd341a41b67bfdea6355bfb60b7bc899c76f7f97431c9fc.
//
// Solidity: event contractTermsUpdated(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationFilterer) WatchContractTermsUpdated(opts *bind.WatchOpts, sink chan<- *ImplementationContractTermsUpdated) (event.Subscription, error) {

	logs, sub, err := _Implementation.contract.WatchLogs(opts, "contractTermsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ImplementationContractTermsUpdated)
				if err := _Implementation.contract.UnpackLog(event, "contractTermsUpdated", log); err != nil {
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

// ParseContractTermsUpdated is a log parse operation binding the contract event 0xfc88033c029a35f78bd341a41b67bfdea6355bfb60b7bc899c76f7f97431c9fc.
//
// Solidity: event contractTermsUpdated(uint256 _speed, uint256 _length, uint32 _version)
func (_Implementation *ImplementationFilterer) ParseContractTermsUpdated(log types.Log) (*ImplementationContractTermsUpdated, error) {
	event := new(ImplementationContractTermsUpdated)
	if err := _Implementation.contract.UnpackLog(event, "contractTermsUpdated", log); err != nil {
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
