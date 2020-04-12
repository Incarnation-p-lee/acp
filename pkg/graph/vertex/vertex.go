package vertex

import (
	"container/list"
)

// Status indicate the status of vertex
type Status struct {
	Visitd bool
}

// Vertex indicates the structure definition for vertex
type Vertex struct {
	ID             string
	SuccessorEdges list.List
	PrecursorEdges list.List
	Value          interface{}
	Status         Status
}
