package main

import (
	"fmt"
	"sort"
)

const CONNECTIONS_TO_PROCESS = 1000
const NUMBERS_TO_MULT = 3

func main() {
	combinations, numPoints := parseInput()

	part1(combinations, numPoints)
	part2(combinations, numPoints)
}

func part1(combinations []PointComb, numPoints int) {
	dsu := NewDSU(numPoints)

	// Iterating over closest 1000 pairs
	for i := range CONNECTIONS_TO_PROCESS {
		comb := combinations[i]
		p1Idx := comb.p1.num
		p2Idx := comb.p2.num

		dsu.Union(p1Idx, p2Idx)
	}

	// Collect all circuit sizes. The size of a circuit is stored at its root.
	var circuitSizes []int
	for i := range numPoints {
		if dsu.parent[i] == i { // 'i' is the root of a set (a circuit)
			circuitSizes = append(circuitSizes, dsu.size[i])
		}
	}

	// Sort the sizes in descending order.
	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})

	fmt.Println(circuitSizes)

	result := 1
	for i := range NUMBERS_TO_MULT {
		result *= circuitSizes[i]
	}

	fmt.Println(result)
}

func part2(combinations []PointComb, numPoints int) {
	dsu := NewDSU(numPoints)

	successCount := 0
	// Iterating over all closest pairs, until they are all in one Set
	for i := range len(combinations) {
		comb := combinations[i]
		p1Idx := comb.p1.num
		p2Idx := comb.p2.num

		if dsu.Union(p1Idx, p2Idx) {
			successCount++
		}

		if successCount == numPoints-1 {
			// This is the final connection
			fmt.Println(comb.p1.x * comb.p2.x)
			break
		}
	}
}
