package stringss

/*
You are given a string of n lines, each substring being n characters long. For example:

s = "abcd\nefgh\nijkl\nmnop"

We will study the "horizontal" and the "vertical" scaling of this square of strings.

A k-horizontal scaling of a string consists of replicating k times each character of the string (except '\n').

Example: 2-horizontal scaling of s: => "aabbccdd\neeffgghh\niijjkkll\nmmnnoopp"
A v-vertical scaling of a string consists of replicating v times each part of the squared string.

Example: 2-vertical scaling of s: => "abcd\nabcd\nefgh\nefgh\nijkl\nijkl\nmnop\nmnop"
Function scale(strng, k, v) will perform a k-horizontal scaling and a v-vertical scaling.
*/

import (
	"fmt"
	"strings"
)

func Scale(s string, k, v int) string {
	if s == "" {
		return ""
	}
	strng := strings.Split(s, "\n")
	result := make([]string, len(strng))
	for i, v := range strng {
		for _, c := range v {
			result[i] += strings.Repeat(string(c), k)
		}
	}
	final := make([]string, len(result))
	for j, r := range result {
		final[j] = strings.Repeat(fmt.Sprintf("%s\n", r), v)
	}
	return strings.TrimRight(strings.Join(final, ""), "\n")
}
