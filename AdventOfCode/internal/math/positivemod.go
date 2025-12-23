package math

func PositiveMod(a, b int) int {
	mod := a % b
	if mod < 0 {
		return mod + b
	}
	return mod
}