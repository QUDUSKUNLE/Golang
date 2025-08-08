package bools

/*
Complete the method that takes a boolean value and return a "Yes" string for true, or a "No" string for false.

*/

func ConvertBool(arg bool) string {
	return map[bool]string{true: "Yes", false: "No"}[arg]
}

func ReturnString(arg bool) string {
	if arg {
		return "Yes"
	}
	return "No"
}
