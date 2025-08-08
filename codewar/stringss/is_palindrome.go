package stringss

import (
	"strings"
)

/*
Write a function that checks if a given string (case insensitive) is a palindrome.

A palindrome is a word, number, phrase, or other sequence of symbols that reads the same backwards as forwards, such as madam or racecar.

*/

func IsPalindrome(arg string) bool {
	if len(arg) == 1 {
		return true
	}
	half, reminder  := len(arg) / 2, len(arg) % 2
	var halfArg string
	for i := len(arg); i > half; i-- {
		halfArg += strings.ToLower(arg[i-1:i])
	}
	return strings.ToLower(arg[0:half+reminder]) == halfArg
}
