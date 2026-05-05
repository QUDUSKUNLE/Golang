package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
	// algo "github.com/QUDUSKUNLE/Golang/tutorial/Algorithm"
)

type value struct {
	mu    sync.Mutex
	value int
}

func main() {
	// var memoryAccess sync.Mutex
	// var data int
	// go func() {
	// 	memoryAccess.Lock()
	// 	data++
	// 	memoryAccess.Unlock()
	// }()
	// memoryAccess.Lock()
	// if data == 0 {
	// 	fmt.Printf("the value is, 0")
	// } else {
	// 	fmt.Printf("the value is %v.\n", data)
	// }
	// memoryAccess.Unlock()
	// fmt.Println(algo.ExpressionMatter(1, 1, 1))

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
