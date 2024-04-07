package solver

import (
	"fmt"
	"math"
	"math/rand"
)

type Board [][]int
type Tile struct {
	X, Y int
}

func GetSeqNumber(n int) int {
	return int(math.Pow(2, float64(n)))
}

func CreateTile(row, col int) Tile {
	return Tile{
		X: row,
		Y: col,
	}
}

const width = 8

func NewBoard() Board {

	b := make([][]int, width)

	for rows := 0; rows < width; rows++ {
		row := make([]int, width)
		for cols := 0; cols < width; cols++ {
			row[cols] = GetSeqNumber(rand.Intn(5) + 1)
		}
		b[rows] = row
	}
	return b
}

func (b Board) Swap(t1, t2 Tile) {
	aux, _ := b.Get(t1.X, t1.Y) // FIXME

	val, _ := b.Get(t2.X, t2.Y)
	b.set(t1.X, t1.Y, val)
	b.set(t2.X, t2.Y, aux)

}

func (b Board) DropTile(target Tile, newValue int) {
	col := target.Y
	for row := target.X; row > 0; row-- {
		b.Swap(CreateTile(row, col), CreateTile(row-1, col))
	}
	b[0][col] = newValue

}

func (b Board) Render() {
	for row := 0; row < width; row++ {
		for column := 0; column < width; column++ {
			fmt.Printf("%02d ", b[row][column])
		}
		fmt.Println()
	}
}

func (b Board) Get(x, y int) (int, error) {

	if x >= width || x < 0 {
		return 0, fmt.Errorf("getting tile value: invalid row: %d", x)
	}

	if y >= width || y < 0 {
		return 0, fmt.Errorf("getting tile value: invalid column: %d", x)
	}

	return b[x][y], nil
}

func (b Board) set(x, y, value int) {
	b[x][y] = value
}
