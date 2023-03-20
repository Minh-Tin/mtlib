package mt

import "github.com/ethereum/go-ethereum/core/types"

func DecodeSwapByInput(tx *types.Transaction) {
	if tx.To() == nil {
		return
	}

}
