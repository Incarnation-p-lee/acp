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
	assert.NotNil(t, v.AdjacencyList)

	assert.NotNil(t, v.ID)
	assert.Equal(t, v.Value, s)
	assert.False(t, v.Status.Visitd)

	assert.Equal(t, v.AdjacencyList.Len(), 0)

	id := New(s).ID

	assert.NotEqual(t, v.ID, id)
}
