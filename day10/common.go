package main

import (
	"os"
	"strconv"
	"strings"
)

type Machine struct {
	indicators uint64
	buttons    []uint64
	altButtons [][]int
	joltages   []int
}

func parseInput() ([]Machine, int) {
	//scanner := bufio.NewScanner(os.Stdin)
	content, _ := os.ReadFile("input.txt")
	//scanner.Scan() scanner.Text()

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	numLines := len(lines)

	var machines []Machine = make([]Machine, numLines)

	for pos, line := range lines {
		machines[pos] = parseMachine(strings.Split(line, " "))
	}

	return machines, numLines
}

func parseMachine(segments []string) Machine {
	buttons, altButtons := parseButtons(segments[1 : len(segments)-1])
	return Machine{
		parseIndicators(segments[0][1 : len(segments[0])-1]),
		buttons,
		altButtons,
		parseJoltages(segments[len(segments)-1]),
	}

}

func parseIndicators(indicators string) uint64 {
	result := uint64(0)

	runes := []rune(indicators)
	for i := len(runes) - 1; i >= 0; i-- {
		result <<= 1

		switch runes[i] {
		case '#':
			result |= 1

		case '.':
			// no-op
		}
	}

	return result
}

func parseButtons(buttonStrings []string) ([]uint64, [][]int) {
	buttons := make([]uint64, len(buttonStrings))
	altButtons := make([][]int, len(buttonStrings))

	for i, buttonString := range buttonStrings {
		result := uint64(0)

		substrings := strings.Split(buttonString[1:len(buttonString)-1], ",")
		a := make([]int, len(substrings))
		for j, v := range substrings {
			idx, _ := strconv.Atoi(v)

			result |= 1 << uint64(idx)
			a[j] = idx
		}

		buttons[i] = result
		altButtons[i] = a
	}

	return buttons, altButtons
}

func parseJoltages(joltageSpec string) []int {
	joltages := strings.Split(joltageSpec[1:len(joltageSpec)-1], ",")
	result := make([]int, len(joltages))
	for i, joltage := range joltages {
		n, _ := strconv.Atoi(joltage)

		result[i] = n
	}

	return result
}
