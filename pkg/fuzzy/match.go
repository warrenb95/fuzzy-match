package fuzzy

import "math"

// Rank will return the Letchstein distance between the input and target.
func Rank(input, target string) int {
	input = " " + input
	target = " " + target
	return lDistanceRecursive([]rune(input), []rune(target), 1, 1)
}

func lDistanceRecursive(input, target []rune, i, j int) int {
	if len(input) == 0 {
		return len(target)
	}
	if len(target) == 0 {
		return len(input)
	}

	if len(input) < i || len(target) < j {
		return -1
	}

	return 0
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(input ...int) int {
	a := math.MaxInt
	for _, v := range input {
		a = int(math.Min(float64(a), float64(v)))
	}
	return a
}
