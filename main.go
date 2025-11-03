package main

import (
	"github.com/fatih/color"
	"fmt"
	"math/big"
) 

func main() {
	// for the cool printing effects...
	color.Green("Calculating pi...")
	fmt.Println("")

	sumSoFar := new(big.Float).SetPrec(256).SetFloat64(0)	
	one := new(big.Float).SetPrec(256).SetFloat64(1)	
	six := new(big.Float).SetPrec(256).SetFloat64(6)	

	for i := 1; true; i++ {
		newI := new(big.Float).SetPrec(256).SetFloat64(float64(i))
		newI.Mul(newI, newI) //square it

		// dividing
		result := new(big.Float).SetPrec(256)
		result.Quo(one, newI)
		sumSoFar.Add(sumSoFar, result)

		if i % 10000 == 0 {
			piSquared := new(big.Float).SetPrec(256)
			piSquared.Mul(six, sumSoFar)
			fmt.Println("\r pi = %v", piSquared)
		}
	}
}
