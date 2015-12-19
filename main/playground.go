package main

import (
	"github.com/matijavizintin/first/basics/tour"
	"fmt"
	"github.com/matijavizintin/first/basics/stringutil"
	"golang.org/x/tour/pic"
	"github.com/matijavizintin/first/methodandinterfaces/methods"
	"github.com/matijavizintin/first/concurrency/channels"
	"time"
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

	pic.Show(tour.Pic)

	v0 := &methods.Vertex{1, 2}
	fmt.Printf("Abs value of vertex is %f\n", v0.Abs())

	// using a reference
	v := &methods.Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())

	// using a value type it works the same
	v1 := methods.Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v1, v1.Abs())
	v1.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v1, v1.Abs())

	// passing to a method that accepts the value type doesn't change the v2
	v2 := &methods.Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v2, v2.Abs())
	v2.Scale1(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v2, v2.Abs())

	array := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}

	// crate a channel
	channel := make(chan int)

	go channels.SumThroughChannel(array[:len(array) / 2], channel)		// first half
	time.Sleep(100 * time.Millisecond)
	go channels.SumThroughChannel(array[len(array) / 2:], channel)		// second half

	// collect results
	x1, y1 := <-channel, <-channel

	fmt.Println(x1, y1, x1+y1)
}
