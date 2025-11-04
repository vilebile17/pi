package main

import (
	"math/big"
)

func Sqrt(num *big.Float) *big.Float {
	two := new(big.Float).SetPrec(256).SetFloat64(2.0)
	z := new(big.Float).SetPrec(256).Quo(num, two)
	temp := new(big.Float).SetPrec(256)
	for _ = range 10 {
		temp.Quo(num, z)
		temp.Add(z, temp)
		z.Quo(temp, two)
	}
	return z
}
