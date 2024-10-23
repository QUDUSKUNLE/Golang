package loop

import (
	"fmt"
)

func ForI(i int) {
	for i <= 3 {
		fmt.Println(i)
		i++
	}
}

func ForJ() {
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}
}

func ForRange() {
	for i := range 3 {
		fmt.Println("Range", i)
	}
}

// func ArrayIntRange() {
// 	nums := []int{2, 3, 4}
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	fmt.Println("sum:", sum)
// }

// func StringRange() {
// 	kvs := map[string]string{"a": "apple", "b": "banana"}
// 	for key, value := range kvs {
// 		fmt.Printf("%s: %s\n", key, value)
// 	}
// 	// for key := range kvs {
// 	// 	fmt.Println(key)
// 	// }
// 	// for i, c := range "go" {
// 	// 	fmt.Println(i, c)
// 	// }
// }
