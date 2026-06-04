# garner

Garner のアルゴリズムによる中国剰余定理です。

## 関数

```go
func Garner(x, m []int, mod int) int
```

`x[i] ≡ result (mod m[i])` が与えられたとき、`result` を `mod` で割った剰余を計算します。`mod` が 0 の場合は、正確な整数を計算します (大きな値ではオーバーフローの可能性があります)。

## 使用例

```go
x := []int{2, 3, 2}
m := []int{3, 5, 7}
result := mylib.Garner(x, m, 0) // 23
// 23 ≡ 2 (mod 3), 23 ≡ 3 (mod 5), 23 ≡ 2 (mod 7)
```
