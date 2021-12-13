package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), ",")

	var lanternFish []uint8

	for _, i := range lines {
		atoi, _ := strconv.Atoi(i)
		lanternFish = append(lanternFish, uint8(atoi))
	}

	var initialGeneration = make(map[uint8]uint64)

	for _, fish := range lanternFish {
		initialGeneration[fish] = initialGeneration[fish] + 1
	}

	fmt.Println("Part 1", totalFish(initialGeneration, 80))
	fmt.Println("Part 2", totalFish(initialGeneration, 256))
}

func totalFish(initialGeneration map[uint8]uint64, days int) uint64 {
	for i := 0; i < days; i++ {
		var nextGeneration = make(map[uint8]uint64)

		for k, v := range initialGeneration {
			if k == 0 {
				nextGeneration[6] = nextGeneration[6] + v
				nextGeneration[8] = nextGeneration[8] + v
			} else {
				nextGeneration[k-1] = nextGeneration[k-1] + v
			}
		}
		initialGeneration = nextGeneration
	}

	var counter uint64 = 0
	for u := range initialGeneration {
		counter += initialGeneration[u]
	}
	return counter
}
