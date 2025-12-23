package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/nairvarun/coding-challenges/AdventOfCode/internal/strutils"
)

func p1(file string) {
	f, _ := os.ReadFile(file)
	platform := strings.Split(string(f), "\n")
	platform = platform[:len(platform)-1]

	R, C := len(platform), len(platform[0])

	blocked := make(map[int]int)
	for c := 0; c < C; c += 1 {
		blocked[c] = -1
	}

	res := 0

	for r := 0; r < R; r += 1 {
		for c := 0; c < C; c += 1 {
			switch platform[r][c] {
				case '#':
					blocked[c] = r
				case 'O':
					beam := blocked[c]
					res += R - (beam + 1)
					blocked[c] = (beam + 1)
			}
		}
	}
	fmt.Println(res)
}

func p2(file string) {
	f, _ := os.ReadFile(file)
	platform := strings.Split(string(f), "\n")
	platform = platform[:len(platform)-1]

	type void struct{}
	var member void
	seen := map[string]void{
		to_key(platform): member,
	}
	platforms := []string{to_key(platform)}

	N := 1000000000
	i := 1
	for ; i < N; i++ {
		platform = cycle(platform)
		key := to_key(platform)
		_, exists := seen[key]
		if exists {
			break
		}
		seen[key] = member
		platforms = append(platforms, key)
	}
	first := slices.Index(platforms, to_key(platform))
	platform = strings.Split(platforms[(N - first) % (i - first) + first], ";")

	res := 0
	for idx, r := range platform {
		for _, ch := range r {
			if ch == 'O' {
				res += len(platform) - idx
			}
		}
	}
	fmt.Println(res)
}

func to_key(platform []string) string {
	return strings.Join(platform, ";")
}

func cycle(platform []string) []string {
	for i := 0; i < 4; i++ {
		platform = strutils.Rotate(tilt(platform))
	}
	return platform
}

func tilt(platform []string) []string {
	for idx, r := range strutils.Transpose(platform) {
		substr := strings.Split(r, "#")
		for idx, s := range substr {
			s := []rune(s)
			slices.Sort(s)
			substr[idx] = strutils.Reversed(string(s))
		}

		platform[idx] = strings.Join(substr, "#")
	}
	return strutils.Transpose(platform)
}

func main() {
	// input := "example.txt"
	input := "input.txt"
	// p1(input)
	p2(input)
}
