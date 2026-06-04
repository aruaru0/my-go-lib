package mylib

import "testing"

func TestSubsets(t *testing.T) {
	got := Subsets(0b1110)
	expected := []int{0b1110, 0b1100, 0b1010, 0b1000, 0b0110, 0b0100, 0b0010}
	if len(got) != len(expected) {
		t.Fatalf("expected %d subsets, got %d: %v", len(expected), len(got), got)
	}
	m := make(map[int]bool)
	for _, v := range got {
		m[v] = true
	}
	for _, e := range expected {
		if !m[e] {
			t.Errorf("missing subset %b", e)
		}
	}
}

func TestSubsetsZero(t *testing.T) {
	got := Subsets(0)
	if len(got) != 0 {
		t.Errorf("expected empty, got %v", got)
	}
}

func TestSubsetsSingleBit(t *testing.T) {
	got := Subsets(1 << 3)
	if len(got) != 1 || got[0] != 1<<3 {
		t.Errorf("expected [%d], got %v", 1<<3, got)
	}
}
