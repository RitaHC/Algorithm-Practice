package main

import (
	"fmt"
	"strings"
)

// QUESTION ! - return if the string is a palindrome or not
func palindromeString(str string) {
	// make the string into all lower case
	strLower := strings.ToLower(str)
	// split string
	splitString := strings.Split(strLower, "")
	// calcaulate the Index
	lastIndex := len(strLower) - 1
	// reverse string
	var reverseString string
	for lastIndex >= 0 {
		reverseString = reverseString + splitString[lastIndex]
		lastIndex--
	}
	// Now check for Palindrome Condition
	if reverseString == str {
		fmt.Println("String is a PALINDROME : ", reverseString)
	} else {
		fmt.Println("Not Palindrome : ", str, " - ", reverseString)
	}
}

// QUESTION 2 - return all pal;indorome's in the slice
// Input: words = ["abc","car","ada","racecar","cool"]
// Output: ["ada","racecar"]

func palindrone(word []string) []string {
	// create a
}

func main() {
	fmt.Println("------ Palindrome STRING -------")
	palindromeString("Hello")
	palindromeString("Samas")
	fmt.Println("------ Palindrome SLICE -------")

}
