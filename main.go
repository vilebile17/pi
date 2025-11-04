package main

import (
	"github.com/fatih/color"
	"fmt"
	"math/big"
) 

func main() {
	color.Green("Calculating pi...")
	fmt.Println("")

	sumSoFar := new(big.Float).SetPrec(256).SetFloat64(0)	
	one := new(big.Float).SetPrec(8).SetFloat64(1)	
	six := new(big.Float).SetPrec(8).SetFloat64(6)	

	for i := 1; true; i++ {
		// calculates i²
		result := new(big.Float).SetPrec(32).SetFloat64(float64(i))
		result.Mul(result, result) //square it

		// calculates 1/i² and adds it to the running total
		result.Quo(one, result)
		sumSoFar.Add(sumSoFar, result)

		if i % 30000 == 0 {
			// solves the equation: pi²/6 = sumSoFar  
			// And then prints the result
			piSquared := new(big.Float).SetPrec(256).Mul(six, sumSoFar)
			result := Sqrt(piSquared)
			fmt.Printf("\r--- ℼ = %.15f ---", result)
		}
	}
}
