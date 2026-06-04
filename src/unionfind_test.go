package mylib

import (
	"testing"
)

func TestUnionFindNew(t *testing.T) {
	uf := NewUnionFind(10)
	if uf == nil {
		t.Fatal("NewUnionFind returned nil")
	}
}

func TestUnionFindInitiallyDisjoint(t *testing.T) {
	uf := NewUnionFind(5)
	for i := 0; i < 5; i++ {
		if uf.Same(i, (i+1)%5) {
			t.Errorf("expected %d and %d to be initially disjoint", i, (i+1)%5)
		}
	}
}

func TestUnionFindUnite(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Unite(0, 1)
	if !uf.Same(0, 1) {
		t.Error("0 and 1 should be in the same set after Unite")
	}
	if uf.Same(0, 2) {
		t.Error("0 and 2 should still be in different sets")
	}
}

func TestUnionFindTransitivity(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Unite(0, 1)
	uf.Unite(1, 2)
	uf.Unite(3, 4)
	if !uf.Same(0, 2) {
		t.Error("0 and 2 should be in the same set via transitivity")
	}
	if uf.Same(0, 3) {
		t.Error("0 and 3 should be in different sets")
	}
}

func TestUnionFindSize(t *testing.T) {
	uf := NewUnionFind(10)
	if s := uf.Size(0); s != 1 {
		t.Errorf("initial size of a set should be 1, got %d", s)
	}
	uf.Unite(0, 1)
	if s := uf.Size(0); s != 2 {
		t.Errorf("size after uniting 0 and 1 should be 2, got %d", s)
	}
	uf.Unite(1, 2)
	if s := uf.Size(0); s != 3 {
		t.Errorf("size after uniting 0,1,2 should be 3, got %d", s)
	}
}

func TestUnionFindMultipleSets(t *testing.T) {
	uf := NewUnionFind(10)
	uf.Unite(0, 1)
	uf.Unite(2, 3)
	uf.Unite(0, 2)
	if !uf.Same(1, 3) {
		t.Error("1 and 3 should be in the same set after merging")
	}
}

func TestUnionFindAllElements(t *testing.T) {
	n := 100
	uf := NewUnionFind(n)
	for i := 0; i < n-1; i++ {
		uf.Unite(i, i+1)
	}
	for i := 0; i < n; i++ {
		if !uf.Same(0, i) {
			t.Errorf("all elements should be in the same set, but %d is not", i)
		}
	}
	if s := uf.Size(0); s != n {
		t.Errorf("size should be %d, got %d", n, s)
	}
}
