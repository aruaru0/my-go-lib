# modint

mylib パッケージは、`Modint` 構造体による設定可能な法を用いたモジュラ演算を提供します。

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
