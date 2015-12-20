package for1
import (
	"testing"
	"fmt"
)

func TestWhile(t *testing.T) {
	sum := 1

	// while loop
	for sum < 1000 {
		sum *= 2
	}
	fmt.Println(sum)
}

func TestNewtonSqrt(t *testing.T) {
	// compute sqrt
	var input float64 = 2
	var initValue float64 = 2
	res := NewtonsSqrt(input, initValue)

	// print
	fmt.Printf("input: %v, res: %v\n", input, res)
}
