package mylib

import "testing"

func TestXorSum(t *testing.T) {
	tests := []struct {
		dist []int
		want int64
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1}, 0},
		{[]int{1, 2}, 3},
		{[]int{1, 1}, 0},
		{[]int{3, 3, 3}, 0},
	}
	for _, tt := range tests {
		got := XorSum(tt.dist)
		if got != tt.want {
			t.Errorf("XorSum(%v) = %d, want %d", tt.dist, got, tt.want)
		}
	}
}
