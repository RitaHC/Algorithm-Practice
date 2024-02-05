// Nested Maps //
package main

import (
	"fmt"
)

func getNamesCounts(names []string) map[rune]map[string]int {
	// create a return map
	counts := make(map[rune]map[string]int)
	// loop over the names
	for _, name := range names {
		// Check if we already have a map associated with the first letter
		// if len(name) == "" {
		// 	continue
		// }
		// Extract the 1st character in the name
		firstChar := (rune(name[0]))
		fmt.Println("First character : ", firstChar) // Convert rune to string for display
		// Now look-up in our map if the firstChar already exists
		_, ok := counts[firstChar]
		// If it does not exist
		if !ok {
			// initialize it
			counts[firstChar] = make(map[string]int)
		}
		// if it already exists
		counts[firstChar][name]++
	}
	return counts
}

func main() {
	fmt.Println("-------")
	names := []string{"billy", "billy", "bob", "joe"}
	fmt.Println(getNamesCounts(names))
}
