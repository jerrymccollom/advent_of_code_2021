package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	Forward = 0
	Up      = -1
	Down    = +1
)

type Command struct {
	Direction int
	Distance  int
}

func readFile(fname string) (cmds []Command, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	cmds = make([]Command, 0, len(lines))

	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		c := strings.Split(l, " ")
		cmd := new(Command)
		switch c[0][0] {
		case 'f':
			cmd.Direction = Forward
			break
		case 'd':
			cmd.Direction = Down
			break
		case 'u':
			cmd.Direction = Up
			break
		default:
			continue
		}
		cmd.Distance, err = strconv.Atoi(c[1])
		if err != nil {
			return nil, err
		}
		cmds = append(cmds, *cmd)
	}
	return cmds, nil
}

func computeLocation(cmds []Command) (X int, Y int) {
	X = 0
	Y = 0
	for _, c := range cmds {
		if c.Direction == Forward {
			X += c.Distance
		} else {
			Y += c.Direction * c.Distance
		}
	}
	return X, Y
}

func main() {
	cmds, err := readFile("solve-commands.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(computeLocation(cmds))
}
