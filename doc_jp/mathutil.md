# mathutil

標準の `math` パッケージでは提供されていない、整数（`int`）用の絶対値、最大値、最小値、累乗、平方根、切り上げ・切り下げ除算などの算術演算を提供する便利ライブラリです。

## 主なユースケース（どういう時に使うか）

- 浮動小数点数へのキャストを避けて、整数のみで安全かつ高速に最大・最小判定や累乗計算を行いたい場合
- 競技プログラミングや、符号付き整数の切り捨て除算の挙動調整

## なぜ使うのか（メリット）

- 型キャストによる精度損失や速度低下を防ぎ、直感的かつタイプセーフな整数演算をインライン風のシンプルなコードで提供します。

---

## 関数

```go
func FloorSum(n, m, a, b int) int
func Crt(r, m []int) [2]int
func PowMod(x, n, m int) int
func SafeMod(x, d int) int
func InvMod(x, m int) int
func InvGcd(a, b int) [2]int
func PrimitiveRoot(m int) int
```

## FloorSum

`sum_{i=0}^{n-1} floor((a*i + b) / m)` を計算します。

## Crt

中国剰余定理です。`x ≡ r[i] (mod m[i])` を満たす `[r, m]` を返します。解がない場合は `[0, 0]` を返します。

## PowMod

二分累乗法: `x^n mod m` を計算します。

## SafeMod

非負の結果を保証する `x mod d` を返します。

## InvMod

`x` の `m` を法とするモジュラ逆数を返します (`gcd(x, m) = 1` が必要です)。

## InvGcd

`[g, inv]` を返します。ただし `g = gcd(a, b)` であり、`inv` は `a/g` の `b/g` を法とするモジュラ逆数です。

## PrimitiveRoot

`m` を法とする最小の原始根を求めます。