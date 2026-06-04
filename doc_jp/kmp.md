# KMP — ジェネリック Knuth-Morris-Pratt パターンマッチング

## 型

- `KMP[T comparable]` — 任意の comparable スライス型のためのジェネリックパターンマッチャー

## 関数

- `NewKMP[T comparable](pattern []T) *KMP[T]` — 新しいパターンマッチャーを作成します
- `(*KMP[T]).Search(text []T) []int` — テキスト内のパターンのすべての開始位置を返します

## 使用例

```go
kmp := mylib.NewKMP([]int{1, 2, 3})
pos := kmp.Search([]int{1, 2, 3, 1, 2, 3})
// pos = [0, 3]

kmpStr := mylib.NewKMP([]byte("abc"))
posStr := kmpStr.Search([]byte("abcabc"))
// posStr = [0, 3]
```
