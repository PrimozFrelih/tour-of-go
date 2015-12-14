package errors
import (
	"time"
	"fmt"
	"math"
)

type MyError struct {
	When time.Time
	What string
}

// implementing error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func createError() error {
	return &MyError{time.Now(), "error description"}
}


// error type
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("%f is negative. Negative numbers are not supported.", float64(e))
}

// modified sqrt
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}