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
	fmt.Println("Part 2:", computeScore(computePolymer(initialString, conversions, 40)))
}

func computeScore(s map[string]uint64) uint64 {
	var lowestCount uint64 = math.MaxUint64
	var highestCount uint64 = 0
	for _, u := range s {
		if u > highestCount {
			highestCount = u
		}
		if u < lowestCount {
			lowestCount = u
		}
	}
	return highestCount - lowestCount
}

func computePolymer(initialString string, conversions map[string]string, steps int) map[string]uint64 {
	pairs := map[string]uint64{}
	letterOccurrence := map[string]uint64{}
	for i := 0; i < len(initialString); i++ {
		if i < len(initialString)-1 {
			pairs[initialString[i:i+2]] += 1
		}
		letterOccurrence[string(initialString[i])] += 1
	}
	for i := 0; i < steps; i++ {
		copyPairs := map[string]uint64{}
		for k, v := range pairs {
			copyPairs[k] = v
		}
		for s, count := range copyPairs {
			n := conversions[s]
			letterOccurrence[n] += count
			pairs[s] -= count
			pairs[string(s[0])+n] += count
			pairs[n+string(s[1])] += count
		}
	}
	return letterOccurrence
}
