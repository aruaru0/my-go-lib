# compress

順序付き型のための座標圧縮を行います。

## 関数

```go
func Compress[T Ordered](s []T) []int
```

## Compress

`s` の各要素に対応する順位 (0 から始まる) を整数スライスとして返します。順位はソートされた一意な値に基づきます。`Ordered` 制約を満たす任意の型 (int 系、float 系、string) をサポートします。
