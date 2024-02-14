package main

import (
	"fmt"
)

// QUESTION 1 : Ugly unmbers are the positive integers whose only prime factors are 2,3 or 5
func uglyNumber() []int {
	// create a return slice
	num := 2
	uglyNumbers := []int{1}
	// To print the first 10 ugly numbers, loop over untill the len(uglyNumbers) > 10
	for len(uglyNumbers) < 10 {
		if num%2 == 0 || num%3 == 0 || num%5 == 0 {
			uglyNumbers = append(uglyNumbers, num)
		}
		num++
	}
	// return the slice
	return uglyNumbers
}

func main() {
	fmt.Println("------------")
	fmt.Println(uglyNumber())

}
