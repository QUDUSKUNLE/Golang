package interfaces

import (
	"fmt"
	"reflect"
)

func DisPlayType(interf interface{}) interface{} {
	fmt.Println(reflect.TypeOf(interf))
	switch theType := interf.(type) {
	case int: 
		return "This is a int " + fmt.Sprintf("%v", theType)
	case float64:
		return "This is a float64 " + fmt.Sprintf("%v", theType)
	case float32:
		return "This is a float32 " + fmt.Sprintf("%v", theType)
	case string:
		return "This is a string " + theType
	default:
		return "unknown"
	}
}

type Employee interface {
	PrintName(name string)
	PrintSalary(basic, tax  int) int
}

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Vehicle interface {
	Structure() []string
	Speed() string
}

type Human interface {
	Structure() []string
	Performance() string
}

// Emp user-defined type
type Emp int
type Pentagon int
type Car string
type Man string

func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t\t", e)
	fmt.Println("Employee Name:\t\t", name)
}

func (e Emp) PrintSalary(basic, tax int) int {
	return basic -  ((basic * tax) / 100)
}

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon:\t", 5*p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Printf("Pentagon has %v sides", p)
}

func (c Car) Structure() []string {
	return []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Tank"}
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}

func (m Man) Structure() []string {
	return []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
}

func (m Man) Performance() string {
	return "8 Hrs/Day"
}
