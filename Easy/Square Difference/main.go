package main

import (
	"fmt"
)

func main() {
// Get the value of N from the user
	var n int
	fmt.Print("Enter the value of N: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

// Calculate the difference
	difference := SquareOfSum(n) - SumOfSquares(n)

// Print the result
	fmt.Printf("The difference between the square of the sum and the sum of the squares of the first %d natural numbers is %d.\n", n, difference)
}

// Function to find square of the sum
func SquareOfSum(n int) int {
	var sqsum int
    for i := 1; i <= n; i++ {
		sqsum+=i    
    }
	return sqsum*sqsum
}

// Function to find sum of the square
func SumOfSquares(n int) int {
	var sumsq int
    for i := 1; i <= n; i++ {
		sumsq+=i*i    
    }
    return sumsq
}
