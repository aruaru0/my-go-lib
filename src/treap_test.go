package mylib

import "testing"

func intLess(a, b int) bool { return a < b }

func TestTreapInsertAndFind(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(7)
	tr.Insert(1)

	if !tr.Find(5) {
		t.Error("Find(5) = false, want true")
	}
	if !tr.Find(3) {
		t.Error("Find(3) = false, want true")
	}
	if !tr.Find(7) {
		t.Error("Find(7) = false, want true")
	}
	if !tr.Find(1) {
		t.Error("Find(1) = false, want true")
	}
	if tr.Find(4) {
		t.Error("Find(4) = true, want false")
	}
}

func TestTreapDelete(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(7)
	tr.Delete(3)
	if tr.Find(3) {
		t.Error("Find(3) after Delete = true, want false")
	}
	if !tr.Find(5) {
		t.Error("Find(5) after Delete = false, want true")
	}
}

func TestTreapMinMax(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(3)
	tr.Insert(7)
	tr.Insert(1)
	tr.Insert(9)

	if got := tr.Min(); got != 1 {
		t.Errorf("Min() = %d, want 1", got)
	}
	if got := tr.Max(); got != 9 {
		t.Errorf("Max() = %d, want 9", got)
	}
}

func TestTreapDuplicates(t *testing.T) {
	tr := NewTreap(intLess)
	tr.Insert(5)
	tr.Insert(5)
	tr.Insert(5)
	tr.Delete(5)
	if !tr.Find(5) {
		t.Error("Find(5) after one Delete = false, want true (2 remaining)")
	}
	tr.Delete(5)
	tr.Delete(5)
	if tr.Find(5) {
		t.Error("Find(5) after three Deletes = true, want false")
	}
}

func TestTreapStringKeys(t *testing.T) {
	tr := NewTreap(func(a, b string) bool { return a < b })
	tr.Insert("c")
	tr.Insert("a")
	tr.Insert("b")
	if got := tr.Min(); got != "a" {
		t.Errorf("Min() = %q, want %q", got, "a")
	}
	if got := tr.Max(); got != "c" {
		t.Errorf("Max() = %q, want %q", got, "c")
	}
	if !tr.Find("b") {
		t.Error("Find(b) = false, want true")
	}
}

func TestTreapSizeAndKth(t *testing.T) {
	tr := NewTreap(intLess)
	if tr.Len() != 0 {
		t.Errorf("Len() at init = %d, want 0", tr.Len())
	}

	// 存在しないキーの削除
	tr.Delete(10)
	if tr.Len() != 0 {
		t.Errorf("Len() after delete non-existent = %d, want 0", tr.Len())
	}

	// 重複挿入時のユニークキー数
	tr.Insert(5)
	tr.Insert(5)
	tr.Insert(3)
	if tr.Len() != 2 {
		t.Errorf("Len() after inserts = %d, want 2", tr.Len())
	}

	// 重複ありの削除
	tr.Delete(5)
	if tr.Len() != 2 {
		t.Errorf("Len() after one delete of duplicate = %d, want 2 (since 5 still exists)", tr.Len())
	}

	tr.Delete(5)
	if tr.Len() != 1 {
		t.Errorf("Len() after deleting all 5 = %d, want 1", tr.Len())
	}

	// Kthのテスト
	tr.Insert(1)
	tr.Insert(7)
	tr.Insert(5) // 現在: 3, 1, 7, 5 (ソート順: 1, 3, 5, 7)
	// ユニーク数: 4 (1, 3, 5, 7)
	if tr.Len() != 4 {
		t.Errorf("Len() before Kth = %d, want 4", tr.Len())
	}

	if got := tr.Kth(0); got != 1 {
		t.Errorf("Kth(0) = %d, want 1", got)
	}
	if got := tr.Kth(1); got != 3 {
		t.Errorf("Kth(1) = %d, want 3", got)
	}
	if got := tr.Kth(2); got != 5 {
		t.Errorf("Kth(2) = %d, want 5", got)
	}
	if got := tr.Kth(3); got != 7 {
		t.Errorf("Kth(3) = %d, want 7", got)
	}
}
