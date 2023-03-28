package mt

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func TestGetLiq(t *testing.T) {
	setup()
	spew.Dump(MtLiqV2.GetBestLiq(common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), common.HexToAddress("0xdfe691F37b6264a90Ff507EB359C45d55037951C")))
}
