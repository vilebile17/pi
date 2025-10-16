package main

import "fmt"

// This struct contains other structs for each method of calculating pi
type Pi struct {
	circlePi CirclePi 
	eulerPi EulerPi 
}
func (p Pi) getMean() float64 {
	return (p.circlePi.val + p.eulerPi.val) / 2
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

	for i:=1; true ; i++ {
		// here we calculate pi using each of the methods
		pi.circlePi.inCircle, pi.circlePi.val = pi.circlePi.calculate(i)
		pi.eulerPi.sumSoFar, pi.eulerPi.val = pi.eulerPi.calculate(i)

		fmt.Printf("ℼ = %v\n", pi.eulerPi.val)
	}
}
