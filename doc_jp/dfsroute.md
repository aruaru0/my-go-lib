# DFSRoute

`DFSRoute` は、深さ優先探索を使用して木または非巡回グラフ内の 2 つのノード間の経路を探索します。

## 型

```go
type Node struct {
    To []int
}
```

## 関数

```go
func DFSRoute(nodes []Node, from, to int) []int
```

`from` から `to` への経路を頂点インデックスのスライスとして返します。経路が存在しない場合は `nil` を返します。グラフはサイクルを含んではいけません。

## 使用例

```go
nodes := []Node{
    {To: []int{1, 2}},
    {To: []int{0, 3}},
    {To: []int{0, 3}},
    {To: []int{1, 2}},
}
route := DFSRoute(nodes, 0, 3)
fmt.Println(route) // [0 1 3] または [0 2 3]
```

## 計算量

O(V + E) 時間、O(V) 空間です。
