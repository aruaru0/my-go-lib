package mylib

type TwoSat struct {
	n        int
	answer   []bool
	sccGraph *SccGraph
}

func NewTwoSat(n int) *TwoSat {
	return &TwoSat{
		n:        n,
		answer:   make([]bool, n),
		sccGraph: NewSccGraph(n * 2),
	}
}

func (ts *TwoSat) AddClause(i int, f bool, j int, g bool) {
	ts.sccGraph.AddEdge(2*i+boolToInt(f, 0, 1), 2*j+boolToInt(g, 1, 0))
	ts.sccGraph.AddEdge(2*j+boolToInt(g, 0, 1), 2*i+boolToInt(f, 1, 0))
}

func boolToInt(f bool, a, b int) int {
	if f {
		return a
	}
	return b
}

func (ts *TwoSat) Satisfiable() bool {
	_, id := ts.sccGraph.sccIds()
	for i := 0; i < ts.n; i++ {
		if id[i*2] == id[2*i+1] {
			return false
		}
		ts.answer[i] = id[2*i] < id[2*i+1]
	}
	return true
}

func (ts *TwoSat) Answer() []bool {
	return ts.answer
}
