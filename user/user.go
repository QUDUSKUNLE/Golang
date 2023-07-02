package user

import (
	"errors"
)

type FullName string

var (
	ErrorEditingFirstName = errors.New("error editing firstname")
	ErrorDeletingLastName = errors.New("error deleting lastname")
)

type Name struct {
	FirstName FullName
	LastName FullName
}

func (name *Name) FullName(full Name) FullName {
	name.FirstName, name.LastName = full.FirstName, full.LastName
	return FullName(name.FirstName + " " + name.LastName) 
}

func (name *Name) EditFirstName(newName FullName) error {
	if newName == name.FirstName {
		return ErrorEditingFirstName
	}
	name.FirstName = newName
	return nil
}

func (name *Name) DeleteLastName(lastName FullName) error {
	if lastName == name.LastName {
		return ErrorDeletingLastName
	}
	name.LastName = ""
	return nil
}

func (name *Name) GetFirstName() FullName {
	return FullName(name.FirstName)
}

func (name *Name) GetLastName() FullName {
	return FullName(name.LastName)
}
