package main // tells the compiler to compile as an executable file

import (
	"fmt"  // import fmt for Printf
	"math" // import math for math.Floor
)

func majority(arr []int) { // arr []int means this takes an array of integers
	fmt.Printf("Given array: %v", arr)
	count := map[int]int{}                          // creates an empty map where the key and value are both integers
	goal := math.Floor(float64((len(arr) / 2) + 1)) // the goal number of appearances converted to float64 so it will be accepted by math.Floor
	fmt.Printf("\nGoal: %v appearances", goal)

	for i := range arr { // loop through the array
		if count[arr[i]] != 0 { // if the key arr[i] is in count, add 1 to its value
			count[arr[i]] += 1
		} else { // else add it and set the value to 1
			count[arr[i]] = 1
		}
	}

	for num, freq := range count { // loop through the map, accessing the key and value
		if freq >= int(goal) {
			fmt.Printf("\n%v is the majority element, appearing %v times\n", num, freq)
		}
	}

}

func main() {
	majority([]int{3, 2, 3})
	majority([]int{2, 2, 1, 1, 1, 2, 2})
}
