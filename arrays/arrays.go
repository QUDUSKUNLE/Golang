package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return Sum(sums)
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}

// Array of 3 strings
var Arrayof3 = [...]string{"First", "Second", "Third"}

// Array Type
var A [5]string

// Slice Type
var B []string

// Map Type
var C map[string]string

const Size = 32;

type Person struct {
	name string
	age int
}

var D [5]string
