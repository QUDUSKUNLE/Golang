package routine

import (
	"fmt"
	"time"
	"sync"
)

type Container struct {
	mu sync.Mutex // guards
	Counters map[string]int
}

func (c *Container) Inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Counters[name]++
}

func Worker(done chan bool) {
	fmt.Print("Working...\n")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func Ping(pings chan <- string, msg string) {
	pings <- msg
}

func Pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings
	pongs <- msg
}
