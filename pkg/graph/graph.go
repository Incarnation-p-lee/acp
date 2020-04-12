package graph

import (
	"github.com/Incarnation-p-lee/acp/pkg/graph/edge"
	"github.com/Incarnation-p-lee/acp/pkg/graph/vertex"
)

// Graph indicates the structure definition for both directed and indirected graph
type Graph struct {
	Attributes Attributes
	Vertexs    map[string]vertex.Vertex
	Edges      map[string]edge.Edge
}

// Attributes indicates some metadata of graph.
type Attributes struct {
	Directed bool
}
