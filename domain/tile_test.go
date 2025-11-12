package domain

import "testing"

func TestTile_isContinous(t *testing.T) {
	tests := []struct {
		name     string
		t1       Tile
		t2       Tile
		expected bool
	}{
		{
			name:     "vertically contiguous - adjacent below",
			t1:       CreateTile(3, 4),
			t2:       CreateTile(4, 4),
			expected: true,
		},
		{
			name:     "vertically contiguous - adjacent above",
			t1:       CreateTile(4, 4),
			t2:       CreateTile(3, 4),
			expected: true,
		},
		{
			name:     "horizontally contiguous - adjacent right",
			t1:       CreateTile(5, 2),
			t2:       CreateTile(5, 3),
			expected: true,
		},
		{
			name:     "horizontally contiguous - adjacent left",
			t1:       CreateTile(5, 3),
			t2:       CreateTile(5, 2),
			expected: true,
		},
		{
			name:     "same column - 2 rows apart",
			t1:       CreateTile(1, 4),
			t2:       CreateTile(3, 4),
			expected: false,
		},
		{
			name:     "same column - far apart",
			t1:       CreateTile(0, 0),
			t2:       CreateTile(7, 0),
			expected: false,
		},
		{
			name:     "same row - 3 columns apart",
			t1:       CreateTile(2, 1),
			t2:       CreateTile(2, 4),
			expected: false,
		},
		{
			name:     "same row - far apart",
			t1:       CreateTile(5, 0),
			t2:       CreateTile(5, 7),
			expected: false,
		},
		{
			name:     "diagonal - down-right",
			t1:       CreateTile(2, 2),
			t2:       CreateTile(3, 3),
			expected: false,
		},
		{
			name:     "diagonal - down-left",
			t1:       CreateTile(4, 5),
			t2:       CreateTile(5, 4),
			expected: false,
		},
		{
			name:     "same tile",
			t1:       CreateTile(3, 3),
			t2:       CreateTile(3, 3),
			expected: false,
		},
		{
			name:     "completely different - opposite corners",
			t1:       CreateTile(0, 0),
			t2:       CreateTile(7, 7),
			expected: false,
		},
		{
			name:     "completely different - random positions",
			t1:       CreateTile(2, 5),
			t2:       CreateTile(6, 1),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.t1.isContinous(tt.t2)
			if result != tt.expected {
				t.Errorf("isContinous(%v, %v) = %v; want %v", tt.t1, tt.t2, result, tt.expected)
			}
		})
	}
}
