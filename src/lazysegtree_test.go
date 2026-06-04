package mylib

import "testing"

type lazyS struct{ sum, len int }
type lazyF struct{ add int }

func TestLazySegTreeRangeAddRangeSum(t *testing.T) {
	v := make([]lazyS, 5)
	for i := range v {
		v[i] = lazyS{0, 1}
	}
	e := func() lazyS { return lazyS{0, 0} }
	merger := func(a, b lazyS) lazyS { return lazyS{a.sum + b.sum, a.len + b.len} }
	mapper := func(f lazyF, x lazyS) lazyS { return lazyS{x.sum + f.add*x.len, x.len} }
	comp := func(f, g lazyF) lazyF { return lazyF{f.add + g.add} }
	id := func() lazyF { return lazyF{0} }

	lseg := NewLazySegTree(v, e, merger, mapper, comp, id)
	lseg.RangeApply(0, 5, lazyF{3})
	if got := lseg.AllProd().sum; got != 15 {
		t.Errorf("AllProd() after range add = %d, want 15", got)
	}
	if got := lseg.Prod(0, 3).sum; got != 9 {
		t.Errorf("Prod(0,3) = %d, want 9", got)
	}
	lseg.RangeApply(1, 4, lazyF{2})
	if got := lseg.AllProd().sum; got != 21 {
		t.Errorf("AllProd() after second range add = %d, want 21", got)
	}
	if got := lseg.Get(0).sum; got != 3 {
		t.Errorf("Get(0) = %d, want 3", got)
	}
	if got := lseg.Get(2).sum; got != 5 {
		t.Errorf("Get(2) = %d, want 5", got)
	}
}

func TestLazySegTreePointSet(t *testing.T) {
	v := []lazyS{{1, 1}, {2, 1}, {3, 1}}
	e := func() lazyS { return lazyS{0, 0} }
	merger := func(a, b lazyS) lazyS { return lazyS{a.sum + b.sum, a.len + b.len} }
	mapper := func(f lazyF, x lazyS) lazyS { return lazyS{x.sum + f.add*x.len, x.len} }
	comp := func(f, g lazyF) lazyF { return lazyF{f.add + g.add} }
	id := func() lazyF { return lazyF{0} }

	lseg := NewLazySegTree(v, e, merger, mapper, comp, id)
	lseg.Set(1, lazyS{100, 1})
	if got := lseg.AllProd().sum; got != 104 {
		t.Errorf("AllProd() after Set = %d, want 104", got)
	}
}

func TestLazySegTreePointApply(t *testing.T) {
	v := []lazyS{{1, 1}, {2, 1}, {3, 1}}
	e := func() lazyS { return lazyS{0, 0} }
	merger := func(a, b lazyS) lazyS { return lazyS{a.sum + b.sum, a.len + b.len} }
	mapper := func(f lazyF, x lazyS) lazyS { return lazyS{x.sum + f.add*x.len, x.len} }
	comp := func(f, g lazyF) lazyF { return lazyF{f.add + g.add} }
	id := func() lazyF { return lazyF{0} }

	lseg := NewLazySegTree(v, e, merger, mapper, comp, id)
	lseg.Apply(1, lazyF{10})
	if got := lseg.Get(1).sum; got != 12 {
		t.Errorf("Get(1) after Apply = %d, want 12", got)
	}
	if got := lseg.AllProd().sum; got != 16 {
		t.Errorf("AllProd() = %d, want 16", got)
	}
}

func TestLazySegTreeMaxRight(t *testing.T) {
	v := []lazyS{{1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}}
	e := func() lazyS { return lazyS{0, 0} }
	merger := func(a, b lazyS) lazyS { return lazyS{a.sum + b.sum, a.len + b.len} }
	mapper := func(f lazyF, x lazyS) lazyS { return lazyS{x.sum + f.add*x.len, x.len} }
	comp := func(f, g lazyF) lazyF { return lazyF{f.add + g.add} }
	id := func() lazyF { return lazyF{0} }

	lseg := NewLazySegTree(v, e, merger, mapper, comp, id)
	// Prod(2,5)=3+4+5=12, not < 12; Prod(2,4)=7 < 12
	got := lseg.MaxRight(2, func(s lazyS) bool { return s.sum < 12 })
	want := 4
	if got != want {
		t.Errorf("MaxRight(2, s.sum < 12) = %d, want %d", got, want)
	}
}

func TestLazySegTreeMinLeft(t *testing.T) {
	v := []lazyS{{1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}}
	e := func() lazyS { return lazyS{0, 0} }
	merger := func(a, b lazyS) lazyS { return lazyS{a.sum + b.sum, a.len + b.len} }
	mapper := func(f lazyF, x lazyS) lazyS { return lazyS{x.sum + f.add*x.len, x.len} }
	comp := func(f, g lazyF) lazyF { return lazyF{f.add + g.add} }
	id := func() lazyF { return lazyF{0} }

	lseg := NewLazySegTree(v, e, merger, mapper, comp, id)
	// Prod(0,4)=1+2+3+4=10 >=10; Prod(1,4)=2+3+4=9 < 10
	got := lseg.MinLeft(4, func(s lazyS) bool { return s.sum < 10 })
	want := 1
	if got != want {
		t.Errorf("MinLeft(4, s.sum < 10) = %d, want %d", got, want)
	}
}
