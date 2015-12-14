package main

import (
	"fmt"
	"github.com/matijavizintin/first/basics/stringutil"
)

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add2(3, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(25))

	languages()
	languages2()
	shortAssignment()
	stringutil.EmptyLine()
}

func add(x int, y int) int {
	return x + y
}

// parameters of same type
func add2(x, y int) int {
	return x + y
}

// multiple returns
func swap(x, y string) (string, string) {
	return y, x
}

// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// multiple variables of boolean
var c, python, java bool

func languages() {
	var i int

	fmt.Println(i, c, python, java)
}

// initialized variables no type required
func languages2() {
	var c, python, java = true, false, "yes!"

	fmt.Println(c, python, java)
}

// short assignment, no var required - available only inside a method
func shortAssignment() {
	c, python, java := false, true, "yes!"

	fmt.Println(c, python, java)
}




