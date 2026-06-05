# GCD / LCM

2つ以上の整数の最大公約数（GCD）および最小公倍数（LCM）を求めるアルゴリズムです。

## 主なユースケース（どういう時に使うか）

- 分数の約分や通分などの算術計算
- **周期的なイベント（例: 3日おきと5日おき）が同時に発生する周期（最小公倍数）の計算**
- 格子点上の線分が通過するグリッドの数の算出

## なぜ使うのか（メリット）

- ユークリッドの互除法を用いて、最悪でも $O(\log(\min(a, b)))$ という極めて少ないステップ数で正確な公約数・公倍数を計算できます。

---

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	// int
	fmt.Println(mylib.Gcd(12, 8))         // 4
	fmt.Println(mylib.Lcm(12, 8))        // 24
	fmt.Println(mylib.Lcm(2, 3, 4))       // 12 (可変引数)

	// int64
	fmt.Println(mylib.Gcd[int64](48, 18)) // 6
	fmt.Println(mylib.Lcm[int64](6, 8))   // 24

	// uint
	fmt.Println(mylib.Gcd[uint](36, 24)) // 12
}
```

## 関数シグネチャ

```go
func Gcd[T Integer](a, b T) T
func Lcm[T Integer](a, b T, integers ...T) T
```

`Integer` は `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr` およびそれらのエイリアスをサポートします。

## アルゴリズム

- **GCD**: ユークリッドの互除法 (再帰)
- **LCM**: 公式: `a / gcd(a, b) * b`（可変引数が指定された場合は、順次計算を連結します）