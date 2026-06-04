# Segment — 線分の交差判定

## 関数

```go
func SegmentsIntersect(ax, ay, bx, by, cx, cy, dx, dy int) bool
```

SegmentsIntersect は、線分 AB と線分 CD が交差する場合に true を返します (端点での接触は交差とみなしません)。

外積の符号テストを使用します。
