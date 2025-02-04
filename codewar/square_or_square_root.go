package codewar

import (
	"math"
)

func SquareOrSquareRoot(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		squareRoot := math.Sqrt(float64(v))
		if v / int(math.Ceil(squareRoot)*math.Ceil(squareRoot)) == 1 {
			result[i] = int(squareRoot)
		} else {
			result[i] = v*v
		}
	}
	return result
}
