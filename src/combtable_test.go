package mylib

import "testing"

func TestCombTable(t *testing.T) {
	ct := NewCombTable[int](10)
	tests := []struct {
		n, k, want int
	}{
		{5, 2, 10},
		{5, 0, 1},
		{5, 5, 1},
		{10, 5, 252},
		{0, 0, 1},
		{10, 11, 0},
		{3, 1, 3},
		{6, 3, 20},
	}
	for _, tt := range tests {
		got := ct.NCk(tt.n, tt.k)
		if got != tt.want {
			t.Errorf("NCk(%d, %d) = %d, want %d", tt.n, tt.k, got, tt.want)
		}
	}
}

func TestCombTableUint64(t *testing.T) {
	ct := NewCombTable[uint64](10)
	got := ct.NCk(5, 2)
	if got != 10 {
		t.Errorf("CombTable[uint64].NCk(5,2) = %d, want %d", got, 10)
	}
}
