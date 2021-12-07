package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), ",")
	var crabs []int
	lowestPosition := 0
	highestPosition := 0
	for _, line := range lines {
		atoi, _ := strconv.Atoi(line)
		crabs = append(crabs, atoi)

		if atoi > highestPosition {
			highestPosition = atoi
		}
	}
	fmt.Println("Part 1:", part1(lowestPosition, highestPosition, crabs))
	fmt.Println("Part 2:", part2(lowestPosition, highestPosition, crabs))
}

func part1(lowestPosition int, highestPosition int, crabs []int) uint64 {
	var minimalFuelCost uint64 = math.MaxUint64
	for i := lowestPosition; i <= highestPosition; i++ {
		var fuelCost uint64 = 0
		for _, crab := range crabs {
			cost := absolute(crab - i)
			fuelCost = fuelCost + uint64(cost)
		}
		if fuelCost < minimalFuelCost {
			minimalFuelCost = fuelCost
		}
	}
	return minimalFuelCost
}

func part2(lowestPosition int, highestPosition int, crabs []int) int {
	var minimalFuelCost = math.MaxInt
	for i := lowestPosition; i <= highestPosition; i++ {
		var fuelCost = 0
		for _, crab := range crabs {
			steps := absolute(crab - i)
			fuelCost += steps * (steps + 1) / 2
		}
		if fuelCost < minimalFuelCost {
			minimalFuelCost = fuelCost
		}
	}
	return minimalFuelCost
}

func absolute(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
