package mylib

import (
	"testing"
)

func TestCeilPow2(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{8, 3},
		{9, 4},
	}
	for _, tt := range tests {
		if got := CeilPow2(tt.n); got != tt.want {
			t.Errorf("CeilPow2(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestBsf(t *testing.T) {
	tests := []struct {
		n    uint
		want int
	}{
		{1, 0},
		{2, 1},
		{4, 2},
		{8, 3},
		{16, 4},
		{6, 1},
	}
	for _, tt := range tests {
		if got := Bsf(tt.n); got != tt.want {
			t.Errorf("Bsf(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
