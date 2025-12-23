package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/input"
	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/math"
)

const Example = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func part1(moves []string) {
	n := 50
	res := 0

	for _, move := range moves {
		direction := move[0]
		turnsStr := move[1:]
		turns, err := strconv.Atoi(turnsStr)
		if err != nil {
			panic(err)
		}

		if direction == 'R' {
			n = (n + turns) % 100
		} else {
			n = math.PositiveMod(n - turns, 100)
		}

		if n == 0 {
			res += 1
		}
	}
	fmt.Println(res)
}

func part2(moves []string) {
	n := 50
	res := 0
	
	for _, move := range moves {
		direction := move[0]
		turnsStr := move[1:]
		turns, err := strconv.Atoi(turnsStr)
		if err != nil {
			panic(err)
		}
		res += turns / 100
		turns %= 100

		if direction == 'R' {
			if n + turns >= 100 {
				res += 1
			}
			n = (n + turns) % 100
		} else {
			if n > 0 && turns >= n {
				res += 1
			}
			n = math.PositiveMod(n - turns, 100)
		}
	}
	fmt.Println(res)
}

func main() {
	part1(strings.Fields(Example))
	part1(strings.Fields(input.Fetch(2025, 1)))
	part2(strings.Fields(Example))
	part2(strings.Fields(input.Fetch(2025, 1)))
}