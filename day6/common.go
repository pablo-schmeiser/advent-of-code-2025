package main

import (
	"os"
	"strconv"
	"strings"
)

const NUM_ROW_COUNT = 4
const MAX_NUM_LENGTH = 10

func parseInput() ([][]int, []bool) {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	var nums [][]int = make([][]int, 0)
	var opIsAdd []bool = make([]bool, 0)
	for i, line := range strings.Split(string(content), "\n") {
		if i < NUM_ROW_COUNT {
			parsedNums := make([]int, 0)
			for s := range strings.SplitSeq(line, " ") {
				n, e := strconv.Atoi(s)

				if e != nil {
					continue
				}

				parsedNums = append(parsedNums, n)
			}
			nums = append(nums, parsedNums)
		} else {
			for _, s := range strings.Split(line, " ") {
				switch s {
				case "+":
					opIsAdd = append(opIsAdd, true)
				case "*":
					opIsAdd = append(opIsAdd, false)
				default:
					continue
				}
			}
		}
	}

	return nums, opIsAdd
}

func parseInput2() ([][][]int, []bool) {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	var nums [][][]int = make([][][]int, 0)
	var opIsAdd []bool = make([]bool, 0)
	for i, line := range strings.Split(string(content), "\n") {
		if i < NUM_ROW_COUNT {
			parsedCells := make([][]int, 0)
			for s := range strings.SplitSeq(line, " ") {
				_, e := strconv.Atoi(s)

				if e != nil {
					continue
				}

				parsedNums := make([]int, MAX_NUM_LENGTH)
				chars := strings.Split(s, "")
				for i := MAX_NUM_LENGTH - 1; i >= 0; i-- {
					if len(chars)-1 < i {
						parsedNums[MAX_NUM_LENGTH-1-i] = -1
					} else {
						n, _ := strconv.Atoi(chars[i])
						parsedNums[MAX_NUM_LENGTH-1-i] = n
					}
				}
				parsedCells = append(parsedCells, parsedNums)
			}
			nums = append(nums, parsedCells)
		} else {
			for s := range strings.SplitSeq(line, " ") {
				switch s {
				case "+":
					opIsAdd = append(opIsAdd, true)
				case "*":
					opIsAdd = append(opIsAdd, false)
				default:
					continue
				}
			}
		}
	}

	return nums, opIsAdd
}
