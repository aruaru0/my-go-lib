package mylib

import "testing"

func TestCountDivisibleSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		p    int
		want int
	}{
		{"123", 1, 6},
		{"123", 2, 2},
		{"123", 3, 3},
		{"123", 5, 0},
		{"181", 2, 2},
		{"181", 5, 0},
		{"1234", 3, 4},
		{"0", 2, 1},
		{"00", 2, 3},
		{"111", 3, 1},
	}
	for _, tt := range tests {
		got := CountDivisibleSubstrings(tt.s, tt.p)
		if got != tt.want {
			t.Errorf("CountDivisibleSubstrings(%q, %d) = %d, want %d", tt.s, tt.p, got, tt.want)
		}
	}
}
