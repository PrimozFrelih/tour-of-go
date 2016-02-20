package http

import (
	"log"
	"net/http"
	"testing"
)

// NOTE: tests get stuck in an infinite loop because they start a server
func TestHTTPServer(t *testing.T) {
	var h HelloServer

	// start http server
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}

func TestHTTPServerExercise(t *testing.T) {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &HTTPStruct{"Hello", ":", "Gophers!"})

	// run a server (with two handlers) that listens on 2 urls
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
