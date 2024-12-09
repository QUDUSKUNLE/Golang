package codewars

/*
Is the string uppercase?
*/

import (
	"strings"
	"unicode"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	return string(s) == strings.ToUpper(string(s))
}

func (s MyString) IsUpper() bool {
	for _, c := range s {
		if unicode.IsLower(c) {
			return false
		}
	}
	return true
}
