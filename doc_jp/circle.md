# Circle — 円と円の交点

## 型

```go
type PointF struct {
	X, Y float64
}
```

## 関数

```go
func CircleCross(p0, p1 PointF, r1, r2 float64) []PointF
```

CircleCross は 2 つの円の交点を返します。円が交差しない場合 (離れている、または同心円) は nil を返します。
