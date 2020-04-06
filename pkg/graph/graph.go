package graph

// edge "github.com/incarnation-p-lee/acp/pkg/graph/edge"
// vertex "github.com/incarnation-p-lee/acp/pkg/graph/vertex"

// Graph indicates the structure definition for both directed and indirected graph
type Graph struct {
	Attributes Attributes
	// Vertexs    map[string]vertex.Vertex
	// Edges      map[string]edge.Edge
}

// Status indicates extra status for vertex.
type Status struct {
	Visited bool
}

// Attributes indicates some metadata of graph.
type Attributes struct {
	Directed bool
}
