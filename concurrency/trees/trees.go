package trees

import (
	"golang.org/x/tour/tree"
)

// tree structure from tour
type Tree struct {
	Left  *Tree
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

func WalkWrapper(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// walk through both trees
	go WalkWrapper(t1, ch1)
	go WalkWrapper(t2, ch2)

	// compare channels
	for {
		v1, open1 := <-ch1
		v2, open2 := <-ch2

		// check value
		if v1 != v2 || open1 != open2 {
			return false
		}

		if !open1 && !open2 {
			break
		}
	}
	return true
}
