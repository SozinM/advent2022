package main

import (
	"slice"
	"strconv"
	"strings"
)

var total = 0
var checkpoint = []int{20, 60, 100, 140, 180, 220}

func main() {
	input := "addx 2\naddx 3\naddx 3\naddx -2\naddx 4\nnoop\naddx 1\naddx 4\naddx 1\nnoop\naddx 4\naddx 1\nnoop\naddx 2\naddx 5\naddx -28\naddx 30\nnoop\naddx 5\naddx 1\nnoop\naddx -38\nnoop\nnoop\nnoop\nnoop\naddx 5\naddx 5\naddx 3\naddx 2\naddx -2\naddx 2\nnoop\nnoop\naddx -2\naddx 12\nnoop\naddx 2\naddx 3\nnoop\naddx 2\naddx -31\naddx 32\naddx 7\nnoop\naddx -2\naddx -37\naddx 1\naddx 5\naddx 1\nnoop\naddx 31\naddx -25\naddx -10\naddx 13\nnoop\nnoop\naddx 18\naddx -11\naddx 3\nnoop\nnoop\naddx 1\naddx 4\naddx -32\naddx 15\naddx 24\naddx -2\nnoop\naddx -37\nnoop\nnoop\nnoop\naddx 5\naddx 5\naddx 21\naddx -20\nnoop\naddx 6\naddx 19\naddx -5\naddx -8\naddx -22\naddx 26\naddx -22\naddx 23\naddx 2\nnoop\nnoop\nnoop\naddx 8\naddx -10\naddx -27\naddx 33\naddx -27\nnoop\naddx 34\naddx -33\naddx 2\naddx 19\naddx -12\naddx 11\naddx -20\naddx 12\naddx 18\naddx -11\naddx -14\naddx 15\naddx 2\nnoop\naddx 3\naddx 2\nnoop\nnoop\nnoop\naddx -33\nnoop\naddx 1\naddx 2\nnoop\naddx 3\naddx 4\nnoop\naddx 1\naddx 2\nnoop\nnoop\naddx 7\naddx 1\nnoop\naddx 4\naddx -17\naddx 18\naddx 5\naddx -1\naddx 5\naddx 1\nnoop\nnoop\nnoop\nnoop"
	cycles := strings.Split(input, "\n")
	counter := 0
	signal := 0
	for _, command := range cycles {
		if command == "noop" {
			counter = tickCounter(signal, counter)
		} else {
			args := strings.Split(command, " ")
			commandSignal, _ := strconv.Atoi(args[1])
			signal += commandSignal
		}
	}
}

func tickCounter(signal int, counter int) int {
	if slice.Contains(checkpoint[:], counter) {
		total += signal
	}
	counter++
	return counter
}
