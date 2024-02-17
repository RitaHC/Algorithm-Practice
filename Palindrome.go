package palindrome

import (
	"fmt"
	"strings"
)

// QUESTION 1 - return if the string is a palindrome or not
func palindromeString(str string) bool {
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
		return true
	} else {
		fmt.Println("Not Palindrome : ", str, " - ", reverseString)
		return false
	}
}

// QUESTION 2 - return all pal;indorome's in the slice
// Input: words = ["abc","car","ada","racecar","cool"]
// Output: ["ada","racecar"]

func palindroneSlice(words []string) []string {
	// create a slice and return it
	var palindromeSlice []string
	// loop over each string in the slice
	for _, word := range words {
		// Now check if the word is a palindrome using the above function
		// if the return value is True, then append to new slice
		isPalindrome := palindromeString(word)

		if isPalindrome {
			// if palindromeString returns true then append it to the new slice
			palindromeSlice = append(palindromeSlice, word)
		}
	}
	// return the slice
	fmt.Println("==== The List below shows all the Palindrome words in the list ======")
	return palindromeSlice
}

func main() {
	fmt.Println("------ Palindrome STRING -------")
	palindromeString("Hello")
	palindromeString("Samas")
	fmt.Println("------ Palindrome SLICE -------")
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(palindroneSlice(words))

}
