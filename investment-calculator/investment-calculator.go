package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 3.5

	investmentAmount := 1000.0
	years := 5.0
	annualInterestRate := 4.25

	var futureInvestmentValue float64 = investmentAmount * math.Pow(1+annualInterestRate/100, years)
	var realFutureInvestmentValue float64 = futureInvestmentValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Accumulated value is", futureInvestmentValue)
	fmt.Println("Real accumulated value is", realFutureInvestmentValue)
}
