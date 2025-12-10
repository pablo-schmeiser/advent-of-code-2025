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
	z   int
	num int
}

type PointComb struct {
	p1   Point
	p2   Point
	dist float64
}

func parseInput() ([]PointComb, int) {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	numPoints := len(lines)

	var points []Point = make([]Point, numPoints)
	var combinations []PointComb = make([]PointComb, 0, numPoints*(numPoints-1)/2)

	for pos, line := range lines {
		coords := make([]int, 3)
		for i, coordString := range strings.Split(line, ",") {
			coord, _ := strconv.Atoi(coordString)
			coords[i] = coord
		}

		newPoint := Point{coords[0], coords[1], coords[2], pos}

		for j := range pos {
			pPrev := points[j]
			combinations = append(combinations, PointComb{pPrev, newPoint, getDist(pPrev, newPoint)})
		}
		points[pos] = newPoint
	}

	sort.Slice(combinations, func(p, q int) bool {
		return combinations[p].dist < combinations[q].dist
	})

	return combinations, numPoints
}

func getDist(p1 Point, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2))
}

// DSU (Disjoint Set Union) structure for managing circuits/sets
type DSU struct {
	parent []int // parent[i] is the parent of element i
	size   []int // size[i] stores the size of the set rooted at i
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range n {
		parent[i] = i // Each element is its own parent
		size[i] = 1   // Each set has size 1
	}
	return &DSU{parent: parent, size: size}
}

// Find operation with Path Compression
func (dsu *DSU) Find(i int) int {
	if dsu.parent[i] == i {
		return i
	}
	dsu.parent[i] = dsu.Find(dsu.parent[i]) // Path compression
	return dsu.parent[i]
}

// Union operation by size. Returns true if a merge occurred (if they were in different circuits).
func (dsu *DSU) Union(i, j int) bool {
	rootI := dsu.Find(i)
	rootJ := dsu.Find(j)

	if rootI != rootJ {
		// Union by size (attach smaller set to larger set)
		if dsu.size[rootI] < dsu.size[rootJ] {
			rootI, rootJ = rootJ, rootI
		}

		dsu.parent[rootJ] = rootI
		dsu.size[rootI] += dsu.size[rootJ]
		return true // Merge occurred
	}
	return false // Already connected
}
