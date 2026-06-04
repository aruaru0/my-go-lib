package mylib

import "testing"

func TestOrderedMapSetAndGet(t *testing.T) {
	m := NewOrderedMap[string, int]()
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)

	v, ok := m.Get("a")
	if !ok || v != 1 {
		t.Errorf(`Get("a") = (%d, %t), want (1, true)`, v, ok)
	}
	v, ok = m.Get("b")
	if !ok || v != 2 {
		t.Errorf(`Get("b") = (%d, %t), want (2, true)`, v, ok)
	}
	v, ok = m.Get("x")
	if ok {
		t.Errorf(`Get("x") = (%d, %t), want (0, false)`, v, ok)
	}
}

func TestOrderedMapOrder(t *testing.T) {
	m := NewOrderedMap[int, string]()
	m.Set(3, "c")
	m.Set(1, "a")
	m.Set(2, "b")

	want := []int{3, 1, 2}
	got := m.Keys()
	for i, k := range got {
		if k != want[i] {
			t.Errorf("Keys()[%d] = %d, want %d", i, k, want[i])
		}
	}
}

func TestOrderedMapDelete(t *testing.T) {
	m := NewOrderedMap[string, int]()
	m.Set("a", 1)
	m.Set("b", 2)
	m.Set("c", 3)
	m.Delete("b")

	if _, ok := m.Get("b"); ok {
		t.Error(`Get("b") after Delete returned true, want false`)
	}
	if got := m.Len(); got != 2 {
		t.Errorf("Len() after Delete = %d, want 2", got)
	}
	want := []string{"a", "c"}
	got := m.Keys()
	for i, k := range got {
		if k != want[i] {
			t.Errorf("Keys()[%d] = %q, want %q", i, k, want[i])
		}
	}
}

func TestOrderedMapOverwrite(t *testing.T) {
	m := NewOrderedMap[string, int]()
	m.Set("a", 1)
	m.Set("a", 99)
	v, ok := m.Get("a")
	if !ok || v != 99 {
		t.Errorf(`Get("a") after overwrite = (%d, %t), want (99, true)`, v, ok)
	}
	if got := m.Len(); got != 1 {
		t.Errorf("Len() after overwrite = %d, want 1", got)
	}
}
