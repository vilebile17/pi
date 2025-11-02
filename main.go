package main

import (
	"github.com/fatih/color"
	"fmt"
) 

func main() {
	// for the cool printing effects...
	color.Cyan("Calculating pi...")
	fmt.Println("")

	// creates the eulerPi struct
	eulerPi := EulerPi {
		sumSoFar: 0.0,
		val: 0.0,
	}

	// calculation loop
	for i:=1; true ; i++ {
		eulerPi.sumSoFar += eulerPi.calculateNextVal(i)

		if (i + 1) % 100000 == 0 {
			eulerPi.val = eulerPi.calculatePi()
			fmt.Printf("\r===  ℼ = %.20f  ===", eulerPi.val)
		}
		
	}
}
