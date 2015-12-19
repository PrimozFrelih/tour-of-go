package select1
import "fmt"

func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {			// select blocks until one of the cases can run
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
