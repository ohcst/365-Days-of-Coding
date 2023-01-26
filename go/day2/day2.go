package main

import (
	"fmt"
	"math"
)

func collatz(num int) {
	var arr []int
	arr = append(arr, num)

	for {
		if num == 1 {
			break
		}
		if num%2 == 0 {
			num = int(math.Floor(float64(num / 2)))
		} else {
			num = (num * 3) + 1
		}
		arr = append(arr, num)
	}

	fmt.Println(arr)
}

func main() {
	collatz(11)
}
