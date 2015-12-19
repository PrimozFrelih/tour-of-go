package select1
import (
	"testing"
	"fmt"
	"time"
)

func TestSelect(t *testing.T) {
	// make two channels
	ch := make(chan int)
	quit := make(chan int)

	// execute goroutine that reads 10 values from ch and then sends exit command
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()

	// push fibonacci numbers through ch
	Fibonacci(ch, quit)
}

func TestSelectDefault(t *testing.T) {
	// channels
	tick := time.Tick(1 * time.Second)
	boom := time.After(10 * time.Second)

	for  {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("boom!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
