package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/input"
)

const Example = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

type Range struct {
	start int
	end   int
}

func part1(ranges []Range) {
	res := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			num := strconv.Itoa(i)
			if len(num) % 2 != 0 {
				continue
			}
			mid := len(num) / 2
			if num[:mid] == num[mid:] {
				res += i
			}
		}
	}
	fmt.Println(res)
}

func part2(ranges []Range) {
	res := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			num := strconv.Itoa(i)
			for j := 1; j <= len(num) / 2; j++ {
				if num == strings.Repeat(num[:j], len(num) / j) {
					res += i
					break
				}
			}
		}
	}
	fmt.Println(res)
}

func processInput(input string) []Range {
	var ranges []Range

	rangesStr := strings.Split(input, ",")
	for _, r := range rangesStr {
		parts := strings.Split(r, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		ranges = append(ranges, Range{start, end})
	}
	return ranges
}

func main() {
	part1(processInput(Example))
	part1(processInput(input.Fetch(2025, 2)))
	part2(processInput(Example))
	part2(processInput(input.Fetch(2025, 2)))
}