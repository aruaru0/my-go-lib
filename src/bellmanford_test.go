package mylib

import (
	"testing"
)

func TestBellmanFordSimple(t *testing.T) {
	edges := []Edge{
		{0, 1, 5},
		{0, 2, 3},
		{1, 3, 2},
		{2, 3, 4},
	}
	d, neg := BellmanFord(4, 0, 3, edges)
	if neg {
		t.Error("no negative cycle expected")
	}
	if d[3] != 7 {
		t.Errorf("expected distance 7, got %d", d[3])
	}
	if d[1] != 5 {
		t.Errorf("expected distance 5 to node 1, got %d", d[1])
	}
}

func TestBellmanFordUnreachable(t *testing.T) {
	edges := []Edge{
		{0, 1, 1},
	}
	d, neg := BellmanFord(3, 0, 2, edges)
	if neg {
		t.Error("no negative cycle expected")
	}
	if d[2] != 1<<60 {
		t.Error("node 2 should be unreachable")
	}
}

func TestBellmanFordNegativeCycle(t *testing.T) {
	edges := []Edge{
		{0, 1, 1},
		{1, 2, -3},
		{2, 0, 1},
	}
	_, neg := BellmanFord(3, 0, 2, edges)
	if !neg {
		t.Error("negative cycle should be detected")
	}
}

func TestBellmanFordNegativeEdgeNoCycle(t *testing.T) {
	edges := []Edge{
		{0, 1, 10},
		{1, 2, -5},
		{2, 3, 3},
	}
	d, neg := BellmanFord(4, 0, 3, edges)
	if neg {
		t.Error("no negative cycle expected")
	}
	if d[3] != 8 {
		t.Errorf("expected distance 8, got %d", d[3])
	}
}

func TestBellmanFordAllUnreachable(t *testing.T) {
	edges := []Edge{}
	_, neg := BellmanFord(3, 0, 1, edges)
	if neg {
		t.Error("no negative cycle expected with no edges")
	}
}
