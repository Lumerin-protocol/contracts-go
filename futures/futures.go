// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package futures

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

// FuturesOrder is an auto generated low-level Go binding around an user-defined struct.
type FuturesOrder struct {
	IsBuy       bool
	Participant common.Address
	DestURL     string
	PricePerDay *big.Int
	DeliveryAt  *big.Int
	CreatedAt   *big.Int
}

// FuturesPosition is an auto generated low-level Go binding around an user-defined struct.
type FuturesPosition struct {
	Seller          common.Address
	Buyer           common.Address
	DestURL         string
	SellPricePerDay *big.Int
	BuyPricePerDay  *big.Int
	DeliveryAt      *big.Int
	CreatedAt       *big.Int
	Paid            bool
}

// FuturesMetaData contains all meta data concerning the Futures contract.
var FuturesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeliveryDateExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeliveryDateNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeliveryDateShouldBeInTheFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeliveryNotFinishedYet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientMarginBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxOrdersPerParticipantReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyPositionBuyer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyValidatorOrPositionParticipant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderNotBelongToSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionAlreadyPaid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionDeliveryExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionDeliveryNotStartedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionDestURLNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PositionNotExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"min\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"max\",\"type\":\"int256\"}],\"name\":\"ValueOutOfRange\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"OrderClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deliveryAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderFee\",\"type\":\"uint256\"}],\"name\":\"OrderFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"positionId\",\"type\":\"bytes32\"}],\"name\":\"PositionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"positionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellPricePerDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyPricePerDay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deliveryAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destURL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"}],\"name\":\"PositionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"positionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"closedBy\",\"type\":\"address\"}],\"name\":\"PositionDeliveryClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"positionId\",\"type\":\"bytes32\"}],\"name\":\"PositionPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"positionId\",\"type\":\"bytes32\"}],\"name\":\"PositionPaymentReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BREACH_PENALTY_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ORDERS_PER_PARTICIPANT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"breachPenaltyRatePerDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_positionId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_blameSeller\",\"type\":\"bool\"}],\"name\":\"closeDelivery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"closeOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deliveryDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_destURL\",\"type\":\"string\"},{\"internalType\":\"int8\",\"name\":\"_qty\",\"type\":\"int8\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deliveryDurationDays\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deliveryIntervalDays\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_positionIds\",\"type\":\"bytes32[]\"}],\"name\":\"depositDeliveryPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deliveryDate\",\"type\":\"uint256\"}],\"name\":\"depositDeliveryPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositReservePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstFutureDeliveryDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"futureDeliveryDatesCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"getCollateralDeficit\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeliveryDates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"getMinMargin\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_entryPricePerDay\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_qty\",\"type\":\"int256\"}],\"name\":\"getMinMarginForPosition\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_orderId\",\"type\":\"bytes32\"}],\"name\":\"getOrderById\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"pricePerDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deliveryAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structFutures.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_positionId\",\"type\":\"bytes32\"}],\"name\":\"getPositionById\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"destURL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sellPricePerDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyPricePerDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deliveryAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"paid\",\"type\":\"bool\"}],\"internalType\":\"structFutures.Position\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_deliveryDate\",\"type\":\"uint256\"}],\"name\":\"getPositionsByParticipantDeliveryDate\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hashrateOracle\",\"outputs\":[{\"internalType\":\"contractHashrateOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20Metadata\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"contractHashrateOracle\",\"name\":\"_hashrateOracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_validatorAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_liquidationMarginPercent\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_speedHps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumPriceIncrement\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_deliveryDurationDays\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_deliveryIntervalDays\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_futureDeliveryDatesCount\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_firstFutureDeliveryDate\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationMarginPercent\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"marginCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumPriceIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"orderFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"removeMargin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_participant\",\"type\":\"address\"}],\"name\":\"removeOutdatedOrdersForParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_breachPenaltyRatePerDay\",\"type\":\"uint256\"}],\"name\":\"setBreachPenaltyRatePerDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_futureDeliveryDatesCount\",\"type\":\"uint8\"}],\"name\":\"setFutureDeliveryDatesCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_liquidationMarginPercent\",\"type\":\"uint8\"}],\"name\":\"setLiquidationMarginPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_orderFee\",\"type\":\"uint256\"}],\"name\":\"setOrderFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"speedHps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_deliveryDate\",\"type\":\"uint256\"}],\"name\":\"withdrawDeliveryPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawReservePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FuturesABI is the input ABI used to generate the binding from.
// Deprecated: Use FuturesMetaData.ABI instead.
var FuturesABI = FuturesMetaData.ABI

// Futures is an auto generated Go binding around an Ethereum contract.
type Futures struct {
	FuturesCaller     // Read-only binding to the contract
	FuturesTransactor // Write-only binding to the contract
	FuturesFilterer   // Log filterer for contract events
}

// FuturesCaller is an auto generated read-only Go binding around an Ethereum contract.
type FuturesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuturesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FuturesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuturesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FuturesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FuturesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FuturesSession struct {
	Contract     *Futures          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FuturesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FuturesCallerSession struct {
	Contract *FuturesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FuturesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FuturesTransactorSession struct {
	Contract     *FuturesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FuturesRaw is an auto generated low-level Go binding around an Ethereum contract.
type FuturesRaw struct {
	Contract *Futures // Generic contract binding to access the raw methods on
}

// FuturesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FuturesCallerRaw struct {
	Contract *FuturesCaller // Generic read-only contract binding to access the raw methods on
}

// FuturesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FuturesTransactorRaw struct {
	Contract *FuturesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFutures creates a new instance of Futures, bound to a specific deployed contract.
func NewFutures(address common.Address, backend bind.ContractBackend) (*Futures, error) {
	contract, err := bindFutures(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Futures{FuturesCaller: FuturesCaller{contract: contract}, FuturesTransactor: FuturesTransactor{contract: contract}, FuturesFilterer: FuturesFilterer{contract: contract}}, nil
}

// NewFuturesCaller creates a new read-only instance of Futures, bound to a specific deployed contract.
func NewFuturesCaller(address common.Address, caller bind.ContractCaller) (*FuturesCaller, error) {
	contract, err := bindFutures(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FuturesCaller{contract: contract}, nil
}

// NewFuturesTransactor creates a new write-only instance of Futures, bound to a specific deployed contract.
func NewFuturesTransactor(address common.Address, transactor bind.ContractTransactor) (*FuturesTransactor, error) {
	contract, err := bindFutures(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FuturesTransactor{contract: contract}, nil
}

// NewFuturesFilterer creates a new log filterer instance of Futures, bound to a specific deployed contract.
func NewFuturesFilterer(address common.Address, filterer bind.ContractFilterer) (*FuturesFilterer, error) {
	contract, err := bindFutures(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FuturesFilterer{contract: contract}, nil
}

// bindFutures binds a generic wrapper to an already deployed contract.
func bindFutures(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FuturesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Futures *FuturesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Futures.Contract.FuturesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Futures *FuturesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Futures.Contract.FuturesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Futures *FuturesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Futures.Contract.FuturesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Futures *FuturesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Futures.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Futures *FuturesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Futures.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Futures *FuturesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Futures.Contract.contract.Transact(opts, method, params...)
}

// BREACHPENALTYDECIMALS is a free data retrieval call binding the contract method 0x979c280f.
//
// Solidity: function BREACH_PENALTY_DECIMALS() view returns(uint8)
func (_Futures *FuturesCaller) BREACHPENALTYDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "BREACH_PENALTY_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BREACHPENALTYDECIMALS is a free data retrieval call binding the contract method 0x979c280f.
//
// Solidity: function BREACH_PENALTY_DECIMALS() view returns(uint8)
func (_Futures *FuturesSession) BREACHPENALTYDECIMALS() (uint8, error) {
	return _Futures.Contract.BREACHPENALTYDECIMALS(&_Futures.CallOpts)
}

// BREACHPENALTYDECIMALS is a free data retrieval call binding the contract method 0x979c280f.
//
// Solidity: function BREACH_PENALTY_DECIMALS() view returns(uint8)
func (_Futures *FuturesCallerSession) BREACHPENALTYDECIMALS() (uint8, error) {
	return _Futures.Contract.BREACHPENALTYDECIMALS(&_Futures.CallOpts)
}

// MAXORDERSPERPARTICIPANT is a free data retrieval call binding the contract method 0xa96bc116.
//
// Solidity: function MAX_ORDERS_PER_PARTICIPANT() view returns(uint8)
func (_Futures *FuturesCaller) MAXORDERSPERPARTICIPANT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "MAX_ORDERS_PER_PARTICIPANT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXORDERSPERPARTICIPANT is a free data retrieval call binding the contract method 0xa96bc116.
//
// Solidity: function MAX_ORDERS_PER_PARTICIPANT() view returns(uint8)
func (_Futures *FuturesSession) MAXORDERSPERPARTICIPANT() (uint8, error) {
	return _Futures.Contract.MAXORDERSPERPARTICIPANT(&_Futures.CallOpts)
}

// MAXORDERSPERPARTICIPANT is a free data retrieval call binding the contract method 0xa96bc116.
//
// Solidity: function MAX_ORDERS_PER_PARTICIPANT() view returns(uint8)
func (_Futures *FuturesCallerSession) MAXORDERSPERPARTICIPANT() (uint8, error) {
	return _Futures.Contract.MAXORDERSPERPARTICIPANT(&_Futures.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Futures *FuturesCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Futures *FuturesSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Futures.Contract.UPGRADEINTERFACEVERSION(&_Futures.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Futures *FuturesCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Futures.Contract.UPGRADEINTERFACEVERSION(&_Futures.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Futures *FuturesCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Futures *FuturesSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Futures.Contract.Allowance(&_Futures.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Futures *FuturesCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Futures.Contract.Allowance(&_Futures.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Futures *FuturesCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Futures *FuturesSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Futures.Contract.BalanceOf(&_Futures.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Futures *FuturesCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Futures.Contract.BalanceOf(&_Futures.CallOpts, account)
}

// BreachPenaltyRatePerDay is a free data retrieval call binding the contract method 0xfb69dd70.
//
// Solidity: function breachPenaltyRatePerDay() view returns(uint256)
func (_Futures *FuturesCaller) BreachPenaltyRatePerDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "breachPenaltyRatePerDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BreachPenaltyRatePerDay is a free data retrieval call binding the contract method 0xfb69dd70.
//
// Solidity: function breachPenaltyRatePerDay() view returns(uint256)
func (_Futures *FuturesSession) BreachPenaltyRatePerDay() (*big.Int, error) {
	return _Futures.Contract.BreachPenaltyRatePerDay(&_Futures.CallOpts)
}

// BreachPenaltyRatePerDay is a free data retrieval call binding the contract method 0xfb69dd70.
//
// Solidity: function breachPenaltyRatePerDay() view returns(uint256)
func (_Futures *FuturesCallerSession) BreachPenaltyRatePerDay() (*big.Int, error) {
	return _Futures.Contract.BreachPenaltyRatePerDay(&_Futures.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Futures *FuturesCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Futures *FuturesSession) Decimals() (uint8, error) {
	return _Futures.Contract.Decimals(&_Futures.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Futures *FuturesCallerSession) Decimals() (uint8, error) {
	return _Futures.Contract.Decimals(&_Futures.CallOpts)
}

// DeliveryDurationDays is a free data retrieval call binding the contract method 0x62ffcab5.
//
// Solidity: function deliveryDurationDays() view returns(uint8)
func (_Futures *FuturesCaller) DeliveryDurationDays(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "deliveryDurationDays")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DeliveryDurationDays is a free data retrieval call binding the contract method 0x62ffcab5.
//
// Solidity: function deliveryDurationDays() view returns(uint8)
func (_Futures *FuturesSession) DeliveryDurationDays() (uint8, error) {
	return _Futures.Contract.DeliveryDurationDays(&_Futures.CallOpts)
}

// DeliveryDurationDays is a free data retrieval call binding the contract method 0x62ffcab5.
//
// Solidity: function deliveryDurationDays() view returns(uint8)
func (_Futures *FuturesCallerSession) DeliveryDurationDays() (uint8, error) {
	return _Futures.Contract.DeliveryDurationDays(&_Futures.CallOpts)
}

// DeliveryIntervalDays is a free data retrieval call binding the contract method 0x4e2cf28f.
//
// Solidity: function deliveryIntervalDays() view returns(uint8)
func (_Futures *FuturesCaller) DeliveryIntervalDays(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "deliveryIntervalDays")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DeliveryIntervalDays is a free data retrieval call binding the contract method 0x4e2cf28f.
//
// Solidity: function deliveryIntervalDays() view returns(uint8)
func (_Futures *FuturesSession) DeliveryIntervalDays() (uint8, error) {
	return _Futures.Contract.DeliveryIntervalDays(&_Futures.CallOpts)
}

// DeliveryIntervalDays is a free data retrieval call binding the contract method 0x4e2cf28f.
//
// Solidity: function deliveryIntervalDays() view returns(uint8)
func (_Futures *FuturesCallerSession) DeliveryIntervalDays() (uint8, error) {
	return _Futures.Contract.DeliveryIntervalDays(&_Futures.CallOpts)
}

// FirstFutureDeliveryDate is a free data retrieval call binding the contract method 0x76615e8d.
//
// Solidity: function firstFutureDeliveryDate() view returns(uint256)
func (_Futures *FuturesCaller) FirstFutureDeliveryDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "firstFutureDeliveryDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstFutureDeliveryDate is a free data retrieval call binding the contract method 0x76615e8d.
//
// Solidity: function firstFutureDeliveryDate() view returns(uint256)
func (_Futures *FuturesSession) FirstFutureDeliveryDate() (*big.Int, error) {
	return _Futures.Contract.FirstFutureDeliveryDate(&_Futures.CallOpts)
}

// FirstFutureDeliveryDate is a free data retrieval call binding the contract method 0x76615e8d.
//
// Solidity: function firstFutureDeliveryDate() view returns(uint256)
func (_Futures *FuturesCallerSession) FirstFutureDeliveryDate() (*big.Int, error) {
	return _Futures.Contract.FirstFutureDeliveryDate(&_Futures.CallOpts)
}

// FutureDeliveryDatesCount is a free data retrieval call binding the contract method 0xcbdc0675.
//
// Solidity: function futureDeliveryDatesCount() view returns(uint8)
func (_Futures *FuturesCaller) FutureDeliveryDatesCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "futureDeliveryDatesCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// FutureDeliveryDatesCount is a free data retrieval call binding the contract method 0xcbdc0675.
//
// Solidity: function futureDeliveryDatesCount() view returns(uint8)
func (_Futures *FuturesSession) FutureDeliveryDatesCount() (uint8, error) {
	return _Futures.Contract.FutureDeliveryDatesCount(&_Futures.CallOpts)
}

// FutureDeliveryDatesCount is a free data retrieval call binding the contract method 0xcbdc0675.
//
// Solidity: function futureDeliveryDatesCount() view returns(uint8)
func (_Futures *FuturesCallerSession) FutureDeliveryDatesCount() (uint8, error) {
	return _Futures.Contract.FutureDeliveryDatesCount(&_Futures.CallOpts)
}

// GetCollateralDeficit is a free data retrieval call binding the contract method 0x441782bd.
//
// Solidity: function getCollateralDeficit(address _participant) view returns(int256)
func (_Futures *FuturesCaller) GetCollateralDeficit(opts *bind.CallOpts, _participant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getCollateralDeficit", _participant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCollateralDeficit is a free data retrieval call binding the contract method 0x441782bd.
//
// Solidity: function getCollateralDeficit(address _participant) view returns(int256)
func (_Futures *FuturesSession) GetCollateralDeficit(_participant common.Address) (*big.Int, error) {
	return _Futures.Contract.GetCollateralDeficit(&_Futures.CallOpts, _participant)
}

// GetCollateralDeficit is a free data retrieval call binding the contract method 0x441782bd.
//
// Solidity: function getCollateralDeficit(address _participant) view returns(int256)
func (_Futures *FuturesCallerSession) GetCollateralDeficit(_participant common.Address) (*big.Int, error) {
	return _Futures.Contract.GetCollateralDeficit(&_Futures.CallOpts, _participant)
}

// GetDeliveryDates is a free data retrieval call binding the contract method 0xe53ac087.
//
// Solidity: function getDeliveryDates() view returns(uint256[])
func (_Futures *FuturesCaller) GetDeliveryDates(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getDeliveryDates")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDeliveryDates is a free data retrieval call binding the contract method 0xe53ac087.
//
// Solidity: function getDeliveryDates() view returns(uint256[])
func (_Futures *FuturesSession) GetDeliveryDates() ([]*big.Int, error) {
	return _Futures.Contract.GetDeliveryDates(&_Futures.CallOpts)
}

// GetDeliveryDates is a free data retrieval call binding the contract method 0xe53ac087.
//
// Solidity: function getDeliveryDates() view returns(uint256[])
func (_Futures *FuturesCallerSession) GetDeliveryDates() ([]*big.Int, error) {
	return _Futures.Contract.GetDeliveryDates(&_Futures.CallOpts)
}

// GetMarketPrice is a free data retrieval call binding the contract method 0x660e16c3.
//
// Solidity: function getMarketPrice() view returns(uint256)
func (_Futures *FuturesCaller) GetMarketPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getMarketPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMarketPrice is a free data retrieval call binding the contract method 0x660e16c3.
//
// Solidity: function getMarketPrice() view returns(uint256)
func (_Futures *FuturesSession) GetMarketPrice() (*big.Int, error) {
	return _Futures.Contract.GetMarketPrice(&_Futures.CallOpts)
}

// GetMarketPrice is a free data retrieval call binding the contract method 0x660e16c3.
//
// Solidity: function getMarketPrice() view returns(uint256)
func (_Futures *FuturesCallerSession) GetMarketPrice() (*big.Int, error) {
	return _Futures.Contract.GetMarketPrice(&_Futures.CallOpts)
}

// GetMinMargin is a free data retrieval call binding the contract method 0x1f83f472.
//
// Solidity: function getMinMargin(address _participant) view returns(int256)
func (_Futures *FuturesCaller) GetMinMargin(opts *bind.CallOpts, _participant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getMinMargin", _participant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinMargin is a free data retrieval call binding the contract method 0x1f83f472.
//
// Solidity: function getMinMargin(address _participant) view returns(int256)
func (_Futures *FuturesSession) GetMinMargin(_participant common.Address) (*big.Int, error) {
	return _Futures.Contract.GetMinMargin(&_Futures.CallOpts, _participant)
}

// GetMinMargin is a free data retrieval call binding the contract method 0x1f83f472.
//
// Solidity: function getMinMargin(address _participant) view returns(int256)
func (_Futures *FuturesCallerSession) GetMinMargin(_participant common.Address) (*big.Int, error) {
	return _Futures.Contract.GetMinMargin(&_Futures.CallOpts, _participant)
}

// GetMinMarginForPosition is a free data retrieval call binding the contract method 0xddad24e3.
//
// Solidity: function getMinMarginForPosition(uint256 _entryPricePerDay, int256 _qty) view returns(int256)
func (_Futures *FuturesCaller) GetMinMarginForPosition(opts *bind.CallOpts, _entryPricePerDay *big.Int, _qty *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getMinMarginForPosition", _entryPricePerDay, _qty)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinMarginForPosition is a free data retrieval call binding the contract method 0xddad24e3.
//
// Solidity: function getMinMarginForPosition(uint256 _entryPricePerDay, int256 _qty) view returns(int256)
func (_Futures *FuturesSession) GetMinMarginForPosition(_entryPricePerDay *big.Int, _qty *big.Int) (*big.Int, error) {
	return _Futures.Contract.GetMinMarginForPosition(&_Futures.CallOpts, _entryPricePerDay, _qty)
}

// GetMinMarginForPosition is a free data retrieval call binding the contract method 0xddad24e3.
//
// Solidity: function getMinMarginForPosition(uint256 _entryPricePerDay, int256 _qty) view returns(int256)
func (_Futures *FuturesCallerSession) GetMinMarginForPosition(_entryPricePerDay *big.Int, _qty *big.Int) (*big.Int, error) {
	return _Futures.Contract.GetMinMarginForPosition(&_Futures.CallOpts, _entryPricePerDay, _qty)
}

// GetOrderById is a free data retrieval call binding the contract method 0xa6399129.
//
// Solidity: function getOrderById(bytes32 _orderId) view returns((bool,address,string,uint256,uint256,uint256))
func (_Futures *FuturesCaller) GetOrderById(opts *bind.CallOpts, _orderId [32]byte) (FuturesOrder, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getOrderById", _orderId)

	if err != nil {
		return *new(FuturesOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(FuturesOrder)).(*FuturesOrder)

	return out0, err

}

// GetOrderById is a free data retrieval call binding the contract method 0xa6399129.
//
// Solidity: function getOrderById(bytes32 _orderId) view returns((bool,address,string,uint256,uint256,uint256))
func (_Futures *FuturesSession) GetOrderById(_orderId [32]byte) (FuturesOrder, error) {
	return _Futures.Contract.GetOrderById(&_Futures.CallOpts, _orderId)
}

// GetOrderById is a free data retrieval call binding the contract method 0xa6399129.
//
// Solidity: function getOrderById(bytes32 _orderId) view returns((bool,address,string,uint256,uint256,uint256))
func (_Futures *FuturesCallerSession) GetOrderById(_orderId [32]byte) (FuturesOrder, error) {
	return _Futures.Contract.GetOrderById(&_Futures.CallOpts, _orderId)
}

// GetPositionById is a free data retrieval call binding the contract method 0xdc563253.
//
// Solidity: function getPositionById(bytes32 _positionId) view returns((address,address,string,uint256,uint256,uint256,uint256,bool))
func (_Futures *FuturesCaller) GetPositionById(opts *bind.CallOpts, _positionId [32]byte) (FuturesPosition, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getPositionById", _positionId)

	if err != nil {
		return *new(FuturesPosition), err
	}

	out0 := *abi.ConvertType(out[0], new(FuturesPosition)).(*FuturesPosition)

	return out0, err

}

// GetPositionById is a free data retrieval call binding the contract method 0xdc563253.
//
// Solidity: function getPositionById(bytes32 _positionId) view returns((address,address,string,uint256,uint256,uint256,uint256,bool))
func (_Futures *FuturesSession) GetPositionById(_positionId [32]byte) (FuturesPosition, error) {
	return _Futures.Contract.GetPositionById(&_Futures.CallOpts, _positionId)
}

// GetPositionById is a free data retrieval call binding the contract method 0xdc563253.
//
// Solidity: function getPositionById(bytes32 _positionId) view returns((address,address,string,uint256,uint256,uint256,uint256,bool))
func (_Futures *FuturesCallerSession) GetPositionById(_positionId [32]byte) (FuturesPosition, error) {
	return _Futures.Contract.GetPositionById(&_Futures.CallOpts, _positionId)
}

// GetPositionsByParticipantDeliveryDate is a free data retrieval call binding the contract method 0x207c1f89.
//
// Solidity: function getPositionsByParticipantDeliveryDate(address _participant, uint256 _deliveryDate) view returns(bytes32[])
func (_Futures *FuturesCaller) GetPositionsByParticipantDeliveryDate(opts *bind.CallOpts, _participant common.Address, _deliveryDate *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "getPositionsByParticipantDeliveryDate", _participant, _deliveryDate)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPositionsByParticipantDeliveryDate is a free data retrieval call binding the contract method 0x207c1f89.
//
// Solidity: function getPositionsByParticipantDeliveryDate(address _participant, uint256 _deliveryDate) view returns(bytes32[])
func (_Futures *FuturesSession) GetPositionsByParticipantDeliveryDate(_participant common.Address, _deliveryDate *big.Int) ([][32]byte, error) {
	return _Futures.Contract.GetPositionsByParticipantDeliveryDate(&_Futures.CallOpts, _participant, _deliveryDate)
}

// GetPositionsByParticipantDeliveryDate is a free data retrieval call binding the contract method 0x207c1f89.
//
// Solidity: function getPositionsByParticipantDeliveryDate(address _participant, uint256 _deliveryDate) view returns(bytes32[])
func (_Futures *FuturesCallerSession) GetPositionsByParticipantDeliveryDate(_participant common.Address, _deliveryDate *big.Int) ([][32]byte, error) {
	return _Futures.Contract.GetPositionsByParticipantDeliveryDate(&_Futures.CallOpts, _participant, _deliveryDate)
}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Futures *FuturesCaller) HashrateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "hashrateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Futures *FuturesSession) HashrateOracle() (common.Address, error) {
	return _Futures.Contract.HashrateOracle(&_Futures.CallOpts)
}

// HashrateOracle is a free data retrieval call binding the contract method 0xf87c68b9.
//
// Solidity: function hashrateOracle() view returns(address)
func (_Futures *FuturesCallerSession) HashrateOracle() (common.Address, error) {
	return _Futures.Contract.HashrateOracle(&_Futures.CallOpts)
}

// LiquidationMarginPercent is a free data retrieval call binding the contract method 0xc45c8418.
//
// Solidity: function liquidationMarginPercent() view returns(uint8)
func (_Futures *FuturesCaller) LiquidationMarginPercent(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "liquidationMarginPercent")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// LiquidationMarginPercent is a free data retrieval call binding the contract method 0xc45c8418.
//
// Solidity: function liquidationMarginPercent() view returns(uint8)
func (_Futures *FuturesSession) LiquidationMarginPercent() (uint8, error) {
	return _Futures.Contract.LiquidationMarginPercent(&_Futures.CallOpts)
}

// LiquidationMarginPercent is a free data retrieval call binding the contract method 0xc45c8418.
//
// Solidity: function liquidationMarginPercent() view returns(uint8)
func (_Futures *FuturesCallerSession) LiquidationMarginPercent() (uint8, error) {
	return _Futures.Contract.LiquidationMarginPercent(&_Futures.CallOpts)
}

// MinimumPriceIncrement is a free data retrieval call binding the contract method 0xdb341765.
//
// Solidity: function minimumPriceIncrement() view returns(uint256)
func (_Futures *FuturesCaller) MinimumPriceIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "minimumPriceIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumPriceIncrement is a free data retrieval call binding the contract method 0xdb341765.
//
// Solidity: function minimumPriceIncrement() view returns(uint256)
func (_Futures *FuturesSession) MinimumPriceIncrement() (*big.Int, error) {
	return _Futures.Contract.MinimumPriceIncrement(&_Futures.CallOpts)
}

// MinimumPriceIncrement is a free data retrieval call binding the contract method 0xdb341765.
//
// Solidity: function minimumPriceIncrement() view returns(uint256)
func (_Futures *FuturesCallerSession) MinimumPriceIncrement() (*big.Int, error) {
	return _Futures.Contract.MinimumPriceIncrement(&_Futures.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Futures *FuturesCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Futures *FuturesSession) Name() (string, error) {
	return _Futures.Contract.Name(&_Futures.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Futures *FuturesCallerSession) Name() (string, error) {
	return _Futures.Contract.Name(&_Futures.CallOpts)
}

// OrderFee is a free data retrieval call binding the contract method 0x1392fb3e.
//
// Solidity: function orderFee() view returns(uint256)
func (_Futures *FuturesCaller) OrderFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "orderFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderFee is a free data retrieval call binding the contract method 0x1392fb3e.
//
// Solidity: function orderFee() view returns(uint256)
func (_Futures *FuturesSession) OrderFee() (*big.Int, error) {
	return _Futures.Contract.OrderFee(&_Futures.CallOpts)
}

// OrderFee is a free data retrieval call binding the contract method 0x1392fb3e.
//
// Solidity: function orderFee() view returns(uint256)
func (_Futures *FuturesCallerSession) OrderFee() (*big.Int, error) {
	return _Futures.Contract.OrderFee(&_Futures.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Futures *FuturesCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Futures *FuturesSession) Owner() (common.Address, error) {
	return _Futures.Contract.Owner(&_Futures.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Futures *FuturesCallerSession) Owner() (common.Address, error) {
	return _Futures.Contract.Owner(&_Futures.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Futures *FuturesCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Futures *FuturesSession) ProxiableUUID() ([32]byte, error) {
	return _Futures.Contract.ProxiableUUID(&_Futures.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Futures *FuturesCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Futures.Contract.ProxiableUUID(&_Futures.CallOpts)
}

// SpeedHps is a free data retrieval call binding the contract method 0x44d4550d.
//
// Solidity: function speedHps() view returns(uint256)
func (_Futures *FuturesCaller) SpeedHps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "speedHps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpeedHps is a free data retrieval call binding the contract method 0x44d4550d.
//
// Solidity: function speedHps() view returns(uint256)
func (_Futures *FuturesSession) SpeedHps() (*big.Int, error) {
	return _Futures.Contract.SpeedHps(&_Futures.CallOpts)
}

// SpeedHps is a free data retrieval call binding the contract method 0x44d4550d.
//
// Solidity: function speedHps() view returns(uint256)
func (_Futures *FuturesCallerSession) SpeedHps() (*big.Int, error) {
	return _Futures.Contract.SpeedHps(&_Futures.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Futures *FuturesCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Futures *FuturesSession) Symbol() (string, error) {
	return _Futures.Contract.Symbol(&_Futures.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Futures *FuturesCallerSession) Symbol() (string, error) {
	return _Futures.Contract.Symbol(&_Futures.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Futures *FuturesCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Futures *FuturesSession) Token() (common.Address, error) {
	return _Futures.Contract.Token(&_Futures.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Futures *FuturesCallerSession) Token() (common.Address, error) {
	return _Futures.Contract.Token(&_Futures.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Futures *FuturesCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Futures *FuturesSession) TotalSupply() (*big.Int, error) {
	return _Futures.Contract.TotalSupply(&_Futures.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Futures *FuturesCallerSession) TotalSupply() (*big.Int, error) {
	return _Futures.Contract.TotalSupply(&_Futures.CallOpts)
}

// ValidatorAddress is a free data retrieval call binding the contract method 0x3fe4676e.
//
// Solidity: function validatorAddress() view returns(address)
func (_Futures *FuturesCaller) ValidatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Futures.contract.Call(opts, &out, "validatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorAddress is a free data retrieval call binding the contract method 0x3fe4676e.
//
// Solidity: function validatorAddress() view returns(address)
func (_Futures *FuturesSession) ValidatorAddress() (common.Address, error) {
	return _Futures.Contract.ValidatorAddress(&_Futures.CallOpts)
}

// ValidatorAddress is a free data retrieval call binding the contract method 0x3fe4676e.
//
// Solidity: function validatorAddress() view returns(address)
func (_Futures *FuturesCallerSession) ValidatorAddress() (common.Address, error) {
	return _Futures.Contract.ValidatorAddress(&_Futures.CallOpts)
}

// AddMargin is a paid mutator transaction binding the contract method 0xa43be948.
//
// Solidity: function addMargin(uint256 _amount) returns()
func (_Futures *FuturesTransactor) AddMargin(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "addMargin", _amount)
}

// AddMargin is a paid mutator transaction binding the contract method 0xa43be948.
//
// Solidity: function addMargin(uint256 _amount) returns()
func (_Futures *FuturesSession) AddMargin(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.AddMargin(&_Futures.TransactOpts, _amount)
}

// AddMargin is a paid mutator transaction binding the contract method 0xa43be948.
//
// Solidity: function addMargin(uint256 _amount) returns()
func (_Futures *FuturesTransactorSession) AddMargin(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.AddMargin(&_Futures.TransactOpts, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Futures *FuturesTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Futures *FuturesSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Approve(&_Futures.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Futures *FuturesTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Approve(&_Futures.TransactOpts, spender, value)
}

// CloseDelivery is a paid mutator transaction binding the contract method 0xdbf8525f.
//
// Solidity: function closeDelivery(bytes32 _positionId, bool _blameSeller) returns()
func (_Futures *FuturesTransactor) CloseDelivery(opts *bind.TransactOpts, _positionId [32]byte, _blameSeller bool) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "closeDelivery", _positionId, _blameSeller)
}

// CloseDelivery is a paid mutator transaction binding the contract method 0xdbf8525f.
//
// Solidity: function closeDelivery(bytes32 _positionId, bool _blameSeller) returns()
func (_Futures *FuturesSession) CloseDelivery(_positionId [32]byte, _blameSeller bool) (*types.Transaction, error) {
	return _Futures.Contract.CloseDelivery(&_Futures.TransactOpts, _positionId, _blameSeller)
}

// CloseDelivery is a paid mutator transaction binding the contract method 0xdbf8525f.
//
// Solidity: function closeDelivery(bytes32 _positionId, bool _blameSeller) returns()
func (_Futures *FuturesTransactorSession) CloseDelivery(_positionId [32]byte, _blameSeller bool) (*types.Transaction, error) {
	return _Futures.Contract.CloseDelivery(&_Futures.TransactOpts, _positionId, _blameSeller)
}

// CloseOrder is a paid mutator transaction binding the contract method 0x3bed6b95.
//
// Solidity: function closeOrder(bytes32 _orderId) returns()
func (_Futures *FuturesTransactor) CloseOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "closeOrder", _orderId)
}

// CloseOrder is a paid mutator transaction binding the contract method 0x3bed6b95.
//
// Solidity: function closeOrder(bytes32 _orderId) returns()
func (_Futures *FuturesSession) CloseOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _Futures.Contract.CloseOrder(&_Futures.TransactOpts, _orderId)
}

// CloseOrder is a paid mutator transaction binding the contract method 0x3bed6b95.
//
// Solidity: function closeOrder(bytes32 _orderId) returns()
func (_Futures *FuturesTransactorSession) CloseOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _Futures.Contract.CloseOrder(&_Futures.TransactOpts, _orderId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x6828a054.
//
// Solidity: function createOrder(uint256 _price, uint256 _deliveryDate, string _destURL, int8 _qty) returns()
func (_Futures *FuturesTransactor) CreateOrder(opts *bind.TransactOpts, _price *big.Int, _deliveryDate *big.Int, _destURL string, _qty int8) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "createOrder", _price, _deliveryDate, _destURL, _qty)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x6828a054.
//
// Solidity: function createOrder(uint256 _price, uint256 _deliveryDate, string _destURL, int8 _qty) returns()
func (_Futures *FuturesSession) CreateOrder(_price *big.Int, _deliveryDate *big.Int, _destURL string, _qty int8) (*types.Transaction, error) {
	return _Futures.Contract.CreateOrder(&_Futures.TransactOpts, _price, _deliveryDate, _destURL, _qty)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x6828a054.
//
// Solidity: function createOrder(uint256 _price, uint256 _deliveryDate, string _destURL, int8 _qty) returns()
func (_Futures *FuturesTransactorSession) CreateOrder(_price *big.Int, _deliveryDate *big.Int, _destURL string, _qty int8) (*types.Transaction, error) {
	return _Futures.Contract.CreateOrder(&_Futures.TransactOpts, _price, _deliveryDate, _destURL, _qty)
}

// DepositDeliveryPayment is a paid mutator transaction binding the contract method 0xe4dc72ba.
//
// Solidity: function depositDeliveryPayment(bytes32[] _positionIds) returns()
func (_Futures *FuturesTransactor) DepositDeliveryPayment(opts *bind.TransactOpts, _positionIds [][32]byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "depositDeliveryPayment", _positionIds)
}

// DepositDeliveryPayment is a paid mutator transaction binding the contract method 0xe4dc72ba.
//
// Solidity: function depositDeliveryPayment(bytes32[] _positionIds) returns()
func (_Futures *FuturesSession) DepositDeliveryPayment(_positionIds [][32]byte) (*types.Transaction, error) {
	return _Futures.Contract.DepositDeliveryPayment(&_Futures.TransactOpts, _positionIds)
}

// DepositDeliveryPayment is a paid mutator transaction binding the contract method 0xe4dc72ba.
//
// Solidity: function depositDeliveryPayment(bytes32[] _positionIds) returns()
func (_Futures *FuturesTransactorSession) DepositDeliveryPayment(_positionIds [][32]byte) (*types.Transaction, error) {
	return _Futures.Contract.DepositDeliveryPayment(&_Futures.TransactOpts, _positionIds)
}

// DepositDeliveryPayment0 is a paid mutator transaction binding the contract method 0xe947cc3e.
//
// Solidity: function depositDeliveryPayment(uint256 _amount, uint256 _deliveryDate) returns(bool)
func (_Futures *FuturesTransactor) DepositDeliveryPayment0(opts *bind.TransactOpts, _amount *big.Int, _deliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "depositDeliveryPayment0", _amount, _deliveryDate)
}

// DepositDeliveryPayment0 is a paid mutator transaction binding the contract method 0xe947cc3e.
//
// Solidity: function depositDeliveryPayment(uint256 _amount, uint256 _deliveryDate) returns(bool)
func (_Futures *FuturesSession) DepositDeliveryPayment0(_amount *big.Int, _deliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositDeliveryPayment0(&_Futures.TransactOpts, _amount, _deliveryDate)
}

// DepositDeliveryPayment0 is a paid mutator transaction binding the contract method 0xe947cc3e.
//
// Solidity: function depositDeliveryPayment(uint256 _amount, uint256 _deliveryDate) returns(bool)
func (_Futures *FuturesTransactorSession) DepositDeliveryPayment0(_amount *big.Int, _deliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositDeliveryPayment0(&_Futures.TransactOpts, _amount, _deliveryDate)
}

// DepositReservePool is a paid mutator transaction binding the contract method 0x144da1cf.
//
// Solidity: function depositReservePool(uint256 _amount) returns()
func (_Futures *FuturesTransactor) DepositReservePool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "depositReservePool", _amount)
}

// DepositReservePool is a paid mutator transaction binding the contract method 0x144da1cf.
//
// Solidity: function depositReservePool(uint256 _amount) returns()
func (_Futures *FuturesSession) DepositReservePool(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositReservePool(&_Futures.TransactOpts, _amount)
}

// DepositReservePool is a paid mutator transaction binding the contract method 0x144da1cf.
//
// Solidity: function depositReservePool(uint256 _amount) returns()
func (_Futures *FuturesTransactorSession) DepositReservePool(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.DepositReservePool(&_Futures.TransactOpts, _amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xaca23867.
//
// Solidity: function initialize(address _token, address _hashrateOracle, address _validatorAddress, uint8 _liquidationMarginPercent, uint256 _speedHps, uint256 _minimumPriceIncrement, uint8 _deliveryDurationDays, uint8 _deliveryIntervalDays, uint8 _futureDeliveryDatesCount, uint256 _firstFutureDeliveryDate) returns()
func (_Futures *FuturesTransactor) Initialize(opts *bind.TransactOpts, _token common.Address, _hashrateOracle common.Address, _validatorAddress common.Address, _liquidationMarginPercent uint8, _speedHps *big.Int, _minimumPriceIncrement *big.Int, _deliveryDurationDays uint8, _deliveryIntervalDays uint8, _futureDeliveryDatesCount uint8, _firstFutureDeliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "initialize", _token, _hashrateOracle, _validatorAddress, _liquidationMarginPercent, _speedHps, _minimumPriceIncrement, _deliveryDurationDays, _deliveryIntervalDays, _futureDeliveryDatesCount, _firstFutureDeliveryDate)
}

// Initialize is a paid mutator transaction binding the contract method 0xaca23867.
//
// Solidity: function initialize(address _token, address _hashrateOracle, address _validatorAddress, uint8 _liquidationMarginPercent, uint256 _speedHps, uint256 _minimumPriceIncrement, uint8 _deliveryDurationDays, uint8 _deliveryIntervalDays, uint8 _futureDeliveryDatesCount, uint256 _firstFutureDeliveryDate) returns()
func (_Futures *FuturesSession) Initialize(_token common.Address, _hashrateOracle common.Address, _validatorAddress common.Address, _liquidationMarginPercent uint8, _speedHps *big.Int, _minimumPriceIncrement *big.Int, _deliveryDurationDays uint8, _deliveryIntervalDays uint8, _futureDeliveryDatesCount uint8, _firstFutureDeliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Initialize(&_Futures.TransactOpts, _token, _hashrateOracle, _validatorAddress, _liquidationMarginPercent, _speedHps, _minimumPriceIncrement, _deliveryDurationDays, _deliveryIntervalDays, _futureDeliveryDatesCount, _firstFutureDeliveryDate)
}

// Initialize is a paid mutator transaction binding the contract method 0xaca23867.
//
// Solidity: function initialize(address _token, address _hashrateOracle, address _validatorAddress, uint8 _liquidationMarginPercent, uint256 _speedHps, uint256 _minimumPriceIncrement, uint8 _deliveryDurationDays, uint8 _deliveryIntervalDays, uint8 _futureDeliveryDatesCount, uint256 _firstFutureDeliveryDate) returns()
func (_Futures *FuturesTransactorSession) Initialize(_token common.Address, _hashrateOracle common.Address, _validatorAddress common.Address, _liquidationMarginPercent uint8, _speedHps *big.Int, _minimumPriceIncrement *big.Int, _deliveryDurationDays uint8, _deliveryIntervalDays uint8, _futureDeliveryDatesCount uint8, _firstFutureDeliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Initialize(&_Futures.TransactOpts, _token, _hashrateOracle, _validatorAddress, _liquidationMarginPercent, _speedHps, _minimumPriceIncrement, _deliveryDurationDays, _deliveryIntervalDays, _futureDeliveryDatesCount, _firstFutureDeliveryDate)
}

// MarginCall is a paid mutator transaction binding the contract method 0x846bf07c.
//
// Solidity: function marginCall(address _participant) returns()
func (_Futures *FuturesTransactor) MarginCall(opts *bind.TransactOpts, _participant common.Address) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "marginCall", _participant)
}

// MarginCall is a paid mutator transaction binding the contract method 0x846bf07c.
//
// Solidity: function marginCall(address _participant) returns()
func (_Futures *FuturesSession) MarginCall(_participant common.Address) (*types.Transaction, error) {
	return _Futures.Contract.MarginCall(&_Futures.TransactOpts, _participant)
}

// MarginCall is a paid mutator transaction binding the contract method 0x846bf07c.
//
// Solidity: function marginCall(address _participant) returns()
func (_Futures *FuturesTransactorSession) MarginCall(_participant common.Address) (*types.Transaction, error) {
	return _Futures.Contract.MarginCall(&_Futures.TransactOpts, _participant)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Futures *FuturesTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Futures *FuturesSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Futures.Contract.Multicall(&_Futures.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Futures *FuturesTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Futures.Contract.Multicall(&_Futures.TransactOpts, data)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xf11f854f.
//
// Solidity: function removeMargin(uint256 _amount) returns()
func (_Futures *FuturesTransactor) RemoveMargin(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "removeMargin", _amount)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xf11f854f.
//
// Solidity: function removeMargin(uint256 _amount) returns()
func (_Futures *FuturesSession) RemoveMargin(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.RemoveMargin(&_Futures.TransactOpts, _amount)
}

// RemoveMargin is a paid mutator transaction binding the contract method 0xf11f854f.
//
// Solidity: function removeMargin(uint256 _amount) returns()
func (_Futures *FuturesTransactorSession) RemoveMargin(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.RemoveMargin(&_Futures.TransactOpts, _amount)
}

// RemoveOutdatedOrdersForParticipant is a paid mutator transaction binding the contract method 0xb129246a.
//
// Solidity: function removeOutdatedOrdersForParticipant(address _participant) returns(uint256 count)
func (_Futures *FuturesTransactor) RemoveOutdatedOrdersForParticipant(opts *bind.TransactOpts, _participant common.Address) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "removeOutdatedOrdersForParticipant", _participant)
}

// RemoveOutdatedOrdersForParticipant is a paid mutator transaction binding the contract method 0xb129246a.
//
// Solidity: function removeOutdatedOrdersForParticipant(address _participant) returns(uint256 count)
func (_Futures *FuturesSession) RemoveOutdatedOrdersForParticipant(_participant common.Address) (*types.Transaction, error) {
	return _Futures.Contract.RemoveOutdatedOrdersForParticipant(&_Futures.TransactOpts, _participant)
}

// RemoveOutdatedOrdersForParticipant is a paid mutator transaction binding the contract method 0xb129246a.
//
// Solidity: function removeOutdatedOrdersForParticipant(address _participant) returns(uint256 count)
func (_Futures *FuturesTransactorSession) RemoveOutdatedOrdersForParticipant(_participant common.Address) (*types.Transaction, error) {
	return _Futures.Contract.RemoveOutdatedOrdersForParticipant(&_Futures.TransactOpts, _participant)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Futures *FuturesTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Futures *FuturesSession) RenounceOwnership() (*types.Transaction, error) {
	return _Futures.Contract.RenounceOwnership(&_Futures.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Futures *FuturesTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Futures.Contract.RenounceOwnership(&_Futures.TransactOpts)
}

// SetBreachPenaltyRatePerDay is a paid mutator transaction binding the contract method 0x96e6fc18.
//
// Solidity: function setBreachPenaltyRatePerDay(uint256 _breachPenaltyRatePerDay) returns()
func (_Futures *FuturesTransactor) SetBreachPenaltyRatePerDay(opts *bind.TransactOpts, _breachPenaltyRatePerDay *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "setBreachPenaltyRatePerDay", _breachPenaltyRatePerDay)
}

// SetBreachPenaltyRatePerDay is a paid mutator transaction binding the contract method 0x96e6fc18.
//
// Solidity: function setBreachPenaltyRatePerDay(uint256 _breachPenaltyRatePerDay) returns()
func (_Futures *FuturesSession) SetBreachPenaltyRatePerDay(_breachPenaltyRatePerDay *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.SetBreachPenaltyRatePerDay(&_Futures.TransactOpts, _breachPenaltyRatePerDay)
}

// SetBreachPenaltyRatePerDay is a paid mutator transaction binding the contract method 0x96e6fc18.
//
// Solidity: function setBreachPenaltyRatePerDay(uint256 _breachPenaltyRatePerDay) returns()
func (_Futures *FuturesTransactorSession) SetBreachPenaltyRatePerDay(_breachPenaltyRatePerDay *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.SetBreachPenaltyRatePerDay(&_Futures.TransactOpts, _breachPenaltyRatePerDay)
}

// SetFutureDeliveryDatesCount is a paid mutator transaction binding the contract method 0xff157997.
//
// Solidity: function setFutureDeliveryDatesCount(uint8 _futureDeliveryDatesCount) returns()
func (_Futures *FuturesTransactor) SetFutureDeliveryDatesCount(opts *bind.TransactOpts, _futureDeliveryDatesCount uint8) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "setFutureDeliveryDatesCount", _futureDeliveryDatesCount)
}

// SetFutureDeliveryDatesCount is a paid mutator transaction binding the contract method 0xff157997.
//
// Solidity: function setFutureDeliveryDatesCount(uint8 _futureDeliveryDatesCount) returns()
func (_Futures *FuturesSession) SetFutureDeliveryDatesCount(_futureDeliveryDatesCount uint8) (*types.Transaction, error) {
	return _Futures.Contract.SetFutureDeliveryDatesCount(&_Futures.TransactOpts, _futureDeliveryDatesCount)
}

// SetFutureDeliveryDatesCount is a paid mutator transaction binding the contract method 0xff157997.
//
// Solidity: function setFutureDeliveryDatesCount(uint8 _futureDeliveryDatesCount) returns()
func (_Futures *FuturesTransactorSession) SetFutureDeliveryDatesCount(_futureDeliveryDatesCount uint8) (*types.Transaction, error) {
	return _Futures.Contract.SetFutureDeliveryDatesCount(&_Futures.TransactOpts, _futureDeliveryDatesCount)
}

// SetLiquidationMarginPercent is a paid mutator transaction binding the contract method 0x1ea423bb.
//
// Solidity: function setLiquidationMarginPercent(uint8 _liquidationMarginPercent) returns()
func (_Futures *FuturesTransactor) SetLiquidationMarginPercent(opts *bind.TransactOpts, _liquidationMarginPercent uint8) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "setLiquidationMarginPercent", _liquidationMarginPercent)
}

// SetLiquidationMarginPercent is a paid mutator transaction binding the contract method 0x1ea423bb.
//
// Solidity: function setLiquidationMarginPercent(uint8 _liquidationMarginPercent) returns()
func (_Futures *FuturesSession) SetLiquidationMarginPercent(_liquidationMarginPercent uint8) (*types.Transaction, error) {
	return _Futures.Contract.SetLiquidationMarginPercent(&_Futures.TransactOpts, _liquidationMarginPercent)
}

// SetLiquidationMarginPercent is a paid mutator transaction binding the contract method 0x1ea423bb.
//
// Solidity: function setLiquidationMarginPercent(uint8 _liquidationMarginPercent) returns()
func (_Futures *FuturesTransactorSession) SetLiquidationMarginPercent(_liquidationMarginPercent uint8) (*types.Transaction, error) {
	return _Futures.Contract.SetLiquidationMarginPercent(&_Futures.TransactOpts, _liquidationMarginPercent)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address addr) returns()
func (_Futures *FuturesTransactor) SetOracle(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "setOracle", addr)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address addr) returns()
func (_Futures *FuturesSession) SetOracle(addr common.Address) (*types.Transaction, error) {
	return _Futures.Contract.SetOracle(&_Futures.TransactOpts, addr)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address addr) returns()
func (_Futures *FuturesTransactorSession) SetOracle(addr common.Address) (*types.Transaction, error) {
	return _Futures.Contract.SetOracle(&_Futures.TransactOpts, addr)
}

// SetOrderFee is a paid mutator transaction binding the contract method 0xc014930f.
//
// Solidity: function setOrderFee(uint256 _orderFee) returns()
func (_Futures *FuturesTransactor) SetOrderFee(opts *bind.TransactOpts, _orderFee *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "setOrderFee", _orderFee)
}

// SetOrderFee is a paid mutator transaction binding the contract method 0xc014930f.
//
// Solidity: function setOrderFee(uint256 _orderFee) returns()
func (_Futures *FuturesSession) SetOrderFee(_orderFee *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.SetOrderFee(&_Futures.TransactOpts, _orderFee)
}

// SetOrderFee is a paid mutator transaction binding the contract method 0xc014930f.
//
// Solidity: function setOrderFee(uint256 _orderFee) returns()
func (_Futures *FuturesTransactorSession) SetOrderFee(_orderFee *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.SetOrderFee(&_Futures.TransactOpts, _orderFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Futures *FuturesTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Futures *FuturesSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Transfer(&_Futures.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_Futures *FuturesTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.Transfer(&_Futures.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Futures *FuturesTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Futures *FuturesSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.TransferFrom(&_Futures.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_Futures *FuturesTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.TransferFrom(&_Futures.TransactOpts, _from, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Futures *FuturesTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Futures *FuturesSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Futures.Contract.TransferOwnership(&_Futures.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Futures *FuturesTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Futures.Contract.TransferOwnership(&_Futures.TransactOpts, newOwner)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Futures *FuturesTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Futures *FuturesSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Futures.Contract.UpgradeToAndCall(&_Futures.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Futures *FuturesTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Futures.Contract.UpgradeToAndCall(&_Futures.TransactOpts, newImplementation, data)
}

// WithdrawDeliveryPayment is a paid mutator transaction binding the contract method 0x1504ae82.
//
// Solidity: function withdrawDeliveryPayment(uint256 _deliveryDate) returns()
func (_Futures *FuturesTransactor) WithdrawDeliveryPayment(opts *bind.TransactOpts, _deliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "withdrawDeliveryPayment", _deliveryDate)
}

// WithdrawDeliveryPayment is a paid mutator transaction binding the contract method 0x1504ae82.
//
// Solidity: function withdrawDeliveryPayment(uint256 _deliveryDate) returns()
func (_Futures *FuturesSession) WithdrawDeliveryPayment(_deliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.WithdrawDeliveryPayment(&_Futures.TransactOpts, _deliveryDate)
}

// WithdrawDeliveryPayment is a paid mutator transaction binding the contract method 0x1504ae82.
//
// Solidity: function withdrawDeliveryPayment(uint256 _deliveryDate) returns()
func (_Futures *FuturesTransactorSession) WithdrawDeliveryPayment(_deliveryDate *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.WithdrawDeliveryPayment(&_Futures.TransactOpts, _deliveryDate)
}

// WithdrawReservePool is a paid mutator transaction binding the contract method 0xa3ccb8ba.
//
// Solidity: function withdrawReservePool(uint256 _amount) returns()
func (_Futures *FuturesTransactor) WithdrawReservePool(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Futures.contract.Transact(opts, "withdrawReservePool", _amount)
}

// WithdrawReservePool is a paid mutator transaction binding the contract method 0xa3ccb8ba.
//
// Solidity: function withdrawReservePool(uint256 _amount) returns()
func (_Futures *FuturesSession) WithdrawReservePool(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.WithdrawReservePool(&_Futures.TransactOpts, _amount)
}

// WithdrawReservePool is a paid mutator transaction binding the contract method 0xa3ccb8ba.
//
// Solidity: function withdrawReservePool(uint256 _amount) returns()
func (_Futures *FuturesTransactorSession) WithdrawReservePool(_amount *big.Int) (*types.Transaction, error) {
	return _Futures.Contract.WithdrawReservePool(&_Futures.TransactOpts, _amount)
}

// FuturesApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Futures contract.
type FuturesApprovalIterator struct {
	Event *FuturesApproval // Event containing the contract specifics and raw log

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
func (it *FuturesApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesApproval)
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
		it.Event = new(FuturesApproval)
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
func (it *FuturesApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesApproval represents a Approval event raised by the Futures contract.
type FuturesApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Futures *FuturesFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FuturesApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FuturesApprovalIterator{contract: _Futures.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Futures *FuturesFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FuturesApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesApproval)
				if err := _Futures.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Futures *FuturesFilterer) ParseApproval(log types.Log) (*FuturesApproval, error) {
	event := new(FuturesApproval)
	if err := _Futures.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Futures contract.
type FuturesInitializedIterator struct {
	Event *FuturesInitialized // Event containing the contract specifics and raw log

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
func (it *FuturesInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesInitialized)
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
		it.Event = new(FuturesInitialized)
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
func (it *FuturesInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesInitialized represents a Initialized event raised by the Futures contract.
type FuturesInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Futures *FuturesFilterer) FilterInitialized(opts *bind.FilterOpts) (*FuturesInitializedIterator, error) {

	logs, sub, err := _Futures.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FuturesInitializedIterator{contract: _Futures.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Futures *FuturesFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FuturesInitialized) (event.Subscription, error) {

	logs, sub, err := _Futures.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesInitialized)
				if err := _Futures.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Futures *FuturesFilterer) ParseInitialized(log types.Log) (*FuturesInitialized, error) {
	event := new(FuturesInitialized)
	if err := _Futures.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesOrderClosedIterator is returned from FilterOrderClosed and is used to iterate over the raw logs and unpacked data for OrderClosed events raised by the Futures contract.
type FuturesOrderClosedIterator struct {
	Event *FuturesOrderClosed // Event containing the contract specifics and raw log

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
func (it *FuturesOrderClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesOrderClosed)
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
		it.Event = new(FuturesOrderClosed)
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
func (it *FuturesOrderClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesOrderClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesOrderClosed represents a OrderClosed event raised by the Futures contract.
type FuturesOrderClosed struct {
	OrderId     [32]byte
	Participant common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderClosed is a free log retrieval operation binding the contract event 0xba23b3f42d60d00e8a99f8faa964276a8b5eb6b1088f9f2d1ea3482c95654fe6.
//
// Solidity: event OrderClosed(bytes32 indexed orderId, address indexed participant)
func (_Futures *FuturesFilterer) FilterOrderClosed(opts *bind.FilterOpts, orderId [][32]byte, participant []common.Address) (*FuturesOrderClosedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "OrderClosed", orderIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &FuturesOrderClosedIterator{contract: _Futures.contract, event: "OrderClosed", logs: logs, sub: sub}, nil
}

// WatchOrderClosed is a free log subscription operation binding the contract event 0xba23b3f42d60d00e8a99f8faa964276a8b5eb6b1088f9f2d1ea3482c95654fe6.
//
// Solidity: event OrderClosed(bytes32 indexed orderId, address indexed participant)
func (_Futures *FuturesFilterer) WatchOrderClosed(opts *bind.WatchOpts, sink chan<- *FuturesOrderClosed, orderId [][32]byte, participant []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "OrderClosed", orderIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesOrderClosed)
				if err := _Futures.contract.UnpackLog(event, "OrderClosed", log); err != nil {
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

// ParseOrderClosed is a log parse operation binding the contract event 0xba23b3f42d60d00e8a99f8faa964276a8b5eb6b1088f9f2d1ea3482c95654fe6.
//
// Solidity: event OrderClosed(bytes32 indexed orderId, address indexed participant)
func (_Futures *FuturesFilterer) ParseOrderClosed(log types.Log) (*FuturesOrderClosed, error) {
	event := new(FuturesOrderClosed)
	if err := _Futures.contract.UnpackLog(event, "OrderClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the Futures contract.
type FuturesOrderCreatedIterator struct {
	Event *FuturesOrderCreated // Event containing the contract specifics and raw log

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
func (it *FuturesOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesOrderCreated)
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
		it.Event = new(FuturesOrderCreated)
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
func (it *FuturesOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesOrderCreated represents a OrderCreated event raised by the Futures contract.
type FuturesOrderCreated struct {
	OrderId     [32]byte
	Participant common.Address
	DestURL     string
	PricePerDay *big.Int
	DeliveryAt  *big.Int
	IsBuy       bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x1f52a6f4a2d2a66b497ba87509c3bf307f623f437d154026f26716ed2d496d3b.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, address indexed participant, string destURL, uint256 pricePerDay, uint256 deliveryAt, bool isBuy)
func (_Futures *FuturesFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderId [][32]byte, participant []common.Address) (*FuturesOrderCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "OrderCreated", orderIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &FuturesOrderCreatedIterator{contract: _Futures.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x1f52a6f4a2d2a66b497ba87509c3bf307f623f437d154026f26716ed2d496d3b.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, address indexed participant, string destURL, uint256 pricePerDay, uint256 deliveryAt, bool isBuy)
func (_Futures *FuturesFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *FuturesOrderCreated, orderId [][32]byte, participant []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "OrderCreated", orderIdRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesOrderCreated)
				if err := _Futures.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x1f52a6f4a2d2a66b497ba87509c3bf307f623f437d154026f26716ed2d496d3b.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, address indexed participant, string destURL, uint256 pricePerDay, uint256 deliveryAt, bool isBuy)
func (_Futures *FuturesFilterer) ParseOrderCreated(log types.Log) (*FuturesOrderCreated, error) {
	event := new(FuturesOrderCreated)
	if err := _Futures.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesOrderFeeUpdatedIterator is returned from FilterOrderFeeUpdated and is used to iterate over the raw logs and unpacked data for OrderFeeUpdated events raised by the Futures contract.
type FuturesOrderFeeUpdatedIterator struct {
	Event *FuturesOrderFeeUpdated // Event containing the contract specifics and raw log

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
func (it *FuturesOrderFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesOrderFeeUpdated)
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
		it.Event = new(FuturesOrderFeeUpdated)
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
func (it *FuturesOrderFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesOrderFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesOrderFeeUpdated represents a OrderFeeUpdated event raised by the Futures contract.
type FuturesOrderFeeUpdated struct {
	OrderFee *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrderFeeUpdated is a free log retrieval operation binding the contract event 0xafb0bb88fecd164b1de860279e72dadcc296d3b91d41cef931856f3416cdd1db.
//
// Solidity: event OrderFeeUpdated(uint256 orderFee)
func (_Futures *FuturesFilterer) FilterOrderFeeUpdated(opts *bind.FilterOpts) (*FuturesOrderFeeUpdatedIterator, error) {

	logs, sub, err := _Futures.contract.FilterLogs(opts, "OrderFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &FuturesOrderFeeUpdatedIterator{contract: _Futures.contract, event: "OrderFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchOrderFeeUpdated is a free log subscription operation binding the contract event 0xafb0bb88fecd164b1de860279e72dadcc296d3b91d41cef931856f3416cdd1db.
//
// Solidity: event OrderFeeUpdated(uint256 orderFee)
func (_Futures *FuturesFilterer) WatchOrderFeeUpdated(opts *bind.WatchOpts, sink chan<- *FuturesOrderFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Futures.contract.WatchLogs(opts, "OrderFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesOrderFeeUpdated)
				if err := _Futures.contract.UnpackLog(event, "OrderFeeUpdated", log); err != nil {
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

// ParseOrderFeeUpdated is a log parse operation binding the contract event 0xafb0bb88fecd164b1de860279e72dadcc296d3b91d41cef931856f3416cdd1db.
//
// Solidity: event OrderFeeUpdated(uint256 orderFee)
func (_Futures *FuturesFilterer) ParseOrderFeeUpdated(log types.Log) (*FuturesOrderFeeUpdated, error) {
	event := new(FuturesOrderFeeUpdated)
	if err := _Futures.contract.UnpackLog(event, "OrderFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Futures contract.
type FuturesOwnershipTransferredIterator struct {
	Event *FuturesOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FuturesOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesOwnershipTransferred)
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
		it.Event = new(FuturesOwnershipTransferred)
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
func (it *FuturesOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesOwnershipTransferred represents a OwnershipTransferred event raised by the Futures contract.
type FuturesOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Futures *FuturesFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FuturesOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FuturesOwnershipTransferredIterator{contract: _Futures.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Futures *FuturesFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FuturesOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesOwnershipTransferred)
				if err := _Futures.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Futures *FuturesFilterer) ParseOwnershipTransferred(log types.Log) (*FuturesOwnershipTransferred, error) {
	event := new(FuturesOwnershipTransferred)
	if err := _Futures.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesPositionClosedIterator is returned from FilterPositionClosed and is used to iterate over the raw logs and unpacked data for PositionClosed events raised by the Futures contract.
type FuturesPositionClosedIterator struct {
	Event *FuturesPositionClosed // Event containing the contract specifics and raw log

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
func (it *FuturesPositionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPositionClosed)
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
		it.Event = new(FuturesPositionClosed)
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
func (it *FuturesPositionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPositionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPositionClosed represents a PositionClosed event raised by the Futures contract.
type FuturesPositionClosed struct {
	PositionId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPositionClosed is a free log retrieval operation binding the contract event 0xb77a99aa6bf8135a033a72ab89373f6f8977dc081d277755b7622c1782a14799.
//
// Solidity: event PositionClosed(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) FilterPositionClosed(opts *bind.FilterOpts, positionId [][32]byte) (*FuturesPositionClosedIterator, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "PositionClosed", positionIdRule)
	if err != nil {
		return nil, err
	}
	return &FuturesPositionClosedIterator{contract: _Futures.contract, event: "PositionClosed", logs: logs, sub: sub}, nil
}

// WatchPositionClosed is a free log subscription operation binding the contract event 0xb77a99aa6bf8135a033a72ab89373f6f8977dc081d277755b7622c1782a14799.
//
// Solidity: event PositionClosed(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) WatchPositionClosed(opts *bind.WatchOpts, sink chan<- *FuturesPositionClosed, positionId [][32]byte) (event.Subscription, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "PositionClosed", positionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPositionClosed)
				if err := _Futures.contract.UnpackLog(event, "PositionClosed", log); err != nil {
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

// ParsePositionClosed is a log parse operation binding the contract event 0xb77a99aa6bf8135a033a72ab89373f6f8977dc081d277755b7622c1782a14799.
//
// Solidity: event PositionClosed(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) ParsePositionClosed(log types.Log) (*FuturesPositionClosed, error) {
	event := new(FuturesPositionClosed)
	if err := _Futures.contract.UnpackLog(event, "PositionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesPositionCreatedIterator is returned from FilterPositionCreated and is used to iterate over the raw logs and unpacked data for PositionCreated events raised by the Futures contract.
type FuturesPositionCreatedIterator struct {
	Event *FuturesPositionCreated // Event containing the contract specifics and raw log

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
func (it *FuturesPositionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPositionCreated)
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
		it.Event = new(FuturesPositionCreated)
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
func (it *FuturesPositionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPositionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPositionCreated represents a PositionCreated event raised by the Futures contract.
type FuturesPositionCreated struct {
	PositionId      [32]byte
	Seller          common.Address
	Buyer           common.Address
	SellPricePerDay *big.Int
	BuyPricePerDay  *big.Int
	DeliveryAt      *big.Int
	DestURL         string
	OrderId         [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPositionCreated is a free log retrieval operation binding the contract event 0x4258e60eecf21b127496b52cfc5b7b5299721db725ba5620a55e2a7c84d43294.
//
// Solidity: event PositionCreated(bytes32 indexed positionId, address indexed seller, address indexed buyer, uint256 sellPricePerDay, uint256 buyPricePerDay, uint256 deliveryAt, string destURL, bytes32 orderId)
func (_Futures *FuturesFilterer) FilterPositionCreated(opts *bind.FilterOpts, positionId [][32]byte, seller []common.Address, buyer []common.Address) (*FuturesPositionCreatedIterator, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "PositionCreated", positionIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &FuturesPositionCreatedIterator{contract: _Futures.contract, event: "PositionCreated", logs: logs, sub: sub}, nil
}

// WatchPositionCreated is a free log subscription operation binding the contract event 0x4258e60eecf21b127496b52cfc5b7b5299721db725ba5620a55e2a7c84d43294.
//
// Solidity: event PositionCreated(bytes32 indexed positionId, address indexed seller, address indexed buyer, uint256 sellPricePerDay, uint256 buyPricePerDay, uint256 deliveryAt, string destURL, bytes32 orderId)
func (_Futures *FuturesFilterer) WatchPositionCreated(opts *bind.WatchOpts, sink chan<- *FuturesPositionCreated, positionId [][32]byte, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "PositionCreated", positionIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPositionCreated)
				if err := _Futures.contract.UnpackLog(event, "PositionCreated", log); err != nil {
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

// ParsePositionCreated is a log parse operation binding the contract event 0x4258e60eecf21b127496b52cfc5b7b5299721db725ba5620a55e2a7c84d43294.
//
// Solidity: event PositionCreated(bytes32 indexed positionId, address indexed seller, address indexed buyer, uint256 sellPricePerDay, uint256 buyPricePerDay, uint256 deliveryAt, string destURL, bytes32 orderId)
func (_Futures *FuturesFilterer) ParsePositionCreated(log types.Log) (*FuturesPositionCreated, error) {
	event := new(FuturesPositionCreated)
	if err := _Futures.contract.UnpackLog(event, "PositionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesPositionDeliveryClosedIterator is returned from FilterPositionDeliveryClosed and is used to iterate over the raw logs and unpacked data for PositionDeliveryClosed events raised by the Futures contract.
type FuturesPositionDeliveryClosedIterator struct {
	Event *FuturesPositionDeliveryClosed // Event containing the contract specifics and raw log

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
func (it *FuturesPositionDeliveryClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPositionDeliveryClosed)
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
		it.Event = new(FuturesPositionDeliveryClosed)
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
func (it *FuturesPositionDeliveryClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPositionDeliveryClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPositionDeliveryClosed represents a PositionDeliveryClosed event raised by the Futures contract.
type FuturesPositionDeliveryClosed struct {
	PositionId [32]byte
	ClosedBy   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPositionDeliveryClosed is a free log retrieval operation binding the contract event 0x4ddc55eb36534e96c19d4ba59a8c3f278889e6277a60101af10cb53f322123a2.
//
// Solidity: event PositionDeliveryClosed(bytes32 indexed positionId, address indexed closedBy)
func (_Futures *FuturesFilterer) FilterPositionDeliveryClosed(opts *bind.FilterOpts, positionId [][32]byte, closedBy []common.Address) (*FuturesPositionDeliveryClosedIterator, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}
	var closedByRule []interface{}
	for _, closedByItem := range closedBy {
		closedByRule = append(closedByRule, closedByItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "PositionDeliveryClosed", positionIdRule, closedByRule)
	if err != nil {
		return nil, err
	}
	return &FuturesPositionDeliveryClosedIterator{contract: _Futures.contract, event: "PositionDeliveryClosed", logs: logs, sub: sub}, nil
}

// WatchPositionDeliveryClosed is a free log subscription operation binding the contract event 0x4ddc55eb36534e96c19d4ba59a8c3f278889e6277a60101af10cb53f322123a2.
//
// Solidity: event PositionDeliveryClosed(bytes32 indexed positionId, address indexed closedBy)
func (_Futures *FuturesFilterer) WatchPositionDeliveryClosed(opts *bind.WatchOpts, sink chan<- *FuturesPositionDeliveryClosed, positionId [][32]byte, closedBy []common.Address) (event.Subscription, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}
	var closedByRule []interface{}
	for _, closedByItem := range closedBy {
		closedByRule = append(closedByRule, closedByItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "PositionDeliveryClosed", positionIdRule, closedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPositionDeliveryClosed)
				if err := _Futures.contract.UnpackLog(event, "PositionDeliveryClosed", log); err != nil {
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

// ParsePositionDeliveryClosed is a log parse operation binding the contract event 0x4ddc55eb36534e96c19d4ba59a8c3f278889e6277a60101af10cb53f322123a2.
//
// Solidity: event PositionDeliveryClosed(bytes32 indexed positionId, address indexed closedBy)
func (_Futures *FuturesFilterer) ParsePositionDeliveryClosed(log types.Log) (*FuturesPositionDeliveryClosed, error) {
	event := new(FuturesPositionDeliveryClosed)
	if err := _Futures.contract.UnpackLog(event, "PositionDeliveryClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesPositionPaidIterator is returned from FilterPositionPaid and is used to iterate over the raw logs and unpacked data for PositionPaid events raised by the Futures contract.
type FuturesPositionPaidIterator struct {
	Event *FuturesPositionPaid // Event containing the contract specifics and raw log

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
func (it *FuturesPositionPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPositionPaid)
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
		it.Event = new(FuturesPositionPaid)
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
func (it *FuturesPositionPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPositionPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPositionPaid represents a PositionPaid event raised by the Futures contract.
type FuturesPositionPaid struct {
	PositionId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPositionPaid is a free log retrieval operation binding the contract event 0x95aac10f3d1c6095ba6d61c88f22bcb31147dd8534a4f0d02b42c2841c271aac.
//
// Solidity: event PositionPaid(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) FilterPositionPaid(opts *bind.FilterOpts, positionId [][32]byte) (*FuturesPositionPaidIterator, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "PositionPaid", positionIdRule)
	if err != nil {
		return nil, err
	}
	return &FuturesPositionPaidIterator{contract: _Futures.contract, event: "PositionPaid", logs: logs, sub: sub}, nil
}

// WatchPositionPaid is a free log subscription operation binding the contract event 0x95aac10f3d1c6095ba6d61c88f22bcb31147dd8534a4f0d02b42c2841c271aac.
//
// Solidity: event PositionPaid(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) WatchPositionPaid(opts *bind.WatchOpts, sink chan<- *FuturesPositionPaid, positionId [][32]byte) (event.Subscription, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "PositionPaid", positionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPositionPaid)
				if err := _Futures.contract.UnpackLog(event, "PositionPaid", log); err != nil {
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

// ParsePositionPaid is a log parse operation binding the contract event 0x95aac10f3d1c6095ba6d61c88f22bcb31147dd8534a4f0d02b42c2841c271aac.
//
// Solidity: event PositionPaid(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) ParsePositionPaid(log types.Log) (*FuturesPositionPaid, error) {
	event := new(FuturesPositionPaid)
	if err := _Futures.contract.UnpackLog(event, "PositionPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesPositionPaymentReceivedIterator is returned from FilterPositionPaymentReceived and is used to iterate over the raw logs and unpacked data for PositionPaymentReceived events raised by the Futures contract.
type FuturesPositionPaymentReceivedIterator struct {
	Event *FuturesPositionPaymentReceived // Event containing the contract specifics and raw log

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
func (it *FuturesPositionPaymentReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesPositionPaymentReceived)
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
		it.Event = new(FuturesPositionPaymentReceived)
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
func (it *FuturesPositionPaymentReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesPositionPaymentReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesPositionPaymentReceived represents a PositionPaymentReceived event raised by the Futures contract.
type FuturesPositionPaymentReceived struct {
	PositionId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPositionPaymentReceived is a free log retrieval operation binding the contract event 0x9ac5613f028902b54981388ba4a28c1df8c2efd2987caea297c6a30330636e34.
//
// Solidity: event PositionPaymentReceived(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) FilterPositionPaymentReceived(opts *bind.FilterOpts, positionId [][32]byte) (*FuturesPositionPaymentReceivedIterator, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "PositionPaymentReceived", positionIdRule)
	if err != nil {
		return nil, err
	}
	return &FuturesPositionPaymentReceivedIterator{contract: _Futures.contract, event: "PositionPaymentReceived", logs: logs, sub: sub}, nil
}

// WatchPositionPaymentReceived is a free log subscription operation binding the contract event 0x9ac5613f028902b54981388ba4a28c1df8c2efd2987caea297c6a30330636e34.
//
// Solidity: event PositionPaymentReceived(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) WatchPositionPaymentReceived(opts *bind.WatchOpts, sink chan<- *FuturesPositionPaymentReceived, positionId [][32]byte) (event.Subscription, error) {

	var positionIdRule []interface{}
	for _, positionIdItem := range positionId {
		positionIdRule = append(positionIdRule, positionIdItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "PositionPaymentReceived", positionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesPositionPaymentReceived)
				if err := _Futures.contract.UnpackLog(event, "PositionPaymentReceived", log); err != nil {
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

// ParsePositionPaymentReceived is a log parse operation binding the contract event 0x9ac5613f028902b54981388ba4a28c1df8c2efd2987caea297c6a30330636e34.
//
// Solidity: event PositionPaymentReceived(bytes32 indexed positionId)
func (_Futures *FuturesFilterer) ParsePositionPaymentReceived(log types.Log) (*FuturesPositionPaymentReceived, error) {
	event := new(FuturesPositionPaymentReceived)
	if err := _Futures.contract.UnpackLog(event, "PositionPaymentReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Futures contract.
type FuturesTransferIterator struct {
	Event *FuturesTransfer // Event containing the contract specifics and raw log

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
func (it *FuturesTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesTransfer)
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
		it.Event = new(FuturesTransfer)
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
func (it *FuturesTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesTransfer represents a Transfer event raised by the Futures contract.
type FuturesTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Futures *FuturesFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FuturesTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FuturesTransferIterator{contract: _Futures.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Futures *FuturesFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FuturesTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesTransfer)
				if err := _Futures.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Futures *FuturesFilterer) ParseTransfer(log types.Log) (*FuturesTransfer, error) {
	event := new(FuturesTransfer)
	if err := _Futures.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FuturesUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Futures contract.
type FuturesUpgradedIterator struct {
	Event *FuturesUpgraded // Event containing the contract specifics and raw log

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
func (it *FuturesUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FuturesUpgraded)
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
		it.Event = new(FuturesUpgraded)
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
func (it *FuturesUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FuturesUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FuturesUpgraded represents a Upgraded event raised by the Futures contract.
type FuturesUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Futures *FuturesFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FuturesUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Futures.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FuturesUpgradedIterator{contract: _Futures.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Futures *FuturesFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FuturesUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Futures.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FuturesUpgraded)
				if err := _Futures.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Futures *FuturesFilterer) ParseUpgraded(log types.Log) (*FuturesUpgraded, error) {
	event := new(FuturesUpgraded)
	if err := _Futures.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
