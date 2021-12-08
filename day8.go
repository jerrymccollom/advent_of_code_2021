package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

// Lifted from example sorting characters in a string
// https://siongui.github.io/2017/05/07/go-sort-string-slice-of-rune/
type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}

func getLengths(s []string) []int {
	result := make([]int, len(s))
	for i, r := range s {
		result[i] = len(r)
	}
	return result
}

var lengths = []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

func count1478(output []int) int {
	count := 0
	for _, s := range output {
		if s == lengths[1] || s == lengths[4] || s == lengths[7] || s == lengths[8] {
			count++
		}
	}
	return count
}

func day8part1(data []string) {
	count := 0
	for _, line := range data {
		pts := strings.Split(line, "|")
		display := getLengths(strings.Split(pts[1], " "))
		count += count1478(display)
	}
	fmt.Println(count)
}

func intersection(a, b string) (result string) {
	counts := make([]int, 256)
	for i := 0; i < 256; i++ {
		counts[i] = 0
	}
	for _, c := range a {
		counts[c]++
	}
	for _, c := range b {
		counts[c]++
	}
	for i := 'a'; i <= 'g'; i++ {
		if counts[i] == 2 {
			result += string(i)
		}
	}
	return result
}

func diff(a, b string) (result string) {
	counts := make([]int, 256)
	for i := 0; i < 256; i++ {
		counts[i] = 0
	}
	for _, c := range a {
		counts[c]++
	}
	for _, c := range b {
		counts[c] = 0
	}
	for i := 'a'; i <= 'g'; i++ {
		if counts[i] == 1 {
			result += string(i)
		}
	}
	return result
}

func decodeWires(readings []string) []string {
	segs := []string{"*", "*", "*", "*", "*", "*", "*", "*", "*", "*"}
	// Program 1478
	for _, r := range readings {
		switch len(r) {
		case 2:
			if segs[1] == "*" {
				segs[1] = r
			}
			break
		case 4:
			if segs[4] == "*" {
				segs[4] = r
			}
			break
		case 3:
			if segs[7] == "*" {
				segs[7] = r
			}
			break
		case 7:
			if segs[8] == "*" {
				segs[8] = r
			}
			break
		}
	}
	// Deduce 6 and 9 and 0 -- 6 does not contain the segments for 1
	for _, r := range readings {
		if len(r) == 6 {
			i := intersection(r, segs[1])
			if len(i) == 1 {
				if segs[6] == "*" {
					segs[6] = r
				}
			} else {
				d2 := diff(r, segs[4])
				if len(d2) == 2 {
					if segs[9] == "*" {
						segs[9] = r
					}
				} else {
					if segs[0] == "*" {
						segs[0] = r
					}
				}
			}
		}
		if len(r) == 5 {
			s := intersection(segs[1], r)
			if len(s) == 2 {
				if segs[3] == "*" {
					segs[3] = r
				}
			} else {
				s = intersection(segs[4], r)
				if len(s) == 2 {
					if segs[2] == "*" {
						segs[2] = r
					}
				} else {
					if segs[5] == "*" {
						segs[5] = r
					}
				}
			}
		}
	}
	for i := 0; i < 10; i++ {
		if segs[i] == "*" {
			fmt.Println("No value for ", i)
		}
	}
	return segs
}

func outputValues(output, segs []string) string {
	result := ""
	for _, s := range output {
		s = SortStringByCharacter(s)
		for i, found := 0, false; i < 10 && !found; i++ {
			if s == SortStringByCharacter(segs[i]) {
				result += fmt.Sprintf("%d", i)
				found = true
			}
		}
	}
	return result
}

func day8part2(data []string) {
	count := 0
	for _, line := range data {
		pts := strings.Split(line, "|")
		readings := strings.Split(pts[0], " ")
		output := strings.Split(pts[1], " ")
		segs := decodeWires(append(readings, output...))
		v, _ := strconv.Atoi(outputValues(output, segs))
		count += v
	}
	fmt.Println(count)
}

func main() {
	data, _ := readLinesFromFile("day8.txt")
	day8part1(data)
	day8part2(data)
}
