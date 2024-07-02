package main

import (
	"fmt"
	"time"

	"github.com/QUDUSKUNLE/kafka/src/consumer"
	"github.com/QUDUSKUNLE/kafka/src/producer"
)


func main() {
	producer.Producer()
	time.Sleep(time.Millisecond * 100000)
	fmt.Println("Message produced successfully")
	consumer.Consumer()
}

