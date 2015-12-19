package goroutines
import "testing"

func TestGoroutine(t *testing.T) {
	go execute("world")		// runs in a new thread
	execute("hello")
}
