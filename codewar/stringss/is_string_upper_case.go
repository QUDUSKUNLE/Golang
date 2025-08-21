package stringss

/*
Is the string uppercase?
*/

import "strings"

func IsStringUpperCase(str string) bool {
	return strings.ToUpper(str) == str
}
