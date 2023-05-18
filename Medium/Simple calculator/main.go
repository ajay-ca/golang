package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, option, result float64
	fmt.Print("Enter the two numbers:\na = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("Select the operation:\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division\n5. Exponent\nEnter your option: ")
calculation:
	fmt.Scan(&option)
	switch option {
	case 1:
		result = a + b
	case 2:
		result = a - b
	case 3:
		result = a * b
	case 4:
		result = a / b
	case 5:
		result = math.Pow(a, b)
	default:
		fmt.Println("Invalid option. Try again: ")
		goto calculation
	}
	fmt.Printf("The result is %.2f\n", result)
}