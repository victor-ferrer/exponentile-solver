package domain

type TileState struct {
	Position Tile
	Value    int
}

type Group struct {
	Tiles []Tile
	Value int
}

type GameEvent struct {
	Type     string
	Sequence int
	Tiles    []TileState
	Score    int
	Group    Group
}

type Board interface {
	Get(x, y int) int
	MakeMove(t1, t2 Tile) []GameEvent
}
