package mylib

import "math/rand"

type treapNode[T any] struct {
	value    T
	priority uint64
	count    int
	left     *treapNode[T]
	right    *treapNode[T]
}

type Treap[T any] struct {
	root *treapNode[T]
	less func(a, b T) bool
	size int
}

func NewTreap[T any](less func(a, b T) bool) *Treap[T] {
	return &Treap[T]{less: less}
}

func (t *Treap[T]) randPriority() uint64 {
	return uint64(rand.Int63())
}

func (t *Treap[T]) rotateRight(n *treapNode[T]) *treapNode[T] {
	l := n.left
	n.left = l.right
	l.right = n
	return l
}

func (t *Treap[T]) rotateLeft(n *treapNode[T]) *treapNode[T] {
	r := n.right
	n.right = r.left
	r.left = n
	return r
}

func (t *Treap[T]) insert(n *treapNode[T], v T) (*treapNode[T], bool) {
	if n == nil {
		return &treapNode[T]{value: v, priority: t.randPriority(), count: 1}, true
	}
	if !t.less(v, n.value) && !t.less(n.value, v) {
		n.count++
		return n, false
	}
	var inserted bool
	if t.less(v, n.value) {
		n.left, inserted = t.insert(n.left, v)
		if n.left.priority < n.priority {
			n = t.rotateRight(n)
		}
	} else {
		n.right, inserted = t.insert(n.right, v)
		if n.right.priority < n.priority {
			n = t.rotateLeft(n)
		}
	}
	return n, inserted
}

func (t *Treap[T]) delete(n *treapNode[T], v T) (*treapNode[T], bool) {
	if n == nil {
		return nil, false
	}
	if t.less(v, n.value) {
		var deleted bool
		n.left, deleted = t.delete(n.left, v)
		return n, deleted
	}
	if t.less(n.value, v) {
		var deleted bool
		n.right, deleted = t.delete(n.right, v)
		return n, deleted
	}
	if n.count > 1 {
		n.count--
		return n, false
	}
	if n.left == nil && n.right == nil {
		return nil, true
	}
	if n.left == nil {
		n = t.rotateLeft(n)
	} else if n.right == nil {
		n = t.rotateRight(n)
	} else if n.left.priority < n.right.priority {
		n = t.rotateRight(n)
	} else {
		n = t.rotateLeft(n)
	}
	var deleted bool
	n, deleted = t.delete(n, v)
	return n, deleted
}

func (t *Treap[T]) Insert(v T) {
	var inserted bool
	t.root, inserted = t.insert(t.root, v)
	if inserted {
		t.size++
	}
}

func (t *Treap[T]) Delete(v T) {
	var deleted bool
	t.root, deleted = t.delete(t.root, v)
	if deleted {
		t.size--
	}
}

func (t *Treap[T]) Find(v T) bool {
	cur := t.root
	for cur != nil {
		if t.less(v, cur.value) {
			cur = cur.left
		} else if t.less(cur.value, v) {
			cur = cur.right
		} else {
			return true
		}
	}
	return false
}

func (t *Treap[T]) Min() T {
	cur := t.root
	for cur.left != nil {
		cur = cur.left
	}
	return cur.value
}

func (t *Treap[T]) Max() T {
	cur := t.root
	for cur.right != nil {
		cur = cur.right
	}
	return cur.value
}

func (t *Treap[T]) Len() int {
	return t.size
}

func countSubtree[T any](n *treapNode[T]) int {
	if n == nil {
		return 0
	}
	return n.count + countSubtree(n.left) + countSubtree(n.right)
}

func (t *Treap[T]) Kth(k int) T {
	n := t.root
	for n != nil {
		left := countSubtree(n.left)
		if k < left {
			n = n.left
		} else if k < left+n.count {
			return n.value
		} else {
			k -= left + n.count
			n = n.right
		}
	}
	var zero T
	return zero
}
