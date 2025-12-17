package main

import (
	"os"
	"strings"
)

type Node struct {
	id       string
	idx      int
	childIdx []int
}

func parseInput() ([]Node, int) {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	var nodes []Node = make([]Node, 0)

	for pos, line := range lines {
		nodes = append(nodes, parseNode(strings.Split(line, " "), pos))
	}

	nodes = append(nodes, Node{
		"out",
		len(nodes),
		make([]int, 0),
	})

	for idx, node := range nodes[:len(nodes)-1] {
		segments := strings.Split(lines[idx], " ")
		for i, seg := range segments[1:] {
			for _, node2 := range nodes {
				if seg == node2.id {
					node.childIdx[i] = node2.idx
				}
			}
		}
	}

	return nodes, len(nodes)
}

func parseNode(segments []string, idx int) Node {
	return Node{
		segments[0][:len(segments[0])-1],
		idx,
		make([]int, len(segments)-1),
	}
}

func (item Node) compare(other Node) int {
	if item.id < other.id {
		return -1
	} else if item.id == other.id {
		return 0
	}
	return 1
}
