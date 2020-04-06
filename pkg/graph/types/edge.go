package types

// Edge indicates the structure definition for edge
type Edge struct {
	ID        string
	Cost      int64
	Precursor *Vertex
	Successor *Vertex
	Status    GraphStatus
}
