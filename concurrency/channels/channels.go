package channels

func SumThroughChannel(array []int, channel chan int) {
	sum := 0
	for _, value := range array {
		sum += value
	}

	// send back the value through the channel
	channel <- sum
}

func Fibonacci(n int, ch chan int)  {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x		// send next fibonacci number through the channel
		x, y = y, x+y
	}

	close(ch)		// close the channel
}
