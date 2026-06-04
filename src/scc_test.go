package mylib

import (
	"testing"
)

func TestSccGraphSingleNode(t *testing.T) {
	g := NewSccGraph(1)
	scc := g.Scc()
	if len(scc) != 1 {
		t.Errorf("expected 1 component, got %d", len(scc))
	}
}

func TestSccGraphDisconnected(t *testing.T) {
	g := NewSccGraph(3)
	scc := g.Scc()
	if len(scc) != 3 {
		t.Errorf("expected 3 components for disconnected graph, got %d", len(scc))
	}
}

func TestSccGraphSimple(t *testing.T) {
	g := NewSccGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 0)
	g.AddEdge(1, 3)
	scc := g.Scc()
	if len(scc) > 3 {
		t.Errorf("expected at most 3 components, got %d", len(scc))
	}
	hasCycle := false
	for _, comp := range scc {
		if len(comp) == 3 {
			hasCycle = true
			m := make(map[int]bool)
			for _, v := range comp {
				m[v] = true
			}
			if !m[0] || !m[1] || !m[2] {
				t.Errorf("cycle component should contain 0,1,2, got %v", comp)
			}
		}
	}
	if !hasCycle {
		t.Error("expected a cycle component with 3 vertices")
	}
}

func TestSccGraphDag(t *testing.T) {
	g := NewSccGraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	scc := g.Scc()
	if len(scc) != 4 {
		t.Errorf("expected 4 components for DAG, got %d", len(scc))
	}
}
