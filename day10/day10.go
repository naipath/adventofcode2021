package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	fmt.Println("Part 1:", part1(lines))

	scores := part2(lines)
	fmt.Println("Part 2:", scores[len(scores)/2])
}

var scoringPart1 = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func part1(lines []string) int {
	score := 0
	for _, line := range lines {
		corruptedChar, _ := findCorruption(line)
		score += scoringPart1[corruptedChar]
	}
	return score
}

var scoringPart2 = map[string]int{
	"(": 1,
	"[": 2,
	"{": 3,
	"<": 4,
}

func part2(lines []string) []int {
	var scores []int
	for _, line := range lines {
		corruption, stack := findCorruption(line)
		if corruption == "" {
			reversedStack := reverse(stack)
			totalScore := 0
			for _, s := range reversedStack {
				totalScore *= 5
				totalScore += scoringPart2[s]
			}
			scores = append(scores, totalScore)
		}
	}

	sort.Ints(scores)
	return scores
}

func findCorruption(line string) (string, stack) {
	stackOpeners := make(stack, 0)

	var closers []string

	for _, c := range line {
		char := string(c)

		if char == "<" || char == "[" || char == "(" || char == "{" {
			stackOpeners = stackOpeners.Push(char)
		}

		if char == ">" || char == "]" || char == ")" || char == "}" {
			s, lastOpener := stackOpeners.Pop()
			stackOpeners = s
			if lastOpener == "<" && char != ">" {
				return char, stackOpeners
			}
			if lastOpener == "[" && char != "]" {
				return char, stackOpeners
			}
			if lastOpener == "(" && char != ")" {
				return char, stackOpeners
			}
			if lastOpener == "{" && char != "}" {
				return char, stackOpeners
			}
			closers = append(closers)
		}
	}
	return "", stackOpeners
}

func reverse(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}
