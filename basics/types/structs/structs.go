package structs

type Vertex struct {
	X int
	Y int
}

// struct initialization
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 5}		// Y is empty
	v3 = Vertex{}			// empty
	p = &Vertex{1, 2}		// pointer to vertex
)
