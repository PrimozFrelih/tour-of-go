package functions

// this function takes a parameter fn = [function with two float parameters and returns a float64] and returns a float64
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// this function returns a function that takes an int argument and returns an int
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1, n0 := 0, 0

	// return a function that returns the next fibonacci number
	return func() int {
		if n0 == 0 {
			n0 = 1
			return 1
		}
		if n1 == 0 {
			n1 = 1
			return 1
		}
 		n2 := n1 + n0
		n0, n1 = n1, n2
		return n2
	}
}

// this function returns two strings
func Swap(x, y string) (string, string) {
	return y, x
}

// named return values
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return			// return is empty - return values are named
}
