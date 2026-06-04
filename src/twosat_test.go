package mylib

import (
	"testing"
)

func TestTwoSatSimple(t *testing.T) {
	ts := NewTwoSat(2)
	ts.AddClause(0, true, 1, true)
	ts.AddClause(0, false, 1, false)
	if !ts.Satisfiable() {
		t.Error("expected satisfiable")
	}
	ans := ts.Answer()
	if !ans[0] && !ans[1] {
		t.Error("at least one must be true (x0 ∨ x1)")
	}
	if ans[0] && ans[1] {
		t.Error("at least one must be false (¬x0 ∨ ¬x1)")
	}
}

func TestTwoSatUnsatisfiable(t *testing.T) {
	ts := NewTwoSat(1)
	ts.AddClause(0, true, 0, true)
	ts.AddClause(0, false, 0, false)
	if ts.Satisfiable() {
		t.Error("expected unsatisfiable")
	}
}

func TestTwoSatThreeClauses(t *testing.T) {
	ts := NewTwoSat(3)
	ts.AddClause(0, true, 1, true)
	ts.AddClause(1, true, 2, true)
	ts.AddClause(2, true, 0, true)
	if !ts.Satisfiable() {
		t.Error("expected satisfiable")
	}
}

func TestTwoSatContradiction(t *testing.T) {
	ts := NewTwoSat(2)
	ts.AddClause(0, true, 0, true)
	ts.AddClause(0, false, 1, true)
	ts.AddClause(1, false, 1, false)
	if ts.Satisfiable() {
		t.Error("expected unsatisfiable")
	}
}
