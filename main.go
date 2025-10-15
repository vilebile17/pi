package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Pi struct {
	circlePi CirclePi 
	eulerPi EulerPi 
}
func (p Pi) getMean() float64 {
	return (p.circlePi.val + p.eulerPi.val) / 2
}

type CirclePi struct {
	width int
	centre float64 
	inCircle float64
	val float64
}
func (c CirclePi) calculate(i int) (float64, float64) {
	x := float64(rand.Intn(9999))
	y := float64(rand.Intn(9999))
	distanceToCentre := float64(math.Pow(math.Pow(x - c.centre, 2) + math.Pow(y - c.centre, 2), 0.5))

	newInCircle := c.inCircle
	if distanceToCentre <= c.centre {
		newInCircle++
	}

	return newInCircle, 4.0 * newInCircle / (float64(i) + 1)
}

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
			width: 9999, 
			inCircle: 0.0, 
		}, 
		eulerPi: EulerPi{
			sumSoFar: 0.0,
		},
	}
	pi.circlePi.centre = (float64(pi.circlePi.width) + 1.0) / 2 

	for i:=1; i < 999999999999999; i++ {
		// here we calculate pi using each of the two formulas
		pi.circlePi.inCircle, pi.circlePi.val = pi.circlePi.calculate(i)
		pi.eulerPi.sumSoFar, pi.eulerPi.val = pi.eulerPi.calculate(i)

		fmt.Printf("circlePi: %v\n", pi.circlePi.val)
		fmt.Printf("eulerPi: %v\n", pi.eulerPi.val)
	}
}
