package main

import (
	"fmt"
	"sync"
	channels "tutorial/channels"
)

var wg sync.WaitGroup

func main() {
	input := make(chan int)
	output := make(chan int)
	done := make(chan bool)

	wg.Add(1)
	go channels.Worker(input, output, done, &wg)

	for i := 0; i < 10; i++ {
		input <- i
	}

	close(input)

	for n := range output {
		fmt.Println(n)
	}

	done <-true
	wg.Wait()
	fmt.Println("done")
}
