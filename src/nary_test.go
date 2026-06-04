package mylib

import (
	"testing"
)

func TestNaryNumbers(t *testing.T) {
	tests := []struct {
		N, digits int
		want      int
	}{
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 8},
		{3, 2, 9},
		{3, 1, 3},
	}
	for _, tt := range tests {
		count := 0
		for range NaryNumbers(tt.N, tt.digits) {
			count++
		}
		if count != tt.want {
			t.Errorf("NaryNumbers(%d,%d) generated %d, want %d", tt.N, tt.digits, count, tt.want)
		}
	}
}

func TestNaryNumbersValues(t *testing.T) {
	var got []string
	for s := range NaryNumbers(3, 2) {
		got = append(got, s)
	}
	if len(got) != 9 {
		t.Fatalf("expected 9, got %d", len(got))
	}
	seen := make(map[string]bool)
	for _, s := range got {
		if seen[s] {
			t.Errorf("duplicate: %s", s)
		}
		seen[s] = true
	}
}
