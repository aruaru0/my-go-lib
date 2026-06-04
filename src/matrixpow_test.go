package mylib

import (
	"testing"
)

func TestPowMatrix(t *testing.T) {
	A := [][]int{{1, 1}, {1, 0}}
	P := PowMatrix(A, 10, 1000000007)
	if P[0][0] != 89 || P[0][1] != 55 || P[1][0] != 55 || P[1][1] != 34 {
		t.Errorf("PowMatrix(A,10) = %v, want [[89 55] [55 34]]", P)
	}
}

func TestPowMatrixIdentity(t *testing.T) {
	A := [][]int{{1, 2}, {3, 4}}
	P := PowMatrix(A, 0, 1000000007)
	if P[0][0] != 1 || P[0][1] != 0 || P[1][0] != 0 || P[1][1] != 1 {
		t.Errorf("PowMatrix(A,0) = %v, want identity", P)
	}
}

func TestPowMatrixGeneric(t *testing.T) {
	A := [][]int64{{2, 0}, {0, 2}}
	P := PowMatrix(A, 3, int64(100))
	if P[0][0] != 8 || P[1][1] != 8 {
		t.Errorf("PowMatrix(A,3) = %v, want [[8 0] [0 8]]", P)
	}
}
