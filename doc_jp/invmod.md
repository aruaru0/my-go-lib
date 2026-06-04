# invmod

O(n) でモジュラ逆数を事前計算します。

## 関数

```go
func InvModTable(n, m int) []int
```

長さ `n` の配列を返します。`ret[i] = i^{-1} mod m` (ただし `1 <= i < n`) です。漸化式 `inv[i] = m - m/i * inv[m%i] % m` を使用します。

## 使用例

```go
inv := mylib.InvModTable(10, 13)
// inv[1..9] は 13 を法とするモジュラ逆数
```
