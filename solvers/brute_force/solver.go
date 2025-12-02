package bruteforce

import "victor-ferrer/solver/domain"

func Solve(board domain.Board) []domain.GameEvent {

	state := board.GetTileState()
	events := []domain.GameEvent{}

	for {
		from, to, err := findMove(state)

		if err != nil {
			break
		}

		events = append(events, board.MakeMove(from, to)...)
	}

	return append(events, domain.GameEvent{
		Type:  domain.EVENT_TYPE_GAME_OVER,
		Tiles: state,
	})

}

func findMove(state []domain.TileState) (domain.Tile, domain.Tile, error) {
	panic("unimplemented")
}
