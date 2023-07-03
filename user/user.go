package user

import (
	"errors"
)

type Name string

var (
	ErrorEditingFirstName = errors.New("error editing firstname")
	ErrorDeletingLastName = errors.New("error deleting lastname")
)

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
