package reader

import (
	"io"
	"unicode"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

// ROT 13 reader
type rot13Reader struct {
	decorated io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
	// read from decorated reader
	n, err := r.decorated.Read(b)

	// make rot13
	rot13(b)

	return n, err
}

func rot13(b []byte) {
	for i := 0; i < len(b); i++ {
		c := rune(b[i]) // rune is used for unicode characters
		if unicode.IsUpper(c) {
			offset := rune(65)
			b[i] = byte(((c - offset + 13) % 26) + offset)
		} else if unicode.IsLower(c) {
			offset := rune(97)
			b[i] = byte(((c - offset + 13) % 26) + offset)
		}
	}
}
