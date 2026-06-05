# Treap — ランダム化二分探索木

重複カウントをサポートする Treap (ランダム化二分探索木) を実装します。

## 関数

- `NewTreap[T any](less func(a, b T) bool) *Treap[T]` — 指定された順序で新しい Treap を作成します
- `(*Treap[T]).Insert(v T)` — 値を挿入します (重複の場合はカウントを増加)
- `(*Treap[T]).Delete(v T)` — 値のインスタンスを 1 つ削除します
- `(*Treap[T]).Find(v T) bool` — 値が存在するかどうかを返します
- `(*Treap[T]).Min() T` — 最小値を返します
- `(*Treap[T]).Max() T` — 最大値を返します
- `(*Treap[T]).Kth(k int) T` — 小さい方から k 番目 (0-indexed) の異なる要素を返します
- `(*Treap[T]).Len() int` — 異なるキーの数を返します
