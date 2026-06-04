package mylib

import "math"

type LCA struct {
	n      int
	log    int
	parent [][]int
	dep    []int
	G      [][]int
}

func NewLCA(n int) *LCA {
	l := &LCA{}
	l.n = n
	l.log = int(math.Log2(float64(n))) + 1
	l.parent = make([][]int, l.log)
	for i := 0; i < l.log; i++ {
		l.parent[i] = make([]int, n)
	}
	l.dep = make([]int, n)
	l.G = make([][]int, n)
	return l
}

func (l *LCA) AddEdge(from, to int) {
	l.G[from] = append(l.G[from], to)
	l.G[to] = append(l.G[to], from)
}

func (l *LCA) dfs(v, p, d int) {
	l.parent[0][v] = p
	l.dep[v] = d
	for _, to := range l.G[v] {
		if to == p {
			continue
		}
		l.dfs(to, v, d+1)
	}
}

func (l *LCA) Build(root int) {
	l.dfs(root, -1, 0)
	for k := 0; k+1 < l.log; k++ {
		for v := 0; v < l.n; v++ {
			if l.parent[k][v] < 0 {
				l.parent[k+1][v] = -1
			} else {
				l.parent[k+1][v] = l.parent[k][l.parent[k][v]]
			}
		}
	}
}

func (l *LCA) LCA(u, v int) int {
	if l.dep[u] > l.dep[v] {
		u, v = v, u
	}
	for k := 0; k < l.log; k++ {
		if (l.dep[v]-l.dep[u])>>k&1 == 1 {
			v = l.parent[k][v]
		}
	}
	if u == v {
		return u
	}
	for k := l.log - 1; k >= 0; k-- {
		if l.parent[k][u] != l.parent[k][v] {
			u = l.parent[k][u]
			v = l.parent[k][v]
		}
	}
	return l.parent[0][u]
}

func (l *LCA) Dist(u, v int) int {
	return l.dep[u] + l.dep[v] - 2*l.dep[l.LCA(u, v)]
}
