package main

import (
	"fmt"
)

func main() {
	number := 100000000
	better := 100_000_000
	arr := [3]int{1, 2, 3}
	sameArr := [...]int{ 1, 2, 3}

	fmt.Println(arr, sameArr)
	fmt.Println(number == better, better)
	fmt.Println("Done")
}
