package edge

import (
	_ "log"
	"testing"

	"github.com/Incarnation-p-lee/acp/pkg/graph/vertex"
	"github.com/stretchr/testify/assert"
)

func TestEdgeNew(t *testing.T) {
	e := New()

	assert.NotNil(t, e.id)

	assert.Equal(t, e.Cost, int64(0))

	assert.Nil(t, e.Precursor)
	assert.Nil(t, e.Successor)

	assert.True(t, e.Status.Enabled)

	assert.NotEqual(t, e.id, New().id)
}

func TestEdgeInit(t *testing.T) {
	e := New()

	e.Precursor, e.Successor = vertex.New(), vertex.New()
	e.Cost = 1234
	e.Status.Enabled = false

	id := e.id

	e.Init()

	assert.Equal(t, e.Cost, int64(0))

	assert.Nil(t, e.Precursor)
	assert.Nil(t, e.Successor)

	assert.True(t, e.Status.Enabled)
	assert.Equal(t, id, e.id)
}

func TestEdgeSame(t *testing.T) {
	e := New()
	e1 := New()

	assert.True(t, e.IsSame(e))
	assert.False(t, e.IsSame(e1))
	assert.False(t, e.IsSame(nil))
}
