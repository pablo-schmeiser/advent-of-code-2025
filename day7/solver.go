package main

import (
	"fmt"
)

const SIZE = 142

func main() {
	numbers := parseInput()
	part1(numbers)
	part2(numbers)
}

func part1(chars [][]string) {
	count := 0

	for x, row := range chars {
		if x == len(chars)-1 {
			break
		}

		for y, c := range row {
			if c == "|" && chars[x+1][y] == "^" {
				chars[x+1][y-1] = "|"
				chars[x+1][y+1] = "|"
				count++
			} else if c == "|" {
				chars[x+1][y] = "|"
			} else if c == "S" {
				chars[x+1][y] = "|"
				break
			}
		}
	}

	fmt.Println(count)
}

func part2(chars [][]string) {
	journeys := [SIZE][SIZE]int{}

	for x, row := range chars {
		if x == len(chars)-1 {
			break
		}

		for y, c := range row {

			if c == "|" && chars[x+1][y] == "^" {
				chars[x+1][y-1] = "|"
				journeys[x+1][y-1] += journeys[x][y]
				chars[x+1][y+1] = "|"
				journeys[x+1][y+1] += journeys[x][y]
			} else if c == "|" {
				chars[x+1][y] = "|"
				journeys[x+1][y] += journeys[x][y]
			} else if c == "S" {
				chars[x+1][y] = "|"
				journeys[x+1][y] += 1
				break
			}
		}
	}

	paths := 0
	for _, universes := range journeys[SIZE-1] {
		paths += universes
	}

	fmt.Println(paths)
}
