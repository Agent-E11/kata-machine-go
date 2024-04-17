package hashmap

// [Hash Table]
//
// [Hash Table]: https://en.wikipedia.org/wiki/Hash_table
type Map[K int | string, V any] struct {}

func New[K int | string, V any]() Map[K, V] {
	return Map[K, V]{}
}

func (m *Map[K, V]) Get(key K) (value V, ok bool) {
	return
}

func (m *Map[K, V]) Set(key K, value V) {}

func (m *Map[K, V]) Delete(key K) (value V, ok bool) {
	return
}

func (m *Map[K, V]) Size() int {
	return 0
}

