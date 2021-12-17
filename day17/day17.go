package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

type Area struct {
	minX, maxX, minY, maxY int
}

func inTargetArea(target Area, x, y int) bool {
	return x >= target.minX &&
		x <= target.maxX &&
		y >= target.minY &&
		y <= target.maxY
}

func getTargetArea(s string) Area {
	fields := strings.Split(s, " ")
	x := strings.Split(fields[2], "..")
	x[0] = strings.Trim(x[0], "x=,")
	x[1] = strings.Trim(x[1], "x=,")
	y := strings.Split(fields[3], "..")
	y[0] = strings.Trim(y[0], "y=,")
	y[1] = strings.Trim(y[1], "y=,")

	result := Area{}
	result.minX, _ = strconv.Atoi(x[0])
	result.maxX, _ = strconv.Atoi(x[1])
	result.minY, _ = strconv.Atoi(y[0])
	result.maxY, _ = strconv.Atoi(y[1])

	return result
}

const NOTREACHED = -987654

func maxYPosition(target Area, xVol, yVol int, debug bool) int {
	x, y := 0, 0
	step := 0
	initX, initY := xVol, yVol
	maxY := 0
	for {
		step++
		if step > 300 {
			return NOTREACHED
		}
		x += xVol
		y += yVol
		if y > maxY {
			maxY = y
		}
		if xVol > 0 {
			xVol--
		} else if xVol < 0 {
			xVol++
		}
		yVol--
		if inTargetArea(target, x, y) {
			if debug {
				fmt.Println(initX, ",", initY, " got to ", maxY)
			}
			return maxY
		}
	}
}

func day17part1(data []string) {
	target := getTargetArea(data[0])
	maxY := 0
	bestYV := -1
	bestXV := -1
	for xVol := 1; xVol < 100; xVol++ {
		for yVol := 1; yVol < 100; yVol++ {
			y := maxYPosition(target, xVol, yVol, false)
			if y != NOTREACHED && y > maxY {
				bestXV = xVol
				bestYV = yVol
				maxY = y
			}
		}
	}
	fmt.Println("MAXY:", maxY, "BEST VELOCITY:", bestXV, ",", bestYV)
}

func day17part2(data []string) {
	target := getTargetArea(data[0])
	maxY := 0
	bestYV := -1
	bestXV := -1
	count := 0
	for xVol := -1000; xVol < 1000; xVol++ {
		for yVol := -1000; yVol < 1000; yVol++ {
			y := maxYPosition(target, xVol, yVol, false)
			if y != NOTREACHED && y > maxY {
				bestXV = xVol
				bestYV = yVol
				maxY = y
			}
			if y != NOTREACHED {
				count++
			}
		}
	}
	fmt.Println("MAXY:", maxY, "BEST VELOCITY:", bestXV, ",", bestYV)
	fmt.Println(count)
}

func main() {
	data, _ := readLinesFromFile("day17.txt")
	day17part1(data)
	fmt.Println()
	day17part2(data)
}
