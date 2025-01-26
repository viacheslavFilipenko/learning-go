package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 3.5

func main() {
	investmentAmount := getUserInput("Enter investment amount: ")
	annualInterestRate := getUserInput("Enter annual interest rate in percentage: ")
	years := getUserInput("Enter number of years: ")

	var futureInvestmentValue float64 = calculateInvestment(investmentAmount, annualInterestRate, years)
	var realFutureInvestmentValue float64 = calculateInvestment(futureInvestmentValue, inflationRate, years)

	fmt.Printf("Accumulated value is %.0f\n", futureInvestmentValue)
	fmt.Printf("Real accumulated value is %.0f\n", realFutureInvestmentValue)
}

func getUserInput(text string) float64 {
	var input float64
	fmt.Print(text)
	fmt.Scanf("%f", &input)
	return input
}

func calculateInvestment(investmentAmount float64, annualInterestRate float64, years float64) float64 {
	return investmentAmount * math.Pow(1+annualInterestRate/100, years)
}
