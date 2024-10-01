package main

import (
	// "fmt"
	"github.com/QUDUSKUNLE/Golang/tutorial/variadic"
	// "strings"
)

func main() {
	variadic.Sum(1, 2)
	variadic.Sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	variadic.Sum(nums...)
}
