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
