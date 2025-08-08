package arrayss

/*
Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.
*/

func Invert(arr []int) []int {
  result := make([]int, len(arr))
  for i, val := range arr {
		result[i] = val * -1
	}
	return result
}
