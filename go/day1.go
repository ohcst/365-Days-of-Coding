package main

import (
	"fmt"
	"math"
)

func majority(arr []int) {
	fmt.Printf("Given array: %v", arr)
	count := map[int]int{}
	goal := math.Floor(float64((len(arr) / 2) + 1))
	fmt.Printf("\nGoal: %v appearances", goal)

	for i := range arr {
		if count[arr[i]] != 0 {
			count[arr[i]] += 1
		} else {
			count[arr[i]] = 1
		}
	}

	for num, freq := range count {
		if freq >= int(goal) {
			fmt.Printf("\n%v is the majority element, appearing %v times\n", num, freq)
		}
	}

}

func main() {
	majority([]int{3, 2, 3})
	majority([]int{2, 2, 1, 1, 1, 2, 2})
}
