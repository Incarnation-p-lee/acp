package edge

import (
	_ "log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeNew(t *testing.T) {
	c := int64(1234)
	e := New(c)

	assert.NotNil(t, e)
	assert.NotNil(t, e.ID)
	assert.Nil(t, e.Precursor)
	assert.Nil(t, e.Successor)
	assert.True(t, e.Status.Enabled)

	id := New(c).ID

	assert.NotEqual(t, e.ID, id)
}
