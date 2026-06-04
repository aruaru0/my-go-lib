package mylib

import (
	"testing"
)

func TestConvexHull(t *testing.T) {
	pts := []Point{
		{0, 0}, {1, 1}, {2, 2}, {0, 2}, {2, 0},
	}
	hull := ConvexHull(pts)
	if len(hull) != 4 {
		t.Errorf("expected 4 points, got %d", len(hull))
	}
}

func TestConvexHullTriangle(t *testing.T) {
	pts := []Point{
		{0, 0}, {2, 0}, {1, 2},
	}
	hull := ConvexHull(pts)
	if len(hull) != 3 {
		t.Errorf("expected 3 points, got %d", len(hull))
	}
}

func TestConvexHullCollinear(t *testing.T) {
	pts := []Point{
		{0, 0}, {1, 1}, {2, 2},
	}
	hull := ConvexHull(pts)
	if len(hull) != 2 {
		t.Errorf("expected 2 points, got %d", len(hull))
	}
}

func TestConvexHullSingle(t *testing.T) {
	pts := []Point{{5, 5}}
	hull := ConvexHull(pts)
	if len(hull) != 1 {
		t.Errorf("expected 1 point, got %d", len(hull))
	}
}
