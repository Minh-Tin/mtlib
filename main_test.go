package mt

import (
	"github.com/Minh-Tin/mtlib/ethutil"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func setup() {
	Setup("ws://0xbmc.com:8845")
}

func TestKeyStore(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("674b85e88577689ada4a054b732bccdca31bb0a8d4088b3c1ebb1c99cf6f33bd")
	if err != nil {
		panic(err)
	}
	k, _ := GenKeyStore(privateKey, "x", "y")
	spew.Dump(k)
	spew.Dump(GetAccountFromKeyStore(k, "x", "y"))
}

func TestGetTxSender(t *testing.T) {
	setup()
	tx, _, _ := C.TransactionByHash(Ctx, common.HexToHash("0x17badf61ba431c4d42148528359a906e68bab0fa17e9eb104e3332779c397398"))
	spew.Dump(ethutil.GetTxSender(tx))
}
