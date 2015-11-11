package hello_test

import (
	"testing"
	"github.com/matijavizintin/first/tour"
)

func TestFor(t *testing.T) {
	max := 100
	f1 := tour.ForExample()
	expected := (max * (max + 1)) / 2

	if f1 != expected {
		t.Errorf("%d != %d", f1, expected)
	}
}

func TestFor2(t *testing.T) {
	test := tour.ForExample2()
	expected := 1024

	if test != expected {
		t.Errorf("%d != %d", test, expected)
	}
}

func TestSqrt(t *testing.T) {
	// test positive
	test := tour.Sqrt(4)
	expected := "2"

	if test != expected {
		t.Errorf("%s != %s", test, expected)
	}

	// test negative
	test = tour.Sqrt(-4)
	expected = "2i"

	if test != expected {
		t.Errorf("%s != %s", test, expected)
	}
}