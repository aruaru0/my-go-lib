package mylib

import "testing"

func TestLis(t *testing.T) {
	tests := []struct {
		in   []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7}, 1},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{5, 4, 3, 2, 1}, 1},
	}
	for _, tt := range tests {
		got := Lis(tt.in)
		if got != tt.want {
			t.Errorf("Lis(%v) = %d, want %d", tt.in, got, tt.want)
		}
	}
}

func TestLowerBound(t *testing.T) {
	a := []int{1, 3, 5, 7, 9}
	tests := []struct {
		x    int
		want int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{5, 2},
		{6, 3},
		{9, 4},
		{10, 5},
	}
	for _, tt := range tests {
		got := LowerBound(a, tt.x)
		if got != tt.want {
			t.Errorf("LowerBound(%v, %d) = %d, want %d", a, tt.x, got, tt.want)
		}
	}
}
