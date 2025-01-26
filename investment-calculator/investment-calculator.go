package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 3.5

	var investmentAmount float64
	var years float64
	var annualInterestRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scanf("%f", &investmentAmount)

	fmt.Print("Enter annual interest rate in percentage: ")
	fmt.Scanf("%f", &annualInterestRate)

	fmt.Print("Enter number of years: ")
	fmt.Scanf("%f", &years)

	var futureInvestmentValue float64 = investmentAmount * math.Pow(1+annualInterestRate/100, years)
	var realFutureInvestmentValue float64 = futureInvestmentValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("Accumulated value is %.2f\n", futureInvestmentValue)
	fmt.Printf("Real accumulated value is %.2f\n", realFutureInvestmentValue)
}
