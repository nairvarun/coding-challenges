package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/set"
)

type Beam struct {
	loc coordinate
	dir coordinate
}

type coordinate struct {
	row int
	col int
}

func calc(start Beam, grid []string) int {
	seen := make(set.Set[Beam])
	energized := make(set.Set[coordinate])

	q := []Beam{start}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		curr.loc.row += curr.dir.row
		curr.loc.col += curr.dir.col

		if curr.loc.row < 0 || curr.loc.row >= len(grid) ||
		curr.loc.col < 0 || curr.loc.col >= len(grid[0]) {
			continue
		}

		switch grid[curr.loc.row][curr.loc.col] {
		case '\\':
			curr.dir.row, curr.dir.col = curr.dir.col, curr.dir.row

		case '/':
			curr.dir.row, curr.dir.col = -curr.dir.col, -curr.dir.row

		case '-':
			if curr.dir.row == 0 {
				break
			}

			opp := curr
			opp.dir.row, opp.dir.col = 0, -1
			if !seen.Contains(opp) {
				seen.Add(opp)
				energized.Add(opp.loc)
				q = append(q, opp)
			}

			curr.dir.row, curr.dir.col = 0, 1

		case '|':
			if curr.dir.col == 0 {
				break
			}

			opp := curr
			opp.dir.row, opp.dir.col = -1, 0
			if !seen.Contains(opp) {
				seen.Add(opp)
				energized.Add(opp.loc)
				q = append(q, opp)
			}

			curr.dir.row, curr.dir.col = 1, 0
		}

		if !seen.Contains(curr) {
			seen.Add(curr)
			q = append(q, curr)
			energized.Add(curr.loc)
		}
	}

	return len(energized)
}

func p1(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	grid := strings.Split(strings.Trim(string(f), "\n"), "\n")

	res := calc(Beam{
		loc: coordinate{row: 0, col: -1},
		dir: coordinate{row: 0, col: 1},
	}, grid)

	fmt.Println(res)
}

func p2(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	grid := strings.Split(strings.Trim(string(f), "\n"), "\n")

	res := 0
	R, C := len(grid), len(grid[0])

	// top & bottom
	for c := 0; c < C; c++ {
		res = max(res, calc(Beam{
			loc: coordinate{row: 0, col: c},
			dir: coordinate{row: 1, col: 0},
		}, grid))

		res = max(res, calc(Beam{
			loc: coordinate{row: R, col: c},
			dir: coordinate{row: -1, col: 0},
		}, grid))
	}

	// left & right
	for r :=  0; r < R; r++ {
		res = max(res, calc(Beam{
			loc: coordinate{row: r, col: 0},
			dir: coordinate{row: 0, col: 1},
		}, grid))

		res = max(res, calc(Beam{
			loc: coordinate{row: r, col: C},
			dir: coordinate{row: 0, col: -1},
		}, grid))
	}
	fmt.Println(res)
}

func main() {
	input := "../example.txt"
	input = "../input.txt"
	p1(input)
	p2(input)
}
