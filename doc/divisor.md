# Divisors

Enumerates all divisors of an integer. Supports generics.

## Example

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	divs := mylib.Divisors(12)
	fmt.Println(divs) // [1 12 2 6 3 4] (unsorted)
}
```

## Function Signatures

```go
func Divisors[T Integer](n T) []T
```

## Algorithm

- Loop i from 1 to √n, collect i and n/i when n % i == 0 (O(√n))
- The result is not sorted
