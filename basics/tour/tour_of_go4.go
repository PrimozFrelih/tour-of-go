package tour
import (
	"math"
	"fmt"
	"runtime"
	"time"
)

func NewtonsSqrt(z, x float64) float64 {
	epsilon := math.Pow(10, -9)

	var previousValue float64
	for i := 0; i < 10; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
		fmt.Printf("%v, %g\n", i, z)

		if math.Abs(z - previousValue) < epsilon {
			break
		}

		previousValue = z
	}

	return z
}

func Switch() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
		// fallthrough so the switch doesn't break automatically
	default:
		fmt.Println(os)
	}
}

func Switch2() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("After tomorrow")
	default:
		fmt.Println("Some day")
	}
}

func Switch3() {
	now := time.Now()

	// no condition
	switch {
	case now.Hour() < 12:
		fmt.Println("Morning")
	case now.Hour() < 18:
		fmt.Println("Afternoon")
	default:
		fmt.Println("Evening")
	}
}