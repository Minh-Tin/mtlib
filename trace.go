package mt

import (
	"github.com/Minh-Tin/mtlib/ethutil"
	"github.com/Minh-Tin/mtlib/trace"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"math/big"
)

var (
	tracerOpt     = "callTracer"
	tracerTimeout = "40s"
	traceConfig   = tracers.TraceConfig{Tracer: &tracerOpt}
)

func TxToCallMsg(tx *types.Transaction) (msg ethereum.CallMsg, err error) {
	from, err := ethutil.GetTxSender(tx)
	if err != nil {
		return msg, err
	}
	msg.From = from
	msg.To = tx.To()
	msg.Value = tx.Value()
	msg.Gas = tx.Gas()
	msg.GasPrice = tx.GasPrice()
	msg.Data = tx.Data()
	msg.GasTipCap = tx.GasTipCap()
	msg.GasFeeCap = tx.GasFeeCap()
	msg.AccessList = tx.AccessList()
	return msg, nil
}

func TxToCallArg(tx *types.Transaction) (interface{}, error) {
	from, err := ethutil.GetTxSender(tx)
	if err != nil {
		return nil, err
	}
	arg := map[string]interface{}{
		"from": from,
		"to":   tx.To(),
	}
	if len(tx.Data()) > 0 {
		arg["data"] = hexutil.Bytes(tx.Data())
	}
	if tx.Value() != nil {
		arg["value"] = (*hexutil.Big)(tx.Value())
	}
	if tx.Gas() != 0 {
		arg["gas"] = hexutil.Uint64(tx.Gas())
	}
	if tx.GasPrice() != nil {
		arg["gasPrice"] = (*hexutil.Big)(tx.GasPrice())
	}
	if tx.GasFeeCap() != nil {
		arg["gasFeeCap"] = (*hexutil.Big)(tx.GasFeeCap())
	}
	if tx.GasTipCap() != nil {
		arg["gasTipCap"] = (*hexutil.Big)(tx.GasTipCap())
	}
	if tx.AccessList() != nil {
		arg["accessList"] = tx.AccessList()
	}
	return arg, nil
}

func TraceCall(tx *types.Transaction) (traceResult *trace.Call, err error) {
	arg, err := TxToCallArg(tx)
	if err != nil {
		return
	}
	if arg == nil {
		return
	}
	err = R.CallContext(Ctx, &traceResult, "debug_traceCall", arg, "latest", &traceConfig)
	return
}

func TraceTx(tx *types.Transaction) (traceResult *trace.Call, err error) {
	return TraceTxHash(tx.Hash())
}

func TraceTxHash(txHash common.Hash) (traceResult *trace.Call, err error) {
	err = R.CallContext(Ctx, &traceResult, "debug_traceTransaction", txHash, &traceConfig)
	if err != nil {
		panic(err)
	}
	return
}
func TraceTxHashToInterface(txHash common.Hash) (traceResult interface{}, err error) {
	err = R.CallContext(Ctx, &traceResult, "debug_traceTransaction", txHash, &traceConfig)
	if err != nil {
		panic(err)
	}
	return
}

func TraceBlockNumber(blockNumber uint64) {
	var result interface{}
	err := R.CallContext(Ctx, &result, "debug_traceBlockByNumber", (*hexutil.Big)(new(big.Int).SetUint64(blockNumber)).String(), &traceConfig)
	if err != nil {
		panic(err)
	}
	spew.Dump(result)
}
