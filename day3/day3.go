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

	lengthOfLine := len(lines[0])
	totalLines := len(lines)

	fmt.Println("Part 1:", part1(lengthOfLine, lines, totalLines))
	fmt.Println("Part 2:", part2(lines, lengthOfLine))
}

func part1(lengthOfLine int, lines []string, totalLines int) int64 {
	var averagePlus = make([]int, lengthOfLine)
	for _, line := range lines {
		for i := range line {
			if line[i] == '1' {
				averagePlus[i] = averagePlus[i] + 1
			}
		}
	}

	plusString := ""
	minusString := ""
	for i := 0; i < lengthOfLine; i++ {
		if averagePlus[i] > totalLines/2 {
			plusString += "1"
			minusString += "0"
		} else {
			plusString += "0"
			minusString += "1"
		}
	}
	gamma, _ := strconv.ParseInt(plusString, 2, 64)
	epsilon, _ := strconv.ParseInt(minusString, 2, 64)
	return gamma * epsilon
}

func part2(lines []string, lengthOfLine int) int64 {
	agamma, _ := strconv.ParseInt(findAverage(lines, lengthOfLine, true), 2, 64)
	aepsilon, _ := strconv.ParseInt(findAverage(lines, lengthOfLine, false), 2, 64)
	return agamma * aepsilon
}

func findAverage(lines []string, lengthOfLine int, shouldBePositive bool) string {
	var current []string
	current = append(current, lines...)
	for i := 0; i < lengthOfLine; i++ {
		averageOnes := 0
		for _, line := range current {
			if line[i] == '1' {
				averageOnes++
			}
		}

		var expected uint8
		if shouldBePositive {
			expected = '0'
			if averageOnes >= len(current)/2 {
				expected = '1'
			}
		} else {
			expected = '1'
			if averageOnes >= len(current)/2 {
				expected = '0'
			}
		}

		var nextCurrent []string
		for _, line := range current {
			if line[i] == expected {
				nextCurrent = append(nextCurrent, line)
			}
		}

		if len(nextCurrent) == 0 {
			break
		}
		current = nextCurrent
	}
	return current[0]
}
