package mylib

import "testing"

func TestMultiSetPutAndCount(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Put(1)
	ms.Put(2)
	ms.Put(1)
	ms.Put(3)

	if got := ms.Count(1); got != 2 {
		t.Errorf("Count(1) = %d, want 2", got)
	}
	if got := ms.Count(2); got != 1 {
		t.Errorf("Count(2) = %d, want 1", got)
	}
	if got := ms.Count(3); got != 1 {
		t.Errorf("Count(3) = %d, want 1", got)
	}
	if got := ms.Count(4); got != 0 {
		t.Errorf("Count(4) = %d, want 0", got)
	}
}

func TestMultiSetRemove(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Put(1)
	ms.Put(1)
	ms.Put(1)
	ms.Remove(1)
	if got := ms.Count(1); got != 2 {
		t.Errorf("Count(1) after one Remove = %d, want 2", got)
	}
	ms.Remove(1)
	ms.Remove(1)
	if got := ms.Count(1); got != 0 {
		t.Errorf("Count(1) after three Removes = %d, want 0", got)
	}
}

func TestMultiSetRemoveNonExistent(t *testing.T) {
	ms := NewMultiSet[string]()
	ms.Put("a")
	ms.Remove("b")
	if got := ms.Count("a"); got != 1 {
		t.Errorf("Count(a) after removing non-existent = %d, want 1", got)
	}
}

func TestMultiSetLen(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Put(1)
	ms.Put(2)
	ms.Put(3)
	if got := ms.Len(); got != 3 {
		t.Errorf("Len() = %d, want 3", got)
	}
	ms.Remove(1)
	if got := ms.Len(); got != 2 {
		t.Errorf("Len() after remove = %d, want 2", got)
	}
}

func TestMultiSetValues(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Put(1)
	ms.Put(2)
	ms.Put(3)
	vals := ms.Values()
	seen := make(map[int]bool)
	for _, v := range vals {
		seen[v] = true
	}
	for _, v := range []int{1, 2, 3} {
		if !seen[v] {
			t.Errorf("Values() missing %d", v)
		}
	}
}
