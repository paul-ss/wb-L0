package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Simple Synchronous Publisher
	err = sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Close connection
	sc.Close()
}
