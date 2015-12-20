package range1
import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	pow := []int{1, 2, 4, 8, 16}

	// print range with for
	for index, element := range pow {
		fmt.Printf("2^%d = %d\n", index, element)
	}

	// reset array
	pow = make([]int, 10)

	// fill with powers
	for idx := range pow {
		pow[idx] = 1 << uint(idx)
	}

	// and print calculated pow array
	for _, value := range pow {        // note: the index is ignored
		fmt.Printf("%d\n", value)
	}
}
