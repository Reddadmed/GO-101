package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Get user input
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Perform the operation
	switch operator {
	case "+":
		fmt.Printf("Result: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Result: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Result: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Result: %.2f\n", num1/num2)
		} else {
			fmt.Println(num1)
		}
	default:
		fmt.Println("Invalid operator. Please use +, -, *, or /.")
	}
}
