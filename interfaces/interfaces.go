package interfaces
import (
	"math"
	"fmt"
)

// interface
type Abser interface {
	Abs() float64
}

// types that will "implement" Abser
type Vertex struct {
	X, Y float64
}

type MyFloat float64

// methods on types
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y + v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}


// interfaces
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWrite interface {
	Reader
	Writer
}

// struct that will implement Stringer
type Person struct {
	Name string
	Age int
}

// implementing Stringer
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}


