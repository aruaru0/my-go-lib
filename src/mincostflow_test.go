package mylib

import (
	"testing"
)

func TestMinCostFlowSimple(t *testing.T) {
	mcf := NewMinCostFlow(4)
	mcf.AddEdge(0, 1, 10, 2)
	mcf.AddEdge(0, 2, 10, 5)
	mcf.AddEdge(1, 2, 10, 1)
	mcf.AddEdge(1, 3, 10, 3)
	mcf.AddEdge(2, 3, 10, 1)
	res := mcf.Flow(0, 3)
	if res[0] != 20 {
		t.Errorf("expected flow 20, got %d", res[0])
	}
}

func TestMinCostFlowNoPath(t *testing.T) {
	mcf := NewMinCostFlow(3)
	mcf.AddEdge(0, 1, 5, 1)
	res := mcf.Flow(0, 2)
	if res[0] != 0 {
		t.Errorf("expected 0 flow, got %d", res[0])
	}
}

func TestMinCostFlowDisconnected(t *testing.T) {
	mcf := NewMinCostFlow(2)
	res := mcf.Flow(0, 1)
	if res[0] != 0 {
		t.Errorf("expected 0 flow for no edges, got %d", res[0])
	}
}
