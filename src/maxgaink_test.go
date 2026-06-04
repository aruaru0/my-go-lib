package mylib

import "testing"

func TestMaxGainK(t *testing.T) {
	p := []int{2, 3, 1}
	c := []int{5, -3, 4}
	got := MaxGainK(0, 3, 5, p, c)
	if got != 7 {
		t.Errorf("expected 7, got %d", got)
	}
}

func TestMaxGainKSmallK(t *testing.T) {
	p := []int{2, 3, 1}
	c := []int{5, -3, 4}
	got := MaxGainK(0, 3, 2, p, c)
	if got != 1 {
		t.Errorf("expected 1, got %d", got)
	}
}

func TestMaxGainKNegative(t *testing.T) {
	p := []int{2, 3, 1}
	c := []int{-1, -2, -3}
	got := MaxGainK(0, 3, 10, p, c)
	if got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
}

func TestMaxGainKOneMove(t *testing.T) {
	p := []int{1}
	c := []int{5}
	got := MaxGainK(0, 1, 3, p, c)
	if got != 15 {
		t.Errorf("expected 15, got %d", got)
	}
}
