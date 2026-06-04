package mylib

import (
	"testing"
)

func TestPriorityQueueIntMinHeap(t *testing.T) {
	pq := NewPriorityQueue[int](func(a, b int) bool { return a < b })
	pq.Push(3)
	pq.Push(1)
	pq.Push(2)

	if pq.Len() != 3 {
		t.Fatalf("expected len 3, got %d", pq.Len())
	}

	if v := pq.Pop(); v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}
	if v := pq.Pop(); v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}
	if v := pq.Pop(); v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}
	if pq.Len() != 0 {
		t.Fatalf("expected len 0, got %d", pq.Len())
	}
}

func TestPriorityQueueIntMaxHeap(t *testing.T) {
	pq := NewPriorityQueue[int](func(a, b int) bool { return a > b })
	pq.Push(1)
	pq.Push(3)
	pq.Push(2)

	if v := pq.Pop(); v != 3 {
		t.Fatalf("expected 3, got %d", v)
	}
	if v := pq.Pop(); v != 2 {
		t.Fatalf("expected 2, got %d", v)
	}
	if v := pq.Pop(); v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}
}

func TestPriorityQueueString(t *testing.T) {
	pq := NewPriorityQueue[string](func(a, b string) bool { return a < b })
	pq.Push("banana")
	pq.Push("apple")
	pq.Push("cherry")

	if v := pq.Pop(); v != "apple" {
		t.Fatalf("expected apple, got %s", v)
	}
	if v := pq.Pop(); v != "banana" {
		t.Fatalf("expected banana, got %s", v)
	}
	if v := pq.Pop(); v != "cherry" {
		t.Fatalf("expected cherry, got %s", v)
	}
}

func TestPriorityQueueEmptyPop(t *testing.T) {
	pq := NewPriorityQueue[int](func(a, b int) bool { return a < b })
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on empty pop")
		}
	}()
	pq.Pop()
}

func TestPriorityQueueLarge(t *testing.T) {
	pq := NewPriorityQueue[int](func(a, b int) bool { return a < b })
	n := 10000
	for i := n; i > 0; i-- {
		pq.Push(i)
	}
	for i := 1; i <= n; i++ {
		if v := pq.Pop(); v != i {
			t.Fatalf("expected %d, got %d", i, v)
		}
	}
}
