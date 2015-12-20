package for1
import (
	"math"
	"fmt"
)

// exercise
func NewtonsSqrt(z, x float64) float64 {
	epsilon := math.Pow(10, -12)

	var previousValue float64
	for i := 0; i < 100; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		fmt.Printf("%v, %g\n", i, z)

		if math.Abs(z - previousValue) < epsilon {
			break
		}

		previousValue = z
	}
	return z
}
