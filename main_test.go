package mt

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestKeyStore(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("674b85e88577689ada4a054b732bccdca31bb0a8d4088b3c1ebb1c99cf6f33bd")
	if err != nil {
		panic(err)
	}
	k, _ := GenKeyStore(privateKey, "x", "y")
	spew.Dump(k)
	spew.Dump(GetAccountFromKeyStore(k, "x", "y"))
}
