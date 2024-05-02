package solver

import "math"

type Tile struct {
	X, Y int
}

func CreateTile(row, col int) Tile {
	return Tile{
		X: row,
		Y: col,
	}
}

func AreTilesContiguous(t1, t2 Tile) bool {
	if t1.X == t2.X {
		return math.Abs(float64(t1.Y-t2.Y)) == 1
	}
	if t1.Y == t2.Y {
		return math.Abs(float64(t1.X-t2.X)) == 1
	}
	return false

}
