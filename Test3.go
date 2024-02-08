package main

// Question : Beautiful Number
// Example 1: if number is 110011 then the maimum number that needs to be changed is 2 i.e. two 0's
// Example 2: 110000 -> maximum number of changes to be made 2 i.e. 2 1's

import (
	"fmt"
	"strconv"
)

func beautifulNumber(x int) int {
	fmt.Println("Digit: ", x)
	// ones and zeros as integers
	var ones, zeros int
	// convert integer to string
	str := strconv.Itoa(x)
	// Iterate over the string
	for _, digit := range str {
		// add digits to either their respective
		if digit == '1' {
			ones++
		} else if digit == '0' {
			zeros++
		}
	}
	// check how many 1's and how many 0's
	// return which ever number is lower or dependeing upon the condition
	if ones > zeros {
		return zeros
	} else if zeros > ones {
		return ones
	} else if zeros == ones {
		return zeros
	} else {
		return 0
	}

}

func main() {
	fmt.Println("---------")
	x := 111100
	fmt.Println("To beautify change the following number of digits: ", beautifulNumber(x))
	fmt.Println("---------")
	a := 1001
	fmt.Println("To beautify change the following number of digits: ", beautifulNumber(a))
	fmt.Println("---------")
	b := 10
	fmt.Println("To beautify change the following number of digits: ", beautifulNumber(b))
	fmt.Println("---------")
	c := 1111
	fmt.Println("To beautify change the following number of digits: ", beautifulNumber(c))
	fmt.Println("---------")
	d := 0000
	fmt.Println("To beautify change the following number of digits: ", beautifulNumber(d))

}
