package errors
import (
	"testing"
	"fmt"
	"strconv"
)

func TestError(t *testing.T) {
	if e := createError(); e != nil {
		fmt.Println(e)
	}
}

func TestError2(t *testing.T) {
	i, err := strconv.Atoi("fff")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
}

func TestError3(t *testing.T) {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}