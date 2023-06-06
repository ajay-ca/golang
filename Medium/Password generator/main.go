package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	LowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits           = "0123456789"
	SpecialChars     = "!@#$%^&*()_+{}[]:;?><.,"
)

//Password generating function
func generator(length int) string {
	rand.Seed(time.Now().UnixNano()) //Initializing the random number generator using the current time as the seed

	var password strings.Builder
	characterSet := LowerCaseLetters + UpperCaseLetters + Digits + SpecialChars

	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characterSet))
		password.WriteByte(characterSet[randomIndex])
	}

	return password.String()


}
//Password strength analysing function
func passStrength(password string, length int) string {
	
	point := 0

	if length >= 24 {
		point += 1
	}

	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, char := range []rune(password) {
		if strings.ContainsRune(LowerCaseLetters, char) && !hasLower {
			point++
			hasLower = true
		} else if strings.ContainsRune(UpperCaseLetters, char) && !hasUpper {
			point++
			hasUpper = true
		} else if strings.ContainsRune(Digits, char) && !hasDigit {
			point++
			hasDigit = true
		} else if strings.ContainsRune(SpecialChars, char) && !hasSpecial {
			point++
			hasSpecial = true
		}
	}
	strength := "-"
	switch point {
	case 1:
		strength = "Very weak"
	case 2:
		strength = "Weak"
	case 3:
		strength = "Good"
	case 4:
		strength = "Strong"
	case 5:
		strength = "Super Strong"
	}
	return strength
}

func main() {
	var length int
	var retry string
	fmt.Print("Enter the desired password length (8-32): ")
	fmt.Scan(&length)

	if length < 8 || length > 32 {
		fmt.Println("Invalid password length. Please enter a length between 8 and 32.")
		return
	}
	//Password generation based on length
	GENERATE:
	password := generator(length)
	strength := passStrength(password, length)

	fmt.Println("\nGenerated password:\t(" + strength + ")")
	fmt.Println(password)

	fmt.Print("\nGenerate another (Y|y): ")
	fmt.Scan(&retry)
	if retry == "Y" || retry == "y" {
		goto GENERATE
	}
}