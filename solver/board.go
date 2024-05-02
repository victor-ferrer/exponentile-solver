package solver

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

type Board interface {
	Swap(t1, t2 Tile)
	Get(x, y int) int
	DropTile(t Tile, newValue int)
	Render()
	FindGroup(x, y int) []Tile
	MakeMove(t1, t2 Tile) ([]Tile, int)
}

type MatriXBoard struct {
	m *mat.Dense
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

func (b MatriXBoard) Swap(t1, t2 Tile) {
	aux := b.m.At(t1.X, t1.Y)

	b.m.Set(t1.X, t1.Y, b.m.At(t2.X, t2.Y))
	b.m.Set(t2.X, t2.Y, aux)
}

func (b MatriXBoard) DropTile(target Tile, newValue int) {
	col := target.Y
	for row := target.X; row > 0; row-- {
		b.Swap(CreateTile(row, col), CreateTile(row-1, col))
	}
	b.m.Set(0, col, float64(newValue))

}

func (b MatriXBoard) Get(x, y int) int {
	return int(b.m.At(x, y))
}

func (b MatriXBoard) Render() {

	for row := 0; row < width; row++ {
		for column := 0; column < width; column++ {
			fmt.Printf("%02d ", int(b.m.At(row, column)))
		}
		fmt.Println()
	}
}

// Find Group: Finds tiles with the same value in the same row
// TODO: Groups of three in the same row only
func (b MatriXBoard) FindGroup(x, y int) []Tile {

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

func (b MatriXBoard) MakeMove(t1, t2 Tile) ([]Tile, int) {
	// TODO
	// If the swap is valid (contiguos tiles)
	if !AreTilesContiguous(t1, t2) {
		return []Tile{}, 0
	}

	b.Swap(t1, t2)

	// Find if there is a group for the tile moved first
	groups := b.FindGroup(t1.X, t1.Y)

	if len(groups) > 0 {
		return []Tile{}, 0
	}

	// If there is a group:
	//  - Calculate the replacement Tile (next power of 2)
	next := GetNextPowerOf2(b.Get(t1.X, t1.Y))
	b.m.Set(t1.X, t1.Y, float64(next))

	//  Drop the rest two tiles and replace them with random tiles
	for _, tileToDrop := range groups {
		if tileToDrop != t1 {
			b.DropTile(tileToDrop, GetSeqNumber(rand.Intn(5)+1))
		}
	}

	//  Return the score increase
	return groups, next

}
