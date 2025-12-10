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
	start int
	end   int
}

func parseLine(line string) (Range, bool) {
	nums := strings.Split(line, "-")
	if len(nums) < 2 {
		return Range{
			start: 0,
			end:   0,
		}, true
	}

	start, _ := strconv.ParseInt(nums[0], 10, 64)
	end, _ := strconv.ParseInt(nums[1], 10, 64)

	return Range{
		start: int(start),
		end:   int(end),
	}, false
}

func part1(input string) {
	freshCount := 0

	lines := strings.Split(input, string('\n'))
	ranges := make([]Range, 0)

	for _, line := range lines {
		newRange, failed := parseLine(line)
		if !failed {
			ranges = append(ranges, newRange)
		} else {
			num, _ := strconv.ParseInt(line, 10, 64)
			for _, possibleRange := range ranges {
				if int(num) >= possibleRange.start && int(num) <= possibleRange.end {
					freshCount++
					// To avoid double counting if it is in multiple ranges
					break
				}
			}
		}
	}
	fmt.Println(freshCount)
}

func part2(input string) {
	// TODO: Fix

	lines := strings.Split(input, string('\n'))
	ranges := make([]Range, 0)

	for _, line := range lines {
		newRange, failed := parseLine(line)
		if !failed {
			ranges = append(ranges, newRange)
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].start == ranges[j].start {
			return ranges[i].end > ranges[j].end
		}
		return ranges[i].start < ranges[j].end
	})

	compInd := 0

	for i := 1; i < len(ranges); i++ {
		compRange := ranges[compInd]

		if compRange.contains(ranges[i]) {
			continue
		}

		// if they touch, we merge
		if compRange.overlaps(ranges[i]) {
			ranges[compInd] = compRange.append(ranges[i])
			continue
		}

		// disjoint ranges
		compInd++
		ranges[compInd] = ranges[i]
	}

	ranges = ranges[:compInd+1]

	sum := 0
	for _, r := range ranges {
		sum += r.Len()
	}

	fmt.Println(sum)
}

func (r1 Range) overlaps(r2 Range) bool {
	return r1.start <= r2.start && r1.end >= r2.start
}

func (r1 Range) contains(r2 Range) bool {
	return r1.start <= r2.start && r1.end >= r2.end
}

func (r1 Range) append(r2 Range) Range {
	return Range{r1.start, r2.end}
}

func (r Range) Len() int {
	return r.end - r.start + 1
}
