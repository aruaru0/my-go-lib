# SuffixArray — SA-IS 接尾辞配列構築

SA-IS アルゴリズムを用いた高速な接尾辞配列構築を行います。

## 関数

- `SuffixArray(s []int, upper int) []int` — [0, upper] の値を持つ整数スライスの接尾辞配列を返します
- `SuffixArrayString(s string) []int` — 文字列 (バイト値) の接尾辞配列を返します
- `SuffixArrayInt(s []int) []int` — 任意の整数スライスの接尾辞配列を返します (内部で値の圧縮を行います)
- `LcpArray(s, sa []int) []int` — 整数スライスとその接尾辞配列から LCP 配列を計算します
- `LcpArrayString(s string, sa []int) []int` — 文字列とその接尾辞配列から LCP 配列を計算します

## 使用例

```go
sa := mylib.SuffixArrayString("banana")
// sa = [5 3 1 0 4 2]

lcp := mylib.LcpArrayString("banana", sa)
// lcp = [1 3 0 0 2]
```
