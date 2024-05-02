package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContigous(t *testing.T) {
	assert.True(t, AreTilesContiguous(CreateTile(1, 0), CreateTile(2, 0)))
	assert.True(t, AreTilesContiguous(CreateTile(1, 0), CreateTile(1, 1)))
	assert.False(t, AreTilesContiguous(CreateTile(1, 0), CreateTile(3, 0)))
	assert.False(t, AreTilesContiguous(CreateTile(7, 0), CreateTile(4, 2)))
}
