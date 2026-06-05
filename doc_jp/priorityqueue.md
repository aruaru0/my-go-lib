# PriorityQueue

バイナリヒープを使用した Go のジェネリック優先度付きキューです。

## インポート

```go
import "github.com/aruaru0/my-go-lib"
```

## 使用方法

### 最小ヒープ

```go
pq := mylib.NewPriorityQueue[int](func(a, b int) bool { return a < b })
pq.Push(3)
pq.Push(1)
pq.Push(2)

for pq.Len() > 0 {
    fmt.Println(pq.Pop()) // 1, 2, 3
}
```

### 最大ヒープ

```go
pq := mylib.NewPriorityQueue[int](func(a, b int) bool { return a > b })
pq.Push(1)
pq.Push(3)
pq.Push(2)

for pq.Len() > 0 {
    fmt.Println(pq.Pop()) // 3, 2, 1
}
```

### カスタム構造体

```go
type Task struct {
    priority int
    name     string
}

pq := mylib.NewPriorityQueue[Task](func(a, b Task) bool {
    return a.priority < b.priority
})
```

## API

### `NewPriorityQueue[T any](less func(a, b T) bool) *PriorityQueue[T]`

新しい優先度付きキューを作成します。`less` 関数が順序を定義します:
- 最小ヒープには `a < b` を使用します (最小値が最高優先度)
- 最大ヒープには `a > b` を使用します (最大値が最高優先度)

### `(pq *PriorityQueue[T]) Push(x T)`

要素をヒープにプッシュします。

### `(pq *PriorityQueue[T]) Pop() T`

`less` で定義された最高優先度の要素を削除して返します。キューが空の場合はパニックします。

### `(pq *PriorityQueue[T]) Peek() T`

最高優先度の要素を削除せずに参照します。キューが空の場合はパニックします。

### `(pq *PriorityQueue[T]) Len() int`

キュー内の要素数を返します。
