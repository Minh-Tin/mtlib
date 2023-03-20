package mt

import (
	"github.com/ethereum/go-ethereum/common"
	"mtlib/abi/OneInchSpotPrice"
)

var (
	NativeToken             common.Address
	OneInchSpotPriceAddress common.Address
	StableCoins             = make(map[common.Address]*Token)
	SpotPrice               *OneInchSpotPrice.OneInchSpotPrice
)

func setupContracts() error {
	switch chainId {
	case 1: //ETH
		NativeToken = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
		OneInchSpotPriceAddress = common.HexToAddress("0x07D91f5fb9Bf7798734C3f606dB065549F6893bb")
		addStableCoin("0xdac17f958d2ee523a2206206994597c13d831ec7") // USDT
		addStableCoin("0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") // USDC
		addStableCoin("0x6b175474e89094c44da98b954eedeac495271d0f") // DAI
		addStableCoin("0x4Fabb145d64652a948d72533023f6E7A623C7C53") // BUSD
		addStableCoin("0x0000000000085d4780B73119b644AE5ecd22b376") // TUSD
		addStableCoin("0x8e870d67f660d95d5be530380d0ec0bd388289e1") // Pax
		addStableCoin("0x0c10bf8fcb7bf5412187a595ab97a3609160b5c6") // USĐ
		addStableCoin("0x853d955acef822db058eb8505911ed77f175b99e") // Frax
		addStableCoin("0x5f98805A4E8be255a32880FDeC7F6728C6568bA0") // LUSD

	case 56: //BSC
		NativeToken = common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c")
		OneInchSpotPriceAddress = common.HexToAddress("0xfbD61B037C325b959c0F6A7e69D8f37770C2c550")
		addStableCoin("0x55d398326f99059fF775485246999027B3197955") //USDT
		addStableCoin("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56") //BUSD
		addStableCoin("0x14016E85a25aeb13065688cAFB43044C2ef86784") //TUSD
		addStableCoin("0x90C97F71E18723b0Cf0dfa30ee176Ab653E89F40") //FRAX
		addStableCoin("0xb3c11196A4f3b1da7c23d9FB0A3dDE9c6340934F") //PAX
		addStableCoin("0xd17479997F34dd9156Deef8F95A52D81D265be9c") //USDD
		addStableCoin("0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3") //DAI
	}
	var err error
	SpotPrice, err = OneInchSpotPrice.NewOneInchSpotPrice(OneInchSpotPriceAddress, C)
	if err != nil {
		return err
	}
	return nil
}

func addStableCoin(addr string) {
	address := common.HexToAddress(addr)
	if _, ok := StableCoins[address]; ok {
		return
	}
	token, err := GetToken(address)
	if err != nil {
		return
	}
	StableCoins[address] = token
}
