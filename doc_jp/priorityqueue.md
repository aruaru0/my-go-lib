# PriorityQueue

要素を動的に追加・取り出しする際、常にもっとも優先度の高い（値が最小、または最大）要素が先頭から取り出されるように管理するデータ構造（ヒープ）です。

## 主なユースケース（どういう時に使うか）

- **ダイクストラ法（最短経路探索）やプリム法（最小全域木）の高速化**
- **大量データのソートや、動的なタスクの優先度順処理（タスクスケジューリング）**

## なぜ使うのか（メリット）

- Goの `container/heap` インターフェースを満たし、要素の挿入と最小（最大）値の取り出しをいずれも $O(\log N)$ で高速に処理できます。

---

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