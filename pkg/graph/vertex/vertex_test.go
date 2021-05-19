package vertex

import (
	_ "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertexNew(t *testing.T) {
	v := New()

	assert.NotNil(t, v)
	assert.NotNil(t, v.ID)

	assert.NotNil(t, v.AdjacencyList)
	assert.Equal(t, v.AdjacencyList.Len(), 0)

	assert.Nil(t, v.Value)
	assert.False(t, v.Status.Visitd)

	assert.NotEqual(t, v.ID, New().ID)
}

func TestVertexInit(t *testing.T) {
	v := New()
	v1 := New()

	v.AdjacencyList.PushBack(v1)
	v.Value = "test-string"
	v.Status.Visitd = true

	id := v.ID

	v.Init()

	assert.Nil(t, v.Value)

	assert.NotNil(t, v.AdjacencyList)
	assert.Equal(t, v.AdjacencyList.Len(), 0)

	assert.False(t, v.Status.Visitd)
	assert.NotEqual(t, id, v.ID)
}
