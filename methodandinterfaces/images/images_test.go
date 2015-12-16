package images
import (
	"testing"
	"image"
	"fmt"
	"golang.org/x/tour/pic"
)

func TestImage(t *testing.T) {
	// crate image
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))

	fmt.Println(m.Bounds())				// print bounds
	fmt.Println(m.At(0, 0).RGBA())		// print top left corner
}

func TestImage2(t *testing.T) {
	m := Image{100, 100}
	pic.ShowImage(m)
}
