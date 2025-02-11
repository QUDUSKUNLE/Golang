package stringss

import (
	"strings"
)

/*
altERnaTIng cAsE <=> ALTerNAtiNG CaSe
Define String.prototype.toAlternatingCase (or a similar function/method such as to_alternating_case/toAlternatingCase/ToAlternatingCase in your selected language; see the initial solution for details) such that each lowercase letter becomes uppercase and each uppercase letter becomes lowercase. For example:
*/

func ToAlternatingCase(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		if str[i:i+1] == strings.ToLower(str[i:i+1]) {
			result += strings.ToUpper(str[i : i+1])
		} else {
			result += strings.ToLower(str[i : i+1])
		}
	}
	return result
}
