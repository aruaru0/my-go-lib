package mylib

type OrderedMap[K comparable, V any] struct {
	keys   []K
	values map[K]V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{values: make(map[K]V)}
}

func (m *OrderedMap[K, V]) Set(key K, value V) {
	if _, ok := m.values[key]; !ok {
		m.keys = append(m.keys, key)
	}
	m.values[key] = value
}

func (m *OrderedMap[K, V]) Get(key K) (V, bool) {
	v, ok := m.values[key]
	return v, ok
}

func (m *OrderedMap[K, V]) Delete(key K) {
	if _, ok := m.values[key]; ok {
		delete(m.values, key)
		for i, k := range m.keys {
			if k == key {
				m.keys = append(m.keys[:i], m.keys[i+1:]...)
				break
			}
		}
	}
}

func (m *OrderedMap[K, V]) Keys() []K {
	res := make([]K, len(m.keys))
	copy(res, m.keys)
	return res
}

func (m *OrderedMap[K, V]) Len() int {
	return len(m.keys)
}
