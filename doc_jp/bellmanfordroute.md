# BellmanFordRoute

ベルマンフォード法を用いて最短経路を求めると同時に、どの頂点を順に通過したか（具体的な経路）を復元する機能を提供します。

## 主なユースケース（どういう時に使うか）

- 負の辺を含むグラフでの、単に最短距離を知るだけでなく実際の移動ルートを出力したい場合
- 負の閉路検出時に、その閉路を構成する具体的な頂点の遷移順を特定したい場合

## なぜ使うのか（メリット）

- 最短経路の計算と並行して先行頂点（前段階の頂点）を記録するため、最短経路の探索後に $O(V)$ の手間で具体的な移動ルートを復元できます。

---

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