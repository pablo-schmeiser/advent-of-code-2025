package main

import (
	"fmt"
	"math"
)

const START_STATE uint64 = 0

func main() {
	machines, numLines := parseInput()

	part1(machines, numLines)
	part2(machines, numLines)
}

func part1(machines []Machine, numLines int) {
	sum := 0

	for _, m := range machines {
		presses, success := solvingButtons(START_STATE, m.buttons, m.indicators, math.MaxInt)

		if !success {
			panic("Critical Error: Not solvable")
		}

		sum += presses
	}

	fmt.Println(sum)
}

func solvingButtons(state uint64, buttons []uint64, targetState uint64, maxPresses int) (int, bool) {
	if state == targetState {
		return 0, true
	} else if maxPresses < 1 {
		return 0, false
	}

	success := false

	for i := range buttons {
		presses, solved := solvingButtons(state^buttons[i], buttons[i+1:], targetState, maxPresses-1)
		if solved {
			success = true
			presses++

			if presses < maxPresses {
				maxPresses = presses
			}
		}
	}

	return maxPresses, success
}

func part2(machmachines []Machine, numLines int) {
	// TODO
}
