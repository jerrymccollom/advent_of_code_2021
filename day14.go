package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func readLinesFromFile(fname string) (result []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	// filter out empty lines
	result = make([]string, 0, len(lines))
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		result = append(result, l)
	}
	return result, nil
}

func computerPolymerPairs(data []string, numSteps int) int {
	template := data[0]
	decoder := make(map[string]string)
	for i := 1; i < len(data); i++ {
		p := strings.Split(data[i], " -> ")
		if len(p) == 2 {
			decoder[p[0]] = p[1]
		}
	}
	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		pairs[template[i:i+2]]++
	}

	for i := 1; i <= numSteps; i++ {
		newPairs := make(map[string]int)
		for pair := range pairs {
			newPairs[string(pair[0])+decoder[pair]] += pairs[pair]
			newPairs[decoder[pair]+string(pair[1])] += pairs[pair]
		}
		pairs = newPairs
	}

	countMap := make(map[string]int)
	for i := range pairs {
		countMap[string(i[0])] += pairs[i]
	}
	countMap[string(template[len(template)-1])] += 1

	minCount := math.MaxInt
	maxCount := 0
	for _, v := range countMap {
		if v > maxCount {
			maxCount = v
		}
		if v < minCount {
			minCount = v
		}
	}

	return maxCount - minCount
}

func day14part1(data []string) {
	fmt.Println(computerPolymerPairs(data, 10))
}

func day14part2(data []string) {
	fmt.Println(computerPolymerPairs(data, 40))
}

func main() {
	data, _ := readLinesFromFile("day14.txt")
	day14part1(data)
	fmt.Println()
	day14part2(data)
}
