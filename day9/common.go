package main

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x   int
	y   int
	num int
}

type Rectangle struct {
	p1   Point
	p2   Point
	area int
}

func parseInput() ([]Rectangle, int) {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	numPoints := len(lines)

	var points []Point = make([]Point, numPoints)
	var rectangles []Rectangle = make([]Rectangle, 0, numPoints*(numPoints-1)/2)

	for pos, line := range lines {
		coords := make([]int, 2)
		for i, coordString := range strings.Split(line, ",") {
			coord, _ := strconv.Atoi(coordString)
			coords[i] = coord
		}

		newPoint := Point{coords[0], coords[1], pos}

		for j := range pos {
			pPrev := points[j]
			rectangles = append(rectangles, Rectangle{pPrev, newPoint, calcArea(pPrev, newPoint)})
		}
		points[pos] = newPoint
	}

	sort.Slice(rectangles, func(p, q int) bool {
		return rectangles[p].area > rectangles[q].area
	})

	return rectangles, numPoints
}

func calcArea(p1 Point, p2 Point) int {
	return int(math.Abs(float64((p1.x - p2.x + 1) * (p1.y - p2.y + 1))))
}
