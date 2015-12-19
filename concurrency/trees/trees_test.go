package trees
import (
	"testing"
	"golang.org/x/tour/tree"
	"fmt"
)

func TestWalk(t *testing.T) {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func TestSame(t *testing.T) {
	result := Same(tree.New(1), tree.New(1))
	fmt.Println(result)

	result = Same(tree.New(1), tree.New(2))
	fmt.Println(result)
}
