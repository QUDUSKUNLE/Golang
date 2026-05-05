package generic

import "fmt"


// Paramatized type in golang

type Radius interface {
	int64 | int8 | float64
}

func Generic_Circumference[r int | float32](radius r) {
	c := 2 * 3 * radius
	fmt.Printf("The circumference is %v", c)
}

func Paramatized_Circumference[R Radius](radius R) {
	c := 2 * 3 * radius
	fmt.Printf("The circumference is %v", c)
}
