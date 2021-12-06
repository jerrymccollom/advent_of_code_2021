package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const BOARDSIZE = 1000
const DEBUG = 0

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

func getPoints(line string) (x1, y1, x2, y2 int) {
	p := strings.Split(line, " -> ")
	p1 := strings.Split(p[0], ",")
	p2 := strings.Split(p[1], ",")
	x1, _ = strconv.Atoi(p1[0])
	y1, _ = strconv.Atoi(p1[1])
	x2, _ = strconv.Atoi(p2[0])
	y2, _ = strconv.Atoi(p2[1])
	return x1, y1, x2, y2
}

var board [BOARDSIZE][BOARDSIZE]int

func updateBoard(data []string) int {
	clearBoard()
	for _, line := range data {
		x1, y1, x2, y2 := getPoints(line)
		if DEBUG != 0 {
			fmt.Printf("%d, %d -> %d, %d\n", x1, y1, x2, y2)
		}
		if x1 == x2 {
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				board[i][x1]++
			}
		} else if y1 == y2 {
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				board[y1][i]++
			}
		} else {
			xi, yi := 0, 0
			if x1 > x2 {
				xi = -1
			}
			if x1 < x2 {
				xi = 1
			}
			if y1 > y2 {
				yi = -1
			}
			if y1 < y2 {
				yi = 1
			}

			y := y1
			x := x1
			for ; x >= 0 && x < BOARDSIZE && x != x2; x, y = x+xi, y+yi {
				board[y][x]++
			}
			board[y][x]++
		}
		printBoard()
	}

	return addOverlaps()

}

func addOverlaps() int {
	count := 0
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			if board[j][i] > 1 {
				count++
			}
		}
	}
	return count
}

func printBoard() {
	if DEBUG == 0 {
		return
	}
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			if board[i][j] > 0 {
				fmt.Printf("%d ", board[i][j])
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func clearBoard() {
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			board[i][j] = 0
		}
	}
}

func main() {
	data, err := readLinesFromFile("day5.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(updateBoard(data))
}
