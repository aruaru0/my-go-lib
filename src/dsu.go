package mylib

type DSU struct {
	parentOrSize []int
	n            int
}

func NewDsu(n int) *DSU {
	d := &DSU{
		n:            n,
		parentOrSize: make([]int, n),
	}
	for i := 0; i < n; i++ {
		d.parentOrSize[i] = -1
	}
	return d
}

func (d *DSU) Merge(a, b int) int {
	x, y := d.Leader(a), d.Leader(b)
	if x == y {
		return x
	}
	if -d.parentOrSize[x] < -d.parentOrSize[y] {
		x, y = y, x
	}
	d.parentOrSize[x] += d.parentOrSize[y]
	d.parentOrSize[y] = x
	return x
}

func (d *DSU) Same(a, b int) bool {
	return d.Leader(a) == d.Leader(b)
}

func (d *DSU) Leader(a int) int {
	if d.parentOrSize[a] < 0 {
		return a
	}
	d.parentOrSize[a] = d.Leader(d.parentOrSize[a])
	return d.parentOrSize[a]
}

func (d *DSU) Size(a int) int {
	return -d.parentOrSize[d.Leader(a)]
}

func (d *DSU) Groups() [][]int {
	m := make(map[int][]int)
	for i := 0; i < d.n; i++ {
		x := d.Leader(i)
		if x < 0 {
			m[i] = append(m[i], i)
		} else {
			m[x] = append(m[x], i)
		}
	}
	ret := make([][]int, len(m))
	idx := 0
	for _, e := range m {
		ret[idx] = make([]int, len(e))
		copy(ret[idx], e)
		idx++
	}
	return ret
}
