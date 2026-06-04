package mylib

type Route struct {
	Path []int
}

func BellmanFordRoute(N, S int, edges []Edge) ([]int, []Route, bool) {
	const inf = 1 << 60
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = inf
	}
	d[S] = 0
	r := make([]Route, N)
	r[S].Path = []int{S}

	negativeLoop := false
	for i := 0; i < N; i++ {
		for _, e := range edges {
			if d[e.To] > d[e.From]+e.Cost {
				d[e.To] = d[e.From] + e.Cost
				r[e.To].Path = append(r[e.From].Path, e.To)
				if i == N-1 {
					negativeLoop = true
				}
			}
		}
	}

	return d, r, negativeLoop
}
