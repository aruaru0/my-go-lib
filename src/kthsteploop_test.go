package mylib

import "testing"

func TestKthStepLoop(t *testing.T) {
	N := 5
	next := []int{1, 2, 3, 4, 0}
	got := KthStepLoop(N, 5, next)
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestKthStepLoopSmall(t *testing.T) {
	N := 3
	next := []int{1, 2, 0}
	got := KthStepLoop(N, 2, next)
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}

func TestKthStepLoopNoLoop(t *testing.T) {
	N := 3
	next := []int{1, 2, 2}
	got := KthStepLoop(N, 10, next)
	if got != 2 {
		t.Errorf("expected 2, got %d", got)
	}
}
