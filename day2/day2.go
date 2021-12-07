package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	forward = "forward"
	down    = "down"
	up      = "up"
)

type instruction struct {
	command string
	amount  int
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	var instructions []instruction
	for _, line := range lines {
		split := strings.Split(line, " ")
		atoi, _ := strconv.Atoi(split[1])
		instructions = append(instructions, instruction{
			command: split[0],
			amount:  atoi,
		})
	}
	fmt.Println("Part 1:", part1(instructions))
	fmt.Println("Part 2:", part2(instructions))
}

func part1(instructions []instruction) int {
	horizontal := 0
	depth := 0
	for _, i := range instructions {
		switch i.command {
		case forward:
			horizontal += i.amount
		case down:
			depth += i.amount
		case up:
			depth -= i.amount
		}
	}
	return horizontal * depth
}

func part2(instructions []instruction) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, i := range instructions {
		switch i.command {
		case forward:
			horizontal += i.amount
			depth += i.amount * aim
		case down:
			aim += i.amount
		case up:
			aim -= i.amount
		}
	}
	return horizontal * depth
}
