package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(string(content))
	part2(string(content))
}

type Range struct {
	start int64
	end   int64
}

func parseLine(line string) Range {
	nums := strings.Split(line, "-")

	start, _ := strconv.ParseInt(nums[0], 10, 64)
	end, _ := strconv.ParseInt(nums[1], 10, 64)

	return Range{
		start: start,
		end:   end,
	}
}

func part1(input string) {
	var freshCount int64 = 0

	lines := strings.Split(input, string('\n'))
	ranges := make([]Range, 0)

	isIDs := false
	for _, line := range lines {
		if line == "" {
			isIDs = true
			continue
		} else if !isIDs {
			ranges = append(ranges, parseLine(line))
		} else {
			num, _ := strconv.ParseInt(line, 10, 64)
			// Check if num is in any of the ranges
			for _, possibleRange := range ranges {
				if num >= possibleRange.start && num <= possibleRange.end {
					// Num is fresh
					freshCount++
					break
				}
			}
		}
	}

	fmt.Println(freshCount)
}

func part2(input string) {
	var freshCount int64 = 0

	lines := strings.Split(input, string('\n'))
	ranges := make([]Range, 0)

	for _, line := range lines {
		if line == "" {
			break
		}

		ranges = append(ranges, parseLine(line))
	}

	// Merge overlapping ranges
	merged := make([]Range, 0)
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	current := ranges[0]
	for _, next := range ranges[1:] {
		if next.start <= current.end+1 {
			if next.end > current.end {
				current.end = next.end
			}
			// Otherwise current fully covers next, do nothing
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	// Count total numbers in ranges
	for _, r := range merged {
		freshCount += r.end - r.start + 1
	}

	fmt.Println(freshCount)
}
