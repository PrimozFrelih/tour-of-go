package main

import (
	"math/cmplx"
	"fmt"
	"github.com/matijavizintin/first/stringutil"
	"math"
)

func main() {
	basicTypes()
	zeroValues()
	typeConversion()
	typeInference()
	constants()
	printNumericConstants()
}

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func basicTypes() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	stringutil.EmptyLine()
}

// default values of uninitialized variables
func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	stringutil.EmptyLine()
}

// converting between types
func typeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	z := int(f)

	fmt.Println(x, y, z)
	stringutil.EmptyLine()
}

// inferring the right type
func typeInference() {
	v := 42.55
	vi := 0.4 + 15i

	fmt.Printf("v is a type of %T\n", v)
	fmt.Printf("vi is a type of %T\n", vi)
	stringutil.EmptyLine()
}

const C1 = "asdf"
const c1 = "qwerty"

func constants() {
	const c2 = false
	const c3 = 5

	fmt.Println(C1, c1, c2, c3)
	stringutil.EmptyLine()
}

const (
	Big = 1 << 10
	Small = Big >> 9
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func printNumericConstants() {
	fmt.Println("Big ", Big)
	fmt.Println("Small ", Small)
	fmt.Println("needInt of Small", needInt(Small))
	fmt.Println("needFloat of Small", needFloat(Small))
	fmt.Println("needFloat of Big", needFloat(Big))
	fmt.Println("needInt of Big", needInt(Big))
	stringutil.EmptyLine()
}
