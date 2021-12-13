package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	re := regexp.MustCompile("([0-9])+")

	var points []string
	var diagonalPoints []string
	for _, line := range lines {
		results := re.FindAllString(line, 4)

		x1, _ := strconv.Atoi(results[0])
		y1, _ := strconv.Atoi(results[1])
		x2, _ := strconv.Atoi(results[2])
		y2, _ := strconv.Atoi(results[3])

		if x1 == x2 {
			for i := min(y1, y2); i <= max(y1, y2); i++ {
				points = append(points, strconv.Itoa(x1)+","+strconv.Itoa(i))
			}
		} else if y1 == y2 {
			for i := min(x1, x2); i <= max(x1, x2); i++ {
				points = append(points, strconv.Itoa(i)+","+strconv.Itoa(y1))
			}
		} else {
			difference := absolute(y1 - y2)

			for i := 0; i <= difference; i++ {
				nextX := 0
				if x1 > x2 {
					nextX = x1 - i
				} else {
					nextX = x1 + i
				}
				nextY := 0
				if y1 > y2 {
					nextY = y1 - i
				} else {
					nextY = y1 + i
				}
				diagonalPoints = append(diagonalPoints, strconv.Itoa(nextX)+","+strconv.Itoa(nextY))
			}
		}
	}

	var occurrences = make(map[string]int)

	findOccurrences(points, occurrences)
	fmt.Println("Part 1:", sum(occurrences))

	findOccurrences(diagonalPoints, occurrences)
	fmt.Println("Part 2:", sum(occurrences))
}

func findOccurrences(points []string, occurrences map[string]int) {
	for _, point := range points {
		occurrences[point] = occurrences[point] + 1
	}
}

func sum(occurences map[string]int) int {
	counter := 0
	for _, i := range occurences {
		if i > 1 {
			counter++
		}
	}
	return counter
}

func absolute(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
