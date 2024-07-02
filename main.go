package main

import (
	"fmt"
	"time"

	"github.com/QUDUSKUNLE/Golang/tutorial/channels"
)

func main() {
	go channels.SayHello()
	time.Sleep(time.Second * 2)
	// fmt.Printf("My name is %[1]s. Yes, heard that right: %[1]s\n", name)
	fmt.Println("Done")
}
