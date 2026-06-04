package mylib

import "testing"

func TestKthStep(t *testing.T) {
	N := 5
	next := []int{1, 2, 3, 4, 0}
	got := KthStep(N, 5, next)
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestKthStepOneStep(t *testing.T) {
	N := 3
	next := []int{1, 2, 0}
	got := KthStep(N, 1, next)
	if got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
}

func TestKthStepBigK(t *testing.T) {
	N := 2
	next := []int{1, 0}
	got := KthStep(N, 1e18, next)
	if got != 0 {
		t.Errorf("expected 0 (even steps), got %d", got)
	}
}

func TestKthStepSelfLoop(t *testing.T) {
	N := 3
	next := []int{0, 0, 0}
	got := KthStep(N, 100, next)
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}
