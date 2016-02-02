package http
import (
	"testing"
	"net/http"
	"log"
)

// disabled to tests aren't stuck in an infinite loop
func testHTTPServer(t *testing.T) {
	var h HelloServer

	// start http server
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}
}

// disabled to tests aren't stuck in an infinite loop
func testHTTPServerExercise(t *testing.T) {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})


	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}
