package mt

import (
	"errors"
	"github.com/Minh-Tin/mtlib/swap"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type SwapParams struct {
	Tx         *types.Transaction
	Sender     common.Address
	Recipient  common.Address
	SwapToken  common.Address
	RxToken    common.Address
	SwapAmount *big.Int // AmountIn / AmountInMax
	RxAmount   *big.Int // AmountOut / AmountOutMin
	Method     swap.SwapMethod
	Calls      []struct {
		Method swap.SwapMethod
		Path   []common.Address
	}
}

func DecodeSwapByInput(tx *types.Transaction) error {
	if tx.To() == nil {
		return errors.New("Invalid transaction")
	}
	return nil
}
