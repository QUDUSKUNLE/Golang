package file

import (
	"os"
	"errors"
	"io/ioutil"
)


func OpenFile(name string) (string, error) {
	f, er := os.Open(name)
	if er != nil {
		return "", errors.New("custom error message: file name is wrong")
	} else {
		return f.Name(), nil
	}
}

func ReadFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("error reading file")
	}
	return string(data), nil
}

func WriteFile(filename, content string) (string, error) {
	file, err := os.Create(filename)
	if err != nil {
		return "", errors.New("error creating a file")
	}
	_, er := file.WriteString(content)
	if er != nil {
		file.Close()
		return "", errors.New("error writing to a file")
	}
	er = file.Close()
	if er != nil {
		return "", errors.New("error closing the file")
	}
	return "Done", nil
}