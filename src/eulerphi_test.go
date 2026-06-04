package mylib

import "testing"

func TestEulerPhi(t *testing.T) {
	tests := []struct {
		n, want int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 4},
		{6, 2},
		{7, 6},
		{8, 4},
		{9, 6},
		{10, 4},
		{12, 4},
		{100, 40},
	}
	for _, tt := range tests {
		got := EulerPhi(tt.n)
		if got != tt.want {
			t.Errorf("EulerPhi(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}
