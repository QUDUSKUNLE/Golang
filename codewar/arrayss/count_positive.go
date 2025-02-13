package arrayss

/*
Given an array of integers.

Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.

If the input is an empty array or is null, return an empty array.

*/

func CountPositiveSumNegatives(numbers []int) []int {
	res := []int{0,0}
	if len(numbers) == 0 {
		return res
	}
	for _, n := range numbers {
		if n < 0 {
			res[1] += n
		} else if n > 0 {
			res[0]++
		}
	}
	return res
}
