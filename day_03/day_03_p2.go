package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func computeValues(readings []string) ([]string, []string) {
	size := len(readings[0])
	oresult := readings
	cresult := readings
	for pos := 0; pos < size; pos++ {
		ones, zeros := onesAndZerosAtPosition(oresult, pos)
		if len(oresult) > 1 {
			var n []string
			for i, c := range oresult {
				if ((ones >= zeros) && c[pos] == '1') || ((zeros > ones) && c[pos] == '0') {
					if i < len(oresult) {
						n = append(n, c)
					}
				}
			}
			oresult = n
		}
		ones, zeros = onesAndZerosAtPosition(cresult, pos)
		if len(cresult) > 1 {
			var n []string
			for i, c := range cresult {
				if ((ones >= zeros) && c[pos] == '0') || ((zeros > ones) && c[pos] == '1') {
					if i < len(cresult) {
						n = append(n, c)
					}
				}
			}
			cresult = n
		}

	}
	return oresult, cresult
}

func onesAndZerosAtPosition(oresult []string, pos int) (int, int) {
	ones := 0
	zeros := 0
	for _, c := range oresult {
		if c[pos] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return ones, zeros
}

func binaryArrayToInt(s []string) int64 {
	result := ""
	for _, c := range s {
		result += c
	}
	v, _ := strconv.ParseInt(result, 2, 64)
	return v
}

func main() {
	readings, err := readFile("solve-data.txt")
	if err != nil {
		panic(err)
	}
	o, c := computeValues(readings)
	oval := binaryArrayToInt(o)
	cval := binaryArrayToInt(c)
	fmt.Println(oval, cval, oval*cval)
}
