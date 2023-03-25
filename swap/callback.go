package swap

import "github.com/ethereum/go-ethereum/core/types"

type CallBack interface {
	SwapCallBack(tx *types.Transaction, sp *Params)
	RemoveLiquidityCallBack(tx *types.Transaction, lq RemoveLiquidity)
}
