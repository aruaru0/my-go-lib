package mylib

type Node struct {
	To []int
}

func DFSRoute(nodes []Node, from, to int) []int {
	var route []int
	var dfs func(v, target, prev int) bool
	dfs = func(v, target, prev int) bool {
		if v == target {
			route = append(route, v)
			return true
		}
		for _, nxt := range nodes[v].To {
			if nxt == prev {
				continue
			}
			if dfs(nxt, target, v) {
				route = append(route, v)
				return true
			}
		}
		return false
	}
	if dfs(from, to, -1) {
		for i, j := 0, len(route)-1; i < j; i, j = i+1, j-1 {
			route[i], route[j] = route[j], route[i]
		}
		return route
	}
	return nil
}
