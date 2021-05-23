package vertex

import (
	_ "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVertexNew(t *testing.T) {
	v := New()

	assert.NotNil(t, v)
	assert.NotNil(t, v.id)

	assert.NotNil(t, v.AdjacencyList)
	assert.Equal(t, v.AdjacencyList.Len(), 0)

	assert.Nil(t, v.Value)
	assert.False(t, v.Status.Visitd)

	assert.NotEqual(t, v.id, New().id)
}

func TestVertexInit(t *testing.T) {
	v := New()

	v.AdjacencyList.PushBack(New())
	v.Value = "test-string"
	v.Status.Visitd = true

	id := v.id

	v.Init()

	assert.Nil(t, v.Value)

	assert.NotNil(t, v.AdjacencyList)
	assert.Equal(t, v.AdjacencyList.Len(), 0)

	assert.False(t, v.Status.Visitd)
	assert.Equal(t, id, v.id)
}

func TestVertexSame(t *testing.T) {
	v := New()
	v1 := New()

	assert.True(t, v.IsSame(v))
	assert.False(t, v.IsSame(v1))
	assert.False(t, v.IsSame(nil))
}
