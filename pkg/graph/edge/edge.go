package edge

import (
	"github.com/Incarnation-p-lee/acp/pkg/graph/vertex"
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
