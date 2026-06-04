# EulerPhi

Euler's totient function (φ function). Supports generics.

## Example

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

## Function Signatures

```go
func EulerPhi[T Integer](n T) T
```

## Algorithm

- Compute φ(n) = n × Π(1 - 1/p) for each prime factor p of n (O(√n))
