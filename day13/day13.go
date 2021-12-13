package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type dot struct {
	x int
	y int
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n\n")

	var dots []dot
	for _, s := range strings.Split(lines[0], "\n") {
		split := strings.Split(s, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		dots = append(dots, dot{x, y})
	}
	var instructions []string
	for _, s := range strings.Split(lines[1], "\n") {
		split := strings.Split(s, "g ")
		instructions = append(instructions, split[1])
	}

	fmt.Println("Part 1:", len(handleFold(dots, instructions[0])))

	var newDots = dots
	for _, instruction := range instructions {
		newDots = handleFold(newDots, instruction)
	}
	fmt.Println("Part 2:")
	printDots(newDots)
}

func printDots(newDots []dot) {
	for y := 0; y < 6; y++ {
		for x := 0; x < 40; x++ {
			symbol := "."
			if findDot(newDots, x, y) {
				symbol = "#"
			}
			fmt.Print(symbol)

		}
		fmt.Println()
	}
}

func handleFold(dots []dot, instruction string) []dot {
	var newDots []dot
	split := strings.Split(instruction, "=")
	amount, _ := strconv.Atoi(split[1])
	for _, d := range dots {
		if split[0] == "x" {
			newDots = append(newDots, dot{
				x: compute(d.x, amount),
				y: d.y,
			})
		} else {
			newDots = append(newDots, dot{
				x: d.x,
				y: compute(d.y, amount),
			})
		}
	}
	return unique(newDots)
}

func compute(i int, amount int) int {
	newX := i
	if i > amount {
		newX = absolute(newX - amount*2)
	}
	return newX
}

func absolute(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func unique(arr []dot) []dot {
	occurred := map[dot]bool{}
	var result []dot
	for e := range arr {
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}

func findDot(dots []dot, x, y int) bool {
	for _, d := range dots {
		if d.x == x && d.y == y {
			return true
		}
	}
	return false
}
