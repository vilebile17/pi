package main

import (
	"math"
	"math/rand"
)

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
