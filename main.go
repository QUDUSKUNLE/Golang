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

	var employee interfaces.Employee = interfaces.Emp(2)
	employee.PrintName("Abdul-Quddus Yekeen")
	fmt.Println("Employee Salary:\t", employee.PrintSalary(25000, 5))

	var polygon interfaces.Polygons = interfaces.Pentagon(5)
	polygon.Perimeter()
	var pentagon interfaces.Pentagon = polygon.(interfaces.Pentagon)
	pentagon.NumberOfSide()
}
