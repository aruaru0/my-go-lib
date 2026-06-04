package mylib

import "testing"

func TestMedianBasic(t *testing.T) {
	m := NewMedian()
	m.Add(3)
	m.Add(1)
	m.Add(2)
	if v := m.Median(); v != 2 {
		t.Fatalf("expected median 2, got %d", v)
	}
}

func TestMedianEven(t *testing.T) {
	m := NewMedian()
	m.Add(1)
	m.Add(2)
	m.Add(3)
	m.Add(4)
	if v := m.Median(); v != 2 {
		t.Fatalf("expected median 2, got %d", v)
	}
}

func TestMedianAscending(t *testing.T) {
	m := NewMedian()
	for i := 1; i <= 5; i++ {
		m.Add(i)
	}
	if v := m.Median(); v != 3 {
		t.Fatalf("expected median 3, got %d", v)
	}
}

func TestMedianDescending(t *testing.T) {
	m := NewMedian()
	for i := 5; i >= 1; i-- {
		m.Add(i)
	}
	if v := m.Median(); v != 3 {
		t.Fatalf("expected median 3, got %d", v)
	}
}

func TestMedianDuplicates(t *testing.T) {
	m := NewMedian()
	m.Add(2)
	m.Add(2)
	m.Add(2)
	if v := m.Median(); v != 2 {
		t.Fatalf("expected median 2, got %d", v)
	}
}

func TestMedianSingle(t *testing.T) {
	m := NewMedian()
	m.Add(42)
	if v := m.Median(); v != 42 {
		t.Fatalf("expected median 42, got %d", v)
	}
}

func TestMedianEmpty(t *testing.T) {
	m := NewMedian()
	if v := m.Median(); v != 0 {
		t.Fatalf("expected median 0, got %d", v)
	}
}

func TestMedianLarge(t *testing.T) {
	m := NewMedian()
	n := 10000
	for i := 1; i <= n; i++ {
		m.Add(i)
	}
	if v := m.Median(); v != (n+1)/2 {
		t.Fatalf("expected median %d, got %d", (n+1)/2, v)
	}
}
