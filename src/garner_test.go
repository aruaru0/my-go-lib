package mylib

import (
	"testing"
)

func TestGarner(t *testing.T) {
	x := []int{2, 3, 2}
	m := []int{3, 5, 7}
	got := Garner(x, m, 0)
	expected := 23
	if got != expected {
		t.Errorf("Garner([2,3,2],[3,5,7],0) = %d, want %d", got, expected)
	}
}

func TestGarnerWithMod(t *testing.T) {
	x := []int{2, 3}
	m := []int{3, 5}
	got := Garner(x, m, 100)
	expected := 8 % 100
	if got != expected {
		t.Errorf("Garner([2,3],[3,5],100) = %d, want %d", got, expected)
	}
}

func TestGarnerEmpty(t *testing.T) {
	if got := Garner([]int{}, []int{}, 0); got != 0 {
		t.Errorf("Garner empty = %d, want 0", got)
	}
}
