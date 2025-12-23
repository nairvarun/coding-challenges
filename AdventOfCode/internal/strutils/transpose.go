package strutils

import (
	"slices"
	"strings"
)

func Transpose(slice []string) []string {
	NUM_ROWS := len(slice[0])
	NUM_COLS := len(slice)
	transposed_slice := make([]string, NUM_ROWS)

	for r := range NUM_ROWS {
		temp := make([]string, NUM_COLS)
		for c := range NUM_COLS {
			temp[c] = string(slice[c][r])
		}
		transposed_slice[r] = strings.Join(temp, "")
	}

	return transposed_slice
}

func Rotate(slice []string) []string {
	slices.Reverse(slice)
	return Transpose(slice)
}
