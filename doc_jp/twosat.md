# TwoSat (2-SAT ソルバー)

各論理変数が 2 つの変数に関する論理和の連言の形で与えられる充足可能性問題（2-SAT）を解き、式を満たす真偽値の割り当てが存在するか判定するアルゴリズムです。

## 主なユースケース（どういう時に使うか）

- スケジュールやタスク競合の解決
- 各種パズルやグラフの彩色問題の適合判定

## なぜ使うのか（メリット）

- 条件式を「ならば（$\implies$）」の有向グラフに変換し、強連結成分分解（SCC）を利用することで、変数と制約の数に対して $O(V + E)$ という線形時間で充足可能性の判定と具体的な解の構成を行えます。

---

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