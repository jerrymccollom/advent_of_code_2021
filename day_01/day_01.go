package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func countIncreases(nums []int) int {
	prev := -1
	count := 0
	for _, v := range nums {
		if prev >= 0 && v > prev {
			count++
		}
		prev = v
	}
	return count
}

func main() {
	nums, err := readFile("numbers.txt")
	if err != nil {
		panic(err)
	}
	fmt.Printf("There are %d measurements larger than the previous measurement.",
		countIncreases(nums))
}
