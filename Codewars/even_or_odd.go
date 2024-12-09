package codewars

/*
Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
*/

func EvenOrOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}
