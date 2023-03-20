// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Erc20Bytes32

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

// Erc20Bytes32MetaData contains all meta data concerning the Erc20Bytes32 contract.
var Erc20Bytes32MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Erc20Bytes32ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20Bytes32MetaData.ABI instead.
var Erc20Bytes32ABI = Erc20Bytes32MetaData.ABI

// Erc20Bytes32 is an auto generated Go binding around an Ethereum contract.
type Erc20Bytes32 struct {
	Erc20Bytes32Caller     // Read-only binding to the contract
	Erc20Bytes32Transactor // Write-only binding to the contract
	Erc20Bytes32Filterer   // Log filterer for contract events
}

// Erc20Bytes32Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Bytes32Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Bytes32Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Bytes32Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Bytes32Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Bytes32Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Bytes32Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Bytes32Session struct {
	Contract     *Erc20Bytes32     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Bytes32CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20Bytes32CallerSession struct {
	Contract *Erc20Bytes32Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Erc20Bytes32TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20Bytes32TransactorSession struct {
	Contract     *Erc20Bytes32Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Erc20Bytes32Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Bytes32Raw struct {
	Contract *Erc20Bytes32 // Generic contract binding to access the raw methods on
}

// Erc20Bytes32CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20Bytes32CallerRaw struct {
	Contract *Erc20Bytes32Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20Bytes32TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20Bytes32TransactorRaw struct {
	Contract *Erc20Bytes32Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20Bytes32 creates a new instance of Erc20Bytes32, bound to a specific deployed contract.
func NewErc20Bytes32(address common.Address, backend bind.ContractBackend) (*Erc20Bytes32, error) {
	contract, err := bindErc20Bytes32(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20Bytes32{Erc20Bytes32Caller: Erc20Bytes32Caller{contract: contract}, Erc20Bytes32Transactor: Erc20Bytes32Transactor{contract: contract}, Erc20Bytes32Filterer: Erc20Bytes32Filterer{contract: contract}}, nil
}

// NewErc20Bytes32Caller creates a new read-only instance of Erc20Bytes32, bound to a specific deployed contract.
func NewErc20Bytes32Caller(address common.Address, caller bind.ContractCaller) (*Erc20Bytes32Caller, error) {
	contract, err := bindErc20Bytes32(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Bytes32Caller{contract: contract}, nil
}

// NewErc20Bytes32Transactor creates a new write-only instance of Erc20Bytes32, bound to a specific deployed contract.
func NewErc20Bytes32Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Bytes32Transactor, error) {
	contract, err := bindErc20Bytes32(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Bytes32Transactor{contract: contract}, nil
}

// NewErc20Bytes32Filterer creates a new log filterer instance of Erc20Bytes32, bound to a specific deployed contract.
func NewErc20Bytes32Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Bytes32Filterer, error) {
	contract, err := bindErc20Bytes32(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Bytes32Filterer{contract: contract}, nil
}

// bindErc20Bytes32 binds a generic wrapper to an already deployed contract.
func bindErc20Bytes32(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20Bytes32ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Bytes32 *Erc20Bytes32Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Bytes32.Contract.Erc20Bytes32Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Bytes32 *Erc20Bytes32Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.Erc20Bytes32Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Bytes32 *Erc20Bytes32Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.Erc20Bytes32Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20Bytes32 *Erc20Bytes32CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20Bytes32.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20Bytes32 *Erc20Bytes32TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20Bytes32 *Erc20Bytes32TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Bytes32.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20Bytes32.Contract.Allowance(&_Erc20Bytes32.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20Bytes32.Contract.Allowance(&_Erc20Bytes32.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Bytes32.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20Bytes32.Contract.BalanceOf(&_Erc20Bytes32.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Erc20Bytes32.Contract.BalanceOf(&_Erc20Bytes32.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Bytes32 *Erc20Bytes32Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20Bytes32.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Bytes32 *Erc20Bytes32Session) Decimals() (uint8, error) {
	return _Erc20Bytes32.Contract.Decimals(&_Erc20Bytes32.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20Bytes32 *Erc20Bytes32CallerSession) Decimals() (uint8, error) {
	return _Erc20Bytes32.Contract.Decimals(&_Erc20Bytes32.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20Bytes32 *Erc20Bytes32Caller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20Bytes32.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20Bytes32 *Erc20Bytes32Session) Name() ([32]byte, error) {
	return _Erc20Bytes32.Contract.Name(&_Erc20Bytes32.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(bytes32)
func (_Erc20Bytes32 *Erc20Bytes32CallerSession) Name() ([32]byte, error) {
	return _Erc20Bytes32.Contract.Name(&_Erc20Bytes32.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Erc20Bytes32 *Erc20Bytes32Caller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Erc20Bytes32.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Erc20Bytes32 *Erc20Bytes32Session) Symbol() ([32]byte, error) {
	return _Erc20Bytes32.Contract.Symbol(&_Erc20Bytes32.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(bytes32)
func (_Erc20Bytes32 *Erc20Bytes32CallerSession) Symbol() ([32]byte, error) {
	return _Erc20Bytes32.Contract.Symbol(&_Erc20Bytes32.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20Bytes32.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32Session) TotalSupply() (*big.Int, error) {
	return _Erc20Bytes32.Contract.TotalSupply(&_Erc20Bytes32.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Erc20Bytes32 *Erc20Bytes32CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20Bytes32.Contract.TotalSupply(&_Erc20Bytes32.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.Approve(&_Erc20Bytes32.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.Approve(&_Erc20Bytes32.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.Transfer(&_Erc20Bytes32.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.Transfer(&_Erc20Bytes32.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.TransferFrom(&_Erc20Bytes32.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20Bytes32 *Erc20Bytes32TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20Bytes32.Contract.TransferFrom(&_Erc20Bytes32.TransactOpts, from, to, value)
}

// Erc20Bytes32ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20Bytes32 contract.
type Erc20Bytes32ApprovalIterator struct {
	Event *Erc20Bytes32Approval // Event containing the contract specifics and raw log

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
func (it *Erc20Bytes32ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Bytes32Approval)
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
		it.Event = new(Erc20Bytes32Approval)
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
func (it *Erc20Bytes32ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20Bytes32ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Bytes32Approval represents a Approval event raised by the Erc20Bytes32 contract.
type Erc20Bytes32Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Bytes32 *Erc20Bytes32Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20Bytes32ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Bytes32.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20Bytes32ApprovalIterator{contract: _Erc20Bytes32.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20Bytes32 *Erc20Bytes32Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Bytes32Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20Bytes32.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Bytes32Approval)
				if err := _Erc20Bytes32.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Erc20Bytes32 *Erc20Bytes32Filterer) ParseApproval(log types.Log) (*Erc20Bytes32Approval, error) {
	event := new(Erc20Bytes32Approval)
	if err := _Erc20Bytes32.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20Bytes32TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20Bytes32 contract.
type Erc20Bytes32TransferIterator struct {
	Event *Erc20Bytes32Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20Bytes32TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Bytes32Transfer)
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
		it.Event = new(Erc20Bytes32Transfer)
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
func (it *Erc20Bytes32TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20Bytes32TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Bytes32Transfer represents a Transfer event raised by the Erc20Bytes32 contract.
type Erc20Bytes32Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Bytes32 *Erc20Bytes32Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20Bytes32TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Bytes32.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20Bytes32TransferIterator{contract: _Erc20Bytes32.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20Bytes32 *Erc20Bytes32Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Bytes32Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20Bytes32.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Bytes32Transfer)
				if err := _Erc20Bytes32.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Erc20Bytes32 *Erc20Bytes32Filterer) ParseTransfer(log types.Log) (*Erc20Bytes32Transfer, error) {
	event := new(Erc20Bytes32Transfer)
	if err := _Erc20Bytes32.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
