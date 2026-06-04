package mylib

import "testing"

func TestFenwickTreeAddAndSum(t *testing.T) {
	f := NewFenwickTree[int](5)
	f.Add(0, 3)
	f.Add(2, 5)
	f.Add(4, 2)

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
		got := f.Sum(tt.i)
		if got != tt.want {
			t.Errorf("Sum(%d) = %d, want %d", tt.i, got, tt.want)
		}
	}
}

func TestFenwickTreeRangeSum(t *testing.T) {
	f := NewFenwickTree[int](5)
	for i := 0; i < 5; i++ {
		f.Add(i, i+1)
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
		got := f.RangeSum(tt.l, tt.r)
		if got != tt.want {
			t.Errorf("RangeSum(%d, %d) = %d, want %d", tt.l, tt.r, got, tt.want)
		}
	}
}

func TestFenwickTreeLowerBound(t *testing.T) {
	f := NewFenwickTree[int](5)
	for i := 0; i < 5; i++ {
		f.Add(i, 10)
	}

	tests := []struct {
		x    int
		want int
	}{
		{1, 0},
		{10, 0},
		{11, 1},
		{20, 1},
		{21, 2},
		{50, 4},
		{51, 5},
	}
	for _, tt := range tests {
		got := f.LowerBound(tt.x)
		if got != tt.want {
			t.Errorf("LowerBound(%d) = %d, want %d", tt.x, got, tt.want)
		}
	}
}

func TestFenwickTreeWithUint(t *testing.T) {
	f := NewFenwickTree[uint](3)
	f.Add(0, 10)
	f.Add(1, 20)
	f.Add(2, 30)

	if got := f.Sum(2); got != uint(60) {
		t.Errorf("Sum(2) = %d, want 60", got)
	}
	if got := f.RangeSum(0, 2); got != uint(30) {
		t.Errorf("RangeSum(0,2) = %d, want 30", got)
	}
}

func TestFenwickTreeLowerBoundUint(t *testing.T) {
	f := NewFenwickTree[uint](4)
	f.Add(0, 5)
	f.Add(1, 5)
	f.Add(2, 5)
	f.Add(3, 5)

	if got := f.LowerBound(12); got != 2 {
		t.Errorf("LowerBound(12) = %d, want 2", got)
	}
}
