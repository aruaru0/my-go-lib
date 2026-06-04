# PriorityQueue

A generic priority queue implementation in Go using a binary heap.

## Import

```go
import "github.com/aruaru0/my-go-lib"
```

## Usage

### Min-Heap

```go
pq := mylib.NewPriorityQueue[int](func(a, b int) bool { return a < b })
pq.Push(3)
pq.Push(1)
pq.Push(2)

for pq.Len() > 0 {
    fmt.Println(pq.Pop()) // 1, 2, 3
}
```

### Max-Heap

```go
pq := mylib.NewPriorityQueue[int](func(a, b int) bool { return a > b })
pq.Push(1)
pq.Push(3)
pq.Push(2)

for pq.Len() > 0 {
    fmt.Println(pq.Pop()) // 3, 2, 1
}
```

### Custom Struct

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

Creates a new priority queue. The `less` function defines the ordering:
- Use `a < b` for a min-heap (smallest value has highest priority)
- Use `a > b` for a max-heap (largest value has highest priority)

### `(pq *PriorityQueue[T]) Push(x T)`

Pushes an element onto the heap.

### `(pq *PriorityQueue[T]) Pop() T`

Removes and returns the element with the highest priority (as defined by `less`).
Panics if the queue is empty.

### `(pq *PriorityQueue[T]) Len() int`

Returns the number of elements in the queue.
