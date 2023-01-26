package main // tells the compiler to compile as an executable file

import (
	"fmt"  // import fmt for Println
	"math" // import math for math.Floor
)

func collatz(num int) {
	var arr []int          // create an empty slice
	arr = append(arr, num) // append the starting number to the slice

	for { // for loop that will run until broken (as replacement for the non-existant while loop)
		if num == 1 { // break as we have gotten to the final number
			break
		}
		if num%2 == 0 { // if the remainder of num/2 is 0 (number is even) we will divide the number by two
			num = int(math.Floor(float64(num / 2)))
		} else { // otherwise we will multiply by 3 and add 1
			num = (num * 3) + 1
		}
		arr = append(arr, num) // append the changed number to the slice
	}

	fmt.Println(arr)
}

func main() {
	collatz(11)
}
