# LIS

Computes the Longest Increasing Subsequence. Supports generics.

## Example

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	a := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(mylib.Lis(a)) // 4

	// LowerBound is also available
	b := []int{1, 3, 5, 7, 9}
	fmt.Println(mylib.LowerBound(b, 5)) // 2
}
```

## Function Signatures

```go
func LowerBound[T Ordered](a []T, x T) int
func Lis[T Ordered](a []T) int
```

## Algorithm

- O(n log n) LIS using binary search (strictly increasing)
- `LowerBound` returns the first position in a sorted slice where the value is >= x
