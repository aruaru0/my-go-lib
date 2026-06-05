# EulerPhi

正の整数 $N$ に対して、 $1$ から $N$ までの整数のうち $N$ と互いに素なものの個数を求める関数（オイラーのトーシェント関数）です。

## 主なユースケース（どういう時に使うか）

- **暗号アルゴリズム（RSA暗号など）におけるキー生成や法 $N$ での逆元の存在判定**
- **余剰数論におけるフェルマーの小定理の一般化（オイラーの定理）の適用**

## なぜ使うのか（メリット）

- $N$ の素因数分解の過程を応用し、 $O(\sqrt{N})$ の時間計算量で正確にオイラーのφ関数の値を求められます。

---

## 使用例

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	fmt.Println(mylib.EulerPhi(12)) // 4
	fmt.Println(mylib.EulerPhi(7))  // 6
}
```

## 関数シグネチャ

```go
func EulerPhi[T Integer](n T) T
```

## アルゴリズム

- φ(n) = n × Π(1 - 1/p) を n の各素因数 p について計算します (O(√n))