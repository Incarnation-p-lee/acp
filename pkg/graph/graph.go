package graph

import (
	"github.com/Incarnation-p-lee/acp/pkg/graph/edge"
	"github.com/Incarnation-p-lee/acp/pkg/graph/vertex"
)

// Graph indicates the structure definition for both directed and indirected graph
type Graph struct {
	Attributes Attributes
	Vertexes   map[string]vertex.Vertex
	Edges      map[string]edge.Edge
}

// Attributes indicates some metadata of graph
type Attributes struct {
	Directed bool
}

// New will return one empty graph
func New() *Graph {
	g := new(Graph)

	g.Vertexes = make(map[string]vertex.Vertex)
	g.Edges = make(map[string]edge.Edge)

	g.Attributes.Directed = true

	return g
}
