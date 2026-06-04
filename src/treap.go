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

func (t *Treap[T]) insert(n *treapNode[T], v T) *treapNode[T] {
	if n == nil {
		return &treapNode[T]{value: v, priority: t.randPriority(), count: 1}
	}
	if !t.less(v, n.value) && !t.less(n.value, v) {
		n.count++
		return n
	}
	if t.less(v, n.value) {
		n.left = t.insert(n.left, v)
		if n.left.priority < n.priority {
			n = t.rotateRight(n)
		}
	} else {
		n.right = t.insert(n.right, v)
		if n.right.priority < n.priority {
			n = t.rotateLeft(n)
		}
	}
	return n
}

func (t *Treap[T]) delete(n *treapNode[T], v T) *treapNode[T] {
	if n == nil {
		return nil
	}
	if t.less(v, n.value) {
		n.left = t.delete(n.left, v)
		return n
	}
	if t.less(n.value, v) {
		n.right = t.delete(n.right, v)
		return n
	}
	if n.count > 1 {
		n.count--
		return n
	}
	if n.left == nil && n.right == nil {
		return nil
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
	return t.delete(n, v)
}

func (t *Treap[T]) Insert(v T) {
	t.root = t.insert(t.root, v)
	t.size++
}

func (t *Treap[T]) Delete(v T) {
	t.root = t.delete(t.root, v)
	t.size--
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
