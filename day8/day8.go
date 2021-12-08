package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type dashboard struct {
	signal      []string
	output      []string
	translation map[string]int
}

func (receiver *dashboard) result() int {
	number := ""
	for _, s := range receiver.output {
		number = number + strconv.Itoa(receiver.translation[s])
	}
	atoi, _ := strconv.Atoi(number)
	return atoi
}

func (receiver *dashboard) resolveLetters() int {
	receiver.translation = make(map[string]int)

	seven := ""
	four := ""
	for _, s := range receiver.signal {
		if len(s) == 2 {
			receiver.translation[s] = 1
		}
		if len(s) == 3 {
			receiver.translation[s] = 7
			seven = s
		}
		if len(s) == 4 {
			receiver.translation[s] = 4
			four = s
		}
		if len(s) == 7 {
			receiver.translation[s] = 8
		}
	}
	three := ""
	for _, s := range receiver.signal {
		if len(s) != 5 {
			continue
		}
		if matches(seven, s) == 3 {
			receiver.translation[s] = 3
			three = s
		}
	}
	nine := ""
	for _, s := range receiver.signal {
		if len(s) != 6 {
			continue
		}
		if matches(four, s) == 4 {
			receiver.translation[s] = 9
			nine = s
		}
	}
	for _, s := range receiver.signal {
		if s == nine || len(s) != 6 {
			continue
		}
		if matches(seven, s) == 3 {
			receiver.translation[s] = 0
		} else {
			receiver.translation[s] = 6
		}
	}
	for _, s := range receiver.signal {
		if s == three || len(s) != 5 {
			continue
		}
		if matches(nine, s) == 5 {
			receiver.translation[s] = 5
		} else {
			receiver.translation[s] = 2
		}
	}

	return receiver.result()
}

func matches(seven string, s string) int {
	count := 0
	for i := range seven {
		if strings.Contains(s, string(seven[i])) {
			count++
		}
	}
	return count
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	var dashboards []dashboard
	for _, line := range lines {
		split := strings.Split(line, "|")
		signal := strings.Fields(split[0])
		for i, field := range signal {
			signal[i] = sortString(field)
		}
		output := strings.Fields(split[1])
		for i, s := range output {
			output[i] = sortString(s)
		}
		dashboards = append(dashboards,
			dashboard{
				signal: signal,
				output: output,
			},
		)
	}

	fmt.Println("Part 1:", part1(dashboards))

	sum := 0
	for _, d := range dashboards {
		sum += d.resolveLetters()
	}
	fmt.Println("Part 2:", sum)
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

func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}
