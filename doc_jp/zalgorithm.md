# ZAlgorithm — Z アルゴリズム (ジェネリック)

## 関数

- `ZAlgorithm[T comparable](s []T) []int` — comparable なスライスの Z 配列を計算します
- `ZAlgorithmString(s string) []int` — 文字列の Z 配列を計算します

Z 配列 `z` において、`z[i]` は `s` と `s[i:]` の最長共通接頭辞の長さです。

## 使用例

```go
z := mylib.ZAlgorithmString("aaaaa")
// z = [5 4 3 2 1]

z2 := mylib.ZAlgorithmString("abcabcabc")
// z2 = [9 0 0 6 0 0 3 0 0]
```
