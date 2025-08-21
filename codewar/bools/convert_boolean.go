package bools

/*
Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.

*/

func ConvertBoolean(b bool) string {
	return map[bool]string{true: "Yes", false: "No"}[b]
}
