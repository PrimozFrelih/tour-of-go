package http

import ("net/http"
	"fmt"
)

// http server implements this interface
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// server struct
type HelloServer struct {}

// implementing ServerHTTP interface that create content for the server
func (h HelloServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello first go web server!")
}

// exercise
type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}
