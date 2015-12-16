package http
import (
	"testing"
	"net/http"
	"log"
)

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
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})


	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
