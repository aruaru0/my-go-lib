package mylib

import "testing"

func TestNCR(t *testing.T) {
	nc := NewNCR(1000000007, 100)
	tests := []struct {
		name string
		fn   func(int, int) int
		n, k int
		want int
	}{
		{"N(5,2)", nc.N, 5, 2, 10},
		{"N(5,0)", nc.N, 5, 0, 1},
		{"N(5,5)", nc.N, 5, 5, 1},
		{"N(10,5)", nc.N, 10, 5, 252},
		{"P(5,2)", nc.P, 5, 2, 20},
		{"P(5,0)", nc.P, 5, 0, 1},
		{"P(5,5)", nc.P, 5, 5, 120},
		{"H(5,2)", nc.H, 5, 2, 15},
		{"H(5,0)", nc.H, 5, 0, 1},
		{"H(0,0)", nc.H, 0, 0, 1},
		{"N(5,6)", nc.N, 5, 6, 0},
		{"P(5,6)", nc.P, 5, 6, 0},
		{"H(3,3)", nc.H, 3, 3, 10},
	}
	for _, tt := range tests {
		got := tt.fn(tt.n, tt.k)
		if got != tt.want {
			t.Errorf("%s = %d, want %d", tt.name, got, tt.want)
		}
	}
}

func TestNCRModConfigurable(t *testing.T) {
	nc := NewNCR(13, 10)
	got := nc.N(5, 2)
	if got != 10 {
		t.Errorf("N(5,2) mod 13 = %d, want %d", got, 10)
	}
	got = nc.N(10, 5)
	if got != 252%13 {
		t.Errorf("N(10,5) mod 13 = %d, want %d", got, 252%13)
	}
}
