package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n\n")

	initialString := lines[0]
	conversions := map[string]string{}
	for _, s := range strings.Split(lines[1], "\n") {
		split := strings.Split(s, " -> ")
		conversions[split[0]] = split[1]
	}

	fmt.Println("Part 1:", computeScore(computePolymer(initialString, conversions, 10)))
}

func computeScore(s string) int {
	counting := map[string]int{}
	for _, i2 := range s {
		counting[string(i2)] = counting[string(i2)] + 1
	}

	lowestCount := math.MaxInt
	highestCount := 0

	for _, u := range counting {
		if u > highestCount {
			highestCount = u
		}
		if u < lowestCount {
			lowestCount = u
		}

	}
	return highestCount - lowestCount
}

func computePolymer(initialString string, conversions map[string]string, steps int) string {
	var b strings.Builder
	for i := 0; i < steps; i++ {
		b.Grow(len(initialString))
		for i := 0; i < len(initialString); i++ {
			currentLetter := initialString[i : i+1]
			if i < len(initialString)-1 {
				result, exists := conversions[initialString[i:i+2]]
				if exists {
					b.WriteString(currentLetter + result)
				} else {
					b.WriteString(currentLetter)
				}
			} else {
				b.WriteString(currentLetter)
			}
		}
		initialString = b.String()
		b.Reset()
	}
	return initialString
}
