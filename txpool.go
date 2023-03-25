package mt

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func SubscribePendingTransactions(callback func(hash common.Hash)) error {
	mustSetup()
	pendingTxsChan := make(chan common.Hash)
	sub, txSubErr := G.SubscribePendingTransactions(context.Background(), pendingTxsChan)
	if txSubErr != nil {
		return txSubErr
	}
	defer sub.Unsubscribe()
	for pendingTx := range pendingTxsChan {
		go callback(pendingTx)
	}
	return nil
}

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

func ContentFrom(addr common.Address) (map[string]map[string]*types.Transaction, error) {
	var result json.RawMessage
	err := R.CallContext(context.Background(), &result, "txpool_contentFrom", addr)
	if err != nil {
		return nil, err
	}
	var data map[string]map[string]*types.Transaction
	if err = json.Unmarshal(result, &data); err != nil {
		return nil, err
	}
	return data, nil
}
