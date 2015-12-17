package functions

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

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
