package maps

import (
	"fmt"
	"testing"
	"golang.org/x/tour/wc"
)

func TestMap(t *testing.T) {
	var m map[string]Vertex

	// init map
	m = make(map[string]Vertex)

	// fill data
	key1 := "something"
	key2 := "something_else"
	m[key1] = Vertex{1, 2}
	m[key2] = Vertex{3, 4}

	fmt.Printf("values for %v = %v and %v = %v\n\n", key1, m[key1], key2, m[key2])

	// init and fill
	m1 := map[string]Vertex {
		"key1": Vertex{11, 12},
		"key2": Vertex{22, 23},
	}

	fmt.Printf("map to string ==> %v\n\n", m1)

	// even better
	m2 := map[string]Vertex{
		"Valentino": {46, 46},
		"Rossi": {46, 46},		// 46 all the way
	}
	fmt.Printf("46 all the way ==> %v\n\n", m2)
}

func TestMutatingMaps(t *testing.T) {
	// init map
	m :=map[string]int{
		"k1": 1,
		"k2": 2,
	}

	// insert element
	val, present := m["k3"]
	fmt.Printf("Present element at k3 = %v val = %v\n", present, val)
	m["k3"] = 3
	val, present = m["k3"]
	fmt.Printf("Present element at k3 = %v val = %v\n", present, val)

	// delete element
	val, present = m["k3"]
	fmt.Printf("Present element at k3 = %v val = %v\n", present, val)
	delete(m, "k3")
	val, present = m["k3"]
	fmt.Printf("Present element at k3 = %v val = %v\n", present, val)
}

func TestExerciseMaps(t *testing.T) {
	wc.Test(WordCount)
}