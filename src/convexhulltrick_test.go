package mylib

import "testing"

func TestConvexHullTrick(t *testing.T) {
	pos := []int{-10, -5, 0, 3, 7, 10}
	cht := NewConvexHullTrick(pos)
	cht.AddLine(1, 0)  // y = x
	cht.AddLine(-1, 5) // y = -x + 5
	cht.AddLine(2, -3) // y = 2x - 3

	tests := []struct {
		x    int
		want int
	}{
		{-10, 15},  // y = -x + 5
		{-5, 10},   // y = -x + 5
		{0, 5},     // y = -x + 5
		{3, 3},     // y = x or y = 2x-3
		{7, 11},    // y = 2x - 3
		{10, 17},   // y = 2x - 3
	}
	for _, tt := range tests {
		got := cht.GetMax(tt.x)
		if got != tt.want {
			t.Errorf("GetMax(%d) = %d, want %d", tt.x, got, tt.want)
		}
	}
}

func TestConvexHullTrickSingleLine(t *testing.T) {
	pos := []int{0, 1, 2, 3, 4, 5}
	cht := NewConvexHullTrick(pos)
	cht.AddLine(3, 1)

	for _, x := range pos {
		want := 3*x + 1
		got := cht.GetMax(x)
		if got != want {
			t.Errorf("GetMax(%d) = %d, want %d", x, got, want)
		}
	}
}
