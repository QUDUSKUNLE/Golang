package arrayss

/*
Complete the function that takes two integers (a, b, where a < b) and return an array of all integers between the input parameters, including them.

For example:
*/

func Between(a, b int) (result []int) {
	result = make([]int, (b - a) + 1)
	result[0] = a
	i := 1
	for (i <= b - a) {
		result[i] = a + i
		i++
	}
	return
}
