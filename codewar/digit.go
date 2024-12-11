package codewar

import (
	"strconv"
)

/*
Convert number to reversed array of digits
Given a random non-negative number, you have to return the digits of this number within an array in reverse order.
*/

func Digitize(n int) []int {
	str := strconv.Itoa(n)
	res := make([]int, 0, len(str))
	for i := len(str) - 1; i >= 0; i-- {
		in, _ := strconv.Atoi(str[i : i+1])
		res = append(res, in)
	}
	return res
}
