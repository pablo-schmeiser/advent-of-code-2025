package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start     int
	end       int
	numDigits int
}

func parseInput() []Range {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var ranges []Range
	for line := range strings.SplitSeq(scanner.Text(), ",") {
		limits := strings.Split(line, "-")
		start, _ := strconv.Atoi(limits[0])
		end, _ := strconv.Atoi(limits[1])
		startDigits, endDigits := digits(start), digits(end)

		if startDigits == endDigits {
			ranges = append(ranges, Range{start, end, startDigits})
		} else {
			ranges = append(ranges, Range{start, pow10(endDigits-1) - 1, startDigits})
			ranges = append(ranges, Range{pow10(startDigits), end, endDigits})
		}
	}

	return ranges
}

func digits(i int) int {
	return int(len(strconv.Itoa(i)))
}

func (r Range) InvalidIDs(chunkLen int) map[int]struct{} {
	repeats := r.numDigits / chunkLen
	invalids := make(map[int]struct{})

	for i := r.start / pow10(r.numDigits-chunkLen); i <= r.end/pow10(r.numDigits-chunkLen)+1; i++ {
		id := buildID(i, repeats)
		if id < r.start {
			continue
		}
		if id > r.end {
			break
		}
		invalids[id] = struct{}{}
	}

	return invalids
}

func buildID(chunk, repeat int) int {
	n := chunk
	for i := 1; i < repeat; i++ {
		n = n*pow10(digits(chunk)) + chunk
	}
	return n
}

func pow10(exp int) int {
	return int(math.Pow10(exp))
}
