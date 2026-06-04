package mylib

import "testing"

func TestRPSWin(t *testing.T) {
	tests := []struct {
		a, b, want byte
	}{
		{'R', 'R', 'R'},
		{'R', 'S', 'R'},
		{'R', 'P', 'P'},
		{'S', 'R', 'R'},
		{'S', 'S', 'S'},
		{'S', 'P', 'S'},
		{'P', 'R', 'P'},
		{'P', 'S', 'S'},
		{'P', 'P', 'P'},
	}
	for _, tt := range tests {
		got := RPSWin(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("RPSWin(%c, %c) = %c, want %c", tt.a, tt.b, got, tt.want)
		}
	}
}

type myByte byte

func TestRPSWinCustomType(t *testing.T) {
	a, b := myByte('R'), myByte('S')
	got := RPSWin(a, b)
	if got != a {
		t.Errorf("RPSWin('R', 'S') with myByte = %c, want 'R'", got)
	}
}
