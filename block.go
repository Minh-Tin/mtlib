package mt

import (
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func BackWard(customNumber uint64, callback func(tx *types.Transaction)) (err error) {
	mustSetup()
	if customNumber == 0 {
		customNumber, err = C.BlockNumber(Ctx)
		if err != nil {
			return err
		}
	}

	for i := customNumber; i >= 0; i-- {
		block, err := C.BlockByNumber(Ctx, new(big.Int).SetUint64(i))
		if err != nil {
			return err
		}
		for _, tx := range block.Transactions() {
			callback(tx)
		}
	}
	return nil
}
