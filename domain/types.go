package domain

type GameEvent struct {
	Type         string
	Board        Board
	Score        int
	GroupedTiles []Tile
}

type Board interface {
	Get(x, y int) int
	MakeMove(t1, t2 Tile) GameEvent
}
