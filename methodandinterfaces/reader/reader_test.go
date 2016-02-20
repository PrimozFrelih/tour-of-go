package reader

import (
	"fmt"
	"golang.org/x/tour/reader"
	"io"
	"os"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b) // read data into slice b
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

func TestReader2(t *testing.T) {
	reader.Validate(MyReader{})
}

func TestRot13Reader(t *testing.T) {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	// copy from reader to std output
	io.Copy(os.Stdout, &r)
}
