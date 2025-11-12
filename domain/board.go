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

type MatriXBoard struct {
	m        *mat.Dense
	score    int
	sequence int
}

const width = 8

func NewBoard() MatriXBoard {

	data := make([]float64, width*width)
	for i := range data {
		data[i] = float64(getSeqNumber(rand.Intn(5) + 1))
	}
	return MatriXBoard{
		m: mat.NewDense(width, width, data),
	}

}

func (b *MatriXBoard) swap(t1, t2 Tile) {
	aux := b.m.At(t1.X, t1.Y)
	b.m.Set(t1.X, t1.Y, b.m.At(t2.X, t2.Y))
	b.m.Set(t2.X, t2.Y, aux)
}

func (b *MatriXBoard) dropTile(target Tile, newValue int) {
	col := target.Y
	for row := target.X; row > 0; row-- {
		b.swap(CreateTile(row, col), CreateTile(row-1, col))
	}
	b.m.Set(0, col, float64(newValue))

}

func (b *MatriXBoard) Get(x, y int) int {
	return int(b.m.At(x, y))
}

func (b *MatriXBoard) getTileStates() []TileState {
	tiles := make([]TileState, 0, width*width)
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			tiles = append(tiles, TileState{
				Position: CreateTile(x, y),
				Value:    b.Get(x, y),
			})
		}
	}
	return tiles
}

// Find Group: Finds all tiles with the same value in straight lines (row and/or column)
// Returns groups of 3 or more contiguous tiles
// If both horizontal and vertical runs exist, returns both
func (b *MatriXBoard) findGroup(x, y int) Group {
	val := b.m.At(x, y)
	result := Group{Value: int(val)}

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
			result.Tiles = append(result.Tiles, CreateTile(x, col))
		}
	}

	// Add vertical run if >= 3
	if vlen >= 3 {
		for row := up; row <= down; row++ {
			tile := CreateTile(row, y)
			// Avoid duplicating the center tile if both runs exist
			if hlen < 3 || row != x {
				result.Tiles = append(result.Tiles, tile)
			}
		}
	}

	return result
}

func (b *MatriXBoard) MakeMove(t1, t2 Tile) []GameEvent {
	b.sequence++

	// Check if the swap is valid (contiguous tiles)
	if !t1.isContinous(t2) {
		return []GameEvent{{
			Type:     EVENT_TYPE_NO_CHANGES,
			Sequence: b.sequence,
			Tiles:    b.getTileStates(),
			Score:    b.score,
		}}
	}

	// Swap the tiles
	b.swap(t1, t2)

	// Find if there is a group for the tile moved first (now at t2 position)
	group := b.findGroup(t2.X, t2.Y)

	// If there is no group, check the other tile
	if len(group.Tiles) == 0 {
		group = b.findGroup(t1.X, t1.Y)
	}

	// If no group was found, swap back and return no changes
	if len(group.Tiles) == 0 {
		b.swap(t1, t2)
		return []GameEvent{{
			Type:     EVENT_TYPE_NO_CHANGES,
			Sequence: b.sequence,
			Tiles:    b.getTileStates(),
			Score:    b.score,
		}}
	}

	// Calculate and increment score before modifying the board
	scoreIncrease := group.GetScore()
	b.score += scoreIncrease

	// Calculate the replacement value (next power of 2)
	nextValue := group.GetReplacementValue()

	// Determine which tile to keep (the one that was moved to create the group)
	// First check t2 (where t1 was moved to), then t1
	var keptTile Tile
	keptTileFound := false
	for _, tile := range group.Tiles {
		if tile.X == t2.X && tile.Y == t2.Y {
			keptTile = tile
			keptTileFound = true
			break
		}
	}
	if !keptTileFound {
		for _, tile := range group.Tiles {
			if tile.X == t1.X && tile.Y == t1.Y {
				keptTile = tile
				keptTileFound = true
				break
			}
		}
	}

	// Upgrade the kept tile
	b.m.Set(keptTile.X, keptTile.Y, float64(nextValue))

	// Drop all other tiles and replace with random tiles
	for _, tile := range group.Tiles {
		if tile.X != keptTile.X || tile.Y != keptTile.Y {
			b.dropTile(tile, getSeqNumber(rand.Intn(5)+1))
		}
	}

	return []GameEvent{{
		Type:     EVENT_TYPE_GAME_UPDATED,
		Sequence: b.sequence,
		Tiles:    b.getTileStates(),
		Score:    b.score,
		Group:    group,
	}}
}
