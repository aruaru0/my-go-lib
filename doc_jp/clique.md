# Clique — 最小クリーク被覆

## 関数

```go
func MinCliqueCover(N int, edges [][]int) int
```

MinCliqueCover は、無向グラフのすべての頂点を被覆するために必要な最小クリーク数を返します。ビットマスク DP を使用します。

クリークとは、すべての頂点対が直接辺で結ばれている頂点の部分集合です。
