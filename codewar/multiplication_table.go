package codewar

import "fmt"

/*
Your goal is to return multiplication table for number that is always an integer from 1 to 10.

For example, a multiplication table (string) for number == 5 looks like below:
*/ 

func MultiTable(n int) (result string) {
	for i := 1; i <= 10; i++ {
		if i == 10 {
			result += fmt.Sprintf("%v * %v = %v", i, n, i * n) 
		} else {
			result += fmt.Sprintf("%v * %v = %v", i, n, i * n) + "\n"
		}
	}
	return
}
