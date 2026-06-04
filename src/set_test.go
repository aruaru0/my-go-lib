package mylib

import "testing"

func TestSetAddAndContains(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if !s.Contains(1) {
		t.Error("Contains(1) = false, want true")
	}
	if !s.Contains(2) {
		t.Error("Contains(2) = false, want true")
	}
	if s.Contains(4) {
		t.Error("Contains(4) = true, want false")
	}
}

func TestSetRemove(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Remove(1)
	if s.Contains(1) {
		t.Error("Contains(1) after Remove = true, want false")
	}
	if !s.Contains(2) {
		t.Error("Contains(2) after Remove = false, want true")
	}
}

func TestSetCardinality(t *testing.T) {
	s := NewSet[int]()
	if got := s.Cardinality(); got != 0 {
		t.Errorf("Cardinality() = %d, want 0", got)
	}
	s.Add(1)
	s.Add(2)
	if got := s.Cardinality(); got != 2 {
		t.Errorf("Cardinality() = %d, want 2", got)
	}
	s.Add(1)
	if got := s.Cardinality(); got != 2 {
		t.Errorf("Cardinality() after duplicate add = %d, want 2", got)
	}
}

func TestSetClear(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Clear()
	if got := s.Cardinality(); got != 0 {
		t.Errorf("Cardinality() after Clear = %d, want 0", got)
	}
}

func TestSetUnion(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)

	u := a.Union(b)
	if !u.Contains(1) || !u.Contains(2) || !u.Contains(3) {
		t.Error("Union missing elements")
	}
	if u.Cardinality() != 3 {
		t.Errorf("Union cardinality = %d, want 3", u.Cardinality())
	}
}

func TestSetIntersect(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	b := NewSet[int]()
	b.Add(2)
	b.Add(3)
	b.Add(4)

	inter := a.Intersect(b)
	if !inter.Contains(2) || !inter.Contains(3) {
		t.Error("Intersect missing elements")
	}
	if inter.Contains(1) || inter.Contains(4) {
		t.Error("Intersect has unexpected elements")
	}
}

func TestSetDifference(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	b := NewSet[int]()
	b.Add(2)

	diff := a.Difference(b)
	if !diff.Contains(1) || !diff.Contains(3) {
		t.Error("Difference missing elements")
	}
	if diff.Cardinality() != 2 {
		t.Errorf("Difference cardinality = %d, want 2", diff.Cardinality())
	}
}

func TestSetEqual(t *testing.T) {
	a := NewSet[int]()
	a.Add(1)
	a.Add(2)
	b := NewSet[int]()
	b.Add(2)
	b.Add(1)

	if !a.Equal(b) {
		t.Error("Equal sets not detected as equal")
	}
	b.Add(3)
	if a.Equal(b) {
		t.Error("Different sets detected as equal")
	}
}

func TestSetNewFromSlice(t *testing.T) {
	s := NewSetFromSlice([]int{1, 2, 3, 2, 1})
	if s.Cardinality() != 3 {
		t.Errorf("Cardinality() = %d, want 3", s.Cardinality())
	}
}

func TestSetAddReturn(t *testing.T) {
	s := NewSet[int]()
	if added := s.Add(1); !added {
		t.Error("Add(1) returned false, want true")
	}
	if added := s.Add(1); added {
		t.Error("Add(1) second time returned true, want false")
	}
}
