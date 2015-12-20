package slices
import "fmt"

// helper function for printing slices
func PrintSlice(sliceName string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", sliceName, len(x), cap(x), x)
}

// exercise
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
