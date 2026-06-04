package mylib

import (
	"sort"
	"testing"
)

func TestAhoCorasickBasic(t *testing.T) {
	ac := NewAhoCorasick()
	ac.Add("he")
	ac.Add("she")
	ac.Add("his")
	ac.Add("hers")
	ac.Build()

	matches := ac.Search("ushers")
	got := make([]AhoMatch, len(matches))
	copy(got, matches)
	sort.Slice(got, func(i, j int) bool {
		if got[i].Pos != got[j].Pos {
			return got[i].Pos < got[j].Pos
		}
		return got[i].ID < got[j].ID
	})
	if len(got) != 3 {
		t.Fatalf("expected 3 matches, got %d: %v", len(got), got)
	}
}

func TestAhoCorasickMatchCount(t *testing.T) {
	ac := NewAhoCorasick()
	ac.Add("a")
	ac.Add("b")
	ac.Add("c")
	ac.Build()

	n := ac.MatchCount("abcabc")
	if n != 6 {
		t.Errorf("MatchCount = %d, want 6", n)
	}
}

func TestAhoCorasickEmpty(t *testing.T) {
	ac := NewAhoCorasick()
	ac.Build()
	matches := ac.Search("hello")
	if len(matches) != 0 {
		t.Errorf("expected 0 matches, got %d", len(matches))
	}
}

func TestAhoCorasickOverlap(t *testing.T) {
	ac := NewAhoCorasick()
	ac.Add("aa")
	ac.Add("aaa")
	ac.Build()

	matches := ac.Search("aaaa")
	if len(matches) == 0 {
		t.Fatal("expected matches, got none")
	}
	found := false
	for _, m := range matches {
		if m.ID == 0 && m.Pos == 2 {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected pattern 0 (aa) at position 2, got %v", matches)
	}
	n := ac.MatchCount("aaaa")
	if n != 5 {
		t.Errorf("MatchCount = %d, want 5 (aa at 0,1,2; aaa at 0,1)", n)
	}
}
