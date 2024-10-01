package main

import (
	"fmt"
	"github.com/QUDUSKUNLE/Golang/tutorial/slice"
	"maps"
	"strings"
)

func main() {
	fmt.Println("uninit:", slice.S, slice.S == nil, len(slice.S) == 0)

	m := make(map[string]int)
	m["k1"] = 7; m["k2"] = 13
	fmt.Printf("map: %v\ntype: %T\n", m, m)

	v1, v3 := m["k1"], m["k3"];
	fmt.Println(v1, v3)
	fmt.Println("len:", len(m))
	clear(m)
	fmt.Println("len:", len(m))

	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2, "car": 3}

	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
	fmt.Println(strings.EqualFold("Boluwatife08971@", "BOLUWATIFE08971@"))
}
