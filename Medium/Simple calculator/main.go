package main

import "fmt"

func main() {
	var a,b,option int
	fmt.Print("Enter the two numbers:")
	fmt.Print("a = ")
	fmt.Scanln("%d",&a)
	fmt.Print("b = ")
	fmt.Scanln("%d",&b)
	fmt.Print("Select the operation:\n 1. Addition\n2.Subtraction\n3.Multiplication\n4.Division\n5.Exponent\nEnter your option: ")
	fmt.Scanln("%d",&option)
	

}