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
			left := findPoint(points, p.x-1, p.y)
			if left.value < 9 && !pointExists(area, left.x, left.y) {
				area = append(area, left)
			}
			right := findPoint(points, p.x+1, p.y)
			if right.value < 9 && !pointExists(area, right.x, right.y) {
				area = append(area, right)
			}
			up := findPoint(points, p.x, p.y-1)
			if up.value < 9 && !pointExists(area, up.x, up.y) {
				area = append(area, up)
			}
			down := findPoint(points, p.x, p.y+1)
			if down.value < 9 && !pointExists(area, down.x, down.y) {
				area = append(area, down)
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
		left := findPoint(points, p.x-1, p.y)
		right := findPoint(points, p.x+1, p.y)
		up := findPoint(points, p.x, p.y-1)
		down := findPoint(points, p.x, p.y+1)
		if !(left.value <= p.value || right.value <= p.value || up.value <= p.value || down.value <= p.value) {
			lowPoints = append(lowPoints, p)
		}
	}
	return lowPoints
}

func findPoint(points []point, x, y int) point {
	for _, p := range points {
		if p.x == x && p.y == y {
			return p
		}
	}
	return point{0, 0, 9}
}

func pointExists(points []point, x, y int) bool {
	for _, p := range points {
		if p.x == x && p.y == y {
			return true
		}
	}
	return false
}
