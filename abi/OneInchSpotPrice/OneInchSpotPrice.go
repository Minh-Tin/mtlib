// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OneInchSpotPrice

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

// OneInchSpotPriceMetaData contains all meta data concerning the OneInchSpotPrice contract.
var OneInchSpotPriceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"_multiWrapper\",\"type\":\"address\"},{\"internalType\":\"contractIOracle[]\",\"name\":\"existingOracles\",\"type\":\"address[]\"},{\"internalType\":\"enumOffchainOracle.OracleType[]\",\"name\":\"oracleTypes\",\"type\":\"uint8[]\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"existingConnectors\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"wBase\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"ConnectorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"ConnectorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractMultiWrapper\",\"name\":\"multiWrapper\",\"type\":\"address\"}],\"name\":\"MultiWrapperUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleType\",\"type\":\"uint8\"}],\"name\":\"OracleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleType\",\"type\":\"uint8\"}],\"name\":\"OracleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"addConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"addOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"connectors\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"allConnectors\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"dstToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useWrappers\",\"type\":\"bool\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useSrcWrappers\",\"type\":\"bool\"}],\"name\":\"getRateToEth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weightedRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multiWrapper\",\"outputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracles\",\"outputs\":[{\"internalType\":\"contractIOracle[]\",\"name\":\"allOracles\",\"type\":\"address[]\"},{\"internalType\":\"enumOffchainOracle.OracleType[]\",\"name\":\"oracleTypes\",\"type\":\"uint8[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"connector\",\"type\":\"address\"}],\"name\":\"removeConnector\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"enumOffchainOracle.OracleType\",\"name\":\"oracleKind\",\"type\":\"uint8\"}],\"name\":\"removeOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractMultiWrapper\",\"name\":\"_multiWrapper\",\"type\":\"address\"}],\"name\":\"setMultiWrapper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OneInchSpotPriceABI is the input ABI used to generate the binding from.
// Deprecated: Use OneInchSpotPriceMetaData.ABI instead.
var OneInchSpotPriceABI = OneInchSpotPriceMetaData.ABI

// OneInchSpotPrice is an auto generated Go binding around an Ethereum contract.
type OneInchSpotPrice struct {
	OneInchSpotPriceCaller     // Read-only binding to the contract
	OneInchSpotPriceTransactor // Write-only binding to the contract
	OneInchSpotPriceFilterer   // Log filterer for contract events
}

// OneInchSpotPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneInchSpotPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchSpotPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneInchSpotPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchSpotPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneInchSpotPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneInchSpotPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneInchSpotPriceSession struct {
	Contract     *OneInchSpotPrice // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneInchSpotPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneInchSpotPriceCallerSession struct {
	Contract *OneInchSpotPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OneInchSpotPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneInchSpotPriceTransactorSession struct {
	Contract     *OneInchSpotPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OneInchSpotPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneInchSpotPriceRaw struct {
	Contract *OneInchSpotPrice // Generic contract binding to access the raw methods on
}

// OneInchSpotPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneInchSpotPriceCallerRaw struct {
	Contract *OneInchSpotPriceCaller // Generic read-only contract binding to access the raw methods on
}

// OneInchSpotPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneInchSpotPriceTransactorRaw struct {
	Contract *OneInchSpotPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneInchSpotPrice creates a new instance of OneInchSpotPrice, bound to a specific deployed contract.
func NewOneInchSpotPrice(address common.Address, backend bind.ContractBackend) (*OneInchSpotPrice, error) {
	contract, err := bindOneInchSpotPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPrice{OneInchSpotPriceCaller: OneInchSpotPriceCaller{contract: contract}, OneInchSpotPriceTransactor: OneInchSpotPriceTransactor{contract: contract}, OneInchSpotPriceFilterer: OneInchSpotPriceFilterer{contract: contract}}, nil
}

// NewOneInchSpotPriceCaller creates a new read-only instance of OneInchSpotPrice, bound to a specific deployed contract.
func NewOneInchSpotPriceCaller(address common.Address, caller bind.ContractCaller) (*OneInchSpotPriceCaller, error) {
	contract, err := bindOneInchSpotPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceCaller{contract: contract}, nil
}

// NewOneInchSpotPriceTransactor creates a new write-only instance of OneInchSpotPrice, bound to a specific deployed contract.
func NewOneInchSpotPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*OneInchSpotPriceTransactor, error) {
	contract, err := bindOneInchSpotPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceTransactor{contract: contract}, nil
}

// NewOneInchSpotPriceFilterer creates a new log filterer instance of OneInchSpotPrice, bound to a specific deployed contract.
func NewOneInchSpotPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*OneInchSpotPriceFilterer, error) {
	contract, err := bindOneInchSpotPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceFilterer{contract: contract}, nil
}

// bindOneInchSpotPrice binds a generic wrapper to an already deployed contract.
func bindOneInchSpotPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneInchSpotPriceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchSpotPrice *OneInchSpotPriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchSpotPrice.Contract.OneInchSpotPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchSpotPrice *OneInchSpotPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.OneInchSpotPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchSpotPrice *OneInchSpotPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.OneInchSpotPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneInchSpotPrice *OneInchSpotPriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OneInchSpotPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneInchSpotPrice *OneInchSpotPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneInchSpotPrice *OneInchSpotPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.contract.Transact(opts, method, params...)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OneInchSpotPrice *OneInchSpotPriceCaller) Connectors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OneInchSpotPrice.contract.Call(opts, &out, "connectors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OneInchSpotPrice *OneInchSpotPriceSession) Connectors() ([]common.Address, error) {
	return _OneInchSpotPrice.Contract.Connectors(&_OneInchSpotPrice.CallOpts)
}

// Connectors is a free data retrieval call binding the contract method 0x65050a68.
//
// Solidity: function connectors() view returns(address[] allConnectors)
func (_OneInchSpotPrice *OneInchSpotPriceCallerSession) Connectors() ([]common.Address, error) {
	return _OneInchSpotPrice.Contract.Connectors(&_OneInchSpotPrice.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPrice *OneInchSpotPriceCaller) GetRate(opts *bind.CallOpts, srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OneInchSpotPrice.contract.Call(opts, &out, "getRate", srcToken, dstToken, useWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPrice *OneInchSpotPriceSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OneInchSpotPrice.Contract.GetRate(&_OneInchSpotPrice.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRate is a free data retrieval call binding the contract method 0x802431fb.
//
// Solidity: function getRate(address srcToken, address dstToken, bool useWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPrice *OneInchSpotPriceCallerSession) GetRate(srcToken common.Address, dstToken common.Address, useWrappers bool) (*big.Int, error) {
	return _OneInchSpotPrice.Contract.GetRate(&_OneInchSpotPrice.CallOpts, srcToken, dstToken, useWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPrice *OneInchSpotPriceCaller) GetRateToEth(opts *bind.CallOpts, srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	var out []interface{}
	err := _OneInchSpotPrice.contract.Call(opts, &out, "getRateToEth", srcToken, useSrcWrappers)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPrice *OneInchSpotPriceSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OneInchSpotPrice.Contract.GetRateToEth(&_OneInchSpotPrice.CallOpts, srcToken, useSrcWrappers)
}

// GetRateToEth is a free data retrieval call binding the contract method 0x7de4fd10.
//
// Solidity: function getRateToEth(address srcToken, bool useSrcWrappers) view returns(uint256 weightedRate)
func (_OneInchSpotPrice *OneInchSpotPriceCallerSession) GetRateToEth(srcToken common.Address, useSrcWrappers bool) (*big.Int, error) {
	return _OneInchSpotPrice.Contract.GetRateToEth(&_OneInchSpotPrice.CallOpts, srcToken, useSrcWrappers)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OneInchSpotPrice *OneInchSpotPriceCaller) MultiWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchSpotPrice.contract.Call(opts, &out, "multiWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OneInchSpotPrice *OneInchSpotPriceSession) MultiWrapper() (common.Address, error) {
	return _OneInchSpotPrice.Contract.MultiWrapper(&_OneInchSpotPrice.CallOpts)
}

// MultiWrapper is a free data retrieval call binding the contract method 0xb77910dc.
//
// Solidity: function multiWrapper() view returns(address)
func (_OneInchSpotPrice *OneInchSpotPriceCallerSession) MultiWrapper() (common.Address, error) {
	return _OneInchSpotPrice.Contract.MultiWrapper(&_OneInchSpotPrice.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OneInchSpotPrice *OneInchSpotPriceCaller) Oracles(opts *bind.CallOpts) (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	var out []interface{}
	err := _OneInchSpotPrice.contract.Call(opts, &out, "oracles")

	outstruct := new(struct {
		AllOracles  []common.Address
		OracleTypes []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AllOracles = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.OracleTypes = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OneInchSpotPrice *OneInchSpotPriceSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OneInchSpotPrice.Contract.Oracles(&_OneInchSpotPrice.CallOpts)
}

// Oracles is a free data retrieval call binding the contract method 0x2857373a.
//
// Solidity: function oracles() view returns(address[] allOracles, uint8[] oracleTypes)
func (_OneInchSpotPrice *OneInchSpotPriceCallerSession) Oracles() (struct {
	AllOracles  []common.Address
	OracleTypes []uint8
}, error) {
	return _OneInchSpotPrice.Contract.Oracles(&_OneInchSpotPrice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchSpotPrice *OneInchSpotPriceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OneInchSpotPrice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchSpotPrice *OneInchSpotPriceSession) Owner() (common.Address, error) {
	return _OneInchSpotPrice.Contract.Owner(&_OneInchSpotPrice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OneInchSpotPrice *OneInchSpotPriceCallerSession) Owner() (common.Address, error) {
	return _OneInchSpotPrice.Contract.Owner(&_OneInchSpotPrice.CallOpts)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactor) AddConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.contract.Transact(opts, "addConnector", connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OneInchSpotPrice *OneInchSpotPriceSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.AddConnector(&_OneInchSpotPrice.TransactOpts, connector)
}

// AddConnector is a paid mutator transaction binding the contract method 0xaa16d4c0.
//
// Solidity: function addConnector(address connector) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactorSession) AddConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.AddConnector(&_OneInchSpotPrice.TransactOpts, connector)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactor) AddOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPrice.contract.Transact(opts, "addOracle", oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPrice *OneInchSpotPriceSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.AddOracle(&_OneInchSpotPrice.TransactOpts, oracle, oracleKind)
}

// AddOracle is a paid mutator transaction binding the contract method 0x9d4d7b1c.
//
// Solidity: function addOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactorSession) AddOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.AddOracle(&_OneInchSpotPrice.TransactOpts, oracle, oracleKind)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactor) RemoveConnector(opts *bind.TransactOpts, connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.contract.Transact(opts, "removeConnector", connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OneInchSpotPrice *OneInchSpotPriceSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.RemoveConnector(&_OneInchSpotPrice.TransactOpts, connector)
}

// RemoveConnector is a paid mutator transaction binding the contract method 0x1a6c6a98.
//
// Solidity: function removeConnector(address connector) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactorSession) RemoveConnector(connector common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.RemoveConnector(&_OneInchSpotPrice.TransactOpts, connector)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactor) RemoveOracle(opts *bind.TransactOpts, oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPrice.contract.Transact(opts, "removeOracle", oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPrice *OneInchSpotPriceSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.RemoveOracle(&_OneInchSpotPrice.TransactOpts, oracle, oracleKind)
}

// RemoveOracle is a paid mutator transaction binding the contract method 0xf0b92e40.
//
// Solidity: function removeOracle(address oracle, uint8 oracleKind) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactorSession) RemoveOracle(oracle common.Address, oracleKind uint8) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.RemoveOracle(&_OneInchSpotPrice.TransactOpts, oracle, oracleKind)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneInchSpotPrice.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchSpotPrice *OneInchSpotPriceSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.RenounceOwnership(&_OneInchSpotPrice.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.RenounceOwnership(&_OneInchSpotPrice.TransactOpts)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactor) SetMultiWrapper(opts *bind.TransactOpts, _multiWrapper common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.contract.Transact(opts, "setMultiWrapper", _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OneInchSpotPrice *OneInchSpotPriceSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.SetMultiWrapper(&_OneInchSpotPrice.TransactOpts, _multiWrapper)
}

// SetMultiWrapper is a paid mutator transaction binding the contract method 0xd0626518.
//
// Solidity: function setMultiWrapper(address _multiWrapper) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactorSession) SetMultiWrapper(_multiWrapper common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.SetMultiWrapper(&_OneInchSpotPrice.TransactOpts, _multiWrapper)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchSpotPrice *OneInchSpotPriceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.TransferOwnership(&_OneInchSpotPrice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OneInchSpotPrice *OneInchSpotPriceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OneInchSpotPrice.Contract.TransferOwnership(&_OneInchSpotPrice.TransactOpts, newOwner)
}

// OneInchSpotPriceConnectorAddedIterator is returned from FilterConnectorAdded and is used to iterate over the raw logs and unpacked data for ConnectorAdded events raised by the OneInchSpotPrice contract.
type OneInchSpotPriceConnectorAddedIterator struct {
	Event *OneInchSpotPriceConnectorAdded // Event containing the contract specifics and raw log

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
func (it *OneInchSpotPriceConnectorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchSpotPriceConnectorAdded)
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
		it.Event = new(OneInchSpotPriceConnectorAdded)
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
func (it *OneInchSpotPriceConnectorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchSpotPriceConnectorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchSpotPriceConnectorAdded represents a ConnectorAdded event raised by the OneInchSpotPrice contract.
type OneInchSpotPriceConnectorAdded struct {
	Connector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConnectorAdded is a free log retrieval operation binding the contract event 0xff88af5d962d47fd25d87755e8267a029fad5a91740c67d0dade2bdbe5268a1d.
//
// Solidity: event ConnectorAdded(address connector)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) FilterConnectorAdded(opts *bind.FilterOpts) (*OneInchSpotPriceConnectorAddedIterator, error) {

	logs, sub, err := _OneInchSpotPrice.contract.FilterLogs(opts, "ConnectorAdded")
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceConnectorAddedIterator{contract: _OneInchSpotPrice.contract, event: "ConnectorAdded", logs: logs, sub: sub}, nil
}

// WatchConnectorAdded is a free log subscription operation binding the contract event 0xff88af5d962d47fd25d87755e8267a029fad5a91740c67d0dade2bdbe5268a1d.
//
// Solidity: event ConnectorAdded(address connector)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) WatchConnectorAdded(opts *bind.WatchOpts, sink chan<- *OneInchSpotPriceConnectorAdded) (event.Subscription, error) {

	logs, sub, err := _OneInchSpotPrice.contract.WatchLogs(opts, "ConnectorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchSpotPriceConnectorAdded)
				if err := _OneInchSpotPrice.contract.UnpackLog(event, "ConnectorAdded", log); err != nil {
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

// ParseConnectorAdded is a log parse operation binding the contract event 0xff88af5d962d47fd25d87755e8267a029fad5a91740c67d0dade2bdbe5268a1d.
//
// Solidity: event ConnectorAdded(address connector)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) ParseConnectorAdded(log types.Log) (*OneInchSpotPriceConnectorAdded, error) {
	event := new(OneInchSpotPriceConnectorAdded)
	if err := _OneInchSpotPrice.contract.UnpackLog(event, "ConnectorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchSpotPriceConnectorRemovedIterator is returned from FilterConnectorRemoved and is used to iterate over the raw logs and unpacked data for ConnectorRemoved events raised by the OneInchSpotPrice contract.
type OneInchSpotPriceConnectorRemovedIterator struct {
	Event *OneInchSpotPriceConnectorRemoved // Event containing the contract specifics and raw log

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
func (it *OneInchSpotPriceConnectorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchSpotPriceConnectorRemoved)
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
		it.Event = new(OneInchSpotPriceConnectorRemoved)
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
func (it *OneInchSpotPriceConnectorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchSpotPriceConnectorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchSpotPriceConnectorRemoved represents a ConnectorRemoved event raised by the OneInchSpotPrice contract.
type OneInchSpotPriceConnectorRemoved struct {
	Connector common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConnectorRemoved is a free log retrieval operation binding the contract event 0x6825b26a0827e9c2ceca01d6289ce4a40e629dc074ec48ea4727d1afbff359f5.
//
// Solidity: event ConnectorRemoved(address connector)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) FilterConnectorRemoved(opts *bind.FilterOpts) (*OneInchSpotPriceConnectorRemovedIterator, error) {

	logs, sub, err := _OneInchSpotPrice.contract.FilterLogs(opts, "ConnectorRemoved")
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceConnectorRemovedIterator{contract: _OneInchSpotPrice.contract, event: "ConnectorRemoved", logs: logs, sub: sub}, nil
}

// WatchConnectorRemoved is a free log subscription operation binding the contract event 0x6825b26a0827e9c2ceca01d6289ce4a40e629dc074ec48ea4727d1afbff359f5.
//
// Solidity: event ConnectorRemoved(address connector)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) WatchConnectorRemoved(opts *bind.WatchOpts, sink chan<- *OneInchSpotPriceConnectorRemoved) (event.Subscription, error) {

	logs, sub, err := _OneInchSpotPrice.contract.WatchLogs(opts, "ConnectorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchSpotPriceConnectorRemoved)
				if err := _OneInchSpotPrice.contract.UnpackLog(event, "ConnectorRemoved", log); err != nil {
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

// ParseConnectorRemoved is a log parse operation binding the contract event 0x6825b26a0827e9c2ceca01d6289ce4a40e629dc074ec48ea4727d1afbff359f5.
//
// Solidity: event ConnectorRemoved(address connector)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) ParseConnectorRemoved(log types.Log) (*OneInchSpotPriceConnectorRemoved, error) {
	event := new(OneInchSpotPriceConnectorRemoved)
	if err := _OneInchSpotPrice.contract.UnpackLog(event, "ConnectorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchSpotPriceMultiWrapperUpdatedIterator is returned from FilterMultiWrapperUpdated and is used to iterate over the raw logs and unpacked data for MultiWrapperUpdated events raised by the OneInchSpotPrice contract.
type OneInchSpotPriceMultiWrapperUpdatedIterator struct {
	Event *OneInchSpotPriceMultiWrapperUpdated // Event containing the contract specifics and raw log

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
func (it *OneInchSpotPriceMultiWrapperUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchSpotPriceMultiWrapperUpdated)
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
		it.Event = new(OneInchSpotPriceMultiWrapperUpdated)
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
func (it *OneInchSpotPriceMultiWrapperUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchSpotPriceMultiWrapperUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchSpotPriceMultiWrapperUpdated represents a MultiWrapperUpdated event raised by the OneInchSpotPrice contract.
type OneInchSpotPriceMultiWrapperUpdated struct {
	MultiWrapper common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMultiWrapperUpdated is a free log retrieval operation binding the contract event 0x1030152fe2062b574a830e6b9f13c65995990df31e4dc708d142533bb3ad0f52.
//
// Solidity: event MultiWrapperUpdated(address multiWrapper)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) FilterMultiWrapperUpdated(opts *bind.FilterOpts) (*OneInchSpotPriceMultiWrapperUpdatedIterator, error) {

	logs, sub, err := _OneInchSpotPrice.contract.FilterLogs(opts, "MultiWrapperUpdated")
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceMultiWrapperUpdatedIterator{contract: _OneInchSpotPrice.contract, event: "MultiWrapperUpdated", logs: logs, sub: sub}, nil
}

// WatchMultiWrapperUpdated is a free log subscription operation binding the contract event 0x1030152fe2062b574a830e6b9f13c65995990df31e4dc708d142533bb3ad0f52.
//
// Solidity: event MultiWrapperUpdated(address multiWrapper)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) WatchMultiWrapperUpdated(opts *bind.WatchOpts, sink chan<- *OneInchSpotPriceMultiWrapperUpdated) (event.Subscription, error) {

	logs, sub, err := _OneInchSpotPrice.contract.WatchLogs(opts, "MultiWrapperUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchSpotPriceMultiWrapperUpdated)
				if err := _OneInchSpotPrice.contract.UnpackLog(event, "MultiWrapperUpdated", log); err != nil {
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

// ParseMultiWrapperUpdated is a log parse operation binding the contract event 0x1030152fe2062b574a830e6b9f13c65995990df31e4dc708d142533bb3ad0f52.
//
// Solidity: event MultiWrapperUpdated(address multiWrapper)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) ParseMultiWrapperUpdated(log types.Log) (*OneInchSpotPriceMultiWrapperUpdated, error) {
	event := new(OneInchSpotPriceMultiWrapperUpdated)
	if err := _OneInchSpotPrice.contract.UnpackLog(event, "MultiWrapperUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchSpotPriceOracleAddedIterator is returned from FilterOracleAdded and is used to iterate over the raw logs and unpacked data for OracleAdded events raised by the OneInchSpotPrice contract.
type OneInchSpotPriceOracleAddedIterator struct {
	Event *OneInchSpotPriceOracleAdded // Event containing the contract specifics and raw log

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
func (it *OneInchSpotPriceOracleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchSpotPriceOracleAdded)
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
		it.Event = new(OneInchSpotPriceOracleAdded)
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
func (it *OneInchSpotPriceOracleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchSpotPriceOracleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchSpotPriceOracleAdded represents a OracleAdded event raised by the OneInchSpotPrice contract.
type OneInchSpotPriceOracleAdded struct {
	Oracle     common.Address
	OracleType uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOracleAdded is a free log retrieval operation binding the contract event 0x5874b2072ff37562df54063dd700c59d45f311bdf6f9cabb5a15f0ffb2e9f622.
//
// Solidity: event OracleAdded(address oracle, uint8 oracleType)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) FilterOracleAdded(opts *bind.FilterOpts) (*OneInchSpotPriceOracleAddedIterator, error) {

	logs, sub, err := _OneInchSpotPrice.contract.FilterLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceOracleAddedIterator{contract: _OneInchSpotPrice.contract, event: "OracleAdded", logs: logs, sub: sub}, nil
}

// WatchOracleAdded is a free log subscription operation binding the contract event 0x5874b2072ff37562df54063dd700c59d45f311bdf6f9cabb5a15f0ffb2e9f622.
//
// Solidity: event OracleAdded(address oracle, uint8 oracleType)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) WatchOracleAdded(opts *bind.WatchOpts, sink chan<- *OneInchSpotPriceOracleAdded) (event.Subscription, error) {

	logs, sub, err := _OneInchSpotPrice.contract.WatchLogs(opts, "OracleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchSpotPriceOracleAdded)
				if err := _OneInchSpotPrice.contract.UnpackLog(event, "OracleAdded", log); err != nil {
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

// ParseOracleAdded is a log parse operation binding the contract event 0x5874b2072ff37562df54063dd700c59d45f311bdf6f9cabb5a15f0ffb2e9f622.
//
// Solidity: event OracleAdded(address oracle, uint8 oracleType)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) ParseOracleAdded(log types.Log) (*OneInchSpotPriceOracleAdded, error) {
	event := new(OneInchSpotPriceOracleAdded)
	if err := _OneInchSpotPrice.contract.UnpackLog(event, "OracleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchSpotPriceOracleRemovedIterator is returned from FilterOracleRemoved and is used to iterate over the raw logs and unpacked data for OracleRemoved events raised by the OneInchSpotPrice contract.
type OneInchSpotPriceOracleRemovedIterator struct {
	Event *OneInchSpotPriceOracleRemoved // Event containing the contract specifics and raw log

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
func (it *OneInchSpotPriceOracleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchSpotPriceOracleRemoved)
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
		it.Event = new(OneInchSpotPriceOracleRemoved)
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
func (it *OneInchSpotPriceOracleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchSpotPriceOracleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchSpotPriceOracleRemoved represents a OracleRemoved event raised by the OneInchSpotPrice contract.
type OneInchSpotPriceOracleRemoved struct {
	Oracle     common.Address
	OracleType uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOracleRemoved is a free log retrieval operation binding the contract event 0x7a7f56716fe703fb190529c336e57df71ab88188ba47e8d786bac684b61ab9a6.
//
// Solidity: event OracleRemoved(address oracle, uint8 oracleType)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) FilterOracleRemoved(opts *bind.FilterOpts) (*OneInchSpotPriceOracleRemovedIterator, error) {

	logs, sub, err := _OneInchSpotPrice.contract.FilterLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceOracleRemovedIterator{contract: _OneInchSpotPrice.contract, event: "OracleRemoved", logs: logs, sub: sub}, nil
}

// WatchOracleRemoved is a free log subscription operation binding the contract event 0x7a7f56716fe703fb190529c336e57df71ab88188ba47e8d786bac684b61ab9a6.
//
// Solidity: event OracleRemoved(address oracle, uint8 oracleType)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) WatchOracleRemoved(opts *bind.WatchOpts, sink chan<- *OneInchSpotPriceOracleRemoved) (event.Subscription, error) {

	logs, sub, err := _OneInchSpotPrice.contract.WatchLogs(opts, "OracleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchSpotPriceOracleRemoved)
				if err := _OneInchSpotPrice.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
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

// ParseOracleRemoved is a log parse operation binding the contract event 0x7a7f56716fe703fb190529c336e57df71ab88188ba47e8d786bac684b61ab9a6.
//
// Solidity: event OracleRemoved(address oracle, uint8 oracleType)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) ParseOracleRemoved(log types.Log) (*OneInchSpotPriceOracleRemoved, error) {
	event := new(OneInchSpotPriceOracleRemoved)
	if err := _OneInchSpotPrice.contract.UnpackLog(event, "OracleRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OneInchSpotPriceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OneInchSpotPrice contract.
type OneInchSpotPriceOwnershipTransferredIterator struct {
	Event *OneInchSpotPriceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OneInchSpotPriceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OneInchSpotPriceOwnershipTransferred)
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
		it.Event = new(OneInchSpotPriceOwnershipTransferred)
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
func (it *OneInchSpotPriceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OneInchSpotPriceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OneInchSpotPriceOwnershipTransferred represents a OwnershipTransferred event raised by the OneInchSpotPrice contract.
type OneInchSpotPriceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OneInchSpotPriceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchSpotPrice.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OneInchSpotPriceOwnershipTransferredIterator{contract: _OneInchSpotPrice.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OneInchSpotPriceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OneInchSpotPrice.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OneInchSpotPriceOwnershipTransferred)
				if err := _OneInchSpotPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OneInchSpotPrice *OneInchSpotPriceFilterer) ParseOwnershipTransferred(log types.Log) (*OneInchSpotPriceOwnershipTransferred, error) {
	event := new(OneInchSpotPriceOwnershipTransferred)
	if err := _OneInchSpotPrice.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
