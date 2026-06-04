# gcdlcm

Greatest common divisor and least common multiple.

## Functions

```go
func Gcd[T Integer](a, b T) T
func Lcm[T Integer](a, b T, integers ...T) T
```

## Example

```go
g := mylib.Gcd(12, 18)    // 6
l := mylib.Lcm(4, 6)      // 12
l2 := mylib.Lcm(2, 3, 4) // 12 (variadic)
```
