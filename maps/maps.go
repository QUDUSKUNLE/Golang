package maps

import "errors"

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot update word because it does not exist")
)

type Dictionary map[string]string

type DictionaryError string

func (dictionaryError DictionaryError) Error() string {
	return string(dictionaryError)
}

func (dictionary Dictionary) Search(word string) (string, error) {
	definition, ok := dictionary[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (dictionary Dictionary) Add(word, definition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		dictionary[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Update(word, newDefinition string) error {
	_, err := dictionary.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		dictionary[word] = newDefinition
	default:
		return err
	}
	return nil
}

func (dictionary Dictionary) Delete(word string) {
	delete(dictionary, word)
}

type LangCategory struct {
	dynamic bool
	strong bool
}
