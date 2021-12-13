package main

import (
	"fmt"
	"io/ioutil"
	"math"
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

	fmt.Println("Part 1:", dumboFlashes(points, 100))
	fmt.Println("Part 2:", dumboFlashes(points, math.MaxInt))
}

func dumboFlashes(inputPoints []point, limit int) int {
	points := make([]point, len(inputPoints))
	copy(points, inputPoints)
	flashes := 0
	for day := 0; ; day++ {
		if day == limit {
			return flashes
		}
		allZero := true
		for _, p := range points {
			if p.value != 0 {
				allZero = false
			}
		}
		if allZero {
			return day
		}
		for i := range points {
			points[i].value += 1
		}
		for true {
			hasFlash := false
			for i := range points {
				if points[i].value > 9 {
					hasFlash = true
					flashes++
					points[i].value = 0

					for i2, p := range points {
						if p.x == points[i].x && p.y == points[i].y {
							continue
						}
						if p.x >= points[i].x-1 && p.x <= points[i].x+1 && p.y >= points[i].y-1 && p.y <= points[i].y+1 {
							if p.value != 0 {
								points[i2].value++
							}
						}
					}
				}
			}
			if !hasFlash {
				break
			}
		}
	}
}
