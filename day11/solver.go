package main

import (
	"fmt"
)

func main() {
	nodes, _ := parseInput()
	part1(nodes)
	part2(nodes)
}

func part1(nodes []Node) {
	sum := 0

	known := make([]Node, 0)
	for _, node := range nodes {
		if node.id == "you" {
			sum = node.solve(nodes, known)
			break
		}
	}

	fmt.Println(sum)
}

func (n Node) solve(nodes []Node, known []Node) int {
	paths := 0

	if n.id == "out" {
		return 1
	}

	known = append(known, n)
	for _, cIdx := range n.childIdx {
		child := nodes[cIdx]
		circular := false
		for _, node := range known {
			if child.id == node.id {
				// Don't allow circular graphs
				circular = true
				break
			}
		}
		if circular {
			continue
		}

		paths += child.solve(nodes, known)
	}

	return paths
}

func part2(nodes []Node) {
	known := make(map[string]struct{})
	svrFft := countPaths(nodes, known, "svr", "fft")
	fmt.Printf("%d Paths from svr to fft\n", svrFft)
	known = make(map[string]struct{})
	fftDac := countPaths(nodes, known, "fft", "dac")
	known = make(map[string]struct{})
	dacOut := countPaths(nodes, known, "dac", "out")

	known = make(map[string]struct{})
	svrDac := countPaths(nodes, known, "svc", "dac")
	known = make(map[string]struct{})
	dacFft := countPaths(nodes, known, "dac", "fft")
	known = make(map[string]struct{})
	fftOut := countPaths(nodes, known, "fft", "out")

	fmt.Println(svrFft*fftDac*dacOut + svrDac*dacFft*fftOut)
}

func (n Node) solve2(nodes []Node, known map[string]struct{}, visitedDAC bool, visitedFFT bool) int {
	// Check for end of recursion
	if n.id == "out" {
		// If DAC and FFT were visited on path: count path
		if visitedDAC && visitedFFT {
			fmt.Println("############################")
			return 1
		}

		return 0
	}

	if _, exists := known[n.id]; exists {
		return 0 // Cycle detected (already on the current path)
	}

	known[n.id] = struct{}{}

	switch n.id {
	case "dac":
		visitedDAC = true
	case "fft":
		visitedFFT = true
	}

	paths := 0
	// Search Subtrees from here
	for _, cIdx := range n.childIdx {
		child := nodes[cIdx]

		paths += child.solve2(nodes, known, visitedDAC, visitedFFT)
	}

	delete(known, n.id)

	return paths
}

func countPaths(nodes []Node, known map[string]struct{}, start string, end string) int {
	for _, n := range nodes {
		if n.id == start {
			return n.dfs(nodes, known, end)
		}
	}
	return 0
}

func (n Node) dfs(nodes []Node, known map[string]struct{}, end string) int {
	// Check for end of recursion
	if n.id == end {
		return 1
	}

	if _, exists := known[n.id]; exists {
		return 0 // Cycle detected (already on the current path)
	}

	known[n.id] = struct{}{}

	paths := 0
	// Search Subtrees from here
	for _, cIdx := range n.childIdx {
		child := nodes[cIdx]

		paths += child.dfs(nodes, known, end)
	}

	delete(known, n.id)

	return paths
}
