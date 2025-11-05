package domain

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
	b.swap(CreateTile(3, 0), CreateTile(4, 0))

	val := b.Get(3, 0)
	assert.Equal(t, 8, val)

	val = b.Get(4, 0)
	assert.Equal(t, 2, val)

}

func TestDrop(t *testing.T) {
	// TODO this could be a table test
	b := getGameBoard()

	b.dropTile(CreateTile(7, 0), 32)
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

	result := b.findGroup(7, 0)
	assert.ElementsMatch(t, result, []Tile{CreateTile(7, 0), CreateTile(7, 1), CreateTile(7, 2)})

	result = b.findGroup(7, 7)
	assert.ElementsMatch(t, result, []Tile{CreateTile(7, 7), CreateTile(7, 6), CreateTile(7, 5)})

	result = b.findGroup(7, 6)
	assert.ElementsMatch(t, result, []Tile{CreateTile(7, 7), CreateTile(7, 6), CreateTile(7, 5)})

	result = b.findGroup(0, 0)
	assert.Empty(t, result)

	result = b.findGroup(0, 1)
	assert.Empty(t, result)

	result = b.findGroup(0, 2)
	assert.Empty(t, result)

}

func TestMakeMove_InvalidSwap(t *testing.T) {
	b := getGameBoard()

	evt := b.MakeMove(CreateTile(0, 0), CreateTile(2, 2))
	assert.Equal(t, EVENT_TYPE_NO_CHANGES, evt.Type)
}

func TestMakeMove_NoGroupFormed(t *testing.T) {
	b := getGameBoard()

	evt := b.MakeMove(CreateTile(0, 0), CreateTile(0, 1))
	assert.Equal(t, EVENT_TYPE_NO_CHANGES, evt.Type)
	assert.Equal(t, 0, evt.Score)
	assert.Equal(t, 2, b.Get(0, 0))
	assert.Equal(t, 8, b.Get(0, 1))
}

func TestMakeMove_GroupFormed(t *testing.T) {
	b := getGameBoard()

	initialValue := b.Get(7, 2)
	assert.Equal(t, 16, initialValue)

	evt := b.MakeMove(CreateTile(7, 3), CreateTile(7, 2))
	assert.Equal(t, EVENT_TYPE_GAME_UPDATED, evt.Type)
	
	middleTileValue := b.Get(7, 1)
	assert.Equal(t, 32, middleTileValue)
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
