# TwoSat (2-SAT ソルバー)

`TwoSat` は 2-Satisfiability 問題を解きます。ブール変数と `(x_i = f) ∨ (x_j = g)` 形式の節が与えられたとき、充足割当が存在するかどうかを判定し、見つけます。

## コンストラクタ

```go
func NewTwoSat(n int) *TwoSat
```

`n` 個のブール変数 (0..n-1) を持つ新しい 2-SAT インスタンスを作成します。

## メソッド

```go
func (ts *TwoSat) AddClause(i int, f bool, j int, g bool)
```
節 `(x_i == f) ∨ (x_j == g)` を追加します。

```go
func (ts *TwoSat) Satisfiable() bool
```
充足割当が存在する場合は `true` を返します。

```go
func (ts *TwoSat) Answer() []bool
```
充足割当を返します (`Satisfiable()` が `true` を返した場合のみ有効です)。

## 使用例

```go
ts := NewTwoSat(2)
ts.AddClause(0, true, 1, true)
ts.AddClause(0, false, 1, false)
if ts.Satisfiable() {
    fmt.Println(ts.Answer()) // [true true]
}
```

## 計算量

O(N + M) 時間、O(N + M) 空間です。N は変数、M は節の数です。
