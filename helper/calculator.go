package helper

import (
	"math/big"
)

func Amount2Human(amount, decimal *big.Int) *big.Float {
	decimapExp := new(big.Int).Exp(big.NewInt(10), decimal, nil)
	if decimapExp.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}
	return new(big.Float).Quo(new(big.Float).SetInt(amount), new(big.Float).SetInt(decimapExp))
}

func CalcPrice(baseAmount, quoteAmount, baseDecimal, quoteDecimal *big.Int) *big.Float {
	if quoteAmount == nil || baseAmount == nil {
		return big.NewFloat(0)
	}
	basePerDecimal := Amount2Human(baseAmount, baseDecimal)
	quotePerDecimal := Amount2Human(quoteAmount, quoteDecimal)
	if quotePerDecimal.Cmp(big.NewFloat(0)) == 0 {
		return big.NewFloat(0)
	}
	return new(big.Float).Quo(quotePerDecimal, basePerDecimal)
}
