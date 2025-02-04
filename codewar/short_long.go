package codewar

import (
	"fmt"
)

/*
Given 2 strings, a and b, return a string of the form short+long+short, with the shorter string on the outside and the longer string on the inside. The strings will not be the same length, but they may be empty ( zero length ).
*/

func Solution(a, b string) string {
	if len(a) > len(b) {
		return fmt.Sprintf("%s%s%s", b, a, b)
	}
	return fmt.Sprintf("%s%s%s", a, b, a)
}
