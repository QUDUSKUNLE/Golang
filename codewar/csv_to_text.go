/*
Create a function that returns the CSV representation of a two-dimensional numeric array.

*/

package codewar

import (
	"fmt"
	"strings"
)

func ToCsvText(arr [][]int) string {
	var result string
	for _, v := range arr {
		for j, val := range v {
			if len(v) - 1 == j {
				result += fmt.Sprintf("%d\n", val)
				continue
			}
			result += fmt.Sprintf("%d,", val)
		}
	}
	return strings.TrimRight(result, "\n")
}

