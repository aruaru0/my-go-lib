# CombTable

A table of binomial coefficients nCk computed using Pascal's triangle.

## Example

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	ct := mylib.NewCombTable[int](10)
	fmt.Println(ct.NCk(5, 2))  // 10
	fmt.Println(ct.NCk(10, 5)) // 252
}
```

## Function Signatures

```go
func NewCombTable[T Integer](n int) *CombTable[T]
func (c *CombTable[T]) NCk(n, k int) T
```

## Algorithm

- Precomputation in O(n²) using Pascal's triangle
- Table size: (n+1) × (n+1)
- Returns 0 if k < 0, n < k, or n exceeds the table size
