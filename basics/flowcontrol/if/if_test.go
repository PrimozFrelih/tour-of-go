package if1
import (
	"testing"
	"fmt"
)

func TestSqrt(t *testing.T) {
	res := Sqrt(10)
	fmt.Println(res)

	res = Sqrt(-10)
	fmt.Println(res)
}

func TestIf(t *testing.T) {
	res := IfWithShortStatement(2, 2, 10)
	fmt.Println(res)

	res = IfWithShortStatement(2, 10, 10)
	fmt.Println(res)
}
