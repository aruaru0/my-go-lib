package mylib

import (
	"math"
	"testing"
)

func TestGachaExpectation(t *testing.T) {
	tests := []struct {
		n    int
		want float64
	}{
		{1, 1.0},
		{2, 3.0},
		{3, 5.5},
		{4, 8.333333333333334},
		{5, 11.416666666666666},
	}
	for _, tt := range tests {
		got := GachaExpectation(tt.n)
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("GachaExpectation(%d) = %.10f, want %.10f", tt.n, got, tt.want)
		}
	}
}
