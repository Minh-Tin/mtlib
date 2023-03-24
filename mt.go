package mt

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gorilla/websocket"
	"math/big"
)

var (
	Ctx     context.Context
	R       *rpc.Client
	C       *ethclient.Client
	G       *gethclient.Client
	W       *websocket.Conn
	ChainId *big.Int
)
var (
	chainId int64
)

func Setup(RpcEndpoint string) error {
	var err error
	Ctx = context.Background()
	R, err = rpc.DialContext(Ctx, RpcEndpoint)
	if err != nil {
		return err
	}
	C = ethclient.NewClient(R)
	G = gethclient.New(R)
	ChainId, err = C.ChainID(Ctx)
	if err != nil {
		return err
	}
	chainId = ChainId.Int64()
	return setupContracts()
}

func mustSetup() {
	if ChainId == nil {
		panic("Setup is required")
	}
}
