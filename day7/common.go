package main

import (
	"os"
	"strings"
)

func parseInput() [][]string {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	var grid [][]string = make([][]string, 0)
	for line := range strings.SplitSeq(string(content), "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}
