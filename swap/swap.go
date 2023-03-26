package swap

import (
	"github.com/Minh-Tin/mtlib/helper"
	"github.com/cockroachdb/errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

var (
	// https://etherscan.io/address/0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45#code#F23#L11
	addrMsgSender = common.BigToAddress(big.NewInt(1)) // the Tx sender
	addrThis      = common.BigToAddress(big.NewInt(2)) // the router
)

type SwapCall struct {
	Method SwapMethod
	Path   []common.Address
}

func (s *SwapCall) AddPath(path []common.Address) {
	for _, p := range path {
		if len(s.Path) == 0 {
			s.Path = append(s.Path, p)
		} else {
			if !helper.IsTheSame(p, s.Path[len(s.Path)-1]) {
				s.Path = append(s.Path, p)
			}
		}
	}
}

type Params struct {
	Tx         *types.Transaction
	Sender     common.Address
	Recipient  common.Address
	SwapToken  common.Address
	RxToken    common.Address
	SwapAmount *big.Int // AmountIn / AmountInMax
	RxAmount   *big.Int // AmountOut / AmountOutMin
	Method     SwapMethod
	Calls      []*SwapCall
}

func (sp *Params) PopulateTradeMethodParams(method SwapMethod, recipient, tokenIn, tokenOut common.Address, amtIn, amtOut *big.Int) error {
	if sp.Method < 1 {
		sp.Method = method
	}
	cSwapToken, cRxToken, cSwapAmount, cRxAmount := sp.SwapToken, sp.RxToken, new(big.Int).Set(sp.SwapAmount), new(big.Int).Set(sp.RxAmount)
	sp.Recipient = InputRecipient(recipient, sp.Sender)
	if !helper.IsSet(sp.SwapToken) {
		sp.SwapToken = tokenIn
		if helper.IsSet(sp.RxToken) {
			return errors.New("multiple rx tokens 1")
		}
		sp.RxToken = tokenOut
		sp.RxAmount.Add(sp.RxAmount, amtOut)
		sp.SwapAmount.Add(sp.SwapAmount, amtIn)

	} else { // swap Token already set
		if helper.IsTheSame(sp.SwapToken, tokenIn) {
			sp.SwapAmount.Add(sp.SwapAmount, amtIn)
		}
		if helper.IsSet(sp.RxToken) {
			if !helper.IsTheSame(sp.RxToken, tokenOut) {
				sp.RxToken = tokenOut
			}
		} else {
			sp.RxToken = tokenOut
		}
		sp.RxAmount.Add(sp.RxAmount, amtOut)
	}
	switch method {
	case ExactInput, ExactInputSingle,
		V3SwapExactIn, V2SwapExactIn,
		SwapExactTokensForTokens,
		SwapExactTokensForTokensSupportingFeeOnTransferTokens,
		SwapExactTokensForETHSupportingFeeOnTransferTokens,
		SwapExactTokensForETH,
		SwapExactETHForTokens:
		if helper.IsTheSame(cSwapToken, tokenIn) {
			sp.SwapAmount = new(big.Int).Add(cSwapAmount, amtIn)
		}

	case ExactOutput, ExactOutputSingle,
		V3SwapExactOut, V2SwapExactOut,
		SwapTokensForExactETH,
		SwapTokensForExactTokens,
		SwapETHForExactTokens:
		if helper.IsTheSame(cRxToken, tokenOut) {
			sp.RxAmount = new(big.Int).Add(cRxAmount, amtOut)
		}
	}
	return nil
}

type RemoveLiquidity struct {
	Token, TokenA, TokenB common.Address
	Liquidity             *big.Int
}

func InputRecipient(swapMethodRecipient, sender common.Address) common.Address {
	if helper.IsTheSame(swapMethodRecipient, addrMsgSender) {
		return sender
	}
	return swapMethodRecipient // could also be addrThis
}

func IsValidRecipient(recipient common.Address) bool {
	if helper.IsTheSame(recipient, addrMsgSender) || helper.IsTheSame(recipient, addrThis) || !helper.IsSet(recipient) {
		return false
	}
	return true
}
