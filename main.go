package main

import (
	"fmt"

	"github.com/QUDUSKUNLE/Golang/tutorial/allocation"
	"github.com/QUDUSKUNLE/Golang/tutorial/pointer"
)

func main() {
	va := 10
	fmt.Println(len(allocation.V))
	fmt.Println(*pointer.Increment(&va))

	// fmt.Printf("My name is %[1]s. Yes, heard that right: %[1]s\n", name)
	fmt.Println("Done")
}
