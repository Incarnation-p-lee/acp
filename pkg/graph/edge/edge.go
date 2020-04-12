package edge

import (
	"github.com/Incarnation-p-lee/acp/pkg/graph/vertex"
	"github.com/google/uuid"
)

// Status indicates the status of edge
type Status struct {
	Enabled bool
}

// Edge indicates the structure definition for edge
type Edge struct {
	ID        string
	Cost      int64
	Precursor *vertex.Vertex
	Successor *vertex.Vertex
	Status    Status
}

// New create one edge and return the pointer to it
func New(cost int64) *Edge {
	e := new(Edge)

	e.ID = uuid.New().String()
	e.Precursor, e.Successor = nil, nil
	e.Status.Enabled = true

	return e
}
