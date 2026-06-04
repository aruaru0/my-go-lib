package mylib

import (
	"testing"
)

func TestMaxFlowSimple(t *testing.T) {
	mf := NewMaxFlow(4)
	mf.AddEdge(0, 1, 10)
	mf.AddEdge(0, 2, 5)
	mf.AddEdge(1, 2, 15)
	mf.AddEdge(1, 3, 10)
	mf.AddEdge(2, 3, 10)
	flow := mf.Flow(0, 3)
	if flow != 15 {
		t.Errorf("expected max flow 15, got %d", flow)
	}
}

func TestMaxFlowNoPath(t *testing.T) {
	mf := NewMaxFlow(3)
	mf.AddEdge(0, 1, 5)
	flow := mf.Flow(0, 2)
	if flow != 0 {
		t.Errorf("expected 0 flow for unreachable sink, got %d", flow)
	}
}

func TestMaxFlowSingleEdge(t *testing.T) {
	mf := NewMaxFlow(2)
	mf.AddEdge(0, 1, 42)
	flow := mf.Flow(0, 1)
	if flow != 42 {
		t.Errorf("expected flow 42, got %d", flow)
	}
}

func TestMaxFlowMinCut(t *testing.T) {
	mf := NewMaxFlow(4)
	mf.AddEdge(0, 1, 3)
	mf.AddEdge(0, 2, 2)
	mf.AddEdge(1, 2, 5)
	mf.AddEdge(1, 3, 2)
	mf.AddEdge(2, 3, 3)
	mf.Flow(0, 3)
	cut := mf.MinCut(0)
	if !cut[0] {
		t.Error("source should be in cut")
	}
}

func TestMaxFlowEdgesList(t *testing.T) {
	mf := NewMaxFlow(2)
	mf.AddEdge(0, 1, 10)
	mf.Flow(0, 1)
	edges := mf.EdgesList()
	if len(edges) != 1 {
		t.Errorf("expected 1 edge, got %d", len(edges))
	}
	if edges[0].Flow != 10 {
		t.Errorf("expected flow 10, got %d", edges[0].Flow)
	}
}
