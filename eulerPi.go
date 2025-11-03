package main

import (
	"math"
	"math/big"
)

type EulerPi struct {
	sumSoFar float64
	val float64
}
// Uses Leonhard Euler's formula: pi²/6 = 1/1² + 1/2² + 1/3² + 1/4² ...
// This func calculates each fraction i.e. 1/i²
func (e *EulerPi) calculateNextVal(i int) {
	e.sumSoFar += float64(math.Pow(float64(i), -2.0))
}
// Uses the sumSoFar value to calculate pi
func (e *EulerPi) calculatePi() {
	e.val = math.Pow(e.sumSoFar * 6,0.5)
}

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
