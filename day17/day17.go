package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	compile, _ := regexp.Compile("-?\\d+")

	var field []int
	for _, s := range compile.FindAllString(string(data), 4) {
		atoi, _ := strconv.Atoi(s)
		field = append(field, atoi)
	}
	distinct, maxHighestPoint := resolveTrajectory(field)

	fmt.Println("Part 1:", maxHighestPoint)
	fmt.Println("Part 2:", distinct)
}

func resolveTrajectory(field []int) (int, int) {
	distinct := 0
	maxHighestPoint := 0
	for x := 0; x < field[1]+1; x++ {
		for y := field[2] - 1; y < 200; y++ {
			trajectory, highestPoint := computeTrajectory(x, y, field)
			if trajectory && highestPoint > maxHighestPoint {
				maxHighestPoint = highestPoint
			}
			if trajectory {
				distinct++
			}
		}
	}
	return distinct, maxHighestPoint
}

func computeTrajectory(xVelocity int, yVelocity int, field []int) (bool, int) {
	highestPoint := 0
	x := 0
	y := 0
	for true {
		x += xVelocity
		y += yVelocity

		if y > highestPoint {
			highestPoint = y
		}

		if xVelocity > 0 {
			xVelocity--
		} else if xVelocity < 0 {
			xVelocity++
		}
		yVelocity--

		if x >= field[0] && x <= field[1] && y >= field[2] && y <= field[3] {
			return true, highestPoint
		}
		if x > field[1] || y < field[2] {
			return false, highestPoint
		}
	}
	return false, highestPoint
}
