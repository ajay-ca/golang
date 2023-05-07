package main

import "fmt"

func main() {
	fmt.Println("Enter Your Year: ")
	var year int
	fmt.Scanln(&year)

	//Condition to check the leap year
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("This is a leap year")
	} else {
		fmt.Println("This is not a leap year")
	}
	

	/*Alternate method for condition (Nested if-loop)
	if year%4 == 0 {
		if year%400 == 0 {
			if year%100 != 0 {
				fmt.Println("This is a leap year")
			}
		}
	} else {
		fmt.Println("This is not a leap year")
	}*/
}