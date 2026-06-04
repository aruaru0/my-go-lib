package mylib

type SccGraph struct {
	n     int
	edges [][2]int
}

type csr struct {
	start []int
	elist []int
}

func NewSccGraph(n int) *SccGraph {
	return &SccGraph{n: n}
}

func (scc *SccGraph) AddEdge(from, to int) {
	scc.edges = append(scc.edges, [2]int{from, to})
}

func (c *csr) init(n int, edges [][2]int) {
	c.start = make([]int, n+1)
	c.elist = make([]int, len(edges))
	for _, e := range edges {
		c.start[e[0]+1]++
	}
	for i := 1; i <= n; i++ {
		c.start[i] += c.start[i-1]
	}
	counter := make([]int, n+1)
	copy(counter, c.start)
	for _, e := range edges {
		c.elist[counter[e[0]]] = e[1]
		counter[e[0]]++
	}
}

func (scc *SccGraph) sccIds() (int, []int) {
	g := new(csr)
	g.init(scc.n, scc.edges)
	nowOrd, groupNum := 0, 0
	visited := make([]int, 0, scc.n)
	low := make([]int, scc.n)
	ord := make([]int, scc.n)
	ids := make([]int, scc.n)
	for i := 0; i < scc.n; i++ {
		ord[i] = -1
	}
	var dfs func(v int)
	dfs = func(v int) {
		low[v], ord[v] = nowOrd, nowOrd
		nowOrd++
		visited = append(visited, v)
		for i := g.start[v]; i < g.start[v+1]; i++ {
			to := g.elist[i]
			if ord[to] == -1 {
				dfs(to)
				if low[v] > low[to] {
					low[v] = low[to]
				}
			} else {
				if low[v] > ord[to] {
					low[v] = ord[to]
				}
			}
		}
		if low[v] == ord[v] {
			for {
				u := visited[len(visited)-1]
				visited = visited[:len(visited)-1]
				ord[u] = scc.n
				ids[u] = groupNum
				if u == v {
					break
				}
			}
			groupNum++
		}
	}
	for i := 0; i < scc.n; i++ {
		if ord[i] == -1 {
			dfs(i)
		}
	}
	for i := 0; i < len(ids); i++ {
		ids[i] = groupNum - 1 - ids[i]
	}
	return groupNum, ids
}

func (scc *SccGraph) Scc() [][]int {
	groupNum, ids := scc.sccIds()
	counts := make([]int, groupNum)
	for _, x := range ids {
		counts[x]++
	}
	groups := make([][]int, groupNum)
	for i := 0; i < groupNum; i++ {
		groups[i] = make([]int, 0, counts[i])
	}
	for i := 0; i < scc.n; i++ {
		groups[ids[i]] = append(groups[ids[i]], i)
	}
	return groups
}
