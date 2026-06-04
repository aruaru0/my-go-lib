# combinations

組み合わせ列挙のユーティリティです。

## 関数

```go
func Combinations[T any](list []T, choose int) <-chan []T
func CombinationsRec(n, k int) <-chan []int
```

## Combinations

`list` から `choose` 個の要素を選ぶすべての組み合わせを生成し、goroutine を使用してチャネルを通じてストリーミングします。

## CombinationsRec

再帰的なバックトラッキングを使用して、すべてのインデックスの組み合わせ (0 から n-1 から k 個を選択) をチャネルを通じて生成します。
