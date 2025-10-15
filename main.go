package main

import (
	"fmt"
	"math"
	"math/rand"
)

// This struct contains other structs for each method of calculating pi
type Pi struct {
	circlePi CirclePi 
	eulerPi EulerPi 
}
func (p Pi) getMean() float64 {
	return (p.circlePi.val + p.eulerPi.val) / 2
}

// The first method implemented, makes a square with an inscribed circle in it and generates a random point in the square and sees if it is in the circle
type CirclePi struct {
	width int
	centre float64 
	inCircle float64
	val float64
}
func (c CirclePi) calculate(i int) (float64, float64) {
	x := float64(rand.Intn(c.width))
	y := float64(rand.Intn(c.width))
	distanceToCentre := float64(math.Pow(math.Pow(x - c.centre, 2) + math.Pow(y - c.centre, 2), 0.5))

	newInCircle := c.inCircle
	if distanceToCentre <= c.centre {
		newInCircle++
	}

	return newInCircle, 4.0 * newInCircle / (float64(i) + 1)
}

// The second method implemented, uses Leonhard Euler's formula: pi² / 6 = 1/1 + 1/4 + 1/9 + 1/16 + 1/25 ...
type EulerPi struct {
	sumSoFar float64
	val float64
}
func (e EulerPi) calculate(i int) (float64, float64) {
	sum := e.sumSoFar + float64(math.Pow(float64(i), -2.0))
	return sum, math.Pow(sum * 6.0, 0.5)	
}


func main() {
	// creating the Pi struct
	pi := Pi{
		circlePi: CirclePi{
			width: 6942067, 
			inCircle: 0.0, 
		}, 
		eulerPi: EulerPi{
			sumSoFar: 0.0,
		},
	}
	pi.circlePi.centre = (float64(pi.circlePi.width) + 1.0) / 2 

	for i:=1; i < 999999999999999; i++ {
		// here we calculate pi using each of the methods
		pi.circlePi.inCircle, pi.circlePi.val = pi.circlePi.calculate(i)
		pi.eulerPi.sumSoFar, pi.eulerPi.val = pi.eulerPi.calculate(i)

		fmt.Printf("ℼ = %v\n", pi.eulerPi.val)
	}
}
