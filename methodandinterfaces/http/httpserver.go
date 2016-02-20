package http

import (
	"fmt"
	"net/http"
)

// http server implements this interface
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// server struct
type HelloServer struct{}

// implementing ServerHTTP interface that create content for the server
func (h HelloServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello first go web server!")
}

// exercise
type String string

type HTTPStruct struct {
	Greeting string
	Punct    string
	Who      string
}

// implement handler (http server) interface on type String
func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

// implement handler interface on HTTPStruct
func (s *HTTPStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}
