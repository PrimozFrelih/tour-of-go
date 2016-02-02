package variables
import (
	"testing"
	"fmt"
)

// print basic types
func TestBasicTypes(t *testing.T) {
	const f = "type: %T       value: %v\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, Z, Z)
}

// print default values of uninitialized variables
func TestZeroValues(t *testing.T) {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// print constants
func TestConstants(t *testing.T) {
	fmt.Println("Big ", Big)
	fmt.Println("Small ", Small)
}
