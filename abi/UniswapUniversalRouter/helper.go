package UniswapUniversalRouter

import "github.com/umbracle/ethgo/abi"

var V3_SWAP_EXACT_IN = abi.MustNewType("tuple(address recipient, uint256 amountIn, uint256 amountOutMin, bytes path, bool payerIsUser)")
var V3_SWAP_EXACT_OUT = abi.MustNewType("tuple(address recipient, uint256 amountOut, uint256 amountInMax, bytes path, bool payerIsUser)")
var V2_SWAP_EXACT_IN = abi.MustNewType("tuple(address recipient, uint256 amountIn, uint256 amountOutMin, address[] path, bool payerIsUser)")
var V2_SWAP_EXACT_OUT = abi.MustNewType("tuple(address recipient, uint256 amountOut, uint256 amountInMax, address[] path, bool payerIsUser)")
var UNWRAP_ETH = abi.MustNewType("tuple(address recipient, uint256 minAmount)")
var SWEEP = abi.MustNewType("tuple(address token, address recipient, uint256 minAmount)")
