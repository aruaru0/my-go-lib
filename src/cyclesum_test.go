package mylib

import "testing"

func TestCycleSumNoCycle(t *testing.T) {
	next := func(x int) int { return x + 1 }
	got := CycleSum(0, 5, 100, next)
	want := 0 + 1 + 2 + 3 + 4
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestCycleSumWithCycle(t *testing.T) {
	next := func(x int) int { return (x + 2) % 5 }
	got := CycleSum(0, 10, 5, next)
	want := 0 + 2 + 4 + 1 + 3 + 0 + 2 + 4 + 1 + 3
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestCycleSumSingle(t *testing.T) {
	next := func(x int) int { return x }
	got := CycleSum(5, 1, 100, next)
	if got != 5 {
		t.Fatalf("expected 5, got %d", got)
	}
}

func TestCycleSumSmallN(t *testing.T) {
	next := func(x int) int { return (x + 1) % 3 }
	got := CycleSum(0, 2, 3, next)
	want := 0 + 1
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}

func TestCycleSumMod7(t *testing.T) {
	next := func(x int) int { return (x * x) % 7 }
	got := CycleSum(2, 10, 7, next)
	want := 0
	for i, x := 0, 2; i < 10; i, x = i+1, next(x) {
		want += x
	}
	if got != want {
		t.Fatalf("expected %d, got %d", want, got)
	}
}
