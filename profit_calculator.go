package main

import "fmt"
import "errors"

func main() {
	var ebt float64
	var profit float64
	var ratio float64

	revenue, err1 := getUserInput("Enter revenue: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	expenses, err2 := getUserInput("Enter expenses: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	taxRate, err3 := getUserInput("Enter tax rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

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

func getUserInput(infoText string) (float64, error) {
	var  userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}