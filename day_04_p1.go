package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Tile struct {
	Value  int
	Marked bool
}
type Board [5][5]Tile

type Draws []int

func getDraws(data []string) Draws {
	values := strings.Split(data[0], ",")
	result := make(Draws, 0, len(values))
	for _, v := range values {
		if len(v) > 0 {
			i, _ := strconv.Atoi(v)
			result = append(result, i)
		}
	}
	return result
}

func getBoards(data []string) []Board {
	count := 0
	i := 0
	boards := make([]Board, 0)
	for ; i < len(data)-1; count++ {
		var b Board
		for j := 1; j < 6; j++ {
			vs := strings.Fields(data[i+j])
			for ki, ks := range vs {
				k, _ := strconv.Atoi(ks)
				b[j-1][ki].Value = k
				b[j-1][ki].Marked = false
			}
		}
		boards = append(boards, b)
		i += 5
	}
	return boards
}

func hasBingo(b Board) bool {
	for r := 0; r < 5; r++ {
		if b[r][0].Marked &&
			b[r][1].Marked &&
			b[r][2].Marked &&
			b[r][3].Marked &&
			b[r][4].Marked {
			return true
		}
	}
	for c := 0; c < 5; c++ {
		if b[0][c].Marked &&
			b[1][c].Marked &&
			b[2][c].Marked &&
			b[3][c].Marked &&
			b[4][c].Marked {
			return true
		}
	}
	return false
}

func markTiles(v int, b *Board) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if (*b)[r][c].Value == v {
				(*b)[r][c].Marked = true
			}
		}
	}
}

func playBingo(d Draws, b []Board) (Board, int) {
	for _, v := range d {
		for i := range b {
			markTiles(v, &b[i])
			bingo := hasBingo(b[i])
			if bingo {
				return b[i], v
			}
		}
	}
	return Board{}, 0
}

func playSquidBingo(d Draws, b []Board) (Board, int) {
	winners := make([]bool, len(b))
	lastWinner := -1
	lastWinnerValue := -1
	for _, v := range d {
		for i := range b {
			if !winners[i] {
				markTiles(v, &b[i])
				bingo := hasBingo(b[i])
				if bingo {
					winners[i] = true
					lastWinner = i
					lastWinnerValue = v
				}
			}
		}
	}
	return b[lastWinner], lastWinnerValue
}

func winnerValue(b Board, v int) int {
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !b[r][c].Marked {
				sum += b[r][c].Value
			}
		}
	}
	return sum * v
}

func main() {
	data, err := readLinesFromFile("day4.txt")
	if err != nil {
		panic(err)
	}
	d := getDraws(data)
	b := getBoards(data)
	winner, v := playBingo(d, b)
	fmt.Println(winnerValue(winner, v))
	secondWinner, v2 := playSquidBingo(d, b)
	fmt.Println(winnerValue(secondWinner, v2))
}
