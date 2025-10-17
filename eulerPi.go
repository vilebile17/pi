package main

import (
	"math"
)

type EulerPi struct {
	sumSoFar float64
	val float64
}
// Uses Leonhard Euler's formula: pi²/6 = 1/1² + 1/2² + 1/3² + 1/4² ...
// This func calculates each fraction i.e. 1/i²
func (e EulerPi) calculateNextVal(i int) float64 {
	return float64(math.Pow(float64(i), -2.0))
}
// Uses the sumSoFar value to calculate pi
func (e EulerPi) calculatePi() float64 {
	return math.Pow(e.sumSoFar * 6,0.5)
}
