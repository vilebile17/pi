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

	ch := make(chan struct{})
	for i:=1; true ; i+=2 {
		// eulerPi can be calculated in parallel so we can use goroutines!
		go func() {
			pi.eulerPi.sumSoFar += pi.eulerPi.calculateNextVal(i)
			ch <- struct{}{}
		}()

		pi.eulerPi.sumSoFar += pi.eulerPi.calculateNextVal(i+1)
		<-ch

		pi.eulerPi.val = pi.eulerPi.calculatePi()
		
		fmt.Printf("ℼ = %v\n", pi.eulerPi.val)
	}
}
