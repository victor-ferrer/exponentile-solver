package solver

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gonum.org/v1/gonum/mat"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

func TestSwap(t *testing.T) {

	b := getGameBoard()
	b.Swap(CreateTile(3, 0), CreateTile(4, 0))

	val := b.Get(3, 0)
	assert.Equal(t, 8, val)

	val = b.Get(4, 0)
	assert.Equal(t, 2, val)

}

func TestDrop(t *testing.T) {
	// TODO this could be a table test
	b := getGameBoard()

	b.DropTile(CreateTile(7, 0), 32)
	val := b.Get(7, 0)
	assert.Equal(t, 8, val)

	val = b.Get(6, 0)
	assert.Equal(t, 4, val)

	val = b.Get(5, 0)
	assert.Equal(t, 8, val)

	val = b.Get(0, 0)
	assert.Equal(t, 32, val)
}

func TestGetGroups(t *testing.T) {
	b := getGameBoard()

	result, err := b.FindGroup(7, 0)
	assert.NoError(t, err)
	assert.ElementsMatch(t, result, []Tile{CreateTile(7, 0), CreateTile(7, 1), CreateTile(7, 2)})

	result, err = b.FindGroup(7, 7)
	assert.NoError(t, err)
	assert.ElementsMatch(t, result, []Tile{CreateTile(7, 7), CreateTile(7, 6), CreateTile(7, 5)})

	result, err = b.FindGroup(7, 6)
	assert.NoError(t, err)
	assert.ElementsMatch(t, result, []Tile{CreateTile(7, 7), CreateTile(7, 6), CreateTile(7, 5)})

	result, err = b.FindGroup(0, 0)
	assert.NoError(t, err)
	assert.Empty(t, result)

	result, err = b.FindGroup(0, 1)
	assert.NoError(t, err)
	assert.Empty(t, result)

	result, err = b.FindGroup(0, 2)
	assert.NoError(t, err)
	assert.Empty(t, result)

}

func getGameBoard() MatriXBoard {

	data := make([]float64, 0)
	data = append(data, []float64{2, 8, 2, 4, 4, 4, 8, 16}...)
	data = append(data, []float64{4, 16, 4, 16, 16, 8, 16, 8}...)
	data = append(data, []float64{4, 8, 2, 8, 4, 16, 8, 16}...)
	data = append(data, []float64{2, 16, 4, 8, 2, 16, 8, 4}...)
	data = append(data, []float64{8, 16, 2, 2, 16, 4, 16, 8}...)
	data = append(data, []float64{4, 2, 8, 8, 2, 8, 4, 8}...)
	data = append(data, []float64{8, 2, 8, 2, 4, 2, 16, 4}...)
	data = append(data, []float64{16, 16, 16, 16, 8, 2, 2, 2}...)

	return MatriXBoard{
		m: mat.NewDense(width, width, data),
	}

}
