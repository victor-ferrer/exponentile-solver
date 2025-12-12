package bruteforce

import "victor-ferrer/solver/domain"

func Solve(board domain.Board) chan (domain.GameEvent) {

	events := make(chan domain.GameEvent)

	go func() {
		const width = 8

		for {
			foundMove := false

			// Try making moves from left to right, bottom to top
			for x := width - 1; x >= 0; x-- {
				for y := 0; y < width; y++ {
					t1 := domain.CreateTile(x, y)

					// Try swapping with right neighbor
					if y+1 < width {
						t2 := domain.CreateTile(x, y+1)
						moveEvents := board.MakeMove(t1, t2)
						for _, event := range moveEvents {
							if event.Type == domain.EVENT_TYPE_NO_CHANGES {
								continue
							}

							events <- event
							foundMove = true

						}
						if foundMove {
							break
						}
					}

					// Try swapping with above neighbor
					if x-1 >= 0 {
						t2 := domain.CreateTile(x-1, y)
						moveEvents := board.MakeMove(t1, t2)
						for _, event := range moveEvents {
							if event.Type == domain.EVENT_TYPE_NO_CHANGES {
								continue
							}

							events <- event
							foundMove = true
						}
						if foundMove {
							break
						}
					}
				}
				if foundMove {
					break
				}
			}

			if !foundMove {
				break
			}
		}

		close(events)

	}()

	return events

}
