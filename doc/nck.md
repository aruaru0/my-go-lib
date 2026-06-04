# NCR

Computes nCk, nPk, and nHk modulo a given value using precomputed factorials. The modulus is specified via the constructor.

## Example

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	nc := mylib.NewNCR(1000000007, 100)
	fmt.Println(nc.N(5, 2))  // 10 (nCk)
	fmt.Println(nc.P(5, 2))  // 20 (nPk)
	fmt.Println(nc.H(5, 2))  // 6  (nHk)
}
```

## Function Signatures

```go
func NewNCR(mod int, maxN int) *NCR
func (nc *NCR) N(n, k int) int
func (nc *NCR) P(n, k int) int
func (nc *NCR) H(n, k int) int
```

## Algorithm

- Precompute factorials and inverse factorials in O(maxN)
- Modular inverses via Fermat's little theorem
- Each query is O(1)
