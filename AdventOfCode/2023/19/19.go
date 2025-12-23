package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/projectroot"
)

type part struct {
	x int
	m int
	a int
	s int
}

type check struct {
	category 	byte
	condition 	byte
	value 		int
	action		string
}

func (p part) getVal(b byte) int {
	switch b {
	case 'x': return p.x
	case 'm': return p.m
	case 'a': return p.a
	case 's': return p.s
	default: return 0
	}
}

func toParts(inp string) []part {
	rawParts := strings.Split(inp, "\n")

	parts := []part{}
	for _, line := range rawParts {
		vals := strings.Split(line[1:len(line)-1], ",")
		x, _ := strconv.Atoi(vals[0][2:])
		m, _ := strconv.Atoi(vals[1][2:])
		a, _ := strconv.Atoi(vals[2][2:])
		s, _ := strconv.Atoi(vals[3][2:])
		part := part{x, m, a, s}
		parts = append(parts, part)
	}

	return parts
}

func toWorkflow(inp string) map[string][]check {
	rawWorkflow := strings.Split(inp, "\n")

	workflow := make(map[string][]check)
	for _, line := range rawWorkflow {
		start := strings.Index(line, "{")

		id := line[:start]
		workflow[id] = []check{}

		checks := strings.Split(line[start+1:len(line)-1], ",")
		for _, c := range checks[:len(checks)-1] {
			colon := strings.Index(c, ":")
			val, _ := strconv.Atoi(c[2:colon])

			check := check {
				c[0],
				c[1],
				val,
				c[colon+1:],
			}

			workflow[id] = append(workflow[id], check)
		}
		workflow[id] = append(workflow[id], check{'\x00', '\x00', 0, checks[len(checks)-1]})
	}

	return workflow
}

func checkPart(part part, entry string, workflows map[string][]check) bool {
	if entry == "A" {
		return true
	}
	if entry == "R" {
		return false
	}

	checks := workflows[entry]
	for _, c := range checks {
		switch c.condition {
		case '<':
			if part.getVal(c.category) < c.value {
				return checkPart(part, c.action, workflows)
			}
		case '>':
			if part.getVal(c.category) > c.value {
				return checkPart(part, c.action, workflows)
			}
		default:
			return checkPart(part, c.action, workflows)
		}
	}

	panic("reached where it shouldnt have")
}

func p1(filename string) {
	f, _ := os.ReadFile(filename)
	inp := strings.Split(strings.Trim(string(f), "\n "), "\n\n")

	workflows := toWorkflow(inp[0])
	parts := toParts(inp[1])

	var ans int
	for _, part := range parts {
		if checkPart(part, "in", workflows) {
			ans += part.x + part.m + part.a + part.s
		}
	}
	fmt.Println(ans)
}

func count(ranges map[byte][]int, name string, workflows map[string][]check) int {
	if name == "R" {
		return 0
	}

	if name == "A" {
		product := 1
		for _, val := range ranges {
			product *= val[1] - (val[0] - 1)
		}
		return product
	}

	total := 0

	checks := workflows[name]
	for _, c := range checks[:len(checks)-1] {
		lo, hi := ranges[c.category][0], ranges[c.category][1]
		t, f := []int{}, []int{}
		switch c.condition {
		case '<':
			t = []int{lo, c.value-1}
			f = []int{c.value, hi}
		case '>':
			t = []int{c.value+1, hi}
			f = []int{lo, c.value}
		default:
		}


		if len(t) > 0 {
			newRanges := make(map[byte][]int)
			for k, v := range ranges {
				newRanges[k] = v
			}
			newRanges[c.category] = t
			total += count(newRanges, c.action, workflows)
		}

		if len(f) > 0 {
			ranges[c.category] = f
		} else {
			return total
		}
	}
	total += count(ranges, checks[len(checks)-1].action, workflows)

	return total
}

func p2(filename string) {
	f, _ := os.ReadFile(filename)
	workflows := toWorkflow(strings.Split(strings.Trim(string(f), "\n "), "\n\n")[0])

	fmt.Println(count(map[byte][]int{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	}, "in", workflows))
}

func main() {
	input := filepath.Join(projectroot.Path, "example.txt")
	input = filepath.Join(projectroot.Path, "input.txt")
	p1(input)
	p2(input)
}
