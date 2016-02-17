package interfaces

import (
	"fmt"
	"math"
	"os"
	"testing"
)

func TestInterfaces(t *testing.T) {
	var a Abser
	f := MyFloat(-math.Sqrt2) // go uses implicit implementation of an interface - by implementing its methods
	v := Vertex{3, 4}

	a = f  // set MyFloat to Abser
	a = &v // set a pointer to Vertex to Abser
	//	a = v	// NOTE: this doesn't work because Abs method on Vertex accepts a pointer

	fmt.Println(a.Abs())
}

func TestInterface2(t *testing.T) {
	var w Writer

	w = os.Stdout                     // Stdout is pointer to a file that implements Writer interface
	fmt.Fprintf(w, "hello, writer\n") // print to Writer
}

func TestStringer(t *testing.T) {
	p1 := Person{"Name1", 30}
	p2 := Person{"Name2", 31}

	fmt.Println(p1, p2) // will be formatted using Person's Stringer method
}

func TestIP(t *testing.T) {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for key, value := range addrs {
		fmt.Printf("%v: %v\n", key, value)
	}
}
