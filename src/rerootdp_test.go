package mylib

import "testing"

func TestTreeDistancesLine(t *testing.T) {
	N := 4
	node := make([][]int, N)
	node[0] = []int{1}
	node[1] = []int{0, 2}
	node[2] = []int{1, 3}
	node[3] = []int{2}

	dist := TreeDistances(N, node)
	expected := []int{3, 2, 2, 3}
	for i := 0; i < N; i++ {
		if dist[i] != expected[i] {
			t.Errorf("node %d: got %d, want %d", i, dist[i], expected[i])
		}
	}
}

func TestTreeDistancesStar(t *testing.T) {
	N := 5
	node := make([][]int, N)
	node[0] = []int{1, 2, 3, 4}
	node[1] = []int{0}
	node[2] = []int{0}
	node[3] = []int{0}
	node[4] = []int{0}

	dist := TreeDistances(N, node)
	if dist[0] != 1 {
		t.Errorf("center expected 1, got %d", dist[0])
	}
	for i := 1; i < N; i++ {
		if dist[i] != 2 {
			t.Errorf("leaf %d expected 2, got %d", i, dist[i])
		}
	}
}

func TestTreeDistancesWithNodeLine(t *testing.T) {
	N := 4
	node := make([][]int, N)
	node[0] = []int{1}
	node[1] = []int{0, 2}
	node[2] = []int{1, 3}
	node[3] = []int{2}

	dists, idxs := TreeDistancesWithNode(N, node)
	expectedDist := []int{3, 2, 2, 3}
	for i := 0; i < N; i++ {
		if dists[i] != expectedDist[i] {
			t.Errorf("dist node %d: got %d, want %d", i, dists[i], expectedDist[i])
		}
	}
	if idxs[0] != 3 || idxs[3] != 0 {
		t.Errorf("farthest node from 0 should be 3, got %d", idxs[0])
	}
}
