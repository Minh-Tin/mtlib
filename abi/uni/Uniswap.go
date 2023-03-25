// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uni

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

// UniMetaData contains all meta data concerning the Uni contract.
var UniMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniABI is the input ABI used to generate the binding from.
// Deprecated: Use UniMetaData.ABI instead.
var UniABI = UniMetaData.ABI

// Uni is an auto generated Go binding around an Ethereum contract.
type Uni struct {
	UniCaller     // Read-only binding to the contract
	UniTransactor // Write-only binding to the contract
	UniFilterer   // Log filterer for contract events
}

// UniCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniSession struct {
	Contract     *Uni              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniCallerSession struct {
	Contract *UniCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UniTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniTransactorSession struct {
	Contract     *UniTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniRaw struct {
	Contract *Uni // Generic contract binding to access the raw methods on
}

// UniCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniCallerRaw struct {
	Contract *UniCaller // Generic read-only contract binding to access the raw methods on
}

// UniTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniTransactorRaw struct {
	Contract *UniTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUni creates a new instance of Uni, bound to a specific deployed contract.
func NewUni(address common.Address, backend bind.ContractBackend) (*Uni, error) {
	contract, err := bindUni(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uni{UniCaller: UniCaller{contract: contract}, UniTransactor: UniTransactor{contract: contract}, UniFilterer: UniFilterer{contract: contract}}, nil
}

// NewUniCaller creates a new read-only instance of Uni, bound to a specific deployed contract.
func NewUniCaller(address common.Address, caller bind.ContractCaller) (*UniCaller, error) {
	contract, err := bindUni(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniCaller{contract: contract}, nil
}

// NewUniTransactor creates a new write-only instance of Uni, bound to a specific deployed contract.
func NewUniTransactor(address common.Address, transactor bind.ContractTransactor) (*UniTransactor, error) {
	contract, err := bindUni(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniTransactor{contract: contract}, nil
}

// NewUniFilterer creates a new log filterer instance of Uni, bound to a specific deployed contract.
func NewUniFilterer(address common.Address, filterer bind.ContractFilterer) (*UniFilterer, error) {
	contract, err := bindUni(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniFilterer{contract: contract}, nil
}

// bindUni binds a generic wrapper to an already deployed contract.
func bindUni(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uni *UniRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uni.Contract.UniCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uni *UniRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uni.Contract.UniTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uni *UniRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uni.Contract.UniTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uni *UniCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uni.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uni *UniTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uni.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uni *UniTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uni.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uni *UniCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uni.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uni *UniSession) Factory() (common.Address, error) {
	return _Uni.Contract.Factory(&_Uni.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Uni *UniCallerSession) Factory() (common.Address, error) {
	return _Uni.Contract.Factory(&_Uni.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uni *UniCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uni.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uni *UniSession) Token0() (common.Address, error) {
	return _Uni.Contract.Token0(&_Uni.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Uni *UniCallerSession) Token0() (common.Address, error) {
	return _Uni.Contract.Token0(&_Uni.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uni *UniCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uni.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uni *UniSession) Token1() (common.Address, error) {
	return _Uni.Contract.Token1(&_Uni.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Uni *UniCallerSession) Token1() (common.Address, error) {
	return _Uni.Contract.Token1(&_Uni.CallOpts)
}
