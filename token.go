package mt

import (
	"encoding/hex"
	"fmt"
	"github.com/Minh-Tin/mtlib/abi/ERC20"
	"github.com/Minh-Tin/mtlib/abi/Erc20Bytes32"
	"github.com/Minh-Tin/mtlib/helper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Token struct {
	instance *ERC20.ERC20
	Name     string         `json:"name"`
	Address  common.Address `json:"address"`
	Symbol   string         `json:"symbol"`
	Decimals *big.Int       `json:"decimals"`
}

func GetToken(address common.Address) (*Token, error) {
	nameOk, symbolOk := true, true
	c, err := ERC20.NewERC20(address, C)
	if err != nil {
		return nil, err
	}
	name, err := c.Name(nil)
	if err != nil {
		nameOk = false
	}
	symbol, err := c.Symbol(nil)
	if err != nil {
		symbolOk = false
	}
	decimals, err := c.Decimals(nil)
	if err != nil {
		return nil, err
	}
	if !nameOk || !symbolOk {
		c2, err := Erc20Bytes32.NewErc20Bytes32(address, C)
		if err != nil {
			return nil, err
		}
		if !nameOk {
			name32, err := c2.Name(nil)
			if err != nil {
				return nil, err
			}
			bs, err := hex.DecodeString(fmt.Sprintf("%x", name32))
			if err != nil {
				return nil, err
			}
			name = string(bs)
		}
		if !symbolOk {
			symbol32, err := c2.Symbol(nil)
			if err != nil {
				return nil, err
			}
			bs, err := hex.DecodeString(fmt.Sprintf("%x", symbol32))
			if err != nil {
				return nil, err
			}
			symbol = string(bs)
		}
	}
	return &Token{
		instance: c,
		Address:  address,
		Name:     name,
		Symbol:   symbol,
		Decimals: new(big.Int).SetUint64(uint64(decimals)),
	}, nil
}

func (t *Token) getInstance() (*ERC20.ERC20, error) {
	var err error
	if t.instance == nil {
		t.instance, err = ERC20.NewERC20(t.Address, C)
	}
	return t.instance, err
}

func (t *Token) GetRateToEth(blockNumber *big.Int) (*big.Int, error) {
	var opts *bind.CallOpts
	if blockNumber != nil {
		opts = &bind.CallOpts{BlockNumber: blockNumber}
	}
	return SpotPrice.GetRateToEth(opts, t.Address, false)
}

func (t *Token) GetRate(quote common.Address, blockNumber *big.Int) (*big.Int, error) {
	var opts *bind.CallOpts
	if blockNumber != nil {
		opts = &bind.CallOpts{BlockNumber: blockNumber}
	}
	return SpotPrice.GetRate(opts, t.Address, quote, false)
}

func (t *Token) IsStableCoin() bool {
	if _, ok := StableCoins[t.Address]; ok {
		return true
	}
	return false
}

func (t *Token) IsNative() bool {
	return helper.IsTheSame(t.Address, NativeToken)
}

func (t *Token) BalanceOf(owner common.Address) (*big.Int, error) {
	if instance, err := t.getInstance(); err != nil {
		return nil, err
	} else {
		return instance.BalanceOf(nil, owner)
	}
}

func (t *Token) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	if instance, err := t.getInstance(); err != nil {
		return nil, err
	} else {
		return instance.Allowance(nil, owner, spender)
	}
}
