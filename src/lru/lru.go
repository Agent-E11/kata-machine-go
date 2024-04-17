package lru

// Cache that implements the Least Recently Used cache replacement policy
//
// [Cache Replacement Policies]
//
// [Cache Replacement Policies]: https://en.wikipedia.org/wiki/Cache_replacement_policies
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
