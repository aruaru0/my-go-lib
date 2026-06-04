package mylib

type PriorityQueue[T any] struct {
	items []T
	less  func(a, b T) bool
}

func NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{less: less}
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.items)
}

func (pq *PriorityQueue[T]) Push(x T) {
	pq.items = append(pq.items, x)
	i := len(pq.items) - 1
	for i > 0 {
		p := (i - 1) / 2
		if !pq.less(pq.items[i], pq.items[p]) {
			break
		}
		pq.items[i], pq.items[p] = pq.items[p], pq.items[i]
		i = p
	}
}

func (pq *PriorityQueue[T]) Peek() T {
	return pq.items[0]
}

func (pq *PriorityQueue[T]) Pop() T {
	n := len(pq.items) - 1
	pq.items[0], pq.items[n] = pq.items[n], pq.items[0]
	pq.heapify(0, n)
	x := pq.items[n]
	pq.items = pq.items[:n]
	return x
}

func (pq *PriorityQueue[T]) heapify(i, n int) {
	for {
		smallest := i
		l := 2*i + 1
		r := 2*i + 2
		if l < n && pq.less(pq.items[l], pq.items[smallest]) {
			smallest = l
		}
		if r < n && pq.less(pq.items[r], pq.items[smallest]) {
			smallest = r
		}
		if smallest == i {
			break
		}
		pq.items[i], pq.items[smallest] = pq.items[smallest], pq.items[i]
		i = smallest
	}
}
