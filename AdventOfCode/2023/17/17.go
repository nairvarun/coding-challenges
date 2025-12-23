package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/set"
)

type coordinate struct {
	row int
	col int
}

type block struct {
	heatLoss int
	loc      coordinate
	dir      coordinate
	dirCount int
}

type seenBlock struct {
	loc      coordinate
	dir      coordinate
	dirCount int
}

type priorityQueue []*block

func (pq priorityQueue) Len() int 			{ return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].heatLoss < pq[j].heatLoss }
func (pq priorityQueue) Swap(i, j int) 		{ pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x any) 		{ *pq = append(*pq, x.(*block)) }
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	blk := old[n-1]
	*pq = old[:n-1]
	return blk
}


func p1(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil { panic(err) }

	inp := strings.Split(strings.Trim(string(f), "\n "), "\n")
	grid := make([][]int, len(inp))
	for row_num, row := range inp {
		for _, ch := range row {
			val, err := strconv.Atoi(string(ch))
			if err != nil { panic(err) }

			grid[row_num] = append(grid[row_num], val)
		}
	}
	// fmt.Println(grid)

	pq := priorityQueue{{
		heatLoss: 0,
		loc:      coordinate{row: 0, col: 0},
		dir:      coordinate{row: 0, col: 0},
		dirCount: 0,
	}}

	heap.Init(&pq)
	seen := make(set.Set[seenBlock])

	for len(pq) > 0 {
		blk := heap.Pop(&pq).(*block)

		if blk.loc.row == len(grid)-1 && blk.loc.col == len(grid[0])-1 {
			fmt.Println(blk.heatLoss)
			break
		}

		if seen.Contains(seenBlock{blk.loc, blk.dir, blk.dirCount}) {
			continue
		}
		seen.Add(seenBlock{blk.loc, blk.dir, blk.dirCount})

		for _, dir := range []coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			if dir == (coordinate{-blk.dir.row, -blk.dir.col}) {
				continue
			}

			if dir == blk.dir && blk.dirCount >= 3 {
				continue
			}

			nextBlk := *blk

			if dir == nextBlk.dir && nextBlk.dirCount < 3 {
				nextBlk.dirCount++
			} else {
				nextBlk.dirCount = 1
			}
			nextBlk.dir.row = dir.row
			nextBlk.dir.col = dir.col
			nextBlk.loc.row += nextBlk.dir.row
			nextBlk.loc.col += nextBlk.dir.col

			if nextBlk.loc.row >= 0 && nextBlk.loc.row < len(grid) &&
			nextBlk.loc.col >= 0 && nextBlk.loc.col < len(grid[0]) {
				nextBlk.heatLoss += grid[nextBlk.loc.row][nextBlk.loc.col]
				heap.Push(&pq, &nextBlk)
			}
		}
	}
}

func p2(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inp := strings.Split(strings.Trim(string(f), "\n "), "\n")
	grid := make([][]int, len(inp))
	for row_num, row := range inp {
		for _, ch := range row {
			val, err := strconv.Atoi(string(ch))
			if err != nil {
				panic(err)
			}
			grid[row_num] = append(grid[row_num], val)
		}
	}

	pq := priorityQueue{{
		heatLoss: 0,
		loc: coordinate{row: 0, col: 0},
		dir: coordinate{row: 0, col: 0},
		dirCount: 0,
	}}

	heap.Init(&pq)
	seen := make(set.Set[seenBlock])

	for len(pq) > 0 {
		blk := heap.Pop(&pq).(*block)
		if blk.loc == (coordinate{row: len(grid)-1, col: len(grid[0])-1}) && blk.dirCount >= 4 {
			fmt.Println(blk.heatLoss)
			break
		}

		s := seenBlock{blk.loc, blk.dir, blk.dirCount}
		if seen.Contains(s) {
			continue
		}
		seen.Add(s)

		for _, dir := range []coordinate{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			if dir == (coordinate{-blk.dir.row, -blk.dir.col}) {
				continue
			}

			if dir == blk.dir && blk.dirCount >= 10 {
				continue
			}

			if blk.dir != (coordinate{0, 0}) && dir != blk.dir && blk.dirCount < 4 {
				continue
			}

			nextBlk := *blk

			if dir == nextBlk.dir {
				nextBlk.dirCount++
			} else {
				nextBlk.dirCount = 1
			}

			nextBlk.dir.row = dir.row
			nextBlk.dir.col = dir.col
			nextBlk.loc.row += nextBlk.dir.row
			nextBlk.loc.col += nextBlk.dir.col

			if nextBlk.loc.row >= 0 && nextBlk.loc.row < len(grid) &&
			nextBlk.loc.col >= 0 && nextBlk.loc.col < len(grid[0]) {
				nextBlk.heatLoss += grid[nextBlk.loc.row][nextBlk.loc.col]
				heap.Push(&pq, &nextBlk)
			}
		}
	}
}

func main() {
	input := "../example.txt"
	input = "../input.txt"
	p1(input)
	p2(input)
}
