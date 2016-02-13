package maps

import (
	"strings"
	"strconv"
)

type Vertex struct {
	Lat, Long float64
}

// int hash set simulation with map
type IntSet map[int]struct{}

func NewIntSet() IntSet {
	return IntSet(make(map[int]struct{}))
}

// functions defined on type IntSet
func (set IntSet) add(value int) {
	set[value] = struct{}{}
}

func (set IntSet) contains(value int) bool {
	_, present := set[value]
	return present
}

// implement fmt.Stringer on type IntSet
func (set IntSet) String() string {
	keys := make([]string, 0, len(set))
	for key := range set {
		keys = append(keys, strconv.Itoa(key))
	}
	return strings.Join(keys, ", ")
}

// exercise
func WordCount(s string) map[string]int {
	// init map
	var wordsMap = make(map[string]int)

	// split into fields
	fields := strings.Fields(s)

	// go through fields
	for _, value := range fields {
		wordsMap[value] = wordsMap[value] + 1
	}

	return wordsMap
}
