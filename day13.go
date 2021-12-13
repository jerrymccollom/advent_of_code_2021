package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

type point struct {
	x, y int
}

func printSheet(sheet [][]int, countOnly bool) {
	count := 0
	for row := 0; row < len(sheet); row++ {
		for col := 0; col < len(sheet[row]); col++ {
			if sheet[row][col] > 0 {
				if !countOnly {
					fmt.Print("#")
				}
				count++
			} else {
				if !countOnly {
					fmt.Print(".")
				}
			}
		}
		if !countOnly {
			fmt.Println()
		}
	}
	if countOnly {
		fmt.Println(count)
	}
	fmt.Println()
}

func foldPaper(data []string, showFirst bool) {
	maxX := 0
	maxY := 0
	folds := 0
	points := make([]*point, 0)
	for i, l := range data {
		p := new(point)
		pt := strings.Split(l, ",")
		if len(pt) < 2 {
			folds = i
			break
		}
		p.x, _ = strconv.Atoi(pt[0])
		if p.x >= maxX {
			maxX = p.x
		}
		p.y, _ = strconv.Atoi(pt[1])
		if p.y >= maxY {
			maxY = p.y
		}
		points = append(points, p)
	}
	maxX++
	maxY++
	sheet := make([][]int, maxY)
	for i := 0; i < maxY; i++ {
		sheet[i] = make([]int, maxX)
	}
	for _, p := range points {
		sheet[p.y][p.x] = 1
	}

	for f := folds; f < len(data); f++ {
		f := strings.Fields(data[f])
		fl := strings.Split(f[2], "=")
		v, _ := strconv.Atoi(fl[1])
		v++
		if fl[0] == "y" {
			for i, r := v, v-2; i < maxY; i++ {
				for c := 0; c < maxX; c++ {
					sheet[r][c] |= sheet[i][c]
				}
				r--
			}
			sheet = sheet[:v-1]
			maxY = v - 1
		} else {
			for i, c := v, v-2; i < maxX; i++ {
				for r := 0; r < maxY; r++ {
					sheet[r][c] |= sheet[r][i]
				}
				c--
			}
			for r := 0; r < maxY; r++ {
				sheet[r] = sheet[r][:v-1]
			}
			maxX = v - 1
		}
		if showFirst {
			printSheet(sheet, true)
			return
		}
	}
	printSheet(sheet, false)
}

func day13part1(data []string) {
	foldPaper(data, true)
}

func day13part2(data []string) {
	foldPaper(data, false)
}

func main() {
	data, _ := readLinesFromFile("day13.txt")
	day13part1(data)
	fmt.Println()
	day13part2(data)
}
