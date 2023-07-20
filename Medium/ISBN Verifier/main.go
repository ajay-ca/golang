package main
import (
    "fmt"
    "unicode"
    )

func main() {
    var sum, count, numlen int = 0, 10, 0
    var isbn string
    fmt.Println("Enter your ISBN number")
    _, err := fmt.Scan(&isbn)
    if err != nil {
        panic(err)
    }
    var length int = len(isbn)
    for _, letter := range isbn {
        if unicode.IsDigit(letter) {
            sum += (int(letter - '0') * count)
            numlen ++
            count --
        } else if letter == 'X' {
            sum += 10
            numlen ++
        } else if int(letter) >= 65 && int(letter) <= 122 {
        	numlen = 0
        }
    }
    if sum != 0 && numlen == 10 && sum%11 == 0 && length >= 10 {
    	fmt.Println("The given input is an ISBN number")
    } else {
        fmt.Println("The given input is not an ISBN number")
    }
}