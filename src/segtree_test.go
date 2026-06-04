package mylib

import "testing"

func TestSegTreeSum(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	seg := NewSegTree(v, func() int { return 0 }, func(a, b int) int { return a + b })

	if got := seg.AllProd(); got != 15 {
		t.Errorf("AllProd() = %d, want 15", got)
	}
	if got := seg.Prod(0, 3); got != 6 {
		t.Errorf("Prod(0,3) = %d, want 6", got)
	}
	if got := seg.Prod(2, 5); got != 12 {
		t.Errorf("Prod(2,5) = %d, want 12", got)
	}
	if got := seg.Get(2); got != 3 {
		t.Errorf("Get(2) = %d, want 3", got)
	}
}

func TestSegTreeSet(t *testing.T) {
	v := []int{1, 2, 3, 4, 5}
	seg := NewSegTree(v, func() int { return 0 }, func(a, b int) int { return a + b })
	seg.Set(2, 10)
	if got := seg.AllProd(); got != 22 {
		t.Errorf("AllProd() after Set = %d, want 22", got)
	}
	if got := seg.Prod(0, 3); got != 13 {
		t.Errorf("Prod(0,3) after Set = %d, want 13", got)
	}
}

func TestSegTreeMaxRight(t *testing.T) {
	v := []int{2, 5, 3, 1, 4}
	seg := NewSegTree(v, func() int { return 0 }, func(a, b int) int { return a + b })
	// Prod(1,2) = 5 < 6, Prod(1,3) = 8 >= 6, so r=2 is rightmost where Prod(1,r) < 6
	got := seg.MaxRight(1, func(s int) bool { return s < 6 })
	want := 2
	if got != want {
		t.Errorf("MaxRight(1, s<6) = %d, want %d", got, want)
	}
}

func TestSegTreeMinLeft(t *testing.T) {
	v := []int{2, 5, 3, 1, 4}
	seg := NewSegTree(v, func() int { return 0 }, func(a, b int) int { return a + b })
	got := seg.MinLeft(4, func(s int) bool { return s < 6 })
	want := 2
	if got != want {
		t.Errorf("MinLeft(4, s<6) = %d, want %d", got, want)
	}
}

func TestSegTreeMin(t *testing.T) {
	v := []int{3, 1, 4, 1, 5}
	seg := NewSegTree(v, func() int { return 1 << 60 }, func(a, b int) int {
		if a < b {
			return a
		}
		return b
	})
	if got := seg.AllProd(); got != 1 {
		t.Errorf("AllProd() = %d, want 1", got)
	}
	if got := seg.Prod(0, 3); got != 1 {
		t.Errorf("Prod(0,3) = %d, want 1", got)
	}
}

func TestSegTreeString(t *testing.T) {
	v := []string{"a", "b", "c"}
	seg := NewSegTree(v, func() string { return "" }, func(a, b string) string { return a + b })
	if got := seg.AllProd(); got != "abc" {
		t.Errorf("AllProd() = %q, want %q", got, "abc")
	}
	if got := seg.Prod(0, 2); got != "ab" {
		t.Errorf("Prod(0,2) = %q, want %q", got, "ab")
	}
	seg.Set(1, "x")
	if got := seg.AllProd(); got != "axc" {
		t.Errorf("AllProd() after Set = %q, want %q", got, "axc")
	}
}
