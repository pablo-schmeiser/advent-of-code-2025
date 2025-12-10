package main

import (
	"fmt"
	"maps"
)

func main() {
	ranges := parseInput()
	part1(ranges)
	part2(ranges)
}

func part1(ranges []Range) {
	sum := 0
	for _, r := range ranges {
		// Skip ranges with uneven digit counts
		if r.numDigits%2 == 1 {
			continue
		}

		for n := range r.InvalidIDs(r.numDigits / 2) {
			sum += n
		}
	}

	fmt.Println(sum)
}

func part2(ranges []Range) {
	sum := 0
	invalids := map[int]struct{}{}

	for _, r := range ranges {
		for chunkLen := 1; chunkLen <= r.numDigits/2; chunkLen++ {

			// filter for full repetitions
			if r.numDigits%chunkLen != 0 {
				continue
			}

			maps.Copy(invalids, r.InvalidIDs(chunkLen))
		}
	}

	for n := range invalids {
		sum += n
	}
	fmt.Println(sum)
}
