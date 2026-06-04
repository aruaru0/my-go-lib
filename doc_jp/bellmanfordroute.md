# BellmanFordRoute

`BellmanFordRoute` は Bellman-Ford 法を拡張し、始点から各頂点への実際の経路 (ルート) も追跡します。

## 型

```go
type Edge struct {
    From, To, Cost int
}

type Route struct {
    Path []int
}
```

## 関数

```go
func BellmanFordRoute(N, S int, edges []Edge) ([]int, []Route, bool)
```

- `N`: 頂点数
- `S`: 始点
- `edges`: 有向辺のリスト
- 戻り値: 距離配列、経路配列 (各頂点ごと)、負のサイクルの有無を示すブール値

## 使用例

```go
edges := []Edge{{0, 1, 2}, {1, 2, 3}, {0, 2, 6}}
dist, routes, neg := BellmanFordRoute(3, 0, edges)
fmt.Println(dist[2])       // 5
fmt.Println(routes[2].Path) // [0 1 2]
```

## 計算量

O(V·E) 時間、ルートについては最悪で O(V·E) 空間です。
