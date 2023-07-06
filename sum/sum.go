package integers

func Add(x, y int) int {
	return x + y
}

// Sum calculates the total of a list of integers
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func ExampleAdd(a, b int) int {
	sum := Add(a, b)
	return sum
}
