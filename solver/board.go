package solver

import (
	"fmt"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

type Board interface {
	// NewBoard() Board
	Swap(t1, t2 Tile)
	Get(x, y int) int
	DropTile(t Tile, newValue int)
	Render()
	FindGroup(x, y int) ([]Tile, error)
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

	b.m.Trace()
	for row := 0; row < width; row++ {
		for column := 0; column < width; column++ {
			fmt.Printf("%02f ", b.m.At(row, column))
		}
		fmt.Println()
	}
}

// Find Group: Finds tiles with the same value in the same row
func (b MatriXBoard) FindGroup(x, y int) ([]Tile, error) {
	// targetVal := b.B.At(x, y)

	// leftTiles, err := b.getMatchingTilesToTheLeft(targetVal, x, y)
	// rightTiles, err := b.getMatchingTilesToTheRight(targetVal, x, y)

	// return append(leftTiles, CreateTile(x, y), rightTiles), nil
	return nil, nil

}
