package codewar


/*
Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

*/ 


func CountPositiveSumNegatives(numbers []int) []int {
	var res []int
	if len(numbers) == 0 {
		return res
	}
	count := 0; sum := 0
	for _, n := range numbers {
		if n < 0 {
			sum += n
		} else if n > 0 {
			count++
		}
	}
	res = append(res, count)
	res = append(res, sum)
	return res
}
