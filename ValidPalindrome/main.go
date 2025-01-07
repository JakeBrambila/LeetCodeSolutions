package main

import (
	"fmt"
	"unicode"
)

func main() {
	s1 := "amanaplanacanalpanama"
	s2 := "raceacar"
	s3 := "A man, a plan, a canal: Panama"

	if isPalindrome(s1) {
		fmt.Println("this is a palindrome: ", s1)
	} else {
		fmt.Println("This is not a palindrome: ", s1)
	}

	fmt.Println()

	if isPalindrome(s2) {
		fmt.Println("this is a palindrome: ", s2)
	} else {
		fmt.Println("This is not a palindrome: ", s2)
	}

	if isPalindrome(s3) {
		fmt.Println("this is a palindrome: ", s3)
	} else {
		fmt.Println("This is not a palindrome: ", s3)
	}

}

// O(n) uses only one loop
// uses the two pointer method
func isPalindrome(s string) bool {
	runeArray := []rune(s)
	start := 0
	end := len(runeArray) - 1

	for start < end {
		for !isAlphaNumeric(runeArray[start]) && start < end {
			start++
		}
		for !isAlphaNumeric(runeArray[end]) && end > start {
			end--
		}
		if unicode.ToLower(runeArray[start]) != unicode.ToLower(runeArray[end]) {
			return false
		}
		start++
		end--
	}
	return true
}

//helper function that uses Ascii values to determine whethere a rune is alphanumerical
func isAlphaNumeric(r rune) bool {
	if (int('A') <= int(r) && int(r) <= int('Z')) || (int('a') <= int(r) && int(r) <= int('z')) || (int('0') <= int(r) && int(r) <= int('9')) {
		return true
	}
	return false
}
