package mt

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestTraceTxHashToInterface(t *testing.T) {
	setup()
	it, err := TraceTxHashToInterface(common.HexToHash("0xc23a6812237be6995c9a76be214682831075074faafb244f0632b3f83bcbe9d2"))
	if err != nil {
		panic(err)
	}
	spew.Dump(it)
}
