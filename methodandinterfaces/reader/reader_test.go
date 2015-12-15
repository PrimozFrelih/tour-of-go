package reader
import (
	"io"
	"fmt"
	"strings"
	"testing"
	"golang.org/x/tour/reader"
	"os"
)

func TestReader(t *testing.T) {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
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
	io.Copy(os.Stdout, &r)
}