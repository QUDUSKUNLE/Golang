package example

import (
	"strconv"
)

func Hello() string {
	return "Hello world"
}

func Convert(str string) (int, error) {
	return strconv.Atoi(str)
}


