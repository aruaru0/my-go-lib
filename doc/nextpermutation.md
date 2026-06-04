# NextPermutation

Generates the next lexicographic permutation. Supports generics.

## Example

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	a := []int{1, 2, 3}
	for {
		fmt.Println(a)
		if !mylib.NextPermutation(a) {
			break
		}
	}
}
```

## Function Signatures

```go
func NextPermutation[T Ordered](x []T) bool
```

`Ordered` supports all comparable basic types.

## Algorithm

- Same algorithm as C++ `std::next_permutation` (O(n))
- Returns `true` if the next permutation exists, `false` otherwise
