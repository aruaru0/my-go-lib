package mylib

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func NewSetFromSlice[T comparable](s []T) Set[T] {
	set := make(Set[T])
	for _, v := range s {
		set.Add(v)
	}
	return set
}

func (s Set[T]) Add(v T) bool {
	_, found := s[v]
	s[v] = struct{}{}
	return !found
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Cardinality() int {
	return len(s)
}

func (s Set[T]) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set[T]) Values() []T {
	res := make([]T, 0, len(s))
	for k := range s {
		res = append(res, k)
	}
	return res
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	res := NewSet[T]()
	for k := range s {
		res.Add(k)
	}
	for k := range other {
		res.Add(k)
	}
	return res
}

func (s Set[T]) Intersect(other Set[T]) Set[T] {
	res := NewSet[T]()
	small, large := s, other
	if len(small) > len(large) {
		small, large = large, small
	}
	for k := range small {
		if large.Contains(k) {
			res.Add(k)
		}
	}
	return res
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	res := NewSet[T]()
	for k := range s {
		if !other.Contains(k) {
			res.Add(k)
		}
	}
	return res
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for k := range s {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}
