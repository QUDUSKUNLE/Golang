package user

import (
	"errors"
	"math/rand"
)

type Name string

var (
	ErrorEditingFirstName = errors.New("error editing firstname")
	ErrorDeletingLastName = errors.New("error deleting lastname")
)
const MaxRnd = 16

type FullName struct {
	FirstName string
	LastName string
}

type Person struct {
	Name FullName
	Age int
}

func (name Person) FullName() Name {
	return Name(name.Name.FirstName + " " + name.Name.LastName)
}

func (name *Person) EditFirstName(newName FullName) error {
	if newName.FirstName == name.Name.FirstName {
		return ErrorEditingFirstName
	}
	name.Name.FirstName = newName.FirstName
	return nil
}

func (name *Person) DeleteLastName(lastName FullName) error {
	if lastName.LastName == name.Name.LastName {
		return ErrorDeletingLastName
	}
	name.Name.LastName = ""
	return nil
}

func (name *Person) GetFirstName() Name {
	return Name(name.Name.FirstName)
}

func (name *Person) GetLastName() Name {
	return Name(name.Name.LastName)
}

func StatRandonNumbers(n int) (int, int) {
	var a, b int
	for i := 0; i < n; i++ {
		if rand.Intn(MaxRnd) < MaxRnd / 2 {
			a = a + 1
		} else {
			b++
		}
	}
	return a, b
}

func SquaresOfSumAndDiff(x, y int) (int, int) {
	return (x + y) * (x + y), (x - y) * (x - y)
}

func CompareLower4Bits(m, n uint32) bool {
	return m&0xF > n&0xF
}

// anonymous function
