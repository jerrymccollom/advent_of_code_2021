package main

import (
	"fmt"
	"strings"
)

func main() {
	// Represent the initial fish sensors as an array of integers
	// fs := strings.Split("3,4,3,1,2", ",")  // initial example
	fs := strings.Split("4,2,4,1,5,1,2,2,4,1,1,2,2,2,4,4,1,2,1,1,4,1,2,1,2,2,2,2,5,2,2,3,1,4,4,4,1,2,3,4,4,5,4,3,5,1,2,5,1,1,5,5,1,4,4,5,1,3,1,4,5,5,5,4,1,2,3,4,2,1,2,1,2,2,1,5,5,1,1,1,1,5,2,2,2,4,2,4,2,4,2,1,2,1,2,4,2,4,1,3,5,5,2,4,4,2,2,2,2,3,3,2,1,1,1,1,4,3,2,5,4,3,5,3,1,5,5,2,4,1,1,2,1,3,5,1,5,3,1,3,1,4,5,1,1,3,2,1,1,1,5,2,1,2,4,2,3,3,2,3,5,1,5,1,2,1,5,2,4,1,2,4,4,1,5,1,1,5,2,2,5,5,3,1,2,2,1,1,4,1,5,4,5,5,2,2,1,1,2,5,4,3,2,2,5,4,2,5,4,4,2,3,1,1,1,5,5,4,5,3,2,5,3,4,5,1,4,1,1,3,4,4,1,1,5,1,4,1,2,1,4,1,1,3,1,5,2,5,1,5,2,5,2,5,4,1,1,4,4,2,3,1,5,2,5,1,5,2,1,1,1,2,1,1,1,4,4,5,4,4,1,4,2,2,2,5,3,2,4,4,5,5,1,1,1,1,3,1,2,1", ",")
	fc := make([]int, 9)
	for _, fv := range fs {
		fc[(int)(fv[0]-'0')]++
	}

	// Grouping all fish by sensor value, compute next state of each group
	// Part 1 compuptes these values over 80 days
	for day := 0; day < 256; day++ {
		d := fc[0]
		for i := 0; i < 8; i++ {
			fc[i] = fc[i+1]
		}
		fc[6] += d
		fc[8] = d
	}

	// Compute the sum of all fish-state counts, which is the total number of extant fish
	sum := 0
	for i := 0; i < 9; i++ {
		sum += fc[i]
	}
	fmt.Println(sum)
}
