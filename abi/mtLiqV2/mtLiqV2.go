// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mtLiqV2

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

// PoolLiq is an auto generated low-level Go binding around an user-defined struct.
type PoolLiq struct {
	Swap    common.Address
	Version uint8
	Pool    common.Address
	Fee     *big.Int
	Liq     *big.Int
}

// MtLiqV2MetaData contains all meta data concerning the MtLiqV2 contract.
var MtLiqV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swap\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"ver\",\"type\":\"uint8\"}],\"name\":\"addSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"liquidity\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"swap\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"liq\",\"type\":\"uint256\"}],\"internalType\":\"structPoolLiq[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swap\",\"type\":\"address\"}],\"name\":\"removeSwap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"viewSwap\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MtLiqV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use MtLiqV2MetaData.ABI instead.
var MtLiqV2ABI = MtLiqV2MetaData.ABI

// MtLiqV2 is an auto generated Go binding around an Ethereum contract.
type MtLiqV2 struct {
	MtLiqV2Caller     // Read-only binding to the contract
	MtLiqV2Transactor // Write-only binding to the contract
	MtLiqV2Filterer   // Log filterer for contract events
}

// MtLiqV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type MtLiqV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MtLiqV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MtLiqV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MtLiqV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MtLiqV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MtLiqV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MtLiqV2Session struct {
	Contract     *MtLiqV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MtLiqV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MtLiqV2CallerSession struct {
	Contract *MtLiqV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MtLiqV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MtLiqV2TransactorSession struct {
	Contract     *MtLiqV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MtLiqV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type MtLiqV2Raw struct {
	Contract *MtLiqV2 // Generic contract binding to access the raw methods on
}

// MtLiqV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MtLiqV2CallerRaw struct {
	Contract *MtLiqV2Caller // Generic read-only contract binding to access the raw methods on
}

// MtLiqV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MtLiqV2TransactorRaw struct {
	Contract *MtLiqV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMtLiqV2 creates a new instance of MtLiqV2, bound to a specific deployed contract.
func NewMtLiqV2(address common.Address, backend bind.ContractBackend) (*MtLiqV2, error) {
	contract, err := bindMtLiqV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MtLiqV2{MtLiqV2Caller: MtLiqV2Caller{contract: contract}, MtLiqV2Transactor: MtLiqV2Transactor{contract: contract}, MtLiqV2Filterer: MtLiqV2Filterer{contract: contract}}, nil
}

// NewMtLiqV2Caller creates a new read-only instance of MtLiqV2, bound to a specific deployed contract.
func NewMtLiqV2Caller(address common.Address, caller bind.ContractCaller) (*MtLiqV2Caller, error) {
	contract, err := bindMtLiqV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MtLiqV2Caller{contract: contract}, nil
}

// NewMtLiqV2Transactor creates a new write-only instance of MtLiqV2, bound to a specific deployed contract.
func NewMtLiqV2Transactor(address common.Address, transactor bind.ContractTransactor) (*MtLiqV2Transactor, error) {
	contract, err := bindMtLiqV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MtLiqV2Transactor{contract: contract}, nil
}

// NewMtLiqV2Filterer creates a new log filterer instance of MtLiqV2, bound to a specific deployed contract.
func NewMtLiqV2Filterer(address common.Address, filterer bind.ContractFilterer) (*MtLiqV2Filterer, error) {
	contract, err := bindMtLiqV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MtLiqV2Filterer{contract: contract}, nil
}

// bindMtLiqV2 binds a generic wrapper to an already deployed contract.
func bindMtLiqV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MtLiqV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MtLiqV2 *MtLiqV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MtLiqV2.Contract.MtLiqV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MtLiqV2 *MtLiqV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MtLiqV2.Contract.MtLiqV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MtLiqV2 *MtLiqV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MtLiqV2.Contract.MtLiqV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MtLiqV2 *MtLiqV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MtLiqV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MtLiqV2 *MtLiqV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MtLiqV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MtLiqV2 *MtLiqV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MtLiqV2.Contract.contract.Transact(opts, method, params...)
}

// Liquidity is a free data retrieval call binding the contract method 0xb2217281.
//
// Solidity: function liquidity(address tokenA, address tokenB) view returns((address,uint8,address,uint24,uint256)[])
func (_MtLiqV2 *MtLiqV2Caller) Liquidity(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) ([]PoolLiq, error) {
	var out []interface{}
	err := _MtLiqV2.contract.Call(opts, &out, "liquidity", tokenA, tokenB)

	if err != nil {
		return *new([]PoolLiq), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolLiq)).(*[]PoolLiq)

	return out0, err

}

// Liquidity is a free data retrieval call binding the contract method 0xb2217281.
//
// Solidity: function liquidity(address tokenA, address tokenB) view returns((address,uint8,address,uint24,uint256)[])
func (_MtLiqV2 *MtLiqV2Session) Liquidity(tokenA common.Address, tokenB common.Address) ([]PoolLiq, error) {
	return _MtLiqV2.Contract.Liquidity(&_MtLiqV2.CallOpts, tokenA, tokenB)
}

// Liquidity is a free data retrieval call binding the contract method 0xb2217281.
//
// Solidity: function liquidity(address tokenA, address tokenB) view returns((address,uint8,address,uint24,uint256)[])
func (_MtLiqV2 *MtLiqV2CallerSession) Liquidity(tokenA common.Address, tokenB common.Address) ([]PoolLiq, error) {
	return _MtLiqV2.Contract.Liquidity(&_MtLiqV2.CallOpts, tokenA, tokenB)
}

// ViewSwap is a free data retrieval call binding the contract method 0x27ff1f0b.
//
// Solidity: function viewSwap() view returns(address[], uint8[])
func (_MtLiqV2 *MtLiqV2Caller) ViewSwap(opts *bind.CallOpts) ([]common.Address, []uint8, error) {
	var out []interface{}
	err := _MtLiqV2.contract.Call(opts, &out, "viewSwap")

	if err != nil {
		return *new([]common.Address), *new([]uint8), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return out0, out1, err

}

// ViewSwap is a free data retrieval call binding the contract method 0x27ff1f0b.
//
// Solidity: function viewSwap() view returns(address[], uint8[])
func (_MtLiqV2 *MtLiqV2Session) ViewSwap() ([]common.Address, []uint8, error) {
	return _MtLiqV2.Contract.ViewSwap(&_MtLiqV2.CallOpts)
}

// ViewSwap is a free data retrieval call binding the contract method 0x27ff1f0b.
//
// Solidity: function viewSwap() view returns(address[], uint8[])
func (_MtLiqV2 *MtLiqV2CallerSession) ViewSwap() ([]common.Address, []uint8, error) {
	return _MtLiqV2.Contract.ViewSwap(&_MtLiqV2.CallOpts)
}

// AddSwap is a paid mutator transaction binding the contract method 0xd84a9ad4.
//
// Solidity: function addSwap(address swap, uint8 ver) returns()
func (_MtLiqV2 *MtLiqV2Transactor) AddSwap(opts *bind.TransactOpts, swap common.Address, ver uint8) (*types.Transaction, error) {
	return _MtLiqV2.contract.Transact(opts, "addSwap", swap, ver)
}

// AddSwap is a paid mutator transaction binding the contract method 0xd84a9ad4.
//
// Solidity: function addSwap(address swap, uint8 ver) returns()
func (_MtLiqV2 *MtLiqV2Session) AddSwap(swap common.Address, ver uint8) (*types.Transaction, error) {
	return _MtLiqV2.Contract.AddSwap(&_MtLiqV2.TransactOpts, swap, ver)
}

// AddSwap is a paid mutator transaction binding the contract method 0xd84a9ad4.
//
// Solidity: function addSwap(address swap, uint8 ver) returns()
func (_MtLiqV2 *MtLiqV2TransactorSession) AddSwap(swap common.Address, ver uint8) (*types.Transaction, error) {
	return _MtLiqV2.Contract.AddSwap(&_MtLiqV2.TransactOpts, swap, ver)
}

// RemoveSwap is a paid mutator transaction binding the contract method 0x92c36b17.
//
// Solidity: function removeSwap(address swap) returns()
func (_MtLiqV2 *MtLiqV2Transactor) RemoveSwap(opts *bind.TransactOpts, swap common.Address) (*types.Transaction, error) {
	return _MtLiqV2.contract.Transact(opts, "removeSwap", swap)
}

// RemoveSwap is a paid mutator transaction binding the contract method 0x92c36b17.
//
// Solidity: function removeSwap(address swap) returns()
func (_MtLiqV2 *MtLiqV2Session) RemoveSwap(swap common.Address) (*types.Transaction, error) {
	return _MtLiqV2.Contract.RemoveSwap(&_MtLiqV2.TransactOpts, swap)
}

// RemoveSwap is a paid mutator transaction binding the contract method 0x92c36b17.
//
// Solidity: function removeSwap(address swap) returns()
func (_MtLiqV2 *MtLiqV2TransactorSession) RemoveSwap(swap common.Address) (*types.Transaction, error) {
	return _MtLiqV2.Contract.RemoveSwap(&_MtLiqV2.TransactOpts, swap)
}
