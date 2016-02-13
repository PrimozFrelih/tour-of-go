package functions

import (
	"testing"
	"math"
	"fmt"
)

func TestFunctions(t *testing.T) {
	// create a function
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x * x + y * y)
	}

	// invoke a function
	fmt.Println(hypot(5, 12))

	// pass a function as an argument
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func TestClosures(t *testing.T) {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func TestFibonacci(t *testing.T) {
	f := fibonacci()

	// print first 10 fibonacci numbers
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func TestMultipleReturn(t *testing.T) {
	i1, i2 := "10", "11"
	v1, v2 := Swap(i1, i2)

	fmt.Printf("input: %s, %s 	output: %s, %s\n", i1, i2, v1, v2)
}

func TestNamedReturnValues(t *testing.T) {
	// check the method for named return values - also the return is empty
	r1, r2 := Split(100)
	fmt.Println(r1, r2)
}
