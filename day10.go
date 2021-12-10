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

var stack []rune

func push(c rune) {
	stack = append(stack, c)
}

func pop() rune {
	n := len(stack) - 1
	if n < 0 {
		return 'x'
	}
	r := stack[n]
	stack = stack[:n]
	return r
}

var tokens = map[rune]rune{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')',
}

var points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func parseLine(s string) int {
	stack = []rune{}
	for _, r := range s {
		if r == '(' || r == '[' || r == '<' || r == '{' {
			// fmt.Println("PUSH ", string(r))
			push(tokens[r])
		} else {
			c := pop()
			// fmt.Println("POP ", string(c))
			if c != r {
				fmt.Println("ERROR: EXPECTED ", string(c), " GOT ", string(r))
				return points[r]
			}
		}
	}
	return 0
}

func day10part1(data []string) {
	sum := 0
	for _, l := range data {
		sum += parseLine(l)
	}
	fmt.Println(sum)
}

func stackToString() string {
	result := ""
	for c := pop(); c != 'x'; c = pop() {
		result += string(c)
	}
	return result
}

var scores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func autoCompleteScore(s string) int {
	score := 0
	for _, c := range s {
		score = score*5 + scores[c]
	}
	return score
}

func completeLine(s string) int {
	stack = []rune{}
	for _, r := range s {
		if r == '(' || r == '[' || r == '<' || r == '{' {
			// fmt.Println("PUSH ", string(r))
			push(tokens[r])
		} else {
			c := pop()
			// fmt.Println("POP ", string(c))
			if c != r {
				fmt.Println("ERROR: EXPECTED ", string(c), " GOT ", string(r))
				return 0
			}
		}
	}
	result := stackToString()
	fmt.Println("AUTOCOMPLETE: ", result, autoCompleteScore(result))
	return autoCompleteScore(result)
}

func day10part2(data []string) {
	scoreList := []int{}
	for _, l := range data {
		s := completeLine(l)
		if s != 0 {
			scoreList = append(scoreList, completeLine(l))
		}
	}
	sort.Ints(scoreList)
	fmt.Println(scoreList[len(scoreList)/2])
}

func main() {
	data, _ := readLinesFromFile("day10.txt")
	day10part1(data)
	day10part2(data)
}
