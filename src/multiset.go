package mylib

type MultiSet[T comparable] struct {
	m map[T]int
}

func NewMultiSet[T comparable]() *MultiSet[T] {
	return &MultiSet[T]{m: make(map[T]int)}
}

func (ms *MultiSet[T]) Put(x T) {
	ms.m[x]++
}

func (ms *MultiSet[T]) Remove(x T) {
	if cnt, ok := ms.m[x]; ok {
		if cnt <= 1 {
			delete(ms.m, x)
		} else {
			ms.m[x] = cnt - 1
		}
	}
}

func (ms *MultiSet[T]) Count(x T) int {
	return ms.m[x]
}

func (ms *MultiSet[T]) Values() []T {
	res := make([]T, 0, len(ms.m))
	for x := range ms.m {
		res = append(res, x)
	}
	return res
}

func (ms *MultiSet[T]) Len() int {
	return len(ms.m)
}
