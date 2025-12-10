package main

import (
	"os"
	"strconv"
	"strings"
)

func parseInput() [][]int {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	var numbers [][]int = make([][]int, 0)
	for line := range strings.SplitSeq(string(content), "\n") {
		numbers = append(numbers, getNumbers(line))
	}

	return numbers
}

func getNumbers(line string) []int {
	var numbers []int = make([]int, 0)
	for s := range strings.SplitSeq(line, "") {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	return numbers
}
