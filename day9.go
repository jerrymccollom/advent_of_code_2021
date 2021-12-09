package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func getMap(data []string) [][]int {
	result := make([][]int, len(data)+2)
	for i := 0; i < len(data)+2; i++ {
		result[i] = make([]int, len(data[0])+2)
		for j := 0; j < len(result[i]); j++ {
			result[i][j] = 10
		}
	}
	for i, l := range data {
		for j := 0; j < len(l); j++ {
			result[i+1][j+1] = int(l[j]) - '0'
		}
	}
	return result
}

func findLows(m [][]int) []int {
	result := make([]int, 0)
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			if m[i][j] < m[i][j-1] &&
				m[i][j] < m[i][j+1] &&
				m[i][j] < m[i-1][j] &&
				m[i][j] < m[i+1][j] {
				result = append(result, m[i][j])
			}
		}
	}
	return result
}

func riskLevel(readings []int) int {
	sum := 0
	for i := 0; i < len(readings); i++ {
		sum += readings[i] + 1
	}
	return sum
}

func day9part1(data []string) {
	fmt.Println("Risk level is", riskLevel(findLows(getMap(data))))
}

type Basin map[int]bool

func addToBasin(b *Basin, m [][]int, x, y int) {
	if x < 1 || x > len(m)+1 ||
		y < 1 || y > len(m[0])+1 {
		return
	}
	p := m[x][y]
	if p < 9 {
		(*b)[x*len(m)+y] = true
		if m[x-1][y] > p {
			addToBasin(b, m, x-1, y)
		}
		if m[x+1][y] > p {
			addToBasin(b, m, x+1, y)
		}
		if m[x][y-1] > p {
			addToBasin(b, m, x, y-1)
		}
		if m[x][y+1] > p {
			addToBasin(b, m, x, y+1)
		}
	}
}

func fundBasins(m [][]int) []*Basin {
	result := make([]*Basin, 0)
	for i := 1; i < len(m); i++ {
		for j := 1; j < len(m[i]); j++ {
			if m[i][j] < m[i][j-1] &&
				m[i][j] < m[i][j+1] &&
				m[i][j] < m[i-1][j] &&
				m[i][j] < m[i+1][j] {
				b := new(Basin)
				*b = make(map[int]bool)
				addToBasin(b, m, i, j)
				result = append(result, b)
			}
		}
	}
	return result
}

func day9part2(data []string) {
	basins := fundBasins(getMap(data))
	fmt.Printf("%d basins, top 3 are ", len(basins))
	result := make([]int, len(basins))
	for i, b := range basins {
		result[i] = len(*b)
	}
	sort.Ints(result)
	result = result[len(result)-3:]
	fmt.Println(result, "thus puzzle answer is", result[0]*result[1]*result[2])
}

func main() {
	data, _ := readLinesFromFile("day9.txt")
	day9part1(data)
	fmt.Println()
	day9part2(data)
}
