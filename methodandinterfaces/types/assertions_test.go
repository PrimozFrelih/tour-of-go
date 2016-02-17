package types

import (
	"fmt"
	"testing"
)

func TestAssertion(t *testing.T) {
	var i interface{} = "hello"

	// assert (and cast) that i is string
	s := i.(string)
	fmt.Println(s)

	// second parameters prevent the panic
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64)		// NOTE: panic is triggered
}
