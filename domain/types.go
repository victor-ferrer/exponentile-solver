package domain

type TileState struct {
	Position Tile
	Value    int
}

type GameEvent struct {
	Type     string
	Sequence int
	Tiles    []TileState
	Score    int
	Group    Group
}

type Board interface {
	GetTileState() []TileState
	MakeMove(t1, t2 Tile) []GameEvent
}
