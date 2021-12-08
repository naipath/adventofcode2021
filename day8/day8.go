package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type dashboard struct {
	signal []string
	output []string
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	var dashboards []dashboard
	for _, line := range lines {
		split := strings.Split(line, "|")
		dashboards = append(dashboards,
			dashboard{
				signal: strings.Fields(split[0]),
				output: strings.Fields(split[1]),
			},
		)
	}

	fmt.Println("Part 1:", part1(dashboards))
}

func part1(dashboards []dashboard) int {
	occurrences := 0
	for _, d := range dashboards {
		for _, s := range d.output {
			if len(s) == 3 || len(s) == 4 || len(s) == 2 || len(s) == 7 {
				occurrences++
			}
		}
	}
	return occurrences
}
