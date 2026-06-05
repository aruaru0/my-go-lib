# PrimeFactor

与えられた正の整数を素数の積の形に分解し、それぞれの素因数とその指数（個数）を求めるアルゴリズムです。

## 主なユースケース（どういう時に使うか）

- **数の約数の個数や約数の和の計算（素因数分解の結果から数式で算出）**
- 暗号理論における公開鍵の安全性評価や数論問題の解決

## なぜ使うのか（メリット）

- $1$ から $\sqrt{N}$ までの探索を行うことで、任意の整数 $N$ に対し $O(\sqrt{N})$ の時間計算量で正確に素因数分解を完了できます。

---

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	// マップ形式: 素因数 -> 指数
	fmt.Println(mylib.PrimeFactorMap(12)) // map[2:2 3:1]

	// スライス形式: 素因数のリスト
	fmt.Println(mylib.PrimeFactors(12)) // [2 2 3]
}
```

## 関数シグネチャ

```go
func PrimeFactorMap[T Integer](n T) map[T]T
func PrimeFactors[T Integer](n T) []T
```

`Integer` は `int`, `int8`, `int16`, `int32`, `int64`, `uint` 型をサポートします。

## アルゴリズム

- 試し割り (O(√n))
- 2 で繰り返し割った後、3 から奇数で割ります