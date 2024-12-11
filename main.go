package main

import (
	"fmt"

	"github.com/QUDUSKUNLE/Golang/tutorial/codewar"
)

func Solve(s string, a, b int) string {
	if b > len(s) {
		return s
	}
	var result string
	between, diff := s[a:b], b-a
	for i, c := range s {
		if len(between)-i <= diff {
			if len(between)-i < 0 {
				result += string(c)
			} else {
				result += string(between[len(between)-i])
			}
		} else {
			result += string(c)
		}
	}
	return result
}

func main() {
	// fmt.Println(eetcode.GroupOpeningDays(eetcode.Data2))
	// fmt.Println(Solve("codewars", 1,5))
	// fmt.Println(len("Print"))
	// fmt.Println(codewars.Contamination("abcd", "z"))
	fmt.Println(codewar.Digitize(1234567))
	fmt.Println(codewar.MultiTable(3456))
	// fmt.Println(codewars.)
	// fmt.Println(codewars.)
	// strings(strings)
}
