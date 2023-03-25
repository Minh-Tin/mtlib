package mt

import (
	"encoding/json"
	"fmt"
	"github.com/Minh-Tin/mtlib/abi/OneInchV5"
	"github.com/Minh-Tin/mtlib/abi/UniswapUniversalRouter"
	"github.com/Minh-Tin/mtlib/abi/uni"
	"github.com/Minh-Tin/mtlib/ethutil"
	"github.com/Minh-Tin/mtlib/helper"
	"github.com/Minh-Tin/mtlib/swap"
	"github.com/cockroachdb/errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/orcaman/concurrent-map/v2"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"math/big"
)

var (
	oneInchNativeToken = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	pairsAndPools      = cmap.New[*poolPair]()
)
var (
	addZero = common.BigToAddress(big.NewInt(0)) // the router
)

type poolHop struct {
	In  common.Address
	Out common.Address
}

type poolPair struct {
	Token0, Token1 common.Address
}

type oneInchSwapInput struct {
	SrcToken        common.Address "json:\"srcToken\""
	DstToken        common.Address "json:\"dstToken\""
	SrcReceiver     common.Address "json:\"srcReceiver\""
	DstReceiver     common.Address "json:\"dstReceiver\""
	Amount          *big.Int       "json:\"amount\""
	MinReturnAmount *big.Int       "json:\"minReturnAmount\""
	Flags           *big.Int       "json:\"flags\""
	Permit          []uint8        "json:\"permit\""
}

func DecodeSwapByInput(tx *types.Transaction, callback swap.CallBack) (*swap.Params, error) {
	if tx.To() == nil || len(tx.Data()) < 4 {
		return nil, errors.Errorf("Not valid transaction")
	}
	if dex, ok := Dexs[*tx.To()]; ok {
		fmt.Println(tx.Hash().Hex())
		sender, err := ethutil.GetTxSender(tx)
		if err != nil {
			return nil, err
		}
		sp := &swap.Params{
			Tx:         tx,
			Sender:     sender,
			SwapAmount: new(big.Int),
			RxAmount:   new(big.Int),
			Method:     -1,
		}
		return sp, decodeSwapByInput(dex, tx.Data(), tx, sp, callback)
	}
	return nil, errors.Errorf("Not valid transaction")
}

func decodeSwapByInput(dex *Dex, d []byte, tx *types.Transaction, sp *swap.Params, callback swap.CallBack) error {
	method, err := dex.abi.MethodById(d)
	if err != nil {
		return err
	}
	if method == nil {
		err = errors.Errorf("method nil for %x", d)
		return err
	}
	iArgs := make(map[string]interface{})
	if err = method.Inputs.UnpackIntoMap(iArgs, d[4:]); err != nil {
		return errors.Wrapf(err, "failed to unpack method '%s' input args", method.Name)
	}

	switch method.Name {
	case "multicall0", "multicall1", "multicall":
		sp.Method = swap.Multicall
		dRaw, ok := iArgs["data"]
		if !ok || dRaw == nil {
			err = errors.Errorf("unable to get '%s' input data", method.Name)
			return err
		}
		dMulti, ok := dRaw.([][]byte)
		if !ok {
			err = errors.Errorf("unable to get '%s' input data 2", method.Name)
			return err
		}
		for _, d := range dMulti {
			if err = decodeSwapByInput(dex, d, tx, sp, callback); err != nil {
				return err
			}
		}
	case "execute", "execute0":
		sp.Method = swap.Execute
		commands, ok := iArgs["commands"].([]uint8)
		if !ok {
			err = errors.Errorf("unable to get '%s' input data 3", method.Name)
			return err
		}
		iRaw, ok := iArgs["inputs"]
		if !ok || iRaw == nil {
			err = errors.Errorf("unable to get '%s' input data 4", method.Name)
			return err
		}
		iMulti, ok := iRaw.([][]byte)
		if !ok {
			err = errors.Errorf("unable to get '%s' input data 5", method.Name)
			return err
		}
		for idx, command := range commands {
			switch command {
			case 0x00: //V3_SWAP_EXACT_IN
				pRaw := iMulti[idx]
				res, err := abi.Decode(UniswapUniversalRouter.V3_SWAP_EXACT_IN, pRaw)
				if err != nil {
					return errors.Wrap(err, "failed to decode V3_SWAP_EXACT_IN")
				}
				spew.Dump("V3_SWAP_EXACT_IN", res)
				input := res.(map[string]interface{})
				inputBytes := input["path"].([]byte)
				hops, err := getHopsFromMultiSwapPath(inputBytes, false)
				if err != nil {
					return errors.Wrap(err, "failed to decode MultiSwap Path arg")
				}
				tokenIn := hops[0].In
				tokenOut := hops[len(hops)-1].Out
				if err = sp.PopulateTradeMethodParams(
					swap.V3SwapExactIn,
					common.HexToAddress(input["recipient"].(ethgo.Address).String()),
					tokenIn, tokenOut,
					input["amountIn"].(*big.Int), input["amountOutMin"].(*big.Int)); err != nil {
					return err
				}
			case 0x01: //V3_SWAP_EXACT_OUT
				pRaw := iMulti[idx]
				res, err := abi.Decode(UniswapUniversalRouter.V3_SWAP_EXACT_OUT, pRaw)
				if err != nil {
					return err
				}
				input := res.(map[string]interface{})
				inputBytes := input["path"].([]byte)

				hops, err := getHopsFromMultiSwapPath(inputBytes, true)
				if err != nil {
					return errors.Wrap(err, "failed to decode V3_SWAP_EXACT_OUT")
				}

				tokenIn := hops[0].In
				tokenOut := hops[len(hops)-1].Out
				if err = sp.PopulateTradeMethodParams(
					swap.V3SwapExactOut,
					common.HexToAddress(input["recipient"].(ethgo.Address).String()),
					tokenIn, tokenOut,
					input["amountInMax"].(*big.Int), input["amountOut"].(*big.Int)); err != nil {
					return err
				}
			case 0x08: //V2_SWAP_EXACT_IN
				pRaw := iMulti[idx]
				res, err := abi.Decode(UniswapUniversalRouter.V2_SWAP_EXACT_IN, pRaw)
				if err != nil {
					return errors.Wrap(err, "failed to decode V2_SWAP_EXACT_IN")
				}
				spew.Dump(res)
				input := res.(map[string]interface{})
				path := input["path"].([]ethgo.Address)

				tokenIn := common.HexToAddress(path[0].String())
				tokenOut := common.HexToAddress(path[len(path)-1].String())
				if err = sp.PopulateTradeMethodParams(
					swap.V2SwapExactIn,
					common.HexToAddress(input["recipient"].(ethgo.Address).String()),
					tokenIn, tokenOut,
					input["amountIn"].(*big.Int), input["amountOutMin"].(*big.Int)); err != nil {
					return err
				}
			case 0x09: //V2_SWAP_EXACT_OUT
				pRaw := iMulti[idx]
				res, err := abi.Decode(UniswapUniversalRouter.V2_SWAP_EXACT_OUT, pRaw)
				if err != nil {
					return errors.Wrap(err, "failed to decode V2_SWAP_EXACT_OUT")
				}
				input := res.(map[string]interface{})
				path := input["path"].([]ethgo.Address)
				tokenIn := common.HexToAddress(path[0].String())
				tokenOut := common.HexToAddress(path[len(path)-1].String())
				if err = sp.PopulateTradeMethodParams(
					swap.V2SwapExactOut,
					common.HexToAddress(input["recipient"].(ethgo.Address).String()),
					tokenIn, tokenOut,
					input["amountInMax"].(*big.Int), input["amountOut"].(*big.Int)); err != nil {
					return err
				}
			case 0x0c:
				if swap.IsValidRecipient(sp.Recipient) {
					break
				}
				pRaw := iMulti[idx]
				res, err := abi.Decode(UniswapUniversalRouter.UNWRAP_ETH, pRaw)
				if err != nil {
					return errors.Wrap(err, "failed to decode UNWRAP_ETH")
				}
				input := res.(map[string]interface{})
				sp.Recipient = swap.InputRecipient(common.HexToAddress(input["recipient"].(ethgo.Address).String()), sp.Sender)
			case 0x04:
				if swap.IsValidRecipient(sp.Recipient) {
					break
				}
				pRaw := iMulti[idx]
				res, err := abi.Decode(UniswapUniversalRouter.SWEEP, pRaw)
				if err != nil {
					return errors.Wrap(err, "failed to decode UNWRAP_ETH")
				}
				input := res.(map[string]interface{})
				spew.Dump(input)
				sp.Recipient = swap.InputRecipient(common.HexToAddress(input["recipient"].(ethgo.Address).String()), sp.Sender)
			}
		}
	case "exactInputSingle":
		pRaw, ok := iArgs["params"]
		if !ok || pRaw == nil {
			err = errors.New("failed to get exactInputSingle params")
			return err
		}

		jRaw, _ := json.Marshal(pRaw)
		var params struct {
			TokenIn           common.Address `json:"tokenIn"`
			TokenOut          common.Address `json:"tokenOut"`
			Fee               *big.Int       `json:"fee"`
			Recipient         common.Address `json:"recipient"`
			AmountIn          *big.Int       `json:"amountIn"`
			AmountOutMinimum  *big.Int       `json:"amountOutMinimum"`
			SqrtPriceLimitX96 *big.Int       `json:"sqrtPriceLimitX96"`
		}
		if err = json.Unmarshal(jRaw, &params); err != nil {
			return errors.Wrapf(err, "failed to unmarshal exactInputSingle params")
		}

		if err = sp.PopulateTradeMethodParams(
			swap.ExactInputSingle,
			params.Recipient,
			params.TokenIn, params.TokenOut,
			params.AmountIn, params.AmountOutMinimum); err != nil {
			return err
		}

	case "exactInput":
		paramsRaw, ok := iArgs["params"]
		if !ok || paramsRaw == nil {
			return errors.New("failed to get exactInput params")
		}
		jRaw, _ := json.Marshal(paramsRaw)
		var params struct {
			Path             []byte         `json:"path"`
			Recipient        common.Address `json:"recipient"`
			AmountIn         *big.Int       `json:"amountIn"`
			AmountOutMinimum *big.Int       `json:"amountOutMinimum"`
		}
		if err = json.Unmarshal(jRaw, &params); err != nil {
			return errors.Wrapf(err, "failed to unmarshal exactInput params")
		}

		hops, err := getHopsFromMultiSwapPath(params.Path, false)
		if err != nil {
			return errors.Wrap(err, "failed to decode MultiSwap Path arg")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.ExactInput,
			params.Recipient,
			hops[0].In, hops[len(hops)-1].Out,
			params.AmountIn, params.AmountOutMinimum); err != nil {
			return err
		}

	case "exactOutputSingle":
		paramsRaw, ok := iArgs["params"]
		if !ok || paramsRaw == nil {
			return errors.New("failed to get exactOutputSingle params")
		}
		jRaw, _ := json.Marshal(paramsRaw)
		var params struct {
			TokenIn           common.Address `json:"tokenIn"`
			TokenOut          common.Address `json:"tokenOut"`
			Fee               *big.Int       `json:"fee"`
			Recipient         common.Address `json:"recipient"`
			AmountOut         *big.Int       `json:"amountOut"`
			AmountInMaximum   *big.Int       `json:"amountInMaximum"`
			SqrtPriceLimitX96 *big.Int       `json:"sqrtPriceLimitX96"`
		}
		if err = json.Unmarshal(jRaw, &params); err != nil {
			return errors.Wrapf(err, "failed to unmarshal exactOutputSingle params")
		}

		if err = sp.PopulateTradeMethodParams(
			swap.ExactOutputSingle,
			params.Recipient,
			params.TokenIn, params.TokenOut,
			params.AmountInMaximum, params.AmountOut); err != nil {
			return err
		}

	case "exactOutput":
		paramsRaw, ok := iArgs["params"]
		if !ok || paramsRaw == nil {
			return errors.New("failed to get exactOutput params")
		}
		jRaw, _ := json.Marshal(paramsRaw)
		var params struct {
			Path            []byte         `json:"path"`
			Recipient       common.Address `json:"recipient"`
			AmountOut       *big.Int       `json:"amountOut"`
			AmountInMaximum *big.Int       `json:"amountInMaximum"`
		}
		if err = json.Unmarshal(jRaw, &params); err != nil {
			return errors.Wrapf(err, "failed to unmarshal exactOutput params")
		}

		hops, err := getHopsFromMultiSwapPath(params.Path, true)
		if err != nil {
			return errors.Wrap(err, "failed to decode MultiSwap Path arg")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.ExactOutput,
			params.Recipient,
			hops[0].In, hops[len(hops)-1].Out,
			params.AmountInMaximum, params.AmountOut); err != nil {
			return err
		}

	case "swapExactTokensForTokens":
		path := iArgs["path"].([]common.Address)
		if len(path) < 2 {
			return errors.Wrapf(err, "Invalid tx")
		}

		if err = sp.PopulateTradeMethodParams(
			swap.SwapExactTokensForTokens,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			iArgs["amountIn"].(*big.Int), iArgs["amountOutMin"].(*big.Int)); err != nil {
			return err
		}
	case "swapExactTokensForTokensSupportingFeeOnTransferTokens":
		path := iArgs["path"].([]common.Address)
		if err = sp.PopulateTradeMethodParams(
			swap.SwapExactTokensForTokensSupportingFeeOnTransferTokens,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			iArgs["amountIn"].(*big.Int), iArgs["amountOutMin"].(*big.Int)); err != nil {
			return err
		}
	case "swapExactTokensForETHSupportingFeeOnTransferTokens":
		path := iArgs["path"].([]common.Address)
		if len(path) < 2 {
			return errors.Wrapf(err, "Invalid tx")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.SwapExactTokensForETHSupportingFeeOnTransferTokens,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			iArgs["amountIn"].(*big.Int), iArgs["amountOutMin"].(*big.Int)); err != nil {
			return err
		}
	case "swapExactTokensForETH":
		path := iArgs["path"].([]common.Address)
		if len(path) < 2 {
			return errors.Wrapf(err, "Invalid tx")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.SwapExactTokensForETH,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			iArgs["amountIn"].(*big.Int), iArgs["amountOutMin"].(*big.Int)); err != nil {
			return err
		}
	case "swapTokensForExactETH":
		path := iArgs["path"].([]common.Address)

		if err = sp.PopulateTradeMethodParams(
			swap.SwapTokensForExactETH,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			iArgs["amountInMax"].(*big.Int), iArgs["amountOut"].(*big.Int)); err != nil {
			return err
		}
	case "swapTokensForExactTokens":
		path := iArgs["path"].([]common.Address)
		if len(path) < 2 {
			return errors.Wrapf(err, "Invalid tx")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.SwapTokensForExactTokens,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			iArgs["amountInMax"].(*big.Int), iArgs["amountOut"].(*big.Int)); err != nil {
			return err
		}
	case "swapExactETHForTokens":
		path := iArgs["path"].([]common.Address)
		if len(path) < 2 {
			return errors.Wrapf(err, "Invalid tx")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.SwapExactETHForTokens,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			tx.Value(), iArgs["amountOutMin"].(*big.Int)); err != nil {
			return err
		}
	case "swapETHForExactTokens":
		path := iArgs["path"].([]common.Address)
		if len(path) < 2 {
			return errors.Wrapf(err, "Invalid tx")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.SwapETHForExactTokens,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			tx.Value(), iArgs["amountOut"].(*big.Int)); err != nil {
			return err
		}
	case "swapExactETHForTokensSupportingFeeOnTransferTokens":
		path := iArgs["path"].([]common.Address)
		if len(path) < 2 {
			return errors.Wrapf(err, "Invalid tx")
		}

		if err = sp.PopulateTradeMethodParams(
			swap.SwapExactETHForTokensSupportingFeeOnTransferTokens,
			iArgs["to"].(common.Address),
			path[0], path[len(path)-1],
			tx.Value(), iArgs["amountOutMin"].(*big.Int)); err != nil {
			return err
		}
	case "refundETH", "selfPermit", "selfPermitAllowed", "pull", "mint":
	case "sweepTokenWithFee0":
		if !helper.IsSet(sp.Recipient) {
			sp.Recipient = iArgs["recipient"].(common.Address)
		}
		if !helper.IsSet(sp.RxToken) {
			sp.RxToken = iArgs["token"].(common.Address)
		}
	case "sweepToken0":
	case "sweepToken":
		if helper.IsSet(sp.RxToken) {
			if !helper.IsSet(sp.Recipient) {
				sp.Recipient = iArgs["recipient"].(common.Address)
			}
			if !helper.IsSet(sp.RxToken) {
				sp.RxToken = iArgs["token"].(common.Address)
			}
		}

	case "swap":
		desc, err := getOneInchSwapInputData(iArgs["desc"])
		if err != nil {
			return err
		}
		if helper.IsTheSame(desc.SrcToken, oneInchNativeToken) {
			desc.SrcToken = NativeToken
		}
		if helper.IsTheSame(desc.DstToken, oneInchNativeToken) {
			desc.DstToken = NativeToken
		}
		if err = sp.PopulateTradeMethodParams(
			swap.OneInchSwap,
			desc.DstReceiver,
			desc.SrcToken, desc.DstToken,
			desc.Amount, desc.MinReturnAmount); err != nil {
			return err
		}
	case "unoswap":
		switch pools := iArgs["pools"].(type) {
		default:
			path, err := getPoolsInfo(tx, pools)
			if err != nil {
				return errors.Wrapf(err, "Invalid tx")
			}
			if err = sp.PopulateTradeMethodParams(
				swap.OneInchUnoSwap,
				addZero,
				path[0], path[len(path)-1],
				iArgs["amount"].(*big.Int), iArgs["minReturn"].(*big.Int)); err != nil {
				return err
			}
		}
	case "unoswapToWithPermit", "unoswapTo":
		switch pools := iArgs["pools"].(type) {
		default:
			path, err := getPoolsInfo(tx, pools)
			if err != nil {
				return errors.Wrapf(err, "Invalid tx")
			}
			if err = sp.PopulateTradeMethodParams(
				swap.OneInchUnoSwapTo,
				iArgs["recipient"].(common.Address),
				path[0], path[len(path)-1],
				iArgs["amount"].(*big.Int), iArgs["minReturn"].(*big.Int)); err != nil {
				return err
			}
		}
	case "uniswapV3Swap":
		pools := iArgs["pools"].([]*big.Int)
		path, err := getPoolsInfo(tx, pools)
		if err != nil {
			return errors.Wrapf(err, "Invalid tx")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.OneInchUniswapV3Swap,
			addZero,
			path[0], path[len(path)-1],
			iArgs["amount"].(*big.Int), iArgs["minReturn"].(*big.Int)); err != nil {
			return err
		}
	case "uniswapV3SwapTo", "uniswapV3SwapToWithPermit":
		pools := iArgs["pools"].([]*big.Int)
		path, err := getPoolsInfo(tx, pools)
		if err != nil {
			return errors.Wrapf(err, "Invalid tx")
		}
		if err = sp.PopulateTradeMethodParams(
			swap.OneInchUniswapV3SwapTo,
			iArgs["recipient"].(common.Address),
			path[0], path[len(path)-1],
			iArgs["amount"].(*big.Int), iArgs["minReturn"].(*big.Int)); err != nil {
			return err
		}
	case "fillOrderRFQCompact":
		params := OneInchV5.OrderRFQLibOrderRFQ(iArgs["order"].(struct {
			Info          *big.Int       "json:\"info\""
			MakerAsset    common.Address "json:\"makerAsset\""
			TakerAsset    common.Address "json:\"takerAsset\""
			Maker         common.Address "json:\"maker\""
			AllowedSender common.Address "json:\"allowedSender\""
			MakingAmount  *big.Int       "json:\"makingAmount\""
			TakingAmount  *big.Int       "json:\"takingAmount\""
		}))
		if err = sp.PopulateTradeMethodParams(
			swap.OneInchFillOrderRFQCompact,
			addZero,
			params.MakerAsset, params.TakerAsset,
			params.MakingAmount, params.TakingAmount); err != nil {
			return err
		}
	case "addLiquidityETH", "addLiquidity", "increaseLiquidity", "increaseLiquidityCurrentRange":
	case "decreaseLiquidityInHalf", "decreaseLiquidity", "removeLiquidity", "removeLiquidityWithPermit":
		if callback != nil {
			callback.RemoveLiquidityCallBack(tx, swap.RemoveLiquidity{
				Token:     addZero,
				TokenA:    iArgs["tokenA"].(common.Address),
				TokenB:    iArgs["tokenB"].(common.Address),
				Liquidity: iArgs["liquidity"].(*big.Int),
			})
		}

	case "removeLiquidityETHWithPermit", "removeLiquidityETHSupportingFeeOnTransferTokens", "removeLiquidityETH", "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens":
		if callback != nil {
			callback.RemoveLiquidityCallBack(tx, swap.RemoveLiquidity{
				Token:     iArgs["token"].(common.Address),
				TokenA:    addZero,
				TokenB:    addZero,
				Liquidity: iArgs["liquidity"].(*big.Int),
			})
		}
	case "collectAllFees":
	case "unwrapWETH90":
	case "unwrapWETH9", "unwrapWETH9WithFee":
		sp.Recipient = iArgs["recipient"].(common.Address)
	case "wrapETH":
	case "collectRewards", "uniswapV3SwapCallback":
	case "cancelOrder":
	default:
		//fmt.Println(tx.Hash())
		//spew.Dump(iArgs)
		//fmt.Println("warn: method not handled:", method.Name)
		//os.Exit(1)
	}
	if helper.IsSet(sp.SwapToken) && callback != nil {
		callback.SwapCallBack(tx, sp)
	}
	return nil
}

func getHopsFromMultiSwapPath(path []byte, isReversed bool) (hops []*poolHop, err error) {
	if len(path) < 3+(2*common.AddressLength) {
		err = errors.New("not long enough for first pool")
		return
	}

	hops = make([]*poolHop, 1, 2)
	hops[0] = new(poolHop)

	if isReversed {
		hops[0].Out = common.BytesToAddress(path[:common.AddressLength])
	} else {
		hops[0].In = common.BytesToAddress(path[:common.AddressLength])
	}

	//fmt.Print("multipath swap: {", swaps[0].In.Hex())
	//defer fmt.Println("}")

	iPath := common.AddressLength
	//Fee := feeFromPath(path[iPath:])
	iPath += 3
	//fmt.Printf(" - %d - ", Fee)

	if isReversed {
		hops[0].In = common.BytesToAddress(path[iPath : iPath+common.AddressLength])
	} else {
		hops[0].Out = common.BytesToAddress(path[iPath : iPath+common.AddressLength])
	}
	iPath += common.AddressLength
	//fmt.Println(swaps[0].Out.Hex())

	// more hops
	for iPath < len(path) {
		if len(path)-iPath < 3+common.AddressLength {
			err = errors.Errorf("not long enough for Fee and second address at offset %d", iPath)
			return
		}

		//Fee := feeFromPath(path[iPath:])
		iPath += 3
		//fmt.Printf(" - %d - ", Fee)

		ph := new(poolHop)
		if isReversed {
			ph.In = common.BytesToAddress(path[iPath : iPath+common.AddressLength])
			ph.Out = hops[len(hops)-1].In
		} else {
			ph.In = hops[len(hops)-1].Out
			ph.Out = common.BytesToAddress(path[iPath : iPath+common.AddressLength])
		}

		hops = append(hops, ph)
		iPath += common.AddressLength

		//fmt.Print(swaps[len(swaps)-1].Out.Hex())
	}

	if isReversed {
		for i, j := 0, len(hops)-1; i < j; i, j = i+1, j-1 {
			hops[i], hops[j] = hops[j], hops[i]
		}
	}
	return
}
func getOneInchSwapInputData(d interface{}) (result oneInchSwapInput, err error) {
	jsonData, err := json.Marshal(d)
	if err != nil {
		return
	}
	err = json.Unmarshal(jsonData, &result)
	return
}

func getPoolsInfo(tx *types.Transaction, v any) (path []common.Address, err error) {
	var ps []string
	switch ev := v.(type) {
	case []*big.Int:
		for _, v := range ev {
			ps = append(ps, v.Text(16))
		}
	case [][32]uint8:
		for _, v := range ev {
			ps = append(ps, fmt.Sprintf("%x", v))
		}
	}

	if len(ps) == 0 {
		return path, errors.New("Empty pool")
	}

	for _, p := range ps {
		in, out, err := getPoolPath(p)
		if err != nil {
			return path, err
		}
		path = append(path, in, out)
	}
	return
}

func getPoolPath(oneInchRaw string) (tokenIn, tokenOut common.Address, err error) {
	poolReverseFlag := false
	var poolAddr common.Address
	if len(oneInchRaw) > 40 {
		lastPoolFlag := oneInchRaw[:2]
		if lastPoolFlag == "c0" || lastPoolFlag == "a0" || lastPoolFlag == "80" {
			poolReverseFlag = true
		}
		poolAddr = common.HexToAddress(oneInchRaw[len(oneInchRaw)-40:])
	} else {
		poolAddr = common.HexToAddress(oneInchRaw)
	}
	lp, err := getPoolInfo(poolAddr)
	if err != nil {
		return addZero, addZero, err
	}
	if poolReverseFlag {
		tokenOut = lp.Token0
		tokenIn = lp.Token1
	} else {
		tokenOut = lp.Token1
		tokenIn = lp.Token0
	}
	return
}

func getPoolInfo(addr common.Address) (*poolPair, error) {
	if v, ok := pairsAndPools.Get(addr.String()); ok {
		return v, nil
	}
	instance, err := uni.NewUni(addr, C)
	if err != nil {
		return nil, err
	}
	token0, err := instance.Token0(nil)
	if err != nil {
		return nil, err
	}
	token1, err := instance.Token1(nil)
	if err != nil {
		return nil, err
	}
	pp := &poolPair{
		Token0: token0,
		Token1: token1,
	}
	pairsAndPools.Set(addr.String(), pp)
	return pp, nil
}
