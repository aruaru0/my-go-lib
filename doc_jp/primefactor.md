# PrimeFactor

素因数分解を行います。ジェネリクスに対応しています。

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
