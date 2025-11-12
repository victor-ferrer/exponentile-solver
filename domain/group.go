package domain

type Group struct {
	Tiles []Tile
	Value int
}

func (g Group) GetReplacementValue() int {
	// Less than 3 tiles, 0 score
	if len(g.Tiles) < 3 {
		return 0
	}

	// Three tiles return the next power of 2 of the current value
	// Each additional tile in the group adds another power of 2
	steps := len(g.Tiles) - 2
	result := g.Value
	for range steps {
		result *= 2
	}

	return result
}

func (g Group) GetScore() int {
	return g.Value * len(g.Tiles)
}
