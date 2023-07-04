package main

import (
	"fmt"
	pointer "home/pointer"
	interfaces "home/interface"
)


func main() {
	age := 10
	pointer.Increment(&age)
	fmt.Println(age)
	fmt.Println(interfaces.DisPlayType(float64(12)))
}
