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

// New create one edge and initialize it
func New() *Edge {
	return new(Edge).Init()
}

// Init initialize the edge all fields.
func (edge *Edge) Init() *Edge {
	edge.ID = uuid.New().String()

	edge.Precursor, edge.Successor = nil, nil
	edge.Status.Enabled = true
	edge.Cost = 0

	return edge
}
