package mylib

type Median struct {
	t *Treap[int]
}

func NewMedian() *Median {
	return &Median{t: NewTreap[int](func(a, b int) bool { return a < b })}
}

func (m *Median) Add(x int) {
	m.t.Insert(x)
}

func (m *Median) Remove(x int) {
	m.t.Delete(x)
}

func (m *Median) Median() int {
	n := m.t.Len()
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return m.t.Kth(n / 2)
	}
	return (m.t.Kth(n/2-1) + m.t.Kth(n/2)) / 2
}
