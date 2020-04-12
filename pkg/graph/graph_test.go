package graph

import (
	_ "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphNew(t *testing.T) {
	g := New()

	assert.NotNil(t, g)
	assert.NotNil(t, g.Vertexes)
	assert.NotNil(t, g.Edges)
	assert.NotNil(t, g.Attributes)
	assert.True(t, g.Attributes.Directed)
}
