package channels
import (
	"testing"
	"fmt"
)

func TestChannel(t *testing.T) {
	array := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// crate a channel
	channel := make(chan int)

	go SumThroughChannel(array[:len(array) / 2], channel)		// first half
	go SumThroughChannel(array[len(array) / 2:], channel)		// second half

	// collect results
	x, y := <-channel, <-channel

	fmt.Println(x, y, x+y)
}

func TestBufferedChannel(t *testing.T) {
	ch := make(chan int, 100)		// buffered

	// write to channel
	for i := 0; i < 100; i++ {
		ch <- i		// fill the buffer
	}
	ch <- 2			// NOTE: and this should block => deadlock    ==> test will fail

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func TestRangeAndClose(t *testing.T) {
	ch := make(chan int, 1000)

	// start calculating fibonacci numbers
	go Fibonacci(cap(ch), ch)
	for number := range ch {
		fmt.Println(number)
	}
}

func TestChannelState(t *testing.T) {
	ch := make(chan int, 10)

	ch <- 1		// write sth
	ch <- 2		// write sth
	close(ch)

	val, state := <-ch
	fmt.Printf("val: %v, state: %v\n", val, state)

	val, state = <-ch
	fmt.Printf("val: %v, state: %v\n", val, state)

	val, state = <-ch
	fmt.Printf("val: %v, state: %v\n", val, state)
}
