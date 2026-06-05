# Types — ジェネリクス型制約

ライブラリ全体で再利用される、ジェネリクス用制約や汎用的なデータ型を定義するモジュールです。

## 主なユースケース（どういう時に使うか）

- 数値型に依存しない汎用的なアルゴリズムの型パラメータ定義の集約

## なぜ使うのか（メリット）

- 型定義を一箇所にまとめることで、ソースコード全体のコードの再利用性を高め、インターフェースの統一性を確保します。

---

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