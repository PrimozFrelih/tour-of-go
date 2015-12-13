package tour
import (
	"testing"
	"fmt"
	"strings"
	"golang.org/x/tour/pic"
)

func TestPointers(t *testing.T) {
	i, j := 42, 2701

	p := &i				// pointer to i
	fmt.Println(*p)		// print through pointers
	*p = 21				// update i via pointer
	if i != 21 {
		t.Errorf("i should be 21 instead of %d", i)
	}
	fmt.Println(i)		// print

	p = &j
	fmt.Println(*p)		// pointer to j
	*p = *p / 37		// divide through pointer
	if j != 73 {
		t.Errorf("i should be 73 instead of %d", i)
	}
	fmt.Println(j)		// print updated value
}

func TestVertex(t *testing.T) {
	v := Vertex{1, 2}
	fmt.Println(v)

	// update field X
	v.X = 5
	fmt.Println(v)

	// pointer to struct
	p := &v
	fmt.Printf("Memory address expected: %p\n", p)
	fmt.Printf("Actual struct expected: %v\n", *p)

	// update through pointer
	p.X = 1e9
	fmt.Printf("New value is: %v\n", *p)
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 5}
	v3 = Vertex{}
	p = &Vertex{1, 2}
)

func TestStructInit(t *testing.T) {
	fmt.Println(v1, v2, v3, p)
}

func TestArrays(t *testing.T) {
	var array [2]string
	array[0] = "Hello"
	array[1] = "Kitty"

	fmt.Printf("%s %s\n", array[0], array[1])
	fmt.Println(array)
}

func TestSlices(t *testing.T) {
	s := []int{1, 3, 5, 7, 9, 14}
	fmt.Printf("%v\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] => %d\n", i, s[i]);
	}
}

func TestTicTacToe(t *testing.T) {
	// board
	game := [][]string {
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
	}

	// update values
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	// print
	for i := 0; i < len(game); i++ {
		fmt.Printf("%s\n", strings.Join(game[i], " "))
	}
}

func TestSlicingSlices(t *testing.T) {
	s := []int{1, 3, 5, 7, 9}
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s[0:1])		// from inclusive, to exclusive => prints element 0
	fmt.Printf("%v\n", s[:3])		// from 0 to 2 inclusive
	fmt.Printf("%v\n", s[2:])		// from 2 to the end
}

func TestSlicesMaker(t *testing.T) {
	a := make([]int, 10)		// 10 zeros
	PrintSlice("a", a)

	b := make([]int, 10, 15)		// 10 zeros but size 15
	PrintSlice("b", b)

	c := b[:2]		// first two elements
	PrintSlice("c", c)

	d := b[2: 5]		// 3 elements
	PrintSlice("d", d)

	e := b[:3]				// pointer
	PrintSlice("e", e)

	// empty slice
	var empty []int
	fmt.Printf("%v, len=%d, cap=%d, isnil=%t\n", empty, len(empty), cap(empty), empty == nil)
}

func TestAppendToSlice(t *testing.T) {
	// empty slice
	var slice []int
	PrintSlice("empty slice", slice)

	// add first element
	slice = append(slice, 1)
	PrintSlice("slice with 1 element", slice)

	// append multiple
	slice = append(slice, 2, 3, 4, 5)
	PrintSlice("slice with multiple elements", slice)
}

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
	for _, value := range pow {		// note: the index is ignored
		fmt.Printf("%d\n", value)
	}
}

func TestExerciseSlices(t *testing.T) {
	pic.Show(Pic)
}


