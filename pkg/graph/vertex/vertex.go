package vertex

import (
	"container/list"

	"github.com/google/uuid"
)

// Status indicate the status of vertex
type Status struct {
	Visitd bool
}

// Vertex indicates the structure definition for vertex
type Vertex struct {
	ID             string
	SuccessorEdges *list.List
	PrecursorEdges *list.List
	Value          interface{}
	Status         Status
}

// New create one vertex and return the pointer to it
func New(value interface{}) *Vertex {
	v := new(Vertex)

	v.ID = uuid.New().String()
	v.SuccessorEdges, v.PrecursorEdges = list.New(), list.New()
	v.Status.Visitd = false

	v.Value = value

	return v
}
