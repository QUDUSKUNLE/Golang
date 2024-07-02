package channels

import (
	"fmt"
	"time"
)


func SayHello() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello world!")
		time.Sleep(time.Millisecond * 100)
	}
}
