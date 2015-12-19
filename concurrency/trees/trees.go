package trees
import (
	"golang.org/x/tour/tree"
	"time"
)

// tree structure from tour
type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 1000)
	ch2 := make(chan int, 1000)

	// walk through both trees
	Walk(t1, ch1)
	Walk(t2, ch2)

	// compare channels
	counter := 10
	for ; counter > 0;  {
		var v1, v2 int

		select {
		case v1 = <-ch1:
			v2 = <-ch2
		case v2 = <-ch2:
			v1 = <-ch1
		default:
			time.Sleep(10 * time.Millisecond)
			counter--
			continue
		}

		// check value
		if v1 != v2 {
			return false
		}
		counter = 10
	}
	return true
}
