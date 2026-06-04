# RerootDP — 全方位木 DP による木の距離

## 関数

```go
func TreeDistances(N int, node [][]int) []int
func TreeDistancesWithNode(N int, node [][]int) ([]int, []int)
```

TreeDistances は、重みなし木 (隣接リスト) において、各ノードから他の任意のノードへの最遠距離を計算します。

TreeDistancesWithNode はさらに、最遠ノードのインデックスを返します (同点の場合は最大のインデックス)。

Rerooting (2 回の DFS) 手法を使用します。1 回目の DFS で部分木の高さを計算し、2 回目の DFS で親側の情報を子に伝搬します。
