package mt

import (
	"github.com/ethereum/go-ethereum/common"
)

var (
	NativeToken, USDT, USDC, WBTC, LUSD, BUSD common.Address
	OneInchSpotPriceAddress                   common.Address
)

func setupContracts() error {
	switch chainId {
	case 1: //ETH
		NativeToken = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
		OneInchSpotPriceAddress = common.HexToAddress("0x07D91f5fb9Bf7798734C3f606dB065549F6893bb")
	case 56: //BSC
		NativeToken = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
		OneInchSpotPriceAddress = common.HexToAddress("0xfbD61B037C325b959c0F6A7e69D8f37770C2c550")
	}
	return nil
}
