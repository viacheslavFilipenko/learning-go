package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var annualInterestRate float64 = 4.25
	var numberOfYears int = 10

	var futureInvestmentValue float64 = investmentAmount * math.Pow(1+annualInterestRate/100, float64(numberOfYears))

	fmt.Println("Accumulated value is", futureInvestmentValue)
}
