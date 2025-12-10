package main

import (
	"fmt"
	"math"
)

func main() {
	numbers, opIsAdd := parseInput()
	part1(numbers, opIsAdd)
	numbers2, opIsAdd2 := parseInput2()
	part2(numbers2, opIsAdd2)
}

func part1(nums [][]int, opIsAdd []bool) {
	if len(nums[0]) != len(opIsAdd) {
		fmt.Println("Error: Length mismatch")
		return
	}

	sum := 0
	for i, isAdd := range opIsAdd {
		numToAdd := 0
		if isAdd {
			for j := range NUM_ROW_COUNT {
				numToAdd += nums[j][i]
			}
		} else {
			numToAdd = 1
			for j := range NUM_ROW_COUNT {
				numToAdd *= nums[j][i]
			}
		}
		sum += numToAdd
	}

	fmt.Println(sum)
}

func part2(chars [][][]int, opIsAdd []bool) {
	if len(chars[0]) != len(opIsAdd) {
		fmt.Println("Error: Length mismatch")
		return
	}

	sum := 0
	for col, isAdd := range opIsAdd {
		nums := make([]int, MAX_NUM_LENGTH)
		for charCount := range MAX_NUM_LENGTH {
			num := 0
			for row := range NUM_ROW_COUNT {
				if chars[row][col][charCount] < 0 {
					continue
				}

				num += int(math.Pow10(NUM_ROW_COUNT-1-row)) * chars[row][col][charCount]
			}
			nums[charCount] = num
		}

		numToAdd := 0
		if isAdd {
			for i := range MAX_NUM_LENGTH {
				numToAdd += nums[i]
			}
		} else {
			numToAdd = 1
			for i := range MAX_NUM_LENGTH {
				numToAdd *= nums[i]
			}
		}

		sum += numToAdd
	}

	fmt.Println(sum)
}
