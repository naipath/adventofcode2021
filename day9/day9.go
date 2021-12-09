package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	points := make(map[int]map[int]int)

	for y, line := range lines {
		points[y] = make(map[int]int)
		for x, i3 := range line {
			atoi, _ := strconv.Atoi(string(i3))
			points[y][x] = atoi
		}
	}

	sum := part1(points)
	fmt.Println("Part 1:", sum)
}

func part1(points map[int]map[int]int) int {
	sum := 0
	for y, m := range points {
		for x, i := range m {
			isLower := true
			left, exists := points[y][x-1]
			if exists && left <= i {
				isLower = false
			}
			right, exists := points[y][x+1]
			if exists && right <= i {
				isLower = false
			}
			up, exists := points[y-1][x]
			if exists && up <= i {
				isLower = false
			}
			down, exists := points[y+1][x]
			if exists && down <= i {
				isLower = false
			}
			if isLower {
				sum += i + 1
			}
		}
	}
	return sum
}
