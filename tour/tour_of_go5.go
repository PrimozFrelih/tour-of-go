package tour
import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i				// pointer to i
	fmt.Println(*p)		// print through pointers
	*p = 21				// update i via pointer
	fmt.Println(i)		// print

	p = &j
	fmt.Println(*p)		// pointer to j
	*p = *p / 37		// divide through pointer
	fmt.Println(j)		// print updated value
	fmt.Println(p)		// print pointer - mem location i suppose
}

type Vertex struct {
	X int
	Y int
}

func PrintSlice(sliceName string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", sliceName, len(x), cap(x), x)
}

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8

	for i := 0; i < dy; i = i + 1 {
		var row []uint8
		for j := 0; j < dx; j = j + 1 {
			row = append(row, uint8((i + j) / 2))
		}
		pic = append(pic, row)
	}
	return pic
}
