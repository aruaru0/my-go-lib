# extgcd

ジェネリクス対応の拡張ユークリッドの互除法です。

## 関数

```go
func ExtGcd[T Integer](a, b T) (T, T, T)
```

`(g, x, y)` を返します。ただし `a*x + b*y = g = gcd(a, b)` を満たします。

## 使用例

```go
g, x, y := mylib.ExtGcd(6, 15)
// g = 3, x = -2, y = 1
// 6*(-2) + 15*1 = 3 = gcd(6, 15)
```
