package mylib

import "testing"

func TestMul3x3(t *testing.T) {
	A := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	B := [3][3]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}
	got := Mul3x3(A, B)
	want := [3][3]int{
		{30, 24, 18},
		{84, 69, 54},
		{138, 114, 90},
	}
	if got != want {
		t.Errorf("Mul3x3 = %v, want %v", got, want)
	}
}

func TestMul3x3Identity(t *testing.T) {
	A := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	I := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	got := Mul3x3(A, I)
	if got != A {
		t.Errorf("Mul3x3(A, I) = %v, want %v", got, A)
	}
}

func TestMulMatrix(t *testing.T) {
	A := [][]int{{1, 2}, {3, 4}}
	B := [][]int{{5, 6}, {7, 8}}
	got := MulMatrix(A, B)
	want := [][]int{{19, 22}, {43, 50}}
	if len(got) != len(want) {
		t.Fatalf("unexpected dimensions")
	}
	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("MulMatrix[%d][%d] = %d, want %d", i, j, got[i][j], want[i][j])
			}
		}
	}
}

func TestMulMatrixNonSquare(t *testing.T) {
	A := [][]int{{1, 2, 3}, {4, 5, 6}}
	B := [][]int{{7, 8}, {9, 10}, {11, 12}}
	got := MulMatrix(A, B)
	want := [][]int{{58, 64}, {139, 154}}
	if len(got) != len(want) {
		t.Fatalf("unexpected dimensions")
	}
	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("MulMatrix[%d][%d] = %d, want %d", i, j, got[i][j], want[i][j])
			}
		}
	}
}

func TestMulMatrixInt64(t *testing.T) {
	A := [][]int64{{1, 2}, {3, 4}}
	B := [][]int64{{5, 6}, {7, 8}}
	got := MulMatrix(A, B)
	want := [][]int64{{19, 22}, {43, 50}}
	for i := range want {
		for j := range want[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("MulMatrix[%d][%d] = %d, want %d", i, j, got[i][j], want[i][j])
			}
		}
	}
}
