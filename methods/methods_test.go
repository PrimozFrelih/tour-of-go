package methods

import (
	"testing"
	"fmt"
)

func TestMethod(t *testing.T) {
	v := &Vertex{1, 2}
	fmt.Printf("Abs value of vertex is %f\n", v.Abs())

	// NOTE: without a reference, value is copied on method call
	v1 := Vertex{1, 2}
	fmt.Printf("Abs value of vertex is %f\n", v1.Abs())
}

func TestMethod2(t *testing.T) {
	f := MyFloat(-10)
	fmt.Printf("Abs value of float -10 is %f\n", f.Abs())
}

func TestMethod3(t *testing.T) {
	// using a reference
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

	// using a value type it works the same
	v1 := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v1, v1.Abs())
	v1.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v1, v1.Abs())

	// passing to a method that accepts the value type doesn't change the v2
	v2 := Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs())
	v2.Scale1(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs())
}
