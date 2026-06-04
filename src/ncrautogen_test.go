package mylib

import (
	"testing"
)

func TestNCRGeneratorNCR(t *testing.T) {
	g := NewNCRGenerator(1000000007)
	if got := g.NCR(5, 2); got != 10 {
		t.Errorf("NCR(5,2) = %d, want 10", got)
	}
	if got := g.NCR(10, 3); got != 120 {
		t.Errorf("NCR(10,3) = %d, want 120", got)
	}
}

func TestNCRGeneratorEdge(t *testing.T) {
	g := NewNCRGenerator(1000000007)
	if got := g.NCR(5, 5); got != 1 {
		t.Errorf("NCR(5,5) = %d, want 1", got)
	}
	if got := g.NCR(5, 0); got != 1 {
		t.Errorf("NCR(5,0) = %d, want 1", got)
	}
	if got := g.NCR(5, 6); got != 0 {
		t.Errorf("NCR(5,6) = %d, want 0", got)
	}
	if got := g.NCR(5, -1); got != 0 {
		t.Errorf("NCR(5,-1) = %d, want 0", got)
	}
}

func TestNCRGeneratorAutoGrow(t *testing.T) {
	g := NewNCRGenerator(1000000007)
	_ = g.NCR(100, 50)
	if got := g.NCR(200, 100); got == 0 {
		t.Errorf("NCR(200,100) should not be 0")
	}
}

func TestNCRGeneratorLarge(t *testing.T) {
	g := NewNCRGenerator(1000000007)
	got := g.NCR(1000, 500)
	if got == 0 {
		t.Errorf("NCR(1000,500) should not be 0")
	}
}
