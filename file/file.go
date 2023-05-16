package main

import (
	"fmt"
	"os"
	"errors"
	"io/ioutil"
)

func main() {
	filename, err := FileOpen("data.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		// Read the content of the file
		filecontent, errr := ReadFile(filename)
		if errr != nil {
			fmt.Println(errr)
		}
		fmt.Println(filecontent)
	}

}

func FileOpen(name string) (string, error) {
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

