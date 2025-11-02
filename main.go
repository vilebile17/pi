package main

import (
	"github.com/fatih/color"
	"fmt"
) 

// This struct contains other structs for each method of calculating pi
type Pi struct {
	circlePi CirclePi 
	eulerPi EulerPi 
}
func (p Pi) getMean() float64 {
	return (p.circlePi.val + p.eulerPi.val) / 2
}

func main() {
	// for the cool printing effects...
	color.Cyan("Calculating pi...")
	fmt.Println("")

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

	for i:=1; true ; i++ {
		pi.eulerPi.sumSoFar += pi.eulerPi.calculateNextVal(i)

		pi.eulerPi.val = pi.eulerPi.calculatePi()
		if (i + 1) % 10000 == 0 {
			fmt.Printf("\r===  ℼ = %.20f  ===", pi.eulerPi.val)
		}
		
	}
}
