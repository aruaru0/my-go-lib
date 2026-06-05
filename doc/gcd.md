# GCD / LCM

Computes the Greatest Common Divisor (GCD) and Least Common Multiple (LCM). Supports generics.

## Example

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
	fmt.Println(mylib.Lcm(2, 3, 4))       // 12 (variadic)

	// int64
	fmt.Println(mylib.Gcd[int64](48, 18)) // 6
	fmt.Println(mylib.Lcm[int64](6, 8))   // 24

	// uint
	fmt.Println(mylib.Gcd[uint](36, 24)) // 12
}
```

## Function Signatures

```go
func Gcd[T Integer](a, b T) T
func Lcm[T Integer](a, b T, integers ...T) T
```

`Integer` supports `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`, and their aliases.

## Algorithm

- **GCD**: Euclidean algorithm (recursive)
- **LCM**: Formula: `a / gcd(a, b) * b` (chained sequentially for variadic inputs)
