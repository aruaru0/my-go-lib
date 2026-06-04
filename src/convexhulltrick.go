package mylib

import "sort"

type ConvexHullTrick[T Number] struct {
	pos  []T
	root *chtNode[T]
}

type chtLine[T Number] struct {
	a, b T
}

type chtNode[T Number] struct {
	l        chtLine[T]
	lch, rch *chtNode[T]
}

func NewConvexHullTrick[T Number](pos []T) *ConvexHullTrick[T] {
	tmp := append([]T{}, pos...)
	sort.Slice(tmp, func(i, j int) bool { return tmp[i] < tmp[j] })
	uniq := make([]T, 0, len(tmp))
	for i, p := range tmp {
		if i == 0 || p != tmp[i-1] {
			uniq = append(uniq, p)
		}
	}
	return &ConvexHullTrick[T]{pos: uniq}
}

func (c *ConvexHullTrick[T]) AddLine(a, b T) {
	c.root = c.modify(c.root, 0, len(c.pos)-1, chtLine[T]{a, b})
}

func (c *ConvexHullTrick[T]) GetMax(x T) T {
	t := sort.Search(len(c.pos), func(i int) bool { return c.pos[i] >= x })
	if t >= len(c.pos) {
		t = len(c.pos) - 1
	}
	return c.sub(c.root, 0, len(c.pos)-1, t)
}

func (l chtLine[T]) get(x T) T { return l.a*x + l.b }

func (c *ConvexHullTrick[T]) modify(p *chtNode[T], lb, ub int, l chtLine[T]) *chtNode[T] {
	if p == nil {
		return &chtNode[T]{l: l}
	}
	pos := c.pos
	curLb := p.l.get(pos[lb])
	curUb := p.l.get(pos[ub])
	newLb := l.get(pos[lb])
	newUb := l.get(pos[ub])
	if curLb >= newLb && curUb >= newUb {
		return p
	}
	if curLb <= newLb && curUb <= newUb {
		p.l = l
		return p
	}
	mid := (lb + ub) / 2
	if p.l.get(pos[mid]) < l.get(pos[mid]) {
		p.l, l = l, p.l
	}
	if p.l.get(pos[lb]) <= l.get(pos[lb]) {
		p.lch = c.modify(p.lch, lb, mid, l)
	} else {
		p.rch = c.modify(p.rch, mid+1, ub, l)
	}
	return p
}

func (c *ConvexHullTrick[T]) sub(p *chtNode[T], lb, ub, t int) T {
	if p == nil {
		var zero T
		return zero
	}
	pos := c.pos
	if ub-lb == 0 {
		return p.l.get(pos[t])
	}
	mid := (lb + ub) / 2
	if t <= mid {
		return max(p.l.get(pos[t]), c.sub(p.lch, lb, mid, t))
	}
	return max(p.l.get(pos[t]), c.sub(p.rch, mid+1, ub, t))
}
