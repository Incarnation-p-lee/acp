package types

// Graph indicates the structure definition for both directed and indirected graph
type Graph struct {
	Attributes GraphAttributes
	Vertexs    map[string]Vertex
	Edges      map[string]Edge
}

// GraphStatus indicates extra status for vertex.
type GraphStatus struct {
	Visited bool
}

// GraphAttributes indicates some metadata of graph.
type GraphAttributes struct {
	Directed bool
}
