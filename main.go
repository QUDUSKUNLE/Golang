package main

import (
	"fmt"
	"home/arrays"
	"home/example"
)

type User struct {
	age int
	sex string
}

var user = User{}

func main() {
	var x float64 = 20.0
	var a, b, c = 3, 4, "foo"
	y := 20
	user.age = 10
	user.sex = "Male"
	agesMap := make(map[string]interface{})
	agesMap["age"] = 12
	agesMap["sex"] = "Male"
	// agesMap["sex"] = "Male"
	fmt.Println("Hello world")
	fmt.Println("Hello World")
	fmt.Println("I am in Go programming world.")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(a + y)
	fmt.Println(a, b, c, ">>>>>>>>>>>>>>>")
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Println(arrays.Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(example.Hello());
	fmt.Println(arrays.Arrayof3)
	fmt.Println(agesMap["sex"])
	fmt.Println(user)
}
