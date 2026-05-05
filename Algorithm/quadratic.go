package algo

/*
In this Kata you are expected to find the coefficients of quadratic equation of the given two roots (x1 and x2).

Equation will be the form of ax^2 + bx + c = 0

Return type is a Vector (tuple in Rust, Array in Ruby) containing coefficients of the equations in the order (a, b, c).

Since there are infinitely many solutions to this problem, we fix a = 1.

Remember, the roots can be written like (x-x1) * (x-x2) = 0

*/

func Quadratic(x1, x2 int) [3]int {
	var result [3]int
	result[0] = 1
	result[1] = - (x1 + x2)
	result[2] = -x1 * -x2
	return result
}
