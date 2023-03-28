package mtLiqV2

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var ZeroContract = common.HexToAddress("0x0000000000000000000000000000000000000000")

func (mtLiq *MtLiqV2) GetBestLiq(token0, token1 common.Address) (common.Address, *big.Int, bool, *big.Int) {
	pool, err := mtLiq.Liquidity(&bind.CallOpts{}, token0, token1)
	if err != nil {
		return common.Address{}, nil, false, nil
	}
	idx := 0
	for i := 1; i < len(pool); i++ {
		if pool[idx].Liq.Cmp(pool[i].Liq) < 0 ||
			(pool[idx].Liq.Cmp(big.NewInt(0)) == 0 && pool[i].Pool != ZeroContract) {
			idx = i
		}
	}
	return pool[idx].Pool, pool[idx].Liq, pool[idx].Version == 3, pool[idx].Fee
}
