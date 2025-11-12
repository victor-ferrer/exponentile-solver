package domain

type Tile struct {
	X, Y int
}

func CreateTile(row, col int) Tile {
	return Tile{
		X: row,
		Y: col,
	}
}

func (t1 Tile) isContinous(t2 Tile) bool {

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
