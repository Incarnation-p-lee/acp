package vertex

import (
	_ "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertexNew(t *testing.T) {
	s := "fake-string"
	v := New(s)

	assert.NotNil(t, v)
	assert.NotNil(t, v.SuccessorEdges)
	assert.NotNil(t, v.PrecursorEdges)

	assert.NotNil(t, v.ID)
	assert.Equal(t, v.Value, s)
	assert.False(t, v.Status.Visitd)

	assert.Equal(t, v.SuccessorEdges.Len(), 0)
	assert.Equal(t, v.PrecursorEdges.Len(), 0)

	id := New(s).ID

	assert.NotEqual(t, v.ID, id)
}
