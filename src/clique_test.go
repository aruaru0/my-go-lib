package mylib

import "testing"

func TestMinCliqueCoverTriangle(t *testing.T) {
	N := 3
	edges := [][]int{
		{1, 2},
		{0, 2},
		{0, 1},
	}
	got := MinCliqueCover(N, edges)
	if got != 1 {
		t.Errorf("expected 1 (complete graph), got %d", got)
	}
}

func TestMinCliqueCoverNoEdges(t *testing.T) {
	N := 3
	edges := [][]int{{}, {}, {}}
	got := MinCliqueCover(N, edges)
	if got != 3 {
		t.Errorf("expected 3 (each node alone), got %d", got)
	}
}

func TestMinCliqueCoverLine(t *testing.T) {
	N := 3
	edges := [][]int{
		{1},
		{0, 2},
		{1},
	}
	got := MinCliqueCover(N, edges)
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}

func TestMinCliqueCoverTwoComponents(t *testing.T) {
	N := 4
	edges := [][]int{
		{1},
		{0},
		{3},
		{2},
	}
	got := MinCliqueCover(N, edges)
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}
