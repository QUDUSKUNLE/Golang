package main

import (
	"fmt"
	"github.com/QUDUSKUNLE/Golang/tutorial/loop"
)

func main() {
	i := 1
	loop.ForI(i)
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}
}
