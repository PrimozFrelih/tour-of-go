package if1
import (
	"fmt"
	"math"
)

// sqrt function
func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if with short statement
func IfWithShortStatement(base, power, limit float64) float64 {

	// init value in if and then make the comparison
	if v := math.Pow(base, power); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)		// if condition var (v) available also in else
	}
	return limit
}
