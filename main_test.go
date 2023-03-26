package mt

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/Minh-Tin/mtlib/ethutil"
	"github.com/Minh-Tin/mtlib/swap"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"
)

func setup() {
	Setup("ws://0xbmc.com:8845")
}

func TestDecodeSwapByInput(t *testing.T) {
	setup()
	//tx, _, _ := C.TransactionByHash(Ctx, common.HexToHash("0x418fcb716d3d502e389e9b21672bda26029bdd3e3b740df4cd917515509d8346"))
	tx, _, _ := C.TransactionByHash(Ctx, common.HexToHash("0xf0a71db6f3cb9d6b234853d9605a13e28ced7b55a81ffb6083c52d966e94055e"))
	spew.Dump(DecodeSwapByInput(tx, nil))
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

func TestBackWard(t *testing.T) {
	setup()
	BackWard(16891388, func(tx *types.Transaction) {
		if sp, err := DecodeSwapByInput(tx, nil); err == nil {
			if sp.Method > -1 && !swap.IsValidRecipient(sp.Recipient) {
				spew.Dump(sp)
				panic(tx.Hash())
			}
		}
	})
}

func TestAccount(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	a := Account{
		PublicKey:  address,
		PrivateKey: privateKey,
	}
	test, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(test))
	return
	var x Account
	json.Unmarshal(test, &x)
	spew.Dump(x)
}

//0x57b5d8c93677e00ac364e2e62d4ca8aeed6ab8942557827125c393bec22bc549 v2 + v3
