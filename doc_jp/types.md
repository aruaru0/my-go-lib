# Types — ジェネリクス型制約

ライブラリ全体で使用されるジェネリクス型制約を定義します。

## 型制約

- `Ordered` — 順序比較（`<`, `>`, `<=`, `>=`）をサポートする型。整数、符号なし整数、浮動小数点数、文字列を含みます。
- `Number` — 数値演算をサポートする型。整数、符号なし整数、浮動小数点数を含みます。

## 定義

```go
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}
```
