package main

import (
	"fmt"
	"log"
	"os"

	"github.com/QUDUSKUNLE/Golang/tutorial/base"
	"github.com/QUDUSKUNLE/Golang/tutorial/codewar"
	"github.com/QUDUSKUNLE/Golang/tutorial/variable"
)

func main() {
	fmt.Println(codewar.BinToDec("1101"))

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <file_path>\n", os.Args[0])
	}
	filePath := os.Args[1]

	// Convert the file to base64
	base64Data, err := base.ConvertFileToBase64(filePath)
	if err != nil {
		log.Fatalf("Failed to convert file to base64: %v\n", err)
	}

	// Print the base64 encoded string
	fmt.Println("Base64 Encoded Data:")
	fmt.Println(base64Data)

	// Create a new Base object
	co := variable.Container{
		Base: variable.Base{
			Num: 42,
		},
		Str: "Hello, World!",
	}
	fmt.Printf("Container: %+v\n", co.Base.Num)
	fmt.Printf("Container: %+v\n", co.Describe())

	type Describer interface {
		Describe() string
	}

	var d Describer = co
	fmt.Println(d.Describe())
}
