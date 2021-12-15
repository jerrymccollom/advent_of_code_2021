package main

import (
	"fmt"
	"github.com/RyanCarrier/dijkstra"
	"io/ioutil"
	"strings"
)

const DEBUG = false

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

func vertexID(r, c int) int {
	return r*maxRow + c
}

func graphFromChart(chart [][]int) *dijkstra.Graph {
	g := dijkstra.NewGraph()
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			nodeID := vertexID(row, col)
			g.AddVertex(nodeID)
		}
	}
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			nodeID := vertexID(row, col)
			rightID := vertexID(row, col+1)
			leftID := vertexID(row, col-1)
			downID := vertexID(row+1, col)
			upID := vertexID(row-1, col)
			if row < maxRow-1 {
				g.AddArc(nodeID, downID, int64(chart[row+1][col]))
			}
			if row > 0 {
				g.AddArc(nodeID, upID, int64(chart[row-1][col]))
			}
			if col < maxCol-1 {
				g.AddArc(nodeID, rightID, int64(chart[row][col+1]))
			}
			if col > 0 {
				g.AddArc(nodeID, leftID, int64(chart[row][col-1]))
			}
		}
	}
	return g
}

func day15part1(data []string) {
	chart := getChart(data)
	g := graphFromChart(chart)
	p, _ := g.Shortest(vertexID(0, 0), vertexID(maxRow-1, maxCol-1))
	if DEBUG {
		printChart(chart, p)
	} else {
		fmt.Println(p.Distance)
	}
}

func nextLevel(risk int) int {
	if risk == 9 {
		return 1
	}
	return risk + 1
}

func newRisk(risk, count int) int {
	for i := 0; i < count; i++ {
		risk = nextLevel(risk)
	}
	return risk
}

func multiplyChart(chart [][]int, times int) [][]int {
	result := make([][]int, maxRow*times)
	for rowCopy := 0; rowCopy < times; rowCopy++ {
		for r := 0; r < maxRow; r++ {
			result[rowCopy*maxRow+r] = make([]int, maxCol*times)
			for colCopy := 0; colCopy < times; colCopy++ {
				for c := 0; c < maxCol; c++ {
					result[maxRow*rowCopy+r][maxCol*colCopy+c] = newRisk(chart[r][c], rowCopy+colCopy)
				}
			}
		}
	}
	maxRow *= times
	maxCol *= times
	return result
}

func printChart(chart [][]int, path dijkstra.BestPath) {

	c := 0
	visit := make(map[int]bool)
	for _, p := range path.Path {
		visit[p] = true
	}
	for row := 0; row < maxRow; row++ {
		for col := 0; col < maxCol; col++ {
			if visit[vertexID(row, col)] {
				c += chart[row][col]
				fmt.Print(colorGreen)
			}
			fmt.Print(chart[row][col])
			if visit[vertexID(row, col)] {
				fmt.Print(colorReset)
			}
		}
		fmt.Println()
	}
	fmt.Println(path.Distance)
}

func day15part2(data []string) {
	chart := getChart(data)
	chart = multiplyChart(chart, 5)
	g := graphFromChart(chart)
	p, _ := g.Shortest(vertexID(0, 0), vertexID(maxRow-1, maxCol-1))
	if DEBUG {
		printChart(chart, p)
	} else {
		fmt.Println(p.Distance)
	}
}

func main() {
	data, _ := readLinesFromFile("day15.txt")
	day15part1(data)
	fmt.Println()
	day15part2(data)
}
