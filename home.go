package main

import "fmt"

func main() {
	var x float64 = 20.0
	var a, b, c = 3, 4, "foo"
	y := 20
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
	fmt.Println(Hello("World", "English"))
}

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix

	switch language {
		case "french":
			prefix = frenchHelloPrefix

		case "spanish":
			prefix = spanishHelloPrefix
		
		default:
			prefix = englishHelloPrefix
	}
	return prefix + name
}

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}
