# PrimeFactor

Performs prime factorization. Supports generics.

## Example

```go
package main

import (
	"fmt"
	"github.com/aruaru0/my-go-lib"
)

func main() {
	// Map form: prime factor -> exponent
	fmt.Println(mylib.PrimeFactorMap(12)) // map[2:2 3:1]

	// Slice form: list of prime factors
	fmt.Println(mylib.PrimeFactors(12)) // [2 2 3]
}
```

## Function Signatures

```go
func PrimeFactorMap[T Integer](n T) map[T]T
func PrimeFactors[T Integer](n T) []T
```

`Integer` supports `int`, `int8`, `int16`, `int32`, `int64`, and `uint` types.

## Algorithm

- Trial division (O(√n))
- Divide by 2 repeatedly, then divide by odd numbers starting from 3
