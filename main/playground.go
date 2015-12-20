package main

import (
	"github.com/matijavizintin/go-first/concurrency/channels"
	"time"
	"fmt"
)


func main() {
	array := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// crate a channel
	channel := make(chan int)

	go channels.SumThroughChannel(array[:len(array) / 2], channel)		// first half
	time.Sleep(100 * time.Millisecond)
	go channels.SumThroughChannel(array[len(array) / 2:], channel)		// second half

	// collect results
	x1, y1 := <-channel, <-channel

	fmt.Println(x1, y1, x1+y1)
}
