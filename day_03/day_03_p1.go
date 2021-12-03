package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(fname string) (readings []string, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	readings = make([]string, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		readings = append(readings, l)
	}
	return readings, nil
}

func computeValues(readings []string) (gamma int, epsilon int) {
	gamma = 0
	epsilon = 0
	size := len(readings[0])
	for pos := 0; pos < size; pos++ {
		ones := 0
		zeros := 0
		for _, c := range readings {
			if c[pos] == '1' {
				ones++
			} else {
				zeros++
			}
		}
		gamma = gamma << 1
		epsilon = epsilon << 1
		if ones > zeros {
			gamma = gamma + 1
		} else {
			epsilon = epsilon + 1
		}
	}
	return gamma, epsilon
}

func main() {
	readings, err := readFile("solve-data.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(computeValues(readings))
}
