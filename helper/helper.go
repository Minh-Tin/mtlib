package helper

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func IsSet(addr common.Address) bool {
	return !IsTheSame(addr, common.Address{})
}

func IsTheSame(addr1, addr2 common.Address) bool {
	return bytes.Equal(addr1.Bytes(), addr2.Bytes())
}

func SortsBefore(tokenA, tokenB common.Address) bool {
	return strings.ToLower(tokenA.String()) < strings.ToLower(tokenB.String())
}
