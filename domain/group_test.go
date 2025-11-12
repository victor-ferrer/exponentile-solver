package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup_GetReplacementValue(t *testing.T) {
	tests := []struct {
		name     string
		tiles    []Tile
		value    int
		expected int
	}{
		{
			name:     "less than 3 tiles returns 0",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1)},
			value:    16,
			expected: 0,
		},
		{
			name:     "empty group returns 0",
			tiles:    []Tile{},
			value:    16,
			expected: 0,
		},
		{
			name:     "single tile returns 0",
			tiles:    []Tile{CreateTile(0, 0)},
			value:    16,
			expected: 0,
		},
		{
			name:     "3 tiles with value 2 returns 4",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2)},
			value:    2,
			expected: 4,
		},
		{
			name:     "3 tiles with value 16 returns 32",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2)},
			value:    16,
			expected: 32,
		},
		{
			name:     "4 tiles with value 2 returns 8",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3)},
			value:    2,
			expected: 8,
		},
		{
			name:     "4 tiles with value 16 returns 64",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3)},
			value:    16,
			expected: 64,
		},
		{
			name:     "5 tiles with value 4 returns 32",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3), CreateTile(0, 4)},
			value:    4,
			expected: 32,
		},
		{
			name:     "6 tiles with value 8 returns 128",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3), CreateTile(0, 4), CreateTile(0, 5)},
			value:    8,
			expected: 128,
		},
		{
			name:     "8 tiles with value 2 returns 128",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3), CreateTile(0, 4), CreateTile(0, 5), CreateTile(0, 6), CreateTile(0, 7)},
			value:    2,
			expected: 128,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Tiles: tt.tiles,
				Value: tt.value,
			}
			result := g.GetReplacementValue()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGroup_GetScore(t *testing.T) {
	tests := []struct {
		name     string
		tiles    []Tile
		value    int
		expected int
	}{
		{
			name:     "3 tiles with value 2 returns 6",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2)},
			value:    2,
			expected: 6,
		},
		{
			name:     "4 tiles with value 16 returns 64",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3)},
			value:    16,
			expected: 64,
		},
		{
			name:     "5 tiles with value 8 returns 40",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3), CreateTile(0, 4)},
			value:    8,
			expected: 40,
		},
		{
			name:     "empty group returns 0",
			tiles:    []Tile{},
			value:    0,
			expected: 0,
		},
		{
			name:     "single tile returns tile value",
			tiles:    []Tile{CreateTile(0, 0)},
			value:    32,
			expected: 32,
		},
		{
			name:     "8 tiles with value 4 returns 32",
			tiles:    []Tile{CreateTile(0, 0), CreateTile(0, 1), CreateTile(0, 2), CreateTile(0, 3), CreateTile(0, 4), CreateTile(0, 5), CreateTile(0, 6), CreateTile(0, 7)},
			value:    4,
			expected: 32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Tiles: tt.tiles,
				Value: tt.value,
			}
			result := g.GetScore()
			assert.Equal(t, tt.expected, result)
		})
	}
}
