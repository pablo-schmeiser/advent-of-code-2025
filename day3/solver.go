package main

import (
	"fmt"
)

func main() {
	numbers := parseInput()
	part1(numbers)
	part2(numbers)
}

func part1(numbers [][]int) {
	sumMaxOutput := 0
	for _, bank := range numbers {
		highestNumPos := findHighestIndex(bank)
		num := 0
		if highestNumPos == len(bank)-1 {
			num = 10*bank[findHighestIndex(bank[0:highestNumPos])] + bank[highestNumPos]
		} else {
			num = 10*bank[highestNumPos] + bank[findHighestIndex(bank[(highestNumPos+1):])+highestNumPos+1]
		}
		sumMaxOutput += num
	}

	fmt.Println(sumMaxOutput)
}

func part2(numbers [][]int) {
	// TODO
}

func findHighestIndex(array []int) int {
	var currentHighestNum int = -1
	var currentHighestIndex int = -1
	for i, e := range array {
		if e > currentHighestNum {
			currentHighestNum = e
			currentHighestIndex = i
		}
	}

	return currentHighestIndex
}
