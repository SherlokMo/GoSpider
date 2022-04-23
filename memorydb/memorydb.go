package memorydb

import "sync"

type MemDB[K Ikey, V any] struct {
	visited map[K]V
	mu      sync.RWMutex
}

func NewMemorydb[K Ikey, V any]() *MemDB[K, V] {
	return &MemDB[K, V]{
		visited: make(map[K]V),
	}
}

func (m *MemDB[K, V]) Store(key K, v V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.visited[key] = v
}

func (m *MemDB[K, T]) Exist(key K) bool {
	_, ok := m.visited[key]
	return ok
}
