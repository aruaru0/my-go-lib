package mylib

import "testing"

func TestLCASimple(t *testing.T) {
	l := NewLCA(5)
	l.AddEdge(0, 1)
	l.AddEdge(0, 2)
	l.AddEdge(1, 3)
	l.AddEdge(1, 4)
	l.Build(0)

	if got := l.LCA(3, 4); got != 1 {
		t.Errorf("LCA(3,4) = %d, want 1", got)
	}
	if got := l.LCA(3, 2); got != 0 {
		t.Errorf("LCA(3,2) = %d, want 0", got)
	}
	if got := l.LCA(0, 1); got != 0 {
		t.Errorf("LCA(0,1) = %d, want 0", got)
	}
	if got := l.LCA(3, 0); got != 0 {
		t.Errorf("LCA(3,0) = %d, want 0", got)
	}
}

func TestLCADist(t *testing.T) {
	l := NewLCA(5)
	l.AddEdge(0, 1)
	l.AddEdge(0, 2)
	l.AddEdge(1, 3)
	l.AddEdge(1, 4)
	l.Build(0)

	if got := l.Dist(3, 4); got != 2 {
		t.Errorf("Dist(3,4) = %d, want 2", got)
	}
	if got := l.Dist(3, 2); got != 3 {
		t.Errorf("Dist(3,2) = %d, want 3", got)
	}
	if got := l.Dist(0, 0); got != 0 {
		t.Errorf("Dist(0,0) = %d, want 0", got)
	}
}

func TestLCASameNode(t *testing.T) {
	l := NewLCA(3)
	l.AddEdge(0, 1)
	l.AddEdge(1, 2)
	l.Build(0)

	if got := l.LCA(1, 1); got != 1 {
		t.Errorf("LCA(1,1) = %d, want 1", got)
	}
}
