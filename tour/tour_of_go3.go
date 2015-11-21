package tour

import (
	"fmt"
	"github.com/matijavizintin/first/stringutil"
	"math"
)

// simple for loop
func ForExample() (sum int) {
	sum = 0

	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	stringutil.EmptyLine()

	return sum
}

// while loop
func ForExample2() int {
	sum := 1

	for ; sum < 1000; {
		sum *= 2
	}
	fmt.Println(sum)
	stringutil.EmptyLine()

	return sum
}

// actually while loop :)
func While() int {
	sum := 1

	for sum < 1000 {
		sum *= 2
	}
	fmt.Println(sum)
	stringutil.EmptyLine()

	return sum
}

// sqrt function
func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if with short statement
func IfWithShortStatement(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)		// if condition available also in else
	}
	return limit
}
