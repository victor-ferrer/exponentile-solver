package domain

type Group struct {
	Tiles []Tile
	Value int
}

func (g Group) GetReplacementValue() int {
	// FIX ME
	return 0
}

func (g Group) GetScore() int {
	score := 0
	for _, _ = range g.Tiles {
		score += g.Value
	}
	return score
}
