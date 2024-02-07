package main

// Create a program which calculates the number of times 1's or 0's occure in a binary integer

import (
	"fmt"
	"strconv"
)

func calculateOccurance(x int) (int, int) {

	// Initialize the return values
	var y, z int
	// First convert int into string
	str := strconv.Itoa(x)
	// Iterate over the string in go
	for _, digit := range str {
		// store the number of times a number is being repeated in a variable
		if digit == '1' {
			y++
		} else if digit == '0' {
			z++
		}
	}
	// return the value's

	fmt.Println("1's :", y)
	fmt.Println("1's :", z)
	return y, z
}

func main() {
	fmt.Println("--------")
	/////// call the function //////
	a := 1001
	b := 10
	c := 11
	calculateOccurance(a)
	fmt.Println("--------")
	calculateOccurance(b)
	fmt.Println("--------")
	calculateOccurance(c)

}
