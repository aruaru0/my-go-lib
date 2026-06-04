package mylib

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationsInt(t *testing.T) {
	list := []int{1, 2, 3, 4}
	var got [][]int
	for comb := range Combinations(list, 2) {
		c := make([]int, len(comb))
		copy(c, comb)
		got = append(got, c)
	}
	want := [][]int{
		{1, 2}, {1, 3}, {1, 4},
		{2, 3}, {2, 4},
		{3, 4},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCombinationsChoose0(t *testing.T) {
	var got [][]int
	for comb := range Combinations([]int{1, 2}, 0) {
		got = append(got, comb)
	}
	if len(got) != 1 || len(got[0]) != 0 {
		t.Errorf("expected [[]], got %v", got)
	}
}

func TestCombinationsChooseAll(t *testing.T) {
	list := []int{1, 2, 3}
	var got [][]int
	for comb := range Combinations(list, 3) {
		got = append(got, comb)
	}
	want := [][]int{{1, 2, 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCombinationsChooseMore(t *testing.T) {
	count := 0
	for range Combinations([]int{1, 2}, 3) {
		count++
	}
	if count != 0 {
		t.Errorf("expected 0, got %d", count)
	}
}

func TestCombinationsString(t *testing.T) {
	list := []string{"a", "b", "c"}
	var got [][]string
	for comb := range Combinations(list, 2) {
		got = append(got, comb)
	}
	want := [][]string{
		{"a", "b"}, {"a", "c"}, {"b", "c"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func collectRec(n, k int) [][]int {
	var res [][]int
	for comb := range CombinationsRec(n, k) {
		c := make([]int, len(comb))
		copy(c, comb)
		res = append(res, c)
	}
	return res
}

func TestCombinationsRec(t *testing.T) {
	got := collectRec(5, 3)
	if len(got) != 10 {
		t.Fatalf("expected 10 combinations, got %d", len(got))
	}
	for _, c := range got {
		if len(c) != 3 {
			t.Errorf("expected len 3, got %v", c)
		}
		sorted := sort.IntsAreSorted(c)
		if !sorted {
			t.Errorf("expected sorted, got %v", c)
		}
	}
}

func TestCombinationsRecSmall(t *testing.T) {
	tests := []struct {
		n, k int
		want int
	}{
		{4, 0, 1},
		{4, 1, 4},
		{4, 2, 6},
		{4, 3, 4},
		{4, 4, 1},
		{4, 5, 0},
	}
	for _, tt := range tests {
		got := collectRec(tt.n, tt.k)
		if len(got) != tt.want {
			t.Errorf("CombinationsRec(%d,%d): got %d, want %d", tt.n, tt.k, len(got), tt.want)
		}
	}
}
