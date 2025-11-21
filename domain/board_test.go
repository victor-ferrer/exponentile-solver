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

	// Row 7: [16, 16, 16, 16, 8, 2, 2, 2] - horizontal run of 4 sixteens
	result := b.findGroup(7, 0)
	assert.ElementsMatch(t, result.Tiles, []Tile{CreateTile(7, 0), CreateTile(7, 1), CreateTile(7, 2), CreateTile(7, 3)})

	// Row 7: last three tiles are 2s
	result = b.findGroup(7, 7)
	assert.ElementsMatch(t, result.Tiles, []Tile{CreateTile(7, 7), CreateTile(7, 6), CreateTile(7, 5)})

	result = b.findGroup(7, 6)
	assert.ElementsMatch(t, result.Tiles, []Tile{CreateTile(7, 7), CreateTile(7, 6), CreateTile(7, 5)})

	// No groups of 3+ for these tiles
	result = b.findGroup(0, 0)
	assert.Empty(t, result.Tiles)

	result = b.findGroup(0, 1)
	assert.Empty(t, result.Tiles)

	result = b.findGroup(0, 2)
	assert.Empty(t, result.Tiles)

}

func TestMakeMove_InvalidSwap(t *testing.T) {
	b := getGameBoard()

	events := b.MakeMove(CreateTile(0, 0), CreateTile(2, 2))
	assert.Len(t, events, 1)
	assert.Equal(t, EVENT_TYPE_NO_CHANGES, events[0].Type)
}

func TestMakeMove_NoGroupFormed(t *testing.T) {
	b := getGameBoard()

	events := b.MakeMove(CreateTile(0, 0), CreateTile(0, 1))
	assert.Len(t, events, 1)
	assert.Equal(t, EVENT_TYPE_NO_CHANGES, events[0].Type)
	assert.Equal(t, 0, events[0].Score)
	assert.Equal(t, 2, b.Get(0, 0))
	assert.Equal(t, 8, b.Get(0, 1))
}

func TestMakeMove_GroupFormed(t *testing.T) {
	b := getGameBoard()

	initialValue := b.Get(7, 2)
	assert.Equal(t, 16, initialValue)

	// Moving tile from (7,3) to (7,2)
	events := b.MakeMove(CreateTile(7, 3), CreateTile(7, 2))
	// May have cascade events, but first event is the initial group
	assert.GreaterOrEqual(t, len(events), 1)
	assert.Equal(t, EVENT_TYPE_GAME_UPDATED, events[0].Type)

	// Row 7 has 4 contiguous 16s
	// The moved tile at (7,2) should be kept and upgraded
	movedTileValue := b.Get(7, 2)
	assert.Equal(t, 64, movedTileValue)

	// Score should be the sum of the 4 tiles (16+16+16+16 = 64)
	assert.Equal(t, 64, events[0].Score)
}

func TestCalculateGroupScore(t *testing.T) {

	// Test with 4 tiles of value 16
	group := Group{[]Tile{CreateTile(7, 0), CreateTile(7, 1), CreateTile(7, 2), CreateTile(7, 3)}, 16}
	score := group.GetScore()
	assert.Equal(t, 64, score)

	// Test with 3 tiles of value 2
	group = Group{[]Tile{CreateTile(7, 5), CreateTile(7, 6), CreateTile(7, 7)}, 2}
	score = group.GetScore()
	assert.Equal(t, 6, score)

	// Test with empty group
	group = Group{[]Tile{}, 0}
	score = group.GetScore()
	assert.Equal(t, 0, score)
}

func TestMakeMove_ScoreIncrement(t *testing.T) {
	b := getGameBoard()

	// First move: Row 7, 4 tiles of value 16 (score = 64)
	events := b.MakeMove(CreateTile(7, 3), CreateTile(7, 2))
	assert.GreaterOrEqual(t, len(events), 1)
	assert.Equal(t, 64, events[0].Score)
	assert.Equal(t, 1, events[0].Sequence)
	// Score is cumulative, including any cascades
	assert.GreaterOrEqual(t, b.score, 64)

	// Second move
	events = b.MakeMove(CreateTile(7, 6), CreateTile(7, 5))
	assert.Equal(t, 2, events[0].Sequence)
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
