package variables
import "math/cmplx"


var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	Z complex128 = cmplx.Sqrt(-5 + 12i)
)

// constants
const (
	Big = 1 << 10
	Small = Big >> 9
)
