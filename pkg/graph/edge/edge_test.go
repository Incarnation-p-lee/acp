package edge

import (
	_ "log"
	"testing"

	"github.com/Incarnation-p-lee/acp/pkg/graph/vertex"
	"github.com/stretchr/testify/assert"
)

func TestEdgeNew(t *testing.T) {
	e := New()

	assert.NotNil(t, e.ID)

	assert.Equal(t, e.Cost, int64(0))

	assert.Nil(t, e.Precursor)
	assert.Nil(t, e.Successor)

	assert.True(t, e.Status.Enabled)

	assert.NotEqual(t, e.ID, New().ID)
}

func TestEdgeInit(t *testing.T) {
	e := New()
	v1, v2 := vertex.New(), vertex.New()

	e.Precursor, e.Successor = v1, v2
	e.Cost = 1234
	e.Status.Enabled = false

	id := e.ID

	e.Init()

	assert.Equal(t, e.Cost, int64(0))

	assert.Nil(t, e.Precursor)
	assert.Nil(t, e.Successor)

	assert.True(t, e.Status.Enabled)
	assert.NotEqual(t, id, e.ID)
}
