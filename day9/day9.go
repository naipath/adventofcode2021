package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	x     int
	y     int
	value int
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")

	var points []point

	for y, line := range lines {
		for x, i3 := range line {
			atoi, _ := strconv.Atoi(string(i3))
			points = append(points, point{x, y, atoi})
		}
	}

	lowPoints := findLowPoints(points)

	fmt.Println("Part 1:", part1(lowPoints))
	fmt.Println("Part 2:", part2(lowPoints, points))
}

func part1(lowPoints []point) int {
	sum := 0
	for _, lowPoint := range lowPoints {
		sum += lowPoint.value + 1
	}
	return sum
}

func part2(lowPoints []point, points []point) int {
	var areaSizes []int
	for _, lowPoint := range lowPoints {
		areaSizes = append(areaSizes, len(findAreaForLowPoint(lowPoint, points)))
	}
	sort.Ints(areaSizes)
	result := 1
	for _, area := range areaSizes[len(areaSizes)-3:] {
		result *= area
	}
	return result
}

func findAreaForLowPoint(lowPoint point, points []point) []point {
	var area []point
	area = append(area, lowPoint)

	for true {
		currentAreaSize := len(area)
		for _, p := range area {
			left, exists := findPoint(points, p.x-1, p.y)
			if exists && left.value != 9 {
				_, existsInArea := findPoint(area, left.x, left.y)
				if !existsInArea {
					area = append(area, left)
				}
			}
			right, exists := findPoint(points, p.x+1, p.y)
			if exists && right.value != 9 {
				_, existsInArea := findPoint(area, right.x, right.y)
				if !existsInArea {
					area = append(area, right)
				}
			}
			up, exists := findPoint(points, p.x, p.y-1)
			if exists && up.value != 9 {
				_, existsInArea := findPoint(area, up.x, up.y)
				if !existsInArea {
					area = append(area, up)
				}
			}
			down, exists := findPoint(points, p.x, p.y+1)
			if exists && down.value != 9 {
				_, existsInArea := findPoint(area, down.x, down.y)
				if !existsInArea {
					area = append(area, down)
				}
			}
		}
		if len(area) == currentAreaSize {
			break
		}
	}
	return area
}

func findLowPoints(points []point) []point {
	var lowPoints []point

	for _, p := range points {
		isLower := true
		left, exists := findPoint(points, p.x-1, p.y)
		if exists && left.value <= p.value {
			isLower = false
		}
		right, exists := findPoint(points, p.x+1, p.y)
		if exists && right.value <= p.value {
			isLower = false
		}
		up, exists := findPoint(points, p.x, p.y-1)
		if exists && up.value <= p.value {
			isLower = false
		}
		down, exists := findPoint(points, p.x, p.y+1)
		if exists && down.value <= p.value {
			isLower = false
		}
		if isLower {
			lowPoints = append(lowPoints, p)
		}
	}
	return lowPoints
}

func findPoint(points []point, x, y int) (point, bool) {
	for _, p := range points {
		if p.x == x && p.y == y {
			return p, true
		}
	}
	return point{}, false
}
