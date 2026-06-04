package mylib

import (
	"math"
	"testing"
)

func TestCircleCrossIntersecting(t *testing.T) {
	pts := CircleCross(PointF{0, 0}, PointF{3, 0}, 2, 2)
	if len(pts) != 2 {
		t.Fatalf("expected 2 points, got %d", len(pts))
	}
	if math.Abs(pts[0].X-1.5) > 1e-9 || math.Abs(pts[0].Y+math.Sqrt(4-2.25)) > 1e-9 {
		t.Errorf("unexpected point: %v", pts[0])
	}
	if math.Abs(pts[1].X-1.5) > 1e-9 || math.Abs(pts[1].Y-math.Sqrt(4-2.25)) > 1e-9 {
		t.Errorf("unexpected point: %v", pts[1])
	}
}

func TestCircleCrossSeparate(t *testing.T) {
	pts := CircleCross(PointF{0, 0}, PointF{10, 0}, 2, 2)
	if pts != nil {
		t.Errorf("expected no intersection, got %v", pts)
	}
}

func TestCircleCrossConcentric(t *testing.T) {
	pts := CircleCross(PointF{0, 0}, PointF{0, 0}, 2, 2)
	if pts != nil {
		t.Errorf("expected no intersection for concentric circles, got %v", pts)
	}
}

func TestCircleCrossTangent(t *testing.T) {
	pts := CircleCross(PointF{0, 0}, PointF{4, 0}, 2, 2)
	if len(pts) != 2 {
		t.Errorf("expected 2 points (same), got %d", len(pts))
	}
}
