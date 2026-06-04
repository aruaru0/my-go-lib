package mylib

import "testing"

func TestSegmentsIntersectCrossing(t *testing.T) {
	if !SegmentsIntersect(0, 0, 2, 2, 0, 2, 2, 0) {
		t.Error("expected crossing segments to intersect")
	}
}

func TestSegmentsIntersectNotCrossing(t *testing.T) {
	if SegmentsIntersect(0, 0, 1, 0, 2, 0, 3, 0) {
		t.Error("expected parallel segments not to intersect")
	}
}

func TestSegmentsIntersectNoTouch(t *testing.T) {
	if SegmentsIntersect(0, 0, 1, 1, 2, 2, 3, 3) {
		t.Error("expected non-touching segments not to intersect")
	}
}

func TestSegmentsIntersectEndToEnd(t *testing.T) {
	if SegmentsIntersect(0, 0, 1, 1, 1, 1, 2, 2) {
		t.Error("expected segments touching at endpoints only not to intersect (strict mode)")
	}
}

func TestSegmentsIntersectTFormation(t *testing.T) {
	if SegmentsIntersect(0, 0, 2, 0, 1, 0, 1, 2) {
		t.Error("expected T-formation segments not to intersect (strict mode)")
	}
}
