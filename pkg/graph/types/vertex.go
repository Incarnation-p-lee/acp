package types

import "container/list"

// Vertex indicates the structure definition for vertex
type Vertex struct {
	ID             string
	SuccessorEdges list.List
	PrecursorEdges list.List
	Value          interface{}
	Status         GraphStatus
}
