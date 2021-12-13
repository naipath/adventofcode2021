package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	nodeMap := make(map[string][]string)
	for _, line := range lines {
		nodes := strings.Split(line, "-")
		nodeMap[nodes[0]] = append(nodeMap[nodes[0]], nodes[1])
		nodeMap[nodes[1]] = append(nodeMap[nodes[1]], nodes[0])
	}

	findPathsPart1(nodeMap, []string{"start"})
	fmt.Println("Part 1", accumulator)
	findPathsPart2(nodeMap, []string{"start"})
	fmt.Println("Part 2", accumulator2)
}

var accumulator = 0
var accumulator2 = 0

func findPathsPart2(nodeMap map[string][]string, currentPath []string) {
	nextNodes := nodeMap[currentPath[len(currentPath)-1]]

	for _, node := range nextNodes {
		if node == "start" {
			continue
		}
		if node == "end" {
			accumulator2++
			continue
		}
		if node != strings.ToUpper(node) {
			isValid := isNewCaveValid(currentPath, node)
			if !isValid {
				continue
			}
		}
		findPathsPart2(nodeMap, append(currentPath, node))
	}
}

func findPathsPart1(nodeMap map[string][]string, currentPath []string) {
	nextNodes := nodeMap[currentPath[len(currentPath)-1]]

	for _, node := range nextNodes {
		if node == "start" {
			continue
		}
		if node == "end" {
			accumulator++
			continue
		}
		if node != strings.ToUpper(node) {
			alreadyHas := false
			for _, n := range currentPath {
				if n == node {
					alreadyHas = true
				}
			}
			if alreadyHas {
				continue
			}
		}
		newCurrentPath := append(currentPath, node)
		findPathsPart1(nodeMap, newCurrentPath)
	}
}

func isNewCaveValid(currentPath []string, node string) bool {
	dict := make(map[string]int)
	for _, cave := range currentPath {
		if cave != strings.ToUpper(cave) {
			dict[cave] = dict[cave] + 1
		}
	}
	dict[node] = dict[node] + 1

	isValid := true
	countsWith2 := 0
	for _, i := range dict {
		if i > 2 {
			isValid = false
		}
		if i == 2 {
			countsWith2++
		}
	}
	if countsWith2 > 1 {
		isValid = false
	}
	return isValid
}
