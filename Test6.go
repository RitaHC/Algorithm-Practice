package main

// You are given an integer array heights representing the heights of buildings, some bricks, and some ladders.

// You start your journey from building 0 and move to the next building by possibly using bricks or ladders.

// While moving from building i to building i+1 (0-indexed),

// If the current building's height is greater than or equal to the next building's height, you do not need a ladder or bricks.
// If the current building's height is less than the next building's height, you can either use one ladder or (h[i+1] - h[i]) bricks.
// Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.

// Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
// Output: 4
// Explanation: Starting at building 0, you can follow these steps:
// - Go to building 1 without using ladders nor bricks since 4 >= 2.
// - Go to building 2 using 5 bricks. You must use either bricks or ladders because 2 < 7.
// - Go to building 3 without using ladders nor bricks since 7 >= 6.
// - Go to building 4 using your only ladder. You must use either bricks or ladders because 6 < 9.
// It is impossible to go beyond building 4 because you do not have any more bricks or ladders.

import (
	"fmt"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	var buildingsCrossed int
	// Check if the height of the next building is higher or lower than the current building
	for i := 0; i < len(heights)-1; i++ {
		// calculate height difference
		heightDifference := heights[i+1] - heights[i]
		// if the current height is more than the next height
		if heights[i] > heights[i+1] {
			// if the height of the current building is higher, simply add to the buildings crossed
			buildingsCrossed++
			// now use the ladder
		} else if heights[i] < heights[i+1] {
			// now check if the number of bricks or ladders is >= heightDifference
			// then use those bricks or ladders
			if bricks >= heightDifference {
				bricks = bricks - heightDifference
				buildingsCrossed++

			} else if ladders > 0 {
				ladders--
				buildingsCrossed++

			} else if ladders == 0 && bricks < heightDifference && heights[i] < heights[i+1] {
				break
			}
		}
	}

	fmt.Println("Input values at the end", heights, bricks, ladders)
	fmt.Println("======")
	// reduce the 1st building from being counted
	// buildingsCrossed = buildingsCrossed - 1
	return buildingsCrossed
}

func main() {
	fmt.Println("--------CASE 1----------")
	building1 := []int{4, 2, 7, 6, 9, 14, 12}
	bricks1 := 5
	ladder1 := 1
	fmt.Println(furthestBuilding(building1, bricks1, ladder1))
	fmt.Println("--------CASE 2----------")
	building2 := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	bricks2 := 10
	ladder2 := 2
	fmt.Println(furthestBuilding(building2, bricks2, ladder2))
	fmt.Println("---------CASE 3---------")
	building3 := []int{14, 3, 19, 3}
	bricks3 := 17
	ladder3 := 0
	fmt.Println(furthestBuilding(building3, bricks3, ladder3))

}
