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
	return (p.circlePi.value + p.eulerPi.value) / 2
}

type CirclePi struct {
	width int
	inCircle int
	value float64
}
func (c CirclePi) calculate(i int) {
	centre := (float64(c.width) + 1.0) / 2.0

	x := float64(rand.Intn(c.width - 1) + 1)
	y := float64(rand.Intn(c.width - 1) + 1)
	distance := math.Pow(math.Pow(x - centre, 2) + math.Pow(y - centre, 2), 0.5)

	if float64(distance) <= centre {
		c.inCircle++
	}

	c.value = 4.0 * float64(c.inCircle) / float64(i)
}

type EulerPi struct {
	sumSoFar float64
	value float64
}
func (e EulerPi) calculate(i int) (float64, float64){
	sum := e.sumSoFar + (1.0 / float64(math.Pow(float64(i), 2)))
	return sum, math.Pow(e.sumSoFar * 6.0, 0.5)
}


func main() {
	// creating the Pi struct
	pi := Pi{
		circlePi: CirclePi{
			width: 9999, 
			inCircle: 0, 
			value: 0.0,
		}, 
		eulerPi: EulerPi{
			sumSoFar: 0.0,
			value: 0.0,
		},
	}


	i := 0

	for true {
		pi.circlePi.calculate(i)
		pi.eulerPi.calculate(i)

		fmt.Printf("circlePi: %v\n", pi.circlePi.value)
		fmt.Printf("eulerPi: %v\n", pi.eulerPi.value)

		i++
	}
}
