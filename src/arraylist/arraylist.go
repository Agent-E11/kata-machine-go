package arraylist

// [Dynamic Array]
//
// [Dynamic Array]: https://en.wikipedia.org/wiki/Dynamic_array
type ArrayList[T comparable] struct {
	length int
}

func New[T comparable](capacity int) ArrayList[T] {
	return ArrayList[T]{}
}

func (al ArrayList[T]) Length() int {
	return 0
}

func (al *ArrayList[T]) Prepend(item T) {}

func (al *ArrayList[T]) InsertAt(item T, idx int) {}

func (al *ArrayList[T]) Append(item T) {}

func (al *ArrayList[T]) Remove(item T) (value T, ok bool) {
	return
}

func (al *ArrayList[T]) Get(idx int) (value T, ok bool) {
	return
}

func (al *ArrayList[T]) RemoveAt(idx int) (value T, ok bool) {
	return
}
