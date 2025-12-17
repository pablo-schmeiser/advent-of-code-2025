package main

import (
	"fmt"
	"math"
)

const NUM_DIGITS_PART2 = 12

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
	var sumMaxOutput int64 = 0
	for _, bank := range numbers {
		// Iterate over each digit, needed
		for i := range NUM_DIGITS_PART2 {
			if len(bank) == 0 {
				break
			}

			// Only check arrays that still have enough digits left
			highestNumPos := findHighestIndex(bank[:len(bank)-NUM_DIGITS_PART2+i+1])
			if highestNumPos <= len(bank)-NUM_DIGITS_PART2+i {
				// Trim the bank to only the part after the highest number found and add the number with correct place value to the sum
				sumMaxOutput += int64(math.Pow10(NUM_DIGITS_PART2-i-1)) * int64(bank[highestNumPos])
				bank = bank[highestNumPos+1:]
			} else {
				break
			}

			if i >= NUM_DIGITS_PART2-1 {
				// If we reached the last digit empty the bank
				bank = nil
			}
		}

		if len(bank) == 0 {
			continue
		} else {
			// If there are still numbers left, we need to add the last digits
			for i := len(bank) - 1; i >= 0; i-- {
				sumMaxOutput += int64(math.Pow10(len(bank)-1-i)) * int64(bank[i])
			}
		}
	}

	fmt.Println(sumMaxOutput)
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
