package mylib

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		s, t string
		want int
	}{
		{"", "", 0},
		{"a", "", 1},
		{"", "a", 1},
		{"abc", "abc", 0},
		{"kitten", "sitting", 3},
		{"saturday", "sunday", 3},
		{"book", "back", 2},
		{"abc", "def", 3},
		{"flaw", "lawn", 2},
	}
	for _, tt := range tests {
		got := LevenshteinDistance(tt.s, tt.t)
		if got != tt.want {
			t.Errorf("LevenshteinDistance(%q, %q) = %d, want %d", tt.s, tt.t, got, tt.want)
		}
	}
}

func TestLevenshteinDistanceReversed(t *testing.T) {
	d1 := LevenshteinDistance("abc", "def")
	d2 := LevenshteinDistance("def", "abc")
	if d1 != d2 {
		t.Errorf("LevenshteinDistance should be symmetric: got %d vs %d", d1, d2)
	}
}
