package methods

import "math"

// struct vertex
type Vertex struct {
	X, Y float64
}

// method on struct vertex
func (v *Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// v is a pointer so the value is read through a reference
func (v *Vertex) Scale(f float64) {
	v.X, v.Y = v.X * f, v.Y * f
}

// v is NOT a pointer so the copy gets mutated
func (v Vertex) Scale1(f float64) {
	v.X, v.Y = v.X * f, v.Y * f
}

// looks like v var of type float
type MyFloat float64

// function defined on MyFloat that is actually a float
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	} else {
		return float64(f)
	}
}


