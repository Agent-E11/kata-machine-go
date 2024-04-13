package lru

type LRU[K comparable, V any] struct {
	length int
	capacity int
}

func New[K comparable, V any](capacity int) LRU[K, V] {
	return LRU[K, V]{}
}

func (l *LRU[K, V]) Update(key K, value V) {}

func (l *LRU[K, V]) Get(key K) (value V, ok bool) {
	return
}
