package edge

import (
	"github.com/Incarnation-p-lee/acp/pkg/graph/vertex"
	"github.com/google/uuid"
)

// Status indicates the status of edge.
type Status struct {
	Enabled bool
}

// Edge indicates the structure definition for edge.
type Edge struct {
	id        string
	Cost      int64
	Precursor *vertex.Vertex
	Successor *vertex.Vertex
	Status    Status
}

// New create one edge and initialize it.
func New() *Edge {
	return new(Edge).Init()
}

// Init initialize the edge all fields.
func (e *Edge) Init() *Edge {
	if e.id == "" {
		e.id = uuid.New().String()
	}

	e.Precursor, e.Successor = nil, nil
	e.Status.Enabled = true
	e.Cost = 0

	return e
}

// IsSame indicate if the object of edge is same or not.
func (e *Edge) IsSame(edge *Edge) bool {
	if edge == nil {
		return false
	}

	return e.id == edge.id
}
