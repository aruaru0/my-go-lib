# xorsum

配列内のすべての順序対に対する XOR の総和を求めます。

## 関数

```go
func XorSum(dist []int) int64
```

## XorSum

すべての対 `(i, j)` (ただし `i < j`) について `dist[i] ^ dist[j]` の総和を計算します。ビットごとの計算を使用します。各ビットについて 1 と 0 の個数を数え、`ones * zeros * 2^bit` の寄与を計算します。
