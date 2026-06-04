package mylib

import (
	"testing"
)

func TestDinicSimple(t *testing.T) {
	d := NewDinic(4)
	d.AddEdge(0, 1, 10)
	d.AddEdge(0, 2, 5)
	d.AddEdge(1, 2, 15)
	d.AddEdge(1, 3, 10)
	d.AddEdge(2, 3, 10)
	flow := d.MaxFlow(0, 3)
	if flow != 15 {
		t.Errorf("expected max flow 15, got %d", flow)
	}
}

func TestDinicNoPath(t *testing.T) {
	d := NewDinic(3)
	d.AddEdge(0, 1, 5)
	flow := d.MaxFlow(0, 2)
	if flow != 0 {
		t.Errorf("expected 0 flow, got %d", flow)
	}
}

func TestDinicDisconnected(t *testing.T) {
	d := NewDinic(3)
	flow := d.MaxFlow(0, 2)
	if flow != 0 {
		t.Errorf("expected 0 flow, got %d", flow)
	}
}

func TestDinicLargeFlow(t *testing.T) {
	d := NewDinic(2)
	d.AddEdge(0, 1, 1000000)
	flow := d.MaxFlow(0, 1)
	if flow != 1000000 {
		t.Errorf("expected 1000000, got %d", flow)
	}
}

func TestDinicMultiplePaths(t *testing.T) {
	d := NewDinic(3)
	d.AddEdge(0, 1, 5)
	d.AddEdge(0, 2, 5)
	d.AddEdge(1, 2, 5)
	flow := d.MaxFlow(0, 2)
	if flow != 10 {
		t.Errorf("expected max flow 10, got %d", flow)
	}
}
