# gcdlcm

最大公約数と最小公倍数を求めます。

## 関数

```go
func Gcd[T Integer](a, b T) T
func Lcm[T Integer](a, b T, integers ...T) T
```

## 使用例

```go
g := mylib.Gcd(12, 18)    // 6
l := mylib.Lcm(4, 6)      // 12
l2 := mylib.Lcm(2, 3, 4) // 12 (可変長引数)
```
