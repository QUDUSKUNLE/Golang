package main

import (
	"fmt"
	"os"
	"log"
	"github.com/QUDUSKUNLE/Golang/tutorial/codewar"
	"github.com/QUDUSKUNLE/Golang/tutorial/base"
	// "github.com/QUDUSKUNLE/Golang/tutorial/codewar/array"
	// "github.com/QUDUSKUNLE/Golang/tutorial/codewar/strin"
	// "github.com/QUDUSKUNLE/Golang/tutorial/mtn"
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
	// fmt.Println(codewar.ArrayDiff([]int{1, 2, 2, 2, 3}, []int{2,3}))
	// fmt.Println(strin.ToAlternatingCase("hello world"))
}
