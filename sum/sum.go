package integers

func Add(x, y int) int {
	return x + y
}

// Sum takes a slice of integers and returns their total
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
