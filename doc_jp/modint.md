# modint

整数演算を行う際、代入や四則演算（加減乗除）を行うたびに自動的に特定の法 $P$ による剰余（Mod）を計算する、カプセル化された整数型です。

## 主なユースケース（どういう時に使うか）

- **組み合わせ数や動的計画法の遷移など、答えが非常に大きくなるため特定の法（10^9+7や998244353など）での出力を求められる計算**
- モジュロ算術のコードをすっきりと書きたい場合

## なぜ使うのか（メリット）

- 四則演算（特に引き算での負数処理や、割り算でのモジュラ逆数の適用）で発生しがちなバグを自動的に回避し、通常の整数と同様の直感的な記述で剰余演算を安全に行えます。

---

## 型

```go
type Modint struct { ... }
```

## 関数

```go
func NewModint(m int) *Modint
```

## メソッド

```go
func (m *Modint) Add(a, b int) int
func (m *Modint) Sub(a, b int) int
func (m *Modint) Mul(a, b int) int
func (m *Modint) Div(a, b int) int
func (m *Modint) Pow(p, n int) int
func (m *Modint) Inv(a int) int
func (m *Modint) MulMatrix(A, B [][]int) [][]int
func (m *Modint) PowMatrix(A [][]int, p int) [][]int
```

## 使用例

```go
m := mylib.NewModint(1000000007)
x := m.Add(10, 20)      // 30
y := m.Sub(10, 20)      // 1000000007-10
z := m.Mul(1000000, 1000000) // 999993007
inv2 := m.Inv(2)         // 2 のモジュラ逆数
d := m.Div(10, 2)        // 5
p := m.Pow(2, 10)        // 1024
```