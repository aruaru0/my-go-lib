# nary

N-ary number enumeration.

## Functions

```go
func NaryNumbers(N, digits int) <-chan string
```

## NaryNumbers

Generates all N-ary numbers with up to `digits` digits as strings,
streaming them through a channel.
