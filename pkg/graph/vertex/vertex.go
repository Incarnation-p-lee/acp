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
	ID            string
	AdjacencyList *list.List
	Value         interface{}
	Status        Status
}

// New create one vertex and initialize it
func New() *Vertex {
	return new(Vertex).Init()
}

// Initialize all fields of vertex
func (vertex *Vertex) Init() *Vertex {
	vertex.ID = uuid.New().String()

	vertex.AdjacencyList = list.New()
	vertex.Status.Visitd = false
	vertex.Value = nil

	return vertex
}
