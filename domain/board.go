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

// Find Group: Finds all tiles with the same value in straight lines (row and/or column)
// Returns groups of 3 or more contiguous tiles
// If both horizontal and vertical runs exist, returns both
func (b MatriXBoard) findGroup(x, y int) []Tile {
	val := b.m.At(x, y)
	result := []Tile{}

	// Find horizontal run extents
	left, right := y, y
	for left-1 >= 0 && b.m.At(x, left-1) == val {
		left--
	}
	for right+1 < width && b.m.At(x, right+1) == val {
		right++
	}
	hlen := right - left + 1

	// Find vertical run extents
	up, down := x, x
	for up-1 >= 0 && b.m.At(up-1, y) == val {
		up--
	}
	for down+1 < width && b.m.At(down+1, y) == val {
		down++
	}
	vlen := down - up + 1

	// Add horizontal run if >= 3
	if hlen >= 3 {
		for col := left; col <= right; col++ {
			result = append(result, CreateTile(x, col))
		}
	}

	// Add vertical run if >= 3
	if vlen >= 3 {
		for row := up; row <= down; row++ {
			tile := CreateTile(row, y)
			// Avoid duplicating the center tile if both runs exist
			if hlen < 3 || row != x {
				result = append(result, tile)
			}
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

	// Keep the middle tile and drop all others with random tiles
	middleIndex := len(group) / 2
	middleTile := group[middleIndex]
	b.m.Set(middleTile.X, middleTile.Y, float64(nextValue))

	// Drop all other tiles and replace with random tiles
	for i, tile := range group {
		if i != middleIndex {
			b.dropTile(tile, GetSeqNumber(rand.Intn(5)+1))
		}
	}

	// Calculate score increase
	score := currentValue * len(group)

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
