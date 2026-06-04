# ConvexHull — Monotone Chain

## 型

```go
type Point struct {
	X, Y int
}
```

## 関数

```go
func ConvexHull(pts []Point) []Point
```

ConvexHull は、Monotone Chain (Andrew のアルゴリズム) を使用して、点集合の凸包を左端の点から反時計回り順で返します。

すべての点が同一直線上にある場合は 2 つの端点を返します。単一の点の場合はその点を返します。
