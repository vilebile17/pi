package main

import (
	"math"
)

type EulerPi struct {
	sumSoFar float64
	val float64
}
// Uses Leonhard Euler's formula: pi²/6 = 1/1² + 1/2² + 1/3² + 1/4² ...
func (e EulerPi) calculate(i int) (float64, float64) {
	sum := e.sumSoFar + float64(math.Pow(float64(i), -2.0))
	return sum, math.Pow(sum * 6.0, 0.5)	
}
