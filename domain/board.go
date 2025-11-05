package domain

import (
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

const (
	EVENT_TYPE_GAME_OVER    = "GAME_OVER"
	EVENT_TYPE_GAME_UPDATED = "GAME_UPDATED"
	EVENT_TYPE_NO_CHANGES   = "NO_CHANGES"
)

type GameEvent struct {
	Type  string
	Board Board
	Score int
}

type Board interface {
	Get(x, y int) int
	MakeMove(t1, t2 Tile) GameEvent
}

type MatriXBoard struct {
	m     *mat.Dense
	score int
}

const width = 8

func NewBoard() MatriXBoard {

	data := make([]float64, width*width)
	for i := range data {
		data[i] = float64(GetSeqNumber(rand.Intn(5) + 1))
	}
	return MatriXBoard{
		m: mat.NewDense(width, width, data),
	}

}

func (b MatriXBoard) swap(t1, t2 Tile) GameEvent {
	aux := b.m.At(t1.X, t1.Y)

	b.m.Set(t1.X, t1.Y, b.m.At(t2.X, t2.Y))
	b.m.Set(t2.X, t2.Y, aux)

	return GameEvent{
		Type:  EVENT_TYPE_GAME_UPDATED,
		Board: b,
		Score: b.score,
	}
}

func (b MatriXBoard) dropTile(target Tile, newValue int) {
	col := target.Y
	for row := target.X; row > 0; row-- {
		b.swap(CreateTile(row, col), CreateTile(row-1, col))
	}
	b.m.Set(0, col, float64(newValue))

}

func (b MatriXBoard) Get(x, y int) int {
	return int(b.m.At(x, y))
}

// Find Group: Finds tiles with the same value in the same row
// TODO: Groups of three in the same row only
func (b MatriXBoard) findGroup(x, y int) []Tile {

	row := b.m.RowView(x)
	val := b.m.At(x, y)

	result := []Tile{}

	if y == width-1 {
		// Look for tiles to the left
		if row.AtVec(y-1) == val && row.AtVec(y-2) == val {
			result = append(result, CreateTile(x, y), CreateTile(x, y-1), CreateTile(x, y-2))
		}

	}
	if y < width-1 && y > 0 {
		// Look on both sides
		if row.AtVec(y-1) == val && row.AtVec(y+1) == val {
			result = append(result, CreateTile(x, y), CreateTile(x, y-1), CreateTile(x, y+1))
		}
	}
	if y == 0 {
		// Look for tiles to the right
		if row.AtVec(y+1) == val && row.AtVec(y+2) == val {
			result = append(result, CreateTile(x, y), CreateTile(x, y+1), CreateTile(x, y+2))

		}
	}

	return result

}

func (b MatriXBoard) MakeMove(t1, t2 Tile) GameEvent {
	// Check if the swap is valid (contiguous tiles)
	if !areContiguous(t1, t2) {
		return GameEvent{
			Board: b,
			Type:  EVENT_TYPE_NO_CHANGES,
			Score: b.score,
		}
	}

	// Swap the tiles
	b.swap(t1, t2)

	// Find if there is a group for the tile moved first (now at t2 position)
	group := b.findGroup(t2.X, t2.Y)

	// If there is no group, check the other tile
	if len(group) == 0 {
		group = b.findGroup(t1.X, t1.Y)
	}

	// If no group was found, swap back and return no changes
	if len(group) == 0 {
		b.swap(t1, t2)
		return GameEvent{
			Board: b,
			Type:  EVENT_TYPE_NO_CHANGES,
			Score: b.score,
		}
	}

	// Calculate the replacement value (next power of 2)
	currentValue := b.Get(group[0].X, group[0].Y)
	nextValue := currentValue * 2

	// Keep the middle tile and drop the other two with random tiles
	middleTile := group[1]
	b.m.Set(middleTile.X, middleTile.Y, float64(nextValue))

	// Drop the first and last tiles and replace with random tiles
	b.dropTile(group[0], GetSeqNumber(rand.Intn(5)+1))
	b.dropTile(group[2], GetSeqNumber(rand.Intn(5)+1))

	// Calculate score increase
	score := currentValue * 3

	return GameEvent{
		Board: b,
		Type:  EVENT_TYPE_GAME_UPDATED,
		Score: score,
	}
}

func areContiguous(t1, t2 Tile) bool {
	// Check if tiles are adjacent horizontally
	if t1.X == t2.X && (t1.Y == t2.Y+1 || t1.Y == t2.Y-1) {
		return true
	}
	// Check if tiles are adjacent vertically
	if t1.Y == t2.Y && (t1.X == t2.X+1 || t1.X == t2.X-1) {
		return true
	}
	return false
}
