package tour
import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i				// pointer to i
	fmt.Println(*p)		// print through pointers
	*p = 21				// update i via pointer
	fmt.Println(i)		// print

	p = &j
	fmt.Println(*p)		// pointer to j
	*p = *p / 37		// divide through pointer
	fmt.Println(j)		// print updated value
	fmt.Println(p)		// print pointer - mem location i suppose
}

type Vertex struct {
	X int
	Y int
}

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
