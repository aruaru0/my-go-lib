package mylib

import (
	"testing"
)

func TestBellmanFordRouteSimple(t *testing.T) {
	edges := []Edge{
		{0, 1, 2},
		{1, 2, 3},
		{0, 2, 6},
	}
	d, r, neg := BellmanFordRoute(3, 0, edges)
	if neg {
		t.Error("no negative cycle expected")
	}
	if d[2] != 5 {
		t.Errorf("expected distance 5, got %d", d[2])
	}
	if len(r[2].Path) == 0 {
		t.Error("expected a route to node 2")
	}
}

func TestBellmanFordRouteNegativeCycle(t *testing.T) {
	edges := []Edge{
		{0, 1, 1},
		{1, 2, -4},
		{2, 0, 1},
	}
	_, _, neg := BellmanFordRoute(3, 0, edges)
	if !neg {
		t.Error("negative cycle should be detected")
	}
}

func TestBellmanFordRouteStartOnly(t *testing.T) {
	edges := []Edge{
		{0, 1, 5},
	}
	d, _, neg := BellmanFordRoute(2, 0, edges)
	if neg {
		t.Error("no negative cycle expected")
	}
	if d[0] != 0 {
		t.Errorf("distance to start should be 0, got %d", d[0])
	}
}
