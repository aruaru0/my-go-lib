package mylib

import "testing"

func TestLCASegSimple(t *testing.T) {
	N := 5
	node := make([][]int, N)
	node[0] = []int{1, 2}
	node[1] = []int{0, 3, 4}
	node[2] = []int{0}
	node[3] = []int{1}
	node[4] = []int{1}

	l := NewLCASeg(0, N, node)

	if got := l.LCA(3, 4); got != 1 {
		t.Errorf("LCA(3,4) = %d, want 1", got)
	}
	if got := l.LCA(3, 2); got != 0 {
		t.Errorf("LCA(3,2) = %d, want 0", got)
	}
}

func TestLCASegDist(t *testing.T) {
	N := 5
	node := make([][]int, N)
	node[0] = []int{1, 2}
	node[1] = []int{0, 3, 4}
	node[2] = []int{0}
	node[3] = []int{1}
	node[4] = []int{1}

	l := NewLCASeg(0, N, node)

	if got := l.Dist(3, 4); got != 2 {
		t.Errorf("Dist(3,4) = %d, want 2", got)
	}
	if got := l.Dist(3, 2); got != 3 {
		t.Errorf("Dist(3,2) = %d, want 3", got)
	}
}

func TestLCASegLine(t *testing.T) {
	N := 4
	node := make([][]int, N)
	node[0] = []int{1}
	node[1] = []int{0, 2}
	node[2] = []int{1, 3}
	node[3] = []int{2}

	l := NewLCASeg(0, N, node)

	if got := l.LCA(0, 3); got != 0 {
		t.Errorf("LCA(0,3) = %d, want 0", got)
	}
	if got := l.LCA(1, 3); got != 1 {
		t.Errorf("LCA(1,3) = %d, want 1", got)
	}
	if got := l.Dist(0, 3); got != 3 {
		t.Errorf("Dist(0,3) = %d, want 3", got)
	}
}
