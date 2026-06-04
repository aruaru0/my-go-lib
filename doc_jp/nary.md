# nary

N 進数の列挙を行います。

## 関数

```go
func NaryNumbers(N, digits int) <-chan string
```

## NaryNumbers

最大 `digits` 桁のすべての N 進数を文字列として生成し、チャネルを通じてストリーミングします。
