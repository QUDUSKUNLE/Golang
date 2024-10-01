package main

import (
	"fmt"
	"github.com/QUDUSKUNLE/Golang/tutorial/closures"
	// "strings"
)

func main() {
	nextInt := closures.InitSequence()
	for i := range 10 {
		i++
		fmt.Println(nextInt())
	}
}
