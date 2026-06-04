package mylib

import "testing"

func TestBITAddAndSum(t *testing.T) {
	b := NewBIT(5)
	b.Add(0, 3)
	b.Add(2, 5)
	b.Add(4, 2)

	tests := []struct {
		i    int
		want int
	}{
		{0, 3},
		{1, 3},
		{2, 8},
		{3, 8},
		{4, 10},
	}
	for _, tt := range tests {
		got := b.Sum(tt.i)
		if got != tt.want {
			t.Errorf("Sum(%d) = %d, want %d", tt.i, got, tt.want)
		}
	}
}

func TestBITRangeSum(t *testing.T) {
	b := NewBIT(5)
	for i := 0; i < 5; i++ {
		b.Add(i, i+1)
	}

	tests := []struct {
		l, r, want int
	}{
		{0, 5, 15},
		{0, 3, 6},
		{1, 4, 9},
		{2, 2, 0},
		{3, 5, 9},
	}
	for _, tt := range tests {
		got := b.RangeSum(tt.l, tt.r)
		if got != tt.want {
			t.Errorf("RangeSum(%d, %d) = %d, want %d", tt.l, tt.r, got, tt.want)
		}
	}
}

func TestBITSumNegativeIndex(t *testing.T) {
	b := NewBIT(3)
	b.Add(0, 10)
	if got := b.Sum(-1); got != 0 {
		t.Errorf("Sum(-1) = %d, want 0", got)
	}
}

func TestBITSumOutOfRange(t *testing.T) {
	b := NewBIT(3)
	b.Add(0, 1)
	b.Add(1, 2)
	b.Add(2, 3)
	if got := b.Sum(10); got != 6 {
		t.Errorf("Sum(10) = %d, want 6", got)
	}
}
