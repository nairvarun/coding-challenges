package main

import (
	"fmt"
	"strings"
	"math"

	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/input"
	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/strutils"
)

const Example = `987654321111111
811111111111119
234234234234278
818181911112111`

func part1(batteryBanks []string) {
	var res int
	for _, batteries := range batteryBanks {
		var d1, d2 int
		for i, batteryRune := range []rune(batteries) {
			battery := strutils.Rtoi(batteryRune)
			if battery > d1 && i != len(batteries) - 1 {
				d1 = battery
				d2 = 0
			} else if battery > d2 {
				d2 = battery
			}
		}
		res += d1 * 10 + d2
	}
	fmt.Println(res)
}

func part2(batteryBanks []string) {
	var res int
	l := 0
	r := len(batteryBanks[0]) - 11
	for _, battries := range batteryBanks {
		batteryRunes := []rune(battries)
		for n := range 12 {
			mx := 0
			for i := l; i < r; i++ {
				b := strutils.Rtoi(batteryRunes[i])
				if b > mx {
					mx = b
				}
			}
			res += mx * int(math.Pow10(n))
		}
	}
	fmt.Println(res)
}

func main() {
	part1(strings.Fields(Example))
	part1(strings.Fields(input.Fetch(2025, 3)))
	part2(strings.Fields(Example))
	part2(strings.Fields(input.Fetch(2025, 3)))
}