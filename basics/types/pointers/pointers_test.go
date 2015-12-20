package pointers
import (
	"fmt"
	"testing"
)

func TestPointers(t *testing.T) {
	i, j := 42, 2701

	p := &i                // pointer to i
	fmt.Println(*p)        // print through pointers
	*p = 21                // update i via pointer
	if i != 21 {
		t.Errorf("i should be 21 instead of %d", i)
	}
	fmt.Println(i)        // print

	p = &j
	fmt.Println(*p)         // pointer to j
	*p = *p / 37        	// divide through pointer
	if j != 73 {
		t.Errorf("i should be 73 instead of %d", i)
	}
	fmt.Println(j)        // print updated value
}
