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

	var numbers []int
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	fmt.Println("Part 1:", part1(numbers))
	fmt.Println("Part 2:", part2(numbers))
}

func part1(numbers []int) int {
	var increases = 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[i-1] {
			increases++
		}
	}
	return increases
}

func part2(numbers []int) int {
	var windowIncreases = 0
	for i := 1; i < len(numbers)-2; i++ {
		totalCurr := numbers[i] + numbers[i+1] + numbers[i+2]
		totalPrevious := numbers[i-1] + numbers[i] + numbers[i+1]
		if totalCurr > totalPrevious {
			windowIncreases++
		}
	}
	return windowIncreases
}
