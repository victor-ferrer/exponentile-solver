package solver

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestSwap(t *testing.T) {

	b := getGameBoard()
	b.Swap(CreateTile(3, 0), CreateTile(4, 0))

	val, err := b.Get(3, 0)
	assert.NoError(t, err)
	assert.Equal(t, 8, val)

	val, err = b.Get(4, 0)
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	_, err = b.Get(8, 0)
	assert.Error(t, err)

}

func getGameBoard() Board {
	b := make([][]int, width)
	b[0] = []int{2, 8, 2, 4, 4, 2, 8, 16}
	b[1] = []int{4, 16, 4, 16, 16, 8, 16, 8}
	b[2] = []int{4, 8, 2, 8, 4, 16, 8, 16}
	b[3] = []int{2, 16, 4, 8, 2, 16, 8, 4}
	b[4] = []int{8, 16, 2, 2, 16, 4, 16, 8}
	b[5] = []int{4, 2, 8, 8, 2, 8, 4, 8}
	b[6] = []int{8, 2, 8, 2, 4, 2, 16, 4}
	b[7] = []int{16, 16, 2, 16, 8, 2, 4, 2}
	return b
}
