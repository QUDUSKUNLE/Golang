package algo

import "fmt"

/*
Given three integers a, b, and c, return the largest number obtained after inserting the operators +, *, and parentheses (). In other words, try every combination of a, b, and c with the operators, without reordering the operands, and return the maximum value.

Example
With the numbers 1, 2, and 3, here are some possible expressions:

1 * (2 + 3) = 5
1 * 2 * 3 = 6
1 + 2 * 3 = 7
(1 + 2) * 3 = 9
The maximum value that can be obtained is 9.
*/ 

func ExpressionMatter(a, b, c int) int {
	var result = [5]int{a * (b + c), a * b * c, a + b * c, (a + b) * c, a + b + c}
	var max = 0
	for i, v := range result {
		fmt.Println(i, v)
		if v > max {
			max = v
		}
	}
	return max
}
