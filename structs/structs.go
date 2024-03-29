package structs

import "math"

type Month string

const (
	January 	Month = "January"
	February 	Month	= "February"
	March    	Month = "March"
	April    	Month = "April"
	May 		 	Month = "May"
	June			Month = "June"
	July			Month = "July"
	August    Month = "August"
	September Month = "September"
	October		Month = "October"
	November	Month = "November"
	December	Month = "December"
)

type Rectangle struct {
	Width float64
	Height float64
}

// Nested structs
type Square struct {
	Width float32
		Geometry struct {
			Area int
			Perimeter int
	}
}

type Salary struct {
	Basic, HRA, TA float64
	Mon Month
}

type Employee struct {
	FirstName, LastName, Email string
	Age 											 int
	MonthlySalary 						 []Salary
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

// Areas
func (triangle *Triangle) Area() float64 {
	return 0.5 * triangle.Base * triangle.Height
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle *Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

//  Perimeter of a rectangle
func (triangle *Triangle) Perimeter() float64 {
	return triangle.Base + triangle.Height
}

func (rectangle *Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (circle *Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}

func (square *Square) Area() int {
	square.Geometry.Area = int(square.Width * square.Width) 
	return square.Geometry.Area
}

func (square *Square) Perimeter() int {
	square.Geometry.Perimeter = int(2 * (square.Width + square.Width))
	return square.Geometry.Perimeter
}

func (employee *Employee) SalaryEarned() (totalsalary, totalHra, totalTax int) {
	totalSalary, totalHra, totalTax := 0, 0, 0
	for _, salary := range employee.MonthlySalary {
		totalSalary = totalSalary + int(salary.Basic)
		totalHra = totalHra + int(salary.HRA)
		totalTax = totalTax + int(salary.TA)
	}
	return totalSalary, totalHra, totalTax
}
