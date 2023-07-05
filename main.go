package main

import (
	"fmt"
	"reflect"
	structs "home/structs"
	example "home/example"
)


func main() {

	res, _ := example.Convert("10")
	fmt.Printf("The value is %v, and is of type %s\t", res, reflect.TypeOf(res))
	fmt.Println(structs.Rectangle{Width: 10.5, Height: 10.2})

	var square = structs.Square{ Width: 10 }

	var square2 = new(structs.Square)
	square2.Width = 20

	var square3 = &structs.Square{ Width: 50 }

	var square4 = &structs.Square{}
	square4.Width = 60

	var salary1 = structs.Salary{Basic: 15000.00, HRA: 5000.00, TA: 2000, Mon: structs.January }
	var salary2 = structs.Salary{Basic: 150000.00, HRA: 500.00, TA: 200, Mon: structs.February }
	var employee = structs.Employee{
		FirstName: "Abdul-Quddus",
		LastName: "Yekeen",
		Email: "qudus@gmail.com",
		Age: 35,
		MonthlySalary: []structs.Salary{salary1, salary2},
	}

	fmt.Println("The Area of the square is:\t", square.Area())
	fmt.Println("The Perimeter of the square is:\t", square.Perimeter())

	fmt.Println("The Ares of the square2 is:\t", square2.Area())
	fmt.Println("The Perimeter of the square2 is:\t", square2.Perimeter())

	fmt.Println("The Ares of the square3 is:\t", square3.Area())
	fmt.Println("The Perimeter of the square3 is:\t", square3.Perimeter())

	fmt.Println("The Ares of the square4 is:\t", square4.Area())
	fmt.Println("The Perimeter of the square4 is:\t", square4.Perimeter())

	fmt.Println(employee.MonthlySalary)
	fmt.Println(employee.SalaryEarned())
}
