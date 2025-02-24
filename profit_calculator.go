package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	var ebt float64
	var profit float64
	var ratio float64

	revenue = getUserInput("Enter revenue: ")
	expenses = getUserInput("Enter expenses: ")
	taxRate = getUserInput("Enter tax rate: ")

	ebt, profit, ratio = financialCalculator(revenue, expenses, taxRate)

	fmt.Printf("%.2f\n", ebt)
	fmt.Printf("%.2f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

//Passing parameters in a function
// func outputText(text string) {
// 	fmt.Print()
// }

func financialCalculator(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var  userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}