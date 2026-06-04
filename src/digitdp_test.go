package mylib

import "testing"

func TestCountContaining(t *testing.T) {
	tests := []struct {
		n, target int
		want      int64
	}{
		{1, 3, 0},
		{3, 3, 1},
		{10, 3, 1},
		{30, 3, 4},
		{100, 3, 19},
		{0, 3, 0},
	}
	for _, tt := range tests {
		got := CountContaining(tt.n, tt.target)
		if got != tt.want {
			t.Errorf("CountContaining(%d,%d) = %d, want %d", tt.n, tt.target, got, tt.want)
		}
	}
}

func TestCountNonZero(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int64
	}{
		{"1", 0, 0},
		{"1", 1, 1},
		{"5", 1, 5},
		{"9", 1, 9},
		{"10", 1, 10},
		{"10", 2, 0},
		{"20", 1, 11},
		{"100", 1, 19},
		{"100", 2, 81},
	}
	for _, tt := range tests {
		got := CountNonZero(tt.s, tt.k)
		if got != tt.want {
			t.Errorf("CountNonZero(%q,%d) = %d, want %d", tt.s, tt.k, got, tt.want)
		}
	}
}
