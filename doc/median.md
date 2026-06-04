# median

Dynamic median data structure with add/remove operations.

## Types

```go
type Median struct { ... }
func NewMedian() *Median
func (m *Median) Add(x int)
func (m *Median) Remove(x int)
func (m *Median) Median() int
```

## Median

Maintains a multiset of integers and supports querying the median value.
Implemented with two heaps (max-heap for lower half, min-heap for upper half).
For even-sized sets, returns the floor of the average of the two middle values.
