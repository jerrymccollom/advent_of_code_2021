package main

import (
	"fmt"
	"io/ioutil"
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

const colorReset = "\033[0m"
const colorGreen = "\033[32m"

var chart [][]int
var flashMap [][]bool
var maxCol, maxRow int

func getChart(data []string) [][]int {
	maxCol = len(data[0])
	maxRow = len(data)
	result := make([][]int, maxRow)
	for row := 0; row < maxRow; row++ {
		result[row] = make([]int, maxCol)
		for col := 0; col < maxCol; col++ {
			result[row][col] = int(data[row][col]) - '0'
		}
	}
	return result
}

func printChart() {
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			if chart[row][col] == 0 {
				fmt.Print(colorGreen)
			}
			fmt.Print(chart[row][col])
			if chart[row][col] == 0 {
				fmt.Print(colorReset)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func clearFlashMap() {
	flashMap = make([][]bool, maxRow)
	for row := 0; row < maxRow; row++ {
		flashMap[row] = make([]bool, maxCol)
	}
}

func propagate(row, col int) int {
	if row < 0 || col < 0 || row >= maxRow || col >= maxCol {
		return 0
	}
	if chart[row][col] > 9 {
		return flashOctopus(row, col)
	}
	if !flashMap[row][col] {
		chart[row][col]++
	}
	return 0
}

func flashOctopus(row, col int) int {
	count := 0
	if !flashMap[row][col] {
		flashMap[row][col] = true
		chart[row][col] = 0
		count++
		count += propagate(row-1, col)
		count += propagate(row+1, col)
		count += propagate(row, col-1)
		count += propagate(row, col+1)
		count += propagate(row-1, col-1)
		count += propagate(row+1, col+1)
		count += propagate(row-1, col+1)
		count += propagate(row+1, col-1)
	}
	return count
}

func computeFlashes() int {
	clearFlashMap()
	result := 0
	for done := false; !done; {
		cur := 0
		for row := 0; row < maxRow; row++ {
			for col := 0; col < maxCol; col++ {
				if chart[row][col] > 9 {
					cur += flashOctopus(row, col)
				}
			}
		}
		if cur == 0 {
			done = true
		}
		result += cur
	}
	return result
}

func doStep(step int, printCharts bool) int {
	// increment energies this step
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			chart[row][col]++
		}
	}
	result := computeFlashes()
	if printCharts && (step <= 10 || step <= 100 && step%10 == 0) {
		fmt.Println("After step ", step, ":")
		printChart()
	}

	return result
}

func day11part1(data []string) {
	chart = getChart(data)
	printChart()
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += doStep(i, true)
	}
	fmt.Println("There were", sum, "flashes after 100 steps")
}

func day11part2(data []string) {
	chart = getChart(data)
	step := 1
	for ; doStep(step, false) != maxRow*maxCol; step++ {
	}

	fmt.Println("All flashed on step", step)
}

func main() {
	data, _ := readLinesFromFile("day11.txt")
	day11part1(data)
	fmt.Println()
	day11part2(data)
}
