package main

import (
	"os"
	"fmt"
	"bufio"
	"unicode"
)

func main() {
	var word string
	flag := true // used to check within the iteration
	
	fmt.Println("Enter your word or phrase: ")
	// used for scanning even the phrases with spaces
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() 
    word = scanner.Text()

    for i,letter := range word {
        if unicode.IsLetter(letter) {			// to avoid the special characters and spaces
            for j,check := range word {
                if i != j && unicode.ToLower(letter) == unicode.ToLower(check) {		// to avoid case-sensitivity
                    flag = false
                }
            }
        }
    }
	if flag == true { 
		fmt.Println("The given word or phrase is an isogram")
	} else { 
		fmt.Println("The given word or phrase is not an isogram")
	}

}
