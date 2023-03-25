package swap

type SwapMethod int

const (
	NotSet SwapMethod = iota
	Multicall
	Execute
	V3SwapExactIn
	V3SwapExactOut
	V2SwapExactIn
	V2SwapExactOut
	ExactInputSingle
	ExactInput
	ExactOutputSingle
	ExactOutput
	SwapExactTokensForTokens
	SwapExactTokensForTokensSupportingFeeOnTransferTokens
	SwapExactTokensForETHSupportingFeeOnTransferTokens
	SwapExactTokensForETH
	SwapTokensForExactETH
	SwapTokensForExactTokens
	SwapExactETHForTokens
	SwapETHForExactTokens
	SwapExactETHForTokensSupportingFeeOnTransferTokens
	OneInchSwap
	OneInchUnoSwap
	OneInchUnoSwapTo
	OneInchUniswapV3Swap
	OneInchUniswapV3SwapTo
	OneInchFillOrderRFQCompact
)
