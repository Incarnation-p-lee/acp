package vertex

import (
	"container/list"

	"github.com/google/uuid"
)

// Status indicate the status of vertex.
type Status struct {
	Visitd bool
}

// Vertex indicates the structure definition for vertex.
type Vertex struct {
	id            string
	AdjacencyList *list.List
	Value         interface{}
	Status        Status
}

// New create one vertex and initialize it.
func New() *Vertex {
	return new(Vertex).Init()
}

// Init initialize all fields of vertex.
func (v *Vertex) Init() *Vertex {
	if v.id == "" {
		v.id = uuid.New().String()
	}

	v.AdjacencyList = list.New()
	v.Status.Visitd = false
	v.Value = nil

	return v
}

// IsSame indicate if the object of edge is same or not.
func (v *Vertex) IsSame(vertex *Vertex) bool {
	if vertex == nil {
		return false
	}

	return v.id == vertex.id
}
