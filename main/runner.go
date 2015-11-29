package main

import (
	"github.com/matijavizintin/first/tour"
	"fmt"
	"github.com/matijavizintin/first.bak/stringutil"
)


func main() {
	tour.ForExample()
	tour.ForExample2()
	tour.While()

	tour.Sqrt(4)
	tour.Sqrt(-4)

	var x, exp, limit float64 = 2, 7, 100
	result := tour.IfWithShortStatement(x, exp, limit)
	fmt.Printf("min of %v on %v and %v is %v\n", x, exp, limit, result)
	stringutil.EmptyLine()

	// newtons sqrt
	for i := 0; i < 5; i++ {
		res := tour.NewtonsSqrt(float64(i), 2.0);
		fmt.Printf("z = %v, x = %v, res = %g\n", i, 2, res)
		stringutil.EmptyLine()
	}

	// switch
	tour.Switch()
	stringutil.EmptyLine()

	tour.Switch2()
	stringutil.EmptyLine()

	tour.Switch3()
	stringutil.EmptyLine()

	tour.Pointers()
}
