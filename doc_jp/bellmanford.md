# BellmanFord

`BellmanFord` は、重み付き有向グラフにおいて単一始点からの最短経路を計算します。負の辺を扱い、始点から到達可能な負のサイクルを検出します。

## 型

```go
type Edge struct {
    From, To, Cost int
}
```

## 関数

```go
func BellmanFord(N, S, E int, edges []Edge) ([]int, bool)
```

- `N`: 頂点数 (0..N-1)
- `S`: 始点
- `E`: 対象頂点 (負のサイクルチェックにのみ使用されます)
- `edges`: 有向辺のリスト
- 戻り値: 距離配列と、`E` が負のサイクルから到達可能かどうかを示すブール値

## 使用例

```go
edges := []Edge{
    {0, 1, 5},
    {0, 2, 3},
    {1, 3, 2},
    {2, 3, 4},
}
dist, neg := BellmanFord(4, 0, 3, edges)
fmt.Println(dist[3]) // 7
fmt.Println(neg)     // false
```

## 計算量

O(V·E) 時間、O(V) 空間です。
