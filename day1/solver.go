package main

import (
	"fmt"
	"log"
	"os"
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

func part1(input string) {
	var pos int = 50
	var count int = 0

	lines := strings.Split(input, string('\n'))
	for _, line := range lines {
		shift, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			log.Fatal(err) // Sanity Check
		}

		if line[0] == 'L' {
			shift *= -1
		}

		pos += int(shift)
		pos %= 100

		if pos == 0 {
			count++
		}
	}

	fmt.Println(count)
}

func part2(input string) {
	var pos int = 50
	var count int = 0

	lines := strings.SplitSeq(input, string('\n'))
	for line := range lines {
		shift, err := strconv.ParseInt(line[1:], 10, 64)
		if err != nil {
			log.Fatal(err) // Sanity Check
		}

		// Count full loops first
		count += int(shift) / 100
		shift %= 100

		if line[0] == 'L' {
			shift *= -1
		}

		// Copy to check normal passes
		prev := pos

		pos += int(shift)

		switch {
		case pos == 0:
			count++
		case pos > 99:
			count++
			pos -= 100
		case pos < 0 && prev != 0:
			count++
			pos += 100
		case pos < 0 && prev == 0:
			pos += 100
		}

	}

	fmt.Println(count)
}
