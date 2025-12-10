package main

import (
	"fmt"
)

func main() {
	rectangles, numPoints := parseInput()

	part1(rectangles, numPoints)
	part2(rectangles, numPoints)
}

func part1(rectangles []Rectangle, numPoints int) {
	fmt.Println(rectangles[0])
}

func part2(rectangles []Rectangle, numPoints int) {
	// TODO
}
