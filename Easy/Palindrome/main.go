package main

import "fmt"

func main (){

	var num,temp,revnum int
	fmt.Println("Enter your number:")
	fmt.Scanln(&num)

	temp = num
	revnum = 0
	for temp > 0 {
		revnum*=10			//increasing the digits place
		revnum+=temp%10		//adding the last digit to reverse
		temp/=10			//removing the last digit
	}
	if num == revnum {
		fmt.Println("The number is a palindrome")
	} else {
		fmt.Println("The number is not a palindrome")
	}

}