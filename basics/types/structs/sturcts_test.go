package structs
import (
	"fmt"
	"testing"
)

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

func TestStructInit(t *testing.T) {
	fmt.Println(v1, v2, v3, p)
}