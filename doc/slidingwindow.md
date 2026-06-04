# Sliding Window

Computes sliding window minimum and maximum values in O(n). Supports generics.

## Example

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	a := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(mylib.SlideMin(a, 3)) // [-1 -3 -3 -3 3 3]
	fmt.Println(mylib.SlideMax(a, 3)) // [3 3 5 5 6 7]
}
```

## Function Signatures

```go
func SlideMin[T Ordered](a []T, k int) []T
func SlideMax[T Ordered](a []T, k int) []T
```

## Algorithm

- O(n) sliding window min/max using a deque
- Based on the implementation from "Ants Book" 4.4 P.300
