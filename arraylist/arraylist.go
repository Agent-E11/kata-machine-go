package arraylist

type ArrayList[T comparable] struct {
	Length int
}

func New[T comparable](capacity int) ArrayList[T] {
	return ArrayList[T]{} // TODO: Remove this
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
