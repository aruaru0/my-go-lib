package mylib

type _dinicEdge struct {
	to, cap, rev int
}

type Dinic struct {
	g     [][]_dinicEdge
	level []int
	iter  []int
}

func NewDinic(N int) *Dinic {
	return &Dinic{
		g: make([][]_dinicEdge, N),
	}
}

func (d *Dinic) AddEdge(from, to, cap int) {
	d.g[from] = append(d.g[from], _dinicEdge{to, cap, len(d.g[to])})
	d.g[to] = append(d.g[to], _dinicEdge{from, 0, len(d.g[from]) - 1})
}

func (d *Dinic) bfs(s int) {
	d.level = make([]int, len(d.g))
	for i := range d.level {
		d.level[i] = -1
	}
	que := []int{s}
	d.level[s] = 0
	for len(que) > 0 {
		v := que[0]
		que = que[1:]
		for _, e := range d.g[v] {
			if e.cap > 0 && d.level[e.to] < 0 {
				d.level[e.to] = d.level[v] + 1
				que = append(que, e.to)
			}
		}
	}
}

func (d *Dinic) dfs(v, t, f int) int {
	if v == t {
		return f
	}
	for i := d.iter[v]; i < len(d.g[v]); i++ {
		e := &d.g[v][i]
		if e.cap > 0 && d.level[v] < d.level[e.to] {
			ret := d.dfs(e.to, t, min(f, e.cap))
			if ret > 0 {
				e.cap -= ret
				d.g[e.to][e.rev].cap += ret
				return ret
			}
		}
		d.iter[v]++
	}
	return 0
}

func (d *Dinic) MaxFlow(s, t int) int {
	flow := 0
	N := len(d.g)
	for {
		d.bfs(s)
		if d.level[t] < 0 {
			return flow
		}
		d.iter = make([]int, N)
		for {
			f := d.dfs(s, t, 1<<60)
			if f <= 0 {
				break
			}
			flow += f
		}
	}
}
