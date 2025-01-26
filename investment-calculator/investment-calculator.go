package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 3.5

func main() {
	var investmentAmount float64
	var years float64
	var annualInterestRate float64

	fmt.Print("Enter investment amount: ")
	fmt.Scanf("%f", &investmentAmount)

	fmt.Print("Enter annual interest rate in percentage: ")
	fmt.Scanf("%f", &annualInterestRate)

	fmt.Print("Enter number of years: ")
	fmt.Scanf("%f", &years)

	var futureInvestmentValue float64 = calculateInvestment(investmentAmount, annualInterestRate, years)
	var realFutureInvestmentValue float64 = calculateInvestment(futureInvestmentValue, inflationRate, years)

	fmt.Printf("Accumulated value is %.0f\n", futureInvestmentValue)
	fmt.Printf("Real accumulated value is %.0f\n", realFutureInvestmentValue)
}

func calculateInvestment(investmentAmount float64, annualInterestRate float64, years float64) float64 {
	return investmentAmount * math.Pow(1+annualInterestRate/100, years)
}
