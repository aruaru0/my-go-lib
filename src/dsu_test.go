package mylib

import (
	"testing"
)

func TestDsuNew(t *testing.T) {
	d := NewDsu(10)
	if d == nil {
		t.Fatal("NewDsu returned nil")
	}
}

func TestDsuInitiallyDisjoint(t *testing.T) {
	d := NewDsu(5)
	for i := 0; i < 5; i++ {
		if d.Same(i, (i+1)%5) {
			t.Errorf("expected %d and %d to be initially disjoint", i, (i+1)%5)
		}
	}
}

func TestDsuMerge(t *testing.T) {
	d := NewDsu(10)
	d.Merge(0, 1)
	if !d.Same(0, 1) {
		t.Error("0 and 1 should be in the same set after Merge")
	}
	if d.Same(0, 2) {
		t.Error("0 and 2 should still be in different sets")
	}
}

func TestDsuTransitivity(t *testing.T) {
	d := NewDsu(10)
	d.Merge(0, 1)
	d.Merge(1, 2)
	d.Merge(3, 4)
	if !d.Same(0, 2) {
		t.Error("0 and 2 should be in the same set via transitivity")
	}
	if d.Same(0, 3) {
		t.Error("0 and 3 should be in different sets")
	}
}

func TestDsuSize(t *testing.T) {
	d := NewDsu(10)
	if s := d.Size(0); s != 1 {
		t.Errorf("initial size of a set should be 1, got %d", s)
	}
	d.Merge(0, 1)
	if s := d.Size(0); s != 2 {
		t.Errorf("size after merging 0 and 1 should be 2, got %d", s)
	}
	d.Merge(1, 2)
	if s := d.Size(0); s != 3 {
		t.Errorf("size after merging 0,1,2 should be 3, got %d", s)
	}
}

func TestDsuLeader(t *testing.T) {
	d := NewDsu(10)
	d.Merge(0, 1)
	d.Merge(2, 3)
	d.Merge(0, 2)
	leader := d.Leader(1)
	if !d.Same(1, leader) {
		t.Error("Leader should be in the same set")
	}
	if d.Leader(4) != 4 {
		t.Error("Singleton should be its own leader")
	}
}

func TestDsuGroups(t *testing.T) {
	d := NewDsu(6)
	d.Merge(0, 1)
	d.Merge(2, 3)
	d.Merge(4, 5)
	d.Merge(0, 2)

	groups := d.Groups()
	sum := 0
	for _, g := range groups {
		sum += len(g)
	}
	if sum != 6 {
		t.Errorf("sum of group sizes should be 6, got %d", sum)
	}

	found := false
	for _, g := range groups {
		if len(g) == 4 {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected a group of size 4 (0,1,2,3)")
	}
}

func TestDsuAllElements(t *testing.T) {
	n := 100
	d := NewDsu(n)
	for i := 0; i < n-1; i++ {
		d.Merge(i, i+1)
	}
	for i := 0; i < n; i++ {
		if !d.Same(0, i) {
			t.Errorf("all elements should be in the same set, but %d is not", i)
		}
	}
	if s := d.Size(0); s != n {
		t.Errorf("size should be %d, got %d", n, s)
	}
}
