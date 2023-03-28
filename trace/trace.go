package trace

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Call struct {
	Type         string          `json:"type"`
	From         common.Address  `json:"from"`
	To           *common.Address `json:"to"`
	Input        string          `json:"input"`
	Output       string          `json:"output"`
	Value        *hexutil.Big    `json:"value"`
	GasUsed      *hexutil.Big    `json:"gasUsed"`
	Revert       bool            `json:"revert"`
	RevertReason string          `json:"revertReason"`
	Error        string          `json:"error,omitempty"`
	Calls        []*Call         `json:"calls,omitempty"`
}
