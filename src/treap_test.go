package mylib

import "testing"

func intLess(a, b int) bool { return a < b }

func TestTreapInsertAndFind(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(7)
	tr.Insert(1)

	if !tr.Find(5) {
		t.Error("Find(5) = false, want true")
	}
	if !tr.Find(3) {
		t.Error("Find(3) = false, want true")
	}
	if !tr.Find(7) {
		t.Error("Find(7) = false, want true")
	}
	if !tr.Find(1) {
		t.Error("Find(1) = false, want true")
	}
	if tr.Find(4) {
		t.Error("Find(4) = true, want false")
	}
}

func TestTreapDelete(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(7)
	tr.Delete(3)
	if tr.Find(3) {
		t.Error("Find(3) after Delete = true, want false")
	}
	if !tr.Find(5) {
		t.Error("Find(5) after Delete = false, want true")
	}
}

func TestTreapMinMax(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(7)
	tr.Insert(1)
	tr.Insert(9)

	if got := tr.Min(); got != 1 {
		t.Errorf("Min() = %d, want 1", got)
	}
	if got := tr.Max(); got != 9 {
		t.Errorf("Max() = %d, want 9", got)
	}
}

func TestTreapDuplicates(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(5)
	tr.Insert(5)
	tr.Delete(5)
	if !tr.Find(5) {
		t.Error("Find(5) after one Delete = false, want true (2 remaining)")
	}
	tr.Delete(5)
	tr.Delete(5)
	if tr.Find(5) {
		t.Error("Find(5) after three Deletes = true, want false")
	}
}

func TestTreapStringKeys(t *testing.T) {
	tr := NewTreap(func(a, b string) bool { return a < b })
	tr.Insert("c")
	tr.Insert("a")
	tr.Insert("b")
	if got := tr.Min(); got != "a" {
		t.Errorf("Min() = %q, want %q", got, "a")
	}
	if got := tr.Max(); got != "c" {
		t.Errorf("Max() = %q, want %q", got, "c")
	}
	if !tr.Find("b") {
		t.Error("Find(b) = false, want true")
	}
}
