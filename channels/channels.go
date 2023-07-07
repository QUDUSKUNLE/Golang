package channels

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)
// Channels basically used for communication between goroutines

// WaitGroup is used to wait for the program to finish goroutines
var Wg sync.WaitGroup

func ResponseSize(url string, nums chan int) {
	defer Wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	nums <-len(body)
}

var i int
// Play and Pause execution
func work() {
	time.Sleep(100 * time.Millisecond)
	i++
	fmt.Println(i)
}

func Routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func Fibronacci(in int, ch chan int) {
	x, y := 0, 1
	for i :=0; i < in; i++ {
		ch <- x
		x, y = y, x + y
	}
	close(ch)
}

// How to gracefully close channels in Golang
func Worker(input chan int, output chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case n := <-input:
			output <- n * 2
		case <-done:
			close(output)
			return
		}
	}
}
