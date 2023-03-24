// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OneInchV5

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
)

// GenericRouterSwapDescription is an auto generated low-level Go binding around an user-defined struct.
type GenericRouterSwapDescription struct {
	SrcToken        common.Address
	DstToken        common.Address
	SrcReceiver     common.Address
	DstReceiver     common.Address
	Amount          *big.Int
	MinReturnAmount *big.Int
	Flags           *big.Int
}

// OrderLibOrder is an auto generated low-level Go binding around an user-defined struct.
type OrderLibOrder struct {
	Salt          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	Receiver      common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
	Offsets       *big.Int
	Interactions  []byte
}

// OrderRFQLibOrderRFQ is an auto generated low-level Go binding around an user-defined struct.
type OrderRFQLibOrderRFQ struct {
	Info          *big.Int
	MakerAsset    common.Address
	TakerAsset    common.Address
	Maker         common.Address
	AllowedSender common.Address
	MakingAmount  *big.Int
	TakingAmount  *big.Int
}

// OneInchV5MetaData contains all meta data concerning the OneInchV5 contract.
var OneInchV5MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIWETH\",\"name\":\"weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessDenied\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdvanceNonceFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyFilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ArbitraryStaticCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadPool\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyPools\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EthDepositRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GetAmountCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMsgValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidatedOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MakingAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOneAmountShouldBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitLengthTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PredicateIsNotTrue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQBadSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQPrivateOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQSwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RFQZeroTargetIsForbidden\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyDetected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RemainingAmountIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReservesCallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReturnAmountIsNotEnough\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafePermitBadLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeTransferFromFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"res\",\"type\":\"bytes\"}],\"name\":\"SimulationResults\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapAmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapWithZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountIncreased\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TakingAmountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromMakerToTakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromTakerToMakerFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongGetter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroMinReturn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroReturnAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroTargetIsForbidden\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingRaw\",\"type\":\"uint256\"}],\"name\":\"OrderCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"}],\"name\":\"OrderFilledRFQ\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"advanceNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"and\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"arbitraryStaticCall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"orderRemaining\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"orderInfo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalMask\",\"type\":\"uint256\"}],\"name\":\"cancelOrderRFQ\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"checkPredicate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"}],\"name\":\"clipperSwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIClipperExchangeInterface\",\"name\":\"clipperExchange\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"outputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"goodUntil\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"clipperSwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"eq\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"vs\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"}],\"name\":\"fillOrderRFQCompact\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderRFQTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"filledMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filledTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"}],\"internalType\":\"structOrderRFQLib.OrderRFQ\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"flagsAndAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderRFQToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order_\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"fillOrderTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualMakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualTakingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"interaction\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skipPermitAndThresholdAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"fillOrderToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"gt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"makerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"takerAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"allowedSender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"takingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"interactions\",\"type\":\"bytes\"}],\"internalType\":\"structOrderLib.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"hashOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increaseNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slot\",\"type\":\"uint256\"}],\"name\":\"invalidatorForOrderRFQ\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"lt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"makerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"makerNonce\",\"type\":\"uint256\"}],\"name\":\"nonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offsets\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"or\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remaining\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"remainingRaw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"orderHashes\",\"type\":\"bytes32[]\"}],\"name\":\"remainingsRaw\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"simulate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAggregationExecutor\",\"name\":\"executor\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"srcReceiver\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"dstReceiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"flags\",\"type\":\"uint256\"}],\"internalType\":\"structGenericRouter.SwapDescription\",\"name\":\"desc\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spentAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"timestampBelow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeNonceAccount\",\"type\":\"uint256\"}],\"name\":\"timestampBelowAndNonceEquals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3Swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"uniswapV3SwapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"}],\"name\":\"unoswapTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"pools\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"permit\",\"type\":\"bytes\"}],\"name\":\"unoswapToWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600436106102f65760003560e01c80637e54f0921161018f578063bf15fcd8116100e1578063d365c6951161008a578063f2fde38b11610064578063f2fde38b14610859578063f78dc25314610879578063fa461e331461088c57600080fd5b8063d365c69514610813578063e449022e14610833578063e5d7bde61461084657600080fd5b8063c805a666116100bb578063c805a66614610799578063ca4ece22146107b9578063cf6fc6e3146107d957600080fd5b8063bf15fcd814610744578063bfa7514314610764578063c53a02921461078457600080fd5b8063942461bb11610143578063bc80f1a81161011d578063bc80f1a8146106f1578063bd61951d14610704578063bddccd351461072457600080fd5b8063942461bb146106915780639570eeee146106be578063bc1ed74c146106d157600080fd5b806383197ef01161017457806383197ef01461064157806384bd6d29146106565780638da5cb5b1461066957600080fd5b80637e54f092146105f4578063825caba11461062157600080fd5b80635a0998431161024857806370ae92d2116101fc57806372c244a8116101d657806372c244a81461059457806374261145146105b457806378e3214f146105d457600080fd5b806370ae92d21461053257806370ccbd311461055f578063715018a61461057f57600080fd5b806363592c2b1161022d57806363592c2b146104d25780636c838250146104f25780636fe7b0ba1461051257600080fd5b80635a099843146104ac57806362e238bb146104bf57600080fd5b80632d9a56f6116102aa5780633eca9c0a116102845780633eca9c0a1461041b5780634f38e2b81461044957806356f161241461046957600080fd5b80632d9a56f6146103bb57806337e7316f146103db5780633c15fd91146103fb57600080fd5b806312aa3caf116102db57806312aa3caf146103435780632521b9301461036b5780632cc2878d1461038b57600080fd5b80630502b1c51461030a578063093d4fa51461033057600080fd5b36610305576103036108ac565b005b600080fd5b61031d61031836600461483f565b6108b6565b6040519081526020015b60405180910390f35b61031d61033e3660046148a9565b6108d0565b610356610351366004614975565b610d16565b60408051928352602083019190915201610327565b34801561037757600080fd5b5061031d610386366004614a17565b610fd1565b34801561039757600080fd5b506103ab6103a6366004614abf565b611001565b6040519015158152602001610327565b3480156103c757600080fd5b506103566103d6366004614af1565b61104b565b3480156103e757600080fd5b5061031d6103f6366004614af1565b61114a565b34801561040757600080fd5b5061031d610416366004614a17565b611164565b61042e610429366004614c15565b611188565b60408051938452602084019290925290820152606001610327565b34801561045557600080fd5b506103ab610464366004614c72565b6111aa565b34801561047557600080fd5b5061031d610484366004614cbe565b6001600160a01b03919091166000908152600360209081526040808320938352929052205490565b61042e6104ba366004614cea565b6111d5565b61042e6104cd366004614d60565b61132b565b3480156104de57600080fd5b506103ab6104ed366004614abf565b421090565b3480156104fe57600080fd5b506103ab61050d366004614af1565b611355565b34801561051e57600080fd5b506103ab61052d366004614c72565b611384565b34801561053e57600080fd5b5061031d61054d366004614e0c565b60016020526000908152604090205481565b34801561056b57600080fd5b5061042e61057a366004614e29565b6113aa565b34801561058b57600080fd5b506103036113f1565b3480156105a057600080fd5b506103036105af366004614ecd565b611403565b3480156105c057600080fd5b506103ab6105cf366004614c72565b6114b2565b3480156105e057600080fd5b506103036105ef366004614cbe565b611524565b34801561060057600080fd5b5061031d61060f366004614abf565b60009081526002602052604090205490565b34801561062d57600080fd5b5061030361063c366004614abf565b611544565b34801561064d57600080fd5b50610303611553565b61031d610664366004614ef0565b61155e565b34801561067557600080fd5b506000546040516001600160a01b039091168152602001610327565b34801561069d57600080fd5b506106b16106ac366004614f67565b611571565b6040516103279190615001565b61042e6106cc366004615045565b61162a565b3480156106dd57600080fd5b5061031d6106ec366004614abf565b611767565b61031d6106ff36600461483f565b6117b7565b34801561071057600080fd5b5061030361071f366004615082565b6117c6565b34801561073057600080fd5b5061030361073f3660046150be565b611867565b34801561075057600080fd5b5061031d61075f366004615082565b611872565b34801561077057600080fd5b506103ab61077f366004614c72565b6118bd565b34801561079057600080fd5b50610303611930565b3480156107a557600080fd5b5061031d6107b43660046150e0565b61193a565b3480156107c557600080fd5b506103ab6107d4366004614c72565b611971565b3480156107e557600080fd5b506103ab6107f4366004614cbe565b6001600160a01b03919091166000908152600160205260409020541490565b34801561081f57600080fd5b5061042e61082e3660046151a4565b611998565b61031d610841366004615295565b611a2c565b61042e6108543660046152e8565b611a3b565b34801561086557600080fd5b50610303610874366004614e0c565b6124cb565b61031d6108873660046153ac565b612558565b34801561089857600080fd5b506103036108a7366004615416565b612573565b6108b4612785565b565b60006108c63387878787876127be565b9695505050505050565b60006001600160a01b0388161580156109085786341461090357604051631841b4e160e01b815260040160405180910390fd5b610a24565b7f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b0316896001600160a01b0316036109f057506001341561096357604051631841b4e160e01b815260040160405180910390fd5b6040516323b872dd60e01b808252336004830152306024830152604482018990527f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc291632e1a7d4d60e01b9060008060648382885af16109c6573d6000823e3d81fd5b8181528a60048201526000806024836000885af16109e7573d6000823e3d81fd5b50505050610a24565b3415610a0f57604051631841b4e160e01b815260040160405180910390fd5b610a246001600160a01b038a16338d8a612b30565b8015610ab85760008b905060006327a9b42460e01b90506040518181528a60048201528960248201528860448201528760648201528c60848201528560ff1c601b0160a48201528660c48201526001600160ff1b03861660e482015261012061010482015264a62929c86960d31b610143820152600080610149838d875af1610ab0573d6000823e3d81fd5b505050610d07565b6001600160a01b0388161580610aff57507f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b0316886001600160a01b0316145b15610c75576040517f4cb6864c00000000000000000000000000000000000000000000000000000000808252600482018b90526024820189905260448201889052606482018790528c918a1560018114610b5e57306084830152610b65565b8d60848301525b508560ff1c601b0160a48201528660c48201526001600160ff1b03861660e482015261012061010482015264a62929c86960d31b610143820152600080610149836000875af1610bb8573d6000823e3d81fd5b507f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b03168a6001600160a01b031603610c6e57604051630d0e30db60e41b8082527f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc29163a9059cbb60e01b906000806004838f885af1610c42573d6000823e3d81fd5b8181528f60048201528b60248201526000806044836000885af1610c69573d6000823e3d81fd5b505050505b5050610d07565b60008b90506000632b651a6c60e01b90506040518181528b60048201528a60248201528960448201528860648201528760848201528c60a48201528560ff1c601b0160c48201528660e48201526001600160ff1b03861661010482015261014061012482015264a62929c86960d31b610163820152600080610169836000875af1610d03573d6000823e3d81fd5b5050505b50939998505050505050505050565b6000808660a00135600003610d57576040517f0262dde400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610d666020890189614e0c565b90506000610d7a60408a0160208b01614e0c565b90506000610d90836001600160a01b0316612bcd565b905060c08a013560021615610dd55780610dab576000610db1565b89608001355b3411610dd057604051631841b4e160e01b815260040160405180910390fd5b610e06565b80610de1576000610de7565b89608001355b3414610e0657604051631841b4e160e01b815260040160405180910390fd5b80610e4f578715610e2557610e256001600160a01b0384168a8a612c06565b610e4f33610e3960608d0160408e01614e0c565b6001600160a01b038616919060808e0135612b30565b610e608b338c608001358a8a612cbf565b60808a01359350610e7a6001600160a01b03831630612d1f565b945084600003610eb6576040517f28ebf24700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000199094019360c08a013560011615610f4f576000610edf6001600160a01b03851630612d1f565b90506001811115610f0e5760001901610ef88186615473565b9450610f0e6001600160a01b0385163383612dca565b610f1c8560a08d0135615486565b610f2a60808d013588615486565b1015610f495760405163f32bec2f60e01b815260040160405180910390fd5b50610f74565b8960a00135851015610f745760405163f32bec2f60e01b815260040160405180910390fd5b600080610f8760808d0160608e01614e0c565b6001600160a01b031614610faa57610fa560808c0160608d01614e0c565b610fac565b335b9050610fc26001600160a01b0384168288612dca565b50505050965096945050505050565b6000610fe76001600160a01b0389168484612c06565b610ff48988888888612eaa565b9998505050505050505050565b600060d082901c60a083901c65ffffffffffff168361101f83421090565b801561104257506001600160a01b03811660009081526001602052604090205482145b95945050505050565b6000803361105f6080850160608601614e0c565b6001600160a01b03161461109f576040517f4ca8886700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110a88361114a565b6000818152600260205260409020549250905060001982016110f6576040517f41a26a6300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080518281526020810184905233917fcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1910160405180910390a26000818152600260205260409020600190559092909150565b600061115e611157613131565b8390613258565b92915050565b600061117a6001600160a01b0389168484612c06565b610ff48989898989896127be565b600080600061119a87878787336111d5565b9250925092509450945094915050565b60008060006111b985856132d7565b915091508180156111c957508581115b925050505b9392505050565b60008060006112356111e5613131565b601f198a0180517f74ab4f0cde46aaf927859983f7d04002116dd057d4c4941f6dbfb775c3e31f4582526101008220915260405161190160f01b8152600281019290925260228201526042902090565b9050600160fe1b8516156112a957600160fd1b851615801590611259575060418614155b15611277576040516317c2b1f160e01b815260040160405180910390fd5b6112878860600151828989613466565b6112a4576040516317c2b1f160e01b815260040160405180910390fd5b6112d6565b6112b988606001518289896134bb565b6112d6576040516317c2b1f160e01b815260040160405180910390fd5b6112e1888686613522565b60408051848152602081018490529295509093507fc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127910160405180910390a1955095509592505050565b60008060006113418b8b8b8b8b8b8b8b33611a3b565b925092509250985098509895505050505050565b600080600061136b61136685613b2e565b6132d7565b9150915081801561137c5750806001145b949350505050565b600080600061139385856132d7565b915091508180156111c95750909414949350505050565b60008060006113d185858c604001516001600160a01b0316612c069092919063ffffffff16565b6113de8a8a8a8a8a6111d5565b9250925092509750975097945050505050565b6113f9613b45565b6108b46000613b9f565b8060ff16600003611440576040517fbd71636d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3360009081526001602052604081205461145e9060ff84169061549d565b336000818152600160205260409081902083905551919250907ffc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db906114a69084815260200190565b60405180910390a25050565b60008080805b63ffffffff87821c1692508215611517576000806114db61136686868a8c6154b0565b915091508180156114ec5750806001145b156114ff576001955050505050506111ce565b50839250611510905060208261549d565b90506114b8565b5060009695505050505050565b61152c613b45565b6115406001600160a01b0383163383612dca565b5050565b61155033826000613c07565b50565b61155b613b45565b33ff5b6000610ff489338a8a8a8a8a8a8a6108d0565b60606000825167ffffffffffffffff81111561158f5761158f614b26565b6040519080825280602002602001820160405280156115b8578160200160208202803683370190505b50905060005b835181101561162357600260008583815181106115dd576115dd6154da565b6020026020010151815260200190815260200160002054828281518110611606576116066154da565b60209081029190910101528061161b816154f0565b9150506115be565b5092915050565b600080600061168a61163a613131565b601f19890180517f74ab4f0cde46aaf927859983f7d04002116dd057d4c4941f6dbfb775c3e31f4582526101008220915260405161190160f01b8152600281019290925260228201526042902090565b9050600160fe1b8416156116e657600160fd1b8416156116d6576116b48760600151828888613c9b565b6116d1576040516317c2b1f160e01b815260040160405180910390fd5b611713565b6116b48760600151828888613d05565b6116f68760600151828888613d5a565b611713576040516317c2b1f160e01b815260040160405180910390fd5b61171e878533613522565b60408051848152602081018490529295509093507fc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127910160405180910390a19450945094915050565b600081815260026020526040812054806117ad576040517fb838de9600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000190192915050565b60006108c68686868686612eaa565b600080846001600160a01b031684846040516117e3929190615509565b600060405180830381855af49150503d806000811461181e576040519150601f19603f3d011682016040523d82523d6000602084013e611823565b606091505b509150915081816040517f1934afc800000000000000000000000000000000000000000000000000000000815260040161185e929190615569565b60405180910390fd5b611540338383613c07565b6000806000611882868686613da9565b9150915081611042576040517f1f1b8f6100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008080805b63ffffffff87821c1692508215611923576000806118e661136686868a8c6154b0565b915091508115806118f8575080600114155b1561190b576000955050505050506111ce565b5083925061191c905060208261549d565b90506118c3565b5060019695505050505050565b6108b46001611403565b60006119506001600160a01b038b168484612c06565b6119618c8c8c8c8c8c8c8c8c6108d0565b9c9b505050505050505050505050565b600080600061198085856132d7565b915091508180156111c9575094909410949350505050565b6000808060148410156119d7576040517fd9e1c6dc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003660006119e68888613dd7565b91945092509050611a016001600160a01b0384168383612c06565b505050611a158e8e8e8e8e8e8e8e8e611a3b565b9250925092509b509b509b98505050505050505050565b60006110423386868686612eaa565b600080806001600160a01b038416611a7f576040517fb0c4d05f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a888c61114a565b6000818152600260205260409020548894508793509091508c906000198101611add576040517fecef366400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611aef60c0840160a08501614e0c565b6001600160a01b031614158015611b1e575033611b1260c0840160a08501614e0c565b6001600160a01b031614155b15611b55576040517fd4dfdafe00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80611c5857611b75611b6d6080840160608501614e0c565b848f8f6134bb565b611bab576040517f5cd5d23300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060c0810135366000611bbd84613e15565b91509150600160ff1b89166000148015611bd8575060148110155b15611c51576000366000611bec8585613dd7565b91945092509050611c076001600160a01b0384168383612c06565b60008881526002602052604090205415611c4d576040517fc5f2be5100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505b5050611c5d565b600019015b6000611c6883613b2e565b90501115611caf57611c7982611355565b611caf576040517fb6629c0200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8415841503611ce9576040517ee2a52200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600003611d795780851115611cfd578094505b611d1b611d0983613e23565b8460c00135888660e001358689613e31565b93506001600160ff1b038716611d318682615486565b611d3b8b87615486565b1115611d73576040517ffb8ae12900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50611e44565b611d97611d8583613e67565b8460e00135878660c001358689613e75565b945080851115611dec57809450611db0611d0983613e23565b935087841115611dec576040517f939c420400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160ff1b038716611e008582615486565b611e0a8a88615486565b1015611e42576040517f481ea39200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505b841580611e4f575083155b15611e86576040517ffba5a27600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b84810390508060010160026000858152602001908152602001600020819055508d6060016020810190611eb99190614e0c565b6001600160a01b03167fb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f028483604051611efc929190918252602082015260400190565b60405180910390a26014611f0f83613e89565b905010611fb2576000366000611f2c611f2786613e89565b613dd7565b919450925090506001600160a01b0383166396a10e3387611f536080890160608a01614e0c565b338c8c8a89896040518963ffffffff1660e01b8152600401611f7c9897969594939291906155ad565b600060405180830381600087803b158015611f9657600080fd5b505af1158015611faa573d6000803e3d6000fd5b505050505050505b611fe5611fc56040840160208501614e0c565b611fd56080850160608601614e0c565b8888611fe087613e97565b613ea5565b61201b576040517f70a03f4800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60148a106120f35760003660006120328e8e613dd7565b9250925092506000836001600160a01b031663ccee33d7338b8b87876040518663ffffffff1660e01b815260040161206e959493929190615600565b6020604051808303816000875af115801561208d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120b1919061562f565b905087811180156120d057506120ce6120c987613e67565b613f00565b155b80156120e557506120e36120c987613e23565b155b156120ee578097505b505050505b6001600160a01b037f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc21661212d6060840160408501614e0c565b6001600160a01b03161480156121435750600034115b15612359578334101561216957604051631841b4e160e01b815260040160405180910390fd5b833411156121df57604051600090339034879003908381818185875af1925050503d80600081146121b6576040519150601f19603f3d011682016040523d82523d6000602084013e6121bb565b606091505b50509050806121dd5760405163b12d13eb60e01b815260040160405180910390fd5b505b7f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561223a57600080fd5b505af115801561224e573d6000803e3d6000fd5b50506001600160a01b037f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc216925063a9059cbb91506000905061229760a0860160808701614e0c565b6001600160a01b0316146122ba576122b560a0850160808601614e0c565b6122ca565b6122ca6080850160608601614e0c565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b1681526001600160a01b039091166004820152602481018790526044016020604051808303816000875af115801561232f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123539190615648565b50612411565b341561237857604051631841b4e160e01b815260040160405180910390fd5b6123db61238b6060840160408501614e0c565b33600061239e60a0870160808801614e0c565b6001600160a01b0316146123c1576123bc60a0860160808701614e0c565b6123d1565b6123d16080860160608701614e0c565b87611fe087613f75565b612411576040517f478a520500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b601461241c83613f83565b9050106124ba576000366000612434611f2786613f83565b919450925090506001600160a01b038316633504ed628761245b6080890160608a01614e0c565b338c8c8a89896040518963ffffffff1660e01b81526004016124849897969594939291906155ad565b600060405180830381600087803b15801561249e57600080fd5b505af11580156124b2573d6000803e3d6000fd5b505050505050505b505099509950999650505050505050565b6124d3613b45565b6001600160a01b03811661254f5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161185e565b61155081613b9f565b60006125688787878787876127be565b979650505050505050565b6125cc565b3d6000803e3d6000fd5b8061258f5761258f612578565b600160005114601f3d11163d151780611540577ff27f64e40000000000000000000000000000000000000000000000000000000060005260046000fd5b604051601581017f0dfe1681d21220a7ddca3f43a9059cbb23b872dd0000000000000000000000008252602081600484335afa61260b5761260b612578565b60208082016004808501335afa61262457612624612578565b602060408201600460088501335afa61263f5761263f612578565b600080600088136001811461265d5760208401519250879150612665565b835192508891505b507fff1f98431c8ad98523631ae4a59f267346ea31f984000000000000000000000084526060832083527fe34f199b19b2b4f47f68442619d555527d244f78a3297ea89325f843f87b8b5460208401526001600160a01b0360558520169250338318156126f6577fb2c027220000000000000000000000000000000000000000000000000000000060005260046000fd5b60843592507f0dfe1681d21220a7ddca3f43a9059cbb23b872dd00000000000000000000000084523083146001811461275757836014860152336034860152816054860152612752602060006064601089016000885af1612582565b61277a565b33601086015281603086015261277a602060006044600c89016000885af1612582565b505050505050505050565b3233036108b4576040517f1b10b0f900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006128b9565b7f0902f1ac0000000000000000000000000000000000000000000000000000000081526000604082600484875afa6127ff576127ff612578565b60603d14612831577f85cd58dc0000000000000000000000000000000000000000000000000000000060005260046000fd5b81516020830151861561284057905b8785029250633b9aca008202830181840204925050507f022c0d9f000000000000000000000000000000000000000000000000000000008252841594508415810260048301528481026024830152866044830152608060648301526000608483015260008060a4846000885af16108c6576108c6612578565b6dffffffffffffffffffffffffffff8511156128f9577fcf0b4d3a0000000000000000000000000000000000000000000000000000000060005260046000fd5b60405160c081016040528260051b84018435886000811461296957341561292b57631841b4e160e01b60005260046000fd5b6323b872dd60e01b84523360048501526001600160a01b03821660248501528860448501526129646020600060648760008f5af1612582565b6129ff565b34891461298157631841b4e160e01b60005260046000fd5b630d0e30db60e41b84526000806004868c73c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25af16129b5576129b5612578565b63a9059cbb60e01b84526001600160a01b0382166004850152886024850152600080604486600073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25af16129ff576129ff612578565b50879350602086015b82811015612a50578035612a446001600160a01b03821663ffffffff60a01b851660a01c600160ff1b86166001600160a01b0387168a8a6127c5565b95509150602001612a08565b50600160fe1b81168015612adf57612a873063ffffffff60a01b841660a01c600160ff1b85166001600160a01b03861689896127c5565b9450632e1a7d4d60e01b8452846004850152600080602486600073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc25af1612ac457612ac4612578565b600080600080888f5af1612ada57612ada612578565b612b0b565b612b088b63ffffffff60a01b841660a01c600160ff1b85166001600160a01b03861689896127c5565b94505b50505050838110156108c65760405163f32bec2f60e01b815260040160405180910390fd5b60006323b872dd60e01b905060006040518281528560048201528460248201528360448201526020600060648360008b5af19150508015612b8e573d8015612b8457600160005114601f3d11169150612b8c565b6000873b1191505b505b80612bc5576040517ff405907100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050565b60006001600160a01b038216158061115e57506001600160a01b03821673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1492915050565b600060e0829003612c4457612c3d847fd505accf000000000000000000000000000000000000000000000000000000008585613f91565b9050612cac565b610100829003612c7a57612c3d847f8fcbaf0c000000000000000000000000000000000000000000000000000000008585613f91565b6040517f6827585700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80612cb957612cb9613fe3565b50505050565b6040517f4b64e4920000000000000000000000000000000000000000000000000000000080825260048201869052908284602483013784836024830101526000808460440183348b5af1612d16573d6000823e3d81fd5b50505050505050565b6000612d2a83612bcd565b15612d4057506001600160a01b0381163161115e565b6040517f70a082310000000000000000000000000000000000000000000000000000000081526001600160a01b0383811660048301528416906370a0823190602401602060405180830381865afa158015612d9f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612dc3919061562f565b905061115e565b8015612ea557612dd983612bcd565b15612e915780471015612e18576040517ff4d678b800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000826001600160a01b03168261138890604051600060405180830381858888f193505050503d8060008114612e6a576040519150601f19603f3d011682016040523d82523d6000602084013e612e6f565b606091505b5050905080612cb95760405163b12d13eb60e01b815260040160405180910390fd5b612ea56001600160a01b0384168383613fef565b505050565b600081808203612ee6576040517f67e7c0f600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8591506000198101341515600080600160fd1b888886818110612f0b57612f0b6154da565b90506020020135161190508115612fb157883414612f3c57604051631841b4e160e01b815260040160405180910390fd5b7f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031663d0e30db08a6040518263ffffffff1660e01b81526004016000604051808303818588803b158015612f9757600080fd5b505af1158015612fab573d6000803e3d6000fd5b50505050505b600184111561305157612feb3083612fc95733612fcb565b305b89896000818110612fde57612fde6154da565b9050602002013588614038565b945060015b838110156130265761301c30308a8a8581811061300f5761300f6154da565b9050602002013589614038565b9550600101612ff0565b5061304a81613035578a613037565b305b30898987818110612fde57612fde6154da565b945061306f565b61306c8161305f578a613061565b305b83612fc95733612fcb565b94505b878510156130905760405163f32bec2f60e01b815260040160405180910390fd5b801561312457604051632e1a7d4d60e01b8152600481018690527f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031690632e1a7d4d90602401600060405180830381600087803b1580156130f857600080fd5b505af115801561310c573d6000803e3d6000fd5b50613124925050506001600160a01b038b16866141d4565b5050505095945050505050565b6000306001600160a01b037f0000000000000000000000001111111254eeb25477b68fb85ed929f73a9605821614801561318a57507f000000000000000000000000000000000000000000000000000000000000000146145b156131b457507f1c0eb4c27d5b523ca136c0b3b83a4dcac8b70225b38be8507ba1a3f2af03cfca90565b50604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6020808301919091527f5c6cbfb2848b981a8f93044b3530be1fac304ecd5042396ca8729cb8fdd718f3828401527fceebf77a833b30520287ddd9478ff51abbdffa30aa90a8d655dba0e8a79ce0c160608301524660808301523060a0808401919091528351808403909101815260c0909201909252805191012090565b6000368161326a61012086018661566a565b60405191935091507f0a244ca8a150ac294c14fcff9277051ced9a5b23e966a0ff0522e989da23116c9082848237828120610140820152610120876020830137818152610160902060405161190160f01b81526002810187905260228101829052604290209094506108c6565b60008060006132e685856142ed565b60e01c905060006132f986866004614317565b9050632cc2878c19820161333057600161331282611001565b61331d576000613320565b60015b90945060ff16925061345f915050565b63bf15fcd88210156133be57636fe7b0ba82101561338257634f38e2b71982016133665760016133128261046489896064614348565b6363592c2a19820161337d57600161331282421090565b61344d565b636fe7b0b91982016133a05760016133128261052d89896064614348565b637426114419820161337d576001613312826105cf89896064614348565b63ca4ece228210156134115763bf15fcd71982016133f35760016133e88261075f89896064614348565b93509350505061345f565b63bfa7514219820161337d5760016133128261077f89896064614348565b63ca4ece2119820161342f576001613312826107d489896064614348565b63cf6fc6e219820161344d576001613312826107f489896024614317565b613458308787613da9565b9350935050505b9250929050565b600080631626ba7e60e01b905060405181815285600482015260406024820152836044820152838560648301376020600085606401838a5afa156134b15760203d1460005183141692505b5050949350505050565b60006001600160a01b0385166134d35750600061137c565b60408214806134e25750604182145b80156135095750846001600160a01b03166134fe858585614378565b6001600160a01b0316145b156135165750600161137c565b61104285858585613466565b6000806001600160a01b038316613565576040517f692e45e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b606085015160808601516001600160a01b031615801590613593575060808601516001600160a01b03163314155b156135ca576040517fe8c6632100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b855167ffffffffffffffff604082901c1680158015906135e957508042115b15613620576040517fc56873ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61362c83836000613c07565b505060a086015160c08701517f0fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8716600081900361366f57829550819450613715565b600160ff1b8816156136ca57828111156136b5576040517faa34b69600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8095506136c383838861442d565b9450613715565b81811115613704576040517f7f902a9300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80945061371283838761445b565b95505b5050508260001480613725575081155b1561375c576040517f07b6e79f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031686602001516001600160a01b03161480156137c257507f1000000000000000000000000000000000000000000000000000000000000000851615155b1561395a576040516323b872dd60e01b81526001600160a01b038281166004830152306024830152604482018590527f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc216906323b872dd906064016020604051808303816000875af115801561383c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138609190615648565b50604051632e1a7d4d60e01b8152600481018490527f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031690632e1a7d4d90602401600060405180830381600087803b1580156138c357600080fd5b505af11580156138d7573d6000803e3d6000fd5b505050506000846001600160a01b03168461138890604051600060405180830381858888f193505050503d806000811461392d576040519150601f19603f3d011682016040523d82523d6000602084013e613932565b606091505b50509050806139545760405163b12d13eb60e01b815260040160405180910390fd5b50613974565b6020860151613974906001600160a01b0316828686612b30565b7f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031686604001516001600160a01b03161480156139b95750600034115b15613aec578134146139de57604051631841b4e160e01b815260040160405180910390fd5b7f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031663d0e30db0836040518263ffffffff1660e01b81526004016000604051808303818588803b158015613a3957600080fd5b505af1158015613a4d573d6000803e3d6000fd5b505060405163a9059cbb60e01b81526001600160a01b038581166004830152602482018790527f000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc216935063a9059cbb925060440190506020604051808303816000875af1158015613ac2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613ae69190615648565b50613b25565b3415613b0b57604051631841b4e160e01b815260040160405180910390fd5b6040860151613b25906001600160a01b0316338385612b30565b50935093915050565b366000613b3c836004614468565b91509150915091565b6000546001600160a01b031633146108b45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161185e565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038316600090815260036020908152604080832066ffffffffffffff600887901c16808552928190529220549091600160ff86161b841791808316839003613c82576040517ff71fbda200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000938452602091909152604090922091179055505050565b600080631626ba7e60e01b905060405181815285600482015260406024820152604160448201528460648201526001600160ff1b03841660848201528360ff1c601b0160a48201536020600060a5838a5afa156134b15750600051143d6020141695945050505050565b600080631626ba7e60e01b905060405181815285600482015260406024820152604060448201528460648201528360848201526020600060a4838a5afa156134b15750600051143d6020141695945050505050565b60006001600160a01b038516613d725750600061137c565b846001600160a01b0316613d878585856144c3565b6001600160a01b031603613d9d5750600161137c565b61104285858585613d05565b60008060405183858237602060008583895afa3d602014169250508115613dcf57506000515b935093915050565b600036816014841015613dfd5760405163779ab6bd60e11b815260040160405180910390fd5b505050813560601c9260149092019160131990910190565b366000613b3c836005614468565b366000613b3c836003614468565b6000868103613e4c57613e4586858761442d565b9050612568565b613e5b88888888888888614532565b98975050505050505050565b366000613b3c836002614468565b6000868103613e4c57613e4584878761445b565b366000613b3c836006614468565b366000613b3c836000614468565b6040516323b872dd60e01b8082526004820187905260248201869052604482018590526000918385606483013760206000856064018360008d5af19050600160005114601f3d11163d15178116925050509695505050505050565b60006001821480156111ce575082826000818110613f2057613f206154da565b9050013560f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167f7800000000000000000000000000000000000000000000000000000000000000149392505050565b366000613b3c836001614468565b366000613b3c836007614468565b6000816004016040518581528385600483013760206000838360008b5af192505050801561137c573d8015613fd257600160005114601f3d11169150613fda565b6000863b1191505b50949350505050565b6040513d6000823e3d81fd5b6140028363a9059cbb60e01b84846146be565b612ea5576040517ffb7f507900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600160ff1b831615801561410e576000846001600160a01b031663128acb08888461406488614700565b604080516001600160a01b038d1660208201526401000276a491016040516020818303038152906040526040518663ffffffff1660e01b81526004016140ae9594939291906156b1565b60408051808303816000875af11580156140cc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906140f091906156eb565b915050614105816141009061570f565b614783565b9250505061137c565b6000846001600160a01b031663128acb08888461412a88614700565b604080516001600160a01b038d16602082015273fffd8963efd1fc6a506488495d951d5263988d2591016040516020818303038152906040526040518663ffffffff1660e01b81526004016141839594939291906156b1565b60408051808303816000875af11580156141a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141c591906156eb565b5090506141056141008261570f565b804710156142245760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604482015260640161185e565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114614271576040519150601f19603f3d011682016040523d82523d6000602084013e614276565b606091505b5050905080612ea55760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d61792068617665207265766572746564000000000000606482015260840161185e565b600060048210156143115760405163779ab6bd60e11b815260040160405180910390fd5b50503590565b60006020820183101561433d5760405163779ab6bd60e11b815260040160405180910390fd5b509190910135919050565b3660008284101561436c5760405163779ab6bd60e11b815260040160405180910390fd5b50509182019291900390565b6000604051826041811461439757604081146143b157600091506143d9565b604085013560001a602083015260408560408401376143d9565b60208501358060ff1c601b01602084015260208660408501376001600160ff1b031660608301525b508015614425577f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a16060820151101561442557848152600080526020600060808360015afa5060005191505b509392505050565b60008360018161443d8686615486565b614447919061549d565b6144519190615473565b61137c919061572b565b6000826144518584615486565b3660008060058460078111156144805761448061574d565b901b905061449261012086018661566a565b6144b79161010088013580851c63ffffffff9081169360209290921b861c16916154b0565b92509250509250929050565b60006001600160ff1b0382167f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a1811015614425576040518581528360ff1c601b016020820152846040820152816060820152600080526020600060808360015afa505060005195945050505050565b600060018790036145be576145478888613f00565b1561458c57858514614585576040517f49986e7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5082612568565b6040517fbec74c8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003660006145cd8b8b613dd7565b925092509250600080846001600160a01b031684848c8b8b6040516020016145f9959493929190615763565b60408051601f198184030181529082905261461391615782565b600060405180830381855afa9150503d806000811461464e576040519150601f19603f3d011682016040523d82523d6000602084013e614653565b606091505b509150915081158061466757508051602014155b1561469e576040517f110b8e7300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808060200190518101906146b2919061562f565b95505050505050612568565b60006040518481528360048201528260248201526020600060448360008a5af1915050801561137c573d8015613fd257600160005114601f3d11169150613fda565b60006001600160ff1b0382111561477f5760405162461bcd60e51b815260206004820152602860248201527f53616665436173743a2076616c756520646f65736e27742066697420696e206160448201527f6e20696e74323536000000000000000000000000000000000000000000000000606482015260840161185e565b5090565b60008082121561477f5760405162461bcd60e51b815260206004820181905260248201527f53616665436173743a2076616c7565206d75737420626520706f736974697665604482015260640161185e565b6001600160a01b038116811461155057600080fd5b80356147f5816147d5565b919050565b60008083601f84011261480c57600080fd5b50813567ffffffffffffffff81111561482457600080fd5b6020830191508360208260051b850101111561345f57600080fd5b60008060008060006080868803121561485757600080fd5b8535614862816147d5565b94506020860135935060408601359250606086013567ffffffffffffffff81111561488c57600080fd5b614898888289016147fa565b969995985093965092949392505050565b60008060008060008060008060006101208a8c0312156148c857600080fd5b89356148d3816147d5565b985060208a01356148e3816147d5565b975060408a01356148f3816147d5565b965060608a0135614903816147d5565b989b979a50959860808101359760a0820135975060c0820135965060e08201359550610100909101359350915050565b60008083601f84011261494557600080fd5b50813567ffffffffffffffff81111561495d57600080fd5b60208301915083602082850101111561345f57600080fd5b60008060008060008086880361014081121561499057600080fd5b873561499b816147d5565b965060e0601f19820112156149af57600080fd5b5060208701945061010087013567ffffffffffffffff808211156149d257600080fd5b6149de8a838b01614933565b90965094506101208901359150808211156149f857600080fd5b50614a0589828a01614933565b979a9699509497509295939492505050565b60008060008060008060008060c0898b031215614a3357600080fd5b8835614a3e816147d5565b97506020890135614a4e816147d5565b96506040890135955060608901359450608089013567ffffffffffffffff80821115614a7957600080fd5b614a858c838d016147fa565b909650945060a08b0135915080821115614a9e57600080fd5b50614aab8b828c01614933565b999c989b5096995094979396929594505050565b600060208284031215614ad157600080fd5b5035919050565b60006101408284031215614aeb57600080fd5b50919050565b600060208284031215614b0357600080fd5b813567ffffffffffffffff811115614b1a57600080fd5b61137c84828501614ad8565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715614b6557614b65614b26565b604052919050565b600060e08284031215614b7f57600080fd5b60405160e0810181811067ffffffffffffffff82111715614ba257614ba2614b26565b604052823581529050806020830135614bba816147d5565b60208201526040830135614bcd816147d5565b60408201526060830135614be0816147d5565b60608201526080830135614bf3816147d5565b8060808301525060a083013560a082015260c083013560c08201525092915050565b6000806000806101208587031215614c2c57600080fd5b614c368686614b6d565b935060e085013567ffffffffffffffff811115614c5257600080fd5b614c5e87828801614933565b959890975094956101000135949350505050565b600080600060408486031215614c8757600080fd5b83359250602084013567ffffffffffffffff811115614ca557600080fd5b614cb186828701614933565b9497909650939450505050565b60008060408385031215614cd157600080fd5b8235614cdc816147d5565b946020939093013593505050565b60008060008060006101408688031215614d0357600080fd5b614d0d8787614b6d565b945060e086013567ffffffffffffffff811115614d2957600080fd5b614d3588828901614933565b9095509350506101008601359150610120860135614d52816147d5565b809150509295509295909350565b60008060008060008060008060c0898b031215614d7c57600080fd5b883567ffffffffffffffff80821115614d9457600080fd5b614da08c838d01614ad8565b995060208b0135915080821115614db657600080fd5b614dc28c838d01614933565b909950975060408b0135915080821115614ddb57600080fd5b50614de88b828c01614933565b999c989b5096999698976060880135976080810135975060a0013595509350505050565b600060208284031215614e1e57600080fd5b81356111ce816147d5565b6000806000806000806000610160888a031215614e4557600080fd5b614e4f8989614b6d565b965060e088013567ffffffffffffffff80821115614e6c57600080fd5b614e788b838c01614933565b90985096506101008a013595506101208a01359150614e96826147d5565b9093506101408901359080821115614ead57600080fd5b50614eba8a828b01614933565b989b979a50959850939692959293505050565b600060208284031215614edf57600080fd5b813560ff811681146111ce57600080fd5b600080600080600080600080610100898b031215614f0d57600080fd5b8835614f18816147d5565b97506020890135614f28816147d5565b96506040890135614f38816147d5565b979a96995096976060810135975060808101359660a0820135965060c0820135955060e0909101359350915050565b60006020808385031215614f7a57600080fd5b823567ffffffffffffffff80821115614f9257600080fd5b818501915085601f830112614fa657600080fd5b813581811115614fb857614fb8614b26565b8060051b9150614fc9848301614b3c565b8181529183018401918481019088841115614fe357600080fd5b938501935b83851015613e5b57843582529385019390850190614fe8565b6020808252825182820181905260009190848201906040850190845b818110156150395783518352928401929184019160010161501d565b50909695505050505050565b600080600080610140858703121561505c57600080fd5b6150668686614b6d565b9660e08601359650610100860135956101200135945092505050565b60008060006040848603121561509757600080fd5b83356150a2816147d5565b9250602084013567ffffffffffffffff811115614ca557600080fd5b600080604083850312156150d157600080fd5b50508035926020909101359150565b60008060008060008060008060008060006101408c8e03121561510257600080fd5b8b3561510d816147d5565b9a5060208c013561511d816147d5565b995060408c013561512d816147d5565b985060608c013561513d816147d5565b975060808c0135965060a08c0135955060c08c0135945060e08c013593506101008c013592506101208c013567ffffffffffffffff81111561517e57600080fd5b61518a8e828f01614933565b915080935050809150509295989b509295989b9093969950565b60008060008060008060008060008060006101008c8e0312156151c657600080fd5b67ffffffffffffffff808d3511156151dd57600080fd5b6151ea8e8e358f01614ad8565b9b508060208e013511156151fd57600080fd5b61520d8e60208f01358f01614933565b909b50995060408d013581101561522357600080fd5b6152338e60408f01358f01614933565b909950975060608d0135965060808d0135955060a08d0135945061525960c08e016147ea565b93508060e08e0135111561526c57600080fd5b5061527d8d60e08e01358e01614933565b81935080925050509295989b509295989b9093969950565b600080600080606085870312156152ab57600080fd5b8435935060208501359250604085013567ffffffffffffffff8111156152d057600080fd5b6152dc878288016147fa565b95989497509550505050565b600080600080600080600080600060e08a8c03121561530657600080fd5b893567ffffffffffffffff8082111561531e57600080fd5b61532a8d838e01614ad8565b9a5060208c013591508082111561534057600080fd5b61534c8d838e01614933565b909a50985060408c013591508082111561536557600080fd5b506153728c828d01614933565b90975095505060608a0135935060808a0135925060a08a0135915060c08a013561539b816147d5565b809150509295985092959850929598565b60008060008060008060a087890312156153c557600080fd5b86356153d0816147d5565b955060208701356153e0816147d5565b94506040870135935060608701359250608087013567ffffffffffffffff81111561540a57600080fd5b614a0589828a016147fa565b6000806000806060858703121561542c57600080fd5b8435935060208501359250604085013567ffffffffffffffff81111561545157600080fd5b6152dc87828801614933565b634e487b7160e01b600052601160045260246000fd5b8181038181111561115e5761115e61545d565b808202811582820484141761115e5761115e61545d565b8082018082111561115e5761115e61545d565b600080858511156154c057600080fd5b838611156154cd57600080fd5b5050820193919092039150565b634e487b7160e01b600052603260045260246000fd5b6000600182016155025761550261545d565b5060010190565b8183823760009101908152919050565b60005b8381101561553457818101518382015260200161551c565b50506000910152565b60008151808452615555816020860160208601615519565b601f01601f19169290920160200192915050565b821515815260406020820152600061137c604083018461553d565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b88815260006001600160a01b03808a1660208401528089166040840152508660608301528560808301528460a083015260e060c08301526155f260e083018486615584565b9a9950505050505050505050565b6001600160a01b0386168152846020820152836040820152608060608201526000612568608083018486615584565b60006020828403121561564157600080fd5b5051919050565b60006020828403121561565a57600080fd5b815180151581146111ce57600080fd5b6000808335601e1984360301811261568157600080fd5b83018035915067ffffffffffffffff82111561569c57600080fd5b60200191503681900382131561345f57600080fd5b60006001600160a01b038088168352861515602084015285604084015280851660608401525060a0608083015261256860a083018461553d565b600080604083850312156156fe57600080fd5b505080516020909101519092909150565b6000600160ff1b82036157245761572461545d565b5060000390565b60008261574857634e487b7160e01b600052601260045260246000fd5b500490565b634e487b7160e01b600052602160045260246000fd5b8486823790930191825260208201526040810191909152606001919050565b60008251615794818460208701615519565b919091019291505056fea264697066735822122040321861ce858a2c911db7a2e1f42f4368d23b5251b80dd661a6f2abf19c358d64736f6c63430008110033",
}

// OneInchV5ABI is the input ABI used to generate the binding from.
// Deprecated: Use OneInchV5MetaData.ABI instead.
var OneInchV5ABI = OneInchV5MetaData.ABI

// OneInchV5Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OneInchV5MetaData.Bin instead.
var OneInchV5Bin = OneInchV5MetaData.Bin

// DeployOneInchV5 deploys a new Ethereum contract, binding an instance of OneInchV5 to it.
func DeployOneInchV5(auth *bind.TransactOpts, backend bind.ContractBackend, weth common.Address) (common.Address, *types.Transaction, *OneInchV5, error) {
	parsed, err := OneInchV5MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OneInchV5Bin), backend, weth)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneInchV5{OneInchV5Caller: OneInchV5Caller{contract: contract}, OneInchV5Transactor: OneInchV5Transactor{contract: contract}, OneInchV5Filterer: OneInchV5Filterer{contract: contract}}, nil
}

// OneInchV5 is an auto generated Go binding around an Ethereum contract.
type OneInchV5 struct {
	OneInchV5Caller     // Read-only binding to the contract
	OneInchV5Transactor // Write-only binding to the contract
	OneInchV5Filterer   // Log filterer for contract events
}

// OneInchV5Caller is an auto generated read-only Go binding around an Ethereum contract.
type OneInchV5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchV5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type OneInchV5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchV5Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneInchV5Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchV5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneInchV5Session struct {
	Contract     *OneInchV5        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneInchV5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneInchV5CallerSession struct {
	Contract *OneInchV5Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OneInchV5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneInchV5TransactorSession struct {
	Contract     *OneInchV5Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OneInchV5Raw is an auto generated low-level Go binding around an Ethereum contract.
type OneInchV5Raw struct {
	Contract *OneInchV5 // Generic contract binding to access the raw methods on
}

// OneInchV5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneInchV5CallerRaw struct {
	Contract *OneInchV5Caller // Generic read-only contract binding to access the raw methods on
}

// OneInchV5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneInchV5TransactorRaw struct {
	Contract *OneInchV5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewOneInchV5 creates a new instance of OneInchV5, bound to a specific deployed contract.
func NewOneInchV5(address common.Address, backend bind.ContractBackend) (*OneInchV5, error) {
	contract, err := bindOneInchV5(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneInchV5{OneInchV5Caller: OneInchV5Caller{contract: contract}, OneInchV5Transactor: OneInchV5Transactor{contract: contract}, OneInchV5Filterer: OneInchV5Filterer{contract: contract}}, nil
}

// NewOneInchV5Caller creates a new read-only instance of OneInchV5, bound to a specific deployed contract.
func NewOneInchV5Caller(address common.Address, caller bind.ContractCaller) (*OneInchV5Caller, error) {
	contract, err := bindOneInchV5(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchV5Caller{contract: contract}, nil
}

// NewOneInchV5Transactor creates a new write-only instance of OneInchV5, bound to a specific deployed contract.
func NewOneInchV5Transactor(address common.Address, transactor bind.ContractTransactor) (*OneInchV5Transactor, error) {
	contract, err := bindOneInchV5(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchV5Transactor{contract: contract}, nil
}

// NewOneInchV5Filterer creates a new log filterer instance of OneInchV5, bound to a specific deployed contract.
func NewOneInchV5Filterer(address common.Address, filterer bind.ContractFilterer) (*OneInchV5Filterer, error) {
	contract, err := bindOneInchV5(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneInchV5Filterer{contract: contract}, nil
}

// bindOneInchV5 binds a generic wrapper to an already deployed contract.
func bindOneInchV5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneInchV5ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchV5 *OneInchV5Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchV5.Contract.OneInchV5Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchV5 *OneInchV5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchV5.Contract.OneInchV5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchV5 *OneInchV5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchV5.Contract.OneInchV5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchV5 *OneInchV5CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchV5.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchV5 *OneInchV5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchV5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchV5 *OneInchV5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchV5.Contract.contract.Transact(opts, method, params...)
}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) And(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "and", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Session) And(offsets *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.And(&_OneInchV5.CallOpts, offsets, data)
}

// And is a free data retrieval call binding the contract method 0xbfa75143.
//
// Solidity: function and(uint256 offsets, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) And(offsets *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.And(&_OneInchV5.CallOpts, offsets, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_OneInchV5 *OneInchV5Caller) ArbitraryStaticCall(opts *bind.CallOpts, target common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "arbitraryStaticCall", target, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_OneInchV5 *OneInchV5Session) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _OneInchV5.Contract.ArbitraryStaticCall(&_OneInchV5.CallOpts, target, data)
}

// ArbitraryStaticCall is a free data retrieval call binding the contract method 0xbf15fcd8.
//
// Solidity: function arbitraryStaticCall(address target, bytes data) view returns(uint256)
func (_OneInchV5 *OneInchV5CallerSession) ArbitraryStaticCall(target common.Address, data []byte) (*big.Int, error) {
	return _OneInchV5.Contract.ArbitraryStaticCall(&_OneInchV5.CallOpts, target, data)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) CheckPredicate(opts *bind.CallOpts, order OrderLibOrder) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "checkPredicate", order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_OneInchV5 *OneInchV5Session) CheckPredicate(order OrderLibOrder) (bool, error) {
	return _OneInchV5.Contract.CheckPredicate(&_OneInchV5.CallOpts, order)
}

// CheckPredicate is a free data retrieval call binding the contract method 0x6c838250.
//
// Solidity: function checkPredicate((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) CheckPredicate(order OrderLibOrder) (bool, error) {
	return _OneInchV5.Contract.CheckPredicate(&_OneInchV5.CallOpts, order)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) Eq(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "eq", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Session) Eq(value *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Eq(&_OneInchV5.CallOpts, value, data)
}

// Eq is a free data retrieval call binding the contract method 0x6fe7b0ba.
//
// Solidity: function eq(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) Eq(value *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Eq(&_OneInchV5.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) Gt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "gt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Session) Gt(value *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Gt(&_OneInchV5.CallOpts, value, data)
}

// Gt is a free data retrieval call binding the contract method 0x4f38e2b8.
//
// Solidity: function gt(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) Gt(value *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Gt(&_OneInchV5.CallOpts, value, data)
}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_OneInchV5 *OneInchV5Caller) HashOrder(opts *bind.CallOpts, order OrderLibOrder) ([32]byte, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "hashOrder", order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_OneInchV5 *OneInchV5Session) HashOrder(order OrderLibOrder) ([32]byte, error) {
	return _OneInchV5.Contract.HashOrder(&_OneInchV5.CallOpts, order)
}

// HashOrder is a free data retrieval call binding the contract method 0x37e7316f.
//
// Solidity: function hashOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) view returns(bytes32)
func (_OneInchV5 *OneInchV5CallerSession) HashOrder(order OrderLibOrder) ([32]byte, error) {
	return _OneInchV5.Contract.HashOrder(&_OneInchV5.CallOpts, order)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_OneInchV5 *OneInchV5Caller) InvalidatorForOrderRFQ(opts *bind.CallOpts, maker common.Address, slot *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "invalidatorForOrderRFQ", maker, slot)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_OneInchV5 *OneInchV5Session) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _OneInchV5.Contract.InvalidatorForOrderRFQ(&_OneInchV5.CallOpts, maker, slot)
}

// InvalidatorForOrderRFQ is a free data retrieval call binding the contract method 0x56f16124.
//
// Solidity: function invalidatorForOrderRFQ(address maker, uint256 slot) view returns(uint256)
func (_OneInchV5 *OneInchV5CallerSession) InvalidatorForOrderRFQ(maker common.Address, slot *big.Int) (*big.Int, error) {
	return _OneInchV5.Contract.InvalidatorForOrderRFQ(&_OneInchV5.CallOpts, maker, slot)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) Lt(opts *bind.CallOpts, value *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "lt", value, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Session) Lt(value *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Lt(&_OneInchV5.CallOpts, value, data)
}

// Lt is a free data retrieval call binding the contract method 0xca4ece22.
//
// Solidity: function lt(uint256 value, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) Lt(value *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Lt(&_OneInchV5.CallOpts, value, data)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_OneInchV5 *OneInchV5Caller) Nonce(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "nonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_OneInchV5 *OneInchV5Session) Nonce(arg0 common.Address) (*big.Int, error) {
	return _OneInchV5.Contract.Nonce(&_OneInchV5.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0x70ae92d2.
//
// Solidity: function nonce(address ) view returns(uint256)
func (_OneInchV5 *OneInchV5CallerSession) Nonce(arg0 common.Address) (*big.Int, error) {
	return _OneInchV5.Contract.Nonce(&_OneInchV5.CallOpts, arg0)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) NonceEquals(opts *bind.CallOpts, makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "nonceEquals", makerAddress, makerNonce)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_OneInchV5 *OneInchV5Session) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _OneInchV5.Contract.NonceEquals(&_OneInchV5.CallOpts, makerAddress, makerNonce)
}

// NonceEquals is a free data retrieval call binding the contract method 0xcf6fc6e3.
//
// Solidity: function nonceEquals(address makerAddress, uint256 makerNonce) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) NonceEquals(makerAddress common.Address, makerNonce *big.Int) (bool, error) {
	return _OneInchV5.Contract.NonceEquals(&_OneInchV5.CallOpts, makerAddress, makerNonce)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) Or(opts *bind.CallOpts, offsets *big.Int, data []byte) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "or", offsets, data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5Session) Or(offsets *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Or(&_OneInchV5.CallOpts, offsets, data)
}

// Or is a free data retrieval call binding the contract method 0x74261145.
//
// Solidity: function or(uint256 offsets, bytes data) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) Or(offsets *big.Int, data []byte) (bool, error) {
	return _OneInchV5.Contract.Or(&_OneInchV5.CallOpts, offsets, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchV5 *OneInchV5Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchV5 *OneInchV5Session) Owner() (common.Address, error) {
	return _OneInchV5.Contract.Owner(&_OneInchV5.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchV5 *OneInchV5CallerSession) Owner() (common.Address, error) {
	return _OneInchV5.Contract.Owner(&_OneInchV5.CallOpts)
}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_OneInchV5 *OneInchV5Caller) Remaining(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "remaining", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_OneInchV5 *OneInchV5Session) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _OneInchV5.Contract.Remaining(&_OneInchV5.CallOpts, orderHash)
}

// Remaining is a free data retrieval call binding the contract method 0xbc1ed74c.
//
// Solidity: function remaining(bytes32 orderHash) view returns(uint256)
func (_OneInchV5 *OneInchV5CallerSession) Remaining(orderHash [32]byte) (*big.Int, error) {
	return _OneInchV5.Contract.Remaining(&_OneInchV5.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_OneInchV5 *OneInchV5Caller) RemainingRaw(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "remainingRaw", orderHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_OneInchV5 *OneInchV5Session) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _OneInchV5.Contract.RemainingRaw(&_OneInchV5.CallOpts, orderHash)
}

// RemainingRaw is a free data retrieval call binding the contract method 0x7e54f092.
//
// Solidity: function remainingRaw(bytes32 orderHash) view returns(uint256)
func (_OneInchV5 *OneInchV5CallerSession) RemainingRaw(orderHash [32]byte) (*big.Int, error) {
	return _OneInchV5.Contract.RemainingRaw(&_OneInchV5.CallOpts, orderHash)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_OneInchV5 *OneInchV5Caller) RemainingsRaw(opts *bind.CallOpts, orderHashes [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "remainingsRaw", orderHashes)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_OneInchV5 *OneInchV5Session) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _OneInchV5.Contract.RemainingsRaw(&_OneInchV5.CallOpts, orderHashes)
}

// RemainingsRaw is a free data retrieval call binding the contract method 0x942461bb.
//
// Solidity: function remainingsRaw(bytes32[] orderHashes) view returns(uint256[])
func (_OneInchV5 *OneInchV5CallerSession) RemainingsRaw(orderHashes [][32]byte) ([]*big.Int, error) {
	return _OneInchV5.Contract.RemainingsRaw(&_OneInchV5.CallOpts, orderHashes)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) TimestampBelow(opts *bind.CallOpts, time *big.Int) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "timestampBelow", time)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_OneInchV5 *OneInchV5Session) TimestampBelow(time *big.Int) (bool, error) {
	return _OneInchV5.Contract.TimestampBelow(&_OneInchV5.CallOpts, time)
}

// TimestampBelow is a free data retrieval call binding the contract method 0x63592c2b.
//
// Solidity: function timestampBelow(uint256 time) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) TimestampBelow(time *big.Int) (bool, error) {
	return _OneInchV5.Contract.TimestampBelow(&_OneInchV5.CallOpts, time)
}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_OneInchV5 *OneInchV5Caller) TimestampBelowAndNonceEquals(opts *bind.CallOpts, timeNonceAccount *big.Int) (bool, error) {
	var out []interface{}
	err := _OneInchV5.contract.Call(opts, &out, "timestampBelowAndNonceEquals", timeNonceAccount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_OneInchV5 *OneInchV5Session) TimestampBelowAndNonceEquals(timeNonceAccount *big.Int) (bool, error) {
	return _OneInchV5.Contract.TimestampBelowAndNonceEquals(&_OneInchV5.CallOpts, timeNonceAccount)
}

// TimestampBelowAndNonceEquals is a free data retrieval call binding the contract method 0x2cc2878d.
//
// Solidity: function timestampBelowAndNonceEquals(uint256 timeNonceAccount) view returns(bool)
func (_OneInchV5 *OneInchV5CallerSession) TimestampBelowAndNonceEquals(timeNonceAccount *big.Int) (bool, error) {
	return _OneInchV5.Contract.TimestampBelowAndNonceEquals(&_OneInchV5.CallOpts, timeNonceAccount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_OneInchV5 *OneInchV5Transactor) AdvanceNonce(opts *bind.TransactOpts, amount uint8) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "advanceNonce", amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_OneInchV5 *OneInchV5Session) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _OneInchV5.Contract.AdvanceNonce(&_OneInchV5.TransactOpts, amount)
}

// AdvanceNonce is a paid mutator transaction binding the contract method 0x72c244a8.
//
// Solidity: function advanceNonce(uint8 amount) returns()
func (_OneInchV5 *OneInchV5TransactorSession) AdvanceNonce(amount uint8) (*types.Transaction, error) {
	return _OneInchV5.Contract.AdvanceNonce(&_OneInchV5.TransactOpts, amount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Transactor) CancelOrder(opts *bind.TransactOpts, order OrderLibOrder) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Session) CancelOrder(order OrderLibOrder) (*types.Transaction, error) {
	return _OneInchV5.Contract.CancelOrder(&_OneInchV5.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x2d9a56f6.
//
// Solidity: function cancelOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order) returns(uint256 orderRemaining, bytes32 orderHash)
func (_OneInchV5 *OneInchV5TransactorSession) CancelOrder(order OrderLibOrder) (*types.Transaction, error) {
	return _OneInchV5.Contract.CancelOrder(&_OneInchV5.TransactOpts, order)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_OneInchV5 *OneInchV5Transactor) CancelOrderRFQ(opts *bind.TransactOpts, orderInfo *big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "cancelOrderRFQ", orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_OneInchV5 *OneInchV5Session) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.CancelOrderRFQ(&_OneInchV5.TransactOpts, orderInfo)
}

// CancelOrderRFQ is a paid mutator transaction binding the contract method 0x825caba1.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo) returns()
func (_OneInchV5 *OneInchV5TransactorSession) CancelOrderRFQ(orderInfo *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.CancelOrderRFQ(&_OneInchV5.TransactOpts, orderInfo)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_OneInchV5 *OneInchV5Transactor) CancelOrderRFQ0(opts *bind.TransactOpts, orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "cancelOrderRFQ0", orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_OneInchV5 *OneInchV5Session) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.CancelOrderRFQ0(&_OneInchV5.TransactOpts, orderInfo, additionalMask)
}

// CancelOrderRFQ0 is a paid mutator transaction binding the contract method 0xbddccd35.
//
// Solidity: function cancelOrderRFQ(uint256 orderInfo, uint256 additionalMask) returns()
func (_OneInchV5 *OneInchV5TransactorSession) CancelOrderRFQ0(orderInfo *big.Int, additionalMask *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.CancelOrderRFQ0(&_OneInchV5.TransactOpts, orderInfo, additionalMask)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) ClipperSwap(opts *bind.TransactOpts, clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "clipperSwap", clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) ClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.ClipperSwap(&_OneInchV5.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwap is a paid mutator transaction binding the contract method 0x84bd6d29.
//
// Solidity: function clipperSwap(address clipperExchange, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) ClipperSwap(clipperExchange common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.ClipperSwap(&_OneInchV5.TransactOpts, clipperExchange, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) ClipperSwapTo(opts *bind.TransactOpts, clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "clipperSwapTo", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.ClipperSwapTo(&_OneInchV5.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapTo is a paid mutator transaction binding the contract method 0x093d4fa5.
//
// Solidity: function clipperSwapTo(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) ClipperSwapTo(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.ClipperSwapTo(&_OneInchV5.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) ClipperSwapToWithPermit(opts *bind.TransactOpts, clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "clipperSwapToWithPermit", clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) ClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.ClipperSwapToWithPermit(&_OneInchV5.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// ClipperSwapToWithPermit is a paid mutator transaction binding the contract method 0xc805a666.
//
// Solidity: function clipperSwapToWithPermit(address clipperExchange, address recipient, address srcToken, address dstToken, uint256 inputAmount, uint256 outputAmount, uint256 goodUntil, bytes32 r, bytes32 vs, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) ClipperSwapToWithPermit(clipperExchange common.Address, recipient common.Address, srcToken common.Address, dstToken common.Address, inputAmount *big.Int, outputAmount *big.Int, goodUntil *big.Int, r [32]byte, vs [32]byte, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.ClipperSwapToWithPermit(&_OneInchV5.TransactOpts, clipperExchange, recipient, srcToken, dstToken, inputAmount, outputAmount, goodUntil, r, vs, permit)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OneInchV5 *OneInchV5Transactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OneInchV5 *OneInchV5Session) Destroy() (*types.Transaction, error) {
	return _OneInchV5.Contract.Destroy(&_OneInchV5.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_OneInchV5 *OneInchV5TransactorSession) Destroy() (*types.Transaction, error) {
	return _OneInchV5.Contract.Destroy(&_OneInchV5.TransactOpts)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Transactor) FillOrder(opts *bind.TransactOpts, order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "fillOrder", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Session) FillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrder(&_OneInchV5.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrder is a paid mutator transaction binding the contract method 0x62e238bb.
//
// Solidity: function fillOrder((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount) payable returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5TransactorSession) FillOrder(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrder(&_OneInchV5.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Transactor) FillOrderRFQ(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "fillOrderRFQ", order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Session) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQ(&_OneInchV5.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQ is a paid mutator transaction binding the contract method 0x3eca9c0a.
//
// Solidity: function fillOrderRFQ((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount) payable returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5TransactorSession) FillOrderRFQ(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQ(&_OneInchV5.TransactOpts, order, signature, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Transactor) FillOrderRFQCompact(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "fillOrderRFQCompact", order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Session) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQCompact(&_OneInchV5.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQCompact is a paid mutator transaction binding the contract method 0x9570eeee.
//
// Solidity: function fillOrderRFQCompact((uint256,address,address,address,address,uint256,uint256) order, bytes32 r, bytes32 vs, uint256 flagsAndAmount) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5TransactorSession) FillOrderRFQCompact(order OrderRFQLibOrderRFQ, r [32]byte, vs [32]byte, flagsAndAmount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQCompact(&_OneInchV5.TransactOpts, order, r, vs, flagsAndAmount)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Transactor) FillOrderRFQTo(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "fillOrderRFQTo", order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Session) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQTo(&_OneInchV5.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQTo is a paid mutator transaction binding the contract method 0x5a099843.
//
// Solidity: function fillOrderRFQTo((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target) payable returns(uint256 filledMakingAmount, uint256 filledTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5TransactorSession) FillOrderRFQTo(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQTo(&_OneInchV5.TransactOpts, order, signature, flagsAndAmount, target)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Transactor) FillOrderRFQToWithPermit(opts *bind.TransactOpts, order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "fillOrderRFQToWithPermit", order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Session) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQToWithPermit(&_OneInchV5.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// FillOrderRFQToWithPermit is a paid mutator transaction binding the contract method 0x70ccbd31.
//
// Solidity: function fillOrderRFQToWithPermit((uint256,address,address,address,address,uint256,uint256) order, bytes signature, uint256 flagsAndAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5TransactorSession) FillOrderRFQToWithPermit(order OrderRFQLibOrderRFQ, signature []byte, flagsAndAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderRFQToWithPermit(&_OneInchV5.TransactOpts, order, signature, flagsAndAmount, target, permit)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Transactor) FillOrderTo(opts *bind.TransactOpts, order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "fillOrderTo", order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5Session) FillOrderTo(order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderTo(&_OneInchV5.TransactOpts, order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderTo is a paid mutator transaction binding the contract method 0xe5d7bde6.
//
// Solidity: function fillOrderTo((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order_, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target) payable returns(uint256 actualMakingAmount, uint256 actualTakingAmount, bytes32 orderHash)
func (_OneInchV5 *OneInchV5TransactorSession) FillOrderTo(order_ OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderTo(&_OneInchV5.TransactOpts, order_, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Transactor) FillOrderToWithPermit(opts *bind.TransactOpts, order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "fillOrderToWithPermit", order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5Session) FillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderToWithPermit(&_OneInchV5.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// FillOrderToWithPermit is a paid mutator transaction binding the contract method 0xd365c695.
//
// Solidity: function fillOrderToWithPermit((uint256,address,address,address,address,address,uint256,uint256,uint256,bytes) order, bytes signature, bytes interaction, uint256 makingAmount, uint256 takingAmount, uint256 skipPermitAndThresholdAmount, address target, bytes permit) returns(uint256, uint256, bytes32)
func (_OneInchV5 *OneInchV5TransactorSession) FillOrderToWithPermit(order OrderLibOrder, signature []byte, interaction []byte, makingAmount *big.Int, takingAmount *big.Int, skipPermitAndThresholdAmount *big.Int, target common.Address, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.FillOrderToWithPermit(&_OneInchV5.TransactOpts, order, signature, interaction, makingAmount, takingAmount, skipPermitAndThresholdAmount, target, permit)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_OneInchV5 *OneInchV5Transactor) IncreaseNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "increaseNonce")
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_OneInchV5 *OneInchV5Session) IncreaseNonce() (*types.Transaction, error) {
	return _OneInchV5.Contract.IncreaseNonce(&_OneInchV5.TransactOpts)
}

// IncreaseNonce is a paid mutator transaction binding the contract method 0xc53a0292.
//
// Solidity: function increaseNonce() returns()
func (_OneInchV5 *OneInchV5TransactorSession) IncreaseNonce() (*types.Transaction, error) {
	return _OneInchV5.Contract.IncreaseNonce(&_OneInchV5.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchV5 *OneInchV5Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchV5 *OneInchV5Session) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchV5.Contract.RenounceOwnership(&_OneInchV5.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchV5 *OneInchV5TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchV5.Contract.RenounceOwnership(&_OneInchV5.TransactOpts)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OneInchV5 *OneInchV5Transactor) RescueFunds(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "rescueFunds", token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OneInchV5 *OneInchV5Session) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.RescueFunds(&_OneInchV5.TransactOpts, token, amount)
}

// RescueFunds is a paid mutator transaction binding the contract method 0x78e3214f.
//
// Solidity: function rescueFunds(address token, uint256 amount) returns()
func (_OneInchV5 *OneInchV5TransactorSession) RescueFunds(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.RescueFunds(&_OneInchV5.TransactOpts, token, amount)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_OneInchV5 *OneInchV5Transactor) Simulate(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "simulate", target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_OneInchV5 *OneInchV5Session) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.Simulate(&_OneInchV5.TransactOpts, target, data)
}

// Simulate is a paid mutator transaction binding the contract method 0xbd61951d.
//
// Solidity: function simulate(address target, bytes data) returns()
func (_OneInchV5 *OneInchV5TransactorSession) Simulate(target common.Address, data []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.Simulate(&_OneInchV5.TransactOpts, target, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_OneInchV5 *OneInchV5Transactor) Swap(opts *bind.TransactOpts, executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "swap", executor, desc, permit, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_OneInchV5 *OneInchV5Session) Swap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.Swap(&_OneInchV5.TransactOpts, executor, desc, permit, data)
}

// Swap is a paid mutator transaction binding the contract method 0x12aa3caf.
//
// Solidity: function swap(address executor, (address,address,address,address,uint256,uint256,uint256) desc, bytes permit, bytes data) payable returns(uint256 returnAmount, uint256 spentAmount)
func (_OneInchV5 *OneInchV5TransactorSession) Swap(executor common.Address, desc GenericRouterSwapDescription, permit []byte, data []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.Swap(&_OneInchV5.TransactOpts, executor, desc, permit, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchV5 *OneInchV5Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchV5 *OneInchV5Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchV5.Contract.TransferOwnership(&_OneInchV5.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchV5 *OneInchV5TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchV5.Contract.TransferOwnership(&_OneInchV5.TransactOpts, newOwner)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) UniswapV3Swap(opts *bind.TransactOpts, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "uniswapV3Swap", amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3Swap(&_OneInchV5.TransactOpts, amount, minReturn, pools)
}

// UniswapV3Swap is a paid mutator transaction binding the contract method 0xe449022e.
//
// Solidity: function uniswapV3Swap(uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) UniswapV3Swap(amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3Swap(&_OneInchV5.TransactOpts, amount, minReturn, pools)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_OneInchV5 *OneInchV5Transactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_OneInchV5 *OneInchV5Session) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3SwapCallback(&_OneInchV5.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes ) returns()
func (_OneInchV5 *OneInchV5TransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3SwapCallback(&_OneInchV5.TransactOpts, amount0Delta, amount1Delta, arg2)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) UniswapV3SwapTo(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "uniswapV3SwapTo", recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3SwapTo(&_OneInchV5.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapTo is a paid mutator transaction binding the contract method 0xbc80f1a8.
//
// Solidity: function uniswapV3SwapTo(address recipient, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) UniswapV3SwapTo(recipient common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3SwapTo(&_OneInchV5.TransactOpts, recipient, amount, minReturn, pools)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) UniswapV3SwapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "uniswapV3SwapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3SwapToWithPermit(&_OneInchV5.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UniswapV3SwapToWithPermit is a paid mutator transaction binding the contract method 0x2521b930.
//
// Solidity: function uniswapV3SwapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) UniswapV3SwapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.UniswapV3SwapToWithPermit(&_OneInchV5.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) Unoswap(opts *bind.TransactOpts, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "unoswap", srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.Unoswap(&_OneInchV5.TransactOpts, srcToken, amount, minReturn, pools)
}

// Unoswap is a paid mutator transaction binding the contract method 0x0502b1c5.
//
// Solidity: function unoswap(address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) Unoswap(srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.Unoswap(&_OneInchV5.TransactOpts, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) UnoswapTo(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "unoswapTo", recipient, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) UnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.UnoswapTo(&_OneInchV5.TransactOpts, recipient, srcToken, amount, minReturn, pools)
}

// UnoswapTo is a paid mutator transaction binding the contract method 0xf78dc253.
//
// Solidity: function unoswapTo(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools) payable returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) UnoswapTo(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int) (*types.Transaction, error) {
	return _OneInchV5.Contract.UnoswapTo(&_OneInchV5.TransactOpts, recipient, srcToken, amount, minReturn, pools)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Transactor) UnoswapToWithPermit(opts *bind.TransactOpts, recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.contract.Transact(opts, "unoswapToWithPermit", recipient, srcToken, amount, minReturn, pools, permit)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5Session) UnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.UnoswapToWithPermit(&_OneInchV5.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// UnoswapToWithPermit is a paid mutator transaction binding the contract method 0x3c15fd91.
//
// Solidity: function unoswapToWithPermit(address recipient, address srcToken, uint256 amount, uint256 minReturn, uint256[] pools, bytes permit) returns(uint256 returnAmount)
func (_OneInchV5 *OneInchV5TransactorSession) UnoswapToWithPermit(recipient common.Address, srcToken common.Address, amount *big.Int, minReturn *big.Int, pools []*big.Int, permit []byte) (*types.Transaction, error) {
	return _OneInchV5.Contract.UnoswapToWithPermit(&_OneInchV5.TransactOpts, recipient, srcToken, amount, minReturn, pools, permit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OneInchV5 *OneInchV5Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchV5.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OneInchV5 *OneInchV5Session) Receive() (*types.Transaction, error) {
	return _OneInchV5.Contract.Receive(&_OneInchV5.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OneInchV5 *OneInchV5TransactorSession) Receive() (*types.Transaction, error) {
	return _OneInchV5.Contract.Receive(&_OneInchV5.TransactOpts)
}

// OneInchV5NonceIncreasedIterator is returned from FilterNonceIncreased and is used to iterate over the raw logs and unpacked data for NonceIncreased events raised by the OneInchV5 contract.
type OneInchV5NonceIncreasedIterator struct {
	Event *OneInchV5NonceIncreased // Event containing the contract specifics and raw log

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
func (it *OneInchV5NonceIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchV5NonceIncreased)
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
		it.Event = new(OneInchV5NonceIncreased)
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
func (it *OneInchV5NonceIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchV5NonceIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchV5NonceIncreased represents a NonceIncreased event raised by the OneInchV5 contract.
type OneInchV5NonceIncreased struct {
	Maker    common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncreased is a free log retrieval operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_OneInchV5 *OneInchV5Filterer) FilterNonceIncreased(opts *bind.FilterOpts, maker []common.Address) (*OneInchV5NonceIncreasedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OneInchV5.contract.FilterLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchV5NonceIncreasedIterator{contract: _OneInchV5.contract, event: "NonceIncreased", logs: logs, sub: sub}, nil
}

// WatchNonceIncreased is a free log subscription operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_OneInchV5 *OneInchV5Filterer) WatchNonceIncreased(opts *bind.WatchOpts, sink chan<- *OneInchV5NonceIncreased, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OneInchV5.contract.WatchLogs(opts, "NonceIncreased", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchV5NonceIncreased)
				if err := _OneInchV5.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
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

// ParseNonceIncreased is a log parse operation binding the contract event 0xfc69110dd11eb791755e4abd6b7d281bae236de95736d38a23782814be5e10db.
//
// Solidity: event NonceIncreased(address indexed maker, uint256 newNonce)
func (_OneInchV5 *OneInchV5Filterer) ParseNonceIncreased(log types.Log) (*OneInchV5NonceIncreased, error) {
	event := new(OneInchV5NonceIncreased)
	if err := _OneInchV5.contract.UnpackLog(event, "NonceIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchV5OrderCanceledIterator is returned from FilterOrderCanceled and is used to iterate over the raw logs and unpacked data for OrderCanceled events raised by the OneInchV5 contract.
type OneInchV5OrderCanceledIterator struct {
	Event *OneInchV5OrderCanceled // Event containing the contract specifics and raw log

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
func (it *OneInchV5OrderCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchV5OrderCanceled)
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
		it.Event = new(OneInchV5OrderCanceled)
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
func (it *OneInchV5OrderCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchV5OrderCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchV5OrderCanceled represents a OrderCanceled event raised by the OneInchV5 contract.
type OneInchV5OrderCanceled struct {
	Maker        common.Address
	OrderHash    [32]byte
	RemainingRaw *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderCanceled is a free log retrieval operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_OneInchV5 *OneInchV5Filterer) FilterOrderCanceled(opts *bind.FilterOpts, maker []common.Address) (*OneInchV5OrderCanceledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OneInchV5.contract.FilterLogs(opts, "OrderCanceled", makerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchV5OrderCanceledIterator{contract: _OneInchV5.contract, event: "OrderCanceled", logs: logs, sub: sub}, nil
}

// WatchOrderCanceled is a free log subscription operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_OneInchV5 *OneInchV5Filterer) WatchOrderCanceled(opts *bind.WatchOpts, sink chan<- *OneInchV5OrderCanceled, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OneInchV5.contract.WatchLogs(opts, "OrderCanceled", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchV5OrderCanceled)
				if err := _OneInchV5.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
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

// ParseOrderCanceled is a log parse operation binding the contract event 0xcbfa7d191838ece7ba4783ca3a30afd316619b7f368094b57ee7ffde9a923db1.
//
// Solidity: event OrderCanceled(address indexed maker, bytes32 orderHash, uint256 remainingRaw)
func (_OneInchV5 *OneInchV5Filterer) ParseOrderCanceled(log types.Log) (*OneInchV5OrderCanceled, error) {
	event := new(OneInchV5OrderCanceled)
	if err := _OneInchV5.contract.UnpackLog(event, "OrderCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchV5OrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the OneInchV5 contract.
type OneInchV5OrderFilledIterator struct {
	Event *OneInchV5OrderFilled // Event containing the contract specifics and raw log

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
func (it *OneInchV5OrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchV5OrderFilled)
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
		it.Event = new(OneInchV5OrderFilled)
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
func (it *OneInchV5OrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchV5OrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchV5OrderFilled represents a OrderFilled event raised by the OneInchV5 contract.
type OneInchV5OrderFilled struct {
	Maker     common.Address
	OrderHash [32]byte
	Remaining *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_OneInchV5 *OneInchV5Filterer) FilterOrderFilled(opts *bind.FilterOpts, maker []common.Address) (*OneInchV5OrderFilledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OneInchV5.contract.FilterLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchV5OrderFilledIterator{contract: _OneInchV5.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_OneInchV5 *OneInchV5Filterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *OneInchV5OrderFilled, maker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	logs, sub, err := _OneInchV5.contract.WatchLogs(opts, "OrderFilled", makerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchV5OrderFilled)
				if err := _OneInchV5.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0xb9ed0243fdf00f0545c63a0af8850c090d86bb46682baec4bf3c496814fe4f02.
//
// Solidity: event OrderFilled(address indexed maker, bytes32 orderHash, uint256 remaining)
func (_OneInchV5 *OneInchV5Filterer) ParseOrderFilled(log types.Log) (*OneInchV5OrderFilled, error) {
	event := new(OneInchV5OrderFilled)
	if err := _OneInchV5.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchV5OrderFilledRFQIterator is returned from FilterOrderFilledRFQ and is used to iterate over the raw logs and unpacked data for OrderFilledRFQ events raised by the OneInchV5 contract.
type OneInchV5OrderFilledRFQIterator struct {
	Event *OneInchV5OrderFilledRFQ // Event containing the contract specifics and raw log

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
func (it *OneInchV5OrderFilledRFQIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchV5OrderFilledRFQ)
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
		it.Event = new(OneInchV5OrderFilledRFQ)
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
func (it *OneInchV5OrderFilledRFQIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchV5OrderFilledRFQIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchV5OrderFilledRFQ represents a OrderFilledRFQ event raised by the OneInchV5 contract.
type OneInchV5OrderFilledRFQ struct {
	OrderHash    [32]byte
	MakingAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOrderFilledRFQ is a free log retrieval operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_OneInchV5 *OneInchV5Filterer) FilterOrderFilledRFQ(opts *bind.FilterOpts) (*OneInchV5OrderFilledRFQIterator, error) {

	logs, sub, err := _OneInchV5.contract.FilterLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return &OneInchV5OrderFilledRFQIterator{contract: _OneInchV5.contract, event: "OrderFilledRFQ", logs: logs, sub: sub}, nil
}

// WatchOrderFilledRFQ is a free log subscription operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_OneInchV5 *OneInchV5Filterer) WatchOrderFilledRFQ(opts *bind.WatchOpts, sink chan<- *OneInchV5OrderFilledRFQ) (event.Subscription, error) {

	logs, sub, err := _OneInchV5.contract.WatchLogs(opts, "OrderFilledRFQ")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchV5OrderFilledRFQ)
				if err := _OneInchV5.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
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

// ParseOrderFilledRFQ is a log parse operation binding the contract event 0xc3b639f02b125bfa160e50739b8c44eb2d1b6908e2b6d5925c6d770f2ca78127.
//
// Solidity: event OrderFilledRFQ(bytes32 orderHash, uint256 makingAmount)
func (_OneInchV5 *OneInchV5Filterer) ParseOrderFilledRFQ(log types.Log) (*OneInchV5OrderFilledRFQ, error) {
	event := new(OneInchV5OrderFilledRFQ)
	if err := _OneInchV5.contract.UnpackLog(event, "OrderFilledRFQ", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchV5OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OneInchV5 contract.
type OneInchV5OwnershipTransferredIterator struct {
	Event *OneInchV5OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneInchV5OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchV5OwnershipTransferred)
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
		it.Event = new(OneInchV5OwnershipTransferred)
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
func (it *OneInchV5OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchV5OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchV5OwnershipTransferred represents a OwnershipTransferred event raised by the OneInchV5 contract.
type OneInchV5OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchV5 *OneInchV5Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneInchV5OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchV5.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchV5OwnershipTransferredIterator{contract: _OneInchV5.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchV5 *OneInchV5Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneInchV5OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchV5.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchV5OwnershipTransferred)
				if err := _OneInchV5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OneInchV5 *OneInchV5Filterer) ParseOwnershipTransferred(log types.Log) (*OneInchV5OwnershipTransferred, error) {
	event := new(OneInchV5OwnershipTransferred)
	if err := _OneInchV5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
