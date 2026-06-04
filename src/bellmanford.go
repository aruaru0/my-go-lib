package mylib

type Edge struct {
	From, To, Cost int
}

func BellmanFord(N, S, E int, edges []Edge) ([]int, bool) {
	const inf = 1 << 60
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = inf
	}
	d[S] = 0
	for i := 0; i < N-1; i++ {
		for _, e := range edges {
			if d[e.From] != inf && d[e.To] > d[e.From]+e.Cost {
				d[e.To] = d[e.From] + e.Cost
			}
		}
	}

	negative := make([]bool, N)
	for i := 0; i < N; i++ {
		for _, e := range edges {
			if d[e.From] != inf && d[e.To] > d[e.From]+e.Cost {
				d[e.To] = d[e.From] + e.Cost
				negative[e.To] = true
			}
			negative[e.To] = negative[e.To] || negative[e.From]
		}
	}

	return d, negative[E]
}
