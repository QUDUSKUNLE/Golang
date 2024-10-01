package main

import (
	"fmt"
	"github.com/QUDUSKUNLE/Golang/tutorial/slice"
	"slices"
)

func main() {
	fmt.Println("uninit:", slice.S, slice.S == nil, len(slice.S) == 0, cap(slice.S))
	slice.S = make([]string, 3)
	fmt.Println("emp", slice.S, "len:", len(slice.S), "cap:", cap(slice.S))

	slice.S[0] = "a"; slice.S[1] = "b"; slice.S[2] = "c";
	fmt.Printf("set: %s\n", slice.S)
	fmt.Printf("get: %s\n", slice.S[2])
	fmt.Printf("len: %d and cap: %d\n", len(slice.S), cap(slice.S))

	slice.S = append(slice.S, "d")
	slice.S = append(slice.S, "e", "f")
	fmt.Println(slice.S)
	t := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}
	
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}
	d := slices.Clone(slice.S)
	fmt.Println(d)
	fmt.Printf("i has value: %v and type: %T\n", d, d)
}
