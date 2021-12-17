package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

func binaryString(hs string) string {
	result := ""
	for _, c := range hs {
		b := c - '0'
		if b > 9 {
			b -= 7
		}
		result += fmt.Sprintf("%04b", b)
	}
	return result
}

func nextVersionPacket(s string) (int, int, int) {
	version := toInt(s[:3])
	v, p, r := nextSubPacket(s[3:])
	return version + v, p + 3, r
}

func lengthBitSize(s string) int {
	if s == "1" {
		return 11
	}
	return 15
}

func nextType(s string) (int, int) {
	if len(s) < 6 {
		return -1, len(s)
	}
	t := toInt(s[:3])
	return t, 3
}

func nextSubPacket(s string) (int, int, int) {
	vTotal := 0
	opResult := 0
	var opResults []int

	pType, startPos := nextType(s)
	if pType == 4 {
		r, newPos := parseLiteral(s[startPos:])
		startPos += newPos
		opResult = r
	} else {
		lType := lengthBitSize(s[startPos : startPos+1])
		startPos++
		pLen := toInt(s[startPos : startPos+lType])
		startPos += lType
		if lType == 15 {
			endPos := startPos + pLen
			for startPos < endPos {
				v, pos, val := nextVersionPacket(s[startPos:])
				vTotal += v
				startPos += pos
				opResults = append(opResults, val)
			}
		} else {
			// pLen is # subpackets
			for i := 0; i < pLen; i++ {
				v, pos, val := nextVersionPacket(s[startPos:])
				vTotal += v
				startPos += pos
				opResults = append(opResults, val)
			}
		}
		opResult = operation(pType, opResults)
	}

	return vTotal, startPos, opResult
}

/*
Packets with type ID 0 are sum packets - their value is the sum of the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
Packets with type ID 1 are product packets - their value is the result of multiplying together the values of their sub-packets. If they only have a single sub-packet, their value is the value of the sub-packet.
Packets with type ID 2 are minimum packets - their value is the minimum of the values of their sub-packets.
Packets with type ID 3 are maximum packets - their value is the maximum of the values of their sub-packets.
Packets with type ID 5 are greater than packets - their value is 1 if the value of the first sub-packet is greater than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
Packets with type ID 6 are less than packets - their value is 1 if the value of the first sub-packet is less than the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
Packets with type ID 7 are equal to packets - their value is 1 if the value of the first sub-packet is equal to the value of the second sub-packet; otherwise, their value is 0. These packets always have exactly two sub-packets.
*/
func operation(opType int, values []int) int {
	result := 0
	switch opType {
	case 0:
		for _, v := range values {
			result += v
		}
		break
	case 1:
		result = 1
		for _, v := range values {
			result *= v
		}
		break
	case 2:
		result = math.MaxInt
		for _, v := range values {
			if v < result {
				result = v
			}
		}
		break
	case 3:
		result = math.MinInt
		for _, v := range values {
			if v > result {
				result = v
			}
		}
		break
	case 5:
		if values[0] > values[1] {
			result = 1
		}
		break
	case 6:
		if values[0] < values[1] {
			result = 1
		}
		break
	case 7:
		if values[0] == values[1] {
			result = 1
		}
		break
	}
	results[opType] += result
	return result
}

func parsePackets(s string) (int, int) {
	startPos := 0
	version, _, result := nextVersionPacket(s[startPos:])
	return version, result
}

func toInt(b string) int {
	v, _ := strconv.ParseInt(b, 2, 64)
	return int(v)
}

func parseLiteral(s string) (int, int) {
	literal := ""
	startPos := 0
	for {
		d := s[startPos : startPos+1]
		v := s[startPos+1 : startPos+5]
		startPos += 5
		literal += fmt.Sprint(v)
		if d != "1" {
			break
		}
	}
	v, _ := strconv.ParseInt(literal, 2, 64)
	return int(v), startPos
}

func day16part1(data []string) {
	for _, d := range data {
		results = make(map[int]int)
		b := binaryString(d)
		v, _ := parsePackets(b)
		fmt.Println("S: ", d, "V: ", v)
	}
}

var results map[int]int

func day16part2(data []string) {
	for _, d := range data {
		results = make(map[int]int)
		b := binaryString(d)
		v, r := parsePackets(b)
		_ = v
		fmt.Println(r)
	}
}

func main() {
	data, _ := readLinesFromFile("day16.txt")
	day16part1(data)
	fmt.Println()
	day16part2(data)
}
