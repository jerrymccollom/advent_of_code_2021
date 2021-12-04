package main

import (
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

func stringsToIntegers(strings []string) (ints []int) {
	// filter out empty lines
	ints = make([]int, 0, len(strings))
	for _, l := range strings {
		if len(l) == 0 {
			continue
		}
		v, _ := strconv.Atoi(l)
		ints = append(ints, v)
	}
	return ints
}
