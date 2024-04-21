package solver

type Tile struct {
	X, Y int
}

func CreateTile(row, col int) Tile {
	return Tile{
		X: row,
		Y: col,
	}
}
