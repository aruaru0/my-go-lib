package mylib

import "testing"

func TestMaxGainSimple(t *testing.T) {
	N := 3
	p := []int{2, 3, 1}
	c := []int{1, 2, 3}
	got := MaxGain(N, 2, p, c)
	if got != 5 {
		t.Errorf("expected 5, got %d", got)
	}
}

func TestMaxGainSmallK(t *testing.T) {
	N := 3
	p := []int{2, 3, 1}
	c := []int{5, 10, -2}
	got := MaxGain(N, 1, p, c)
	if got != 10 {
		t.Errorf("expected 10, got %d", got)
	}
}

func TestMaxGainNegativeLoop(t *testing.T) {
	N := 2
	p := []int{2, 1}
	c := []int{-5, -10}
	got := MaxGain(N, 10, p, c)
	if got != 0 {
		t.Errorf("expected 0 (stay at start), got %d", got)
	}
}
