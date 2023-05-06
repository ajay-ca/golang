package main

import "fmt"

func main (){
	//declaring numbers as a and b
	var a,b int = 9,16
	
	/*Another method of multiple declaration
	var (
		a int = 9
		b int = 16
		c int = a + b
	)*/

	var c int = a + b
	fmt.Println(c)
}