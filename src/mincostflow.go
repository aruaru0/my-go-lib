package mylib

import "container/heap"

type _mcfEdge struct {
	to   int
	rev  int
	capa int
	cost int
}

type MCFEdge struct {
	From, To, Capa, Flow, Cost int
}

type MinCostFlow struct {
	n   int
	pos [][2]int
	g   [][]_mcfEdge
}

func NewMinCostFlow(n int) *MinCostFlow {
	return &MinCostFlow{n: n, g: make([][]_mcfEdge, n)}
}

func (mcf *MinCostFlow) AddEdge(from, to, capa, cost int) int {
	m := len(mcf.pos)
	mcf.pos = append(mcf.pos, [2]int{from, len(mcf.g[from])})
	mcf.g[from] = append(mcf.g[from], _mcfEdge{to, len(mcf.g[to]), capa, cost})
	mcf.g[to] = append(mcf.g[to], _mcfEdge{from, len(mcf.g[from]) - 1, 0, -cost})
	return m
}

func (mcf *MinCostFlow) GetEdge(i int) MCFEdge {
	e := mcf.g[mcf.pos[i][0]][mcf.pos[i][1]]
	re := mcf.g[e.to][e.rev]
	return MCFEdge{mcf.pos[i][0], e.to, e.capa + re.capa, re.capa, e.cost}
}

func (mcf *MinCostFlow) Edges() []MCFEdge {
	m := len(mcf.pos)
	res := make([]MCFEdge, m)
	for i := 0; i < m; i++ {
		res[i] = mcf.GetEdge(i)
	}
	return res
}

func (mcf *MinCostFlow) Flow(s, t int) [2]int {
	res := mcf.Slope(s, t)
	return res[len(res)-1]
}

func (mcf *MinCostFlow) FlowL(s, t, flowLim int) [2]int {
	res := mcf.SlopeL(s, t, flowLim)
	return res[len(res)-1]
}

func (mcf *MinCostFlow) Slope(s, t int) [][2]int {
	return mcf.SlopeL(s, t, 1<<60)
}

func (mcf *MinCostFlow) SlopeL(s, t, flowLim int) [][2]int {
	const inf = 1 << 60
	n := mcf.n
	dual := make([]int, n)
	dist := make([]int, n)
	pv := make([]int, n)
	pe := make([]int, n)
	vis := make([]bool, n)

	dualRef := func() bool {
		for i := 0; i < n; i++ {
			dist[i] = inf
			pv[i] = -1
			pe[i] = -1
			vis[i] = false
		}
		pq := &_mcfPQ{}
		heap.Init(pq)
		dist[s] = 0
		heap.Push(pq, &_mcfItem{value: s, priority: 0})
		for pq.Len() > 0 {
			v := heap.Pop(pq).(*_mcfItem).value
			if vis[v] {
				continue
			}
			vis[v] = true
			if v == t {
				break
			}
			for i := 0; i < len(mcf.g[v]); i++ {
				e := mcf.g[v][i]
				if vis[e.to] || e.capa == 0 {
					continue
				}
				cost := e.cost - dual[e.to] + dual[v]
				if dist[e.to]-dist[v] > cost {
					dist[e.to] = dist[v] + cost
					pv[e.to] = v
					pe[e.to] = i
					heap.Push(pq, &_mcfItem{value: e.to, priority: dist[e.to]})
				}
			}
		}
		if !vis[t] {
			return false
		}
		for v := 0; v < n; v++ {
			if !vis[v] {
				continue
			}
			dual[v] -= dist[t] - dist[v]
		}
		return true
	}

	flow := 0
	cost := 0
	res := make([][2]int, 0, n)
	res = append(res, [2]int{flow, cost})
	for flow < flowLim {
		if !dualRef() {
			break
		}
		c := flowLim - flow
		for v := t; v != s; v = pv[v] {
			c = min(c, mcf.g[pv[v]][pe[v]].capa)
		}
		for v := t; v != s; v = pv[v] {
			mcf.g[pv[v]][pe[v]].capa -= c
			mcf.g[v][mcf.g[pv[v]][pe[v]].rev].capa += c
		}
		d := -dual[s]
		flow += c
		cost += c * d
		res = append(res, [2]int{flow, cost})
	}
	return res
}

type _mcfItem struct {
	value    int
	priority int
	index    int
}

type _mcfPQ []*_mcfItem

func (pq _mcfPQ) Len() int { return len(pq) }

func (pq _mcfPQ) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq _mcfPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *_mcfPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*_mcfItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *_mcfPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
